package profile_test

import (
	"context"
	"encoding/json"
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

func TestGet_TaxFields(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"taxes": [
				{"id": 7, "name": "MwSt.", "enabled": true, "includeInPrice": true, "defaultTax": 0,
				 "rules": [{"zoneId": "zone-de", "tax": 19}, {"zoneId": "zone-eu", "tax": 7.5}]}
			],
			"taxSettings": {"automaticTaxEnabled": true, "pricesIncludeTax": true, "taxExemptBusiness": false},
			"zones": [{"id": "zone-de", "name": "Deutschland", "countryCodes": ["DE"]}]
		}`))
	}))
	defer srv.Close()

	p, err := newTestService(t, srv).Get(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(p.Taxes) != 1 || p.Taxes[0].ID != 7 {
		t.Fatalf("expected one top-level tax with id 7, got %+v", p.Taxes)
	}
	if len(p.Taxes[0].Rules) != 2 || p.Taxes[0].Rules[0].ZoneID != "zone-de" || p.Taxes[0].Rules[0].Tax != 19 {
		t.Errorf("expected typed tax rules, got %+v", p.Taxes[0].Rules)
	}
	if p.Taxes[0].Rules[1].Tax != 7.5 {
		t.Errorf("expected fractional rule tax 7.5, got %v", p.Taxes[0].Rules[1].Tax)
	}
	if p.TaxSettings == nil || p.TaxSettings.PricesIncludeTax == nil || !*p.TaxSettings.PricesIncludeTax {
		t.Error("expected taxSettings.pricesIncludeTax=true")
	}
	if p.TaxSettings.TaxExemptBusiness == nil || *p.TaxSettings.TaxExemptBusiness {
		t.Error("expected taxSettings.taxExemptBusiness=false")
	}
}

func TestUpdate_TaxRoundTrip(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var got profile.UpdateRequest
		if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
			t.Fatalf("decode update body: %v", err)
		}
		if len(got.Taxes) != 1 || got.Taxes[0].Name != "Kleinunternehmer" {
			t.Errorf("expected top-level tax to round-trip, got %+v", got.Taxes)
		}
		if len(got.Taxes[0].Rules) != 1 || got.Taxes[0].Rules[0].ZoneID != "zone-de" {
			t.Errorf("expected tax rules to round-trip, got %+v", got.Taxes[0].Rules)
		}
		if got.TaxSettings == nil || got.TaxSettings.TaxExemptBusiness == nil || !*got.TaxSettings.TaxExemptBusiness {
			t.Error("expected taxSettings.taxExemptBusiness=true to round-trip")
		}
		if len(got.Zones) != 1 || got.Zones[0].ID != "zone-de" {
			t.Errorf("expected zones to round-trip, got %+v", got.Zones)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	enabled := true
	exempt := true
	_, err := newTestService(t, srv).Update(context.Background(), &profile.UpdateRequest{
		Taxes: []profile.Tax{{
			Name:    "Kleinunternehmer",
			Enabled: &enabled,
			Rules:   []profile.TaxRule{{ZoneID: "zone-de", Tax: 0}},
		}},
		TaxSettings: &profile.TaxSettings{TaxExemptBusiness: &exempt},
		Zones:       []profile.Zone{{ID: "zone-de", Name: "Deutschland", CountryCodes: []string{"DE"}}},
	})
	if err != nil {
		t.Fatal(err)
	}
}
