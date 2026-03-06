// Package config provides configuration loading for the Ecwid API client and CLI.
//
// Configuration is loaded with the following precedence (highest to lowest):
//  1. Explicit values passed programmatically
//  2. Environment variables (ECWID_STORE_ID, ECWID_TOKEN, etc.)
//  3. Config file (~/.ecwid.yaml)
//  4. Default values
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	// DefaultBaseURL is the default Ecwid API base URL.
	DefaultBaseURL = "https://app.ecwid.com/api/v3"

	// DefaultConfigFile is the default config file name in the user's home directory.
	DefaultConfigFile = ".ecwid.yaml"

	// DefaultOutput is the default output format.
	DefaultOutput = "json"

	// DefaultLogLevel is the default log level.
	DefaultLogLevel = "info"
)

// Environment variable names.
const (
	EnvStoreID  = "ECWID_STORE_ID"
	EnvToken    = "ECWID_TOKEN"
	EnvBaseURL  = "ECWID_BASE_URL"
	EnvOutput   = "ECWID_OUTPUT"
	EnvLogLevel = "ECWID_LOG_LEVEL"
)

// Config holds all configuration for the Ecwid client and CLI.
type Config struct {
	// StoreID is the Ecwid store ID.
	StoreID string `json:"store_id" yaml:"store_id"`

	// Token is the API access token.
	Token string `json:"token" yaml:"token"`

	// BaseURL is the API base URL. Defaults to DefaultBaseURL.
	BaseURL string `json:"base_url,omitempty" yaml:"base_url,omitempty"`

	// Output format: "json" or "table". Defaults to "json".
	Output string `json:"output,omitempty" yaml:"output,omitempty"`

	// LogLevel: "debug", "info", "warn", "error". Defaults to "info".
	LogLevel string `json:"log_level,omitempty" yaml:"log_level,omitempty"`

	// MaxRetries is the maximum number of retries for rate-limited requests.
	// 0 means no retries. Defaults to 0.
	MaxRetries int `json:"max_retries,omitempty" yaml:"max_retries,omitempty"`
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

// LoadFromEnv returns a Config populated from environment variables.
// Only non-empty environment variables are set.
func LoadFromEnv() Config {
	var cfg Config

	if v := os.Getenv(EnvStoreID); v != "" {
		cfg.StoreID = v
	}
	if v := os.Getenv(EnvToken); v != "" {
		cfg.Token = v
	}
	if v := os.Getenv(EnvBaseURL); v != "" {
		cfg.BaseURL = v
	}
	if v := os.Getenv(EnvOutput); v != "" {
		cfg.Output = v
	}
	if v := os.Getenv(EnvLogLevel); v != "" {
		cfg.LogLevel = v
	}
	if v := os.Getenv("ECWID_MAX_RETRIES"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			cfg.MaxRetries = n
		}
	}

	return cfg
}

// LoadFromFile reads a config file in a simple YAML-like key: value format.
// This uses a minimal parser to avoid external dependencies.
// Supported keys: store_id, token, base_url, output, log_level, max_retries.
func LoadFromFile(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return Config{}, nil
		}
		return Config{}, fmt.Errorf("read config file: %w", err)
	}

	return parseSimpleYAML(data)
}

// LoadFromDefaultFile loads config from ~/.ecwid.yaml if it exists.
func LoadFromDefaultFile() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, nil // Silently skip if home dir is unknown
	}
	return LoadFromFile(filepath.Join(home, DefaultConfigFile))
}

// Merge returns a new Config where non-zero values from overlay take precedence over base.
func Merge(base, overlay Config) Config {
	if overlay.StoreID != "" {
		base.StoreID = overlay.StoreID
	}
	if overlay.Token != "" {
		base.Token = overlay.Token
	}
	if overlay.BaseURL != "" {
		base.BaseURL = overlay.BaseURL
	}
	if overlay.Output != "" {
		base.Output = overlay.Output
	}
	if overlay.LogLevel != "" {
		base.LogLevel = overlay.LogLevel
	}
	if overlay.MaxRetries != 0 {
		base.MaxRetries = overlay.MaxRetries
	}
	return base
}

// Load loads configuration with full precedence: file < env < explicit overrides.
// The configPath is optional; if empty, ~/.ecwid.yaml is used.
func Load(configPath string, overrides Config) (Config, error) {
	// 1. Load from file
	var fileCfg Config
	var err error
	if configPath != "" {
		fileCfg, err = LoadFromFile(configPath)
	} else {
		fileCfg, err = LoadFromDefaultFile()
	}
	if err != nil {
		return Config{}, err
	}

	// 2. Load from env
	envCfg := LoadFromEnv()

	// 3. Merge: file < env < overrides
	cfg := Merge(fileCfg, envCfg)
	cfg = Merge(cfg, overrides)

	// 4. Apply defaults
	cfg = cfg.WithDefaults()

	return cfg, nil
}

// RedactedToken returns the token with all but the last 4 characters masked.
// Use this for logging — never log the full token.
func (c *Config) RedactedToken() string {
	if len(c.Token) <= 4 {
		return "****"
	}
	return strings.Repeat("*", len(c.Token)-4) + c.Token[len(c.Token)-4:]
}

// parseSimpleYAML parses a minimal subset of YAML (key: value pairs, one per line).
// This avoids pulling in a YAML library for the config module.
func parseSimpleYAML(data []byte) (Config, error) {
	var cfg Config
	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return Config{}, fmt.Errorf("config line %d: invalid format %q", i+1, line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		// Strip surrounding quotes
		value = strings.Trim(value, `"'`)

		switch key {
		case "store_id":
			cfg.StoreID = value
		case "token":
			cfg.Token = value
		case "base_url":
			cfg.BaseURL = value
		case "output":
			cfg.Output = value
		case "log_level":
			cfg.LogLevel = value
		case "max_retries":
			n, err := strconv.Atoi(value)
			if err != nil {
				return Config{}, fmt.Errorf("config line %d: invalid max_retries %q: %w", i+1, value, err)
			}
			cfg.MaxRetries = n
		default:
			// Ignore unknown keys for forward compatibility
		}
	}

	return cfg, nil
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
