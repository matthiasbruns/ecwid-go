package domains_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/domains"
	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

func newTestService(t *testing.T, srv *httptest.Server) *domains.Service {
	t.Helper()
	return domains.NewService(api.NewHTTPClient(api.HTTPClientConfig{
		BaseURL: srv.URL + "/api/v3",
		StoreID: "12345",
		Token:   "secret_test",
	}))
}

func TestGet(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/domains" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"instantSiteDomain": {
				"primaryInstantSiteDomain": "1003",
				"ecwidSubdomain": "1003",
				"instantSiteIpAddress": "18.213.217.106",
				"instantSiteUrl": "https://1003.company.site"
			},
			"purchasedDomains": [
				{
					"id": 42,
					"name": "mystore.com",
					"status": "connected",
					"connectedToInstantSite": true,
					"primaryDomain": true
				}
			]
		}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Get(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if result.InstantSiteDomain == nil {
		t.Fatal("expected instantSiteDomain to be non-nil")
	}
	if result.InstantSiteDomain.EcwidSubdomain != "1003" {
		t.Errorf("expected ecwidSubdomain=1003, got %s", result.InstantSiteDomain.EcwidSubdomain)
	}
	if len(result.PurchasedDomains) != 1 {
		t.Fatalf("expected 1 purchased domain, got %d", len(result.PurchasedDomains))
	}
	if result.PurchasedDomains[0].Name != "mystore.com" {
		t.Errorf("expected name=mystore.com, got %s", result.PurchasedDomains[0].Name)
	}
}

func TestGet_NoPurchasedDomains(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"instantSiteDomain":{"ecwidSubdomain":"test"},"purchasedDomains":[]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Get(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(result.PurchasedDomains) != 0 {
		t.Errorf("expected 0 purchased domains, got %d", len(result.PurchasedDomains))
	}
}

func TestPurchase(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/domains/purchase" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":99,"name":"mynewstore.com","status":"pending","connectedToInstantSite":false,"primaryDomain":false}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Purchase(context.Background(), &domains.PurchaseRequest{
		DomainName:  "mynewstore.com",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john@example.com",
		CountryCode: "US",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.ID != 99 {
		t.Errorf("expected id=99, got %d", result.ID)
	}
	if result.Name != "mynewstore.com" {
		t.Errorf("expected name=mynewstore.com, got %s", result.Name)
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
}
