package server

import (
	"encoding/json"
	"net/http"
)

// handleHealth serves GET /_mock/health with 200 {"status":"ok"} for CI
// readiness polling.
func handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Encoding a fixed map cannot fail; ignore the error to keep the handler
	// simple, and there is nothing actionable to do if the client has gone away.
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
