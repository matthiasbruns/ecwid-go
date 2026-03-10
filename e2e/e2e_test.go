// Package e2e runs end-to-end tests against a real Ecwid store.
//
// Requires environment variables:
//
//	ECWID_E2E=1          — gate flag
//	ECWID_STORE_ID       — Ecwid store ID (required for API tests)
//	ECWID_TOKEN          — API access token (required for API tests)
package e2e

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/matthiasbruns/ecwid-go/config"
	"github.com/matthiasbruns/ecwid-go/ecwid"
)

var testClient *ecwid.Client

// binaryPath is the path to the compiled CLI binary, built once in TestMain.
var binaryPath string

// defaultTimeout is the per-request timeout for E2E tests.
const defaultTimeout = 30 * time.Second

// testContext returns a context with the default E2E timeout.
func testContext(t *testing.T) context.Context {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	t.Cleanup(cancel)
	return ctx
}

// requireClient skips the test if the API client is not available (missing credentials).
func requireClient(t *testing.T) {
	t.Helper()
	if testClient == nil {
		t.Skip("skipping: ECWID_STORE_ID and ECWID_TOKEN required for API tests")
	}
}

// skipIfForbidden skips the test if the error is a 403 Forbidden API error.
// This allows tests to pass when the API token lacks required scopes.
func skipIfForbidden(t *testing.T, err error) {
	t.Helper()
	var apiErr *ecwid.APIError
	if errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusForbidden {
		t.Skipf("skipping: token lacks required scope (403 Forbidden)")
	}
}

func TestMain(m *testing.M) {
	if os.Getenv("ECWID_E2E") != "1" {
		_, _ = fmt.Fprintln(os.Stderr, "ECWID_E2E environment variable not set to 1, skipping E2E tests")
		os.Exit(0)
	}

	// Build the CLI binary once for all CLI E2E tests.
	tmpDir, err := os.MkdirTemp("", "ecwid-e2e-*")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create temp dir: %v\n", err)
		os.Exit(1)
	}
	binaryPath = filepath.Join(tmpDir, "ecwid")
	buildCmd := exec.Command("go", "build", "-o", binaryPath, "../cli/")
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to build CLI binary: %v\n", err)
		os.Exit(1)
	}

	// Set up API client if credentials are available.
	cfg := config.Config{
		StoreID:    os.Getenv("ECWID_STORE_ID"),
		Token:      os.Getenv("ECWID_TOKEN"),
		MaxRetries: 3,
	}
	cfg = cfg.WithDefaults()

	if err := cfg.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "API credentials not configured: %v\n", err)
		fmt.Fprintf(os.Stderr, "API client tests will be skipped; CLI-only tests will still run\n")
	} else {
		httpClient := &http.Client{Timeout: 60 * time.Second}
		testClient = ecwid.NewClient(cfg, ecwid.WithHTTPClient(httpClient))
	}

	code := m.Run()
	_ = os.RemoveAll(tmpDir)
	os.Exit(code)
}
