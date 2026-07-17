package webhooks

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// maxBodyBytes caps a webhook body. Ecwid's payloads are a few hundred bytes;
// this is generous enough to never reject a real one and small enough that an
// unauthenticated POST cannot exhaust memory.
const maxBodyBytes = 1 << 20 // 1 MiB

// isSuccessCode reports whether Ecwid accepts status as a successful delivery.
// Everything else, including 203, 208 and every 3xx, is a failed delivery that
// Ecwid will retry.
func isSuccessCode(status int) bool {
	switch status {
	case http.StatusOK, // 200
		http.StatusCreated,   // 201
		http.StatusAccepted,  // 202
		http.StatusNoContent, // 204
		209:                  // no net/http constant; not a registered IANA code
		return true
	default:
		return false
	}
}

// successCodes lists those same codes for error messages and tests. It returns a
// fresh slice per call, so there is no package-level state to mutate.
func successCodes() []int {
	return []int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent, 209}
}

// Deduper records which events have been processed, so a replayed or retried
// webhook runs the callback at most once.
//
// Implementations must be safe for concurrent use, and should record eventID and
// report its prior presence atomically — a check-then-set that is not atomic
// lets two concurrent deliveries of the same event both run the callback.
//
// Entries only need to outlive Ecwid's 24h retry window, but because a signature
// never expires, retaining them longer is what blunts a replay. There is no
// built-in implementation: dedupe state must be shared across every replica of
// your service, so it belongs in whatever store you already run.
//
// Seen is called on receipt, before the callback runs, because marking an event
// only after the callback would let two concurrent deliveries — or a replay
// arriving mid-callback — both through, which is the case this exists to stop.
// The cost is a crash window: if the process dies between Seen and the callback
// completing, Ecwid's retry is deduped and the event is dropped. If losing an
// event is worse for you than handling one twice, have the callback record its
// own completion and let Seen consult that.
type Deduper interface {
	// Seen records eventID and reports whether it was already present.
	Seen(ctx context.Context, eventID string) (bool, error)
}

// Options configures a [Handler]. The zero value is valid.
type Options struct {
	// SuccessCode is the status returned for an accepted webhook. It must be one
	// Ecwid accepts: 200, 201, 202, 204 or 209. Defaults to 200.
	SuccessCode int

	// MaxAge rejects events whose EventCreated is older than this, which bounds
	// how long a captured webhook stays replayable. Ecwid retries a failed
	// delivery for up to 24h, so a window shorter than that can drop legitimate
	// retries of an event your endpoint was down for; weigh that against the
	// replay exposure. Zero disables the check, accepting events of any age.
	MaxAge time.Duration

	// Deduper suppresses events whose EventID has already been processed.
	// Optional; nil disables deduplication, and the callback then runs once per
	// delivery rather than once per event.
	Deduper Deduper

	// OnError is called when a webhook is rejected or dropped, for logging and
	// metrics. It must not log the signature or the client secret. Optional.
	//
	// It reports problems the response cannot: a rejected request is answered
	// long before the callback would run, and a Deduper failure is invisible to
	// the caller entirely.
	OnError func(r *http.Request, err error)
}

// Handler verifies Ecwid webhooks and dispatches them to a callback.
//
// It responds to Ecwid before invoking the callback, so a slow callback cannot
// trip Ecwid's 10s response timeout and cause a redelivery. The callback still
// runs on the request's goroutine — it occupies a server connection until it
// returns, so hand off genuinely long work to your own queue.
//
// Requests are answered with the configured success code once the signature
// verifies; with 401 if it does not, 400 for an unparseable body, 405 for a
// non-POST, and 413 for an oversized one.
type Handler struct {
	clientSecret string
	handle       func(context.Context, Event)
	successCode  int
	maxAge       time.Duration
	deduper      Deduper
	onError      func(*http.Request, error)
	now          func() time.Time // overridden in tests
}

// NewHandler builds a [Handler] that verifies webhooks with clientSecret — the
// app's client_secret, not an access token — and passes those that verify to
// handle. opts may be nil.
//
// The callback must be idempotent: Ecwid retries a failed delivery up to 27
// times over 24h, so the same event can arrive more than once. It must also be
// safe for concurrent use, as deliveries are not serialized.
//
// Treat [Event.Data] as a hint about what changed, not as fact — it is not
// covered by the signature. Re-fetch the entity by [Event.EntityID] over the
// REST API before acting on it. Order events are the exception: their EntityID
// is the internal order ID, so re-fetch those by [OrderData.OrderID] instead.
func NewHandler(clientSecret string, handle func(ctx context.Context, e Event), opts *Options) (*Handler, error) {
	if clientSecret == "" {
		return nil, errors.New("webhooks: clientSecret must not be empty")
	}
	if handle == nil {
		return nil, errors.New("webhooks: handle must not be nil")
	}
	if opts == nil {
		opts = &Options{}
	}

	code := opts.SuccessCode
	if code == 0 {
		code = http.StatusOK
	}
	if !isSuccessCode(code) {
		return nil, fmt.Errorf("webhooks: SuccessCode %d is not one Ecwid accepts (want one of %v)", code, successCodes())
	}
	if opts.MaxAge < 0 {
		return nil, fmt.Errorf("webhooks: MaxAge must not be negative, got %s", opts.MaxAge)
	}

	return &Handler{
		clientSecret: clientSecret,
		handle:       handle,
		successCode:  code,
		maxAge:       opts.MaxAge,
		deduper:      opts.Deduper,
		onError:      opts.OnError,
		now:          time.Now,
	}, nil
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		h.reject(w, r, http.StatusMethodNotAllowed, fmt.Errorf("webhooks: method %s not allowed", r.Method))
		return
	}

	body, err := io.ReadAll(http.MaxBytesReader(w, r.Body, maxBodyBytes))
	if err != nil {
		if _, ok := errors.AsType[*http.MaxBytesError](err); ok {
			h.reject(w, r, http.StatusRequestEntityTooLarge, fmt.Errorf("webhooks: body exceeds %d bytes", maxBodyBytes))
			return
		}
		h.reject(w, r, http.StatusBadRequest, fmt.Errorf("webhooks: read body: %w", err))
		return
	}

	var e Event
	if err := json.Unmarshal(body, &e); err != nil {
		h.reject(w, r, http.StatusBadRequest, fmt.Errorf("webhooks: parse body: %w", err))
		return
	}

	// Verify before anything else touches the event: until this passes, every
	// field above is attacker-controlled.
	if err := Verify(r.Header.Get(SignatureHeader), e.EventCreated, e.EventID, h.clientSecret); err != nil {
		status := http.StatusUnauthorized
		if !errors.Is(err, ErrInvalidSignature) {
			// A misconfigured secret is our fault, not the sender's.
			status = http.StatusInternalServerError
		}
		h.reject(w, r, status, err)
		return
	}

	// Answer with a success code even for events we drop below: they are
	// authentic and a failure response would only earn 27 retries of an event we
	// have already decided not to process.
	if h.maxAge > 0 {
		if age := h.now().Sub(time.Unix(e.EventCreated, 0)); age > h.maxAge {
			w.WriteHeader(h.successCode)
			h.report(r, fmt.Errorf("webhooks: dropped event %s: age %s exceeds MaxAge %s", e.EventID, age.Round(time.Second), h.maxAge))
			return
		}
	}

	w.WriteHeader(h.successCode)
	// Flush so Ecwid has its response before the callback starts; without this
	// the response could sit buffered until ServeHTTP returns.
	if err := http.NewResponseController(w).Flush(); err != nil {
		// Unsupported by this ResponseWriter, or the peer is gone. Neither is
		// worth dropping an authentic event over — the response is already
		// written, and the callback below does not depend on it.
		h.report(r, fmt.Errorf("webhooks: flush response: %w", err))
	}

	// Detach from the request context: it is canceled when ServeHTTP returns,
	// and Ecwid may close the connection as soon as it reads the response above.
	// Neither should abort a callback that has already been promised success.
	ctx := context.WithoutCancel(r.Context())

	if h.deduper != nil {
		seen, err := h.deduper.Seen(ctx, e.EventID)
		if err != nil {
			// Fail open. Delivery is at-least-once regardless, so a callback that
			// is not idempotent is already broken; dropping the event instead
			// would turn a flaky dedupe store into silent data loss.
			h.report(r, fmt.Errorf("webhooks: dedupe event %s: %w", e.EventID, err))
		} else if seen {
			return
		}
	}

	h.handle(ctx, e)
}

// reject writes an error status and reports why.
func (h *Handler) reject(w http.ResponseWriter, r *http.Request, status int, err error) {
	// The body is deliberately generic: the sender is unauthenticated, so err
	// goes to OnError, not over the wire.
	http.Error(w, http.StatusText(status), status)
	h.report(r, err)
}

func (h *Handler) report(r *http.Request, err error) {
	if h.onError != nil {
		h.onError(r, err)
	}
}
