package ecwid

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

// retryTransport wraps an [http.RoundTripper] to retry requests that receive a 429 response.
// It respects the Retry-After header from the Ecwid API.
type retryTransport struct {
	base       http.RoundTripper
	maxRetries int
	logger     *slog.Logger
}

// RoundTrip implements [http.RoundTripper].
func (t *retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	transport := t.base
	if transport == nil {
		transport = http.DefaultTransport
	}

	var resp *http.Response
	var err error

	for attempt := 0; attempt <= t.maxRetries; attempt++ {
		resp, err = transport.RoundTrip(req)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusTooManyRequests {
			return resp, nil
		}

		// Parse Retry-After header.
		wait := 1 * time.Second // Default wait.
		if after := resp.Header.Get("Retry-After"); after != "" {
			if secs, parseErr := strconv.Atoi(after); parseErr == nil {
				wait = time.Duration(secs) * time.Second
			}
		}

		// Don't retry if this was the last attempt.
		if attempt == t.maxRetries {
			return resp, nil
		}

		t.logger.Warn("rate limited, retrying",
			"attempt", attempt+1,
			"max_retries", t.maxRetries,
			"retry_after", wait,
			"method", req.Method,
			"path", req.URL.Path,
		)

		// Close the body before retrying.
		_ = resp.Body.Close()

		// Wait before retrying. Respect context cancellation.
		timer := time.NewTimer(wait)
		select {
		case <-req.Context().Done():
			timer.Stop()
			return nil, req.Context().Err()
		case <-timer.C:
		}
	}

	return resp, nil
}
