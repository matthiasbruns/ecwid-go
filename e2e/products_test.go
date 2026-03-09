package e2e

import (
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/products"
)

func TestProducts_Search(t *testing.T) {
	ctx := testContext(t)

	result, err := testClient.Products.Search(ctx, &products.SearchOptions{Limit: 5})
	if err != nil {
		skipIfForbidden(t, err)
		t.Fatalf("Products.Search: %v", err)
	}
	if result.Items == nil {
		t.Error("expected non-nil items slice")
	}
}

func TestProducts_SearchWithFilters(t *testing.T) {
	ctx := testContext(t)

	// Create a product with a known price to search for.
	created, err := testClient.Products.Create(ctx, &products.Product{
		Name:  "ecwid-go-filter-test",
		Price: 42.50,
		SKU:   "ECWID-GO-FILTER-TEST",
	})
	if err != nil {
		skipIfForbidden(t, err)
		t.Fatalf("Products.Create: %v", err)
	}
	prodID := created.ID
	t.Cleanup(func() {
		_, _ = testClient.Products.Delete(testContext(t), prodID)
	})

	// Search by keyword.
	result, err := testClient.Products.Search(ctx, &products.SearchOptions{
		Keyword: "ecwid-go-filter-test",
		Limit:   10,
	})
	if err != nil {
		t.Fatalf("Products.Search(keyword): %v", err)
	}
	if result.Total == 0 {
		t.Error("expected at least one product matching keyword")
	}

	// Search by price range.
	priceFrom := 40.0
	priceTo := 45.0
	result, err = testClient.Products.Search(ctx, &products.SearchOptions{
		PriceFrom: &priceFrom,
		PriceTo:   &priceTo,
		Limit:     10,
	})
	if err != nil {
		t.Fatalf("Products.Search(price): %v", err)
	}
	if result.Total == 0 {
		t.Error("expected at least one product in price range 40-45")
	}

	// Search by SKU.
	result, err = testClient.Products.Search(ctx, &products.SearchOptions{
		SKU:   "ECWID-GO-FILTER-TEST",
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("Products.Search(sku): %v", err)
	}
	if result.Total == 0 {
		t.Error("expected at least one product matching SKU")
	}
}

func TestProducts_CRUD(t *testing.T) {
	ctx := testContext(t)

	// Create
	created, err := testClient.Products.Create(ctx, &products.Product{
		Name:        "ecwid-go-test-product",
		Price:       19.99,
		SKU:         "ECWID-GO-TEST",
		Description: "Test product created by ecwid-go E2E tests",
	})
	if err != nil {
		skipIfForbidden(t, err)
		t.Fatalf("Products.Create: %v", err)
	}
	if created.ID == 0 {
		t.Fatal("expected non-zero product ID")
	}
	prodID := created.ID

	// Clean up at the end.
	defer func() {
		_, _ = testClient.Products.Delete(testContext(t), prodID)
	}()

	// Get
	prod, err := testClient.Products.Get(ctx, prodID)
	if err != nil {
		t.Fatalf("Products.Get: %v", err)
	}
	if prod.Name != "ecwid-go-test-product" {
		t.Errorf("expected name=ecwid-go-test-product, got %s", prod.Name)
	}
	if prod.Price != 19.99 {
		t.Errorf("expected price=19.99, got %.2f", prod.Price)
	}
	if prod.SKU != "ECWID-GO-TEST" {
		t.Errorf("expected sku=ECWID-GO-TEST, got %s", prod.SKU)
	}

	// Update
	updated, err := testClient.Products.Update(ctx, prodID, &products.Product{
		Name:  "ecwid-go-test-product-updated",
		Price: 29.99,
	})
	if err != nil {
		t.Fatalf("Products.Update: %v", err)
	}
	if updated.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", updated.UpdateCount)
	}

	// Verify update
	prod, err = testClient.Products.Get(ctx, prodID)
	if err != nil {
		t.Fatalf("Products.Get (after update): %v", err)
	}
	if prod.Name != "ecwid-go-test-product-updated" {
		t.Errorf("expected updated name, got %s", prod.Name)
	}

	// Delete
	deleted, err := testClient.Products.Delete(ctx, prodID)
	if err != nil {
		t.Fatalf("Products.Delete: %v", err)
	}
	if deleted.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", deleted.DeleteCount)
	}
}
