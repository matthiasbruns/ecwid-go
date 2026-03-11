package e2e

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/orders"
)

func TestOrders_Search(t *testing.T) {
	requireClient(t)
	ctx := testContext(t)

	result, err := testClient.Orders.Search(ctx, &orders.SearchOptions{Limit: 5})
	if err != nil {
		skipIfForbidden(t, err)
		t.Fatalf("Orders.Search: %v", err)
	}
	if result.Items == nil {
		t.Error("expected non-nil items slice")
	}
}

func TestOrders_CRUD(t *testing.T) {
	requireClient(t)
	ctx := testContext(t)

	items, _ := json.Marshal([]map[string]any{
		{
			"name":     "Test Item",
			"price":    9.99,
			"quantity": 1,
		},
	})

	// Create
	created, err := testClient.Orders.Create(ctx, &orders.CreateRequest{
		Subtotal:          9.99,
		Total:             9.99,
		Email:             "ecwid-go-test@example.com",
		PaymentStatus:     "PAID",
		FulfillmentStatus: "AWAITING_PROCESSING",
		Items:             items,
	})
	if err != nil {
		skipIfForbidden(t, err)
		t.Fatalf("Orders.Create: %v", err)
	}
	if created.ID == 0 {
		t.Fatal("expected non-zero order internal ID")
	}
	orderID := created.OrderID
	if orderID == "" {
		// Fall back to string ID from internal ID.
		orderID = fmt.Sprintf("%d", created.ID)
	}

	// Clean up at the end.
	defer func() {
		_, _ = testClient.Orders.Delete(testContext(t), orderID)
	}()

	// Get
	order, err := testClient.Orders.Get(ctx, orderID)
	if err != nil {
		t.Fatalf("Orders.Get: %v", err)
	}
	if order.Email != "ecwid-go-test@example.com" {
		t.Errorf("expected email=ecwid-go-test@example.com, got %s", order.Email)
	}
	if order.PaymentStatus != "PAID" {
		t.Errorf("expected paymentStatus=PAID, got %s", order.PaymentStatus)
	}

	// Update
	updated, err := testClient.Orders.Update(ctx, orderID, &orders.UpdateRequest{
		FulfillmentStatus: "SHIPPED",
		TrackingNumber:    "ECWID-GO-TRACK-123",
	})
	if err != nil {
		t.Fatalf("Orders.Update: %v", err)
	}
	if updated.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", updated.UpdateCount)
	}

	// Verify update
	order, err = testClient.Orders.Get(ctx, orderID)
	if err != nil {
		t.Fatalf("Orders.Get (after update): %v", err)
	}
	if order.FulfillmentStatus != "SHIPPED" {
		t.Errorf("expected fulfillmentStatus=SHIPPED, got %s", order.FulfillmentStatus)
	}

	// Delete
	deleted, err := testClient.Orders.Delete(ctx, orderID)
	if err != nil {
		t.Fatalf("Orders.Delete: %v", err)
	}
	if deleted.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", deleted.DeleteCount)
	}
}
