package server

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/matthiasbruns/ecwid-go/ecwid/appauth"
	"github.com/matthiasbruns/ecwid-go/ecwid/webhooks"
	"github.com/matthiasbruns/ecwid-go/mock/internal/config"
	"github.com/matthiasbruns/ecwid-go/mock/internal/webhook"
)

// These tests boot the whole mock over a real loopback socket and drive a full
// app session against it — iframe payload -> SDK app storage -> webhook delivery
// — exercising the same appauth and webhooks code a real integration runs. They
// are hermetic: no network, no real store, no headless browser. The app iframed
// by the shell is a tiny embedded HTML page (fakeAppHTML) served locally; the
// HTTP contract is asserted on rendered HTML rather than in a browser, per #71.

// integSecret is the client_secret used across the integration tests. It is at
// least 16 bytes so appauth can derive an AES-128 key and webhooks can sign.
const integSecret = "integration_client_secret_0123456789"

// fakeAppHTML is a stand-in for the developer's iframed app. It exists only so
// the shell has a real, reachable origin to point its iframe at; nothing in it
// executes under test since there is no browser.
//
//go:embed testdata/fakeapp.html
var fakeAppHTML []byte

// ── Harness ─────────────────────────────────────────────────────────────

// newFakeApp serves the embedded fake app page, standing in for the developer's
// app the shell iframes.
func newFakeApp(t *testing.T) *httptest.Server {
	t.Helper()
	app := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write(fakeAppHTML)
	}))
	t.Cleanup(app.Close)
	return app
}

// integConfig is a coherent default-mode config: the required access token
// matches the token the shell injects into the payload (mockAccessToken), so a
// session can read the token out of the payload and use it as the storage bearer
// exactly as the real SDK does.
func integConfig(appURL string) config.Config {
	return config.Config{
		AppURL:       appURL,
		ClientID:     config.DefaultClientID,
		ClientSecret: integSecret,
		StoreID:      testStoreID,
		AuthMode:     config.AuthModeDefault,
		AccessToken:  mockAccessToken(testStoreID),
		Port:         0,
	}
}

// mockOpts carries optional overrides for a booted mock.
type mockOpts struct {
	// upstreamBase points the proxy at a test upstream instead of the real Ecwid
	// API. Empty leaves the default (which no hermetic test may reach).
	upstreamBase string
}

// bootMock builds a Server, wraps it in a real loopback httptest server (which
// binds port 0 and hands back the assigned address), waits for readiness via the
// health endpoint, and returns both. The listener is closed on cleanup.
func bootMock(t *testing.T, cfg config.Config, opts mockOpts) (srv *Server, ts *httptest.Server) {
	t.Helper()
	srv = New(cfg, discardLogger())
	if opts.upstreamBase != "" {
		srv.upstreamBase = opts.upstreamBase
	}
	ts = httptest.NewServer(srv.Handler())
	t.Cleanup(ts.Close)
	waitHealthy(t, ts.URL)
	return srv, ts
}

// waitHealthy polls GET /_mock/health until it answers 200, so a test never
// races the server's startup. It uses a bounded ticker rather than a fixed
// sleep: the loop returns the instant the endpoint is ready.
func waitHealthy(t *testing.T, baseURL string) {
	t.Helper()
	const timeout = 5 * time.Second
	deadline := time.Now().Add(timeout)
	// Bound each probe so a stalled health handler cannot block past the overall
	// deadline (a plain http.Get has no timeout and could hang CI).
	client := &http.Client{Timeout: 250 * time.Millisecond}
	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()
	for {
		resp, err := client.Get(baseURL + "/_mock/health")
		if err == nil {
			healthy := resp.StatusCode == http.StatusOK
			_ = resp.Body.Close()
			if healthy {
				return
			}
		}
		if time.Now().After(deadline) {
			t.Fatalf("mock server did not become healthy within %s (last error: %v)", timeout, err)
		}
		<-ticker.C
	}
}

// getShellHTML fetches the admin shell at / and returns its rendered HTML.
func getShellHTML(t *testing.T, baseURL string) string {
	t.Helper()
	resp, err := http.Get(baseURL + "/")
	if err != nil {
		t.Fatalf("GET shell: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("shell status = %d, want 200", resp.StatusCode)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read shell body: %v", err)
	}
	return string(b)
}

// assertFakeAppReachable confirms the URL the shell points its iframe at is a
// real, reachable app page. In default mode the payload rides the URL fragment,
// which a client never sends, so a plain GET returns the app page unchanged.
func assertFakeAppReachable(t *testing.T, iframeSrc string) {
	t.Helper()
	resp, err := http.Get(iframeSrc)
	if err != nil {
		t.Fatalf("GET iframe app URL %q: %v", iframeSrc, err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("iframe app URL status = %d, want 200", resp.StatusCode)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read iframe app body: %v", err)
	}
	if !strings.Contains(string(b), `id="fake-app"`) {
		t.Error("iframe app URL did not serve the fake app page")
	}
}

// ── Storage helpers ─────────────────────────────────────────────────────

// storageURL builds an absolute app-storage URL for a key (or the collection
// when key is empty).
func storageURL(baseURL, key string) string {
	if key == "" {
		return baseURL + "/api/v3/" + testStoreID + "/storage"
	}
	return baseURL + "/api/v3/" + testStoreID + "/storage/" + key
}

// storageRequest issues a storage request with the given bearer token (omitted
// when token is empty) and returns the live response for the caller to inspect
// and close.
func storageRequest(t *testing.T, method, baseURL, key, token string, body io.Reader) *http.Response {
	t.Helper()
	req, err := http.NewRequest(method, storageURL(baseURL, key), body)
	if err != nil {
		t.Fatalf("build storage request: %v", err)
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("%s storage: %v", method, err)
	}
	return resp
}

// getStorageEntry GETs a key and decodes the {"key","value"} entry, failing if
// the status is not 200.
func getStorageEntry(t *testing.T, baseURL, key, token string) storageEntry {
	t.Helper()
	resp := storageRequest(t, http.MethodGet, baseURL, key, token, nil)
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GET storage %q status = %d, want 200", key, resp.StatusCode)
	}
	var e storageEntry
	if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
		t.Fatalf("decode storage entry: %v", err)
	}
	return e
}

// ── Webhook helpers ─────────────────────────────────────────────────────

// capturedWebhook records one delivery a receiver observed.
type capturedWebhook struct {
	body      []byte
	readErr   error
	signature string
	hasSig    bool
}

// webhookReceiver is a test endpoint that records every delivery and answers
// with a fixed status code, standing in for the developer's webhook handler.
type webhookReceiver struct {
	mu       sync.Mutex
	captured []capturedWebhook
	status   int
	srv      *httptest.Server
}

// newWebhookReceiver starts a receiver that replies with status.
func newWebhookReceiver(t *testing.T, status int) *webhookReceiver {
	t.Helper()
	rc := &webhookReceiver{status: status}
	rc.srv = httptest.NewServer(http.HandlerFunc(rc.handle))
	t.Cleanup(rc.srv.Close)
	return rc
}

func (rc *webhookReceiver) handle(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	rc.mu.Lock()
	rc.captured = append(rc.captured, capturedWebhook{
		body:      body,
		readErr:   err,
		signature: r.Header.Get(webhooks.SignatureHeader),
		hasSig:    len(r.Header.Values(webhooks.SignatureHeader)) > 0,
	})
	rc.mu.Unlock()
	w.WriteHeader(rc.status)
}

// last returns the most recent delivery. Deliveries are synchronous — the mock
// blocks in Fire until the receiver responds — so a delivery is always recorded
// by the time the trigger call returns.
func (rc *webhookReceiver) last(t *testing.T) capturedWebhook {
	t.Helper()
	rc.mu.Lock()
	defer rc.mu.Unlock()
	if len(rc.captured) == 0 {
		t.Fatal("webhook receiver captured no request")
	}
	got := rc.captured[len(rc.captured)-1]
	if got.readErr != nil {
		t.Fatalf("webhook receiver failed to read the delivery body: %v", got.readErr)
	}
	return got
}

// triggerWebhook POSTs to the control API and returns the delivery Result and
// the trigger call's own HTTP status.
func triggerWebhook(t *testing.T, baseURL string, reqBody map[string]any) (res webhook.Result, status int) {
	t.Helper()
	raw, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("marshal trigger request: %v", err)
	}
	resp, err := http.Post(baseURL+"/_mock/webhooks/trigger", "application/json", bytes.NewReader(raw))
	if err != nil {
		t.Fatalf("trigger POST: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			t.Fatalf("decode trigger result: %v", err)
		}
	}
	return res, resp.StatusCode
}

// rawFields decodes a webhook body into its top-level fields, so a test can
// inspect the exact wire shape (entityId typing, data-key presence) a real
// handler receives.
func rawFields(t *testing.T, body []byte) map[string]json.RawMessage {
	t.Helper()
	var m map[string]json.RawMessage
	if err := json.Unmarshal(body, &m); err != nil {
		t.Fatalf("received webhook is not a JSON object: %v", err)
	}
	return m
}

// entityIDTypeInvalid is what wireEntityIDType reports for an entityId that is
// absent or neither a quoted string nor a valid JSON number. It never matches a
// spec's EntityIDType (which is only "string" or "number"), so an omitted or
// malformed entityId always fails an assertion rather than being silently taken
// for a number.
const entityIDTypeInvalid = "invalid"

// wireEntityIDType reports how entityId was serialized: a quoted JSON string, a
// bare JSON number (matching EventSpec.EntityIDType's vocabulary), or
// entityIDTypeInvalid when it is absent or malformed.
func wireEntityIDType(raw json.RawMessage) string {
	trimmed := bytes.TrimSpace(raw)
	switch {
	case len(trimmed) == 0 || bytes.Equal(trimmed, []byte("null")):
		return entityIDTypeInvalid
	case trimmed[0] == '"':
		return webhook.EntityIDString
	default:
		var n json.Number
		if err := json.Unmarshal(trimmed, &n); err != nil {
			return entityIDTypeInvalid
		}
		return webhook.EntityIDNumber
	}
}

// verifyReceived asserts a captured delivery carried a signature that verifies
// with secret over its own eventCreated/eventId — proving the mock and a real
// handler agree on the signature.
func verifyReceived(t *testing.T, rcv capturedWebhook, secret string) {
	t.Helper()
	if !rcv.hasSig {
		t.Fatal("receiver got no signature header on a valid-signature webhook")
	}
	var ev webhooks.Event
	if err := json.Unmarshal(rcv.body, &ev); err != nil {
		t.Fatalf("received webhook body is not a well-formed event: %v", err)
	}
	if err := webhooks.Verify(rcv.signature, ev.EventCreated, ev.EventID, secret); err != nil {
		t.Errorf("received webhook signature did not verify: %v", err)
	}
}

// ── Full session ────────────────────────────────────────────────────────

// TestIntegration_FullSession_DefaultMode drives the whole path the mock exists
// to prove: the shell injects a payload the app can decode, the token in that
// payload authorizes app storage, and a fired webhook is delivered and verifies.
func TestIntegration_FullSession_DefaultMode(t *testing.T) {
	app := newFakeApp(t)
	receiver := newWebhookReceiver(t, http.StatusOK)
	cfg := integConfig(app.URL)
	cfg.WebhookURL = receiver.srv.URL
	_, ts := bootMock(t, cfg, mockOpts{})

	// 1) The shell renders the app iframe carrying the default-mode payload in a
	//    hex fragment; decode it to the store context the SDK would read.
	body := getShellHTML(t, ts.URL)
	src := extractIframeSrc(t, body)
	if !strings.HasPrefix(src, app.URL) {
		t.Fatalf("iframe src %q does not point at the app URL %q", src, app.URL)
	}
	u, err := url.Parse(src)
	if err != nil {
		t.Fatalf("parse iframe src: %v", err)
	}
	payload, err := appauth.DecodeHex(u.Fragment)
	if err != nil {
		t.Fatalf("decode hex payload: %v", err)
	}
	if payload.StoreID != 1003 {
		t.Errorf("payload store_id = %d, want 1003", payload.StoreID)
	}
	wantToken := mockAccessToken(testStoreID)
	if payload.AccessToken != wantToken {
		// Report only redacted forms — test output reaches CI logs.
		t.Errorf("payload access_token = %q, want %q", redactToken(payload.AccessToken), redactToken(wantToken))
	}
	// The token round-trips intact through the payload, yet must never appear in
	// the rendered HTML — only its redacted form does.
	if strings.Contains(body, wantToken) {
		t.Error("rendered shell leaked the plaintext access_token")
	}
	if !strings.Contains(body, redactToken(wantToken)) {
		t.Error("rendered shell missing the redacted access_token")
	}

	// 2) The URL the shell points the iframe at is actually reachable.
	assertFakeAppReachable(t, src)

	// 3) The SDK uses the payload's access_token as the bearer for app storage;
	//    a raw-body PUT round-trips as an unwrapped {"key","value"} entry.
	token := payload.AccessToken
	resp := storageRequest(t, http.MethodPut, ts.URL, "settings", token, strings.NewReader(`{"theme":"dark"}`))
	_ = resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("storage PUT status = %d, want 200", resp.StatusCode)
	}
	got := getStorageEntry(t, ts.URL, "settings", token)
	if got.Key != "settings" || got.Value != `{"theme":"dark"}` {
		t.Errorf("storage round-trip = %+v, want key=settings with the value stored verbatim", got)
	}

	// 4) A fired webhook is delivered to the receiver and its signature verifies
	//    with the same client_secret.
	res, code := triggerWebhook(t, ts.URL, map[string]any{
		"eventType": string(webhooks.EventOrderCreated),
		"signature": "valid",
	})
	if code != http.StatusOK {
		t.Fatalf("trigger status = %d", code)
	}
	if !res.Delivered {
		t.Errorf("order.created not delivered to a 200 receiver: %s", res.Reason)
	}
	verifyReceived(t, receiver.last(t), integSecret)
}

// ── Payload round-trip (both modes) ─────────────────────────────────────

// TestIntegration_PayloadRoundTrip proves both auth modes round-trip through
// ecwid/appauth and that the access_token is redacted in the HTML but intact in
// the decoded payload.
func TestIntegration_PayloadRoundTrip(t *testing.T) {
	wantToken := mockAccessToken(testStoreID)

	tests := []struct {
		name            string
		authMode        string
		wantCacheKiller bool
		decode          func(t *testing.T, src string) *appauth.Payload
	}{
		{
			name:     "default mode hex fragment",
			authMode: config.AuthModeDefault,
			decode: func(t *testing.T, src string) *appauth.Payload {
				t.Helper()
				u, err := url.Parse(src)
				if err != nil {
					t.Fatalf("parse src %q: %v", src, err)
				}
				if u.Fragment == "" {
					t.Fatalf("default mode: no hex fragment in %q", src)
				}
				p, err := appauth.DecodeHex(u.Fragment)
				if err != nil {
					t.Fatalf("DecodeHex: %v", err)
				}
				return p
			},
		},
		{
			name:            "enhanced mode encrypted query",
			authMode:        config.AuthModeEnhanced,
			wantCacheKiller: true,
			decode: func(t *testing.T, src string) *appauth.Payload {
				t.Helper()
				u, err := url.Parse(src)
				if err != nil {
					t.Fatalf("parse src %q: %v", src, err)
				}
				blob := u.Query().Get("payload")
				if blob == "" {
					t.Fatalf("enhanced mode: no payload query param in %q", src)
				}
				p, err := appauth.Decrypt(blob, integSecret)
				if err != nil {
					t.Fatalf("Decrypt: %v", err)
				}
				return p
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			app := newFakeApp(t)
			cfg := integConfig(app.URL)
			cfg.AuthMode = tc.authMode
			_, ts := bootMock(t, cfg, mockOpts{})

			body := getShellHTML(t, ts.URL)
			src := extractIframeSrc(t, body)
			payload := tc.decode(t, src)

			if payload.StoreID != 1003 {
				t.Errorf("store_id = %d, want 1003", payload.StoreID)
			}
			if payload.AccessToken != wantToken {
				// Report only redacted forms — test output reaches CI logs.
				t.Errorf("access_token = %q, want %q", redactToken(payload.AccessToken), redactToken(wantToken))
			}
			// Redacted in the rendered HTML, correct in the decoded payload.
			if strings.Contains(body, wantToken) {
				t.Error("rendered HTML leaked the plaintext access_token")
			}
			if !strings.Contains(body, redactToken(wantToken)) {
				t.Error("rendered HTML missing the redacted access_token")
			}
			if tc.wantCacheKiller {
				u, _ := url.Parse(src)
				if u.Query().Get("cache-killer") == "" {
					t.Errorf("enhanced mode missing cache-killer: %s", src)
				}
			}
		})
	}
}

// ── Storage (as the SDK actually calls it) ──────────────────────────────

// TestIntegration_Storage covers the app-storage contract the JS SDK relies on:
// raw-body values, a real 404 for missing keys, the SDK's literal
// "Content-type: undefined" header, the private/public size limits, and bearer
// enforcement.
func TestIntegration_Storage(t *testing.T) {
	app := newFakeApp(t)
	cfg := integConfig(app.URL)
	_, ts := bootMock(t, cfg, mockOpts{})
	token := mockAccessToken(testStoreID)

	t.Run("raw body round-trips unwrapped", func(t *testing.T) {
		resp := storageRequest(t, http.MethodPut, ts.URL, "greeting", token, strings.NewReader("hello world"))
		_ = resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("PUT status = %d, want 200", resp.StatusCode)
		}
		got := getStorageEntry(t, ts.URL, "greeting", token)
		if got.Key != "greeting" || got.Value != "hello world" {
			t.Errorf("got %+v, want {greeting, hello world}", got)
		}
	})

	t.Run("missing key is 404", func(t *testing.T) {
		resp := storageRequest(t, http.MethodGet, ts.URL, "does-not-exist", token, nil)
		_ = resp.Body.Close()
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("GET missing key status = %d, want 404 (the SDK maps this to null)", resp.StatusCode)
		}
	})

	t.Run("Content-type undefined succeeds", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPut, storageURL(ts.URL, "cfg"), strings.NewReader("v"))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "Bearer "+token)
		// The SDK's ajax() sends this literal string; rejecting it would reject
		// the SDK's own writes.
		req.Header.Set("Content-type", "undefined")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("PUT: %v", err)
		}
		_ = resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("PUT with Content-type: undefined status = %d, want 200", resp.StatusCode)
		}
	})

	t.Run("over-limit private write rejected", func(t *testing.T) {
		over := strings.Repeat("a", maxPrivateValueBytes+1)
		resp := storageRequest(t, http.MethodPut, ts.URL, "big", token, strings.NewReader(over))
		_ = resp.Body.Close()
		if resp.StatusCode != http.StatusRequestEntityTooLarge {
			t.Errorf("over-limit private status = %d, want 413", resp.StatusCode)
		}
	})

	t.Run("over-limit public write rejected", func(t *testing.T) {
		over := strings.Repeat("a", maxPublicValueBytes+1)
		resp := storageRequest(t, http.MethodPut, ts.URL, reservedPublicKey, token, strings.NewReader(over))
		_ = resp.Body.Close()
		if resp.StatusCode != http.StatusRequestEntityTooLarge {
			t.Errorf("over-limit public status = %d, want 413", resp.StatusCode)
		}
	})

	// The exact boundaries must be accepted, so an off-by-one that rejects a
	// value sitting precisely at the limit is caught.
	t.Run("at-limit private write accepted", func(t *testing.T) {
		atLimit := strings.Repeat("a", maxPrivateValueBytes)
		resp := storageRequest(t, http.MethodPut, ts.URL, "exact", token, strings.NewReader(atLimit))
		_ = resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("at-limit private status = %d, want 200", resp.StatusCode)
		}
	})

	t.Run("at-limit public write accepted", func(t *testing.T) {
		atLimit := strings.Repeat("a", maxPublicValueBytes)
		resp := storageRequest(t, http.MethodPut, ts.URL, reservedPublicKey, token, strings.NewReader(atLimit))
		_ = resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("at-limit public status = %d, want 200", resp.StatusCode)
		}
	})

	t.Run("wrong bearer is 401", func(t *testing.T) {
		resp := storageRequest(t, http.MethodGet, ts.URL, "greeting", "not-the-real-token", nil)
		_ = resp.Body.Close()
		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("wrong bearer status = %d, want 401", resp.StatusCode)
		}
	})
}

// ── Webhooks ────────────────────────────────────────────────────────────

// TestIntegration_Webhooks_AllEventTypes fires every one of the 42 catalog
// events and asserts each is delivered, its signature verifies, and its wire
// entityId typing and data-key presence match the catalog.
func TestIntegration_Webhooks_AllEventTypes(t *testing.T) {
	app := newFakeApp(t)
	receiver := newWebhookReceiver(t, http.StatusOK)
	cfg := integConfig(app.URL)
	cfg.WebhookURL = receiver.srv.URL
	_, ts := bootMock(t, cfg, mockOpts{})

	specs := webhook.Catalog()
	if len(specs) != 42 {
		t.Fatalf("catalog has %d events, want 42", len(specs))
	}

	for _, spec := range specs {
		t.Run(string(spec.Type), func(t *testing.T) {
			res, code := triggerWebhook(t, ts.URL, map[string]any{
				"eventType": string(spec.Type),
				"signature": "valid",
			})
			if code != http.StatusOK {
				t.Fatalf("trigger status = %d", code)
			}
			if !res.Delivered {
				t.Errorf("%s not delivered to a 200 receiver: %s", spec.Type, res.Reason)
			}

			rcv := receiver.last(t)
			verifyReceived(t, rcv, integSecret)

			fields := rawFields(t, rcv.body)
			entityID, ok := fields["entityId"]
			if !ok {
				t.Fatalf("%s delivery omitted the entityId field", spec.Type)
			}
			if got := wireEntityIDType(entityID); got != spec.EntityIDType() {
				t.Errorf("%s entityId wire type = %s, want %s (raw %s)",
					spec.Type, got, spec.EntityIDType(), entityID)
			}
			if _, hasData := fields["data"]; hasData != spec.HasData() {
				t.Errorf("%s data key present = %v, want %v", spec.Type, hasData, spec.HasData())
			}
		})
	}
}

// TestIntegration_Webhooks_EntityIDTypingAndDataAbsence pins the specific typing
// quirks #71 calls out, as an oracle independent of the catalog: order.created's
// entityId is a number, application.installed's is a string, and product.created
// carries no data key at all.
func TestIntegration_Webhooks_EntityIDTypingAndDataAbsence(t *testing.T) {
	app := newFakeApp(t)
	receiver := newWebhookReceiver(t, http.StatusOK)
	cfg := integConfig(app.URL)
	cfg.WebhookURL = receiver.srv.URL
	_, ts := bootMock(t, cfg, mockOpts{})

	tests := []struct {
		event       webhooks.EventType
		wantIDType  string
		wantHasData bool
	}{
		{webhooks.EventOrderCreated, webhook.EntityIDNumber, true},
		{webhooks.EventApplicationInstalled, webhook.EntityIDString, false},
		{webhooks.EventProductCreated, webhook.EntityIDNumber, false},
	}

	for _, tc := range tests {
		t.Run(string(tc.event), func(t *testing.T) {
			_, code := triggerWebhook(t, ts.URL, map[string]any{
				"eventType": string(tc.event),
				"signature": "valid",
			})
			if code != http.StatusOK {
				t.Fatalf("trigger status = %d", code)
			}
			fields := rawFields(t, receiver.last(t).body)
			entityID, ok := fields["entityId"]
			if !ok {
				t.Fatalf("%s delivery omitted the entityId field", tc.event)
			}
			if got := wireEntityIDType(entityID); got != tc.wantIDType {
				t.Errorf("%s entityId wire type = %s, want %s", tc.event, got, tc.wantIDType)
			}
			if _, hasData := fields["data"]; hasData != tc.wantHasData {
				t.Errorf("%s data key present = %v, want %v", tc.event, hasData, tc.wantHasData)
			}
		})
	}
}

// TestIntegration_Webhooks_InvalidAndMissingSignature proves fail-closed testing
// works: a wrong signature and an absent one both make webhooks.Verify reject the
// delivery.
func TestIntegration_Webhooks_InvalidAndMissingSignature(t *testing.T) {
	app := newFakeApp(t)
	receiver := newWebhookReceiver(t, http.StatusOK)
	cfg := integConfig(app.URL)
	cfg.WebhookURL = receiver.srv.URL
	_, ts := bootMock(t, cfg, mockOpts{})

	for _, mode := range []string{"invalid", "missing"} {
		t.Run(mode, func(t *testing.T) {
			_, code := triggerWebhook(t, ts.URL, map[string]any{
				"eventType": string(webhooks.EventOrderCreated),
				"signature": mode,
			})
			if code != http.StatusOK {
				t.Fatalf("trigger status = %d", code)
			}
			rcv := receiver.last(t)
			var ev webhooks.Event
			if err := json.Unmarshal(rcv.body, &ev); err != nil {
				t.Fatalf("received body not a well-formed event: %v", err)
			}
			// Verify must reject both a wrong signature and an absent one.
			if err := webhooks.Verify(rcv.signature, ev.EventCreated, ev.EventID, integSecret); err == nil {
				t.Errorf("%s signature: Verify accepted it, want rejection", mode)
			}
			if mode == "missing" && rcv.hasSig {
				t.Error("missing mode still sent a signature header")
			}
		})
	}
}

// TestIntegration_Webhooks_DeliveryClassification asserts Ecwid's delivery
// carve-outs: 200/201/202/204/209 count as delivered while 203/208 and every 3xx
// do not.
func TestIntegration_Webhooks_DeliveryClassification(t *testing.T) {
	app := newFakeApp(t)

	tests := []struct {
		status    int
		delivered bool
	}{
		{http.StatusOK, true},                    // 200
		{http.StatusCreated, true},               // 201
		{http.StatusAccepted, true},              // 202
		{http.StatusNoContent, true},             // 204
		{209, true},                              // Ecwid-specific success code
		{http.StatusNonAuthoritativeInfo, false}, // 203 — a 2xx trap
		{http.StatusAlreadyReported, false},      // 208 — a 2xx trap
		{http.StatusMovedPermanently, false},     // 301 — redirects are never followed
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("status_%d", tc.status), func(t *testing.T) {
			receiver := newWebhookReceiver(t, tc.status)
			cfg := integConfig(app.URL)
			cfg.WebhookURL = receiver.srv.URL
			_, ts := bootMock(t, cfg, mockOpts{})

			res, code := triggerWebhook(t, ts.URL, map[string]any{
				"eventType": string(webhooks.EventProductUpdated),
				"signature": "valid",
			})
			if code != http.StatusOK {
				t.Fatalf("trigger status = %d", code)
			}
			if res.Delivered != tc.delivered {
				t.Errorf("status %d: delivered = %v, want %v (%s)",
					tc.status, res.Delivered, tc.delivered, res.Reason)
			}
			if res.StatusCode != tc.status {
				t.Errorf("result StatusCode = %d, want %d", res.StatusCode, tc.status)
			}
		})
	}
}

// TestIntegration_Webhooks_SlowReceiver_NoPanic drives a receiver too slow to
// answer in time and asserts the mock unwinds cleanly — the server survives, so
// there was no panic in the delivery path.
func TestIntegration_Webhooks_SlowReceiver_NoPanic(t *testing.T) {
	app := newFakeApp(t)

	release := make(chan struct{})
	slow := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Hold the response until released or the request is canceled.
		select {
		case <-release:
			w.WriteHeader(http.StatusOK)
		case <-r.Context().Done():
		}
	}))
	defer slow.Close()

	cfg := integConfig(app.URL)
	cfg.WebhookURL = slow.URL
	_, ts := bootMock(t, cfg, mockOpts{})

	// A short client timeout stands in for giving up on a slow trigger: it
	// cancels the request context, which the mock propagates to its outbound
	// delivery. The point is that this unwinds without a panic.
	client := &http.Client{Timeout: 250 * time.Millisecond}
	resp, err := client.Post(ts.URL+"/_mock/webhooks/trigger", "application/json",
		strings.NewReader(`{"eventType":"order.created","signature":"valid"}`))
	if err == nil {
		_ = resp.Body.Close()
		t.Fatal("expected the slow trigger to time out on the client side")
	}

	// Release the receiver and confirm the server is still healthy — a panic in
	// the delivery path would have taken it down.
	close(release)
	waitHealthy(t, ts.URL)
}

// ── Proxy ───────────────────────────────────────────────────────────────

// TestIntegration_Proxy covers the REST proxy: store-ID rewrite and token swap
// when forwarding, the read-only default blocking writes, and /storage staying
// local even with proxying on.
func TestIntegration_Proxy(t *testing.T) {
	app := newFakeApp(t)

	t.Run("rewrites store id and swaps token", func(t *testing.T) {
		var (
			mu               sync.Mutex
			gotPath, gotAuth string
		)
		upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mu.Lock()
			gotPath, gotAuth = r.URL.Path, r.Header.Get("Authorization")
			mu.Unlock()
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"ok":true}`)
		}))
		defer upstream.Close()

		cfg := integConfig(app.URL)
		cfg.ProxyStore = "999"
		cfg.ProxyToken = "real-store-token"
		_, ts := bootMock(t, cfg, mockOpts{upstreamBase: upstream.URL})

		// An endpoint the mock does not implement locally is proxied upstream.
		resp, err := http.Get(ts.URL + "/api/v3/" + testStoreID + "/orders")
		if err != nil {
			t.Fatalf("proxied GET: %v", err)
		}
		_ = resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("proxied GET status = %d, want 200", resp.StatusCode)
		}

		mu.Lock()
		defer mu.Unlock()
		if want := "/api/v3/999/orders"; gotPath != want {
			t.Errorf("upstream path = %q, want %q (store id not rewritten)", gotPath, want)
		}
		// Assert the swap without printing the bearer token — test output reaches
		// CI logs.
		if gotAuth != "Bearer "+cfg.ProxyToken {
			t.Errorf("upstream Authorization was not swapped to the proxy token (got %q)", redactToken(gotAuth))
		}
	})

	t.Run("read-only default blocks every mutating verb", func(t *testing.T) {
		upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			t.Error("upstream must not be reached for a blocked mutation")
			w.WriteHeader(http.StatusOK)
		}))
		defer upstream.Close()

		cfg := integConfig(app.URL)
		cfg.ProxyStore = "999"
		cfg.ProxyToken = "real-store-token"
		cfg.ProxyReadonly = true
		_, ts := bootMock(t, cfg, mockOpts{upstreamBase: upstream.URL})

		// Every write verb must be blocked, not just POST — only GET/HEAD are read
		// methods.
		for _, method := range []string{http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete} {
			t.Run(method, func(t *testing.T) {
				req, err := http.NewRequest(method, ts.URL+"/api/v3/"+testStoreID+"/orders", strings.NewReader("{}"))
				if err != nil {
					t.Fatalf("build %s request: %v", method, err)
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					t.Fatalf("proxied %s: %v", method, err)
				}
				_ = resp.Body.Close()
				if resp.StatusCode != http.StatusForbidden {
					t.Errorf("proxied %s status = %d, want 403 under the read-only default", method, resp.StatusCode)
				}
			})
		}
	})

	t.Run("storage stays local even with proxying on", func(t *testing.T) {
		upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			t.Error("storage must never be proxied")
			w.WriteHeader(http.StatusOK)
		}))
		defer upstream.Close()

		cfg := integConfig(app.URL)
		cfg.ProxyStore = "999"
		cfg.ProxyToken = "real-store-token"
		cfg.ProxyReadonly = false // even with writes allowed, storage stays local
		_, ts := bootMock(t, cfg, mockOpts{upstreamBase: upstream.URL})
		token := mockAccessToken(testStoreID)

		resp := storageRequest(t, http.MethodPut, ts.URL, "local", token, strings.NewReader("scratch"))
		_ = resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("local storage PUT status = %d, want 200", resp.StatusCode)
		}
		got := getStorageEntry(t, ts.URL, "local", token)
		if got.Value != "scratch" {
			t.Errorf("local storage value = %q, want scratch", got.Value)
		}
	})
}
