// Package webhooks provides typed Ecwid webhook events, signature verification,
// and an [http.Handler] that verifies and dispatches them.
//
// # Security model
//
// Read this before trusting anything a webhook tells you.
//
// The Ecwid webhook signature does NOT cover the request body. It is an HMAC over
// only two fields — "<eventCreated>.<eventId>" — taken from the parsed JSON body
// and keyed with the app's client_secret. Three consequences follow:
//
//   - The payload is unauthenticated. storeId, eventType, entityId and data are
//     not covered by the signature. An attacker who observes one valid
//     (eventCreated, eventId, signature) triple can attach any body to it.
//   - Webhooks are fully replayable. The signature never expires, so a captured
//     request stays valid forever.
//   - Only possession of a signature for a given (eventCreated, eventId) pair is
//     proven — not the integrity of what arrived with it.
//
// Defend in depth:
//
//   - Dedupe on EventID, so a replayed event is processed at most once. See
//     [Deduper].
//   - Enforce a recency window on EventCreated, so old captures are rejected.
//     See [Options.MaxAge].
//   - Re-fetch the entity over the REST API using EntityID rather than trusting
//     Data. Treat Data as a hint about what changed, never as the source of truth.
//
// [Handler] wires up the first two; the third is yours to do in the callback.
//
// # Delivery semantics
//
// Ecwid allows 3s to connect and 10s for the response. [Handler] therefore
// responds before invoking your callback.
//
// Ecwid counts only 200, 201, 202, 204 and 209 as success. Note the traps: 203
// and 208 are failures despite being 2xx, and every 3xx is a failure — an
// endpoint that 301-redirects HTTP to HTTPS silently never receives a webhook.
// [Handler] responds 200 by default.
//
// A failed delivery is retried 27 times over 24 hours (15min, 30min, 45min, 1h,
// 2h, 3h, 4h, 5h, 6h, ... with the final attempt at 24h; the intervals between
// attempts 10 and 26 are elided in Ecwid's documentation and are unconfirmed).
// If an endpoint fails to respond for two weeks, webhooks for the app are
// blocked entirely. Combined with the retries, delivery is at-least-once:
// your callback must be idempotent.
//
// Webhooks fire regardless of what caused the change — creating an order through
// the REST API still emits order.created. Integrations that write back to Ecwid
// in response to a webhook can trigger themselves in a loop.
//
// # Usage
//
//	h, err := webhooks.NewHandler(clientSecret, func(ctx context.Context, e webhooks.Event) {
//		if e.EventType != webhooks.EventOrderCreated {
//			return
//		}
//		// Order events are the exception to re-fetching by EntityID: theirs is
//		// the internal ID, which the order endpoints do not accept. The ID they
//		// want is in the payload.
//		d, err := e.OrderData()
//		if err != nil {
//			return
//		}
//		// Re-fetch rather than trusting the rest of e.Data.
//		order, err := client.Orders.Get(ctx, d.OrderID)
//		// ...
//	}, &webhooks.Options{MaxAge: 5 * time.Minute})
//	if err != nil {
//		return err
//	}
//	mux.Handle("POST /webhooks/ecwid", h)
package webhooks
