package profile_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
	"github.com/matthiasbruns/ecwid-go/ecwid/profile"
)

func newTestService(t *testing.T, srv *httptest.Server) *profile.Service {
	t.Helper()
	return profile.NewService(api.NewHTTPClient(api.HTTPClientConfig{
		BaseURL: srv.URL + "/api/v3",
		StoreID: "12345",
		Token:   "secret_test",
	}))
}

func TestGet(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/profile" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"generalInfo": {"storeId": 12345, "storeUrl": "https://store.example.com"},
			"account": {"accountName": "Test Store", "accountEmail": "test@example.com"},
			"settings": {"storeName": "My Store", "closed": false},
			"company": {"companyName": "ACME Inc.", "city": "Berlin", "countryCode": "DE"}
		}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	p, err := svc.Get(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if p.GeneralInfo == nil || p.GeneralInfo.StoreID != 12345 {
		t.Error("expected storeId=12345")
	}
	if p.Account == nil || p.Account.AccountName != "Test Store" {
		t.Error("expected accountName=Test Store")
	}
	if p.Settings == nil || p.Settings.StoreName != "My Store" {
		t.Error("expected storeName=My Store")
	}
	if p.Company == nil || p.Company.CompanyName != "ACME Inc." {
		t.Error("expected companyName=ACME Inc.")
	}
}

func TestUpdate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/profile" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Error("expected Content-Type: application/json")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Update(context.Background(), &profile.UpdateRequest{
		Settings: &profile.Settings{StoreName: "Updated Store"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
	}
}

func TestGet_Error(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"errorMessage":"access denied","errorCode":"403"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	_, err := svc.Get(context.Background())
	if err == nil {
		t.Fatal("expected error")
	}

	var apiErr *api.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *api.APIError, got %T", err)
	}
	if apiErr.StatusCode != http.StatusForbidden {
		t.Errorf("expected 403, got %d", apiErr.StatusCode)
	}
}
