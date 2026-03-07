package reviews_test

import (
	"context"
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

func TestList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/products/42/reviews" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":1,"items":[{"id":1,"productId":42,"rating":5,"reviewText":"Great!","reviewerName":"John"}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.List(context.Background(), 42)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 1 {
		t.Errorf("expected total=1, got %d", result.Total)
	}
	if result.Items[0].ReviewerName != "John" {
		t.Errorf("expected reviewerName=John, got %s", result.Items[0].ReviewerName)
	}
}

func TestList_ZeroID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fatal("should not reach server")
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	_, err := svc.List(context.Background(), 0)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestGet(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/products/42/reviews/1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":1,"productId":42,"rating":5,"reviewText":"Great!","reviewerName":"John"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	review, err := svc.Get(context.Background(), 42, 1)
	if err != nil {
		t.Fatal(err)
	}
	if review.ID != 1 {
		t.Errorf("expected id=1, got %d", review.ID)
	}
}

func TestGet_ZeroProductID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fatal("should not reach server")
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	_, err := svc.Get(context.Background(), 0, 1)
	if err == nil {
		t.Fatal("expected error for zero productID")
	}
}

func TestGet_ZeroReviewID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fatal("should not reach server")
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	_, err := svc.Get(context.Background(), 42, 0)
	if err == nil {
		t.Fatal("expected error for zero reviewID")
	}
}

func TestCreate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42/reviews" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":2}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Create(context.Background(), 42, &reviews.Review{Rating: 5, ReviewText: "Awesome!"})
	if err != nil {
		t.Fatal(err)
	}
	if result.ID != 2 {
		t.Errorf("expected id=2, got %d", result.ID)
	}
}

func TestUpdate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42/reviews/1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Update(context.Background(), 42, 1, &reviews.Review{Rating: 4})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
	}
}

func TestDelete(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42/reviews/1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Delete(context.Background(), 42, 1)
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", result.DeleteCount)
	}
}

func TestDeleteAll(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/products/42/reviews" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":5}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.DeleteAll(context.Background(), 42)
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 5 {
		t.Errorf("expected deleteCount=5, got %d", result.DeleteCount)
	}
}
