package server

import (
	"html"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/appauth"
	"github.com/matthiasbruns/ecwid-go/mock/internal/config"
)

// testSecret is at least 16 bytes so appauth can derive an AES-128 key.
const testSecret = "shell_test_client_secret_0123456789"

func defaultShellConfig() config.Config {
	return config.Config{
		AppURL:       "http://localhost:3000",
		ClientID:     "mock-app",
		ClientSecret: testSecret,
		StoreID:      "1003",
		AuthMode:     config.AuthModeDefault,
		Port:         0,
	}
}

func getShell(t *testing.T, cfg config.Config, rawQuery string) *httptest.ResponseRecorder {
	t.Helper()
	srv := New(cfg, discardLogger())
	target := "/"
	if rawQuery != "" {
		target += "?" + rawQuery
	}
	req := httptest.NewRequest(http.MethodGet, target, http.NoBody)
	rec := httptest.NewRecorder()
	srv.Handler().ServeHTTP(rec, req)
	return rec
}

func TestShell_DefaultMode(t *testing.T) {
	rec := getShell(t, defaultShellConfig(), "")

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); !strings.HasPrefix(ct, "text/html") {
		t.Errorf("Content-Type = %q, want text/html", ct)
	}
	body := rec.Body.String()

	if !strings.Contains(body, "auth: default") {
		t.Error("body missing default auth-mode badge")
	}
	// Default mode injects the payload in the fragment. The iframe src must carry
	// a '#<hex>' fragment on the app URL.
	if !strings.Contains(body, `src="http://localhost:3000#`) {
		t.Errorf("iframe src missing hex fragment; body:\n%s", body)
	}
	// The enhanced-mode warning must NOT appear in default mode.
	if strings.Contains(body, "getPayload() will be undefined") {
		t.Error("default mode should not show the enhanced getPayload warning")
	}
	// The postMessage config must expose the app origin for validation.
	if !strings.Contains(body, `"appOrigin":"http://localhost:3000"`) {
		t.Error("shell-config missing appOrigin for postMessage validation")
	}
}

func TestShell_RedactsAccessToken(t *testing.T) {
	rec := getShell(t, defaultShellConfig(), "")
	body := rec.Body.String()

	// The plaintext mock access_token must never be rendered in the decoded view.
	plain := mockAccessToken("1003")
	if strings.Contains(body, plain) {
		t.Errorf("rendered UI leaked plaintext access_token %q", plain)
	}
	// The redacted form (all but last 4 masked) must be present instead.
	if !strings.Contains(body, redactToken(plain)) {
		t.Errorf("rendered UI missing redacted access_token %q", redactToken(plain))
	}
	// The decoded store_id must still be visible so the table clearly rendered.
	if !strings.Contains(body, ">1003<") {
		t.Error("decoded payload table missing store_id 1003")
	}
}

func TestShell_EnhancedMode(t *testing.T) {
	cfg := defaultShellConfig()
	cfg.AuthMode = config.AuthModeEnhanced

	rec := getShell(t, cfg, "")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	body := rec.Body.String()

	if !strings.Contains(body, "auth: enhanced") {
		t.Error("body missing enhanced auth-mode badge")
	}
	if !strings.Contains(body, "getPayload() will be undefined") {
		t.Error("enhanced mode must surface the getPayload()-undefined warning")
	}
	// The iframe src must carry ?payload= and a cache-killer.
	src := extractIframeSrc(t, body)
	u, err := url.Parse(src)
	if err != nil {
		t.Fatalf("parse iframe src %q: %v", src, err)
	}
	q := u.Query()
	blob := q.Get("payload")
	if blob == "" {
		t.Fatalf("iframe src missing payload param: %s", src)
	}
	if q.Get("cache-killer") == "" {
		t.Errorf("iframe src missing cache-killer: %s", src)
	}
	// The encrypted blob must round-trip through appauth.Decrypt back to the
	// expected store context.
	got, err := appauth.Decrypt(blob, testSecret)
	if err != nil {
		t.Fatalf("Decrypt(payload) error = %v", err)
	}
	if got.StoreID != 1003 {
		t.Errorf("decrypted store_id = %d, want 1003", got.StoreID)
	}
	if got.AccessToken != mockAccessToken("1003") {
		t.Errorf("decrypted access_token = %q, want %q", got.AccessToken, mockAccessToken("1003"))
	}
}

func TestShell_QueryOverrides(t *testing.T) {
	rec := getShell(t, defaultShellConfig(), "store_id=42&lang=de&view_mode=INLINE&app_state=orderId%3A7")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	src := extractIframeSrc(t, rec.Body.String())
	u, err := url.Parse(src)
	if err != nil {
		t.Fatalf("parse iframe src: %v", err)
	}
	got, err := appauth.DecodeHex(strings.TrimPrefix(u.Fragment, "#"))
	if err != nil {
		t.Fatalf("DecodeHex(fragment) error = %v", err)
	}
	if got.StoreID != 42 || got.Lang != "de" || got.ViewMode != "INLINE" || got.AppState != "orderId:7" {
		t.Errorf("overrides not applied: %+v", got)
	}
}

func TestShell_NonCanonicalStoreID(t *testing.T) {
	// store_id=042 parses to 42; the derived tokens must encode the canonical
	// "42", not the raw "042", so the payload stays internally consistent.
	rec := getShell(t, defaultShellConfig(), "store_id=042")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	src := extractIframeSrc(t, rec.Body.String())
	u, err := url.Parse(src)
	if err != nil {
		t.Fatalf("parse iframe src: %v", err)
	}
	got, err := appauth.DecodeHex(u.Fragment)
	if err != nil {
		t.Fatalf("DecodeHex(fragment) error = %v", err)
	}
	if got.StoreID != 42 {
		t.Errorf("StoreID = %d, want 42", got.StoreID)
	}
	if got.AccessToken != mockAccessToken("42") {
		t.Errorf("access_token = %q, want token derived from canonical id 42", got.AccessToken)
	}
}

func TestShell_Devpayload(t *testing.T) {
	fixture := &appauth.Payload{
		StoreID:     7777,
		Lang:        "fr",
		AccessToken: "secret_fixture_token",
		ViewMode:    "PAGE",
	}
	hexPayload, err := appauth.EncodeHex(fixture)
	if err != nil {
		t.Fatalf("EncodeHex: %v", err)
	}

	rec := getShell(t, defaultShellConfig(), "devpayload="+hexPayload)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	body := rec.Body.String()

	// The decoded view must reflect the fixture, not the config defaults.
	if !strings.Contains(body, ">7777<") {
		t.Error("devpayload store_id 7777 not reflected in decoded view")
	}
	// The iframe must drive the app via ?devpayload=, not the fragment.
	src := extractIframeSrc(t, body)
	u, err := url.Parse(src)
	if err != nil {
		t.Fatalf("parse iframe src: %v", err)
	}
	if u.Query().Get("devpayload") != hexPayload {
		t.Errorf("iframe src missing devpayload query param: %s", src)
	}
	if u.Fragment != "" {
		t.Errorf("devpayload mode should not use a fragment, got %q", u.Fragment)
	}
	// The fixture's plaintext token must still be redacted in the display.
	if strings.Contains(body, "secret_fixture_token") {
		t.Error("devpayload view leaked plaintext access_token")
	}
}

func TestShell_InvalidStoreID(t *testing.T) {
	rec := getShell(t, defaultShellConfig(), "store_id=not-a-number")
	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400 for non-numeric store_id", rec.Code)
	}
}

func TestShell_InvalidDevpayload(t *testing.T) {
	rec := getShell(t, defaultShellConfig(), "devpayload=zzzz-not-hex")
	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400 for undecodable devpayload", rec.Code)
	}
}

func TestShell_UnknownPathNotFound(t *testing.T) {
	// The shell is registered at the root only ("GET /{$}"); an unknown path must
	// 404 rather than be swallowed by a catch-all.
	srv := New(defaultShellConfig(), discardLogger())
	req := httptest.NewRequest(http.MethodGet, "/does-not-exist", http.NoBody)
	rec := httptest.NewRecorder()
	srv.Handler().ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404 for unknown path (shell must not be a catch-all)", rec.Code)
	}
}

func TestRedactToken(t *testing.T) {
	tests := []struct{ in, want string }{
		{"", ""},
		{"abcd", "****"},
		{"ab", "****"},
		{"secret_token", "********oken"},
	}
	for _, tt := range tests {
		if got := redactToken(tt.in); got != tt.want {
			t.Errorf("redactToken(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

// extractIframeSrc pulls the app-frame src attribute out of the rendered HTML.
func extractIframeSrc(t *testing.T, body string) string {
	t.Helper()
	const marker = `id="app-frame" src="`
	_, rest, found := strings.Cut(body, marker)
	if !found {
		t.Fatalf("iframe app-frame not found in body")
	}
	end := strings.IndexByte(rest, '"')
	if end < 0 {
		t.Fatalf("iframe src attribute not terminated")
	}
	// html/template HTML-escapes attribute values (e.g. '&' -> '&amp;'); a real
	// browser decodes these before requesting the URL, so undo it for parsing.
	return html.UnescapeString(rest[:end])
}
