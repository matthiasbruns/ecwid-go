package discounts_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/discounts"
	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

func newTestService(t *testing.T, srv *httptest.Server) *discounts.Service {
	t.Helper()
	return discounts.NewService(api.NewHTTPClient(api.HTTPClientConfig{
		BaseURL: srv.URL + "/api/v3",
		StoreID: "12345",
		Token:   "secret_test",
	}))
}

// --- Promotions ---

func TestSearchPromotions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/promotions" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":1,"count":1,"offset":0,"limit":100,"items":[{"id":608502,"name":"10% off","enabled":true,"discountBase":"SUBTOTAL","discountType":"PERCENT","amount":10}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.SearchPromotions(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 1 {
		t.Errorf("expected total=1, got %d", result.Total)
	}
	if result.Items[0].Name != "10% off" {
		t.Errorf("expected name=10%% off, got %s", result.Items[0].Name)
	}
	if result.Items[0].DiscountBase != "SUBTOTAL" {
		t.Errorf("expected discountBase=SUBTOTAL, got %s", result.Items[0].DiscountBase)
	}
}

func TestCreatePromotion(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/api/v3/12345/promotions" {
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":999}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.CreatePromotion(context.Background(), &discounts.Promotion{Name: "Test"})
	if err != nil {
		t.Fatal(err)
	}
	if result.ID != 999 {
		t.Errorf("expected id=999, got %d", result.ID)
	}
}

func TestUpdatePromotion(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut || r.URL.Path != "/api/v3/12345/promotions/42" {
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.UpdatePromotion(context.Background(), 42, &discounts.Promotion{Name: "Updated"})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
	}
}

func TestUpdatePromotion_ZeroID(t *testing.T) {
	svc := discounts.NewService(nil)
	_, err := svc.UpdatePromotion(context.Background(), 0, &discounts.Promotion{})
	if err == nil {
		t.Fatal("expected error for zero promotionID")
	}
}

func TestDeletePromotion(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete || r.URL.Path != "/api/v3/12345/promotions/42" {
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.DeletePromotion(context.Background(), 42)
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", result.DeleteCount)
	}
}

// --- Coupons ---

func TestSearchCoupons(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/discount_coupons" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("code") != "DISC" {
			t.Errorf("expected code=DISC, got %s", r.URL.Query().Get("code"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":1,"count":1,"offset":0,"limit":100,"items":[{"id":162428889,"name":"10% OFF","code":"DISC","discountType":"ABS","status":"ACTIVE","discount":10,"usesLimit":"UNLIMITED","applicationLimit":"UNLIMITED","repeatCustomerOnly":false}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.SearchCoupons(context.Background(), &discounts.CouponSearchOptions{Code: "DISC"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 1 {
		t.Errorf("expected total=1, got %d", result.Total)
	}
	if result.Items[0].ApplicationLimit != "UNLIMITED" {
		t.Errorf("expected applicationLimit=UNLIMITED, got %s", result.Items[0].ApplicationLimit)
	}
}

func TestGetCoupon(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/discount_coupons/162428889" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":162428889,"name":"10% OFF","code":"DISC","discountType":"ABS","status":"ACTIVE","discount":10,"usesLimit":"UNLIMITED","applicationLimit":"UNLIMITED"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.GetCoupon(context.Background(), 162428889)
	if err != nil {
		t.Fatal(err)
	}
	if result.ID != 162428889 {
		t.Errorf("expected id=162428889, got %d", result.ID)
	}
}

func TestGetCoupon_ZeroID(t *testing.T) {
	svc := discounts.NewService(nil)
	_, err := svc.GetCoupon(context.Background(), 0)
	if err == nil {
		t.Fatal("expected error for zero couponID")
	}
}

func TestCreateCoupon(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":999}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.CreateCoupon(context.Background(), &discounts.Coupon{Name: "Test", Code: "TEST10"})
	if err != nil {
		t.Fatal(err)
	}
	if result.ID != 999 {
		t.Errorf("expected id=999, got %d", result.ID)
	}
}

func TestDeleteCoupon(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.DeleteCoupon(context.Background(), 99)
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", result.DeleteCount)
	}
}

func TestSearchPromotions_Error(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"errorMessage":"access denied","errorCode":"403"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	_, err := svc.SearchPromotions(context.Background())
	if err == nil {
		t.Fatal("expected error")
	}
	var apiErr *api.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *api.APIError, got %T", err)
	}
}
