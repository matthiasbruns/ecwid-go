package webhook

import (
	"fmt"
	"slices"

	"github.com/matthiasbruns/ecwid-go/ecwid/webhooks"
)

// Classification is the verdict on a delivery's HTTP status: whether Ecwid would
// count it as delivered, and why. The reason is the point — a developer needs to
// know that a 208 or a 301 silently fails, not just the raw status.
type Classification struct {
	// Delivered reports whether Ecwid counts the status as a successful delivery.
	Delivered bool
	// Reason explains the verdict in words, e.g.
	// "208 -> NOT delivered: a 2xx Ecwid does not count as success".
	Reason string
}

// classify maps an HTTP status to Ecwid's delivery rules, using
// [webhooks.SuccessCodes] as the single source of truth for the success set so
// the mock can never drift from what the library enforces.
func classify(status int) Classification {
	codes := webhooks.SuccessCodes()
	if slices.Contains(codes, status) {
		return Classification{
			Delivered: true,
			Reason:    fmt.Sprintf("%d -> delivered: one of Ecwid's success codes %v", status, codes),
		}
	}

	switch {
	case status >= 300 && status < 400:
		return Classification{Reason: fmt.Sprintf(
			"%d -> NOT delivered: every 3xx is a failure — Ecwid does not follow redirects, so an endpoint that redirects (e.g. HTTP->HTTPS) silently never receives the webhook",
			status)}
	case status >= 200 && status < 300:
		return Classification{Reason: fmt.Sprintf(
			"%d -> NOT delivered: a 2xx that Ecwid does not count as success — only %v do (203 and 208 are the traps)",
			status, codes)}
	default:
		return Classification{Reason: fmt.Sprintf(
			"%d -> NOT delivered: only %v count as success", status, codes)}
	}
}
