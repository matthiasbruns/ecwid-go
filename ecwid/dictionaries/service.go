package dictionaries

import (
	"context"
	"net/url"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

// Service provides access to the Ecwid dictionaries API.
type Service struct {
	requester api.Requester
}

// NewService creates a new dictionaries service.
func NewService(requester api.Requester) *Service {
	return &Service{requester: requester}
}

// Countries returns a list of all countries.
//
// API: GET /countries
func (s *Service) Countries(ctx context.Context, opts *CountriesOptions) ([]Country, error) {
	q := url.Values{}
	if opts != nil {
		if opts.Lang != "" {
			q.Set("lang", opts.Lang)
		}
		if opts.WithStates {
			q.Set("withStates", "true")
		}
	}

	var result []Country
	if err := s.requester.Get(ctx, "/countries", q, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// Currencies returns a list of all currencies.
//
// API: GET /currencies
func (s *Service) Currencies(ctx context.Context, lang string) ([]Currency, error) {
	q := url.Values{}
	if lang != "" {
		q.Set("lang", lang)
	}

	var result []Currency
	if err := s.requester.Get(ctx, "/currencies", q, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CurrencyByCountry returns the currency for a specific country.
//
// API: GET /currencyByCountry
func (s *Service) CurrencyByCountry(ctx context.Context, countryCode, lang string) ([]CountryCurrency, error) {
	q := url.Values{}
	q.Set("countryCode", countryCode)
	if lang != "" {
		q.Set("lang", lang)
	}

	var result []CountryCurrency
	if err := s.requester.Get(ctx, "/currencyByCountry", q, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// States returns a list of states for a specific country.
//
// API: GET /states
func (s *Service) States(ctx context.Context, countryCode, lang string) ([]State, error) {
	q := url.Values{}
	q.Set("countryCode", countryCode)
	if lang != "" {
		q.Set("lang", lang)
	}

	var result []State
	if err := s.requester.Get(ctx, "/states", q, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// TaxClasses returns a list of tax classes for a specific country.
//
// API: GET /taxClasses
func (s *Service) TaxClasses(ctx context.Context, countryCode, lang string) ([]TaxClass, error) {
	q := url.Values{}
	q.Set("countryCode", countryCode)
	if lang != "" {
		q.Set("lang", lang)
	}

	var result []TaxClass
	if err := s.requester.Get(ctx, "/taxClasses", q, &result); err != nil {
		return nil, err
	}
	return result, nil
}
