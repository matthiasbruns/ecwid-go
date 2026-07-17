package appauth

import (
	"errors"
	"strings"
	"testing"
)

// clientSecret used across the crypto tests. Its first 16 bytes are the AES key.
const testClientSecret = "test_client_secret_0123456789"

// fixedVector is a committed Enhanced-mode payload generated once with the
// documented algorithm (AES-128-CBC, key = first 16 bytes of client_secret,
// IV = "appauth-fixed-iv", PKCS#7, URL-safe base64). It is NOT the docs' hex
// example — that string is hex, not base64, and is not a valid AES payload.
//
// Its plaintext begins with `{"access_token":`, so the C#-style bug that
// corrupts the first ciphertext block would break JSON parsing and fail
// TestDecrypt_FixedVector.
const fixedVector = "YXBwYXV0aC1maXhlZC1pdhIzjFPUtYpLNVCK76EMJ_yHnHmCWZinMHcLWoW0fvWXh4f_Bh3q9o8Qh0IhXHvVqJkC8FpvYcoFcfcSTiHKZwPmEoo72ZNHVmTN2UoMwNmfI1oyaqW-qoXmmUcj75Eig1eQ7NFgc3cHQ5_JpzqFN9ZitEv9wWBPUy8E7_fGI0IpPXIcsMadp-azBhdITMq6mnl6rpir2wZYiDyty3Q5arXRSn9627zTiQJAaHSzr7mGpe3GwJsO-p37z_w7VXLZwA=="

func samplerPayload() *Payload {
	return &Payload{
		StoreID:     1003,
		Lang:        "en",
		AccessToken: "secret_abc123def456token",
		PublicToken: "public_zzz999yyy888token",
		ViewMode:    "PAGE",
		AppState:    "orderId%3A12",
		Domain:      "app.example.test",
	}
}

func assertEqual(t *testing.T, got, want *Payload) {
	t.Helper()
	if *got != *want {
		t.Fatalf("payload mismatch:\n got = %+v\nwant = %+v", *got, *want)
	}
}

// TestDecrypt_FixedVector cross-checks against a committed ciphertext produced
// by the documented algorithm. It also proves correct IV handling: a
// first-block-corrupting decode would fail here.
func TestDecrypt_FixedVector(t *testing.T) {
	got, err := Decrypt(fixedVector, testClientSecret)
	if err != nil {
		t.Fatalf("Decrypt() error = %v", err)
	}
	assertEqual(t, got, samplerPayload())
}

func TestEncryptDecrypt_RoundTrip(t *testing.T) {
	want := samplerPayload()

	enc, err := Encrypt(want, testClientSecret)
	if err != nil {
		t.Fatalf("Encrypt() error = %v", err)
	}
	got, err := Decrypt(enc, testClientSecret)
	if err != nil {
		t.Fatalf("Decrypt() error = %v", err)
	}
	assertEqual(t, got, want)
}

// TestEncrypt_RandomIV asserts a fresh IV per call: identical input yields
// different ciphertext, and both still decrypt back to the original.
func TestEncrypt_RandomIV(t *testing.T) {
	p := samplerPayload()

	a, err := Encrypt(p, testClientSecret)
	if err != nil {
		t.Fatalf("Encrypt() error = %v", err)
	}
	b, err := Encrypt(p, testClientSecret)
	if err != nil {
		t.Fatalf("Encrypt() error = %v", err)
	}
	if a == b {
		t.Fatal("Encrypt produced identical output twice; IV is not random")
	}
	for _, enc := range []string{a, b} {
		got, err := Decrypt(enc, testClientSecret)
		if err != nil {
			t.Fatalf("Decrypt() error = %v", err)
		}
		assertEqual(t, got, p)
	}
}

func TestDecrypt_Errors(t *testing.T) {
	// A valid blob to mutate for the block-size case: reuse the round-trip output.
	valid, err := Encrypt(samplerPayload(), testClientSecret)
	if err != nil {
		t.Fatalf("Encrypt() error = %v", err)
	}

	tests := []struct {
		name    string
		payload string
		secret  string
		want    error
	}{
		{"short secret", fixedVector, "too-short", ErrShortSecret},
		{"short blob", encodeURLSafeBase64([]byte("only-a-few-bytes")), testClientSecret, ErrShortBlob},
		{"bad padding", tamperLastBlock(t, valid), testClientSecret, ErrPadding},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Decrypt(tt.payload, tt.secret)
			if err == nil {
				t.Fatal("expected error, got nil")
			}
			if tt.want != nil && !errors.Is(err, tt.want) {
				t.Fatalf("error = %v, want %v", err, tt.want)
			}
		})
	}
}

// TestDecrypt_NonMultipleCiphertext feeds a blob whose ciphertext is 8 bytes
// short of a block boundary and asserts ErrBlockSize specifically.
func TestDecrypt_NonMultipleCiphertext(t *testing.T) {
	// 16-byte IV + 24 bytes: 24 is not a multiple of 16.
	blob := make([]byte, 16+24)
	for i := range blob {
		blob[i] = byte(i)
	}
	_, err := Decrypt(encodeURLSafeBase64(blob), testClientSecret)
	if !errors.Is(err, ErrBlockSize) {
		t.Fatalf("error = %v, want ErrBlockSize", err)
	}
}

// tamperLastBlock deterministically corrupts the PKCS#7 padding without changing
// the length. In CBC, plaintext block N depends on ciphertext block N XOR
// ciphertext block N-1, so flipping a byte of block N-1 flips exactly the
// corresponding byte of plaintext block N. Flipping the last byte of the
// second-to-last ciphertext block (offset len-17) thus flips only the final
// padding byte — always invalid. Flipping the final ciphertext block instead
// would randomize the whole last plaintext block via the avalanche effect,
// leaving a ~1/256 chance the result is still valid padding and the test flakes.
func tamperLastBlock(t *testing.T, enc string) string {
	t.Helper()
	decoded, err := decodeURLSafeBase64(enc)
	if err != nil {
		t.Fatalf("decode: %v", err)
	}
	decoded[len(decoded)-17] ^= 0xFF
	return encodeURLSafeBase64(decoded)
}

func TestEncrypt_NilAndShortSecret(t *testing.T) {
	if _, err := Encrypt(nil, testClientSecret); !errors.Is(err, ErrNilPayload) {
		t.Fatalf("Encrypt(nil) error = %v, want ErrNilPayload", err)
	}
	if _, err := Encrypt(samplerPayload(), "short"); !errors.Is(err, ErrShortSecret) {
		t.Fatalf("Encrypt(short secret) error = %v, want ErrShortSecret", err)
	}
}

func TestEncodeDecodeHex_RoundTrip(t *testing.T) {
	want := samplerPayload()

	enc, err := EncodeHex(want)
	if err != nil {
		t.Fatalf("EncodeHex() error = %v", err)
	}
	if strings.ContainsAny(enc, "#") {
		t.Fatal("EncodeHex output must be bare hex without a leading '#'")
	}
	got, err := DecodeHex(enc)
	if err != nil {
		t.Fatalf("DecodeHex() error = %v", err)
	}
	assertEqual(t, got, want)
}

func TestDecodeHex_Errors(t *testing.T) {
	if _, err := DecodeHex("nothex!!"); err == nil {
		t.Fatal("expected error for invalid hex")
	}
	// Valid hex, but not JSON.
	if _, err := DecodeHex("6e6f742d6a736f6e"); err == nil {
		t.Fatal("expected error for hex that is not JSON")
	}
}

func TestEncodeHex_Nil(t *testing.T) {
	if _, err := EncodeHex(nil); !errors.Is(err, ErrNilPayload) {
		t.Fatalf("EncodeHex(nil) error = %v, want ErrNilPayload", err)
	}
}

// TestPadBase64 guards the off-by-four trap: a length already a multiple of 4
// must gain no padding.
func TestPadBase64(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"a", "a==="},
		{"ab", "ab=="},
		{"abc", "abc="},
		{"abcd", "abcd"},         // len%4 == 0: no padding added
		{"abcdefgh", "abcdefgh"}, // len%4 == 0: no padding added
	}
	for _, tt := range tests {
		if got := padBase64(tt.in); got != tt.want {
			t.Errorf("padBase64(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
