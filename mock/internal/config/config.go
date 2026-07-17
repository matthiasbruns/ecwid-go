// Package config holds the mock server's configuration.
//
// The mock is an *app developer's* tool: it stands in for a real Ecwid store so
// a Go developer can exercise their embedded app locally. Its configuration is
// therefore deliberately different in shape from the top-level config.Config
// (which carries StoreID/Token for real API calls). Those two must not be
// merged — config.Config is a published compatibility contract and these
// developer-tool fields do not belong to API clients.
package config

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	appconfig "github.com/matthiasbruns/ecwid-go/config"
	"github.com/matthiasbruns/ecwid-go/ecwid/appauth"
)

const (
	// DefaultClientID is the default app client_id, also used as
	// EcwidApp.init({app_id}).
	DefaultClientID = "mock-app"

	// DefaultStoreID is the default store ID placed in the auth payload.
	DefaultStoreID = "1003"

	// DefaultAuthMode is the default iframe auth mode.
	DefaultAuthMode = AuthModeDefault

	// DefaultPort is the default listen port.
	DefaultPort = 8080

	// AuthModeDefault is Default User Auth: a hex-encoded plaintext payload in
	// the URL fragment (see ecwid/appauth).
	AuthModeDefault = "default"

	// AuthModeEnhanced is Enhanced Security User Auth: an AES-encrypted payload
	// in the query string (see ecwid/appauth).
	AuthModeEnhanced = "enhanced"

	// MinClientSecretLen is the minimum client_secret length in bytes. appauth
	// derives an AES-128 key from the first 16 bytes of the secret, so anything
	// shorter cannot sign webhooks or encrypt payloads.
	MinClientSecretLen = 16

	// generatedSecretBytes is the number of random bytes used when generating a
	// client_secret. Hex-encoding doubles the length, comfortably clearing
	// MinClientSecretLen.
	generatedSecretBytes = 16
)

// Config is the mock server's runtime configuration.
type Config struct {
	// AppURL is the URL of the app to iframe (e.g. http://localhost:3000).
	// Required.
	AppURL string

	// ClientID is the app client_id, also used as EcwidApp.init({app_id}).
	ClientID string

	// ClientSecret signs webhooks and encrypts payloads. Must be at least
	// MinClientSecretLen bytes. Generated if not provided.
	ClientSecret string

	// StoreID is the store ID placed in the auth payload.
	StoreID string

	// AuthMode selects the iframe auth mode: AuthModeDefault or AuthModeEnhanced.
	AuthMode string

	// WebhookURL is where triggered webhooks POST. Optional.
	WebhookURL string

	// Port is the listen port.
	Port int

	// ProxyStore is the store ID to forward unimplemented REST calls to.
	// Optional.
	ProxyStore string

	// ProxyToken is the access token for the proxy store. Optional.
	ProxyToken string

	// SecretGenerated reports whether ClientSecret was generated rather than
	// supplied. When true, the startup banner prints it so the developer can
	// configure their app to match.
	SecretGenerated bool
}

// RedactedClientSecret returns the client_secret with all but the last 4
// characters masked. Never log the raw secret; use this for diagnostics.
func (c *Config) RedactedClientSecret() string {
	return redact(c.ClientSecret)
}

// RedactedProxyToken returns the proxy token with all but the last 4 characters
// masked, or "" when no proxy token is configured.
func (c *Config) RedactedProxyToken() string {
	if c.ProxyToken == "" {
		return ""
	}
	return redact(c.ProxyToken)
}

// redact masks a secret for logging, reusing config.RedactedToken so the mock
// and the API client mask credentials identically.
func redact(secret string) string {
	rc := appconfig.Config{Token: secret}
	return rc.RedactedToken()
}

// Validate checks required fields and value constraints.
func (c *Config) Validate() error {
	var errs []string

	if c.AppURL == "" {
		errs = append(errs, "--app-url is required (env: ECWID_MOCK_APP_URL)")
	} else if u, err := url.Parse(c.AppURL); err != nil || u.Scheme == "" || u.Host == "" {
		errs = append(errs, fmt.Sprintf("--app-url %q is not a valid absolute URL", c.AppURL))
	}

	if len(c.ClientSecret) < MinClientSecretLen {
		errs = append(errs, fmt.Sprintf(
			"--client-secret must be at least %d bytes for AES-128 (%v), got %d",
			MinClientSecretLen, appauth.ErrShortSecret, len(c.ClientSecret)))
	}

	if c.AuthMode != AuthModeDefault && c.AuthMode != AuthModeEnhanced {
		errs = append(errs, fmt.Sprintf(
			"--auth-mode %q is invalid (must be %q or %q)", c.AuthMode, AuthModeDefault, AuthModeEnhanced))
	}

	if c.Port < 1 || c.Port > 65535 {
		errs = append(errs, fmt.Sprintf("--port %d is out of range (1-65535)", c.Port))
	}

	if len(errs) > 0 {
		return fmt.Errorf("config validation: %s", strings.Join(errs, "; "))
	}
	return nil
}

// Load resolves configuration with precedence flags > env > defaults, matching
// the config module's behavior. cmd supplies the parsed flags; a flag counts
// only when explicitly set (Changed), so an unset flag falls through to env
// then default. When no client_secret is supplied one is generated and
// SecretGenerated is set.
func Load(cmd *cobra.Command) (Config, error) {
	cfg := Config{
		ClientID: DefaultClientID,
		StoreID:  DefaultStoreID,
		AuthMode: DefaultAuthMode,
		Port:     DefaultPort,
	}

	cfg.AppURL = resolveString(cmd, "app-url", "ECWID_MOCK_APP_URL", "")
	cfg.ClientID = resolveString(cmd, "client-id", "ECWID_MOCK_CLIENT_ID", DefaultClientID)
	cfg.ClientSecret = resolveString(cmd, "client-secret", "ECWID_MOCK_CLIENT_SECRET", "")
	cfg.StoreID = resolveString(cmd, "store-id", "ECWID_MOCK_STORE_ID", DefaultStoreID)
	cfg.AuthMode = resolveString(cmd, "auth-mode", "ECWID_MOCK_AUTH_MODE", DefaultAuthMode)
	cfg.WebhookURL = resolveString(cmd, "webhook-url", "ECWID_MOCK_WEBHOOK_URL", "")
	cfg.ProxyStore = resolveString(cmd, "proxy-store", "ECWID_MOCK_PROXY_STORE", "")
	cfg.ProxyToken = resolveString(cmd, "proxy-token", "ECWID_MOCK_PROXY_TOKEN", "")

	port, err := resolvePort(cmd)
	if err != nil {
		return Config{}, err
	}
	cfg.Port = port

	if cfg.ClientSecret == "" {
		secret, err := generateSecret()
		if err != nil {
			return Config{}, fmt.Errorf("generate client_secret: %w", err)
		}
		cfg.ClientSecret = secret
		cfg.SecretGenerated = true
	}

	return cfg, nil
}

// resolveString applies flags > env > default for a string value.
func resolveString(cmd *cobra.Command, flag, env, def string) string {
	if cmd != nil {
		if f := cmd.Flags().Lookup(flag); f != nil && f.Changed {
			return f.Value.String()
		}
	}
	if v, ok := os.LookupEnv(env); ok {
		return v
	}
	return def
}

// resolvePort applies flags > env > default for the listen port.
func resolvePort(cmd *cobra.Command) (int, error) {
	if cmd != nil {
		if f := cmd.Flags().Lookup("port"); f != nil && f.Changed {
			p, err := cmd.Flags().GetInt("port")
			if err != nil {
				return 0, fmt.Errorf("parse --port: %w", err)
			}
			return p, nil
		}
	}
	if v, ok := os.LookupEnv("ECWID_MOCK_PORT"); ok {
		p, err := strconv.Atoi(v)
		if err != nil {
			return 0, fmt.Errorf("parse ECWID_MOCK_PORT %q: %w", v, err)
		}
		return p, nil
	}
	return DefaultPort, nil
}

// generateSecret returns a hex-encoded random client_secret of at least
// MinClientSecretLen bytes.
func generateSecret() (string, error) {
	buf := make([]byte, generatedSecretBytes)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}
