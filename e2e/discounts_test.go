package e2e

import (
	"testing"
)

func TestDiscounts_SearchPromotions(t *testing.T) {
	requireClient(t)
	ctx := testContext(t)

	result, err := testClient.Discounts.SearchPromotions(ctx)
	if err != nil {
		skipIfForbidden(t, err)
		t.Fatalf("Discounts.SearchPromotions: %v", err)
	}
	t.Logf("found %d promotions", result.Total)
}

func TestDiscounts_SearchCoupons(t *testing.T) {
	requireClient(t)
	ctx := testContext(t)

	result, err := testClient.Discounts.SearchCoupons(ctx, nil)
	if err != nil {
		skipIfForbidden(t, err)
		t.Fatalf("Discounts.SearchCoupons: %v", err)
	}
	t.Logf("found %d coupons", result.Total)
}
