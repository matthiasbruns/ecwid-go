package webhook

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/matthiasbruns/ecwid-go/ecwid/webhooks"
)

// Delivery timeouts mirror Ecwid's: 3s to connect, 10s for the response.
const (
	connectTimeout  = 3 * time.Second
	responseTimeout = 10 * time.Second
)

// maxResponseBody caps how much of the endpoint's response body is read back
// into the result, so a chatty or malicious endpoint cannot balloon a result.
const maxResponseBody = 4 << 10 // 4 KiB

// SignatureMode selects how a triggered webhook is signed, so a developer can
// prove their handler fails closed on a bad or absent signature.
type SignatureMode string

const (
	// SignatureValid signs with the configured client_secret; the signature
	// verifies with [webhooks.Verify].
	SignatureValid SignatureMode = "valid"
	// SignatureInvalid sends a well-formed but wrong signature; [webhooks.Verify]
	// rejects it. A handler that accepts it is not verifying.
	SignatureInvalid SignatureMode = "invalid"
	// SignatureMissing sends no signature header at all. A handler that accepts it
	// fails open — the exact bug Ecwid's own PHP example ships.
	SignatureMissing SignatureMode = "missing"
)

// valid reports whether m is a known signature mode.
func (m SignatureMode) valid() bool {
	switch m {
	case SignatureValid, SignatureInvalid, SignatureMissing:
		return true
	default:
		return false
	}
}

// Request is a webhook trigger. Only EventType is required; EntityID and Data
// default to the event's fixture, and Mode defaults to [SignatureValid].
type Request struct {
	// EventType is the event to fire, e.g. "order.updated".
	EventType webhooks.EventType
	// EntityID overrides the fixture entityId. Empty uses the fixture default.
	EntityID string
	// Data overrides the fixture data. Nil uses the fixture default; a non-nil
	// JSON null ("null") forces no data key.
	Data json.RawMessage
	// Mode selects signature validity. Empty means [SignatureValid].
	Mode SignatureMode
}

// Result is the outcome of a trigger: the composed event, the exact bytes sent,
// and the delivery verdict. It carries the eventCreated/eventId/signature a test
// needs to re-verify with [webhooks.Verify].
type Result struct {
	EventType     webhooks.EventType `json:"eventType"`
	EventID       string             `json:"eventId"`
	EventCreated  int64              `json:"eventCreated"`
	EntityID      string             `json:"entityId"`
	EntityIDType  string             `json:"entityIdType"`
	SignatureMode SignatureMode      `json:"signatureMode"`
	// Signature is the value sent in the signature header, empty for
	// [SignatureMissing]. It is not a secret — Ecwid transmits it in the clear —
	// so it is safe to surface; the client_secret never appears here.
	Signature   string `json:"signature"`
	RequestBody string `json:"requestBody"`
	URL         string `json:"url"`

	// Delivered and Reason classify the response per Ecwid's rules.
	Delivered bool   `json:"delivered"`
	Reason    string `json:"reason"`
	// StatusCode is the endpoint's HTTP status, or 0 when no response arrived.
	StatusCode int `json:"statusCode"`
	// ResponseBody is the endpoint's response, truncated to maxResponseBody.
	ResponseBody string `json:"responseBody"`
	// LatencyMS is the round-trip time in milliseconds.
	LatencyMS int64 `json:"latencyMs"`
	// Error is set when delivery failed before a response (timeout, refused
	// connection, DNS). It is a transport failure, distinct from a delivered-but-
	// unsuccessful status.
	Error string `json:"error,omitempty"`
}

// ErrNoWebhookURL reports that no delivery target is configured. The UI disables
// Fire and the control API returns 409 rather than attempting a delivery.
var ErrNoWebhookURL = errors.New("webhook: no --webhook-url configured")

// ErrUnknownEvent reports an eventType with no fixture.
var ErrUnknownEvent = errors.New("webhook: unknown event type")

// Trigger composes, signs, and delivers webhooks to a configured URL.
type Trigger struct {
	clientSecret string
	storeID      int64
	url          string
	client       *http.Client
	now          func() time.Time
}

// Config configures a [Trigger].
type Config struct {
	// ClientSecret signs webhooks. Required.
	ClientSecret string
	// StoreID is placed in every event's storeId.
	StoreID int64
	// URL is the delivery target. Empty disables delivery; the Trigger still
	// composes events but [Trigger.Fire] returns [ErrNoWebhookURL].
	URL string
	// Client overrides the HTTP client. Optional; a client with Ecwid's connect
	// and response timeouts is used when nil.
	Client *http.Client
	// Now overrides the clock, for tests. Optional; defaults to [time.Now].
	Now func() time.Time
}

// NewTrigger builds a [Trigger] from cfg.
func NewTrigger(cfg Config) *Trigger {
	client := cfg.Client
	if client == nil {
		client = &http.Client{
			Transport: &http.Transport{
				DialContext:           (&net.Dialer{Timeout: connectTimeout}).DialContext,
				ResponseHeaderTimeout: responseTimeout,
			},
			// Do not follow redirects: Ecwid does not, so an endpoint that 301s
			// (e.g. HTTP->HTTPS) silently never receives the webhook. Surfacing
			// that is a headline goal — following the redirect would hide it and
			// mislabel the delivery as a success.
			CheckRedirect: func(*http.Request, []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	}
	now := cfg.Now
	if now == nil {
		now = time.Now
	}
	return &Trigger{
		clientSecret: cfg.ClientSecret,
		storeID:      cfg.StoreID,
		url:          cfg.URL,
		client:       client,
		now:          now,
	}
}

// Enabled reports whether a delivery URL is configured.
func (t *Trigger) Enabled() bool { return t.url != "" }

// URL returns the configured delivery target, for display.
func (t *Trigger) URL() string { return t.url }

// compose builds a signed event and its wire bytes from req, without delivering.
func (t *Trigger) compose(req Request) (Result, []byte, error) {
	spec, ok := Lookup(req.EventType)
	if !ok {
		return Result{}, nil, fmt.Errorf("%w: %q", ErrUnknownEvent, req.EventType)
	}

	mode := req.Mode
	if mode == "" {
		mode = SignatureValid
	}
	if !mode.valid() {
		return Result{}, nil, fmt.Errorf("webhook: invalid signature mode %q (want valid, invalid, or missing)", mode)
	}

	entityID := req.EntityID
	if entityID == "" {
		entityID = spec.EntityID
	}

	// Nil Data uses the fixture default; an explicit JSON null forces no data key.
	data := req.Data
	if data == nil {
		data = spec.Data
	} else if bytes.Equal(bytes.TrimSpace(data), []byte("null")) {
		data = nil
	}

	eventID, err := newUUID()
	if err != nil {
		return Result{}, nil, fmt.Errorf("webhook: generate eventId: %w", err)
	}
	created := t.now().Unix()

	event := webhooks.Event{
		EventID:      eventID,
		EventCreated: created,
		StoreID:      t.storeID,
		EventType:    req.EventType,
		EntityID:     entityID,
		Data:         data,
	}
	body, err := json.Marshal(event)
	if err != nil {
		return Result{}, nil, fmt.Errorf("webhook: marshal event: %w", err)
	}

	res := Result{
		EventType:     req.EventType,
		EventID:       eventID,
		EventCreated:  created,
		EntityID:      entityID,
		EntityIDType:  spec.EntityIDType(),
		SignatureMode: mode,
		Signature:     signature(mode, created, eventID, t.clientSecret),
		RequestBody:   string(body),
		URL:           t.url,
	}
	return res, body, nil
}

// signature produces the header value for a mode: a valid signature, a
// deliberately wrong one, or none.
func signature(mode SignatureMode, eventCreated int64, eventID, clientSecret string) string {
	switch mode {
	case SignatureMissing:
		return ""
	case SignatureInvalid:
		// A real HMAC under a wrong key: well-formed base64 that webhooks.Verify
		// rejects with a digest mismatch, exactly like a spoofed sender.
		return webhooks.Sign(eventCreated, eventID, clientSecret+"-wrong-key")
	default:
		return webhooks.Sign(eventCreated, eventID, clientSecret)
	}
}

// Compose builds and signs an event without delivering it. Useful for tests and
// for showing the exact payload before firing.
func (t *Trigger) Compose(req Request) (Result, error) {
	res, _, err := t.compose(req)
	return res, err
}

// Fire composes, signs, and POSTs a webhook, returning the delivery verdict. It
// returns [ErrNoWebhookURL] if no target is configured and [ErrUnknownEvent] for
// an unknown event; a reachable endpoint that returns any status yields a Result
// with no error, its Delivered/Reason set by Ecwid's classification.
func (t *Trigger) Fire(ctx context.Context, req Request) (Result, error) {
	if !t.Enabled() {
		return Result{}, ErrNoWebhookURL
	}

	res, body, err := t.compose(req)
	if err != nil {
		return Result{}, err
	}

	// Bound the whole exchange, so a stalled body read cannot exceed Ecwid's
	// budget: ResponseHeaderTimeout caps only time-to-first-header, not the body.
	ctx, cancel := context.WithTimeout(ctx, connectTimeout+responseTimeout)
	defer cancel()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, t.url, bytes.NewReader(body))
	if err != nil {
		return Result{}, fmt.Errorf("webhook: build request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	if res.SignatureMode != SignatureMissing {
		httpReq.Header.Set(webhooks.SignatureHeader, res.Signature)
	}

	start := t.now()
	resp, err := t.client.Do(httpReq)
	if err != nil {
		res.LatencyMS = t.now().Sub(start).Milliseconds()
		res.Reason = "delivery failed before a response (timeout, refused connection, or DNS failure)"
		res.Error = deliveryError(err)
		return res, nil
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, _ := io.ReadAll(io.LimitReader(resp.Body, maxResponseBody))
	res.LatencyMS = t.now().Sub(start).Milliseconds()
	res.StatusCode = resp.StatusCode
	res.ResponseBody = string(respBody)

	verdict := classify(resp.StatusCode)
	res.Delivered = verdict.Delivered
	res.Reason = verdict.Reason
	return res, nil
}

// deliveryError renders a transport error, mapping a deadline to a message that
// names the response timeout rather than leaking a raw context string.
func deliveryError(err error) string {
	var netErr net.Error
	if errors.Is(err, context.DeadlineExceeded) || (errors.As(err, &netErr) && netErr.Timeout()) {
		return fmt.Sprintf("timed out (connect %s / response %s)", connectTimeout, responseTimeout)
	}
	return err.Error()
}

// newUUID returns a random RFC 4122 version-4 UUID string.
func newUUID() (string, error) {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", err
	}
	b[6] = (b[6] & 0x0f) | 0x40 // version 4
	b[8] = (b[8] & 0x3f) | 0x80 // variant 10
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16]), nil
}
