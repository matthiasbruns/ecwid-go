package webhook

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/webhooks"
)

// newTestServer wires a Handler over a Trigger delivering to url onto a mux.
func newTestServer(t *testing.T, url string) *httptest.Server {
	t.Helper()
	mux := http.NewServeMux()
	NewHandler(newTestTrigger(t, url)).Routes(mux)
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return srv
}

func TestHTTP_Trigger_NumericEntityIDInRequest(t *testing.T) {
	var gotBody []byte
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotBody = readBody(t, r)
		w.WriteHeader(http.StatusOK)
	}))
	defer backend.Close()

	control := newTestServer(t, backend.URL)

	// entityId as a bare JSON number, mirroring the issue's example.
	body := `{"eventType":"order.updated","entityId":103878161,"data":{"orderId":"XJ12H","newPaymentStatus":"PAID"}}`
	resp, err := http.Post(control.URL+"/_mock/webhooks/trigger", "application/json", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", resp.StatusCode)
	}
	var res Result
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		t.Fatal(err)
	}
	if !res.Delivered {
		t.Errorf("delivered = false, reason %q", res.Reason)
	}
	if res.EntityID != "103878161" {
		t.Errorf("entityId = %q, want 103878161", res.EntityID)
	}

	var e webhooks.Event
	if err := json.Unmarshal(gotBody, &e); err != nil {
		t.Fatal(err)
	}
	if e.EntityID != "103878161" {
		t.Errorf("delivered entityId = %q, want 103878161", e.EntityID)
	}
}

func TestHTTP_Trigger_NoURLReturns409(t *testing.T) {
	control := newTestServer(t, "")
	resp, err := http.Post(control.URL+"/_mock/webhooks/trigger", "application/json",
		strings.NewReader(`{"eventType":"order.created"}`))
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusConflict {
		t.Errorf("status = %d, want 409 when no --webhook-url", resp.StatusCode)
	}
}

func TestHTTP_Trigger_BadRequests(t *testing.T) {
	control := newTestServer(t, "http://127.0.0.1:0")
	cases := map[string]string{
		"missing eventType":  `{"entityId":1}`,
		"unknown event":      `{"eventType":"no.such.event"}`,
		"malformed JSON":     `{`,
		"unknown field":      `{"eventType":"order.created","nope":1}`,
		"bad signature mode": `{"eventType":"order.created","signature":"bogus"}`,
	}
	for name, body := range cases {
		t.Run(name, func(t *testing.T) {
			resp, err := http.Post(control.URL+"/_mock/webhooks/trigger", "application/json", strings.NewReader(body))
			if err != nil {
				t.Fatal(err)
			}
			defer func() { _ = resp.Body.Close() }()
			if resp.StatusCode != http.StatusBadRequest {
				t.Errorf("status = %d, want 400", resp.StatusCode)
			}
		})
	}
}

func TestHTTP_Events(t *testing.T) {
	control := newTestServer(t, "")
	resp, err := http.Get(control.URL + "/_mock/webhooks/events")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = resp.Body.Close() }()

	var payload struct {
		Events []eventInfo `json:"events"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		t.Fatal(err)
	}
	if len(payload.Events) != wantEventCount {
		t.Fatalf("events = %d, want %d", len(payload.Events), wantEventCount)
	}
	for _, ev := range payload.Events {
		if ev.EventType == "" || ev.Group == "" || ev.EntityIDType == "" {
			t.Errorf("event %q has an empty field: %+v", ev.EventType, ev)
		}
	}
}

func TestHTTP_UI(t *testing.T) {
	enabled := newTestServer(t, "http://example.test/hook")
	body := getBody(t, enabled.URL+"/_mock/webhooks/ui")
	if !strings.Contains(body, "Webhook Trigger") {
		t.Error("UI missing its heading")
	}
	if !strings.Contains(body, "http://example.test/hook") {
		t.Error("UI does not show the configured webhook URL")
	}
	if strings.Contains(body, `type="submit" disabled`) {
		t.Error("UI Fire button disabled while a webhook URL is configured")
	}

	disabled := newTestServer(t, "")
	body = getBody(t, disabled.URL+"/_mock/webhooks/ui")
	if !strings.Contains(body, `type="submit" disabled`) {
		t.Error("UI Fire button not disabled without a webhook URL")
	}
}

func TestNormalizeEntityID(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{`103878161`, "103878161"},
		{`"1003"`, "1003"},
		{`  42 `, "42"},
		{``, ""},
		{`null`, ""},
	}
	for _, c := range cases {
		got, err := normalizeEntityID(json.RawMessage(c.in))
		if err != nil {
			t.Errorf("normalizeEntityID(%q) error: %v", c.in, err)
			continue
		}
		if got != c.want {
			t.Errorf("normalizeEntityID(%q) = %q, want %q", c.in, got, c.want)
		}
	}
	if _, err := normalizeEntityID(json.RawMessage(`{}`)); err == nil {
		t.Error("normalizeEntityID of an object should error")
	}
}

func readBody(t *testing.T, r *http.Request) []byte {
	t.Helper()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func getBody(t *testing.T, url string) string {
	t.Helper()
	resp, err := http.Get(url) //nolint:noctx // test helper
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GET %s = %d, want 200", url, resp.StatusCode)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}
