// Package e2e runs end-to-end tests against a real Ecwid store.
//
// Requires environment variables:
//
//	ECWID_E2E=1          — gate flag
//	ECWID_STORE_ID       — Ecwid store ID
//	ECWID_TOKEN          — API access token
package e2e

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/matthiasbruns/ecwid-go/config"
	"github.com/matthiasbruns/ecwid-go/ecwid"
)

var testClient *ecwid.Client

// defaultTimeout is the per-request timeout for E2E tests.
const defaultTimeout = 30 * time.Second

// testContext returns a context with the default E2E timeout.
func testContext(t *testing.T) context.Context {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	t.Cleanup(cancel)
	return ctx
}

func TestMain(m *testing.M) {
	if os.Getenv("ECWID_E2E") != "1" {
		os.Exit(0)
	}

	cfg := config.Config{
		StoreID:    os.Getenv("ECWID_STORE_ID"),
		Token:      os.Getenv("ECWID_TOKEN"),
		MaxRetries: 3,
	}
	cfg = cfg.WithDefaults()

	if err := cfg.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "E2E config invalid: %v\n", err)
		fmt.Fprintf(os.Stderr, "Required: ECWID_STORE_ID and ECWID_TOKEN environment variables\n")
		os.Exit(1)
	}

	httpClient := &http.Client{Timeout: 60 * time.Second}
	testClient = ecwid.NewClient(cfg, ecwid.WithHTTPClient(httpClient))

	os.Exit(m.Run())
}
