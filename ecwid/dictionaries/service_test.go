package dictionaries_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/dictionaries"
	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

func newTestService(t *testing.T, srv *httptest.Server) *dictionaries.Service {
	t.Helper()
	requester := api.NewHTTPClient(api.HTTPClientConfig{
		BaseURL: srv.URL + "/api/v3",
		StoreID: "12345",
		Token:   "secret_test",
	})
	return dictionaries.NewService(requester)
}

func TestCountries(t *testing.T) {
	fixture, err := os.ReadFile("testdata/countries.json")
	if err != nil {
		t.Fatal(err)
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/countries" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(fixture)
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	countries, err := svc.Countries(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(countries) != 2 {
		t.Fatalf("expected 2 countries, got %d", len(countries))
	}
	if countries[0].Code != "US" {
		t.Errorf("expected US, got %s", countries[0].Code)
	}
}

func TestCountriesWithStates(t *testing.T) {
	fixture, err := os.ReadFile("testdata/countries.json")
	if err != nil {
		t.Fatal(err)
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("withStates") != "true" {
			t.Error("expected withStates=true")
		}
		if r.URL.Query().Get("lang") != "de" {
			t.Error("expected lang=de")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(fixture)
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	countries, err := svc.Countries(context.Background(), &dictionaries.CountriesOptions{
		Lang:       "de",
		WithStates: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(countries[0].States) != 2 {
		t.Fatalf("expected 2 states, got %d", len(countries[0].States))
	}
}

func TestCurrencies(t *testing.T) {
	fixture, err := os.ReadFile("testdata/currencies.json")
	if err != nil {
		t.Fatal(err)
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/currencies" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(fixture)
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	currencies, err := svc.Currencies(context.Background(), "")
	if err != nil {
		t.Fatal(err)
	}
	if len(currencies) != 2 {
		t.Fatalf("expected 2 currencies, got %d", len(currencies))
	}
	if currencies[0].Code != "USD" {
		t.Errorf("expected USD, got %s", currencies[0].Code)
	}
}

func TestCurrencyByCountry(t *testing.T) {
	fixture, err := os.ReadFile("testdata/currency_by_country.json")
	if err != nil {
		t.Fatal(err)
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("countryCode") != "US" {
			t.Error("expected countryCode=US")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(fixture)
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	currencies, err := svc.CurrencyByCountry(context.Background(), "US", "en")
	if err != nil {
		t.Fatal(err)
	}
	if currencies[0].CurrencyCode != "USD" {
		t.Errorf("expected USD, got %s", currencies[0].CurrencyCode)
	}
	if currencies[0].Prefix != "$" {
		t.Errorf("expected $, got %s", currencies[0].Prefix)
	}
}

func TestStates(t *testing.T) {
	fixture, err := os.ReadFile("testdata/states.json")
	if err != nil {
		t.Fatal(err)
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("countryCode") != "US" {
			t.Error("expected countryCode=US")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(fixture)
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	states, err := svc.States(context.Background(), "US", "")
	if err != nil {
		t.Fatal(err)
	}
	if len(states) != 2 {
		t.Fatalf("expected 2 states, got %d", len(states))
	}
}

func TestTaxClasses(t *testing.T) {
	fixture, err := os.ReadFile("testdata/tax_classes.json")
	if err != nil {
		t.Fatal(err)
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("countryCode") != "US" {
			t.Error("expected countryCode=US")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(fixture)
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	taxClasses, err := svc.TaxClasses(context.Background(), "US", "en")
	if err != nil {
		t.Fatal(err)
	}
	if len(taxClasses) != 2 {
		t.Fatalf("expected 2 tax classes, got %d", len(taxClasses))
	}
	if taxClasses[0].TaxClassCode != "default" {
		t.Errorf("expected default, got %s", taxClasses[0].TaxClassCode)
	}
}

func TestCountries_Error(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"errorMessage":"internal error","errorCode":500}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	_, err := svc.Countries(context.Background(), nil)
	if err == nil {
		t.Fatal("expected error")
	}
}
