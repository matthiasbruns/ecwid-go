package carts

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
	cfg := config.Config{
		StoreID: "12345",
		Token:   "test-token",
		BaseURL: srv.URL,
	}
	return ecwid.NewClient(cfg, ecwid.WithHTTPClient(srv.Client()))
}

func executeCmd(t *testing.T, client *ecwid.Client, args []string) (string, error) {
	t.Helper()

	root := &cobra.Command{Use: "ecwid"}
	root.PersistentFlags().String("output", "", "output format")

	root.AddCommand(Cmd)
	cmdutil.AppClient = client

	var buf bytes.Buffer
	root.SetOut(&buf)

	root.SetArgs(args)
	err := root.ExecuteContext(context.Background())
	return buf.String(), err
}

func TestListCmd_JSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	client := newTestClient(t, srv)
	out, err := executeCmd(t, client, []string{"carts", "list"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "abc123") {
		t.Errorf("expected cartId in output, got:\n%s", out)
	}
}

func TestGetCmd_MissingArg(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(srv.Close)

	client := newTestClient(t, srv)
	_, err := executeCmd(t, client, []string{"carts", "get"})
	if err == nil {
		t.Fatal("expected error when cartId arg is missing")
	}
}

func TestUpdateCmd_MissingArg(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(srv.Close)

	client := newTestClient(t, srv)
	_, err := executeCmd(t, client, []string{"carts", "update"})
	if err == nil {
		t.Fatal("expected error when cartId arg is missing")
	}
}

func TestPlaceCmd_MissingArg(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(srv.Close)

	client := newTestClient(t, srv)
	_, err := executeCmd(t, client, []string{"carts", "place"})
	if err == nil {
		t.Fatal("expected error when cartId arg is missing")
	}
}
