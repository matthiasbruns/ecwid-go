// Package webhook builds, signs, delivers, and classifies Ecwid webhooks for the
// mock's trigger control API and UI.
//
// Every event is composed with [webhooks.Event] and signed with [webhooks.Sign],
// so the mock exercises the exact wire format, entityId typing, and signature a
// real integration receives — running the same code users run is the whole point
// of the tool.
package webhook

import (
	"encoding/json"
	"strings"

	"github.com/matthiasbruns/ecwid-go/ecwid/webhooks"
)

// entityId wire types. Ecwid is not consistent: application.* events send a
// quoted string while every other family sends a bare JSON number. The mock
// reproduces that quirk so tooling built against it hits the same bug class real
// integrations do.
const (
	// EntityIDNumber marks a family whose entityId is a bare JSON number.
	EntityIDNumber = "number"
	// EntityIDString marks a family (application.*) whose entityId is a quoted
	// JSON string.
	EntityIDString = "string"
)

// applicationPrefix is the event family whose entityId serializes as a string.
const applicationPrefix = "application."

// EventSpec is the fixture for one event type: the default entityId and data
// payload used when the caller supplies neither, so the UI's Fire button works
// with zero input.
type EventSpec struct {
	// Type is the event, e.g. [webhooks.EventOrderCreated].
	Type webhooks.EventType
	// EntityID is the default entityId in string form. It serializes as a JSON
	// number for every family except application.*, which serializes a quoted
	// string; see [EventSpec.EntityIDType].
	EntityID string
	// Data is the default data payload, or nil for the families that carry no
	// data key at all (product.*, category.*, discount_coupon.*,
	// application.installed/uninstalled, and the class, group, and GDPR events).
	Data json.RawMessage
}

// Group is the event's display group: the entity prefix before the first dot,
// e.g. "order" for order.created, "customer_group" for customer_group.updated.
func (s EventSpec) Group() string {
	prefix, _, _ := strings.Cut(string(s.Type), ".")
	return prefix
}

// EntityIDType reports whether the event serializes entityId as a JSON number
// ([EntityIDNumber]) or a quoted string ([EntityIDString]).
func (s EventSpec) EntityIDType() string {
	if strings.HasPrefix(string(s.Type), applicationPrefix) {
		return EntityIDString
	}
	return EntityIDNumber
}

// HasData reports whether the event carries a data key.
func (s EventSpec) HasData() bool {
	return len(s.Data) > 0
}

// catalog is the fixture for all 42 Ecwid webhook events, ordered by family so
// the UI dropdown groups sensibly. Data shapes follow ecwid/webhooks; families
// documented to carry no data are left with a nil Data and emit no data key.
var catalog = []EventSpec{
	// Store profile. entityId is the numeric store ID.
	{Type: webhooks.EventProfileUpdated, EntityID: "1003"},
	{Type: webhooks.EventProfileSubscriptionStatusChanged, EntityID: "1003", Data: json.RawMessage(`{"oldSubscriptionName":"ECWID_BUSINESS","newSubscriptionName":"ECWID_UNLIMITED"}`)},
	{Type: webhooks.EventProfilePersonalDataRemovalReq, EntityID: "1003"},
	{Type: webhooks.EventProfilePersonalDataExportReq, EntityID: "1003"},

	// Orders. entityId is the internal order ID (a number); the human-readable ID
	// is data.orderId.
	{Type: webhooks.EventOrderCreated, EntityID: "103878161", Data: json.RawMessage(`{"orderId":"XJ12H","newPaymentStatus":"PAID","newFulfillmentStatus":"AWAITING_PROCESSING"}`)},
	{Type: webhooks.EventOrderUpdated, EntityID: "103878161", Data: json.RawMessage(`{"orderId":"XJ12H","oldPaymentStatus":"AWAITING_PAYMENT","newPaymentStatus":"PAID","oldFulfillmentStatus":"AWAITING_PROCESSING","newFulfillmentStatus":"PROCESSING"}`)},
	{Type: webhooks.EventOrderDeleted, EntityID: "103878161", Data: json.RawMessage(`{"orderId":"XJ12H"}`)},

	// Abandoned carts.
	{Type: webhooks.EventUnfinishedOrderCreated, EntityID: "103878162", Data: json.RawMessage(`{"orderId":"RA1SD","cartId":"a1b2c3d4-0000-4000-8000-000000000000"}`)},
	{Type: webhooks.EventUnfinishedOrderUpdated, EntityID: "103878162", Data: json.RawMessage(`{"orderId":"RA1SD","cartId":"a1b2c3d4-0000-4000-8000-000000000000"}`)},
	{Type: webhooks.EventUnfinishedOrderDeleted, EntityID: "103878162", Data: json.RawMessage(`{"orderId":"RA1SD","cartId":"a1b2c3d4-0000-4000-8000-000000000000"}`)},

	// Products. No data key.
	{Type: webhooks.EventProductCreated, EntityID: "12345678"},
	{Type: webhooks.EventProductUpdated, EntityID: "12345678"},
	{Type: webhooks.EventProductDeleted, EntityID: "12345678"},

	// Categories. No data key.
	{Type: webhooks.EventCategoryCreated, EntityID: "8901234"},
	{Type: webhooks.EventCategoryUpdated, EntityID: "8901234"},
	{Type: webhooks.EventCategoryDeleted, EntityID: "8901234"},

	// Product types (classes). No data key.
	{Type: webhooks.EventProductClassCreated, EntityID: "1000001"},
	{Type: webhooks.EventProductClassUpdated, EntityID: "1000001"},
	{Type: webhooks.EventProductClassDeleted, EntityID: "1000001"},

	// Customers.
	{Type: webhooks.EventCustomerCreated, EntityID: "5001", Data: json.RawMessage(`{"customerEmail":"jane@example.com"}`)},
	{Type: webhooks.EventCustomerUpdated, EntityID: "5001", Data: json.RawMessage(`{"customerEmail":"jane@example.com"}`)},
	{Type: webhooks.EventCustomerDeleted, EntityID: "5001", Data: json.RawMessage(`{"customerEmail":"jane@example.com"}`)},

	// Customer GDPR. No data key; the customer is identified by entityId.
	{Type: webhooks.EventCustomerPersonalDataRemovalReq, EntityID: "5001"},
	{Type: webhooks.EventCustomerPersonalDataExportReq, EntityID: "5001"},

	// Customer groups. No data key.
	{Type: webhooks.EventCustomerGroupCreated, EntityID: "10"},
	{Type: webhooks.EventCustomerGroupUpdated, EntityID: "10"},
	{Type: webhooks.EventCustomerGroupDeleted, EntityID: "10"},

	// Discount coupons. No data key.
	{Type: webhooks.EventDiscountCouponCreated, EntityID: "701"},
	{Type: webhooks.EventDiscountCouponUpdated, EntityID: "701"},
	{Type: webhooks.EventDiscountCouponDeleted, EntityID: "701"},

	// Promotions (advanced discounts).
	{Type: webhooks.EventPromotionCreated, EntityID: "301", Data: json.RawMessage(`{"version":1}`)},
	{Type: webhooks.EventPromotionUpdated, EntityID: "301", Data: json.RawMessage(`{"version":2}`)},
	{Type: webhooks.EventPromotionDeleted, EntityID: "301", Data: json.RawMessage(`{"version":3}`)},

	// Product reviews. status is absent on review.deleted.
	{Type: webhooks.EventReviewCreated, EntityID: "9001", Data: json.RawMessage(`{"productId":12345678,"orderId":103878161,"status":"MODERATED"}`)},
	{Type: webhooks.EventReviewUpdated, EntityID: "9001", Data: json.RawMessage(`{"productId":12345678,"orderId":103878161,"status":"PUBLISHED"}`)},
	{Type: webhooks.EventReviewDeleted, EntityID: "9001", Data: json.RawMessage(`{"productId":12345678,"orderId":103878161}`)},

	// Order tax invoices. There is no invoice.updated event.
	{Type: webhooks.EventInvoiceCreated, EntityID: "103878161", Data: json.RawMessage(`{"orderId":"GH781"}`)},
	{Type: webhooks.EventInvoiceDeleted, EntityID: "103878161", Data: json.RawMessage(`{"orderId":"GH781"}`)},

	// Applications. entityId is a quoted STRING, not a number. installed and
	// uninstalled carry no data.
	{Type: webhooks.EventApplicationInstalled, EntityID: "1003"},
	{Type: webhooks.EventApplicationUninstalled, EntityID: "1003"},
	{Type: webhooks.EventApplicationSubscriptionStatusChanged, EntityID: "1003", Data: json.RawMessage(`{"oldSubscriptionStatus":"TRIAL","newSubscriptionStatus":"ACTIVE"}`)},
	{Type: webhooks.EventApplicationStorageChanged, EntityID: "1003", Data: json.RawMessage(`{"key":"config"}`)},
}

// byType indexes the catalog for O(1) lookup.
var byType = func() map[webhooks.EventType]EventSpec {
	m := make(map[webhooks.EventType]EventSpec, len(catalog))
	for _, s := range catalog {
		m[s.Type] = s
	}
	return m
}()

// Lookup returns the fixture for an event type, and whether it is a known event.
func Lookup(t webhooks.EventType) (EventSpec, bool) {
	s, ok := byType[t]
	return s, ok
}

// Catalog returns a copy of every event fixture, ordered by family. The copy
// keeps callers from mutating the shared fixtures.
func Catalog() []EventSpec {
	out := make([]EventSpec, len(catalog))
	copy(out, catalog)
	return out
}
