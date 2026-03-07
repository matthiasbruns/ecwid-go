package e2e

import (
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/categories"
)

func TestCategories_Search(t *testing.T) {
	ctx := testContext(t)

	result, err := testClient.Categories.Search(ctx, &categories.SearchOptions{Limit: 5})
	if err != nil {
		t.Fatalf("Categories.Search: %v", err)
	}
	if result.Items == nil {
		t.Error("expected non-nil items slice")
	}
}

func TestCategories_CRUD(t *testing.T) {
	ctx := testContext(t)

	// Create
	created, err := testClient.Categories.Create(ctx, &categories.Category{
		Name: "ecwid-go-test-category",
	})
	if err != nil {
		t.Fatalf("Categories.Create: %v", err)
	}
	if created.ID == 0 {
		t.Fatal("expected non-zero category ID")
	}
	catID := created.ID

	// Clean up at the end
	defer func() {
		_, _ = testClient.Categories.Delete(testContext(t), catID)
	}()

	// Get
	cat, err := testClient.Categories.Get(ctx, catID)
	if err != nil {
		t.Fatalf("Categories.Get: %v", err)
	}
	if cat.Name != "ecwid-go-test-category" {
		t.Errorf("expected name=ecwid-go-test-category, got %s", cat.Name)
	}

	// Update
	updated, err := testClient.Categories.Update(ctx, catID, &categories.Category{
		Name: "ecwid-go-test-category-updated",
	})
	if err != nil {
		t.Fatalf("Categories.Update: %v", err)
	}
	if updated.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", updated.UpdateCount)
	}

	// Delete
	deleted, err := testClient.Categories.Delete(ctx, catID)
	if err != nil {
		t.Fatalf("Categories.Delete: %v", err)
	}
	if deleted.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", deleted.DeleteCount)
	}
}
