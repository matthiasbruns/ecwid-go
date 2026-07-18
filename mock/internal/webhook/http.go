package webhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/matthiasbruns/ecwid-go/ecwid/webhooks"
)

// Handler serves the webhook control API and UI panel. It is the client-facing
// surface over a [Trigger]:
//
//	POST /_mock/webhooks/trigger   fire a webhook, return the delivery result
//	GET  /_mock/webhooks/events    the 42 event types and their data shapes
//	GET  /_mock/webhooks/ui        the developer-facing trigger panel
//
// Register the routes with [Handler.Routes].
type Handler struct {
	trigger *Trigger
}

// NewHandler builds a control-API handler over trigger.
func NewHandler(trigger *Trigger) *Handler {
	return &Handler{trigger: trigger}
}

// Routes registers the webhook control-plane and UI handlers on mux.
func (h *Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("POST /_mock/webhooks/trigger", h.handleTrigger)
	mux.HandleFunc("GET /_mock/webhooks/events", h.handleEvents)
	mux.HandleFunc("GET /_mock/webhooks/ui", h.handleUI)
}

// triggerRequest is the JSON body of POST /_mock/webhooks/trigger. entityId
// accepts a JSON number or a quoted string, mirroring the wire quirk, and is
// normalized to a string.
type triggerRequest struct {
	EventType string          `json:"eventType"`
	EntityID  json.RawMessage `json:"entityId"`
	Data      json.RawMessage `json:"data"`
	Signature SignatureMode   `json:"signature"`
}

func (h *Handler) handleTrigger(w http.ResponseWriter, r *http.Request) {
	var body triggerRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body: "+err.Error())
		return
	}
	if body.EventType == "" {
		writeError(w, http.StatusBadRequest, "eventType is required")
		return
	}

	entityID, err := normalizeEntityID(body.EntityID)
	if err != nil {
		writeError(w, http.StatusBadRequest, "entityId must be a JSON string or number: "+err.Error())
		return
	}

	res, err := h.trigger.Fire(r.Context(), Request{
		EventType: webhooks.EventType(body.EventType),
		EntityID:  entityID,
		Data:      body.Data,
		Mode:      body.Signature,
	})
	switch {
	case errors.Is(err, ErrNoWebhookURL):
		writeError(w, http.StatusConflict, "no --webhook-url configured; set it to deliver webhooks")
		return
	case errors.Is(err, ErrUnknownEvent):
		writeError(w, http.StatusBadRequest, err.Error())
		return
	case err != nil:
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, res)
}

// eventInfo is one entry of GET /_mock/webhooks/events.
type eventInfo struct {
	EventType    string          `json:"eventType"`
	Group        string          `json:"group"`
	EntityIDType string          `json:"entityIdType"`
	HasData      bool            `json:"hasData"`
	EntityID     string          `json:"entityId"`
	Data         json.RawMessage `json:"data,omitempty"`
}

func (h *Handler) handleEvents(w http.ResponseWriter, _ *http.Request) {
	specs := Catalog()
	events := make([]eventInfo, len(specs))
	for i, s := range specs {
		events[i] = eventInfo{
			EventType:    string(s.Type),
			Group:        s.Group(),
			EntityIDType: s.EntityIDType(),
			HasData:      s.HasData(),
			EntityID:     s.EntityID,
			Data:         s.Data,
		}
	}
	writeJSON(w, http.StatusOK, map[string]any{"events": events})
}

// normalizeEntityID reduces a raw JSON entityId to a string: "" for absent/null,
// the string contents for a quoted string, or the digits for a number.
func normalizeEntityID(raw json.RawMessage) (string, error) {
	raw = bytes.TrimSpace(raw)
	if len(raw) == 0 || bytes.Equal(raw, []byte("null")) {
		return "", nil
	}
	if raw[0] == '"' {
		var s string
		if err := json.Unmarshal(raw, &s); err != nil {
			return "", err
		}
		return s, nil
	}
	var n json.Number
	if err := json.Unmarshal(raw, &n); err != nil {
		return "", err
	}
	return n.String(), nil
}

// writeJSON writes v as an indented JSON response with the given status.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	// The client may have gone away; there is nothing actionable to do on error.
	_ = enc.Encode(v)
}

// writeError writes a JSON error envelope.
func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}
