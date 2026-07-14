// Package config defines the configuration struct and validation for the Ecwid API client and CLI.
package config

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	// DefaultBaseURL is the default Ecwid API base URL.
	DefaultBaseURL = "https://app.ecwid.com/api/v3"

	// DefaultOutput is the default output format.
	DefaultOutput = "json"

	// DefaultLogLevel is the default log level.
	DefaultLogLevel = "info"

	// DefaultInstantSiteBaseURL is the default base URL for the Instant Site v1
	// API (profile, pages, tiles, themes, text labels, publish/discard/clone).
	//
	// NOTE: this host is unverified against a live store — the Ecwid docs
	// consistently render "vuega.ecwid.com" for these v1 endpoints, and the
	// profile response's ecwidApiUrl field hints the host may be store-specific.
	// Override via InstantSiteBaseURL if a live call reveals a different host.
	DefaultInstantSiteBaseURL = "https://vuega.ecwid.com/api/v1"

	// DefaultInstantSiteAuthURL is the default base URL for the Instant Site
	// OAuth token-exchange endpoint (POST /oauth/token).
	DefaultInstantSiteAuthURL = "https://app.ecwid.com/instantsite"
)

// Config holds all configuration for the Ecwid client and CLI.
type Config struct {
	// StoreID is the Ecwid store ID.
	StoreID string `json:"store_id" yaml:"store_id" mapstructure:"store_id"`

	// Token is the API access token.
	Token string `json:"token" yaml:"token" mapstructure:"token"`

	// BaseURL is the API base URL. Defaults to DefaultBaseURL.
	BaseURL string `json:"base_url,omitempty" yaml:"base_url,omitempty" mapstructure:"base_url"`

	// Output format: "json" or "table". Defaults to "json".
	Output string `json:"output,omitempty" yaml:"output,omitempty" mapstructure:"output"`

	// LogLevel: "debug", "info", "warn", "error". Defaults to "info".
	LogLevel string `json:"log_level,omitempty" yaml:"log_level,omitempty" mapstructure:"log_level"`

	// MaxRetries is the maximum number of retries for rate-limited requests.
	// 0 means no retries. Defaults to 0.
	MaxRetries int `json:"max_retries,omitempty" yaml:"max_retries,omitempty" mapstructure:"max_retries"`

	// InstantSiteBaseURL is the base URL for the Instant Site v1 API.
	// Defaults to DefaultInstantSiteBaseURL.
	InstantSiteBaseURL string `json:"instant_site_base_url,omitempty" yaml:"instant_site_base_url,omitempty" mapstructure:"instant_site_base_url"`

	// InstantSiteAuthURL is the base URL for the Instant Site OAuth token
	// exchange. Defaults to DefaultInstantSiteAuthURL.
	InstantSiteAuthURL string `json:"instant_site_auth_url,omitempty" yaml:"instant_site_auth_url,omitempty" mapstructure:"instant_site_auth_url"`

	// InstantSiteToken is the separate access token for the Instant Site v1 API,
	// obtained via the token-exchange endpoint (24h lifetime). Empty by default;
	// v1 Instant Site calls require it.
	InstantSiteToken string `json:"instant_site_token,omitempty" yaml:"instant_site_token,omitempty" mapstructure:"instant_site_token"`
}

// Validate checks that required fields are present and values are valid.
func (c *Config) Validate() error {
	var errs []string

	if c.StoreID == "" {
		errs = append(errs, "store_id is required")
	}
	if c.Token == "" {
		errs = append(errs, "token is required")
	}

	if c.Output != "" && c.Output != "json" && c.Output != "table" {
		errs = append(errs, fmt.Sprintf("invalid output format %q (must be json or table)", c.Output))
	}

	validLevels := map[string]bool{"debug": true, "info": true, "warn": true, "error": true, "": true}
	if !validLevels[c.LogLevel] {
		errs = append(errs, fmt.Sprintf("invalid log level %q (must be debug, info, warn, or error)", c.LogLevel))
	}

	if c.MaxRetries < 0 {
		errs = append(errs, "max_retries must be >= 0")
	}

	if len(errs) > 0 {
		return fmt.Errorf("config validation: %s", strings.Join(errs, "; "))
	}

	return nil
}

// WithDefaults returns a copy of the config with default values applied for empty fields.
func (c Config) WithDefaults() Config {
	if c.BaseURL == "" {
		c.BaseURL = DefaultBaseURL
	}
	if c.Output == "" {
		c.Output = DefaultOutput
	}
	if c.LogLevel == "" {
		c.LogLevel = DefaultLogLevel
	}
	if c.InstantSiteBaseURL == "" {
		c.InstantSiteBaseURL = DefaultInstantSiteBaseURL
	}
	if c.InstantSiteAuthURL == "" {
		c.InstantSiteAuthURL = DefaultInstantSiteAuthURL
	}
	return c
}

// RedactedInstantSiteToken returns the Instant Site token with all but the last
// 4 characters masked. Use this for logging — never log the full token.
func (c *Config) RedactedInstantSiteToken() string {
	if c.InstantSiteToken == "" {
		return ""
	}
	if len(c.InstantSiteToken) <= 4 {
		return "****"
	}
	return strings.Repeat("*", len(c.InstantSiteToken)-4) + c.InstantSiteToken[len(c.InstantSiteToken)-4:]
}

// RedactedToken returns the token with all but the last 4 characters masked.
// Use this for logging — never log the full token.
func (c *Config) RedactedToken() string {
	if len(c.Token) <= 4 {
		return "****"
	}
	return strings.Repeat("*", len(c.Token)-4) + c.Token[len(c.Token)-4:]
}

// MarshalJSON implements custom JSON marshaling that redacts secret tokens.
func (c Config) MarshalJSON() ([]byte, error) {
	type alias Config
	safe := struct {
		alias
		Token            string `json:"token"`
		InstantSiteToken string `json:"instant_site_token,omitempty"`
	}{
		alias:            alias(c),
		Token:            c.RedactedToken(),
		InstantSiteToken: c.RedactedInstantSiteToken(),
	}
	return json.Marshal(safe)
}
