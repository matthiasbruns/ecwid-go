package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/matthiasbruns/ecwid-go/ecwid/customers"
	"github.com/matthiasbruns/ecwid-go/ecwid/profile"
)

// The fixtures control API lives under the /_mock/ control plane (never under
// /api/v3/, so it can never collide with a real Ecwid route). It is the
// programmatic hook an out-of-process consumer — e.g. an integration test suite
// that boots the mock as a subprocess — uses to install its own fixtures before
// driving the simulated REST endpoints. Like the webhook control API it is a
// local dev surface and requires no bearer token.

// maxFixtureBodyBytes bounds a fixtures-control request body.
const maxFixtureBodyBytes = 1 << 20

// fixtureStoreID resolves which store a fixtures-control request targets: the
// ?storeId= query value when present, otherwise the mock's configured store. A
// consumer seeding a single store can omit it; a multi-tenant consumer names the
// store explicitly.
func (s *Server) fixtureStoreID(r *http.Request) string {
	if v := r.URL.Query().Get("storeId"); v != "" {
		return v
	}
	return s.cfg.StoreID
}

// handleFixtureCustomersPut serves POST /_mock/fixtures/customers: it upserts one
// customer (a JSON object) or many (a JSON array) into the target store. A
// customer with no id is assigned one. It replies with {"id":N} for a single
// customer or {"ids":[...]} for an array, so the caller can address the seeded
// records afterward.
func (s *Server) handleFixtureCustomersPut(w http.ResponseWriter, r *http.Request) {
	storeID := s.fixtureStoreID(r)

	body, err := io.ReadAll(io.LimitReader(r.Body, maxFixtureBodyBytes))
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "failed to read request body")
		return
	}

	// Try an array first; a single object fails that and falls through to the
	// object decode below.
	var many []customers.Customer
	if err := json.Unmarshal(body, &many); err == nil {
		ids := make([]int64, 0, len(many))
		for i := range many {
			ids = append(ids, s.fixtures.putCustomer(storeID, many[i]))
		}
		writeJSON(w, http.StatusOK, map[string]any{"ids": ids})
		return
	}

	var one customers.Customer
	if err := json.Unmarshal(body, &one); err != nil {
		writeJSONError(w, http.StatusBadRequest, "body must be a customer object or an array of customers")
		return
	}
	id := s.fixtures.putCustomer(storeID, one)
	writeJSON(w, http.StatusOK, map[string]any{"id": id})
}

// handleFixtureProfilePut serves PUT /_mock/fixtures/profile: it installs the
// store profile for the target store from the request body.
func (s *Server) handleFixtureProfilePut(w http.ResponseWriter, r *http.Request) {
	storeID := s.fixtureStoreID(r)

	body, err := io.ReadAll(io.LimitReader(r.Body, maxFixtureBodyBytes))
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "failed to read request body")
		return
	}
	var p profile.Profile
	if err := json.Unmarshal(body, &p); err != nil {
		writeJSONError(w, http.StatusBadRequest, "body must be a store profile object")
		return
	}
	s.fixtures.setProfile(storeID, &p)
	writeJSON(w, http.StatusOK, map[string]bool{"success": true})
}
