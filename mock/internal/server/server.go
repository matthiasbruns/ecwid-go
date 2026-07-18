// Package server assembles the mock HTTP server: its route namespaces, request
// logging, and graceful lifecycle.
//
// # Route namespaces
//
// The mux is partitioned into three namespaces that later issues build on:
//
//	/                        -> admin shell (the developer-facing UI)
//	/api/v3/{storeId}/...    -> simulated Ecwid REST (proxy/501 fallback)
//	/_mock/...               -> the mock's own control API
//
// The /_mock/ prefix is reserved for the mock's control plane so it can never
// collide with a real Ecwid REST route. This skeleton registers only
// GET /_mock/health; the other namespaces are established here and filled in by
// later issues.
package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/matthiasbruns/ecwid-go/mock/internal/config"
	"github.com/matthiasbruns/ecwid-go/mock/internal/webhook"
)

// readHeaderTimeout bounds how long the server waits for request headers,
// mitigating slow-header (Slowloris) clients. Required by gosec/golangci.
const readHeaderTimeout = 10 * time.Second

// readTimeout bounds how long the server waits for a full request (headers plus
// body), so a slow client cannot hold a connection open indefinitely.
const readTimeout = 30 * time.Second

// idleTimeout bounds how long an idle keep-alive connection is retained.
const idleTimeout = 120 * time.Second

// shutdownTimeout bounds graceful shutdown before in-flight requests are
// abandoned.
const shutdownTimeout = 10 * time.Second

// Server wraps the HTTP server and its configuration.
type Server struct {
	cfg     config.Config
	log     *slog.Logger
	mux     *http.ServeMux
	http    *http.Server
	store   *appStorage
	trigger *webhook.Trigger

	// upstreamBase is the scheme+host proxied requests are forwarded to. It
	// defaults to defaultUpstreamBase and is overridden by tests to point at an
	// httptest server instead of the real Ecwid API.
	upstreamBase string
	// proxyClient forwards proxied requests. It is a plain http.Client (not the
	// ecwid internal/api transport, which is JSON-oriented and cannot pass an
	// opaque response through) with a bounded timeout.
	proxyClient *http.Client
}

// New builds a Server with all routes registered. The provided logger is used
// for request and lifecycle logging.
func New(cfg config.Config, log *slog.Logger) *Server {
	if log == nil {
		log = slog.Default()
	}

	mux := http.NewServeMux()
	addr := net.JoinHostPort("", strconv.Itoa(cfg.Port))

	s := &Server{
		cfg:   cfg,
		log:   log,
		mux:   mux,
		store: newAppStorage(),
		trigger: webhook.NewTrigger(webhook.Config{
			ClientSecret: cfg.ClientSecret,
			StoreID:      parseStoreID(cfg.StoreID, log),
			URL:          cfg.WebhookURL,
		}),
		upstreamBase: defaultUpstreamBase,
		proxyClient:  &http.Client{Timeout: proxyTimeout},
	}

	s.routes()

	s.http = &http.Server{
		Addr:              addr,
		Handler:           logRequests(log, mux),
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
		IdleTimeout:       idleTimeout,
		// WriteTimeout is intentionally left unset: later issues add streaming /
		// long-poll control endpoints (e.g. the admin shell's live webhook
		// delivery log) that a fixed write deadline would sever mid-response.
	}

	return s
}

// routes registers the mock's HTTP handlers. The admin shell and simulated REST
// namespaces are filled in by later issues; the control plane carries the health
// check and the webhook trigger API and UI.
func (s *Server) routes() {
	s.mux.HandleFunc("GET /_mock/health", handleHealth)

	// The admin shell lives at the site root only; unknown paths fall through to
	// ServeMux's 404 rather than being swallowed by a catch-all.
	s.mux.HandleFunc("GET /{$}", s.handleShell)

	// Simulated Ecwid REST: the app-storage endpoints the JS SDK actually calls
	// over HTTP (getAppStorage/setAppStorage and the public-config variants).
	s.mux.HandleFunc("GET /api/v3/{storeId}/storage", s.handleStorageList)
	s.mux.HandleFunc("GET /api/v3/{storeId}/storage/{key}", s.handleStorageGet)
	s.mux.HandleFunc("PUT /api/v3/{storeId}/storage/{key}", s.handleStoragePut)
	s.mux.HandleFunc("POST /api/v3/{storeId}/storage/{key}", s.handleStoragePut)
	s.mux.HandleFunc("DELETE /api/v3/{storeId}/storage/{key}", s.handleStorageDelete)

	// Simulated Ecwid REST namespace. This catch-all backstops every REST route
	// the mock does not implement locally: it proxies to a real store when
	// configured, otherwise returns an informative 501. Locally-implemented
	// routes (e.g. /storage above) register more specific patterns that
	// ServeMux prefers over this one; the handler also guards /storage directly
	// so it is never proxied.
	s.mux.HandleFunc("/api/v3/{storeId}/{rest...}", s.handleRESTFallback)

	// Control plane: the webhook trigger API and its UI panel.
	webhook.NewHandler(s.trigger).Routes(s.mux)
}

// Run starts the server and blocks until ctx is canceled (e.g. on SIGINT or
// SIGTERM), then shuts down gracefully. It returns nil on a clean shutdown.
func (s *Server) Run(ctx context.Context) error {
	errCh := make(chan error, 1)

	go func() {
		s.log.Info("mock server listening",
			"addr", s.http.Addr,
			"store_id", s.cfg.StoreID,
			"auth_mode", s.cfg.AuthMode,
			"client_id", s.cfg.ClientID,
		)
		if err := s.http.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
			return
		}
		errCh <- nil
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		s.log.Info("shutdown signal received, draining connections")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()
		if err := s.http.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("graceful shutdown: %w", err)
		}
		s.log.Info("mock server stopped")
		return nil
	}
}

// Handler exposes the underlying mux for tests.
func (s *Server) Handler() http.Handler {
	return s.mux
}

// parseStoreID converts the configured store ID to the numeric form Ecwid puts
// in a webhook's storeId. A non-numeric value is not fatal — the mock still runs
// and fires webhooks — so it logs and falls back to 0 rather than failing.
func parseStoreID(id string, log *slog.Logger) int64 {
	n, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Warn("store ID is not numeric; webhooks will use storeId 0", "store_id", id)
		return 0
	}
	return n
}
