package webhooks

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
)

// recorder captures the events a handler dispatches.
type recorder struct {
	mu     sync.Mutex
	events []Event
}

func (r *recorder) handle(_ context.Context, e Event) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.events = append(r.events, e)
}

func (r *recorder) get() []Event {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]Event(nil), r.events...)
}

// mapDeduper is an in-memory [Deduper] for tests only. Real deployments need one
// backed by a store shared across replicas.
type mapDeduper struct {
	mu   sync.Mutex
	seen map[string]bool
	err  error
}

func (d *mapDeduper) Seen(_ context.Context, eventID string) (bool, error) {
	if d.err != nil {
		return false, d.err
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.seen == nil {
		d.seen = map[string]bool{}
	}
	was := d.seen[eventID]
	d.seen[eventID] = true
	return was, nil
}

const orderCreatedBody = `{"eventId":"80aece08-40e8-4145-8764-6c2f0d38678","eventCreated":1234567,"storeId":1003,"entityId":103878161,"eventType":"order.created","data":{"orderId":"XJ12H","newPaymentStatus":"PAID","newFulfillmentStatus":"SHIPPED"}}`

// newRequest builds a webhook request, signing it correctly unless signature is
// non-nil, in which case that literal value is used.
func newRequest(body string, signature *string) *http.Request {
	r := httptest.NewRequest(http.MethodPost, "/webhooks", strings.NewReader(body))
	if signature != nil {
		if *signature != "" {
			r.Header.Set(SignatureHeader, *signature)
		}
		return r
	}
	var e Event
	if err := json.Unmarshal([]byte(body), &e); err == nil {
		r.Header.Set(SignatureHeader, sign(e.EventCreated, e.EventID, testSecret))
	}
	return r
}

// newTestHandler builds a handler with a fixed clock, so MaxAge tests do not
// depend on the wall clock.
func newTestHandler(t *testing.T, handle func(context.Context, Event), opts *Options) *Handler {
	t.Helper()
	h, err := NewHandler(testSecret, handle, opts)
	if err != nil {
		t.Fatalf("NewHandler() = %v", err)
	}
	h.now = func() time.Time { return time.Unix(testEventCreated, 0) }
	return h
}

func TestHandler_ValidEvent(t *testing.T) {
	var rec recorder
	h := newTestHandler(t, rec.handle, nil)

	w := httptest.NewRecorder()
	h.ServeHTTP(w, newRequest(orderCreatedBody, nil))

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", w.Code, http.StatusOK)
	}
	got := rec.get()
	if len(got) != 1 {
		t.Fatalf("callback fired %d times, want 1", len(got))
	}
	if got[0].EventType != EventOrderCreated {
		t.Errorf("EventType = %q, want %q", got[0].EventType, EventOrderCreated)
	}
	if got[0].EntityID != "103878161" {
		t.Errorf("EntityID = %q, want %q", got[0].EntityID, "103878161")
	}
	d, err := got[0].OrderData()
	if err != nil {
		t.Fatalf("OrderData() = %v", err)
	}
	if d.OrderID != "XJ12H" {
		t.Errorf("OrderData().OrderID = %q, want %q", d.OrderID, "XJ12H")
	}
}

func TestHandler_RejectsBadSignature(t *testing.T) {
	bad := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
	empty := ""

	tests := []struct {
		name       string
		body       string
		signature  *string
		wantStatus int
	}{
		{"wrong signature", orderCreatedBody, &bad, http.StatusUnauthorized},
		{"missing signature header fails closed", orderCreatedBody, &empty, http.StatusUnauthorized},
		{
			// The signature covers only eventCreated and eventId, so a body whose
			// other fields were swapped still verifies. Confirm that a body whose
			// *signed* fields were altered does not.
			name:       "tampered eventId",
			body:       strings.Replace(orderCreatedBody, `"eventId":"80aece08-40e8-4145-8764-6c2f0d38678"`, `"eventId":"00000000-0000-0000-0000-000000000000"`, 1),
			signature:  new(sign(testEventCreated, testEventID, testSecret)),
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "signature from a different secret",
			body:       orderCreatedBody,
			signature:  new(sign(testEventCreated, testEventID, "some_other_secret")),
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rec recorder
			h := newTestHandler(t, rec.handle, nil)

			w := httptest.NewRecorder()
			h.ServeHTTP(w, newRequest(tt.body, tt.signature))

			if w.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", w.Code, tt.wantStatus)
			}
			if n := len(rec.get()); n != 0 {
				t.Errorf("callback fired %d times, want 0", n)
			}
		})
	}
}

func TestHandler_RejectsMalformedRequests(t *testing.T) {
	tests := []struct {
		name       string
		newReq     func() *http.Request
		wantStatus int
	}{
		{
			name:       "non-POST",
			newReq:     func() *http.Request { return httptest.NewRequest(http.MethodGet, "/webhooks", http.NoBody) },
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:       "invalid JSON",
			newReq:     func() *http.Request { return newRequest(`{not json`, new(testSignature)) },
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "entityId of an unexpected JSON type",
			newReq:     func() *http.Request { return newRequest(`{"eventId":"e","entityId":true}`, new(testSignature)) },
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "oversized body",
			newReq: func() *http.Request {
				return newRequest(`{"eventId":"`+strings.Repeat("A", maxBodyBytes)+`"}`, new(testSignature))
			},
			wantStatus: http.StatusRequestEntityTooLarge,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rec recorder
			h := newTestHandler(t, rec.handle, nil)

			w := httptest.NewRecorder()
			h.ServeHTTP(w, tt.newReq())

			if w.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", w.Code, tt.wantStatus)
			}
			if n := len(rec.get()); n != 0 {
				t.Errorf("callback fired %d times, want 0", n)
			}
		})
	}
}

func TestHandler_MethodNotAllowedAdvertisesPOST(t *testing.T) {
	h := newTestHandler(t, func(context.Context, Event) {}, nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/webhooks", http.NoBody))

	if got := w.Header().Get("Allow"); got != http.MethodPost {
		t.Errorf("Allow = %q, want %q", got, http.MethodPost)
	}
}

func TestHandler_MaxAge(t *testing.T) {
	tests := []struct {
		name         string
		maxAge       time.Duration
		eventCreated int64
		wantCalls    int
	}{
		{"fresh event within window", time.Hour, testEventCreated - 60, 1},
		{"event exactly at the window edge", time.Hour, testEventCreated - int64(time.Hour.Seconds()), 1},
		{"stale event outside window", time.Hour, testEventCreated - 7200, 0},
		{"zero MaxAge accepts an ancient event", 0, 1, 1},
		// Clock skew: an event stamped slightly in the future has a negative age,
		// which must not read as stale.
		{"event from the near future", time.Hour, testEventCreated + 60, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rec recorder
			h := newTestHandler(t, rec.handle, &Options{MaxAge: tt.maxAge})

			body := fmt.Sprintf(`{"eventId":%q,"eventCreated":%d,"storeId":1003,"entityId":1,"eventType":"product.created"}`, testEventID, tt.eventCreated)
			w := httptest.NewRecorder()
			h.ServeHTTP(w, newRequest(body, nil))

			// A stale event is authentic, so it is still answered with success:
			// failing it would only earn 27 retries of an event we will not process.
			if w.Code != http.StatusOK {
				t.Errorf("status = %d, want %d", w.Code, http.StatusOK)
			}
			if n := len(rec.get()); n != tt.wantCalls {
				t.Errorf("callback fired %d times, want %d", n, tt.wantCalls)
			}
		})
	}
}

func TestHandler_Deduper(t *testing.T) {
	var rec recorder
	dd := &mapDeduper{}
	h := newTestHandler(t, rec.handle, &Options{Deduper: dd})

	// The same signed event replayed: valid every time, since the signature never
	// expires. Only the first delivery should reach the callback.
	for range 3 {
		w := httptest.NewRecorder()
		h.ServeHTTP(w, newRequest(orderCreatedBody, nil))
		if w.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", w.Code, http.StatusOK)
		}
	}

	if n := len(rec.get()); n != 1 {
		t.Errorf("callback fired %d times for a replayed event, want 1", n)
	}
}

// A broken dedupe store must not swallow events: delivery is at-least-once
// anyway, so failing open costs a duplicate while failing closed loses data.
func TestHandler_DeduperErrorFailsOpen(t *testing.T) {
	var rec recorder
	var reported []error
	dd := &mapDeduper{err: errors.New("store unavailable")}
	h := newTestHandler(t, rec.handle, &Options{
		Deduper: dd,
		OnError: func(_ *http.Request, err error) { reported = append(reported, err) },
	})

	w := httptest.NewRecorder()
	h.ServeHTTP(w, newRequest(orderCreatedBody, nil))

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", w.Code, http.StatusOK)
	}
	if n := len(rec.get()); n != 1 {
		t.Errorf("callback fired %d times, want 1", n)
	}
	if len(reported) != 1 {
		t.Fatalf("OnError called %d times, want 1", len(reported))
	}
	if !strings.Contains(reported[0].Error(), "store unavailable") {
		t.Errorf("OnError got %v, want it to wrap the store error", reported[0])
	}
}

func TestHandler_OnErrorNeverLeaksSecrets(t *testing.T) {
	var reported []error
	h := newTestHandler(t, func(context.Context, Event) {}, &Options{
		OnError: func(_ *http.Request, err error) { reported = append(reported, err) },
	})

	w := httptest.NewRecorder()
	h.ServeHTTP(w, newRequest(orderCreatedBody, new("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=")))

	if len(reported) != 1 {
		t.Fatalf("OnError called %d times, want 1", len(reported))
	}
	if msg := reported[0].Error(); strings.Contains(msg, testSecret) {
		t.Errorf("OnError message contains the client secret: %s", msg)
	}
	// The response body must not describe why verification failed.
	if body := w.Body.String(); strings.Contains(body, "signature") {
		t.Errorf("response body leaks failure detail: %s", body)
	}
}

func TestHandler_SuccessCode(t *testing.T) {
	for _, code := range successCodes() {
		t.Run(fmt.Sprint(code), func(t *testing.T) {
			var rec recorder
			h := newTestHandler(t, rec.handle, &Options{SuccessCode: code})

			w := httptest.NewRecorder()
			h.ServeHTTP(w, newRequest(orderCreatedBody, nil))

			if w.Code != code {
				t.Errorf("status = %d, want %d", w.Code, code)
			}
			if n := len(rec.get()); n != 1 {
				t.Errorf("callback fired %d times, want 1", n)
			}
		})
	}
}

// Ecwid counts 203, 208 and every 3xx as a failed delivery despite 203/208 being
// 2xx. Configuring one is a silent outage, so it must not be accepted.
func TestNewHandler_RejectsCodesEcwidTreatsAsFailure(t *testing.T) {
	for _, code := range []int{203, 208, 301, 302, 400, 500, 999, -1} {
		if _, err := NewHandler(testSecret, func(context.Context, Event) {}, &Options{SuccessCode: code}); err == nil {
			t.Errorf("NewHandler(SuccessCode: %d) = nil, want an error", code)
		}
	}
}

func TestNewHandler_Validation(t *testing.T) {
	tests := []struct {
		name         string
		clientSecret string
		handle       func(context.Context, Event)
		opts         *Options
	}{
		{"empty client secret", "", func(context.Context, Event) {}, nil},
		{"nil callback", testSecret, nil, nil},
		{"negative MaxAge", testSecret, func(context.Context, Event) {}, &Options{MaxAge: -time.Second}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewHandler(tt.clientSecret, tt.handle, tt.opts); err == nil {
				t.Error("NewHandler() = nil, want an error")
			}
		})
	}
}

func TestNewHandler_NilOptionsDefaultsTo200(t *testing.T) {
	h, err := NewHandler(testSecret, func(context.Context, Event) {}, nil)
	if err != nil {
		t.Fatalf("NewHandler() = %v", err)
	}
	if h.successCode != http.StatusOK {
		t.Errorf("successCode = %d, want %d", h.successCode, http.StatusOK)
	}
}

// The response must reach Ecwid before the callback runs, or a slow callback
// trips Ecwid's 10s response timeout and earns a redelivery.
func TestHandler_RespondsBeforeCallback(t *testing.T) {
	responded := make(chan struct{})
	release := make(chan struct{})

	h, err := NewHandler(testSecret, func(context.Context, Event) {
		// Blocks until the client has read the response. If the handler waited
		// for this callback before responding, the test would deadlock.
		<-responded
		close(release)
	}, nil)
	if err != nil {
		t.Fatalf("NewHandler() = %v", err)
	}

	srv := httptest.NewServer(h)
	defer srv.Close()

	req, err := http.NewRequestWithContext(t.Context(), http.MethodPost, srv.URL, strings.NewReader(orderCreatedBody))
	if err != nil {
		t.Fatalf("NewRequest() = %v", err)
	}
	req.Header.Set(SignatureHeader, sign(testEventCreated, testEventID, testSecret))

	// Do runs asynchronously: should the handler ever wait for the callback
	// before responding, it and the callback would wait on each other forever,
	// and a synchronous Do would hang the test rather than fail it.
	type result struct {
		status int
		err    error
	}
	done := make(chan result, 1)
	go func() {
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			done <- result{err: err}
			return
		}
		defer func() { _ = resp.Body.Close() }()
		done <- result{status: resp.StatusCode}
	}()

	select {
	case got := <-done:
		if got.err != nil {
			t.Fatalf("Do() = %v", got.err)
		}
		if got.status != http.StatusOK {
			t.Errorf("status = %d, want %d", got.status, http.StatusOK)
		}
	case <-time.After(5 * time.Second):
		close(responded) // Unblock the callback so the server can shut down.
		t.Fatal("handler did not respond until after the callback ran")
	}

	close(responded)

	select {
	case <-release:
	case <-time.After(5 * time.Second):
		t.Fatal("callback did not run after the response was sent")
	}
}

// The callback outlives the request context, which is canceled once ServeHTTP
// returns and may be cut short by Ecwid closing the connection.
func TestHandler_CallbackContextNotCanceled(t *testing.T) {
	var (
		mu     sync.Mutex
		ctxErr error
	)
	done := make(chan struct{})

	h, err := NewHandler(testSecret, func(ctx context.Context, _ Event) {
		mu.Lock()
		ctxErr = ctx.Err()
		mu.Unlock()
		close(done)
	}, nil)
	if err != nil {
		t.Fatalf("NewHandler() = %v", err)
	}

	srv := httptest.NewServer(h)
	defer srv.Close()

	req, err := http.NewRequestWithContext(t.Context(), http.MethodPost, srv.URL, strings.NewReader(orderCreatedBody))
	if err != nil {
		t.Fatalf("NewRequest() = %v", err)
	}
	req.Header.Set(SignatureHeader, sign(testEventCreated, testEventID, testSecret))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Do() = %v", err)
	}
	_ = resp.Body.Close()

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("callback never ran")
	}

	mu.Lock()
	defer mu.Unlock()
	if ctxErr != nil {
		t.Errorf("callback context was canceled: %v", ctxErr)
	}
}

func TestHandler_ConcurrentDeliveries(t *testing.T) {
	var rec recorder
	dd := &mapDeduper{}
	h := newTestHandler(t, rec.handle, &Options{Deduper: dd})

	const n = 20
	var wg sync.WaitGroup
	for i := range n {
		wg.Go(func() {
			eventID := fmt.Sprintf("event-%d", i)
			body := fmt.Sprintf(`{"eventId":%q,"eventCreated":%d,"storeId":1003,"entityId":%d,"eventType":"product.created"}`, eventID, testEventCreated, i)
			w := httptest.NewRecorder()
			h.ServeHTTP(w, newRequest(body, nil))
			if w.Code != http.StatusOK {
				t.Errorf("status = %d, want %d", w.Code, http.StatusOK)
			}
		})
	}
	wg.Wait()

	if got := len(rec.get()); got != n {
		t.Errorf("callback fired %d times, want %d", got, n)
	}
}
