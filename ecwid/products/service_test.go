package products_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
	"github.com/matthiasbruns/ecwid-go/ecwid/products"
)

func newTestService(t *testing.T, srv *httptest.Server) *products.Service {
	t.Helper()
	return products.NewService(api.NewHTTPClient(api.HTTPClientConfig{
		BaseURL: srv.URL + "/api/v3",
		StoreID: "12345",
		Token:   "secret_test",
	}))
}

// ── Core ────────────────────────────────────────────────────────────────

func TestSearch(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/products" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("keyword") != "shirt" {
			t.Error("expected keyword=shirt")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":1,"count":1,"offset":0,"limit":100,"items":[{"id":42,"name":"Blue Shirt","price":29.99}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Search(context.Background(), &products.SearchOptions{Keyword: "shirt"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 1 {
		t.Errorf("expected total=1, got %d", result.Total)
	}
	if result.Items[0].ID != 42 {
		t.Errorf("expected id=42, got %d", result.Items[0].ID)
	}
}

func TestSearch_WithQueryParams(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("priceFrom") != "10.00" {
			t.Errorf("expected priceFrom=10.00, got %q", q.Get("priceFrom"))
		}
		if q.Get("category") != "5" {
			t.Errorf("expected category=5, got %q", q.Get("category"))
		}
		if q.Get("limit") != "50" {
			t.Errorf("expected limit=50, got %q", q.Get("limit"))
		}
		if q.Get("enabled") != "true" {
			t.Errorf("expected enabled=true, got %q", q.Get("enabled"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":0,"count":0,"offset":0,"limit":50,"items":[]}`))
	}))
	defer srv.Close()

	priceFrom := 10.0
	enabled := true
	svc := newTestService(t, srv)
	_, err := svc.Search(context.Background(), &products.SearchOptions{
		PriceFrom: &priceFrom,
		Category:  5,
		Limit:     50,
		Enabled:   &enabled,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearch_NilOpts(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":0,"count":0,"offset":0,"limit":100,"items":[]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Search(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 0 {
		t.Errorf("expected total=0, got %d", result.Total)
	}
}

func TestGet(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/products/42" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":42,"name":"Blue Shirt","price":29.99,"sku":"BS-001"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	product, err := svc.Get(context.Background(), 42)
	if err != nil {
		t.Fatal(err)
	}
	if product.Name != "Blue Shirt" {
		t.Errorf("expected Blue Shirt, got %s", product.Name)
	}
	if product.SKU != "BS-001" {
		t.Errorf("expected BS-001, got %s", product.SKU)
	}
}

func TestGet_ZeroID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.Get(context.Background(), 0)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestCreate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":100}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Create(context.Background(), &products.Product{
		Name:  "New Product",
		Price: 19.99,
		SKU:   "NP-001",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.ID != 100 {
		t.Errorf("expected id=100, got %d", result.ID)
	}
}

func TestUpdate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Update(context.Background(), 42, &products.Product{
		Name: "Updated Product",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
	}
}

func TestUpdate_ZeroID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.Update(context.Background(), 0, nil)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestDelete(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Delete(context.Background(), 42)
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", result.DeleteCount)
	}
}

func TestDelete_ZeroID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.Delete(context.Background(), 0)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestDeleteAll(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusAccepted)
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	err := svc.DeleteAll(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}

// ── Combinations ────────────────────────────────────────────────────────

func TestListCombinations(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/products/42/combinations" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"id":1,"sku":"VAR-A"},{"id":2,"sku":"VAR-B"}]`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	combos, err := svc.ListCombinations(context.Background(), 42)
	if err != nil {
		t.Fatal(err)
	}
	if len(combos) != 2 {
		t.Fatalf("expected 2 combinations, got %d", len(combos))
	}
	if combos[0].SKU != "VAR-A" {
		t.Errorf("expected VAR-A, got %s", combos[0].SKU)
	}
}

func TestListCombinations_ZeroID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.ListCombinations(context.Background(), 0)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestGetCombination(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/products/42/combinations/10" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":10,"sku":"VAR-A"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	combo, err := svc.GetCombination(context.Background(), 42, 10)
	if err != nil {
		t.Fatal(err)
	}
	if combo.SKU != "VAR-A" {
		t.Errorf("expected VAR-A, got %s", combo.SKU)
	}
}

func TestGetCombination_ZeroProductID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.GetCombination(context.Background(), 0, 10)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestGetCombination_ZeroCombinationID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.GetCombination(context.Background(), 42, 0)
	if err == nil {
		t.Fatal("expected error for zero combinationID")
	}
}

func TestCreateCombination(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42/combinations" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":10}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.CreateCombination(context.Background(), 42, &products.Combination{SKU: "VAR-NEW"})
	if err != nil {
		t.Fatal(err)
	}
	if result.ID != 10 {
		t.Errorf("expected id=10, got %d", result.ID)
	}
}

func TestCreateCombination_ZeroID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.CreateCombination(context.Background(), 0, nil)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestUpdateCombination(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42/combinations/10" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.UpdateCombination(context.Background(), 42, 10, &products.Combination{SKU: "VAR-UPD"})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
	}
}

func TestUpdateCombination_ZeroProductID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.UpdateCombination(context.Background(), 0, 10, nil)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestUpdateCombination_ZeroCombinationID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.UpdateCombination(context.Background(), 42, 0, nil)
	if err == nil {
		t.Fatal("expected error for zero combinationID")
	}
}

func TestDeleteCombination(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42/combinations/10" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.DeleteCombination(context.Background(), 42, 10)
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", result.DeleteCount)
	}
}

func TestDeleteCombination_ZeroProductID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.DeleteCombination(context.Background(), 0, 10)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestDeleteCombination_ZeroCombinationID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.DeleteCombination(context.Background(), 42, 0)
	if err == nil {
		t.Fatal("expected error for zero combinationID")
	}
}

func TestDeleteAllCombinations(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42/combinations" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":5}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.DeleteAllCombinations(context.Background(), 42)
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 5 {
		t.Errorf("expected deleteCount=5, got %d", result.DeleteCount)
	}
}

func TestDeleteAllCombinations_ZeroID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.DeleteAllCombinations(context.Background(), 0)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

// ── Inventory ───────────────────────────────────────────────────────────

func TestAdjustInventory(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42/inventory" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.AdjustInventory(context.Background(), 42, &products.InventoryAdjust{QuantityDelta: -10})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
	}
}

func TestAdjustInventory_ZeroID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.AdjustInventory(context.Background(), 0, nil)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestAdjustCombinationInventory(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42/combinations/10/inventory" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.AdjustCombinationInventory(context.Background(), 42, 10, &products.InventoryAdjust{QuantityDelta: 5})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
	}
}

func TestAdjustCombinationInventory_ZeroProductID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.AdjustCombinationInventory(context.Background(), 0, 10, nil)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestAdjustCombinationInventory_ZeroCombinationID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.AdjustCombinationInventory(context.Background(), 42, 0, nil)
	if err == nil {
		t.Fatal("expected error for zero combinationID")
	}
}

// ── Images & Gallery (delete) ───────────────────────────────────────────

func TestDeleteImage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42/image" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.DeleteImage(context.Background(), 42)
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", result.DeleteCount)
	}
}

func TestDeleteImage_ZeroID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.DeleteImage(context.Background(), 0)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestDeleteAllGalleryImages(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42/gallery" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":3}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.DeleteAllGalleryImages(context.Background(), 42)
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 3 {
		t.Errorf("expected deleteCount=3, got %d", result.DeleteCount)
	}
}

func TestDeleteGalleryImage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42/gallery/5" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.DeleteGalleryImage(context.Background(), 42, 5)
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", result.DeleteCount)
	}
}

func TestDeleteGalleryImage_ZeroProductID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.DeleteGalleryImage(context.Background(), 0, 5)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestDeleteGalleryImage_ZeroImageID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.DeleteGalleryImage(context.Background(), 42, 0)
	if err == nil {
		t.Fatal("expected error for zero galleryImageID")
	}
}

func TestDeleteCombinationImage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42/combinations/10/image" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.DeleteCombinationImage(context.Background(), 42, 10)
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", result.DeleteCount)
	}
}

func TestDeleteCombinationImage_ZeroProductID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.DeleteCombinationImage(context.Background(), 0, 10)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestDeleteCombinationImage_ZeroCombinationID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.DeleteCombinationImage(context.Background(), 42, 0)
	if err == nil {
		t.Fatal("expected error for zero combinationID")
	}
}

// ── Video ───────────────────────────────────────────────────────────────

func TestDeleteVideo(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42/video" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.DeleteVideo(context.Background(), 42)
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", result.DeleteCount)
	}
}

func TestDeleteVideo_ZeroID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.DeleteVideo(context.Background(), 0)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestGetGalleryVideo(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/products/42/gallery/video/7" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":7,"url":"https://example.com/video.mp4","title":"Demo"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	video, err := svc.GetGalleryVideo(context.Background(), 42, 7)
	if err != nil {
		t.Fatal(err)
	}
	if video.Title != "Demo" {
		t.Errorf("expected Demo, got %s", video.Title)
	}
}

func TestGetGalleryVideo_ZeroProductID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.GetGalleryVideo(context.Background(), 0, 7)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestGetGalleryVideo_ZeroVideoID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.GetGalleryVideo(context.Background(), 42, 0)
	if err == nil {
		t.Fatal("expected error for zero galleryVideoID")
	}
}

func TestDeleteGalleryVideo(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42/gallery/7" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.DeleteGalleryVideo(context.Background(), 42, 7)
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", result.DeleteCount)
	}
}

func TestDeleteGalleryVideo_ZeroProductID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.DeleteGalleryVideo(context.Background(), 0, 7)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestDeleteGalleryVideo_ZeroVideoID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.DeleteGalleryVideo(context.Background(), 42, 0)
	if err == nil {
		t.Fatal("expected error for zero galleryVideoID")
	}
}

// ── Files ───────────────────────────────────────────────────────────────

func TestDeleteAllFiles(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42/files" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":2}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.DeleteAllFiles(context.Background(), 42)
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 2 {
		t.Errorf("expected deleteCount=2, got %d", result.DeleteCount)
	}
}

func TestDeleteAllFiles_ZeroID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.DeleteAllFiles(context.Background(), 0)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestUpdateFileDescription(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42/files/3" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.UpdateFileDescription(context.Background(), 42, 3, &products.ProductFileUpdate{Description: "updated description"})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
	}
}

func TestDeleteFile(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42/files/3" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.DeleteFile(context.Background(), 42, 3)
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", result.DeleteCount)
	}
}

func TestDeleteFile_ZeroProductID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.DeleteFile(context.Background(), 0, 3)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestDeleteFile_ZeroFileID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.DeleteFile(context.Background(), 42, 0)
	if err == nil {
		t.Fatal("expected error for zero fileID")
	}
}

// ── Other ───────────────────────────────────────────────────────────────

func TestUpdateMedia(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42/media" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.UpdateMedia(context.Background(), 42, &products.MediaUpdate{})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
	}
}

func TestUpdateMedia_ZeroID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.UpdateMedia(context.Background(), 0, nil)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestGetSortOrder(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/products/sort" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("parentCategory") != "5" {
			t.Errorf("expected parentCategory=5, got %q", r.URL.Query().Get("parentCategory"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"sortedIds":[1,2,3]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.GetSortOrder(context.Background(), 5)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.SortedIDs) != 3 {
		t.Errorf("expected 3 sorted IDs, got %d", len(result.SortedIDs))
	}
}

func TestGetSortOrder_ZeroParentCategory(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/products/sort" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("parentCategory") != "0" {
			t.Errorf("expected parentCategory=0, got %q", r.URL.Query().Get("parentCategory"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"sortedIds":[1,2,3]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.GetSortOrder(context.Background(), 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.SortedIDs) != 3 {
		t.Errorf("expected 3 sorted IDs, got %d", len(result.SortedIDs))
	}
}

func TestGetSortOrder_NegativeParentCategory(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.GetSortOrder(context.Background(), -1)
	if err == nil {
		t.Fatal("expected error for negative parentCategory")
	}
}

func TestUpdateSortOrder(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Query().Get("parentCategory") != "5" {
			t.Errorf("expected parentCategory=5, got %q", r.URL.Query().Get("parentCategory"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.UpdateSortOrder(context.Background(), &products.SortOrderUpdate{
		SortedIDs:      []int64{3, 1, 2},
		ParentCategory: 5,
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
	}
}

func TestUpdateSortOrder_Nil(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.UpdateSortOrder(context.Background(), nil)
	if err == nil {
		t.Fatal("expected error for nil sort")
	}
}

func TestGetDeleted(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/products/deleted" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("from_date") != "2024-01-01" {
			t.Errorf("expected from_date=2024-01-01, got %q", r.URL.Query().Get("from_date"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":1,"count":1,"offset":0,"limit":100,"items":[{"id":99,"date":"2024-06-15"}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.GetDeleted(context.Background(), &products.DeletedProductsOptions{FromDate: "2024-01-01"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 1 {
		t.Errorf("expected total=1, got %d", result.Total)
	}
	if result.Items[0].ID != 99 {
		t.Errorf("expected id=99, got %d", result.Items[0].ID)
	}
}

func TestGetFilters(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"productCount":2,"filters":{"price":{"minValue":0,"maxValue":100,"status":"SUCCESS"}}}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.GetFilters(context.Background(), &products.FiltersRequest{
		Params: &products.FiltersParams{
			FilterFields: "price",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.ProductCount != 2 {
		t.Errorf("expected productCount=2, got %d", result.ProductCount)
	}
	if result.Filters == nil {
		t.Error("expected non-nil filters")
	}
}

// ── Product Classes ─────────────────────────────────────────────────────

func TestListClasses(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/classes" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"id":1,"name":"General"},{"id":2,"name":"Clothing"}]`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	classes, err := svc.ListClasses(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(classes) != 2 {
		t.Fatalf("expected 2 classes, got %d", len(classes))
	}
	if classes[0].Name != "General" {
		t.Errorf("expected General, got %s", classes[0].Name)
	}
}

func TestGetClass(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/classes/1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":1,"name":"General"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	class, err := svc.GetClass(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
	if class.Name != "General" {
		t.Errorf("expected General, got %s", class.Name)
	}
}

func TestGetClass_ZeroID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.GetClass(context.Background(), 0)
	if err == nil {
		t.Fatal("expected error for zero classID")
	}
}

func TestCreateClass(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/classes" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":3}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.CreateClass(context.Background(), &products.ProductClass{Name: "Electronics"})
	if err != nil {
		t.Fatal(err)
	}
	if result.ID != 3 {
		t.Errorf("expected id=3, got %d", result.ID)
	}
}

func TestUpdateClass(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/classes/1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.UpdateClass(context.Background(), 1, &products.ProductClass{Name: "Updated"})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
	}
}

func TestUpdateClass_ZeroID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.UpdateClass(context.Background(), 0, nil)
	if err == nil {
		t.Fatal("expected error for zero classID")
	}
}

func TestDeleteClass(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/classes/1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.DeleteClass(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", result.DeleteCount)
	}
}

func TestDeleteClass_ZeroID(t *testing.T) {
	svc := products.NewService(nil)
	_, err := svc.DeleteClass(context.Background(), 0)
	if err == nil {
		t.Fatal("expected error for zero classID")
	}
}

// ── Brands & Swatches ───────────────────────────────────────────────────

func TestListBrands(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/brands" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":1,"count":1,"offset":0,"limit":100,"items":[{"name":"Nike","productsFilteredByBrandUrl":"https://example.com/search?attribute_Brand=Nike"}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.ListBrands(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 1 {
		t.Errorf("expected total=1, got %d", result.Total)
	}
	if result.Items[0].Name != "Nike" {
		t.Errorf("expected Nike, got %s", result.Items[0].Name)
	}
}

func TestListSwatches(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/swatches" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"colors":[{"name":"Red","hexCode":"#ff0000"}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.ListSwatches(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Colors) != 1 {
		t.Fatalf("expected 1 color, got %d", len(result.Colors))
	}
	if result.Colors[0].Name != "Red" {
		t.Errorf("expected Red, got %s", result.Colors[0].Name)
	}
}

// ── Error handling ──────────────────────────────────────────────────────

func TestSearch_Error(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"errorMessage":"access denied","errorCode":"403"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	_, err := svc.Search(context.Background(), nil)
	if err == nil {
		t.Fatal("expected error")
	}

	var apiErr *api.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *api.APIError, got %T", err)
	}
}
