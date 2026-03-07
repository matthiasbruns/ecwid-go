package e2e

import (
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/customers"
)

func TestCustomers_Search(t *testing.T) {
	ctx := testContext(t)

	result, err := testClient.Customers.Search(ctx, &customers.SearchOptions{Limit: 5})
	if err != nil {
		t.Fatalf("Customers.Search: %v", err)
	}
	if result.Items == nil {
		t.Error("expected non-nil items slice")
	}
}

func TestCustomers_CRUD(t *testing.T) {
	ctx := testContext(t)

	// Create
	created, err := testClient.Customers.Create(ctx, &customers.Customer{
		Email:     "ecwid-go-test@example.com",
		FirstName: "Test",
		LastName:  "User",
	})
	if err != nil {
		t.Fatalf("Customers.Create: %v", err)
	}
	if created.ID == 0 {
		t.Fatal("expected non-zero customer ID")
	}
	custID := created.ID

	// Clean up at the end
	defer func() {
		_, _ = testClient.Customers.Delete(testContext(t), custID)
	}()

	// Get
	cust, err := testClient.Customers.Get(ctx, custID)
	if err != nil {
		t.Fatalf("Customers.Get: %v", err)
	}
	if cust.Email != "ecwid-go-test@example.com" {
		t.Errorf("expected email=ecwid-go-test@example.com, got %s", cust.Email)
	}

	// Update
	updated, err := testClient.Customers.Update(ctx, custID, &customers.Customer{
		FirstName: "Updated",
	})
	if err != nil {
		t.Fatalf("Customers.Update: %v", err)
	}
	if updated.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", updated.UpdateCount)
	}

	// Delete
	deleted, err := testClient.Customers.Delete(ctx, custID)
	if err != nil {
		t.Fatalf("Customers.Delete: %v", err)
	}
	if deleted.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", deleted.DeleteCount)
	}
}
