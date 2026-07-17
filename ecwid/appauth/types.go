package appauth

import (
	"fmt"
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

// redact masks all but the last 4 characters of a secret, matching
// config.RedactedToken. An empty secret stays empty.
func redact(s string) string {
	if s == "" {
		return ""
	}
	if len(s) <= 4 {
		return "****"
	}
	return strings.Repeat("*", len(s)-4) + s[len(s)-4:]
}
