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
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/domain" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"customDomain":"shop.example.com","sslEnabled":true,"storeFrontUrl":"https://shop.example.com"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	settings, err := svc.Get(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if settings.CustomDomain != "shop.example.com" {
		t.Errorf("expected shop.example.com, got %s", settings.CustomDomain)
	}
	if !settings.SSLEnabled {
		t.Error("expected sslEnabled=true")
	}
}

func TestUpdate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/domain" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Update(context.Background(), &domains.DomainSettings{
		CustomDomain: "new.example.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
	}
}

func TestTemplates(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/domain/templates" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"name":"basic","price":"9.99"},{"name":"premium","price":"19.99"}]`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	templates, err := svc.Templates(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(templates) != 2 {
		t.Fatalf("expected 2 templates, got %d", len(templates))
	}
	if templates[0].Name != "basic" {
		t.Errorf("expected basic, got %s", templates[0].Name)
	}
}

func TestWhois(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/domain/whois/example.com" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"available":true,"domain":"example.com"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Whois(context.Background(), "example.com")
	if err != nil {
		t.Fatal(err)
	}
	if !result.Available {
		t.Error("expected available=true")
	}
	if result.Domain != "example.com" {
		t.Errorf("expected example.com, got %s", result.Domain)
	}
}

func TestWhois_EmptyDomain(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fatal("should not reach server")
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	_, err := svc.Whois(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty domainName")
	}
}

func TestPurchase(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/domain/purchase" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"pending"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Purchase(context.Background(), &domains.PurchaseRequest{
		DomainName: "example.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != "pending" {
		t.Errorf("expected status=pending, got %s", result.Status)
	}
}

func TestRemove(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/domain" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Remove(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", result.DeleteCount)
	}
}

func TestGet_Error(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"errorMessage":"internal error","errorCode":"500"}`))
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
	if apiErr.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", apiErr.StatusCode)
	}
}
