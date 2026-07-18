package webhook

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/matthiasbruns/ecwid-go/ecwid/webhooks"
)

const testSecret = "test_client_secret_1234567890"

func newTestTrigger(t *testing.T, url string) *Trigger {
	t.Helper()
	return NewTrigger(Config{ClientSecret: testSecret, StoreID: 1003, URL: url})
}

// Every catalog event must compose into a body that ecwid/webhooks parses back
// intact and whose signature verifies with webhooks.Verify — the mock exercising
// the same code an integration runs is the whole point.
func TestCompose_AllEventsParseAndVerify(t *testing.T) {
	tr := newTestTrigger(t, "")
	for _, spec := range Catalog() {
		t.Run(string(spec.Type), func(t *testing.T) {
			res, err := tr.Compose(Request{EventType: spec.Type})
			if err != nil {
				t.Fatalf("Compose(%q) error: %v", spec.Type, err)
			}

			var e webhooks.Event
			if err := json.Unmarshal([]byte(res.RequestBody), &e); err != nil {
				t.Fatalf("unmarshal body %q: %v", res.RequestBody, err)
			}
			if e.EventType != spec.Type {
				t.Errorf("eventType = %q, want %q", e.EventType, spec.Type)
			}
			if e.EntityID != spec.EntityID {
				t.Errorf("entityId = %q, want %q", e.EntityID, spec.EntityID)
			}
			if e.StoreID != 1003 {
				t.Errorf("storeId = %d, want 1003", e.StoreID)
			}
			if e.EventID != res.EventID || e.EventCreated != res.EventCreated {
				t.Errorf("event/result eventId/eventCreated mismatch")
			}

			if err := webhooks.Verify(res.Signature, res.EventCreated, res.EventID, testSecret); err != nil {
				t.Errorf("Verify() = %v, want nil for a valid signature", err)
			}

			// Data presence must match the fixture: no key at all when absent.
			raw := rawFields(t, res.RequestBody)
			if _, present := raw["data"]; present != spec.HasData() {
				t.Errorf("data key present = %t, want %t", present, spec.HasData())
			}
		})
	}
}

// entityId is a bare number for order/product families and a quoted string for
// application.* — the inconsistency the mock exists to reproduce.
func TestCompose_EntityIDWireType(t *testing.T) {
	tr := newTestTrigger(t, "")

	res, err := tr.Compose(Request{EventType: webhooks.EventOrderCreated})
	if err != nil {
		t.Fatal(err)
	}
	if raw := rawFields(t, res.RequestBody)["entityId"]; raw[0] == '"' {
		t.Errorf("order.created entityId = %s, want a bare number", raw)
	}

	res, err = tr.Compose(Request{EventType: webhooks.EventApplicationInstalled})
	if err != nil {
		t.Fatal(err)
	}
	if raw := rawFields(t, res.RequestBody)["entityId"]; raw[0] != '"' {
		t.Errorf("application.installed entityId = %s, want a quoted string", raw)
	}
}

// product.created carries no data — the body must omit the key, not send an
// empty object.
func TestCompose_ProductHasNoDataKey(t *testing.T) {
	tr := newTestTrigger(t, "")
	res, err := tr.Compose(Request{EventType: webhooks.EventProductCreated})
	if err != nil {
		t.Fatal(err)
	}
	if _, present := rawFields(t, res.RequestBody)["data"]; present {
		t.Errorf("product.created body has a data key: %s", res.RequestBody)
	}
}

func TestCompose_SignatureModes(t *testing.T) {
	tr := newTestTrigger(t, "")

	valid, err := tr.Compose(Request{EventType: webhooks.EventOrderCreated, Mode: SignatureValid})
	if err != nil {
		t.Fatal(err)
	}
	if err := webhooks.Verify(valid.Signature, valid.EventCreated, valid.EventID, testSecret); err != nil {
		t.Errorf("valid mode: Verify = %v, want nil", err)
	}

	invalid, err := tr.Compose(Request{EventType: webhooks.EventOrderCreated, Mode: SignatureInvalid})
	if err != nil {
		t.Fatal(err)
	}
	if invalid.Signature == "" {
		t.Error("invalid mode produced an empty signature; it should be a wrong-but-well-formed one")
	}
	if err := webhooks.Verify(invalid.Signature, invalid.EventCreated, invalid.EventID, testSecret); !errors.Is(err, webhooks.ErrInvalidSignature) {
		t.Errorf("invalid mode: Verify = %v, want ErrInvalidSignature", err)
	}

	missing, err := tr.Compose(Request{EventType: webhooks.EventOrderCreated, Mode: SignatureMissing})
	if err != nil {
		t.Fatal(err)
	}
	if missing.Signature != "" {
		t.Errorf("missing mode signature = %q, want empty", missing.Signature)
	}
	if err := webhooks.Verify(missing.Signature, missing.EventCreated, missing.EventID, testSecret); !errors.Is(err, webhooks.ErrInvalidSignature) {
		t.Errorf("missing mode: Verify = %v, want ErrInvalidSignature", err)
	}
}

func TestCompose_DefaultsAndOverrides(t *testing.T) {
	tr := newTestTrigger(t, "")

	// Empty mode defaults to valid.
	res, err := tr.Compose(Request{EventType: webhooks.EventOrderCreated})
	if err != nil {
		t.Fatal(err)
	}
	if res.SignatureMode != SignatureValid {
		t.Errorf("default mode = %q, want valid", res.SignatureMode)
	}

	// entityId override is honored.
	res, err = tr.Compose(Request{EventType: webhooks.EventOrderCreated, EntityID: "999"})
	if err != nil {
		t.Fatal(err)
	}
	if res.EntityID != "999" {
		t.Errorf("entityId = %q, want 999", res.EntityID)
	}

	// Explicit JSON null forces no data key even for a data-bearing event.
	res, err = tr.Compose(Request{EventType: webhooks.EventOrderCreated, Data: json.RawMessage("null")})
	if err != nil {
		t.Fatal(err)
	}
	if _, present := rawFields(t, res.RequestBody)["data"]; present {
		t.Errorf("explicit null data still emitted a data key: %s", res.RequestBody)
	}
}

func TestCompose_UnknownEvent(t *testing.T) {
	tr := newTestTrigger(t, "")
	if _, err := tr.Compose(Request{EventType: "no.such.event"}); !errors.Is(err, ErrUnknownEvent) {
		t.Errorf("Compose(unknown) = %v, want ErrUnknownEvent", err)
	}
}

func TestFire_Classification(t *testing.T) {
	tests := []struct {
		status        int
		wantDelivered bool
	}{
		{200, true}, {201, true}, {202, true}, {204, true}, {209, true},
		{203, false}, {208, false}, {301, false}, {500, false},
	}
	for _, tt := range tests {
		t.Run(http.StatusText(tt.status), func(t *testing.T) {
			var gotSig string
			var gotBody []byte
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotSig = r.Header.Get(webhooks.SignatureHeader)
				gotBody, _ = io.ReadAll(r.Body)
				w.WriteHeader(tt.status)
			}))
			defer srv.Close()

			tr := newTestTrigger(t, srv.URL)
			res, err := tr.Fire(context.Background(), Request{EventType: webhooks.EventOrderCreated})
			if err != nil {
				t.Fatalf("Fire error: %v", err)
			}
			if res.StatusCode != tt.status {
				t.Errorf("status = %d, want %d", res.StatusCode, tt.status)
			}
			if res.Delivered != tt.wantDelivered {
				t.Errorf("delivered = %t, want %t (%s)", res.Delivered, tt.wantDelivered, res.Reason)
			}
			// The endpoint received a verifiable webhook.
			var e webhooks.Event
			if err := json.Unmarshal(gotBody, &e); err != nil {
				t.Fatalf("endpoint got unparseable body: %v", err)
			}
			if err := webhooks.Verify(gotSig, e.EventCreated, e.EventID, testSecret); err != nil {
				t.Errorf("endpoint could not verify signature: %v", err)
			}
		})
	}
}

func TestFire_MissingSignatureHeaderAbsent(t *testing.T) {
	var hadHeader bool
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, hadHeader = r.Header[http.CanonicalHeaderKey(webhooks.SignatureHeader)]
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	tr := newTestTrigger(t, srv.URL)
	if _, err := tr.Fire(context.Background(), Request{EventType: webhooks.EventOrderCreated, Mode: SignatureMissing}); err != nil {
		t.Fatal(err)
	}
	if hadHeader {
		t.Error("missing mode still sent the signature header")
	}
}

func TestFire_Timeout(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		time.Sleep(300 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	// A client whose response timeout is far below the server's delay, so the
	// delivery fails on timeout rather than waiting Ecwid's real 10s.
	tr := NewTrigger(Config{
		ClientSecret: testSecret,
		StoreID:      1003,
		URL:          srv.URL,
		Client: &http.Client{Transport: &http.Transport{
			ResponseHeaderTimeout: 50 * time.Millisecond,
		}},
	})

	res, err := tr.Fire(context.Background(), Request{EventType: webhooks.EventOrderCreated})
	if err != nil {
		t.Fatalf("Fire returned a Go error, want a failed-delivery Result: %v", err)
	}
	if res.Delivered {
		t.Error("timed-out delivery reported as delivered")
	}
	if res.StatusCode != 0 {
		t.Errorf("statusCode = %d, want 0 on timeout", res.StatusCode)
	}
	if res.Error == "" {
		t.Error("timed-out delivery has no Error set")
	}
}

func TestFire_NoURL(t *testing.T) {
	tr := newTestTrigger(t, "")
	if tr.Enabled() {
		t.Error("Enabled() = true with no URL")
	}
	if _, err := tr.Fire(context.Background(), Request{EventType: webhooks.EventOrderCreated}); !errors.Is(err, ErrNoWebhookURL) {
		t.Errorf("Fire without URL = %v, want ErrNoWebhookURL", err)
	}
}

// rawFields decodes a JSON object into its top-level raw fields.
func rawFields(t *testing.T, body string) map[string]json.RawMessage {
	t.Helper()
	var m map[string]json.RawMessage
	if err := json.Unmarshal([]byte(body), &m); err != nil {
		t.Fatalf("decode fields from %q: %v", body, err)
	}
	return m
}
