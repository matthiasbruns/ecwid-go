package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPClient_Get(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.Header.Get("Authorization") != "Bearer test_token" {
			t.Error("missing or wrong Authorization header")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()

	c := NewHTTPClient(HTTPClientConfig{
		BaseURL: srv.URL + "/api/v3",
		StoreID: "123",
		Token:   "test_token",
	})

	var result struct{ OK bool }
	if err := c.Get(context.Background(), "/test", nil, &result); err != nil {
		t.Fatal(err)
	}
	if !result.OK {
		t.Error("expected ok=true")
	}
}

func TestHTTPClient_ErrorResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"errorMessage":"not found","errorCode":"404"}`))
	}))
	defer srv.Close()

	c := NewHTTPClient(HTTPClientConfig{
		BaseURL: srv.URL + "/api/v3",
		StoreID: "123",
		Token:   "test_token",
	})

	var result struct{}
	err := c.Get(context.Background(), "/missing", nil, &result)
	if err == nil {
		t.Fatal("expected error")
	}

	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.StatusCode != 404 {
		t.Errorf("expected 404, got %d", apiErr.StatusCode)
	}
}

func TestHTTPClient_RateLimitError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Retry-After", "30")
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"errorMessage":"rate limited","errorCode":"429"}`))
	}))
	defer srv.Close()

	c := NewHTTPClient(HTTPClientConfig{
		BaseURL: srv.URL + "/api/v3",
		StoreID: "123",
		Token:   "test_token",
	})

	err := c.Get(context.Background(), "/test", nil, nil)
	if err == nil {
		t.Fatal("expected error")
	}

	rlErr, ok := err.(*RateLimitError)
	if !ok {
		t.Fatalf("expected *RateLimitError, got %T", err)
	}
	if rlErr.RetryAfter.Seconds() != 30 {
		t.Errorf("expected 30s, got %s", rlErr.RetryAfter)
	}
}

func TestHTTPClient_NilBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":1}`))
	}))
	defer srv.Close()

	c := NewHTTPClient(HTTPClientConfig{
		BaseURL: srv.URL + "/api/v3",
		StoreID: "123",
		Token:   "test_token",
	})

	// v=nil should drain body without error.
	if err := c.Get(context.Background(), "/test", nil, nil); err != nil {
		t.Fatal(err)
	}
}

func TestHTTPClient_204NoContent(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c := NewHTTPClient(HTTPClientConfig{
		BaseURL: srv.URL + "/api/v3",
		StoreID: "123",
		Token:   "test_token",
	})

	var result struct{ ID int }
	if err := c.Delete(context.Background(), "/test", &result); err != nil {
		t.Fatalf("204 with non-nil v should not error, got: %v", err)
	}
}

func TestHTTPClient_Retries(t *testing.T) {
	attempts := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		attempts++
		if attempts < 3 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write([]byte(`{"errorMessage":"rate limited"}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()

	c := NewHTTPClient(HTTPClientConfig{
		BaseURL:    srv.URL + "/api/v3",
		StoreID:    "123",
		Token:      "test_token",
		MaxRetries: 3,
	})

	var result struct{ OK bool }
	if err := c.Get(context.Background(), "/test", nil, &result); err != nil {
		t.Fatal(err)
	}
	if !result.OK {
		t.Error("expected ok=true after retries")
	}
	if attempts != 3 {
		t.Errorf("expected 3 attempts, got %d", attempts)
	}
}
