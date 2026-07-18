package webhook

import (
	"bytes"
	"embed"
	"encoding/json"
	"html/template"
	"net/http"
)

//go:embed templates/panel.html
var templateFS embed.FS

// panelTemplate is parsed once at startup; a parse failure is a build-time bug in
// the embedded template, so it panics rather than deferring to request time.
var panelTemplate = template.Must(template.ParseFS(templateFS, "templates/panel.html"))

// panelData is the data rendered into the trigger UI.
type panelData struct {
	// Enabled reports whether a delivery URL is configured. When false the Fire
	// button is disabled with an explanation rather than erroring on click.
	Enabled bool
	// URL is the delivery target, shown for context.
	URL string
	// Groups are the event families for the dropdown, in catalog order.
	Groups []panelGroup
	// CatalogJSON is the per-event fixture map the page's JS reads to prefill the
	// entityId and data fields. Go's json escapes <, > and & to \u00xx, so it is
	// safe to embed inside a <script> block.
	CatalogJSON template.JS
}

// panelGroup is one <optgroup> of events.
type panelGroup struct {
	Label  string
	Events []string
}

// panelEvent is the prefill fixture for one event, keyed by event type in the
// embedded catalog JSON.
type panelEvent struct {
	EntityID     string          `json:"entityId"`
	EntityIDType string          `json:"entityIdType"`
	HasData      bool            `json:"hasData"`
	Data         json.RawMessage `json:"data,omitempty"`
}

func (h *Handler) handleUI(w http.ResponseWriter, _ *http.Request) {
	data, err := h.buildPanelData()
	if err != nil {
		http.Error(w, "render webhook panel", http.StatusInternalServerError)
		return
	}
	// Render into a buffer so a template error becomes a clean 500 rather than a
	// half-written 200 with a truncated page.
	var buf bytes.Buffer
	if err := panelTemplate.Execute(&buf, data); err != nil {
		http.Error(w, "render webhook panel", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = buf.WriteTo(w)
}

// buildPanelData assembles the template model from the event catalog.
func (h *Handler) buildPanelData() (panelData, error) {
	specs := Catalog()

	// Group events in catalog order, preserving first-seen group order.
	var groups []panelGroup
	index := map[string]int{}
	fixtures := map[string]panelEvent{}
	for _, s := range specs {
		g := s.Group()
		i, ok := index[g]
		if !ok {
			i = len(groups)
			index[g] = i
			groups = append(groups, panelGroup{Label: g})
		}
		groups[i].Events = append(groups[i].Events, string(s.Type))
		fixtures[string(s.Type)] = panelEvent{
			EntityID:     s.EntityID,
			EntityIDType: s.EntityIDType(),
			HasData:      s.HasData(),
			Data:         s.Data,
		}
	}

	catalogJSON, err := json.Marshal(fixtures)
	if err != nil {
		return panelData{}, err
	}

	return panelData{
		Enabled:     h.trigger.Enabled(),
		URL:         h.trigger.URL(),
		Groups:      groups,
		CatalogJSON: template.JS(catalogJSON),
	}, nil
}
