package appauth

import (
	"fmt"
	"io"
	"log/slog"
	"strings"
)

// Payload is the store context Ecwid injects into a native app's iframe. It is
// the decoded form of both the default-mode (hex) and enhanced-mode (AES)
// payloads.
type Payload struct {
	StoreID     int64  `json:"store_id"`
	Lang        string `json:"lang"`
	AccessToken string `json:"access_token"`
	PublicToken string `json:"public_token"` // only with public_storefront scope
	ViewMode    string `json:"view_mode"`    // PAGE | POPUP | INLINE (POPUP currently disabled)
	AppState    string `json:"app_state"`    // URL-encoded; deep-linking only
	Domain      string `json:"domain"`       // undocumented; overrides the API host
}

// String implements [fmt.Stringer] with access_token and public_token redacted,
// so a Payload is safe to print.
func (p Payload) String() string {
	return fmt.Sprintf(
		"appauth.Payload{StoreID:%d Lang:%q AccessToken:%q PublicToken:%q ViewMode:%q AppState:%q Domain:%q}",
		p.StoreID, p.Lang, redact(p.AccessToken), redact(p.PublicToken), p.ViewMode, p.AppState, p.Domain,
	)
}

// Format implements [fmt.Formatter] so every verb renders the redacted form.
// [fmt.Stringer] only covers %v/%s/%+v: bare %#v falls back to the raw struct
// and %d embeds field values in fmt's error text, both of which would leak the
// tokens. Routing all verbs through here closes those holes.
func (p Payload) Format(f fmt.State, verb rune) {
	if verb == 'v' && f.Flag('#') {
		io.WriteString(f, p.goString()) //nolint:errcheck // writing to fmt.State never errors usefully
		return
	}
	io.WriteString(f, p.String()) //nolint:errcheck // writing to fmt.State never errors usefully
}

// goString renders the Go-syntax representation used for %#v, with tokens
// redacted.
func (p Payload) goString() string {
	return fmt.Sprintf(
		"appauth.Payload{StoreID:%d, Lang:%q, AccessToken:%q, PublicToken:%q, ViewMode:%q, AppState:%q, Domain:%q}",
		p.StoreID, p.Lang, redact(p.AccessToken), redact(p.PublicToken), p.ViewMode, p.AppState, p.Domain,
	)
}

// LogValue implements [slog.LogValuer] with access_token and public_token
// redacted, so a Payload is safe to log.
func (p Payload) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int64("store_id", p.StoreID),
		slog.String("lang", p.Lang),
		slog.String("access_token", redact(p.AccessToken)),
		slog.String("public_token", redact(p.PublicToken)),
		slog.String("view_mode", p.ViewMode),
		slog.String("app_state", p.AppState),
		slog.String("domain", p.Domain),
	)
}

// redact masks all but the last 4 characters of a secret, the same all-but-last-4
// policy as config's Config.RedactedToken. Unlike that method it leaves an empty
// secret empty rather than returning "****", so an absent optional token (e.g.
// public_token) is not mistaken for a hidden one.
func redact(s string) string {
	if s == "" {
		return ""
	}
	if len(s) <= 4 {
		return "****"
	}
	return strings.Repeat("*", len(s)-4) + s[len(s)-4:]
}
