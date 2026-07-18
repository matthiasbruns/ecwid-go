package server

import (
	"crypto/rand"
	"embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/matthiasbruns/ecwid-go/ecwid/appauth"
	"github.com/matthiasbruns/ecwid-go/mock/internal/config"
)

// SDK / payload defaults for the shell. These are developer-facing conveniences,
// not part of any published contract.
const (
	// defaultLang is the store language placed in the payload when the shell is
	// loaded without an explicit lang override.
	defaultLang = "en"

	// defaultViewMode is the render mode Ecwid uses for settings pages. POPUP is
	// currently disabled by Ecwid and INLINE is rare, so PAGE is the sensible
	// default.
	defaultViewMode = "PAGE"

	// sdkURL is the real Ecwid native-app JS SDK. The docs and the sample app
	// both pin 1.3.0; 1.3.1 exists on the CDN but is undocumented. The developer
	// includes this script in their own app — the shell only surfaces it for
	// reference.
	sdkURL = "https://djqizrxa6f10j.cloudfront.net/ecwid-sdk/js/1.3.0/ecwid-app.js"
)

//go:embed templates/shell.html
var shellFS embed.FS

// shellTmpl is parsed once at startup. A parse failure is a programming error in
// the embedded template, so panic via template.Must rather than deferring it to
// the first request.
var shellTmpl = template.Must(template.ParseFS(shellFS, "templates/shell.html"))

// payloadView is the decoded payload as shown in the UI, with both tokens
// redacted. The raw tokens never reach the template — only this view does.
type payloadView struct {
	StoreID     int64
	Lang        string
	AccessToken string // redacted
	PublicToken string // redacted
	ViewMode    string
	AppState    string
	Domain      string
}

// shellView is the full template model for one render of the admin shell.
type shellView struct {
	AppURL    string
	AppOrigin string
	ClientID  string
	AuthMode  string
	Enhanced  bool

	// IframeURL is the fully-built src for the app iframe, carrying the payload
	// in the fragment (default mode), the query (enhanced mode), or as a
	// devpayload query param (fixture hook). It is template.URL because it is
	// server-built from config and must not be filtered by html/template.
	IframeURL template.URL

	// Payload is the decoded, token-redacted payload for display.
	Payload payloadView

	// RawKind labels the raw payload transport ("hex fragment" or "encrypted
	// query blob"). RawPayload is the copy-to-clipboard value; it is the same
	// bytes injected into the iframe, so it is offered via a copy affordance
	// rather than printed prominently.
	RawKind    string
	RawPayload string

	// CacheKiller is the enhanced-mode cache-busting query value, shown for
	// reference. Empty in default mode.
	CacheKiller string

	// Control values reflected back into the form inputs.
	StoreID  string
	Lang     string
	AppState string
	ViewMode string

	// SdkURL is the reference SDK script URL.
	SdkURL string

	// ConfigJS is the JSON config consumed by the page's postMessage listener:
	// the app origin to validate against and the client_id namespace to match.
	ConfigJS template.JS
}

// handleShell serves GET / — the admin shell that iframes the developer's app
// with a correctly-built payload for the configured auth mode.
//
// Query overrides let the developer re-render without restarting the server:
// store_id, lang, app_state, view_mode. The devpayload fixture hook (a hex
// payload) overrides all of them and drives the iframe via a ?devpayload= query
// param instead of the fragment, so tests can assert on a known payload without
// touching the URL hash.
func (s *Server) handleShell(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	view, err := s.buildShellView(q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := shellTmpl.ExecuteTemplate(w, "shell.html", view); err != nil {
		// Headers/partial body may already be flushed; there is nothing useful to
		// send the client, so just record it.
		s.log.Error("render admin shell", "error", err)
	}
}

// buildShellView assembles the template model from config plus query overrides.
// It returns an error only for client-correctable input (a non-numeric store_id
// or an undecodable devpayload); everything else falls back to config defaults.
func (s *Server) buildShellView(q url.Values) (shellView, error) {
	appOrigin, err := originOf(s.cfg.AppURL)
	if err != nil {
		return shellView{}, fmt.Errorf("app URL %q has no origin: %w", s.cfg.AppURL, err)
	}

	var (
		payload     *appauth.Payload
		iframeURL   string
		rawKind     string
		rawPayload  string
		cacheKiller string
	)

	if dev := q.Get("devpayload"); dev != "" {
		// Fixture hook: decode the supplied hex payload and drive the iframe via a
		// ?devpayload= query param, bypassing the fragment entirely.
		payload, err = appauth.DecodeHex(dev)
		if err != nil {
			return shellView{}, fmt.Errorf("invalid devpayload: %w", err)
		}
		iframeURL, err = withQueryParam(s.cfg.AppURL, "devpayload", dev)
		if err != nil {
			return shellView{}, err
		}
		rawKind = "hex fragment (via devpayload)"
		rawPayload = dev
	} else {
		payload, err = s.payloadFromQuery(q)
		if err != nil {
			return shellView{}, err
		}
		iframeURL, rawKind, rawPayload, cacheKiller, err = s.buildIframeURL(payload)
		if err != nil {
			return shellView{}, err
		}
	}

	cfgJSON, err := json.Marshal(map[string]string{
		"appOrigin": appOrigin,
		"clientId":  s.cfg.ClientID,
	})
	if err != nil {
		return shellView{}, fmt.Errorf("marshal shell config: %w", err)
	}

	return shellView{
		AppURL:      s.cfg.AppURL,
		AppOrigin:   appOrigin,
		ClientID:    s.cfg.ClientID,
		AuthMode:    s.cfg.AuthMode,
		Enhanced:    s.cfg.AuthMode == config.AuthModeEnhanced,
		IframeURL:   template.URL(iframeURL), //nolint:gosec // server-built from operator config, not user input
		Payload:     redactedView(payload),
		RawKind:     rawKind,
		RawPayload:  rawPayload,
		CacheKiller: cacheKiller,
		StoreID:     strconv.FormatInt(payload.StoreID, 10),
		Lang:        payload.Lang,
		AppState:    payload.AppState,
		ViewMode:    payload.ViewMode,
		SdkURL:      sdkURL,
		ConfigJS:    template.JS(cfgJSON), //nolint:gosec // JSON-encoded, HTML-escaped by encoding/json
	}, nil
}

// payloadFromQuery builds the payload from config defaults, applying the
// store_id / lang / app_state / view_mode query overrides. The tokens are
// deterministic mock values so the simulated REST API (#5) can recognize them.
func (s *Server) payloadFromQuery(q url.Values) (*appauth.Payload, error) {
	storeIDStr := firstNonEmpty(q.Get("store_id"), s.cfg.StoreID)
	storeID, err := strconv.ParseInt(storeIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid store_id %q: must be an integer", storeIDStr)
	}

	return &appauth.Payload{
		StoreID:     storeID,
		Lang:        firstNonEmpty(q.Get("lang"), defaultLang),
		AccessToken: mockAccessToken(storeIDStr),
		PublicToken: mockPublicToken(storeIDStr),
		ViewMode:    firstNonEmpty(q.Get("view_mode"), defaultViewMode),
		AppState:    q.Get("app_state"),
		// Domain is left empty: forcing the SDK's API host at the app breaks apps
		// that do not expect it. The shell documents the getEcwidSdkApiDomain()
		// and domain-field overrides instead of imposing them.
	}, nil
}

// buildIframeURL encodes the payload for the configured auth mode and returns
// the iframe src, a human label for the raw transport, the raw payload string
// (for the copy affordance), and the enhanced-mode cache-killer ("" in default
// mode).
func (s *Server) buildIframeURL(p *appauth.Payload) (iframeURL, rawKind, rawPayload, cacheKiller string, err error) {
	switch s.cfg.AuthMode {
	case config.AuthModeEnhanced:
		blob, err := appauth.Encrypt(p, s.cfg.ClientSecret)
		if err != nil {
			return "", "", "", "", fmt.Errorf("encrypt payload: %w", err)
		}
		cacheKiller, err = randomCacheKiller()
		if err != nil {
			return "", "", "", "", err
		}
		u, err := url.Parse(s.cfg.AppURL)
		if err != nil {
			return "", "", "", "", fmt.Errorf("parse app URL: %w", err)
		}
		query := u.Query()
		query.Set("payload", blob)
		if p.AppState != "" {
			query.Set("app_state", p.AppState)
		}
		query.Set("cache-killer", cacheKiller)
		u.RawQuery = query.Encode()
		return u.String(), "encrypted query blob (?payload=)", blob, cacheKiller, nil

	default: // config.AuthModeDefault
		hexPayload, err := appauth.EncodeHex(p)
		if err != nil {
			return "", "", "", "", fmt.Errorf("hex-encode payload: %w", err)
		}
		u, err := url.Parse(s.cfg.AppURL)
		if err != nil {
			return "", "", "", "", fmt.Errorf("parse app URL: %w", err)
		}
		u.Fragment = hexPayload
		return u.String(), "hex fragment (#)", hexPayload, "", nil
	}
}

// redactedView copies the payload into its display form with both tokens masked.
func redactedView(p *appauth.Payload) payloadView {
	return payloadView{
		StoreID:     p.StoreID,
		Lang:        p.Lang,
		AccessToken: redactToken(p.AccessToken),
		PublicToken: redactToken(p.PublicToken),
		ViewMode:    p.ViewMode,
		AppState:    p.AppState,
		Domain:      p.Domain,
	}
}

// withQueryParam returns rawURL with key=value set on its query string.
func withQueryParam(rawURL, key, value string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("parse app URL: %w", err)
	}
	q := u.Query()
	q.Set(key, value)
	u.RawQuery = q.Encode()
	return u.String(), nil
}

// originOf returns the scheme://host origin used to validate postMessage events.
func originOf(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	if u.Scheme == "" || u.Host == "" {
		return "", fmt.Errorf("not an absolute URL")
	}
	return u.Scheme + "://" + u.Host, nil
}

// randomCacheKiller returns a short random hex string used to defeat caching of
// the enhanced-mode iframe request, mirroring Ecwid's cache-killer param.
func randomCacheKiller() (string, error) {
	buf := make([]byte, 6)
	if _, err := rand.Read(buf); err != nil {
		return "", fmt.Errorf("generate cache-killer: %w", err)
	}
	return hex.EncodeToString(buf), nil
}

// mockAccessToken derives a deterministic stand-in access_token for a store, so
// the payload carries a plausible token the simulated REST API can recognize.
func mockAccessToken(storeID string) string {
	return "secret_mock_" + storeID + "_access"
}

// mockPublicToken derives a deterministic stand-in public_token for a store.
func mockPublicToken(storeID string) string {
	return "public_mock_" + storeID + "_public"
}

// firstNonEmpty returns a if it is non-empty, else b.
func firstNonEmpty(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

// redactToken masks all but the last four characters, matching the codebase's
// all-but-last-4 policy. An empty token stays empty so an absent optional token
// (e.g. public_token) is not mistaken for a hidden one.
func redactToken(t string) string {
	switch {
	case t == "":
		return ""
	case len(t) <= 4:
		return "****"
	default:
		return strings.Repeat("*", len(t)-4) + t[len(t)-4:]
	}
}
