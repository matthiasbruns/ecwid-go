package appauth

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
)

// keySize is the AES-128 key length: the first 16 bytes of the client_secret,
// raw ASCII — not hashed or KDF-derived.
const keySize = 16

// Sentinel errors from [Decrypt]. They are deliberately distinct and carry no
// payload bytes, plaintext, or key material.
var (
	// ErrShortSecret means the client_secret is under 16 bytes, so there is no
	// AES-128 key to derive.
	ErrShortSecret = errors.New("appauth: client_secret must be at least 16 bytes")
	// ErrShortBlob means the decoded payload is under 32 bytes, so it cannot hold
	// a 16-byte IV plus at least one ciphertext block.
	ErrShortBlob = errors.New("appauth: encrypted payload too short (need IV + at least one block)")
	// ErrBlockSize means the ciphertext length is not a multiple of the AES block
	// size.
	ErrBlockSize = errors.New("appauth: ciphertext length is not a multiple of the AES block size")
	// ErrPadding means the decrypted plaintext has invalid PKCS#7 padding.
	ErrPadding = errors.New("appauth: invalid PKCS#7 padding")
	// ErrNilPayload means a nil *Payload was passed to an encode function.
	ErrNilPayload = errors.New("appauth: nil payload")
)

// Decrypt decodes an Enhanced Security User Auth payload: a URL-safe base64
// blob whose first 16 bytes are the CBC IV and whose remainder is AES-128-CBC
// ciphertext with PKCS#7 padding, keyed by the first 16 bytes of clientSecret.
func Decrypt(payload, clientSecret string) (*Payload, error) {
	if len(clientSecret) < keySize {
		return nil, ErrShortSecret
	}

	decoded, err := decodeURLSafeBase64(payload)
	if err != nil {
		return nil, fmt.Errorf("appauth: decode payload: %w", err)
	}
	// Need a full IV plus at least one ciphertext block.
	if len(decoded) < aes.BlockSize*2 {
		return nil, ErrShortBlob
	}

	// The IV is the first block; the ciphertext is everything after it. The
	// official C# sample passes the whole blob as ciphertext and then strips
	// leading bytes — it only appears to work because CBC self-synchronizes and
	// corrupts just the first block. Slice properly instead.
	iv := decoded[:aes.BlockSize]
	ciphertext := decoded[aes.BlockSize:]
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, ErrBlockSize
	}

	block, err := aes.NewCipher([]byte(clientSecret[:keySize]))
	if err != nil {
		return nil, fmt.Errorf("appauth: new cipher: %w", err)
	}

	plaintext := make([]byte, len(ciphertext))
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(plaintext, ciphertext)

	plaintext, err = pkcs7Unpad(plaintext)
	if err != nil {
		return nil, err
	}

	var p Payload
	if err := json.Unmarshal(plaintext, &p); err != nil {
		return nil, fmt.Errorf("appauth: unmarshal payload JSON: %w", err)
	}
	return &p, nil
}

// Encrypt produces an Enhanced Security User Auth payload from p, keyed by the
// first 16 bytes of clientSecret. It generates a random IV, so successive calls
// on the same input yield different output. It is the inverse of [Decrypt] and
// exists for the mock admin shell and for callers' own tests.
func Encrypt(p *Payload, clientSecret string) (string, error) {
	if p == nil {
		return "", ErrNilPayload
	}
	if len(clientSecret) < keySize {
		return "", ErrShortSecret
	}

	plaintext, err := json.Marshal(p)
	if err != nil {
		return "", fmt.Errorf("appauth: marshal payload JSON: %w", err)
	}

	block, err := aes.NewCipher([]byte(clientSecret[:keySize]))
	if err != nil {
		return "", fmt.Errorf("appauth: new cipher: %w", err)
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return "", fmt.Errorf("appauth: generate IV: %w", err)
	}

	padded := pkcs7Pad(plaintext, aes.BlockSize)
	ciphertext := make([]byte, len(padded))
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(ciphertext, padded)

	blob := make([]byte, 0, len(iv)+len(ciphertext))
	blob = append(blob, iv...)
	blob = append(blob, ciphertext...)
	return encodeURLSafeBase64(blob), nil
}

// DecodeHex decodes a Default User Auth payload: hex-encoded plaintext JSON from
// the URL fragment. The input is the fragment content without the leading '#'.
//
// The result is unauthenticated — no secret is involved. Re-validate its
// access_token against the Ecwid API before trusting anything it carries.
func DecodeHex(fragment string) (*Payload, error) {
	raw, err := hex.DecodeString(fragment)
	if err != nil {
		return nil, fmt.Errorf("appauth: decode hex fragment: %w", err)
	}

	var p Payload
	if err := json.Unmarshal(raw, &p); err != nil {
		return nil, fmt.Errorf("appauth: unmarshal payload JSON: %w", err)
	}
	return &p, nil
}

// EncodeHex produces a Default User Auth payload from p: hex-encoded plaintext
// JSON, to be placed in the URL fragment after a '#'. It is the inverse of
// [DecodeHex].
func EncodeHex(p *Payload) (string, error) {
	if p == nil {
		return "", ErrNilPayload
	}

	raw, err := json.Marshal(p)
	if err != nil {
		return "", fmt.Errorf("appauth: marshal payload JSON: %w", err)
	}
	return hex.EncodeToString(raw), nil
}

// decodeURLSafeBase64 decodes URL-safe base64 (- and _), tolerating input that
// is missing its trailing padding.
func decodeURLSafeBase64(s string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(padBase64(s))
}

// encodeURLSafeBase64 encodes b as URL-safe base64 (- and _) with '=' padding.
func encodeURLSafeBase64(b []byte) string {
	return base64.URLEncoding.EncodeToString(b)
}

// padBase64 pads s with '=' up to a multiple of 4. Note the trap that has bitten
// existing implementations: when len(s) is already a multiple of 4 this must add
// nothing. The C#-style `len + (4 - len%4)` wrongly appends 4 '=' in that case.
func padBase64(s string) string {
	for len(s)%4 != 0 {
		s += "="
	}
	return s
}

// pkcs7Pad appends PKCS#7 padding so the result is a multiple of blockSize.
func pkcs7Pad(data []byte, blockSize int) []byte {
	pad := blockSize - len(data)%blockSize
	return append(data, bytes.Repeat([]byte{byte(pad)}, pad)...)
}

// pkcs7Unpad removes and validates PKCS#7 padding.
func pkcs7Unpad(data []byte) ([]byte, error) {
	if len(data) == 0 || len(data)%aes.BlockSize != 0 {
		return nil, ErrPadding
	}
	pad := int(data[len(data)-1])
	if pad == 0 || pad > aes.BlockSize || pad > len(data) {
		return nil, ErrPadding
	}
	for _, b := range data[len(data)-pad:] {
		if int(b) != pad {
			return nil, ErrPadding
		}
	}
	return data[:len(data)-pad], nil
}
