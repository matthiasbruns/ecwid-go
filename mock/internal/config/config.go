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
	"slices"
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

	// DefaultProxyReadonly is the default for --proxy-readonly. It is true on
	// purpose: a proxied write mutates a REAL store and fires REAL webhooks from
	// it, so the mock refuses mutations until the developer explicitly opts in.
	DefaultProxyReadonly = true

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

	// generatedTokenBytes is the number of random bytes used when generating an
	// access_token. Hex-encoding doubles the resulting string length.
	generatedTokenBytes = 16
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

	// AccessToken is the access_token the mock requires as the Bearer credential
	// on simulated REST calls. Generated if not supplied. The admin shell (#4)
	// will inject this same value into the iframe payload so the SDK sends it
	// back automatically; until then it is supplied to the app out of band.
	AccessToken string

	// ProxyReadonly gates proxied mutations. When true (the default), only
	// GET/HEAD are forwarded and write methods return 403 — proxied writes
	// mutate a REAL store and fire REAL webhooks from it, so the safe default is
	// "cannot wreck your store". Set false to allow mutating proxy requests.
	ProxyReadonly bool

	// SecretGenerated reports whether ClientSecret was generated rather than
	// supplied. When true, the startup banner prints it so the developer can
	// configure their app to match.
	SecretGenerated bool

	// Scopes is the set of Ecwid access scopes the mock's token is granted. It
	// gates the simulated REST endpoints: a request needing a scope not in this
	// set gets the same 403 a real store returns, so a consumer can exercise the
	// unhappy path. An empty set means "all scopes granted" (the default), so the
	// fixtures answer the happy path with no configuration.
	Scopes []string
}

// HasScope reports whether the mock's token is granted scope. An empty Scopes
// set grants everything — the out-of-the-box default so fixtures just work;
// configure --scopes to a narrower set to test scope-denied paths.
func (c *Config) HasScope(scope string) bool {
	if len(c.Scopes) == 0 {
		return true
	}
	return slices.Contains(c.Scopes, scope)
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

// ProxyEnabled reports whether request forwarding is active. Both the target
// store and its token are required; one without the other cannot proxy.
func (c *Config) ProxyEnabled() bool {
	return c.ProxyStore != "" && c.ProxyToken != ""
}

// Validate checks required fields and value constraints.
func (c *Config) Validate() error {
	var errs []string

	if c.AppURL == "" {
		errs = append(errs, "--app-url is required (env: ECWID_MOCK_APP_URL)")
	} else if u, err := url.Parse(c.AppURL); err != nil || u.Scheme == "" || u.Host == "" {
		errs = append(errs, fmt.Sprintf("--app-url %q is not a valid absolute URL", c.AppURL))
	} else if u.Scheme != "http" && u.Scheme != "https" {
		// The app is loaded in a browser iframe and validated via the
		// postMessage event.origin, both of which only make sense over http(s).
		// Reject any other scheme at startup rather than failing obscurely at
		// render time.
		errs = append(errs, fmt.Sprintf("--app-url %q must use http or https, got scheme %q", c.AppURL, u.Scheme))
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

	// Proxying needs both halves: a target store and its token. One without the
	// other is a misconfiguration, not a partial proxy.
	if (c.ProxyStore == "") != (c.ProxyToken == "") {
		errs = append(errs, "--proxy-store and --proxy-token must be set together to enable proxying")
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
		ClientID:      DefaultClientID,
		StoreID:       DefaultStoreID,
		AuthMode:      DefaultAuthMode,
		Port:          DefaultPort,
		ProxyReadonly: DefaultProxyReadonly,
	}

	cfg.AppURL = resolveString(cmd, "app-url", "ECWID_MOCK_APP_URL", "")
	cfg.ClientID = resolveString(cmd, "client-id", "ECWID_MOCK_CLIENT_ID", DefaultClientID)
	cfg.ClientSecret = resolveString(cmd, "client-secret", "ECWID_MOCK_CLIENT_SECRET", "")
	cfg.StoreID = resolveString(cmd, "store-id", "ECWID_MOCK_STORE_ID", DefaultStoreID)
	cfg.AuthMode = resolveString(cmd, "auth-mode", "ECWID_MOCK_AUTH_MODE", DefaultAuthMode)
	cfg.WebhookURL = resolveString(cmd, "webhook-url", "ECWID_MOCK_WEBHOOK_URL", "")
	cfg.ProxyStore = resolveString(cmd, "proxy-store", "ECWID_MOCK_PROXY_STORE", "")
	cfg.ProxyToken = resolveString(cmd, "proxy-token", "ECWID_MOCK_PROXY_TOKEN", "")
	cfg.AccessToken = resolveString(cmd, "access-token", "ECWID_MOCK_ACCESS_TOKEN", "")
	cfg.Scopes = parseScopes(resolveString(cmd, "scopes", "ECWID_MOCK_SCOPES", ""))

	port, err := resolvePort(cmd)
	if err != nil {
		return Config{}, err
	}
	cfg.Port = port

	readonly, err := resolveBool(cmd, "proxy-readonly", "ECWID_MOCK_PROXY_READONLY", DefaultProxyReadonly)
	if err != nil {
		return Config{}, err
	}
	cfg.ProxyReadonly = readonly

	if cfg.ClientSecret == "" {
		secret, err := generateSecret()
		if err != nil {
			return Config{}, fmt.Errorf("generate client_secret: %w", err)
		}
		cfg.ClientSecret = secret
		cfg.SecretGenerated = true
	}

	if cfg.AccessToken == "" {
		token, err := generateAccessToken()
		if err != nil {
			return Config{}, fmt.Errorf("generate access_token: %w", err)
		}
		cfg.AccessToken = token
	}

	return cfg, nil
}

// parseScopes splits a comma-separated scope list into a trimmed, non-empty
// slice. An empty or all-whitespace input yields nil, which HasScope treats as
// "all scopes granted".
func parseScopes(raw string) []string {
	if strings.TrimSpace(raw) == "" {
		return nil
	}
	var scopes []string
	for s := range strings.SplitSeq(raw, ",") {
		if s = strings.TrimSpace(s); s != "" {
			scopes = append(scopes, s)
		}
	}
	return scopes
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

// resolveBool applies flags > env > default for a boolean value. The flag wins
// only when explicitly set; otherwise the env var is parsed with
// strconv.ParseBool ("1", "t", "true", "0", "f", "false", …); otherwise def.
func resolveBool(cmd *cobra.Command, flag, env string, def bool) (bool, error) {
	if cmd != nil {
		if f := cmd.Flags().Lookup(flag); f != nil && f.Changed {
			b, err := cmd.Flags().GetBool(flag)
			if err != nil {
				return false, fmt.Errorf("parse --%s: %w", flag, err)
			}
			return b, nil
		}
	}
	if v, ok := os.LookupEnv(env); ok {
		b, err := strconv.ParseBool(v)
		if err != nil {
			return false, fmt.Errorf("parse %s %q: %w", env, v, err)
		}
		return b, nil
	}
	return def, nil
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
	return randomHex(generatedSecretBytes)
}

// generateAccessToken returns a hex-encoded random access_token.
func generateAccessToken() (string, error) {
	return randomHex(generatedTokenBytes)
}

// randomHex returns the hex encoding of n cryptographically random bytes.
func randomHex(n int) (string, error) {
	buf := make([]byte, n)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}
