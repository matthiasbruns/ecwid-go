package ecwid

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

// retryTransport wraps an [http.RoundTripper] to retry requests that receive a 429 response.
// It respects the Retry-After header from the Ecwid API.
//
// Note: retries reuse the same *http.Request. For requests with bodies (POST/PUT),
// the body may already be consumed after the first attempt. This is acceptable for
// Ecwid's rate limiting (429 is returned before body consumption), but callers
// needing reliable POST/PUT retries should handle retry logic at a higher level.
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

		// Parse Retry-After header (seconds or HTTP-date).
		wait := parseRetryAfter(resp.Header.Get("Retry-After"))

		// Don't retry if this was the last attempt.
		if attempt == t.maxRetries {
			return resp, nil
		}

		logger := t.logger
		if logger == nil {
			logger = slog.Default()
		}
		logger.Warn("rate limited, retrying",
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

// parseRetryAfter parses a Retry-After header value.
// Supports delta-seconds (e.g., "30") and HTTP-date (e.g., "Fri, 06 Mar 2026 21:00:00 GMT").
// Returns a minimum of 1s on parse failure or if the date is in the past.
func parseRetryAfter(value string) time.Duration {
	const minWait = 1 * time.Second

	if value == "" {
		return minWait
	}

	// Try delta-seconds first.
	if secs, err := strconv.Atoi(value); err == nil {
		d := time.Duration(secs) * time.Second
		if d < minWait {
			return minWait
		}
		return d
	}

	// Try HTTP-date.
	if t, err := http.ParseTime(value); err == nil {
		d := time.Until(t)
		if d < minWait {
			return minWait
		}
		return d
	}

	return minWait
}
