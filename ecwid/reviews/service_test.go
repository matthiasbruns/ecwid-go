package reviews_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
	"github.com/matthiasbruns/ecwid-go/ecwid/reviews"
)

func newTestService(t *testing.T, srv *httptest.Server) *reviews.Service {
	t.Helper()
	return reviews.NewService(api.NewHTTPClient(api.HTTPClientConfig{
		BaseURL: srv.URL + "/api/v3",
		StoreID: "12345",
		Token:   "secret_test",
	}))
}

func TestSearch(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/reviews" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":1,"count":1,"offset":0,"limit":100,"items":[{"id":712737671,"status":"published","customerId":8108179152,"productId":123456,"orderId":"2D31G","rating":5,"review":"Just what I need","reviewerInfo":{"name":"Abraham Smith","email":"a@example.com","city":"New York","orders":2},"createDate":"2025-02-26 13:37:46 +0000","updateDate":"2025-02-27 04:40:33 +0000","createTimestamp":1740562666,"updateTimestamp":1740616833}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Search(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 1 {
		t.Errorf("expected total=1, got %d", result.Total)
	}
	r := result.Items[0]
	if r.ID != 712737671 {
		t.Errorf("expected id=712737671, got %d", r.ID)
	}
	if r.Status != "published" {
		t.Errorf("expected status=published, got %s", r.Status)
	}
	if r.CustomerID != 8108179152 {
		t.Errorf("expected customerId=8108179152, got %d", r.CustomerID)
	}
	if r.ReviewerInfo == nil || r.ReviewerInfo.Name != "Abraham Smith" {
		t.Errorf("expected reviewerInfo.name=Abraham Smith")
	}
}

func TestSearch_WithFilters(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("status") != "published" {
			t.Errorf("expected status=published, got %s", r.URL.Query().Get("status"))
		}
		if r.URL.Query().Get("rating") != "5" {
			t.Errorf("expected rating=5, got %s", r.URL.Query().Get("rating"))
		}
		if r.URL.Query().Get("productId") != "123" {
			t.Errorf("expected productId=123, got %s", r.URL.Query().Get("productId"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":0,"count":0,"offset":0,"limit":100,"items":[]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	_, err := svc.Search(context.Background(), &reviews.SearchOptions{
		Status:    "published",
		Rating:    5,
		ProductID: 123,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut || r.URL.Path != "/api/v3/12345/reviews/712737671" {
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.UpdateStatus(context.Background(), 712737671, "moderated")
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
	}
}

func TestUpdateStatus_ZeroID(t *testing.T) {
	svc := reviews.NewService(nil)
	_, err := svc.UpdateStatus(context.Background(), 0, "published")
	if err == nil {
		t.Fatal("expected error for zero reviewID")
	}
}

func TestDelete(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete || r.URL.Path != "/api/v3/12345/reviews/999" {
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Delete(context.Background(), 999)
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", result.DeleteCount)
	}
}

func TestDelete_ZeroID(t *testing.T) {
	svc := reviews.NewService(nil)
	_, err := svc.Delete(context.Background(), 0)
	if err == nil {
		t.Fatal("expected error for zero reviewID")
	}
}

func TestBulkUpdate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut || r.URL.Path != "/api/v3/12345/reviews/mass_update" {
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":5}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	del := false
	result, err := svc.BulkUpdate(context.Background(), &reviews.BulkUpdateRequest{
		SelectMode: "ALL",
		Delete:     &del,
		NewStatus:  "published",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 5 {
		t.Errorf("expected updateCount=5, got %d", result.UpdateCount)
	}
}

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
