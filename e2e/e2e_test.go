// Package e2e runs end-to-end tests against a real Ecwid store.
//
// Requires environment variables:
//
//	ECWID_E2E=1          — gate flag
//	ECWID_STORE_ID       — Ecwid store ID
//	ECWID_TOKEN          — API access token
package e2e

import (
	"os"
	"testing"

	"github.com/matthiasbruns/ecwid-go/config"
	"github.com/matthiasbruns/ecwid-go/ecwid"
)

var testClient *ecwid.Client

func TestMain(m *testing.M) {
	if os.Getenv("ECWID_E2E") != "1" {
		// Skip all E2E tests when not explicitly enabled.
		os.Exit(0)
	}

	cfg := config.Config{
		StoreID:    os.Getenv("ECWID_STORE_ID"),
		Token:      os.Getenv("ECWID_TOKEN"),
		MaxRetries: 3,
	}
	cfg = cfg.WithDefaults()

	if err := cfg.Validate(); err != nil {
		panic("E2E config invalid: " + err.Error())
	}

	testClient = ecwid.NewClient(cfg)

	os.Exit(m.Run())
}
