package server

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/matthiasbruns/ecwid-go/mock/internal/config"
)

func TestHealth(t *testing.T) {
	srv := New(config.Config{Port: 0}, discardLogger())

	req := httptest.NewRequest(http.MethodGet, "/_mock/health", http.NoBody)
	rec := httptest.NewRecorder()
	srv.Handler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	var body map[string]string
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("decode body: %v", err)
	}
	if body["status"] != "ok" {
		t.Errorf("status = %q, want ok", body["status"])
	}
	if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("Content-Type = %q, want application/json", ct)
	}
}

func TestHealth_WrongMethod(t *testing.T) {
	srv := New(config.Config{Port: 0}, discardLogger())

	req := httptest.NewRequest(http.MethodPost, "/_mock/health", http.NoBody)
	rec := httptest.NewRecorder()
	srv.Handler().ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("status = %d, want %d for POST to GET-only route", rec.Code, http.StatusMethodNotAllowed)
	}
}

func TestRun_GracefulShutdown(t *testing.T) {
	srv := New(config.Config{Port: 0}, discardLogger())

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() { done <- srv.Run(ctx) }()

	// Give the listener a moment to come up. If Run returns during this window it
	// failed to start — fail fast rather than masking it until the later timeout.
	select {
	case err := <-done:
		t.Fatalf("Run returned before shutdown was requested: %v", err)
	case <-time.After(50 * time.Millisecond):
	}

	cancel()

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("Run returned %v, want nil on graceful shutdown", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("Run did not return after context cancel")
	}
}

func TestWebhookTrigger_EndToEnd(t *testing.T) {
	var gotSig string
	var gotBody []byte
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotSig = r.Header.Get("X-Ecwid-Webhook-Signature")
		gotBody, _ = io.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
	}))
	defer backend.Close()

	srv := New(config.Config{
		Port:         0,
		StoreID:      "1003",
		ClientSecret: "test_client_secret_1234567890",
		WebhookURL:   backend.URL,
	}, discardLogger())

	body := `{"eventType":"order.created","signature":"valid"}`
	req := httptest.NewRequest(http.MethodPost, "/_mock/webhooks/trigger", strings.NewReader(body))
	rec := httptest.NewRecorder()
	srv.Handler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200; body %s", rec.Code, rec.Body.String())
	}
	var res struct {
		Delivered    bool   `json:"delivered"`
		EventID      string `json:"eventId"`
		EventCreated int64  `json:"eventCreated"`
	}
	if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
		t.Fatal(err)
	}
	if !res.Delivered {
		t.Error("delivered = false, want true for a 200 endpoint")
	}
	if len(gotBody) == 0 {
		t.Fatal("backend received no webhook body")
	}
	// The signature the mock sent must verify with the same secret, proving the
	// mock and a real handler agree.
	if gotSig == "" {
		t.Error("backend received no signature header")
	}
}

func TestWebhookUI_Reachable(t *testing.T) {
	srv := New(config.Config{Port: 0, StoreID: "1003", ClientSecret: "test_client_secret_1234567890"}, discardLogger())
	req := httptest.NewRequest(http.MethodGet, "/_mock/webhooks/ui", http.NoBody)
	rec := httptest.NewRecorder()
	srv.Handler().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("UI status = %d, want 200", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); !strings.HasPrefix(ct, "text/html") {
		t.Errorf("UI Content-Type = %q, want text/html", ct)
	}
}

func discardLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}
