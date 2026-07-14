package billing_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid"
	"github.com/matthiasbruns/ecwid-go/ecwid/billing"
	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

func newTestService(t *testing.T, srv *httptest.Server) *billing.Service {
	t.Helper()
	return billing.NewService(api.NewHTTPClient(api.HTTPClientConfig{
		BaseURL: srv.URL + "/api/v3",
		StoreID: "12345",
		Token:   "secret_test",
	}))
}

func TestCharge(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/billing/transactions" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}

		var body billing.ChargeRequest
		data, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		if err := json.Unmarshal(data, &body); err != nil {
			t.Fatalf("decode request body: %v", err)
		}
		if body.IdempotencyKey != "bill-1" {
			t.Errorf("expected idempotencyKey=bill-1, got %q", body.IdempotencyKey)
		}
		if body.Amount != 12.34 {
			t.Errorf("expected amount=12.34, got %v", body.Amount)
		}
		if body.Currency != "EUR" {
			t.Errorf("expected currency=EUR, got %q", body.Currency)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"transactionId":"txn-99","idempotencyKeyInUse":false}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Charge(context.Background(), &billing.ChargeRequest{
		IdempotencyKey: "bill-1",
		Amount:         12.34,
		Currency:       "EUR",
		Description:    "test charge",
		Metadata:       map[string]any{"shopID": "12345"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.TransactionID != "txn-99" {
		t.Errorf("expected transactionId=txn-99, got %q", result.TransactionID)
	}
}

func TestCharge_ErrorStatusSurfacesAPIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusPaymentRequired)
		_, _ = w.Write([]byte(`{"errorMessage":"CHARGE_LIMIT_EXCEEDED"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	_, err := svc.Charge(context.Background(), &billing.ChargeRequest{
		IdempotencyKey: "bill-2",
		Amount:         5,
		Currency:       "EUR",
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	var apiErr *ecwid.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *ecwid.APIError, got %T", err)
	}
	if apiErr.StatusCode != http.StatusPaymentRequired {
		t.Errorf("expected status 402, got %d", apiErr.StatusCode)
	}
	if apiErr.Message != "CHARGE_LIMIT_EXCEEDED" {
		t.Errorf("expected message CHARGE_LIMIT_EXCEEDED, got %q", apiErr.Message)
	}
}

func TestCharge_NilRequest(t *testing.T) {
	svc := billing.NewService(api.NewHTTPClient(api.HTTPClientConfig{
		BaseURL: "https://example.invalid/api/v3",
		StoreID: "12345",
		Token:   "secret_test",
	}))
	if _, err := svc.Charge(context.Background(), nil); err == nil {
		t.Fatal("expected error for nil request, got nil")
	}
}
