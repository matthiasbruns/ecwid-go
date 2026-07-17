package webhooks

import (
	"errors"
	"testing"
)

// Known-good vector, computed independently of this package's own signing code:
//
//	printf '1234567.80aece08-40e8-4145-8764-6c2f0d38678' |
//	  openssl dgst -sha256 -hmac 'test_client_secret' -binary | base64
const (
	testSecret       = "test_client_secret"
	testEventID      = "80aece08-40e8-4145-8764-6c2f0d38678"
	testEventCreated = int64(1234567)
	testSignature    = "ekmGPQc/IpCwO6woAiM+L8uvY6chew4n6JVQCiDrEVw="
)

func TestVerify_KnownGoodVector(t *testing.T) {
	if err := Verify(testSignature, testEventCreated, testEventID, testSecret); err != nil {
		t.Fatalf("Verify() on known-good vector = %v, want nil", err)
	}
}

// The signature covers "<eventCreated>.<eventID>" and nothing else, so it must
// not verify once either half is altered.
func TestVerify_Rejects(t *testing.T) {
	tests := []struct {
		name                       string
		signature                  string
		eventCreated               int64
		eventID, clientSecret      string
		wantInvalidSignatureSentin bool
	}{
		{
			name:                       "missing signature fails closed",
			signature:                  "",
			eventCreated:               testEventCreated,
			eventID:                    testEventID,
			clientSecret:               testSecret,
			wantInvalidSignatureSentin: true,
		},
		{
			name:                       "tampered signature",
			signature:                  "Xkm GPQc/IpCwO6woAiM+L8uvY6chew4n6JVQCiDrEVw=",
			eventCreated:               testEventCreated,
			eventID:                    testEventID,
			clientSecret:               testSecret,
			wantInvalidSignatureSentin: true,
		},
		{
			name:                       "valid base64 but wrong digest",
			signature:                  "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
			eventCreated:               testEventCreated,
			eventID:                    testEventID,
			clientSecret:               testSecret,
			wantInvalidSignatureSentin: true,
		},
		{
			name:                       "signature is not base64",
			signature:                  "not-base64-$$$",
			eventCreated:               testEventCreated,
			eventID:                    testEventID,
			clientSecret:               testSecret,
			wantInvalidSignatureSentin: true,
		},
		{
			name:                       "wrong secret",
			signature:                  testSignature,
			eventCreated:               testEventCreated,
			eventID:                    testEventID,
			clientSecret:               "wrong_client_secret",
			wantInvalidSignatureSentin: true,
		},
		{
			name:                       "tampered eventCreated",
			signature:                  testSignature,
			eventCreated:               testEventCreated + 1,
			eventID:                    testEventID,
			clientSecret:               testSecret,
			wantInvalidSignatureSentin: true,
		},
		{
			name:                       "tampered eventID",
			signature:                  testSignature,
			eventCreated:               testEventCreated,
			eventID:                    "00000000-0000-0000-0000-000000000000",
			clientSecret:               testSecret,
			wantInvalidSignatureSentin: true,
		},
		{
			// A misconfiguration, not a bad sender: it must not masquerade as a
			// signature mismatch, so callers can tell the two apart.
			name:                       "empty client secret",
			signature:                  testSignature,
			eventCreated:               testEventCreated,
			eventID:                    testEventID,
			clientSecret:               "",
			wantInvalidSignatureSentin: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Verify(tt.signature, tt.eventCreated, tt.eventID, tt.clientSecret)
			if err == nil {
				t.Fatal("Verify() = nil, want an error")
			}
			if got := errors.Is(err, ErrInvalidSignature); got != tt.wantInvalidSignatureSentin {
				t.Errorf("errors.Is(err, ErrInvalidSignature) = %t, want %t (err = %v)", got, tt.wantInvalidSignatureSentin, err)
			}
		})
	}
}

// The "." is a literal separator, not a formatting artifact: the two fields must
// not be concatenable into the same message by different values.
func TestVerify_SeparatorIsNotAmbiguous(t *testing.T) {
	// "1.23" and "12.3" must sign differently.
	a := sign(1, "23", testSecret)
	b := sign(12, "3", testSecret)
	if a == b {
		t.Error("sign(1, \"23\") == sign(12, \"3\"), separator is being ignored")
	}
	if err := Verify(a, 12, "3", testSecret); err == nil {
		t.Error("Verify() accepted a signature for a different (eventCreated, eventID) split")
	}
}

func TestVerify_NegativeAndZeroEventCreated(t *testing.T) {
	for _, created := range []int64{0, -1} {
		if err := Verify(sign(created, testEventID, testSecret), created, testEventID, testSecret); err != nil {
			t.Errorf("Verify() with eventCreated = %d = %v, want nil", created, err)
		}
	}
}

func TestSign_MatchesKnownGoodVector(t *testing.T) {
	if got := sign(testEventCreated, testEventID, testSecret); got != testSignature {
		t.Errorf("sign() = %q, want %q", got, testSignature)
	}
}
