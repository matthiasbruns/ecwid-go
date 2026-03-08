package orders

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/cli/internal/cmdutil"
	"github.com/matthiasbruns/ecwid-go/config"
	"github.com/matthiasbruns/ecwid-go/ecwid"
)

func newTestClient(t *testing.T, srv *httptest.Server) *ecwid.Client {
	t.Helper()
	return ecwid.NewClient(config.Config{
		StoreID: "12345",
		Token:   "secret_test",
		BaseURL: srv.URL + "/api/v3",
	})
}

func newTestCmd(format string) *cobra.Command {
	cmd := &cobra.Command{Use: "test"}
	cmd.Flags().String("output", format, "")
	cmd.SetContext(context.Background())
	return cmd
}

func TestListCmd(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":1,"count":1,"offset":10,"limit":5,"items":[{"id":"ORD-1","email":"user@example.com"}]}`))
	}))
	defer srv.Close()

	cmdutil.AppClient = newTestClient(t, srv)

	cmd := newTestCmd("json")
	cmd.Flags().String("payment-status", "PAID", "")
	cmd.Flags().String("fulfillment-status", "SHIPPED", "")
	cmd.Flags().String("email", "user@example.com", "")
	cmd.Flags().String("keyword", "widget", "")
	cmd.Flags().Int("limit", 5, "")
	cmd.Flags().Int("offset", 10, "")

	var buf bytes.Buffer
	cmd.SetOut(&buf)

	err := listCmd.RunE(cmd, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(buf.String(), "ORD-1") {
		t.Errorf("expected order ID in output, got:\n%s", buf.String())
	}
}

func TestGetCmd(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"ORD-42","email":"buyer@example.com"}`))
	}))
	defer srv.Close()

	cmdutil.AppClient = newTestClient(t, srv)

	cmd := newTestCmd("json")
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	err := getCmd.RunE(cmd, []string{"ORD-42"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(buf.String(), "ORD-42") {
		t.Errorf("expected order ID in output, got:\n%s", buf.String())
	}
}

func TestCreateCmd(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":999,"orderid":"ORD-999"}`))
	}))
	defer srv.Close()

	cmdutil.AppClient = newTestClient(t, srv)

	payload, _ := json.Marshal(map[string]any{"email": "new@example.com", "total": 49.99})

	cmd := newTestCmd("json")
	cmd.Flags().String("data", string(payload), "")
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	err := createCmd.RunE(cmd, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(buf.String(), "ORD-999") {
		t.Errorf("expected created order ID in output, got:\n%s", buf.String())
	}
}

func TestCreateCmd_InvalidJSON(t *testing.T) {
	cmdutil.AppClient = ecwid.NewClient(config.Config{
		StoreID: "12345",
		Token:   "secret_test",
		BaseURL: "http://localhost:1",
	})

	cmd := newTestCmd("json")
	cmd.Flags().String("data", "not-json", "")
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	err := createCmd.RunE(cmd, nil)
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
	if !strings.Contains(err.Error(), "parse JSON") {
		t.Errorf("expected parse JSON error, got: %v", err)
	}
}

func TestDeleteCmd(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	cmdutil.AppClient = newTestClient(t, srv)

	cmd := newTestCmd("json")
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	err := deleteCmd.RunE(cmd, []string{"ORD-5"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(buf.String(), "deleteCount") {
		t.Errorf("expected deleteCount in output, got:\n%s", buf.String())
	}
}
