package config

import (
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
