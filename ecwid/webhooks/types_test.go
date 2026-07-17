package webhooks

import (
	"encoding/json"
	"strings"
	"testing"
)

// allEventTypes is every event this package defines. Payloads below are the
// documented examples with Ecwid's JavaScript-style comments and trailing commas
// removed — their docs' snippets are not valid JSON as printed.
var allEventTypes = []EventType{
	EventProfileUpdated,
	EventProfileSubscriptionStatusChanged,
	EventProfilePersonalDataRemovalReq,
	EventProfilePersonalDataExportReq,
	EventOrderCreated,
	EventOrderUpdated,
	EventOrderDeleted,
	EventUnfinishedOrderCreated,
	EventUnfinishedOrderUpdated,
	EventUnfinishedOrderDeleted,
	EventProductCreated,
	EventProductUpdated,
	EventProductDeleted,
	EventCategoryCreated,
	EventCategoryUpdated,
	EventCategoryDeleted,
	EventProductClassCreated,
	EventProductClassUpdated,
	EventProductClassDeleted,
	EventCustomerCreated,
	EventCustomerUpdated,
	EventCustomerDeleted,
	EventCustomerPersonalDataRemovalReq,
	EventCustomerPersonalDataExportReq,
	EventCustomerGroupCreated,
	EventCustomerGroupUpdated,
	EventCustomerGroupDeleted,
	EventDiscountCouponCreated,
	EventDiscountCouponUpdated,
	EventDiscountCouponDeleted,
	EventPromotionCreated,
	EventPromotionUpdated,
	EventPromotionDeleted,
	EventReviewCreated,
	EventReviewUpdated,
	EventReviewDeleted,
	EventInvoiceCreated,
	EventInvoiceDeleted,
	EventApplicationInstalled,
	EventApplicationUninstalled,
	EventApplicationSubscriptionStatusChanged,
	EventApplicationStorageChanged,
}

func TestEventTypes_AllDefined(t *testing.T) {
	const want = 42
	if len(allEventTypes) != want {
		t.Errorf("allEventTypes has %d entries, want %d", len(allEventTypes), want)
	}

	seen := make(map[EventType]bool, len(allEventTypes))
	for _, et := range allEventTypes {
		if seen[et] {
			t.Errorf("duplicate event type %q", et)
		}
		seen[et] = true

		entity, action, ok := strings.Cut(string(et), ".")
		if !ok || entity == "" || action == "" {
			t.Errorf("event type %q is not in entity.action form", et)
		}
	}

	// Ecwid publishes no invoice.updated, and inventing one would have callers
	// switch on an event that never arrives.
	if seen["invoice.updated"] {
		t.Error("invoice.updated is defined, but Ecwid has no such event")
	}
}

func TestEvent_UnmarshalEntityID(t *testing.T) {
	tests := []struct {
		name         string
		body         string
		wantEntityID string
		wantType     EventType
	}{
		{
			// Order and product events send entityId as a bare JSON number.
			name: "numeric entityId (order.created)",
			body: `{
				"eventId":"80aece08-40e8-4145-8764-6c2f0d38678",
				"eventCreated":1234567,
				"storeId":1003,
				"entityId":103878161,
				"eventType":"order.created",
				"data":{"orderId":"XJ12H","newPaymentStatus":"PAID","newFulfillmentStatus":"SHIPPED"}
			}`,
			wantEntityID: "103878161",
			wantType:     EventOrderCreated,
		},
		{
			// application.* events send the same field as a quoted string.
			name: "string entityId (application.installed)",
			body: `{
				"eventId":"80aece08-40e8-4145-8764-6c2f0d38678",
				"eventCreated":1234567,
				"storeId":1003,
				"entityId":"1003",
				"eventType":"application.installed"
			}`,
			wantEntityID: "1003",
			wantType:     EventApplicationInstalled,
		},
		{
			name: "entityId exceeding float64 precision keeps every digit",
			body: `{"eventId":"e","eventCreated":1,"storeId":1003,"entityId":9007199254740993,"eventType":"product.updated"}`,
			// Decoding via float64 would round this to ...92.
			wantEntityID: "9007199254740993",
			wantType:     EventProductUpdated,
		},
		{
			name:         "absent entityId",
			body:         `{"eventId":"e","eventCreated":1,"storeId":1003,"eventType":"product.updated"}`,
			wantEntityID: "",
			wantType:     EventProductUpdated,
		},
		{
			name:         "null entityId",
			body:         `{"eventId":"e","eventCreated":1,"storeId":1003,"entityId":null,"eventType":"product.updated"}`,
			wantEntityID: "",
			wantType:     EventProductUpdated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var e Event
			if err := json.Unmarshal([]byte(tt.body), &e); err != nil {
				t.Fatalf("Unmarshal() = %v, want nil", err)
			}
			if e.EntityID != tt.wantEntityID {
				t.Errorf("EntityID = %q, want %q", e.EntityID, tt.wantEntityID)
			}
			if e.EventType != tt.wantType {
				t.Errorf("EventType = %q, want %q", e.EventType, tt.wantType)
			}
		})
	}
}

func TestEvent_UnmarshalEntityID_Invalid(t *testing.T) {
	for _, body := range []string{
		`{"eventId":"e","entityId":true,"eventType":"product.updated"}`,
		`{"eventId":"e","entityId":{"id":1},"eventType":"product.updated"}`,
		`{"eventId":"e","entityId":[1],"eventType":"product.updated"}`,
	} {
		var e Event
		if err := json.Unmarshal([]byte(body), &e); err == nil {
			t.Errorf("Unmarshal(%s) = nil, want an error", body)
		}
	}
}

func TestEvent_UnmarshalCommonFields(t *testing.T) {
	const body = `{
		"eventId":"95ed494f-362d-4fca-933d-f5c7064b04bb",
		"eventCreated":1634209655,
		"storeId":1003,
		"entityId":293225803,
		"eventType":"unfinished_order.created",
		"data":{"orderId":"RA1SD","cartId":"320AB5DE-5EA6-4419-A4CF-0B04D2913190"}
	}`

	var e Event
	if err := json.Unmarshal([]byte(body), &e); err != nil {
		t.Fatalf("Unmarshal() = %v, want nil", err)
	}
	if e.EventID != "95ed494f-362d-4fca-933d-f5c7064b04bb" {
		t.Errorf("EventID = %q", e.EventID)
	}
	if e.EventCreated != 1634209655 {
		t.Errorf("EventCreated = %d, want 1634209655", e.EventCreated)
	}
	if e.StoreID != 1003 {
		t.Errorf("StoreID = %d, want 1003", e.StoreID)
	}
}

// product.*, category.* and discount_coupon.* omit the data key entirely.
func TestEvent_NoDataKey(t *testing.T) {
	for _, body := range []string{
		`{"eventId":"08a78904-4c1a-0aa0-953a-2e33c56236f1","eventCreated":1469429915,"storeId":1003,"entityId":667251253,"eventType":"product.created"}`,
		`{"eventId":"08a78904-0aa0-4c1a-953a-2e33c56236f0","eventCreated":1469429912,"storeId":1003,"entityId":66722483,"eventType":"category.created"}`,
		`{"eventId":"80aece08-40e8-4145-8764-6c2f0d38678","eventCreated":7891234567,"storeId":1003,"entityId":12345678,"eventType":"discount_coupon.created"}`,
		`{"eventId":"80aece08-40e8-4145-8764-6c2f0d38678","eventCreated":1234567,"storeId":1003,"entityId":"1003","eventType":"application.uninstalled"}`,
	} {
		var e Event
		if err := json.Unmarshal([]byte(body), &e); err != nil {
			t.Errorf("Unmarshal(%s) = %v, want nil", body, err)
			continue
		}
		if e.Data != nil {
			t.Errorf("Data = %s, want nil for %s", e.Data, e.EventType)
		}
	}
}

func TestEvent_DataAccessors(t *testing.T) {
	// Each body is the documented example for that event.
	t.Run("order.created", func(t *testing.T) {
		e := mustUnmarshal(t, `{"eventId":"80aece08-40e8-4145-8764-6c2f0d38678","eventCreated":1234567,"storeId":1003,"entityId":103878161,"eventType":"order.created","data":{"orderId":"XJ12H","newPaymentStatus":"PAID","newFulfillmentStatus":"SHIPPED"}}`)
		d, err := e.OrderData()
		if err != nil {
			t.Fatalf("OrderData() = %v", err)
		}
		want := OrderData{OrderID: "XJ12H", NewPaymentStatus: "PAID", NewFulfillmentStatus: "SHIPPED"}
		if d != want {
			t.Errorf("OrderData() = %+v, want %+v", d, want)
		}
		// entityId is the internal ID; the usable one lives in data.
		if e.EntityID == d.OrderID {
			t.Error("EntityID and data.orderId should differ for order events")
		}
	})

	t.Run("order.updated", func(t *testing.T) {
		e := mustUnmarshal(t, `{"eventId":"80aece08-40e8-4145-8764-6c2f0d38678","eventCreated":1234567,"storeId":1003,"entityId":450012387,"eventType":"order.updated","data":{"orderId":"B8HGD","oldPaymentStatus":"PAID","newPaymentStatus":"PAID","oldFulfillmentStatus":"PROCESSING","newFulfillmentStatus":"SHIPPED"}}`)
		d, err := e.OrderData()
		if err != nil {
			t.Fatalf("OrderData() = %v", err)
		}
		want := OrderData{
			OrderID:              "B8HGD",
			OldPaymentStatus:     "PAID",
			NewPaymentStatus:     "PAID",
			OldFulfillmentStatus: "PROCESSING",
			NewFulfillmentStatus: "SHIPPED",
		}
		if d != want {
			t.Errorf("OrderData() = %+v, want %+v", d, want)
		}
	})

	t.Run("order.deleted", func(t *testing.T) {
		e := mustUnmarshal(t, `{"eventId":"80aece08-40e8-4145-8764-6c2f0d38678","eventCreated":1234567,"storeId":1003,"entityId":103878161,"eventType":"order.deleted","data":{"orderId":"XJ12H"}}`)
		d, err := e.OrderData()
		if err != nil {
			t.Fatalf("OrderData() = %v", err)
		}
		if want := (OrderData{OrderID: "XJ12H"}); d != want {
			t.Errorf("OrderData() = %+v, want %+v", d, want)
		}
	})

	t.Run("unfinished_order.created", func(t *testing.T) {
		e := mustUnmarshal(t, `{"eventId":"95ed494f-362d-4fca-933d-f5c7064b04bb","eventCreated":1634209655,"storeId":1003,"entityId":293225803,"eventType":"unfinished_order.created","data":{"orderId":"RA1SD","cartId":"320AB5DE-5EA6-4419-A4CF-0B04D2913190"}}`)
		d, err := e.UnfinishedOrderData()
		if err != nil {
			t.Fatalf("UnfinishedOrderData() = %v", err)
		}
		want := UnfinishedOrderData{OrderID: "RA1SD", CartID: "320AB5DE-5EA6-4419-A4CF-0B04D2913190"}
		if d != want {
			t.Errorf("UnfinishedOrderData() = %+v, want %+v", d, want)
		}
	})

	t.Run("review.created", func(t *testing.T) {
		e := mustUnmarshal(t, `{"eventId":"95ed494f-362d-4fca-933d-f5c7064b04bb","eventCreated":1634209655,"storeId":1003,"entityId":293225803,"eventType":"review.created","data":{"productId":62752,"orderId":63254,"status":"MODERATED"}}`)
		d, err := e.ReviewData()
		if err != nil {
			t.Fatalf("ReviewData() = %v", err)
		}
		// review's orderId is a number, unlike order.*/invoice.*'s string form.
		want := ReviewData{ProductID: 62752, OrderID: 63254, Status: "MODERATED"}
		if d != want {
			t.Errorf("ReviewData() = %+v, want %+v", d, want)
		}
	})

	t.Run("review.deleted has no status", func(t *testing.T) {
		e := mustUnmarshal(t, `{"eventId":"95ed494f-362d-4fca-933d-f5c7064b04bb","eventCreated":1634209655,"storeId":1003,"entityId":293225803,"eventType":"review.deleted","data":{"productId":62752,"orderId":63254}}`)
		d, err := e.ReviewData()
		if err != nil {
			t.Fatalf("ReviewData() = %v", err)
		}
		if want := (ReviewData{ProductID: 62752, OrderID: 63254}); d != want {
			t.Errorf("ReviewData() = %+v, want %+v", d, want)
		}
	})

	t.Run("customer.created", func(t *testing.T) {
		e := mustUnmarshal(t, `{"eventId":"80aece08-40e8-4145-8764-6c2f0d38678","eventCreated":7891234567,"storeId":1003,"entityId":1663830,"eventType":"customer.created","data":{"customerEmail":"user@example.com"}}`)
		d, err := e.CustomerData()
		if err != nil {
			t.Fatalf("CustomerData() = %v", err)
		}
		if want := (CustomerData{CustomerEmail: "user@example.com"}); d != want {
			t.Errorf("CustomerData() = %+v, want %+v", d, want)
		}
	})

	t.Run("promotion.created", func(t *testing.T) {
		e := mustUnmarshal(t, `{"eventId":"95ed494f-362d-4fca-933d-f5c7064b04bb","eventCreated":1469429915,"storeId":1003,"entityId":667251253,"eventType":"promotion.created","data":{"version":1}}`)
		d, err := e.PromotionData()
		if err != nil {
			t.Fatalf("PromotionData() = %v", err)
		}
		if want := (PromotionData{Version: 1}); d != want {
			t.Errorf("PromotionData() = %+v, want %+v", d, want)
		}
	})

	t.Run("invoice.created", func(t *testing.T) {
		e := mustUnmarshal(t, `{"eventId":"08a78904-4c1a-0aa0-953a-2e33c56236f1","eventCreated":1469429915,"storeId":1003,"entityId":667251253,"eventType":"invoice.created","data":{"orderId":"GH781"}}`)
		d, err := e.InvoiceData()
		if err != nil {
			t.Fatalf("InvoiceData() = %v", err)
		}
		if want := (InvoiceData{OrderID: "GH781"}); d != want {
			t.Errorf("InvoiceData() = %+v, want %+v", d, want)
		}
	})

	// The two subscriptionStatusChanged events look alike but use different keys.
	t.Run("application.subscriptionStatusChanged uses status", func(t *testing.T) {
		e := mustUnmarshal(t, `{"eventId":"80aece08-40e8-4145-8764-6c2f0d38678","eventCreated":1234567,"storeId":1003,"entityId":"1003","eventType":"application.subscriptionStatusChanged","data":{"oldSubscriptionStatus":"TRIAL","newSubscriptionStatus":"ACTIVE"}}`)
		d, err := e.AppSubscriptionData()
		if err != nil {
			t.Fatalf("AppSubscriptionData() = %v", err)
		}
		want := AppSubscriptionData{OldSubscriptionStatus: "TRIAL", NewSubscriptionStatus: "ACTIVE"}
		if d != want {
			t.Errorf("AppSubscriptionData() = %+v, want %+v", d, want)
		}
	})

	t.Run("profile.subscriptionStatusChanged uses name", func(t *testing.T) {
		e := mustUnmarshal(t, `{"eventId":"80aece08-40e8-4145-8764-6c2f0d38678","eventCreated":1494503041,"storeId":1421002,"entityId":1421002,"eventType":"profile.subscriptionStatusChanged","data":{"oldSubscriptionName":"ECWID_BUSINESS","newSubscriptionName":"ECWID_UNLIMITED"}}`)
		d, err := e.ProfileSubscriptionData()
		if err != nil {
			t.Fatalf("ProfileSubscriptionData() = %v", err)
		}
		want := ProfileSubscriptionData{OldSubscriptionName: "ECWID_BUSINESS", NewSubscriptionName: "ECWID_UNLIMITED"}
		if d != want {
			t.Errorf("ProfileSubscriptionData() = %+v, want %+v", d, want)
		}
	})
}

// An accessor must not hand back a zero value for an event that never carries
// that payload — silently reading data.orderId off a product event is exactly
// the bug typed accessors exist to prevent.
func TestEvent_DataAccessors_WrongEventType(t *testing.T) {
	e := mustUnmarshal(t, `{"eventId":"e","eventCreated":1,"storeId":1003,"entityId":667251253,"eventType":"product.created"}`)

	if _, err := e.OrderData(); err == nil {
		t.Error("OrderData() on product.created = nil, want an error")
	}
	if _, err := e.ReviewData(); err == nil {
		t.Error("ReviewData() on product.created = nil, want an error")
	}
	if _, err := e.AppSubscriptionData(); err == nil {
		t.Error("AppSubscriptionData() on product.created = nil, want an error")
	}

	// Right family, but this member carries no data.
	app := mustUnmarshal(t, `{"eventId":"e","eventCreated":1,"storeId":1003,"entityId":"1003","eventType":"application.installed"}`)
	if _, err := app.AppSubscriptionData(); err == nil {
		t.Error("AppSubscriptionData() on application.installed = nil, want an error")
	}
}

// Marshal must reproduce the wire format it parses, so an Event round-trips and
// a mock sender emits what Ecwid really sends.
func TestEvent_MarshalJSON_EntityIDShape(t *testing.T) {
	tests := []struct {
		name     string
		event    Event
		want     string
		wantType EventType
	}{
		{
			name:  "order event emits a bare number",
			event: Event{EventID: "e", EventCreated: 1234567, StoreID: 1003, EventType: EventOrderCreated, EntityID: "103878161"},
			want:  `"entityId":103878161`,
		},
		{
			name:  "application event emits a quoted string",
			event: Event{EventID: "e", EventCreated: 1234567, StoreID: 1003, EventType: EventApplicationInstalled, EntityID: "1003"},
			want:  `"entityId":"1003"`,
		},
		{
			name:  "non-numeric entityId stays quoted",
			event: Event{EventID: "e", EventCreated: 1, StoreID: 1003, EventType: EventProductCreated, EntityID: "abc"},
			want:  `"entityId":"abc"`,
		},
		{
			name:  "empty entityId stays quoted rather than emitting invalid JSON",
			event: Event{EventID: "e", EventCreated: 1, StoreID: 1003, EventType: EventProductCreated},
			want:  `"entityId":""`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := json.Marshal(tt.event)
			if err != nil {
				t.Fatalf("Marshal() = %v", err)
			}
			if !strings.Contains(string(b), tt.want) {
				t.Errorf("Marshal() = %s, want it to contain %s", b, tt.want)
			}

			var got Event
			if err := json.Unmarshal(b, &got); err != nil {
				t.Fatalf("Unmarshal(Marshal()) = %v", err)
			}
			if got.EntityID != tt.event.EntityID {
				t.Errorf("round-tripped EntityID = %q, want %q", got.EntityID, tt.event.EntityID)
			}
			if got.EventType != tt.event.EventType {
				t.Errorf("round-tripped EventType = %q, want %q", got.EventType, tt.event.EventType)
			}
		})
	}
}

func TestEvent_MarshalJSON_OmitsAbsentData(t *testing.T) {
	b, err := json.Marshal(Event{EventID: "e", EventCreated: 1, StoreID: 1003, EventType: EventProductCreated, EntityID: "1"})
	if err != nil {
		t.Fatalf("Marshal() = %v", err)
	}
	if strings.Contains(string(b), "data") {
		t.Errorf("Marshal() = %s, want no data key", b)
	}
}

func mustUnmarshal(t *testing.T, body string) Event {
	t.Helper()
	var e Event
	if err := json.Unmarshal([]byte(body), &e); err != nil {
		t.Fatalf("Unmarshal(%s) = %v", body, err)
	}
	return e
}
