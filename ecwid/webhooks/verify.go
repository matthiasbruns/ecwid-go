package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
)

// SignatureHeader is the HTTP header Ecwid sends the webhook signature in.
const SignatureHeader = "X-Ecwid-Webhook-Signature"

// ErrInvalidSignature reports that a webhook's signature was missing, malformed,
// or did not match. Reject such a request; it is never a transient condition, so
// a retry cannot fix it.
var ErrInvalidSignature = errors.New("webhooks: invalid signature")

// Verify checks an Ecwid webhook signature, returning nil only if it matches.
//
// The signature is base64(HMAC-SHA256(key = clientSecret, msg = "<eventCreated>.<eventID>")).
// It covers ONLY those two values, not the request body — see the package
// documentation for what that does and does not prove.
//
// Pass the app's client_secret as clientSecret. An access token (secret_*) is
// the wrong key and will never verify.
//
// Verify fails closed: an empty signature is rejected, never skipped. Ecwid's own
// PHP example does the opposite — it reads its $signatureHeaderPresent flag
// before assigning it, so a request carrying no signature header at all skips
// verification and is accepted. Do not reproduce that.
func Verify(signature string, eventCreated int64, eventID, clientSecret string) error {
	if clientSecret == "" {
		return errors.New("webhooks: clientSecret must not be empty")
	}
	if signature == "" {
		return fmt.Errorf("%w: no signature provided", ErrInvalidSignature)
	}

	got, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("%w: not valid base64", ErrInvalidSignature)
	}
	if !hmac.Equal(got, mac(eventCreated, eventID, clientSecret)) {
		return fmt.Errorf("%w: signature mismatch", ErrInvalidSignature)
	}
	return nil
}

// mac computes the raw HMAC-SHA256 of "<eventCreated>.<eventID>".
func mac(eventCreated int64, eventID, clientSecret string) []byte {
	h := hmac.New(sha256.New, []byte(clientSecret))
	// Writes to a hash.Hash are documented never to return an error.
	h.Write(strconv.AppendInt(nil, eventCreated, 10))
	h.Write([]byte("."))
	h.Write([]byte(eventID))
	return h.Sum(nil)
}

// sign produces the signature Verify expects. Kept unexported: real callers only
// ever verify, and an exported signer would invite using this package to forge
// the very events it exists to authenticate.
func sign(eventCreated int64, eventID, clientSecret string) string {
	return base64.StdEncoding.EncodeToString(mac(eventCreated, eventID, clientSecret))
}
