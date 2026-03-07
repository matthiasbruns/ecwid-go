package ecwid

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/matthiasbruns/ecwid-go/config"
)

func testClient(t *testing.T, handler http.HandlerFunc) *Client {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	cfg := config.Config{
		StoreID: "12345",
		Token:   "test-token",
		BaseURL: srv.URL,
	}

	return NewClient(cfg)
}

func loadFixture(t *testing.T, name string) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", name))
	if err != nil {
		t.Fatalf("load fixture %s: %v", name, err)
	}
	return data
}

func TestClient_AuthorizationHeader(t *testing.T) {
	var gotAuth string
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	})

	var result map[string]any
	err := c.get(context.Background(), "/test", nil, &result)
	if err != nil {
		t.Fatal(err)
	}

	if gotAuth != "Bearer test-token" {
		t.Errorf("Authorization = %q, want %q", gotAuth, "Bearer test-token")
	}
}

func TestClient_AcceptGzip(t *testing.T) {
	var gotEncoding string
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotEncoding = r.Header.Get("Accept-Encoding")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	})

	var result map[string]any
	err := c.get(context.Background(), "/test", nil, &result)
	if err != nil {
		t.Fatal(err)
	}

	if gotEncoding != "gzip" {
		t.Errorf("Accept-Encoding = %q, want %q", gotEncoding, "gzip")
	}
}

func TestClient_JSONContentType(t *testing.T) {
	var gotContentType string
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotContentType = r.Header.Get("Content-Type")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	})

	err := c.post(context.Background(), "/test", map[string]string{"key": "value"}, nil)
	if err != nil {
		t.Fatal(err)
	}

	if gotContentType != "application/json" {
		t.Errorf("Content-Type = %q, want %q", gotContentType, "application/json")
	}
}

func TestClient_RequestBody(t *testing.T) {
	var gotBody map[string]string
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		data, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(data, &gotBody)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	})

	body := map[string]string{"name": "Test Product"}
	err := c.post(context.Background(), "/products", body, nil)
	if err != nil {
		t.Fatal(err)
	}

	if gotBody["name"] != "Test Product" {
		t.Errorf("body name = %q, want %q", gotBody["name"], "Test Product")
	}
}

func TestClient_ErrorResponse_404(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write(loadFixture(t, "error_404.json"))
	})

	var result map[string]any
	err := c.get(context.Background(), "/products/999999", nil, &result)

	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected APIError, got %T: %v", err, err)
	}

	if apiErr.StatusCode != 404 {
		t.Errorf("StatusCode = %d, want 404", apiErr.StatusCode)
	}
	if apiErr.Code != "RESOURCE_NOT_FOUND" {
		t.Errorf("Code = %q, want %q", apiErr.Code, "RESOURCE_NOT_FOUND")
	}
}

func TestClient_ErrorResponse_429(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Retry-After", "30")
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write(loadFixture(t, "error_429.json"))
	})

	var result map[string]any
	err := c.get(context.Background(), "/products", nil, &result)

	var rlErr *RateLimitError
	if !errors.As(err, &rlErr) {
		t.Fatalf("expected RateLimitError, got %T: %v", err, err)
	}

	if rlErr.RetryAfter != 30*time.Second {
		t.Errorf("RetryAfter = %v, want 30s", rlErr.RetryAfter)
	}

	// Should also unwrap to APIError.
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Error("RateLimitError should unwrap to APIError")
	}
}

func TestClient_ErrorResponse_500(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(loadFixture(t, "error_500.json"))
	})

	var result map[string]any
	err := c.get(context.Background(), "/products", nil, &result)

	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected APIError, got %T: %v", err, err)
	}

	if apiErr.StatusCode != 500 {
		t.Errorf("StatusCode = %d, want 500", apiErr.StatusCode)
	}
}

func TestClient_ErrorResponse_EmptyBody(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
	})

	var result map[string]any
	err := c.get(context.Background(), "/products", nil, &result)

	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected APIError, got %T: %v", err, err)
	}

	if apiErr.StatusCode != 502 {
		t.Errorf("StatusCode = %d, want 502", apiErr.StatusCode)
	}
}

func TestClient_SuccessfulDecode(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"total":42,"count":10}`))
	})

	var result struct {
		Total int `json:"total"`
		Count int `json:"count"`
	}

	err := c.get(context.Background(), "/products", nil, &result)
	if err != nil {
		t.Fatal(err)
	}

	if result.Total != 42 {
		t.Errorf("Total = %d, want 42", result.Total)
	}
	if result.Count != 10 {
		t.Errorf("Count = %d, want 10", result.Count)
	}
}

func TestClient_BaseURL(t *testing.T) {
	var gotPath string
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	})

	var result map[string]any
	err := c.get(context.Background(), "/products", nil, &result)
	if err != nil {
		t.Fatal(err)
	}

	// baseURL is srv.URL + "/" + storeID, so path should be /12345/products
	if gotPath != "/12345/products" {
		t.Errorf("path = %q, want %q", gotPath, "/12345/products")
	}
}

func TestClient_ContextCancellation(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately.

	var result map[string]any
	err := c.get(ctx, "/products", nil, &result)
	if err == nil {
		t.Fatal("expected error from cancelled context")
	}
}

func TestClient_NilResponseBody(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	})

	// v=nil should discard body without error.
	err := c.get(context.Background(), "/products", nil, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewClient_Services(t *testing.T) {
	c := NewClient(config.Config{StoreID: "1", Token: "t"})

	if c.Products == nil {
		t.Error("Products service is nil")
	}
	if c.Orders == nil {
		t.Error("Orders service is nil")
	}
	if c.Categories == nil {
		t.Error("Categories service is nil")
	}
	if c.Customers == nil {
		t.Error("Customers service is nil")
	}
	if c.Carts == nil {
		t.Error("Carts service is nil")
	}
	if c.Subscriptions == nil {
		t.Error("Subscriptions service is nil")
	}
	if c.Promotions == nil {
		t.Error("Promotions service is nil")
	}
	if c.Coupons == nil {
		t.Error("Coupons service is nil")
	}
	if c.Profile == nil {
		t.Error("Profile service is nil")
	}
	if c.Reviews == nil {
		t.Error("Reviews service is nil")
	}
	if c.Staff == nil {
		t.Error("Staff service is nil")
	}
	if c.Domains == nil {
		t.Error("Domains service is nil")
	}
	if c.Dictionaries == nil {
		t.Error("Dictionaries service is nil")
	}
	if c.Reports == nil {
		t.Error("Reports service is nil")
	}
}

func TestClient_PutRequest(t *testing.T) {
	var gotMethod string
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	})

	err := c.put(context.Background(), "/products/123", map[string]string{"name": "Updated"}, nil)
	if err != nil {
		t.Fatal(err)
	}

	if gotMethod != http.MethodPut {
		t.Errorf("Method = %q, want PUT", gotMethod)
	}
}

func TestClient_DeleteRequest(t *testing.T) {
	var gotMethod string
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	})

	err := c.delete(context.Background(), "/products/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	if gotMethod != http.MethodDelete {
		t.Errorf("Method = %q, want DELETE", gotMethod)
	}
}

func TestClient_Retries429ThenSucceeds(t *testing.T) {
	var attempts int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts == 1 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write(loadFixture(t, "error_429.json"))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	t.Cleanup(srv.Close)

	c := NewClient(config.Config{
		StoreID:    "12345",
		Token:      "test-token",
		BaseURL:    srv.URL,
		MaxRetries: 1,
	})

	var result map[string]any
	if err := c.get(context.Background(), "/products", nil, &result); err != nil {
		t.Fatal(err)
	}
	if attempts != 2 {
		t.Fatalf("attempts = %d, want 2", attempts)
	}
}

func TestClient_RetriesExhausted(t *testing.T) {
	var attempts int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		w.Header().Set("Retry-After", "0")
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write(loadFixture(t, "error_429.json"))
	}))
	t.Cleanup(srv.Close)

	c := NewClient(config.Config{
		StoreID:    "12345",
		Token:      "test-token",
		BaseURL:    srv.URL,
		MaxRetries: 2,
	})

	var result map[string]any
	err := c.get(context.Background(), "/products", nil, &result)

	var rlErr *RateLimitError
	if !errors.As(err, &rlErr) {
		t.Fatalf("expected RateLimitError after retries exhausted, got %T: %v", err, err)
	}
	if attempts != 3 { // 1 initial + 2 retries
		t.Fatalf("attempts = %d, want 3", attempts)
	}
}

func TestClient_WithNilOptions(t *testing.T) {
	// Should not panic with nil options.
	c := NewClient(config.Config{StoreID: "1", Token: "t"},
		WithHTTPClient(nil),
		WithLogger(nil),
	)
	if c.httpClient == nil {
		t.Error("httpClient should not be nil")
	}
	if c.logger == nil {
		t.Error("logger should not be nil")
	}
}
