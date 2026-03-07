package discounts_test

import (
	"context"
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
		if r.URL.Path != "/api/v3/12345/discount_rules" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":1,"items":[{"id":100,"name":"Summer Sale","enabled":true}]}`))
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
	if result.Items[0].Name != "Summer Sale" {
		t.Errorf("expected name=Summer Sale, got %s", result.Items[0].Name)
	}
}

func TestGetPromotion(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/discount_rules/100" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":100,"name":"Summer Sale","enabled":true}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	promo, err := svc.GetPromotion(context.Background(), 100)
	if err != nil {
		t.Fatal(err)
	}
	if promo.ID != 100 {
		t.Errorf("expected id=100, got %d", promo.ID)
	}
}

func TestGetPromotion_ZeroID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fatal("should not reach server")
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	_, err := svc.GetPromotion(context.Background(), 0)
	if err == nil {
		t.Fatal("expected error for zero ruleID")
	}
}

func TestCreatePromotion(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":101}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.CreatePromotion(context.Background(), &discounts.Promotion{Name: "New Promo"})
	if err != nil {
		t.Fatal(err)
	}
	if result.ID != 101 {
		t.Errorf("expected id=101, got %d", result.ID)
	}
}

func TestUpdatePromotion(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/discount_rules/100" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.UpdatePromotion(context.Background(), 100, &discounts.Promotion{Name: "Updated"})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
	}
}

func TestDeletePromotion(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/discount_rules/100" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.DeletePromotion(context.Background(), 100)
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
		if r.URL.Query().Get("code") != "SAVE10" {
			t.Error("expected code=SAVE10")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":1,"count":1,"offset":0,"limit":100,"items":[{"id":200,"name":"Save 10","code":"SAVE10","discountType":"PERCENT","status":"ACTIVE","discount":10}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.SearchCoupons(context.Background(), &discounts.CouponSearchOptions{Code: "SAVE10"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 1 {
		t.Errorf("expected total=1, got %d", result.Total)
	}
	if result.Items[0].Code != "SAVE10" {
		t.Errorf("expected code=SAVE10, got %s", result.Items[0].Code)
	}
}

func TestGetCoupon(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/discount_coupons/200" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":200,"name":"Save 10","code":"SAVE10","discountType":"PERCENT","status":"ACTIVE","discount":10}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	coupon, err := svc.GetCoupon(context.Background(), 200)
	if err != nil {
		t.Fatal(err)
	}
	if coupon.ID != 200 {
		t.Errorf("expected id=200, got %d", coupon.ID)
	}
}

func TestGetCoupon_ZeroID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fatal("should not reach server")
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
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
		_, _ = w.Write([]byte(`{"id":201}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.CreateCoupon(context.Background(), &discounts.Coupon{Name: "New Coupon", Code: "NEW10"})
	if err != nil {
		t.Fatal(err)
	}
	if result.ID != 201 {
		t.Errorf("expected id=201, got %d", result.ID)
	}
}

func TestUpdateCoupon(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/discount_coupons/200" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.UpdateCoupon(context.Background(), 200, &discounts.Coupon{Name: "Updated Coupon"})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
	}
}

func TestDeleteCoupon(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/discount_coupons/200" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.DeleteCoupon(context.Background(), 200)
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", result.DeleteCount)
	}
}
