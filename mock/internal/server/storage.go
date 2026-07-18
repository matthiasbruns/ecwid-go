package server

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"sync"
)

// App-storage value size limits. Ecwid caps private keys at 1 MB and the
// reserved public key at 256 KB; the mock enforces the same so a developer hits
// the limit locally before production does.
const (
	maxPrivateValueBytes = 1 << 20   // 1 MB
	maxPublicValueBytes  = 256 << 10 // 256 KB

	// reservedPublicKey is the single key whose value the SDK exposes to the
	// storefront via getAppPublicConfig/setAppPublicConfig. It rides the same
	// routes as any other key but carries the smaller size limit.
	reservedPublicKey = "public"
)

// storageEntry is the wire shape of a single stored key/value pair. Values are
// opaque strings kept verbatim — the mock never parses or coerces them.
type storageEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// appStorage is the mock's in-memory app-storage backend, keyed by store ID and
// then by key. A fresh, empty store per process run is intentional: the mock is
// a dev tool. It is safe for concurrent use.
type appStorage struct {
	mu   sync.RWMutex
	data map[string]map[string]string // storeID -> key -> raw value
}

// newAppStorage returns an empty app-storage backend.
func newAppStorage() *appStorage {
	return &appStorage{data: make(map[string]map[string]string)}
}

// get returns the raw value for a key and whether it was present.
func (s *appStorage) get(storeID, key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.data[storeID][key]
	return v, ok
}

// all returns every entry for a store, sorted by key for a stable response.
// The result is always non-nil so it encodes as [] rather than null.
func (s *appStorage) all(storeID string) []storageEntry {
	s.mu.RLock()
	defer s.mu.RUnlock()

	store := s.data[storeID]
	keys := make([]string, 0, len(store))
	for k := range store {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	entries := make([]storageEntry, 0, len(keys))
	for _, k := range keys {
		entries = append(entries, storageEntry{Key: k, Value: store[k]})
	}
	return entries
}

// set stores value verbatim under key for the given store.
func (s *appStorage) set(storeID, key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.data[storeID] == nil {
		s.data[storeID] = make(map[string]string)
	}
	s.data[storeID][key] = value
}

// delete removes key from the given store. Deleting an absent key is a no-op.
func (s *appStorage) delete(storeID, key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data[storeID], key)
}

// valueLimit returns the byte limit for a key: the smaller public-config limit
// for the reserved "public" key, the private limit otherwise.
func valueLimit(key string) int {
	if key == reservedPublicKey {
		return maxPublicValueBytes
	}
	return maxPrivateValueBytes
}

// authorized reports whether the request carries the access_token the mock
// issued, as an "Authorization: Bearer <token>" header. The comparison is
// constant-time to avoid leaking the token through timing.
func (s *Server) authorized(r *http.Request) bool {
	const prefix = "Bearer "
	// An empty configured token would otherwise accept "Authorization: Bearer "
	// (empty credential), so a Server built with a bare config.Config{} would run
	// with effectively no auth. Reject that explicitly.
	if s.cfg.AccessToken == "" {
		return false
	}
	h := r.Header.Get("Authorization")
	if !strings.HasPrefix(h, prefix) {
		return false
	}
	got := strings.TrimPrefix(h, prefix)
	return subtle.ConstantTimeCompare([]byte(got), []byte(s.cfg.AccessToken)) == 1
}

// handleStorageList serves GET /api/v3/{storeId}/storage: every stored entry as
// a JSON array.
func (s *Server) handleStorageList(w http.ResponseWriter, r *http.Request) {
	if !s.authorized(r) {
		writeStorageError(w, http.StatusUnauthorized, "invalid or missing access token")
		return
	}
	writeJSON(w, http.StatusOK, s.store.all(r.PathValue("storeId")))
}

// handleStorageGet serves GET /api/v3/{storeId}/storage/{key}: the single entry,
// or 404 if the key is absent. The SDK maps that 404 to null, so it must be a
// real 404 and never a 200 with an empty value.
func (s *Server) handleStorageGet(w http.ResponseWriter, r *http.Request) {
	if !s.authorized(r) {
		writeStorageError(w, http.StatusUnauthorized, "invalid or missing access token")
		return
	}
	key := r.PathValue("key")
	value, ok := s.store.get(r.PathValue("storeId"), key)
	if !ok {
		writeStorageError(w, http.StatusNotFound, fmt.Sprintf("no value stored for key %q", key))
		return
	}
	writeJSON(w, http.StatusOK, storageEntry{Key: key, Value: value})
}

// handleStoragePut serves PUT and POST /api/v3/{storeId}/storage/{key}. The
// request body is the raw value string, stored verbatim — not a {"value":...}
// wrapper. Content-Type is deliberately not validated: the SDK's ajax() sends a
// literal "Content-type: undefined", and rejecting it would reject the SDK's own
// writes.
func (s *Server) handleStoragePut(w http.ResponseWriter, r *http.Request) {
	if !s.authorized(r) {
		writeStorageError(w, http.StatusUnauthorized, "invalid or missing access token")
		return
	}
	key := r.PathValue("key")
	limit := valueLimit(key)

	// Read one byte past the limit so an over-limit write is detected without
	// buffering an unbounded body.
	body, err := io.ReadAll(io.LimitReader(r.Body, int64(limit)+1))
	if err != nil {
		writeStorageError(w, http.StatusBadRequest, "failed to read request body")
		return
	}
	if len(body) > limit {
		writeStorageError(w, http.StatusRequestEntityTooLarge,
			fmt.Sprintf("value for key %q exceeds the %d-byte limit", key, limit))
		return
	}

	s.store.set(r.PathValue("storeId"), key, string(body))
	writeJSON(w, http.StatusOK, successResponse())
}

// handleStorageDelete serves DELETE /api/v3/{storeId}/storage/{key}. Deleting a
// missing key still reports success, matching Ecwid's idempotent delete.
func (s *Server) handleStorageDelete(w http.ResponseWriter, r *http.Request) {
	if !s.authorized(r) {
		writeStorageError(w, http.StatusUnauthorized, "invalid or missing access token")
		return
	}
	s.store.delete(r.PathValue("storeId"), r.PathValue("key"))
	writeJSON(w, http.StatusOK, successResponse())
}

// successResponse is the {"success": true} body the write and delete routes
// return.
func successResponse() map[string]bool {
	return map[string]bool{"success": true}
}

// writeJSON writes v as a JSON response with the given status.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	// A client that has gone away is the only realistic encode failure and there
	// is nothing actionable to do about it here.
	_ = json.NewEncoder(w).Encode(v)
}

// writeStorageError writes a JSON error body with the given status.
func writeStorageError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"errorMessage": msg})
}
