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
	return c
}

// RedactedToken returns the token with all but the last 4 characters masked.
// Use this for logging — never log the full token.
func (c *Config) RedactedToken() string {
	if len(c.Token) <= 4 {
		return "****"
	}
	return strings.Repeat("*", len(c.Token)-4) + c.Token[len(c.Token)-4:]
}

// MarshalJSON implements custom JSON marshaling that redacts the token.
func (c Config) MarshalJSON() ([]byte, error) {
	type alias Config
	safe := struct {
		alias
		Token string `json:"token"`
	}{
		alias: alias(c),
		Token: c.RedactedToken(),
	}
	return json.Marshal(safe)
}
