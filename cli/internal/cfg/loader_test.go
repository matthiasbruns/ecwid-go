package cfg

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/config"
)

func TestLoad_FromFile(t *testing.T) {
	content := `store_id: "file-store"
token: "file-token"
output: table
log_level: warn
max_retries: 3
`
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(path, nil)
	if err != nil {
		t.Fatal(err)
	}

	if cfg.StoreID != "file-store" {
		t.Errorf("StoreID = %q, want %q", cfg.StoreID, "file-store")
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
	if cfg.MaxRetries != 3 {
		t.Errorf("MaxRetries = %d, want %d", cfg.MaxRetries, 3)
	}
}

func TestLoad_MissingFile(t *testing.T) {
	cfg, err := Load("/nonexistent/path/config.yaml", nil)
	if err != nil {
		t.Fatalf("unexpected error for missing file: %v", err)
	}
	// Should still get defaults.
	if cfg.Output != config.DefaultOutput {
		t.Errorf("Output = %q, want default %q", cfg.Output, config.DefaultOutput)
	}
}

func TestLoad_EnvOverridesFile(t *testing.T) {
	content := `store_id: "file-store"
token: "file-token"
`
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}

	t.Setenv("ECWID_TOKEN", "env-token")

	cfg, err := Load(path, nil)
	if err != nil {
		t.Fatal(err)
	}

	if cfg.StoreID != "file-store" {
		t.Errorf("StoreID = %q, want file value", cfg.StoreID)
	}
	if cfg.Token != "env-token" {
		t.Errorf("Token = %q, want env value %q", cfg.Token, "env-token")
	}
}

func TestLoad_FlagsOverrideEnv(t *testing.T) {
	t.Setenv("ECWID_OUTPUT", "table")

	cmd := &cobra.Command{}
	cmd.Flags().String("store-id", "", "")
	cmd.Flags().String("token", "", "")
	cmd.Flags().String("output", "", "")
	cmd.Flags().String("log-level", "", "")

	// Simulate flag being set.
	if err := cmd.Flags().Set("output", "json"); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load("", cmd)
	if err != nil {
		t.Fatal(err)
	}

	if cfg.Output != "json" {
		t.Errorf("Output = %q, want flag value %q", cfg.Output, "json")
	}
}

func TestLoad_Defaults(t *testing.T) {
	cfg, err := Load("", nil)
	if err != nil {
		t.Fatal(err)
	}

	if cfg.BaseURL != config.DefaultBaseURL {
		t.Errorf("BaseURL = %q, want %q", cfg.BaseURL, config.DefaultBaseURL)
	}
	if cfg.Output != config.DefaultOutput {
		t.Errorf("Output = %q, want %q", cfg.Output, config.DefaultOutput)
	}
	if cfg.LogLevel != config.DefaultLogLevel {
		t.Errorf("LogLevel = %q, want %q", cfg.LogLevel, config.DefaultLogLevel)
	}
}
