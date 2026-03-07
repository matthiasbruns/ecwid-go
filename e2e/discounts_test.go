package e2e

import (
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/discounts"
)

func TestDiscounts_SearchPromotions(t *testing.T) {
	ctx := testContext(t)

	result, err := testClient.Discounts.SearchPromotions(ctx)
	if err != nil {
		t.Fatalf("SearchPromotions: %v", err)
	}
	_ = result
}

func TestDiscounts_PromotionCRUD(t *testing.T) {
	ctx := testContext(t)

	enabled := true
	created, err := testClient.Discounts.CreatePromotion(ctx, &discounts.Promotion{
		Name:    "E2E Test Promo",
		Enabled: &enabled,
	})
	if err != nil {
		t.Fatalf("CreatePromotion: %v", err)
	}
	if created.ID == 0 {
		t.Fatal("expected non-zero promotion ID")
	}

	promo, err := testClient.Discounts.GetPromotion(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetPromotion: %v", err)
	}
	if promo.Name != "E2E Test Promo" {
		t.Errorf("expected name=E2E Test Promo, got %s", promo.Name)
	}

	updated, err := testClient.Discounts.UpdatePromotion(ctx, created.ID, &discounts.Promotion{
		Name: "E2E Test Promo Updated",
	})
	if err != nil {
		t.Fatalf("UpdatePromotion: %v", err)
	}
	if updated.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", updated.UpdateCount)
	}

	deleted, err := testClient.Discounts.DeletePromotion(ctx, created.ID)
	if err != nil {
		t.Fatalf("DeletePromotion: %v", err)
	}
	if deleted.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", deleted.DeleteCount)
	}
}

func TestDiscounts_SearchCoupons(t *testing.T) {
	ctx := testContext(t)

	result, err := testClient.Discounts.SearchCoupons(ctx, nil)
	if err != nil {
		t.Fatalf("SearchCoupons: %v", err)
	}
	_ = result
}

func TestDiscounts_CouponCRUD(t *testing.T) {
	ctx := testContext(t)

	created, err := testClient.Discounts.CreateCoupon(ctx, &discounts.Coupon{
		Name:         "E2E Test Coupon",
		Code:         "E2ETEST10",
		DiscountType: "PERCENT",
		Discount:     10,
	})
	if err != nil {
		t.Fatalf("CreateCoupon: %v", err)
	}
	if created.ID == 0 {
		t.Fatal("expected non-zero coupon ID")
	}

	coupon, err := testClient.Discounts.GetCoupon(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetCoupon: %v", err)
	}
	if coupon.Code != "E2ETEST10" {
		t.Errorf("expected code=E2ETEST10, got %s", coupon.Code)
	}

	updated, err := testClient.Discounts.UpdateCoupon(ctx, created.ID, &discounts.Coupon{
		Name: "E2E Test Coupon Updated",
	})
	if err != nil {
		t.Fatalf("UpdateCoupon: %v", err)
	}
	if updated.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", updated.UpdateCount)
	}

	deleted, err := testClient.Discounts.DeleteCoupon(ctx, created.ID)
	if err != nil {
		t.Fatalf("DeleteCoupon: %v", err)
	}
	if deleted.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", deleted.DeleteCount)
	}
}
