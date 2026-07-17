// Package appauth encodes and decodes the payload Ecwid injects into a native
// app's iframe.
//
// # Two auth modes
//
// Ecwid hands an embedded app its store context in one of two mutually
// exclusive ways. They are completely different and this package handles both.
//
// Default User Auth is what every native app gets unless it is switched over by
// Ecwid. The payload is hex-encoded plaintext JSON in the URL fragment:
//
//	GET https://app.example.com/iframe#53035362...
//
// A fragment is never sent to the server, so a Go backend cannot see this
// payload at all — it is read client-side via EcwidApp.getPayload(). Decode it
// with [DecodeHex]; synthesize one with [EncodeHex].
//
// Enhanced Security User Auth is opt-in. The payload is AES-encrypted and sits
// in the query string, so a Go server can act on it:
//
//	GET https://app.example.com/iframe?payload=<urlsafe-b64>&app_state=...&cache-killer=...
//
// Decrypt it with [Decrypt]; produce one with [Encrypt].
//
// # Enhanced mode availability — UNCONFIRMED
//
// Enhanced Security User Auth is gated behind emailing Ecwid
// (ec.apps@lightspeedhq.com) and survives only in 2020-era documentation.
// Whether Ecwid still grants it to new apps is unconfirmed. Do not build on the
// AES path assuming your app has been switched to it — verify with Ecwid first.
// The two-mode distinction itself comes from the legacy ecwid-api-docs
// _add_to_cp.md, which the current docs dropped.
//
// # Default-mode payloads are unauthenticated
//
// A default-mode (hex) payload is plaintext with no secret involved. Anything it
// carries is client-supplied and must not be trusted server-side. In particular,
// re-validate the access_token against the Ecwid API before acting on it; do not
// take the store_id or tokens in a hex payload at face value.
//
// # Tokens are secret
//
// [Payload] carries an access_token and a public_token. Its String and LogValue
// methods redact both, so a Payload is safe to pass to fmt and log/slog. Never
// log the raw field values.
package appauth
