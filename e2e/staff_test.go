package e2e

import (
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/staff"
)

func TestStaff_List(t *testing.T) {
	ctx := testContext(t)

	result, err := testClient.Staff.List(ctx)
	if err != nil {
		t.Fatalf("Staff.List: %v", err)
	}
	if result.Total < 0 {
		t.Errorf("expected non-negative total, got %d", result.Total)
	}
}

func TestStaff_CRUD(t *testing.T) {
	ctx := testContext(t)

	// Create
	created, err := testClient.Staff.Create(ctx, &staff.StaffMember{
		Email:     "e2e-test@example.com",
		FirstName: "E2E",
		LastName:  "Test",
		Role:      "MANAGER",
	})
	if err != nil {
		t.Fatalf("Staff.Create: %v", err)
	}
	if created.ID == 0 {
		t.Fatal("expected non-zero ID from Create")
	}

	// Get
	member, err := testClient.Staff.Get(ctx, created.ID)
	if err != nil {
		t.Fatalf("Staff.Get: %v", err)
	}
	if member.Email != "e2e-test@example.com" {
		t.Errorf("expected e2e-test@example.com, got %s", member.Email)
	}

	// Update
	updated, err := testClient.Staff.Update(ctx, created.ID, &staff.StaffMember{
		FirstName: "Updated",
	})
	if err != nil {
		t.Fatalf("Staff.Update: %v", err)
	}
	if updated.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", updated.UpdateCount)
	}

	// Delete
	deleted, err := testClient.Staff.Delete(ctx, created.ID)
	if err != nil {
		t.Fatalf("Staff.Delete: %v", err)
	}
	if deleted.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", deleted.DeleteCount)
	}
}
