package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/matthiasbruns/ecwid-go/mock/internal/config"
)

const (
	testToken   = "test-access-token-abc123"
	testStoreID = "1003"
)

// newStorageServer builds a server with a known access token for storage tests.
func newStorageServer() *Server {
	return New(config.Config{Port: 0, AccessToken: testToken}, discardLogger())
}

// storageReq builds an authorized request against the store's storage routes.
func storageReq(method, target string, body io.Reader) *http.Request {
	req := httptest.NewRequest(method, target, body)
	req.Header.Set("Authorization", "Bearer "+testToken)
	return req
}

func do(srv *Server, req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	srv.Handler().ServeHTTP(rec, req)
	return rec
}

func path(key string) string {
	if key == "" {
		return "/api/v3/" + testStoreID + "/storage"
	}
	return "/api/v3/" + testStoreID + "/storage/" + key
}

// TestStorage_RawBodyRoundTrip asserts the body is stored verbatim as the value,
// not wrapped in {"value": ...}.
func TestStorage_RawBodyRoundTrip(t *testing.T) {
	srv := newStorageServer()

	put := do(srv, storageReq(http.MethodPut, path("greeting"), strings.NewReader("hello")))
	if put.Code != http.StatusOK {
		t.Fatalf("PUT status = %d, want %d", put.Code, http.StatusOK)
	}
	var ok map[string]bool
	if err := json.NewDecoder(put.Body).Decode(&ok); err != nil {
		t.Fatalf("decode PUT body: %v", err)
	}
	if !ok["success"] {
		t.Errorf("PUT body = %v, want success=true", ok)
	}

	get := do(srv, storageReq(http.MethodGet, path("greeting"), http.NoBody))
	if get.Code != http.StatusOK {
		t.Fatalf("GET status = %d, want %d", get.Code, http.StatusOK)
	}
	var entry storageEntry
	if err := json.NewDecoder(get.Body).Decode(&entry); err != nil {
		t.Fatalf("decode GET body: %v", err)
	}
	if entry.Key != "greeting" || entry.Value != "hello" {
		t.Errorf("entry = %+v, want {greeting hello}", entry)
	}
}

// TestStorage_MissingKey404 asserts a real 404 for an absent key — the SDK maps
// it to null.
func TestStorage_MissingKey404(t *testing.T) {
	srv := newStorageServer()

	rec := do(srv, storageReq(http.MethodGet, path("installed"), http.NoBody))
	if rec.Code != http.StatusNotFound {
		t.Fatalf("GET missing key status = %d, want %d", rec.Code, http.StatusNotFound)
	}
}

// TestStorage_ContentTypeUndefined asserts the SDK's literal
// "Content-type: undefined" header is accepted, since Content-Type is not
// validated.
func TestStorage_ContentTypeUndefined(t *testing.T) {
	srv := newStorageServer()

	req := storageReq(http.MethodPut, path("k"), strings.NewReader("v"))
	req.Header.Set("Content-Type", "undefined")
	rec := do(srv, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("PUT with Content-type: undefined status = %d, want %d", rec.Code, http.StatusOK)
	}
}

// TestStorage_OverLimit asserts private (>1MB) and public (>256KB) over-limit
// writes are rejected.
func TestStorage_OverLimit(t *testing.T) {
	tests := []struct {
		name string
		key  string
		size int
	}{
		{"private", "big", maxPrivateValueBytes + 1},
		{"public", reservedPublicKey, maxPublicValueBytes + 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := newStorageServer()
			body := bytes.Repeat([]byte("a"), tt.size)
			rec := do(srv, storageReq(http.MethodPut, path(tt.key), bytes.NewReader(body)))
			if rec.Code != http.StatusRequestEntityTooLarge {
				t.Fatalf("over-limit PUT status = %d, want %d", rec.Code, http.StatusRequestEntityTooLarge)
			}
		})
	}
}

// TestStorage_AtLimit asserts a write exactly at the limit is accepted.
func TestStorage_AtLimit(t *testing.T) {
	srv := newStorageServer()
	body := bytes.Repeat([]byte("a"), maxPublicValueBytes)
	rec := do(srv, storageReq(http.MethodPut, path(reservedPublicKey), bytes.NewReader(body)))
	if rec.Code != http.StatusOK {
		t.Fatalf("at-limit PUT status = %d, want %d", rec.Code, http.StatusOK)
	}
}

// TestStorage_PublicKeyRoundTrips asserts the reserved "public" key uses the
// same routes as any other key.
func TestStorage_PublicKeyRoundTrips(t *testing.T) {
	srv := newStorageServer()

	value := `{"theme":"dark"}`
	if rec := do(srv, storageReq(http.MethodPut, path(reservedPublicKey), strings.NewReader(value))); rec.Code != http.StatusOK {
		t.Fatalf("PUT public status = %d, want %d", rec.Code, http.StatusOK)
	}

	get := do(srv, storageReq(http.MethodGet, path(reservedPublicKey), http.NoBody))
	if get.Code != http.StatusOK {
		t.Fatalf("GET public status = %d, want %d", get.Code, http.StatusOK)
	}
	var entry storageEntry
	if err := json.NewDecoder(get.Body).Decode(&entry); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if entry.Value != value {
		t.Errorf("public value = %q, want %q", entry.Value, value)
	}
}

// TestStorage_Unauthorized asserts wrong and missing bearer tokens are rejected
// with 401 across all storage routes.
func TestStorage_Unauthorized(t *testing.T) {
	srv := newStorageServer()
	// Seed a key so a GET would otherwise succeed.
	srv.store.set(testStoreID, "seeded", "v")

	cases := []struct {
		name string
		req  *http.Request
	}{
		{"no header", httptest.NewRequest(http.MethodGet, path("seeded"), http.NoBody)},
		{"list no header", httptest.NewRequest(http.MethodGet, path(""), http.NoBody)},
		{"put no header", httptest.NewRequest(http.MethodPut, path("k"), strings.NewReader("v"))},
		{"delete no header", httptest.NewRequest(http.MethodDelete, path("seeded"), http.NoBody)},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if rec := do(srv, tc.req); rec.Code != http.StatusUnauthorized {
				t.Errorf("status = %d, want %d", rec.Code, http.StatusUnauthorized)
			}
		})
	}

	wrong := httptest.NewRequest(http.MethodGet, path("seeded"), http.NoBody)
	wrong.Header.Set("Authorization", "Bearer not-the-token")
	if rec := do(srv, wrong); rec.Code != http.StatusUnauthorized {
		t.Errorf("wrong token status = %d, want %d", rec.Code, http.StatusUnauthorized)
	}
}

// TestStorage_List asserts the list route returns every entry as a sorted JSON
// array, and an empty store returns [] rather than null.
func TestStorage_List(t *testing.T) {
	srv := newStorageServer()

	empty := do(srv, storageReq(http.MethodGet, path(""), http.NoBody))
	if empty.Code != http.StatusOK {
		t.Fatalf("empty list status = %d, want %d", empty.Code, http.StatusOK)
	}
	if got := strings.TrimSpace(empty.Body.String()); got != "[]" {
		t.Errorf("empty list body = %q, want []", got)
	}

	for k, v := range map[string]string{"b": "2", "a": "1"} {
		if rec := do(srv, storageReq(http.MethodPut, path(k), strings.NewReader(v))); rec.Code != http.StatusOK {
			t.Fatalf("seed PUT %q status = %d", k, rec.Code)
		}
	}

	rec := do(srv, storageReq(http.MethodGet, path(""), http.NoBody))
	var entries []storageEntry
	if err := json.NewDecoder(rec.Body).Decode(&entries); err != nil {
		t.Fatalf("decode list: %v", err)
	}
	want := []storageEntry{{Key: "a", Value: "1"}, {Key: "b", Value: "2"}}
	if fmt.Sprintf("%v", entries) != fmt.Sprintf("%v", want) {
		t.Errorf("list = %v, want %v", entries, want)
	}
}

// TestStorage_Delete asserts a deleted key subsequently 404s, and deleting an
// absent key still reports success.
func TestStorage_Delete(t *testing.T) {
	srv := newStorageServer()

	if rec := do(srv, storageReq(http.MethodPut, path("k"), strings.NewReader("v"))); rec.Code != http.StatusOK {
		t.Fatalf("PUT status = %d", rec.Code)
	}
	if rec := do(srv, storageReq(http.MethodDelete, path("k"), http.NoBody)); rec.Code != http.StatusOK {
		t.Fatalf("DELETE status = %d, want %d", rec.Code, http.StatusOK)
	}
	if rec := do(srv, storageReq(http.MethodGet, path("k"), http.NoBody)); rec.Code != http.StatusNotFound {
		t.Errorf("GET after delete status = %d, want %d", rec.Code, http.StatusNotFound)
	}
	// Deleting an absent key is idempotent.
	if rec := do(srv, storageReq(http.MethodDelete, path("k"), http.NoBody)); rec.Code != http.StatusOK {
		t.Errorf("DELETE absent key status = %d, want %d", rec.Code, http.StatusOK)
	}
}

// TestStorage_Post asserts POST writes identically to PUT.
func TestStorage_Post(t *testing.T) {
	srv := newStorageServer()

	if rec := do(srv, storageReq(http.MethodPost, path("k"), strings.NewReader("posted"))); rec.Code != http.StatusOK {
		t.Fatalf("POST status = %d, want %d", rec.Code, http.StatusOK)
	}
	get := do(srv, storageReq(http.MethodGet, path("k"), http.NoBody))
	var entry storageEntry
	if err := json.NewDecoder(get.Body).Decode(&entry); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if entry.Value != "posted" {
		t.Errorf("value = %q, want posted", entry.Value)
	}
}
