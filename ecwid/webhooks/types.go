package webhooks

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"
)

// EventType identifies what happened in the store, in "entity.action" form.
type EventType string

// Store profile events. Require the read_store_profile scope.
const (
	EventProfileUpdated                   EventType = "profile.updated"
	EventProfileSubscriptionStatusChanged EventType = "profile.subscriptionStatusChanged"
	EventProfilePersonalDataRemovalReq    EventType = "profile.personalDataRemovalRequest"
	EventProfilePersonalDataExportReq     EventType = "profile.personalDataExportRequest"
)

// Order events. Require the read_orders scope.
const (
	EventOrderCreated EventType = "order.created"
	EventOrderUpdated EventType = "order.updated"
	EventOrderDeleted EventType = "order.deleted"
)

// Abandoned cart events. Require the read_orders scope.
const (
	EventUnfinishedOrderCreated EventType = "unfinished_order.created"
	EventUnfinishedOrderUpdated EventType = "unfinished_order.updated"
	EventUnfinishedOrderDeleted EventType = "unfinished_order.deleted"
)

// Product events. Require the read_catalog scope.
const (
	EventProductCreated EventType = "product.created"
	EventProductUpdated EventType = "product.updated"
	EventProductDeleted EventType = "product.deleted"
)

// Category events. Require the read_catalog scope.
const (
	EventCategoryCreated EventType = "category.created"
	EventCategoryUpdated EventType = "category.updated"
	EventCategoryDeleted EventType = "category.deleted"
)

// Product type (product class) events. Require the read_catalog scope.
const (
	EventProductClassCreated EventType = "product_class.created"
	EventProductClassUpdated EventType = "product_class.updated"
	EventProductClassDeleted EventType = "product_class.deleted"
)

// Customer events. Require the read_customers scope.
const (
	EventCustomerCreated EventType = "customer.created"
	EventCustomerUpdated EventType = "customer.updated"
	EventCustomerDeleted EventType = "customer.deleted"
)

// Customer GDPR events. These require the read_store_profile scope, NOT
// read_customers as the neighbouring customer.* events do.
const (
	EventCustomerPersonalDataRemovalReq EventType = "customer.personalDataRemovalRequest"
	EventCustomerPersonalDataExportReq  EventType = "customer.personalDataExportRequest"
)

// Customer group events. Require the read_customers scope.
const (
	EventCustomerGroupCreated EventType = "customer_group.created"
	EventCustomerGroupUpdated EventType = "customer_group.updated"
	EventCustomerGroupDeleted EventType = "customer_group.deleted"
)

// Discount coupon events. Require the read_discount_coupons scope.
const (
	EventDiscountCouponCreated EventType = "discount_coupon.created"
	EventDiscountCouponUpdated EventType = "discount_coupon.updated"
	EventDiscountCouponDeleted EventType = "discount_coupon.deleted"
)

// Promotion (advanced discount) events. Require the read_promotion scope.
const (
	EventPromotionCreated EventType = "promotion.created"
	EventPromotionUpdated EventType = "promotion.updated"
	EventPromotionDeleted EventType = "promotion.deleted"
)

// Product review events. Require the read_reviews scope.
const (
	EventReviewCreated EventType = "review.created"
	EventReviewUpdated EventType = "review.updated"
	EventReviewDeleted EventType = "review.deleted"
)

// Order tax invoice events. Require the read_invoices scope. There is no
// invoice.updated event.
const (
	EventInvoiceCreated EventType = "invoice.created"
	EventInvoiceDeleted EventType = "invoice.deleted"
)

// Application events. Require the read_store_profile scope. They are delivered
// only to the app they concern.
const (
	EventApplicationInstalled                 EventType = "application.installed"
	EventApplicationUninstalled               EventType = "application.uninstalled"
	EventApplicationSubscriptionStatusChanged EventType = "application.subscriptionStatusChanged"
	EventApplicationStorageChanged            EventType = "application.storageChanged"
)

// applicationPrefix marks the event family whose entityId is sent as a quoted
// JSON string rather than a number.
const applicationPrefix = "application."

// Event is a webhook delivered by Ecwid.
//
// Only EventCreated and EventID are covered by the signature; every other field
// is unauthenticated. See the package documentation.
type Event struct {
	// EventID is a UUID, unique per webhook. Use it to dedupe replays and retries.
	EventID string `json:"eventId"`
	// EventCreated is the UNIX timestamp, in seconds, of when the event happened.
	EventCreated int64 `json:"eventCreated"`
	// StoreID is the Ecwid store ID.
	StoreID int64 `json:"storeId"`
	// EventType is the entity and action, e.g. [EventOrderCreated].
	EventType EventType `json:"eventType"`
	// EntityID identifies the affected item: order, product, customer, etc.
	// Use it to re-fetch the entity over the REST API.
	//
	// It is normalized to a string because Ecwid does not type it consistently
	// on the wire: order and product events send a JSON number
	// (`"entityId":103878161`) while application.* events send a quoted string
	// (`"entityId":"1003"`). Both decode into this field.
	//
	// For order events this is the *internal* order ID, which the REST API's
	// order endpoints do not accept. The human-readable order ID is
	// [OrderData.OrderID].
	EntityID string `json:"entityId"`
	// Data carries extra details for the events that have them, and is absent
	// entirely for those that do not (product.*, category.*, discount_coupon.*,
	// application.installed, application.uninstalled). Decode it with the
	// accessor for the event family, e.g. [Event.OrderData].
	Data json.RawMessage `json:"data,omitempty"`
}

// UnmarshalJSON decodes an Ecwid webhook body, accepting entityId as either a
// JSON number or a quoted string and normalizing it to [Event.EntityID].
func (e *Event) UnmarshalJSON(b []byte) error {
	// alias sheds the methods on Event, so json does not recurse into this one.
	type alias Event
	aux := struct {
		EntityID json.RawMessage `json:"entityId"`
		*alias
	}{alias: (*alias)(e)}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	raw := bytes.TrimSpace(aux.EntityID)
	switch {
	case len(raw) == 0, bytes.Equal(raw, []byte("null")):
		e.EntityID = ""
	case raw[0] == '"':
		if err := json.Unmarshal(raw, &e.EntityID); err != nil {
			return fmt.Errorf("webhooks: decode entityId: %w", err)
		}
	case isJSONNumber(string(raw)):
		e.EntityID = string(raw)
	default:
		return fmt.Errorf("webhooks: entityId must be a JSON string or number, got %s", raw)
	}
	return nil
}

// MarshalJSON encodes the event in Ecwid's wire format. It reproduces the
// entityId typing quirk that [Event.UnmarshalJSON] absorbs: application.* events
// get a quoted string, everything else a bare JSON number.
func (e Event) MarshalJSON() ([]byte, error) {
	entityID, err := json.Marshal(e.EntityID)
	if err != nil {
		return nil, fmt.Errorf("webhooks: encode entityId: %w", err)
	}
	if !strings.HasPrefix(string(e.EventType), applicationPrefix) && isJSONNumber(e.EntityID) {
		entityID = []byte(e.EntityID)
	}

	type alias Event
	return json.Marshal(struct {
		EntityID json.RawMessage `json:"entityId"`
		*alias
	}{EntityID: entityID, alias: (*alias)(&e)})
}

// isJSONNumber reports whether s is exactly one JSON number literal.
func isJSONNumber(s string) bool {
	dec := json.NewDecoder(strings.NewReader(s))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false
	}
	if _, ok := tok.(json.Number); !ok {
		return false
	}
	// Reject trailing junk, e.g. "1 2".
	_, err = dec.Token()
	return errors.Is(err, io.EOF)
}

// OrderData is the data payload of order.* events.
type OrderData struct {
	// OrderID is the human-readable order ID, e.g. "XJ12H" — the one the REST
	// API's order endpoints accept, unlike [Event.EntityID].
	OrderID string `json:"orderId"`
	// OldPaymentStatus is absent on order.created and order.deleted.
	OldPaymentStatus string `json:"oldPaymentStatus,omitempty"`
	// NewPaymentStatus is absent on order.deleted.
	NewPaymentStatus string `json:"newPaymentStatus,omitempty"`
	// OldFulfillmentStatus is absent on order.created and order.deleted.
	OldFulfillmentStatus string `json:"oldFulfillmentStatus,omitempty"`
	// NewFulfillmentStatus is absent on order.deleted.
	NewFulfillmentStatus string `json:"newFulfillmentStatus,omitempty"`
}

// UnfinishedOrderData is the data payload of unfinished_order.* events.
type UnfinishedOrderData struct {
	// OrderID is the human-readable order ID, e.g. "RA1SD".
	OrderID string `json:"orderId"`
	// CartID is the UUID of the abandoned cart.
	CartID string `json:"cartId"`
}

// ReviewData is the data payload of review.* events.
type ReviewData struct {
	// ProductID is the reviewed product.
	ProductID int64 `json:"productId"`
	// OrderID is the internal order ID as a number — unlike [OrderData.OrderID]
	// and [InvoiceData.OrderID], which are the human-readable string form.
	OrderID int64 `json:"orderId"`
	// Status is "MODERATED" or "PUBLISHED", and is absent on review.deleted.
	Status string `json:"status,omitempty"`
}

// CustomerData is the data payload of customer.* events.
type CustomerData struct {
	CustomerEmail string `json:"customerEmail"`
}

// PromotionData is the data payload of promotion.* events.
type PromotionData struct {
	// Version is Ecwid's internal iterated version of the discount.
	Version int `json:"version"`
}

// InvoiceData is the data payload of invoice.* events.
type InvoiceData struct {
	// OrderID is the human-readable ID of the invoiced order, e.g. "GH781".
	OrderID string `json:"orderId"`
}

// AppSubscriptionData is the data payload of application.subscriptionStatusChanged.
//
// Note the asymmetry with the similarly named
// [ProfileSubscriptionData]: this event reports a subscription *status*
// (e.g. "TRIAL" to "ACTIVE"), that one a subscription *name*.
type AppSubscriptionData struct {
	OldSubscriptionStatus string `json:"oldSubscriptionStatus"`
	NewSubscriptionStatus string `json:"newSubscriptionStatus"`
}

// ProfileSubscriptionData is the data payload of profile.subscriptionStatusChanged.
//
// Note the asymmetry with the similarly named [AppSubscriptionData]: this event
// reports a subscription *name* (e.g. "ECWID_BUSINESS" to "ECWID_UNLIMITED"),
// that one a subscription *status*.
type ProfileSubscriptionData struct {
	OldSubscriptionName string `json:"oldSubscriptionName"`
	NewSubscriptionName string `json:"newSubscriptionName"`
}

// OrderData decodes the payload of an order.* event.
func (e Event) OrderData() (OrderData, error) {
	var d OrderData
	err := e.decodeData(&d, EventOrderCreated, EventOrderUpdated, EventOrderDeleted)
	return d, err
}

// UnfinishedOrderData decodes the payload of an unfinished_order.* event.
func (e Event) UnfinishedOrderData() (UnfinishedOrderData, error) {
	var d UnfinishedOrderData
	err := e.decodeData(&d, EventUnfinishedOrderCreated, EventUnfinishedOrderUpdated, EventUnfinishedOrderDeleted)
	return d, err
}

// ReviewData decodes the payload of a review.* event.
func (e Event) ReviewData() (ReviewData, error) {
	var d ReviewData
	err := e.decodeData(&d, EventReviewCreated, EventReviewUpdated, EventReviewDeleted)
	return d, err
}

// CustomerData decodes the payload of a customer.* event.
func (e Event) CustomerData() (CustomerData, error) {
	var d CustomerData
	err := e.decodeData(&d, EventCustomerCreated, EventCustomerUpdated, EventCustomerDeleted)
	return d, err
}

// PromotionData decodes the payload of a promotion.* event.
func (e Event) PromotionData() (PromotionData, error) {
	var d PromotionData
	err := e.decodeData(&d, EventPromotionCreated, EventPromotionUpdated, EventPromotionDeleted)
	return d, err
}

// InvoiceData decodes the payload of an invoice.* event.
func (e Event) InvoiceData() (InvoiceData, error) {
	var d InvoiceData
	err := e.decodeData(&d, EventInvoiceCreated, EventInvoiceDeleted)
	return d, err
}

// AppSubscriptionData decodes the payload of application.subscriptionStatusChanged.
func (e Event) AppSubscriptionData() (AppSubscriptionData, error) {
	var d AppSubscriptionData
	err := e.decodeData(&d, EventApplicationSubscriptionStatusChanged)
	return d, err
}

// ProfileSubscriptionData decodes the payload of profile.subscriptionStatusChanged.
func (e Event) ProfileSubscriptionData() (ProfileSubscriptionData, error) {
	var d ProfileSubscriptionData
	err := e.decodeData(&d, EventProfileSubscriptionStatusChanged)
	return d, err
}

// decodeData unmarshals Data into dst, refusing event types that do not carry
// that payload so a caller cannot silently read a zero value off the wrong event.
func (e Event) decodeData(dst any, allowed ...EventType) error {
	if !slices.Contains(allowed, e.EventType) {
		return fmt.Errorf("webhooks: event type %q carries no such data (want one of %v)", e.EventType, allowed)
	}
	if len(e.Data) == 0 || bytes.Equal(bytes.TrimSpace(e.Data), []byte("null")) {
		return fmt.Errorf("webhooks: event %q has no data", e.EventType)
	}
	if err := json.Unmarshal(e.Data, dst); err != nil {
		return fmt.Errorf("webhooks: decode %q data: %w", e.EventType, err)
	}
	return nil
}
