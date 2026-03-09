// Package dictionaries provides access to the Ecwid dictionaries API (read-only reference data).
package dictionaries

import "encoding/json"

// Country represents a country returned by the Ecwid dictionaries API.
type Country struct {
	Code   string  `json:"code"`
	Name   string  `json:"name"`
	States []State `json:"states,omitempty"`
}

// State represents a state or region within a country.
type State struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Currency represents a currency returned by the Ecwid dictionaries API.
type Currency struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// CountryCurrency represents a currency associated with a specific country.
type CountryCurrency struct {
	CountryCode   string `json:"countryCode"`
	CurrencyCode  string `json:"currencyCode"`
	CurrencyName  string `json:"currencyName"`
	Prefix        string `json:"prefix"`
	Suffix        string `json:"suffix"`
	DecimalPlaces string `json:"decimalPlaces"`
}

// TaxClass represents a tax class returned by the Ecwid dictionaries API.
type TaxClass struct {
	StateCode    string          `json:"stateCode"`
	TaxClassCode string          `json:"taxClassCode"`
	TaxClassRate string          `json:"taxClassRate"`
	Localization json.RawMessage `json:"localization,omitempty"`
}

// CountriesOptions holds optional parameters for the Countries request.
type CountriesOptions struct {
	// Lang is the language ISO code (e.g. "en", "de"). Default: "en".
	Lang string
	// WithStates adds the states field for each country in the response.
	WithStates bool
}
