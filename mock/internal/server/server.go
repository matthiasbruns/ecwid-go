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
	cfg  config.Config
	log  *slog.Logger
	mux  *http.ServeMux
	http *http.Server
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
		cfg: cfg,
		log: log,
		mux: mux,
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

// routes registers the mock's HTTP handlers. Only the control-plane health
// check exists in this skeleton; the admin shell, simulated REST, and remaining
// control endpoints are added by later issues.
func (s *Server) routes() {
	s.mux.HandleFunc("GET /_mock/health", handleHealth)
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
