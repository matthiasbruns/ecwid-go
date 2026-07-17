package appauth

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"strings"
	"testing"
)

func TestPayload_RedactsTokens(t *testing.T) {
	p := samplerPayload()
	secrets := []string{p.AccessToken, p.PublicToken}

	// fmt "%v" and "%s" and "%+v" all route through String().
	for _, format := range []string{"%v", "%s", "%+v"} {
		out := fmt.Sprintf(format, p)
		for _, secret := range secrets {
			if strings.Contains(out, secret) {
				t.Errorf("Sprintf(%q) leaked secret %q: %s", format, secret, out)
			}
		}
		// Non-secret fields should still be visible for diagnostics.
		if !strings.Contains(out, p.Lang) {
			t.Errorf("Sprintf(%q) dropped non-secret field lang: %s", format, out)
		}
	}

	// A value (not pointer) must redact too — String has a value receiver.
	pv := *p
	if out := pv.String(); strings.Contains(out, p.AccessToken) {
		t.Errorf("value Payload leaked access_token: %s", out)
	}
}

func TestPayload_LogValueRedacts(t *testing.T) {
	p := samplerPayload()

	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, nil))
	logger.LogAttrs(context.Background(), slog.LevelInfo, "iframe", slog.Any("payload", p))

	out := buf.String()
	for _, secret := range []string{p.AccessToken, p.PublicToken} {
		if strings.Contains(out, secret) {
			t.Errorf("slog leaked secret %q: %s", secret, out)
		}
	}
	if !strings.Contains(out, "store_id") {
		t.Errorf("slog dropped store_id group field: %s", out)
	}
}

func TestRedact(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"abcd", "****"},
		{"ab", "****"},
		{"secret_token", "********oken"},
	}
	for _, tt := range tests {
		if got := redact(tt.in); got != tt.want {
			t.Errorf("redact(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
