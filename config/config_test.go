package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     Config
		wantErr bool
	}{
		{
			name:    "valid config",
			cfg:     Config{StoreID: "123", Token: "abc"},
			wantErr: false,
		},
		{
			name:    "missing store_id",
			cfg:     Config{Token: "abc"},
			wantErr: true,
		},
		{
			name:    "missing token",
			cfg:     Config{StoreID: "123"},
			wantErr: true,
		},
		{
			name:    "invalid output",
			cfg:     Config{StoreID: "123", Token: "abc", Output: "xml"},
			wantErr: true,
		},
		{
			name:    "invalid log level",
			cfg:     Config{StoreID: "123", Token: "abc", LogLevel: "trace"},
			wantErr: true,
		},
		{
			name:    "negative max_retries",
			cfg:     Config{StoreID: "123", Token: "abc", MaxRetries: -1},
			wantErr: true,
		},
		{
			name:    "valid with all fields",
			cfg:     Config{StoreID: "123", Token: "abc", Output: "table", LogLevel: "debug", MaxRetries: 3},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig_WithDefaults(t *testing.T) {
	cfg := Config{StoreID: "123", Token: "abc"}.WithDefaults()

	if cfg.BaseURL != DefaultBaseURL {
		t.Errorf("BaseURL = %q, want %q", cfg.BaseURL, DefaultBaseURL)
	}
	if cfg.Output != DefaultOutput {
		t.Errorf("Output = %q, want %q", cfg.Output, DefaultOutput)
	}
	if cfg.LogLevel != DefaultLogLevel {
		t.Errorf("LogLevel = %q, want %q", cfg.LogLevel, DefaultLogLevel)
	}
}

func TestConfig_WithDefaults_NoOverwrite(t *testing.T) {
	cfg := Config{
		StoreID:  "123",
		Token:    "abc",
		BaseURL:  "https://custom.example.com",
		Output:   "table",
		LogLevel: "debug",
	}.WithDefaults()

	if cfg.BaseURL != "https://custom.example.com" {
		t.Errorf("BaseURL overwritten: %q", cfg.BaseURL)
	}
	if cfg.Output != "table" {
		t.Errorf("Output overwritten: %q", cfg.Output)
	}
	if cfg.LogLevel != "debug" {
		t.Errorf("LogLevel overwritten: %q", cfg.LogLevel)
	}
}

func TestLoadFromEnv(t *testing.T) {
	t.Setenv(EnvStoreID, "99999")
	t.Setenv(EnvToken, "env-token")
	t.Setenv(EnvBaseURL, "https://test.example.com")
	t.Setenv(EnvOutput, "table")
	t.Setenv(EnvLogLevel, "debug")
	t.Setenv("ECWID_MAX_RETRIES", "5")

	cfg := LoadFromEnv()

	if cfg.StoreID != "99999" {
		t.Errorf("StoreID = %q, want %q", cfg.StoreID, "99999")
	}
	if cfg.Token != "env-token" {
		t.Errorf("Token = %q, want %q", cfg.Token, "env-token")
	}
	if cfg.BaseURL != "https://test.example.com" {
		t.Errorf("BaseURL = %q, want %q", cfg.BaseURL, "https://test.example.com")
	}
	if cfg.Output != "table" {
		t.Errorf("Output = %q, want %q", cfg.Output, "table")
	}
	if cfg.LogLevel != "debug" {
		t.Errorf("LogLevel = %q, want %q", cfg.LogLevel, "debug")
	}
	if cfg.MaxRetries != 5 {
		t.Errorf("MaxRetries = %d, want %d", cfg.MaxRetries, 5)
	}
}

func TestLoadFromFile(t *testing.T) {
	content := `# Ecwid config
store_id: "12345"
token: "file-token"
output: table
log_level: warn
max_retries: 2
`
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}

	cfg, err := LoadFromFile(path)
	if err != nil {
		t.Fatal(err)
	}

	if cfg.StoreID != "12345" {
		t.Errorf("StoreID = %q, want %q", cfg.StoreID, "12345")
	}
	if cfg.Token != "file-token" {
		t.Errorf("Token = %q, want %q", cfg.Token, "file-token")
	}
	if cfg.Output != "table" {
		t.Errorf("Output = %q, want %q", cfg.Output, "table")
	}
	if cfg.LogLevel != "warn" {
		t.Errorf("LogLevel = %q, want %q", cfg.LogLevel, "warn")
	}
	if cfg.MaxRetries != 2 {
		t.Errorf("MaxRetries = %d, want %d", cfg.MaxRetries, 2)
	}
}

func TestLoadFromFile_NotFound(t *testing.T) {
	cfg, err := LoadFromFile("/nonexistent/path/config.yaml")
	if err != nil {
		t.Errorf("expected nil error for missing file, got %v", err)
	}
	if cfg.StoreID != "" {
		t.Errorf("expected empty config, got StoreID=%q", cfg.StoreID)
	}
}

func TestMerge(t *testing.T) {
	base := Config{
		StoreID:  "base-id",
		Token:    "base-token",
		BaseURL:  "https://base.example.com",
		Output:   "json",
		LogLevel: "info",
	}
	overlay := Config{
		StoreID: "overlay-id",
		Token:   "overlay-token",
	}

	merged := Merge(base, overlay)

	if merged.StoreID != "overlay-id" {
		t.Errorf("StoreID = %q, want %q", merged.StoreID, "overlay-id")
	}
	if merged.Token != "overlay-token" {
		t.Errorf("Token = %q, want %q", merged.Token, "overlay-token")
	}
	if merged.BaseURL != "https://base.example.com" {
		t.Errorf("BaseURL = %q, want base value", merged.BaseURL)
	}
	if merged.Output != "json" {
		t.Errorf("Output = %q, want base value", merged.Output)
	}
}

func TestLoad_Precedence(t *testing.T) {
	// File config
	content := `store_id: "file-id"
token: "file-token"
output: table
`
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}

	// Env overrides file
	t.Setenv(EnvToken, "env-token")

	// Explicit overrides override env
	overrides := Config{Output: "json"}

	cfg, err := Load(path, overrides)
	if err != nil {
		t.Fatal(err)
	}

	if cfg.StoreID != "file-id" {
		t.Errorf("StoreID = %q, want file value", cfg.StoreID)
	}
	if cfg.Token != "env-token" {
		t.Errorf("Token = %q, want env value", cfg.Token)
	}
	if cfg.Output != "json" {
		t.Errorf("Output = %q, want override value", cfg.Output)
	}
}

func TestConfig_RedactedToken(t *testing.T) {
	tests := []struct {
		token string
		want  string
	}{
		{"secret_abc123", "*********c123"},
		{"abcde", "*bcde"},
		{"abcd", "****"},
		{"abc", "****"},
		{"", "****"},
	}

	for _, tt := range tests {
		cfg := Config{Token: tt.token}
		got := cfg.RedactedToken()
		if got != tt.want {
			t.Errorf("RedactedToken(%q) = %q, want %q", tt.token, got, tt.want)
		}
	}
}

func TestParseSimpleYAML_Comments(t *testing.T) {
	content := `# This is a comment
store_id: "123"
# Another comment
token: abc
`
	cfg, err := parseSimpleYAML([]byte(content))
	if err != nil {
		t.Fatal(err)
	}
	if cfg.StoreID != "123" {
		t.Errorf("StoreID = %q, want %q", cfg.StoreID, "123")
	}
	if cfg.Token != "abc" {
		t.Errorf("Token = %q, want %q", cfg.Token, "abc")
	}
}

func TestConfig_MarshalJSON_RedactsToken(t *testing.T) {
	cfg := Config{
		StoreID: "12345",
		Token:   "super_secret_token",
	}

	data, err := cfg.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	jsonStr := string(data)

	// Must NOT contain the raw token.
	if strings.Contains(jsonStr, "super_secret_token") {
		t.Errorf("MarshalJSON leaked raw token: %s", jsonStr)
	}

	// Must contain the redacted version.
	if !strings.Contains(jsonStr, cfg.RedactedToken()) {
		t.Errorf("MarshalJSON missing redacted token %q in: %s", cfg.RedactedToken(), jsonStr)
	}
}

func TestParseSimpleYAML_InvalidLine(t *testing.T) {
	content := `store_id: "123"
this is not valid yaml
`
	_, err := parseSimpleYAML([]byte(content))
	if err == nil {
		t.Error("expected error for invalid line")
	}
}
