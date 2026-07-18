package server

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	// defaultUpstreamBase is the real Ecwid API origin proxied requests target.
	// Only the scheme+host are used; the path is rebuilt per request with the
	// proxy store ID swapped in.
	defaultUpstreamBase = "https://app.ecwid.com"

	// proxyTimeout bounds a single forwarded request end to end so a slow or
	// hung upstream cannot pin a mock connection open indefinitely.
	proxyTimeout = 30 * time.Second

	// apiPrefix is the REST path prefix shared by the mock and the real API.
	apiPrefix = "/api/v3/"
)

// hopByHopHeaders are connection-scoped headers that must not be forwarded
// between the client<->mock and mock<->upstream hops (RFC 7230 §6.1).
var hopByHopHeaders = map[string]struct{}{
	"Connection":          {},
	"Proxy-Connection":    {},
	"Keep-Alive":          {},
	"Proxy-Authenticate":  {},
	"Proxy-Authorization": {},
	"Te":                  {},
	"Trailer":             {},
	"Transfer-Encoding":   {},
	"Upgrade":             {},
}

// newProxyTransport clones the default transport and disables automatic
// compression. Otherwise net/http would add Accept-Encoding: gzip and
// transparently decode gzip responses, rewriting Content-Encoding/Content-Length
// and breaking the intended byte-for-byte passthrough; with it off, whatever the
// upstream returns (gzipped or not) is forwarded unchanged.
func newProxyTransport() http.RoundTripper {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.DisableCompression = true
	return t
}

// handleRESTFallback backstops every /api/v3/{storeId}/... route the mock does
// not simulate locally. With proxying configured it forwards to the real store
// (subject to the read-only gate); otherwise it returns an informative 501.
//
// /storage is always served locally, even when proxying, so a developer's
// scratch state never leaks into the real store's app storage.
func (s *Server) handleRESTFallback(w http.ResponseWriter, r *http.Request) {
	rest := r.PathValue("rest")

	// Local routes take precedence over the proxy. /storage is owned by the
	// storage feature and must stay local; guard it here so it is never
	// forwarded even before a dedicated /storage route exists.
	if isStoragePath(rest) {
		s.writeNotImplemented(w, r, rest, false)
		return
	}

	if !s.cfg.ProxyEnabled() {
		s.writeNotImplemented(w, r, rest, true)
		return
	}

	s.proxyRequest(w, r, rest)
}

// isStoragePath reports whether the endpoint (the path after the store ID) is
// the storage resource or a sub-resource of it.
func isStoragePath(rest string) bool {
	return rest == "storage" || strings.HasPrefix(rest, "storage/")
}

// writeNotImplemented returns 501 naming the actual endpoint. When proxyable is
// true it also states the remedy (run with --proxy-store/--proxy-token); for
// storage — which is deliberately never proxied — it does not, since proxying
// storage would write dev scratch state into the real store.
func (s *Server) writeNotImplemented(w http.ResponseWriter, r *http.Request, rest string, proxyable bool) {
	endpoint := r.Method + " " + apiPrefix + "{storeId}/" + rest
	msg := "mock: " + endpoint + " is not implemented by ecwid-mock."
	if proxyable {
		msg += " Run with --proxy-store and --proxy-token to forward unimplemented endpoints to the real Ecwid API."
	} else {
		msg += " It is served locally and is never proxied."
	}
	writeJSONError(w, http.StatusNotImplemented, msg)
}

// proxyRequest forwards r to the configured real store, rewriting the store ID
// and swapping in the proxy token, and passes the upstream response through.
func (s *Server) proxyRequest(w http.ResponseWriter, r *http.Request, rest string) {
	// Read-only gate: proxied writes mutate the real store and fire real
	// webhooks from it, so block them unless explicitly opted in.
	if s.cfg.ProxyReadonly && !isReadMethod(r.Method) {
		s.log.Warn("proxy blocked mutation",
			"method", r.Method,
			"path", r.URL.Path,
			"target_store", s.cfg.ProxyStore,
		)
		writeJSONError(w, http.StatusForbidden,
			"mock: "+r.Method+" "+apiPrefix+"{storeId}/"+rest+
				" is blocked because --proxy-readonly is on (default). A proxied write mutates the real store "+
				s.cfg.ProxyStore+" and fires real webhooks from it. Re-run with --proxy-readonly=false to allow proxied mutations.")
		return
	}

	// Every proxied request is logged at warn so forwarding is never silent.
	// Only method + path are logged — never the proxy token.
	s.log.Warn("proxy forwarding",
		"method", r.Method,
		"path", r.URL.Path,
		"target_store", s.cfg.ProxyStore,
	)

	targetURL := s.upstreamBase + apiPrefix + s.cfg.ProxyStore + "/" + rest
	if r.URL.RawQuery != "" {
		targetURL += "?" + r.URL.RawQuery
	}

	outReq, err := http.NewRequestWithContext(r.Context(), r.Method, targetURL, r.Body)
	if err != nil {
		writeJSONError(w, http.StatusBadGateway, "mock: could not build proxy request: "+err.Error())
		return
	}
	copyProxyRequestHeaders(outReq.Header, r.Header)
	// Preserve the declared body length so the upstream sees a real
	// Content-Length instead of chunked transfer. NewRequestWithContext cannot
	// infer it from an opaque request body.
	outReq.ContentLength = r.ContentLength
	// Swap the mock's fake access_token for the real proxy token.
	outReq.Header.Set("Authorization", "Bearer "+s.cfg.ProxyToken)

	resp, err := s.proxyClient.Do(outReq)
	if err != nil {
		// Timeout, connection refused, DNS failure, etc. — surface a sane error,
		// never a panic. The error string is the transport's and carries no token.
		s.log.Warn("proxy upstream error", "method", r.Method, "path", r.URL.Path, "error", err)
		writeJSONError(w, http.StatusBadGateway, "mock: proxy request to the real Ecwid API failed: "+err.Error())
		return
	}
	defer func() { _ = resp.Body.Close() }()

	copyProxyResponseHeaders(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	// Stream the body straight through — io.Copy does not buffer the whole body,
	// so there is no memory concern and no reason to cap (which would silently
	// corrupt a large legitimate response). A copy error (client gone, upstream
	// truncated) is not actionable once the status is written.
	_, _ = io.Copy(w, resp.Body)
}

// isReadMethod reports whether m is a non-mutating HTTP method safe to proxy
// under --proxy-readonly.
func isReadMethod(m string) bool {
	return m == http.MethodGet || m == http.MethodHead
}

// hopByHopSet returns the headers that must not be forwarded across a hop: the
// fixed connection-scoped set plus any headers named in src's own Connection
// header (RFC 7230 §6.1). Names are canonicalized for lookup.
func hopByHopSet(src http.Header) map[string]struct{} {
	drop := make(map[string]struct{}, len(hopByHopHeaders))
	for k := range hopByHopHeaders {
		drop[k] = struct{}{}
	}
	for _, v := range src.Values("Connection") {
		for tok := range strings.SplitSeq(v, ",") {
			if tok = strings.TrimSpace(tok); tok != "" {
				drop[http.CanonicalHeaderKey(tok)] = struct{}{}
			}
		}
	}
	return drop
}

// copyProxyRequestHeaders copies client request headers to the upstream request,
// dropping hop-by-hop headers (including those named by Connection) and
// Authorization (which the caller resets to the proxy token) so the mock's fake
// token is never forwarded.
func copyProxyRequestHeaders(dst, src http.Header) {
	drop := hopByHopSet(src)
	drop["Authorization"] = struct{}{}
	copyHeaders(dst, src, drop)
}

// copyProxyResponseHeaders copies upstream response headers back to the client,
// dropping hop-by-hop headers (including those named by Connection).
func copyProxyResponseHeaders(dst, src http.Header) {
	copyHeaders(dst, src, hopByHopSet(src))
}

// copyHeaders copies every header from src to dst except those whose
// canonicalized name is in drop.
func copyHeaders(dst, src http.Header, drop map[string]struct{}) {
	for k, vv := range src {
		if _, skip := drop[http.CanonicalHeaderKey(k)]; skip {
			continue
		}
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

// writeJSONError writes a JSON body of {"errorMessage": msg} with the given
// status, matching Ecwid's error shape.
func writeJSONError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"errorMessage": msg})
}
