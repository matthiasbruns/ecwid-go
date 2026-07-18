package config

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// newCmd builds a command with the same flags serve registers, so Load can be
// exercised in isolation.
func newCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "serve"}
	f := cmd.Flags()
	f.String("app-url", "", "")
	f.String("client-id", DefaultClientID, "")
	f.String("client-secret", "", "")
	f.String("store-id", DefaultStoreID, "")
	f.String("auth-mode", DefaultAuthMode, "")
	f.String("webhook-url", "", "")
	f.Int("port", DefaultPort, "")
	f.String("proxy-store", "", "")
	f.String("proxy-token", "", "")
	f.String("access-token", "", "")
	f.String("scopes", "", "")
	f.Bool("proxy-readonly", DefaultProxyReadonly, "")
	return cmd
}

func TestHasScope(t *testing.T) {
	tests := []struct {
		name   string
		scopes []string
		query  string
		want   bool
	}{
		{name: "empty set grants everything", scopes: nil, query: "read_customers", want: true},
		{name: "granted scope is allowed", scopes: []string{"read_customers", "update_customers"}, query: "read_customers", want: true},
		{name: "withheld scope is denied", scopes: []string{"read_store_profile"}, query: "read_customers", want: false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c := Config{Scopes: tc.scopes}
			if got := c.HasScope(tc.query); got != tc.want {
				t.Errorf("HasScope(%q) with %v = %v, want %v", tc.query, tc.scopes, got, tc.want)
			}
		})
	}
}

func TestParseScopes(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want []string
	}{
		{name: "empty is nil", raw: "", want: nil},
		{name: "whitespace is nil", raw: "  ", want: nil},
		{name: "trims and splits", raw: " read_customers , update_customers ", want: []string{"read_customers", "update_customers"}},
		{name: "drops empty entries", raw: "read_customers,,", want: []string{"read_customers"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := parseScopes(tc.raw)
			if strings.Join(got, ",") != strings.Join(tc.want, ",") {
				t.Errorf("parseScopes(%q) = %v, want %v", tc.raw, got, tc.want)
			}
		})
	}
}

func TestLoad_Scopes(t *testing.T) {
	cmd := newCmd()
	if err := cmd.Flags().Set("scopes", "read_store_profile,read_customers"); err != nil {
		t.Fatalf("set scopes flag: %v", err)
	}
	cfg, err := Load(cmd)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if !cfg.HasScope("read_customers") {
		t.Error("expected read_customers to be granted")
	}
	if cfg.HasScope("update_customers") {
		t.Error("expected update_customers to be withheld")
	}
}

func TestLoad_Defaults(t *testing.T) {
	cmd := newCmd()
	if err := cmd.Flags().Set("app-url", "http://localhost:3000"); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(cmd)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}

	if cfg.ClientID != DefaultClientID {
		t.Errorf("ClientID = %q, want %q", cfg.ClientID, DefaultClientID)
	}
	if cfg.StoreID != DefaultStoreID {
		t.Errorf("StoreID = %q, want %q", cfg.StoreID, DefaultStoreID)
	}
	if cfg.AuthMode != DefaultAuthMode {
		t.Errorf("AuthMode = %q, want %q", cfg.AuthMode, DefaultAuthMode)
	}
	if cfg.Port != DefaultPort {
		t.Errorf("Port = %d, want %d", cfg.Port, DefaultPort)
	}
	if !cfg.SecretGenerated {
		t.Error("SecretGenerated = false, want true when no secret supplied")
	}
	if len(cfg.ClientSecret) < MinClientSecretLen {
		t.Errorf("generated ClientSecret len = %d, want >= %d", len(cfg.ClientSecret), MinClientSecretLen)
	}
	if cfg.AccessToken == "" {
		t.Error("AccessToken is empty, want a generated token when none supplied")
	}
}

func TestLoad_SuppliedAccessToken(t *testing.T) {
	cmd := newCmd()
	if err := cmd.Flags().Set("app-url", "http://localhost:3000"); err != nil {
		t.Fatal(err)
	}
	if err := cmd.Flags().Set("access-token", "supplied-token"); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(cmd)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.AccessToken != "supplied-token" {
		t.Errorf("AccessToken = %q, want supplied value", cfg.AccessToken)
	}
}

func TestLoad_FlagsOverrideEnv(t *testing.T) {
	t.Setenv("ECWID_MOCK_STORE_ID", "env-store")
	t.Setenv("ECWID_MOCK_PORT", "9999")

	cmd := newCmd()
	for k, v := range map[string]string{
		"app-url":  "http://localhost:3000",
		"store-id": "flag-store",
		"port":     "1234",
	} {
		if err := cmd.Flags().Set(k, v); err != nil {
			t.Fatal(err)
		}
	}

	cfg, err := Load(cmd)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}

	if cfg.StoreID != "flag-store" {
		t.Errorf("StoreID = %q, want flag to win over env", cfg.StoreID)
	}
	if cfg.Port != 1234 {
		t.Errorf("Port = %d, want flag 1234 to win over env", cfg.Port)
	}
}

func TestLoad_EnvOverridesDefault(t *testing.T) {
	t.Setenv("ECWID_MOCK_APP_URL", "http://env-app:3000")
	t.Setenv("ECWID_MOCK_STORE_ID", "env-store")
	t.Setenv("ECWID_MOCK_PORT", "9999")

	cfg, err := Load(newCmd())
	if err != nil {
		t.Fatalf("Load: %v", err)
	}

	if cfg.AppURL != "http://env-app:3000" {
		t.Errorf("AppURL = %q, want env value", cfg.AppURL)
	}
	if cfg.StoreID != "env-store" {
		t.Errorf("StoreID = %q, want env value", cfg.StoreID)
	}
	if cfg.Port != 9999 {
		t.Errorf("Port = %d, want env 9999", cfg.Port)
	}
}

func TestLoad_SuppliedSecretNotGenerated(t *testing.T) {
	cmd := newCmd()
	if err := cmd.Flags().Set("app-url", "http://localhost:3000"); err != nil {
		t.Fatal(err)
	}
	if err := cmd.Flags().Set("client-secret", "0123456789abcdef0123"); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(cmd)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}

	if cfg.SecretGenerated {
		t.Error("SecretGenerated = true, want false when secret supplied")
	}
	if cfg.ClientSecret != "0123456789abcdef0123" {
		t.Errorf("ClientSecret = %q, want supplied value", cfg.ClientSecret)
	}
}

func TestLoad_InvalidEnvPort(t *testing.T) {
	t.Setenv("ECWID_MOCK_PORT", "not-a-number")

	if _, err := Load(newCmd()); err == nil {
		t.Fatal("Load: want error for non-numeric ECWID_MOCK_PORT, got nil")
	}
}

func TestValidate(t *testing.T) {
	validSecret := "0123456789abcdef" // exactly 16 bytes

	tests := []struct {
		name    string
		cfg     Config
		wantErr string
	}{
		{
			name: "valid",
			cfg:  Config{AppURL: "http://localhost:3000", ClientSecret: validSecret, AuthMode: AuthModeDefault, Port: 8080},
		},
		{
			name:    "missing app-url",
			cfg:     Config{ClientSecret: validSecret, AuthMode: AuthModeDefault, Port: 8080},
			wantErr: "--app-url is required",
		},
		{
			name:    "invalid app-url",
			cfg:     Config{AppURL: "not-a-url", ClientSecret: validSecret, AuthMode: AuthModeDefault, Port: 8080},
			wantErr: "not a valid absolute URL",
		},
		{
			name:    "non-http(s) app-url scheme",
			cfg:     Config{AppURL: "ftp://localhost:3000", ClientSecret: validSecret, AuthMode: AuthModeDefault, Port: 8080},
			wantErr: "must use http or https",
		},
		{
			name:    "short secret",
			cfg:     Config{AppURL: "http://localhost:3000", ClientSecret: "tooshort", AuthMode: AuthModeDefault, Port: 8080},
			wantErr: "at least 16 bytes",
		},
		{
			name:    "invalid auth mode",
			cfg:     Config{AppURL: "http://localhost:3000", ClientSecret: validSecret, AuthMode: "bogus", Port: 8080},
			wantErr: "--auth-mode",
		},
		{
			name:    "port out of range",
			cfg:     Config{AppURL: "http://localhost:3000", ClientSecret: validSecret, AuthMode: AuthModeDefault, Port: 70000},
			wantErr: "out of range",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Validate()
			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("Validate() = %v, want nil", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("Validate() = nil, want error containing %q", tt.wantErr)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("Validate() = %v, want error containing %q", err, tt.wantErr)
			}
		})
	}
}

func TestLoad_ProxyReadonlyDefaultsTrue(t *testing.T) {
	cmd := newCmd()
	if err := cmd.Flags().Set("app-url", "http://localhost:3000"); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(cmd)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if !cfg.ProxyReadonly {
		t.Error("ProxyReadonly = false, want true by default (writes must be blocked unless opted in)")
	}
}

func TestLoad_ProxyReadonlyResolution(t *testing.T) {
	t.Run("flag false wins over env true", func(t *testing.T) {
		t.Setenv("ECWID_MOCK_PROXY_READONLY", "true")
		cmd := newCmd()
		if err := cmd.Flags().Set("app-url", "http://localhost:3000"); err != nil {
			t.Fatal(err)
		}
		if err := cmd.Flags().Set("proxy-readonly", "false"); err != nil {
			t.Fatal(err)
		}
		cfg, err := Load(cmd)
		if err != nil {
			t.Fatalf("Load: %v", err)
		}
		if cfg.ProxyReadonly {
			t.Error("ProxyReadonly = true, want flag false to win over env")
		}
	})

	t.Run("env false overrides default", func(t *testing.T) {
		t.Setenv("ECWID_MOCK_PROXY_READONLY", "false")
		cfg, err := Load(newCmd())
		if err != nil {
			t.Fatalf("Load: %v", err)
		}
		if cfg.ProxyReadonly {
			t.Error("ProxyReadonly = true, want env false to override default")
		}
	})

	t.Run("invalid env is an error", func(t *testing.T) {
		t.Setenv("ECWID_MOCK_PROXY_READONLY", "maybe")
		if _, err := Load(newCmd()); err == nil {
			t.Fatal("Load: want error for non-boolean ECWID_MOCK_PROXY_READONLY, got nil")
		}
	})
}

func TestValidate_ProxyStoreAndTokenPairing(t *testing.T) {
	validSecret := "0123456789abcdef"
	base := func() Config {
		return Config{AppURL: "http://localhost:3000", ClientSecret: validSecret, AuthMode: AuthModeDefault, Port: 8080}
	}

	tests := []struct {
		name    string
		mutate  func(*Config)
		wantErr string
	}{
		{name: "neither set", mutate: func(*Config) {}},
		{name: "both set", mutate: func(c *Config) { c.ProxyStore = "42"; c.ProxyToken = "t" }},
		{name: "store without token", mutate: func(c *Config) { c.ProxyStore = "42" }, wantErr: "must be set together"},
		{name: "token without store", mutate: func(c *Config) { c.ProxyToken = "t" }, wantErr: "must be set together"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := base()
			tt.mutate(&cfg)
			err := cfg.Validate()
			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("Validate() = %v, want nil", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("Validate() = %v, want error containing %q", err, tt.wantErr)
			}
		})
	}
}

func TestProxyEnabled(t *testing.T) {
	if (&Config{}).ProxyEnabled() {
		t.Error("ProxyEnabled() = true for empty config, want false")
	}
	if (&Config{ProxyStore: "42"}).ProxyEnabled() {
		t.Error("ProxyEnabled() = true with store but no token, want false")
	}
	if !(&Config{ProxyStore: "42", ProxyToken: "t"}).ProxyEnabled() {
		t.Error("ProxyEnabled() = false with both set, want true")
	}
}

func TestRedactedSecrets(t *testing.T) {
	cfg := Config{ClientSecret: "super_secret_value", ProxyToken: "proxy_token_value"}

	if got := cfg.RedactedClientSecret(); strings.Contains(got, "super_secret") {
		t.Errorf("RedactedClientSecret leaked secret: %q", got)
	}
	if got := cfg.RedactedProxyToken(); strings.Contains(got, "proxy_token") {
		t.Errorf("RedactedProxyToken leaked token: %q", got)
	}

	empty := Config{}
	if got := empty.RedactedProxyToken(); got != "" {
		t.Errorf("RedactedProxyToken() = %q, want empty when no token set", got)
	}
}
