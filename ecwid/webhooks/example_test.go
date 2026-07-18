package webhooks_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/matthiasbruns/ecwid-go/ecwid/webhooks"
)

// ExampleVerify checks an Ecwid webhook signature. The signature is
// base64(HMAC-SHA256(key = client_secret, msg = "<eventCreated>.<eventId>")): it
// covers only those two fields, not the request body, so a match proves the
// sender holds the client_secret — nothing about the integrity of the body.
func ExampleVerify() {
	const clientSecret = "test_client_secret"
	// The header value Ecwid sends, plus the two body fields it is computed over.
	const signature = "ekmGPQc/IpCwO6woAiM+L8uvY6chew4n6JVQCiDrEVw="
	eventCreated := int64(1234567)
	eventID := "80aece08-40e8-4145-8764-6c2f0d38678"

	err := webhooks.Verify(signature, eventCreated, eventID, clientSecret)
	fmt.Println("valid:", err == nil)

	// The same signature over a different eventId no longer matches.
	err = webhooks.Verify(signature, eventCreated, "tampered-event-id", clientSecret)
	fmt.Println("valid:", err == nil)

	// Output:
	// valid: true
	// valid: false
}

// ExampleHandler wires a webhook endpoint end to end: NewHandler verifies the
// signature and dispatches verified events to the callback. Here the request is
// composed and signed locally (as the mock server or a test harness would) and
// driven through httptest.
func ExampleHandler() {
	const clientSecret = "test_client_secret"

	h, err := webhooks.NewHandler(clientSecret, func(_ context.Context, e webhooks.Event) {
		// Data is not covered by the signature — treat it as a hint and re-fetch the
		// entity over REST before acting on it. Order events are the exception:
		// EntityID is the internal order ID the order endpoints don't accept, so
		// re-fetch by OrderData().OrderID instead.
		fmt.Println("handled:", e.EventType, "entity", e.EntityID)
	}, &webhooks.Options{MaxAge: 5 * time.Minute})
	if err != nil {
		log.Fatal(err)
	}

	// Compose a webhook exactly as Ecwid would send it, then sign it. Sign shares
	// the MAC that Verify checks, so signing and verification cannot drift.
	e := webhooks.Event{
		EventID:      "80aece08-40e8-4145-8764-6c2f0d38678",
		EventCreated: time.Now().Unix(),
		StoreID:      1003,
		EventType:    webhooks.EventOrderCreated,
		EntityID:     "103878161",
	}
	body, err := json.Marshal(e)
	if err != nil {
		log.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/webhooks/ecwid", bytes.NewReader(body))
	req.Header.Set(webhooks.SignatureHeader, webhooks.Sign(e.EventCreated, e.EventID, clientSecret))

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	// ServeHTTP runs the callback synchronously — after flushing the response but
	// before returning — so the callback's line prints before this status line.
	fmt.Println("status:", rec.Code)

	// Output:
	// handled: order.created entity 103878161
	// status: 200
}
