package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/config"
	"github.com/matthiasbruns/ecwid-go/ecwid"
)

// newOrdersTestClient creates an ecwid.Client pointing at the given test server.
// The server receives requests at /<storeID>/... relative to srv.URL.
func newOrdersTestClient(t *testing.T, srv *httptest.Server) *ecwid.Client {
	t.Helper()
	return ecwid.NewClient(config.Config{
		StoreID: "12345",
		Token:   "secret_test",
		BaseURL: srv.URL + "/api/v3",
	})
}

// newTestCmdWithContext creates a cobra test command with output format and a background context.
func newTestCmdWithContext(format string) *cobra.Command {
	cmd := newCmdWithOutput(format)
	cmd.SetContext(context.Background())
	return cmd
}

func TestOrdersListCmd(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/orders" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("paymentStatus"); got != "PAID" {
			t.Errorf("expected paymentStatus=PAID, got %q", got)
		}
		if got := r.URL.Query().Get("fulfillmentStatus"); got != "SHIPPED" {
			t.Errorf("expected fulfillmentStatus=SHIPPED, got %q", got)
		}
		if got := r.URL.Query().Get("email"); got != "user@example.com" {
			t.Errorf("expected email=user@example.com, got %q", got)
		}
		if got := r.URL.Query().Get("keywords"); got != "widget" {
			t.Errorf("expected keywords=widget, got %q", got)
		}
		if got := r.URL.Query().Get("limit"); got != "5" {
			t.Errorf("expected limit=5, got %q", got)
		}
		if got := r.URL.Query().Get("offset"); got != "10" {
			t.Errorf("expected offset=10, got %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":1,"count":1,"offset":10,"limit":5,"items":[{"id":"ORD-1","email":"user@example.com"}]}`))
	}))
	defer srv.Close()

	AppClient = newOrdersTestClient(t, srv)

	cmd := newTestCmdWithContext("json")
	cmd.Flags().String("payment-status", "PAID", "")
	cmd.Flags().String("fulfillment-status", "SHIPPED", "")
	cmd.Flags().String("email", "user@example.com", "")
	cmd.Flags().String("keyword", "widget", "")
	cmd.Flags().Int("limit", 5, "")
	cmd.Flags().Int("offset", 10, "")

	var buf bytes.Buffer
	cmd.SetOut(&buf)

	err := ordersListCmd.RunE(cmd, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "ORD-1") {
		t.Errorf("expected order ID in output, got:\n%s", out)
	}
}

func TestOrdersGetCmd(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/orders/ORD-42" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"ORD-42","email":"buyer@example.com","paymentStatus":"PAID"}`))
	}))
	defer srv.Close()

	AppClient = newOrdersTestClient(t, srv)

	cmd := newTestCmdWithContext("json")
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	err := ordersGetCmd.RunE(cmd, []string{"ORD-42"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "ORD-42") {
		t.Errorf("expected order ID in output, got:\n%s", out)
	}
	if !strings.Contains(out, "buyer@example.com") {
		t.Errorf("expected email in output, got:\n%s", out)
	}
}

func TestOrdersCreateCmd(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/orders" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":999,"orderid":"ORD-999"}`))
	}))
	defer srv.Close()

	AppClient = newOrdersTestClient(t, srv)

	payload, _ := json.Marshal(map[string]any{
		"email":             "new@example.com",
		"total":             49.99,
		"paymentStatus":     "AWAITING_PAYMENT",
		"fulfillmentStatus": "AWAITING_PROCESSING",
	})

	cmd := newTestCmdWithContext("json")
	cmd.Flags().String("data", string(payload), "")
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	err := ordersCreateCmd.RunE(cmd, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "ORD-999") {
		t.Errorf("expected created order ID in output, got:\n%s", out)
	}
}

func TestOrdersCreateCmd_Stdin(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":1,"orderid":"ORD-1"}`))
	}))
	defer srv.Close()

	AppClient = newOrdersTestClient(t, srv)

	payload := `{"email":"stdin@example.com","total":10.0,"paymentStatus":"PAID","fulfillmentStatus":"SHIPPED"}`

	cmd := newTestCmdWithContext("json")
	cmd.Flags().String("data", "", "")
	cmd.SetIn(strings.NewReader(payload))
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	err := ordersCreateCmd.RunE(cmd, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "ORD-1") {
		t.Errorf("expected order ID in output, got:\n%s", out)
	}
}

func TestOrdersUpdateCmd_Stdin(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/orders/ORD-7" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	AppClient = newOrdersTestClient(t, srv)

	payload := `{"fulfillmentStatus":"SHIPPED"}`

	cmd := newTestCmdWithContext("json")
	cmd.Flags().String("data", "", "")
	cmd.SetIn(strings.NewReader(payload))
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	err := ordersUpdateCmd.RunE(cmd, []string{"ORD-7"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "updateCount") {
		t.Errorf("expected updateCount in output, got:\n%s", out)
	}
}

func TestOrdersUpdateCmd(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/orders/ORD-7" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	AppClient = newOrdersTestClient(t, srv)

	payload, _ := json.Marshal(map[string]any{
		"fulfillmentStatus": "SHIPPED",
	})

	cmd := newTestCmdWithContext("json")
	cmd.Flags().String("data", string(payload), "")
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	err := ordersUpdateCmd.RunE(cmd, []string{"ORD-7"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "updateCount") {
		t.Errorf("expected updateCount in output, got:\n%s", out)
	}
}

func TestOrdersDeleteCmd(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/orders/ORD-5" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	AppClient = newOrdersTestClient(t, srv)

	cmd := newTestCmdWithContext("json")
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	err := ordersDeleteCmd.RunE(cmd, []string{"ORD-5"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "deleteCount") {
		t.Errorf("expected deleteCount in output, got:\n%s", out)
	}
}

func TestOrdersCreateCmd_InvalidJSON(t *testing.T) {
	AppClient = ecwid.NewClient(config.Config{
		StoreID: "12345",
		Token:   "secret_test",
		BaseURL: "http://localhost:1",
	})

	cmd := newTestCmdWithContext("json")
	cmd.Flags().String("data", "not-json", "")
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	err := ordersCreateCmd.RunE(cmd, nil)
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
	if !strings.Contains(err.Error(), "parse JSON") {
		t.Errorf("expected parse JSON error, got: %v", err)
	}
}

func TestOrdersUpdateCmd_InvalidJSON(t *testing.T) {
	AppClient = ecwid.NewClient(config.Config{
		StoreID: "12345",
		Token:   "secret_test",
		BaseURL: "http://localhost:1",
	})

	cmd := newTestCmdWithContext("json")
	cmd.Flags().String("data", "{bad json", "")
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	err := ordersUpdateCmd.RunE(cmd, []string{"ORD-1"})
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
	if !strings.Contains(err.Error(), "parse JSON") {
		t.Errorf("expected parse JSON error, got: %v", err)
	}
}
