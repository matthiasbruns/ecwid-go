package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/matthiasbruns/ecwid-go/mock/internal/config"
)

func TestPrintBanner_NoProxy(t *testing.T) {
	var buf bytes.Buffer
	cfg := config.Config{AppURL: "http://localhost:3000", StoreID: "1003", AuthMode: "default", ClientID: "mock-app", Port: 8080}
	if err := printBanner(&buf, &cfg); err != nil {
		t.Fatalf("printBanner: %v", err)
	}
	if strings.Contains(buf.String(), "PROXY ENABLED") {
		t.Errorf("banner shows proxy block when proxying is off:\n%s", buf.String())
	}
}

func TestPrintBanner_ProxyEnabled_NamesStoreAndRedactsToken(t *testing.T) {
	const token = "super-secret-proxy-token"
	var buf bytes.Buffer
	cfg := config.Config{
		AppURL: "http://localhost:3000", StoreID: "1003", AuthMode: "default", ClientID: "mock-app", Port: 8080,
		ProxyStore: "42", ProxyToken: token, ProxyReadonly: true,
	}
	if err := printBanner(&buf, &cfg); err != nil {
		t.Fatalf("printBanner: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "PROXY ENABLED") {
		t.Errorf("banner missing proxy block:\n%s", out)
	}
	if !strings.Contains(out, "42") {
		t.Errorf("banner does not name the target store:\n%s", out)
	}
	if strings.Contains(out, token) {
		t.Errorf("banner leaked the proxy token:\n%s", out)
	}
	if !strings.Contains(out, "read-only:    true") {
		t.Errorf("banner does not state read-only status:\n%s", out)
	}
}

func TestPrintBanner_ProxyWritable_WarnsLoudly(t *testing.T) {
	var buf bytes.Buffer
	cfg := config.Config{
		AppURL: "http://localhost:3000", StoreID: "1003", AuthMode: "default", ClientID: "mock-app", Port: 8080,
		ProxyStore: "42", ProxyToken: "t", ProxyReadonly: false,
	}
	if err := printBanner(&buf, &cfg); err != nil {
		t.Fatalf("printBanner: %v", err)
	}
	if !strings.Contains(buf.String(), "MUTATE") {
		t.Errorf("writable proxy banner does not warn about real mutations:\n%s", buf.String())
	}
}
