package subscriptions_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
	"github.com/matthiasbruns/ecwid-go/ecwid/subscriptions"
)

func newTestService(t *testing.T, srv *httptest.Server) *subscriptions.Service {
	t.Helper()
	return subscriptions.NewService(api.NewHTTPClient(api.HTTPClientConfig{
		BaseURL: srv.URL + "/api/v3",
		StoreID: "12345",
		Token:   "secret_test",
	}))
}

func TestSearch(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/subscriptions" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("status") != "ACTIVE" {
			t.Error("expected status=ACTIVE")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":1,"count":1,"offset":0,"limit":100,"items":[{"subscriptionId":42,"customerId":7,"status":"ACTIVE"}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Search(context.Background(), &subscriptions.SearchOptions{Status: "ACTIVE"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 1 {
		t.Errorf("expected total=1, got %d", result.Total)
	}
	if result.Items[0].SubscriptionID != 42 {
		t.Errorf("expected subscriptionId=42, got %d", result.Items[0].SubscriptionID)
	}
}

func TestGet(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/subscriptions/42" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"subscriptionId":42,"customerId":7,"status":"ACTIVE","chargeSettings":{"recurringInterval":"MONTH","recurringIntervalCount":1}}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	sub, err := svc.Get(context.Background(), 42)
	if err != nil {
		t.Fatal(err)
	}
	if sub.Status != "ACTIVE" {
		t.Errorf("expected ACTIVE, got %s", sub.Status)
	}
	if sub.ChargeSettings == nil || sub.ChargeSettings.RecurringInterval != "MONTH" {
		t.Error("expected chargeSettings with MONTH interval")
	}
}

func TestUpdate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/subscriptions/42" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Update(context.Background(), 42, &subscriptions.UpdateRequest{
		Status: "CANCELLED",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
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
