package server

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/matthiasbruns/ecwid-go/mock/internal/config"
)

// proxyServer builds a Server whose proxy target is upstream (an httptest URL,
// never the real API) and whose logs are captured in logBuf.
func proxyServer(cfg config.Config, upstream string, logBuf io.Writer) *Server {
	log := slog.New(slog.NewTextHandler(logBuf, &slog.HandlerOptions{Level: slog.LevelDebug}))
	srv := New(cfg, log)
	srv.upstreamBase = upstream
	return srv
}

func serve(srv *Server, req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	srv.Handler().ServeHTTP(rec, req)
	return rec
}

func decodeErrorMessage(t *testing.T, body *bytes.Buffer) string {
	t.Helper()
	var parsed struct {
		ErrorMessage string `json:"errorMessage"`
	}
	if err := json.NewDecoder(body).Decode(&parsed); err != nil {
		t.Fatalf("decode error body: %v", err)
	}
	return parsed.ErrorMessage
}

func TestFallback_NotImplemented_NamesEndpointAndRemedy(t *testing.T) {
	srv := proxyServer(config.Config{StoreID: "1003"}, defaultUpstreamBase, io.Discard)

	req := httptest.NewRequest(http.MethodGet, "/api/v3/1003/products", http.NoBody)
	rec := serve(srv, req)

	if rec.Code != http.StatusNotImplemented {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusNotImplemented)
	}
	msg := decodeErrorMessage(t, rec.Body)
	if !strings.Contains(msg, "GET /api/v3/{storeId}/products") {
		t.Errorf("message does not name the endpoint: %q", msg)
	}
	if !strings.Contains(msg, "--proxy-store") || !strings.Contains(msg, "--proxy-token") {
		t.Errorf("message does not state the proxy remedy: %q", msg)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("Content-Type = %q, want application/json", ct)
	}
}

func TestFallback_Proxy_ForwardsRewritingStoreAndToken(t *testing.T) {
	var gotPath, gotQuery, gotAuth, gotMethod string
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod, gotPath, gotQuery = r.Method, r.URL.Path, r.URL.RawQuery
		gotAuth = r.Header.Get("Authorization")
		w.Header().Set("X-Upstream", "yes")
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, `{"items":[]}`)
	}))
	defer upstream.Close()

	cfg := config.Config{StoreID: "1003", ProxyStore: "42", ProxyToken: "real-token", ProxyReadonly: true}
	srv := proxyServer(cfg, upstream.URL, io.Discard)

	req := httptest.NewRequest(http.MethodGet, "/api/v3/1003/products?limit=10", http.NoBody)
	req.Header.Set("Authorization", "Bearer mock-access-token")
	rec := serve(srv, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("upstream method = %q, want GET", gotMethod)
	}
	if gotPath != "/api/v3/42/products" {
		t.Errorf("upstream path = %q, want store ID rewritten to /api/v3/42/products", gotPath)
	}
	if gotQuery != "limit=10" {
		t.Errorf("upstream query = %q, want limit=10", gotQuery)
	}
	if gotAuth != "Bearer real-token" {
		t.Errorf("upstream Authorization = %q, want the proxy token swapped in", gotAuth)
	}
	if body := rec.Body.String(); !strings.Contains(body, `"items"`) {
		t.Errorf("response body not passed through: %q", body)
	}
	if rec.Header().Get("X-Upstream") != "yes" {
		t.Error("upstream response header not passed through")
	}
}

func TestFallback_Proxy_PassesLargeBodyThroughUntruncated(t *testing.T) {
	// A body larger than a naive buffer/cap would truncate; assert it arrives
	// whole so response passthrough never silently corrupts big payloads.
	const size = 2 << 20 // 2 MiB
	payload := bytes.Repeat([]byte("x"), size)
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(payload)
	}))
	defer upstream.Close()

	cfg := config.Config{StoreID: "1003", ProxyStore: "42", ProxyToken: "real-token", ProxyReadonly: true}
	srv := proxyServer(cfg, upstream.URL, io.Discard)

	rec := serve(srv, httptest.NewRequest(http.MethodGet, "/api/v3/1003/products", http.NoBody))
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if got := rec.Body.Len(); got != size {
		t.Errorf("response body length = %d, want %d (must not be truncated)", got, size)
	}
}

func TestFallback_ReadonlyDefault_BlocksMutations(t *testing.T) {
	var upstreamHits int
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		upstreamHits++
		w.WriteHeader(http.StatusOK)
	}))
	defer upstream.Close()

	cfg := config.Config{StoreID: "1003", ProxyStore: "42", ProxyToken: "real-token", ProxyReadonly: true}
	srv := proxyServer(cfg, upstream.URL, io.Discard)

	// GET proxies through.
	getRec := serve(srv, httptest.NewRequest(http.MethodGet, "/api/v3/1003/products", http.NoBody))
	if getRec.Code != http.StatusOK {
		t.Errorf("GET status = %d, want 200 (proxied)", getRec.Code)
	}

	// POST is blocked with 403 and never reaches the upstream.
	postRec := serve(srv, httptest.NewRequest(http.MethodPost, "/api/v3/1003/products", strings.NewReader(`{}`)))
	if postRec.Code != http.StatusForbidden {
		t.Fatalf("POST status = %d, want 403 under --proxy-readonly", postRec.Code)
	}
	msg := decodeErrorMessage(t, postRec.Body)
	if !strings.Contains(msg, "--proxy-readonly=false") {
		t.Errorf("403 message does not name the override flag: %q", msg)
	}
	if upstreamHits != 1 {
		t.Errorf("upstream hits = %d, want 1 (only the GET); the blocked POST must not reach upstream", upstreamHits)
	}
}

func TestFallback_ReadonlyDisabled_ProxiesMutations(t *testing.T) {
	var gotMethod, gotBody string
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		b, _ := io.ReadAll(r.Body)
		gotBody = string(b)
		w.WriteHeader(http.StatusCreated)
	}))
	defer upstream.Close()

	cfg := config.Config{StoreID: "1003", ProxyStore: "42", ProxyToken: "real-token", ProxyReadonly: false}
	srv := proxyServer(cfg, upstream.URL, io.Discard)

	rec := serve(srv, httptest.NewRequest(http.MethodPost, "/api/v3/1003/products", strings.NewReader(`{"name":"x"}`)))
	if rec.Code != http.StatusCreated {
		t.Fatalf("POST status = %d, want 201 (proxied with --proxy-readonly=false)", rec.Code)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("upstream method = %q, want POST", gotMethod)
	}
	if gotBody != `{"name":"x"}` {
		t.Errorf("upstream body = %q, want request body forwarded", gotBody)
	}
}

func TestFallback_StorageServedLocallyEvenWhenProxying(t *testing.T) {
	var upstreamHits int
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		upstreamHits++
		w.WriteHeader(http.StatusOK)
	}))
	defer upstream.Close()

	cfg := config.Config{StoreID: "1003", ProxyStore: "42", ProxyToken: "real-token", ProxyReadonly: true}
	srv := proxyServer(cfg, upstream.URL, io.Discard)

	for _, path := range []string{"/api/v3/1003/storage", "/api/v3/1003/storage/my-key"} {
		rec := serve(srv, httptest.NewRequest(http.MethodGet, path, http.NoBody))
		// /storage is served by the local storage handler, which rejects a
		// missing bearer token with 401 — proving the request was handled
		// locally rather than proxied (a proxied miss would be 501).
		if rec.Code != http.StatusUnauthorized {
			t.Errorf("%s status = %d, want 401 served locally", path, rec.Code)
		}
		msg := decodeErrorMessage(t, rec.Body)
		if strings.Contains(msg, "--proxy-store") {
			t.Errorf("%s: storage response must not suggest proxying: %q", path, msg)
		}
	}
	if upstreamHits != 0 {
		t.Errorf("upstream hits = %d, want 0 — /storage must never be proxied", upstreamHits)
	}
}

func TestFallback_ProxyTokenNeverLogged(t *testing.T) {
	const token = "super-secret-proxy-token-value"
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer upstream.Close()

	var logBuf bytes.Buffer
	cfg := config.Config{StoreID: "1003", ProxyStore: "42", ProxyToken: token, ProxyReadonly: false}
	srv := proxyServer(cfg, upstream.URL, &logBuf)
	// Exercise the log paths: a forwarded read, a forwarded write, and a blocked
	// write (with readonly toggled on) all log at warn.
	serve(srv, httptest.NewRequest(http.MethodGet, "/api/v3/1003/products", http.NoBody))
	serve(srv, httptest.NewRequest(http.MethodPost, "/api/v3/1003/products", strings.NewReader(`{}`)))
	srv.cfg.ProxyReadonly = true
	serve(srv, httptest.NewRequest(http.MethodDelete, "/api/v3/1003/products/1", http.NoBody))

	if strings.Contains(logBuf.String(), token) {
		t.Errorf("proxy token leaked into logs:\n%s", logBuf.String())
	}
}

func TestFallback_Proxy_DropsConnectionNamedHopByHopHeaders(t *testing.T) {
	var gotUpstream http.Header
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotUpstream = r.Header.Clone()
		// Name a hop-by-hop header on the response side too.
		w.Header().Set("Connection", "X-Resp-Hop")
		w.Header().Set("X-Resp-Hop", "secret")
		w.Header().Set("X-Resp-Keep", "kept")
		w.WriteHeader(http.StatusOK)
	}))
	defer upstream.Close()

	cfg := config.Config{StoreID: "1003", ProxyStore: "42", ProxyToken: "real-token", ProxyReadonly: true}
	srv := proxyServer(cfg, upstream.URL, io.Discard)

	req := httptest.NewRequest(http.MethodGet, "/api/v3/1003/products", http.NoBody)
	req.Header.Set("Connection", "X-Req-Hop")
	req.Header.Set("X-Req-Hop", "should-be-dropped")
	req.Header.Set("X-Req-Keep", "kept")
	rec := serve(srv, req)

	if v := gotUpstream.Get("X-Req-Hop"); v != "" {
		t.Errorf("Connection-named request header forwarded upstream: X-Req-Hop = %q, want dropped", v)
	}
	if v := gotUpstream.Get("X-Req-Keep"); v != "kept" {
		t.Errorf("ordinary request header X-Req-Keep = %q, want forwarded", v)
	}
	if v := rec.Header().Get("X-Resp-Hop"); v != "" {
		t.Errorf("Connection-named response header returned to client: X-Resp-Hop = %q, want dropped", v)
	}
	if v := rec.Header().Get("X-Resp-Keep"); v != "kept" {
		t.Errorf("ordinary response header X-Resp-Keep = %q, want forwarded", v)
	}
}

func TestFallback_Proxy_PassesGzipBodyThroughUndecoded(t *testing.T) {
	var raw bytes.Buffer
	gz := gzip.NewWriter(&raw)
	if _, err := gz.Write([]byte(`{"items":[1,2,3]}`)); err != nil {
		t.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		t.Fatal(err)
	}
	gzBytes := raw.Bytes()

	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(gzBytes)
	}))
	defer upstream.Close()

	cfg := config.Config{StoreID: "1003", ProxyStore: "42", ProxyToken: "real-token", ProxyReadonly: true}
	srv := proxyServer(cfg, upstream.URL, io.Discard)

	rec := serve(srv, httptest.NewRequest(http.MethodGet, "/api/v3/1003/products", http.NoBody))

	if ce := rec.Header().Get("Content-Encoding"); ce != "gzip" {
		t.Errorf("Content-Encoding = %q, want gzip preserved (no transparent decompression)", ce)
	}
	if !bytes.Equal(rec.Body.Bytes(), gzBytes) {
		t.Errorf("body was not passed through as raw gzip bytes; got %d bytes, want %d", rec.Body.Len(), len(gzBytes))
	}
}

func TestFallback_UpstreamError_NoPanic(t *testing.T) {
	// Point the proxy at a closed listener so Do() fails immediately.
	upstream := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	closedURL := upstream.URL
	upstream.Close()

	cfg := config.Config{StoreID: "1003", ProxyStore: "42", ProxyToken: "real-token", ProxyReadonly: true}
	srv := proxyServer(cfg, closedURL, io.Discard)

	rec := serve(srv, httptest.NewRequest(http.MethodGet, "/api/v3/1003/products", http.NoBody))
	if rec.Code != http.StatusBadGateway {
		t.Fatalf("status = %d, want 502 on upstream failure", rec.Code)
	}
	msg := decodeErrorMessage(t, rec.Body)
	if !strings.Contains(msg, "failed") {
		t.Errorf("502 message = %q, want a sane upstream-failure message", msg)
	}
}
