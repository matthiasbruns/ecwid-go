package carts_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/carts"
	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

func newTestService(t *testing.T, srv *httptest.Server) *carts.Service {
	t.Helper()
	return carts.NewService(api.NewHTTPClient(api.HTTPClientConfig{
		BaseURL: srv.URL + "/api/v3",
		StoreID: "12345",
		Token:   "secret_test",
	}))
}

func TestSearch(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/carts" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("customerId") != "7" {
			t.Error("expected customerId=7")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":1,"count":1,"offset":0,"limit":100,"items":[{"cartId":"abc123","total":49.99,"email":"test@example.com"}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Search(context.Background(), &carts.SearchOptions{CustomerID: 7})
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 1 {
		t.Errorf("expected total=1, got %d", result.Total)
	}
	if result.Items[0].CartID != "abc123" {
		t.Errorf("expected cartId=abc123, got %s", result.Items[0].CartID)
	}
}

func TestGet(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/carts/abc123" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"cartId":"abc123","total":49.99,"email":"test@example.com"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	cart, err := svc.Get(context.Background(), "abc123")
	if err != nil {
		t.Fatal(err)
	}
	if cart.Total != 49.99 {
		t.Errorf("expected total=49.99, got %f", cart.Total)
	}
}

func TestGet_EmptyID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fatal("should not reach server")
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	_, err := svc.Get(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty cartID")
	}
}

func TestUpdate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Update(context.Background(), "abc123", &carts.UpdateRequest{Hidden: true})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
	}
}

func TestPlace(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/carts/abc123/place" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"orderNumber":1001}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Place(context.Background(), "abc123")
	if err != nil {
		t.Fatal(err)
	}
	if result.OrderNumber != 1001 {
		t.Errorf("expected orderNumber=1001, got %d", result.OrderNumber)
	}
}

func TestSearch_Error(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"errorMessage":"unauthorized","errorCode":"401"}`))
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
	if apiErr.StatusCode != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", apiErr.StatusCode)
	}
}
