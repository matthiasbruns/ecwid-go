package e2e

import (
	"context"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/dictionaries"
)

func TestDictionaries_Countries(t *testing.T) {
	ctx := context.Background()

	countries, err := testClient.Dictionaries.Countries(ctx, nil)
	if err != nil {
		t.Fatalf("Countries: %v", err)
	}
	if len(countries) == 0 {
		t.Fatal("expected at least one country")
	}

	// Spot-check a well-known country.
	found := false
	for _, c := range countries {
		if c.Code == "US" {
			found = true
			if c.Name == "" {
				t.Error("US country name is empty")
			}
			break
		}
	}
	if !found {
		t.Error("US not found in countries list")
	}
}

func TestDictionaries_CountriesWithStates(t *testing.T) {
	ctx := context.Background()

	countries, err := testClient.Dictionaries.Countries(ctx, &dictionaries.CountriesOptions{
		WithStates: true,
		Lang:       "en",
	})
	if err != nil {
		t.Fatalf("Countries(withStates): %v", err)
	}

	// US should have states when withStates=true.
	for _, c := range countries {
		if c.Code == "US" {
			if len(c.States) == 0 {
				t.Error("US has no states with withStates=true")
			}
			return
		}
	}
	t.Error("US not found")
}

func TestDictionaries_Currencies(t *testing.T) {
	ctx := context.Background()

	currencies, err := testClient.Dictionaries.Currencies(ctx, "en")
	if err != nil {
		t.Fatalf("Currencies: %v", err)
	}
	if len(currencies) == 0 {
		t.Fatal("expected at least one currency")
	}

	found := false
	for _, c := range currencies {
		if c.Code == "USD" {
			found = true
			break
		}
	}
	if !found {
		t.Error("USD not found in currencies list")
	}
}

func TestDictionaries_CurrencyByCountry(t *testing.T) {
	ctx := context.Background()

	currencies, err := testClient.Dictionaries.CurrencyByCountry(ctx, "DE", "en")
	if err != nil {
		t.Fatalf("CurrencyByCountry(DE): %v", err)
	}
	if len(currencies) == 0 {
		t.Fatal("expected currency for DE")
	}
	if currencies[0].CurrencyCode != "EUR" {
		t.Errorf("expected EUR for DE, got %s", currencies[0].CurrencyCode)
	}
}

func TestDictionaries_States(t *testing.T) {
	ctx := context.Background()

	states, err := testClient.Dictionaries.States(ctx, "US", "en")
	if err != nil {
		t.Fatalf("States(US): %v", err)
	}
	if len(states) == 0 {
		t.Fatal("expected states for US")
	}

	found := false
	for _, s := range states {
		if s.Code == "CA" {
			found = true
			break
		}
	}
	if !found {
		t.Error("CA not found in US states")
	}
}

func TestDictionaries_TaxClasses(t *testing.T) {
	ctx := context.Background()

	taxClasses, err := testClient.Dictionaries.TaxClasses(ctx, "US", "en")
	if err != nil {
		t.Fatalf("TaxClasses(US): %v", err)
	}
	// Tax classes may be empty for some countries, just verify no error.
	_ = taxClasses
}
