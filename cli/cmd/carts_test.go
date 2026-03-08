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

// newTestClientFromServer creates an ecwid.Client pointed at the given test server.
func newTestClientFromServer(t *testing.T, srv *httptest.Server) *ecwid.Client {
	t.Helper()
	cfg := config.Config{
		StoreID: "12345",
		Token:   "test-token",
		BaseURL: srv.URL,
	}
	return ecwid.NewClient(cfg, ecwid.WithHTTPClient(srv.Client()))
}

// executeCartsCmd runs cartsCmd as a sub-command of a fresh root command,
// without triggering PersistentPreRunE (which needs real credentials).
func executeCartsCmd(t *testing.T, client *ecwid.Client, args []string) (string, error) {
	t.Helper()

	root := &cobra.Command{Use: "ecwid"}
	root.PersistentFlags().String("output", "", "output format")

	// Wire the cart sub-commands under our test root.
	root.AddCommand(cartsCmd)
	AppClient = client

	var buf bytes.Buffer
	root.SetOut(&buf)

	root.SetArgs(args)
	err := root.ExecuteContext(context.Background())
	return buf.String(), err
}

func TestCartsListCmd_JSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/12345/carts" {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		resp := map[string]any{
			"total": 1, "count": 1, "offset": 0, "limit": 100,
			"items": []map[string]any{
				{"cartId": "abc123", "total": 42.5, "email": "user@example.com"},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
	t.Cleanup(srv.Close)

	client := newTestClientFromServer(t, srv)
	out, err := executeCartsCmd(t, client, []string{"carts", "list"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "abc123") {
		t.Errorf("expected cartId in output, got:\n%s", out)
	}
}

func TestCartsListCmd_WithFlags(t *testing.T) {
	var gotQuery string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.RawQuery
		resp := map[string]any{"total": 0, "count": 0, "offset": 5, "limit": 10, "items": []any{}}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
	t.Cleanup(srv.Close)

	client := newTestClientFromServer(t, srv)
	_, err := executeCartsCmd(t, client, []string{"carts", "list", "--customer-id=99", "--limit=10", "--offset=5"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(gotQuery, "customerId=99") {
		t.Errorf("expected customerId in query, got: %s", gotQuery)
	}
	if !strings.Contains(gotQuery, "limit=10") {
		t.Errorf("expected limit in query, got: %s", gotQuery)
	}
	if !strings.Contains(gotQuery, "offset=5") {
		t.Errorf("expected offset in query, got: %s", gotQuery)
	}
}

func TestCartsGetCmd_JSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/12345/carts/cart-xyz" {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		resp := map[string]any{"cartId": "cart-xyz", "total": 99.0, "email": "buyer@example.com"}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
	t.Cleanup(srv.Close)

	client := newTestClientFromServer(t, srv)
	out, err := executeCartsCmd(t, client, []string{"carts", "get", "cart-xyz"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "cart-xyz") {
		t.Errorf("expected cartId in output, got:\n%s", out)
	}
}

func TestCartsGetCmd_MissingArg(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(srv.Close)

	client := newTestClientFromServer(t, srv)
	_, err := executeCartsCmd(t, client, []string{"carts", "get"})
	if err == nil {
		t.Fatal("expected error when cartId arg is missing")
	}
}

func TestCartsUpdateCmd_MissingArg(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(srv.Close)

	client := newTestClientFromServer(t, srv)
	_, err := executeCartsCmd(t, client, []string{"carts", "update"})
	if err == nil {
		t.Fatal("expected error when cartId arg is missing")
	}
}

func TestCartsPlaceCmd_MissingArg(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(srv.Close)

	client := newTestClientFromServer(t, srv)
	_, err := executeCartsCmd(t, client, []string{"carts", "place"})
	if err == nil {
		t.Fatal("expected error when cartId arg is missing")
	}
}

func TestCartsUpdateCmd_Hidden(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "bad method", http.StatusMethodNotAllowed)
			return
		}
		resp := map[string]any{"updateCount": 1}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
	t.Cleanup(srv.Close)

	client := newTestClientFromServer(t, srv)
	out, err := executeCartsCmd(t, client, []string{"carts", "update", "cart-abc", "--hidden=true"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "1") {
		t.Errorf("expected updateCount in output, got:\n%s", out)
	}
}

func TestCartsPlaceCmd_JSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "bad method", http.StatusMethodNotAllowed)
			return
		}
		resp := map[string]any{
			"id":                "O-12345",
			"orderNumber":       42,
			"vendorOrderNumber": "V-42",
			"cartId":            "cart-abc",
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
	t.Cleanup(srv.Close)

	client := newTestClientFromServer(t, srv)
	out, err := executeCartsCmd(t, client, []string{"carts", "place", "cart-abc"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "O-12345") {
		t.Errorf("expected order id in output, got:\n%s", out)
	}
}
