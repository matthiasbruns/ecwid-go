package e2e

import (
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/dictionaries"
)

func TestDictionaries_Countries(t *testing.T) {
	ctx := testContext(t)

	countries, err := testClient.Dictionaries.Countries(ctx, nil)
	if err != nil {
		t.Fatalf("Countries: %v", err)
	}
	if len(countries) == 0 {
		t.Fatal("expected at least one country")
	}

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
	ctx := testContext(t)

	countries, err := testClient.Dictionaries.Countries(ctx, &dictionaries.CountriesOptions{
		WithStates: true,
		Lang:       "en",
	})
	if err != nil {
		t.Fatalf("Countries(withStates): %v", err)
	}

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
	ctx := testContext(t)

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
	ctx := testContext(t)

	currencies, err := testClient.Dictionaries.CurrencyByCountry(ctx, "DE", "en")
	if err != nil {
		t.Fatalf("CurrencyByCountry(DE): %v", err)
	}
	if len(currencies) == 0 {
		t.Fatal("expected currency for DE")
	}

	found := false
	for _, c := range currencies {
		if c.CurrencyCode == "EUR" {
			found = true
			break
		}
	}
	if !found {
		t.Error("EUR not found in currencies for DE")
	}
}

func TestDictionaries_States(t *testing.T) {
	ctx := testContext(t)

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
	ctx := testContext(t)

	taxClasses, err := testClient.Dictionaries.TaxClasses(ctx, "US", "en")
	if err != nil {
		t.Fatalf("TaxClasses(US): %v", err)
	}
	_ = taxClasses
}
