// Package ecwid provides a Go client for the Ecwid REST API.
//
// Create a client with [NewClient], then access domain services:
//
//	cfg := config.Config{StoreID: "12345", Token: "secret_xxx"}
//	client := ecwid.NewClient(cfg)
//	countries, err := client.Dictionaries.Countries(ctx, nil)
package ecwid

import (
	"log/slog"
	"net/http"

	"github.com/matthiasbruns/ecwid-go/config"
	"github.com/matthiasbruns/ecwid-go/ecwid/billing"
	"github.com/matthiasbruns/ecwid-go/ecwid/carts"
	"github.com/matthiasbruns/ecwid-go/ecwid/categories"
	"github.com/matthiasbruns/ecwid-go/ecwid/customers"
	"github.com/matthiasbruns/ecwid-go/ecwid/dictionaries"
	"github.com/matthiasbruns/ecwid-go/ecwid/discounts"
	"github.com/matthiasbruns/ecwid-go/ecwid/domains"
	"github.com/matthiasbruns/ecwid-go/ecwid/instantsite"
	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
	"github.com/matthiasbruns/ecwid-go/ecwid/orders"
	"github.com/matthiasbruns/ecwid-go/ecwid/products"
	"github.com/matthiasbruns/ecwid-go/ecwid/profile"
	"github.com/matthiasbruns/ecwid-go/ecwid/reports"
	"github.com/matthiasbruns/ecwid-go/ecwid/reviews"
	"github.com/matthiasbruns/ecwid-go/ecwid/staff"
	"github.com/matthiasbruns/ecwid-go/ecwid/subscriptions"
)

// Client is the Ecwid API client. It is safe for concurrent use.
type Client struct {
	Billing       *billing.Service
	Carts         *carts.Service
	Categories    *categories.Service
	Customers     *customers.Service
	Dictionaries  *dictionaries.Service
	Discounts     *discounts.Service
	Domains       *domains.Service
	InstantSite   *instantsite.Service
	Orders        *orders.Service
	Products      *products.Service
	Profile       *profile.Service
	Reports       *reports.Service
	Reviews       *reviews.Service
	Staff         *staff.Service
	Subscriptions *subscriptions.Service
}

// options holds optional configuration for the Client.
type options struct {
	httpClient *http.Client
	logger     *slog.Logger
}

// Option configures the Client.
type Option func(*options)

// WithHTTPClient sets a custom HTTP client. A nil value is ignored.
func WithHTTPClient(c *http.Client) Option {
	return func(o *options) {
		if c != nil {
			o.httpClient = c
		}
	}
}

// WithLogger sets a custom slog logger. A nil value is ignored. By default the
// client logs nothing (a no-op handler); pass a logger to opt into request
// (debug) and rate-limit retry (warn) logging. Credentials are never logged.
func WithLogger(l *slog.Logger) Option {
	return func(o *options) {
		if l != nil {
			o.logger = l
		}
	}
}

// NewClient creates a new Ecwid API client with the given configuration.
//
// The client applies defaults to the config (BaseURL, Output, LogLevel)
// but does NOT validate it — call [config.Config.Validate] beforehand if needed.
func NewClient(cfg config.Config, opts ...Option) *Client {
	cfg = cfg.WithDefaults()

	o := &options{}
	for _, opt := range opts {
		opt(o)
	}

	requester := api.NewHTTPClient(api.HTTPClientConfig{
		BaseURL:    cfg.BaseURL,
		StoreID:    cfg.StoreID,
		Token:      cfg.Token,
		MaxRetries: cfg.MaxRetries,
		HTTPClient: o.httpClient,
		Logger:     o.logger,
	})

	// The Instant Site v1 endpoints live on a separate host and use a separate
	// token; the token-exchange endpoint uses the auth host with no store-ID
	// path segment and no bearer token.
	instantSiteV1 := api.NewHTTPClient(api.HTTPClientConfig{
		BaseURL:    cfg.InstantSiteBaseURL,
		StoreID:    cfg.StoreID,
		Token:      cfg.InstantSiteToken,
		MaxRetries: cfg.MaxRetries,
		HTTPClient: o.httpClient,
		Logger:     o.logger,
	})
	instantSiteAuth := api.NewHTTPClient(api.HTTPClientConfig{
		BaseURL:    cfg.InstantSiteAuthURL,
		StoreID:    "",
		Token:      "",
		MaxRetries: cfg.MaxRetries,
		HTTPClient: o.httpClient,
		Logger:     o.logger,
	})

	return &Client{
		Billing:       billing.NewService(requester),
		Carts:         carts.NewService(requester),
		Categories:    categories.NewService(requester),
		Customers:     customers.NewService(requester),
		Dictionaries:  dictionaries.NewService(requester),
		Discounts:     discounts.NewService(requester),
		Domains:       domains.NewService(requester),
		InstantSite:   instantsite.NewService(requester, instantSiteV1, instantSiteAuth),
		Orders:        orders.NewService(requester),
		Products:      products.NewService(requester),
		Profile:       profile.NewService(requester),
		Reports:       reports.NewService(requester),
		Reviews:       reviews.NewService(requester),
		Staff:         staff.NewService(requester),
		Subscriptions: subscriptions.NewService(requester),
	}
}
