package ecwid

import (
	"fmt"
	"time"
)

// APIError represents an error response from the Ecwid API.
type APIError struct {
	// StatusCode is the HTTP status code.
	StatusCode int `json:"-"`

	// Code is the error code from the API response.
	Code string `json:"errorCode"`

	// Message is the human-readable error message.
	Message string `json:"errorMessage"`
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("ecwid: %d %s: %s", e.StatusCode, e.Code, e.Message)
	}
	return fmt.Sprintf("ecwid: %d: %s", e.StatusCode, e.Message)
}

// RateLimitError is returned when the API responds with 429 Too Many Requests.
type RateLimitError struct {
	// APIError contains the underlying API error details.
	APIError

	// RetryAfter is the duration to wait before retrying, parsed from the Retry-After header.
	RetryAfter time.Duration
}

// Error implements the error interface.
func (e *RateLimitError) Error() string {
	return fmt.Sprintf("ecwid: rate limited (429), retry after %s", e.RetryAfter)
}

// Unwrap returns the underlying APIError.
func (e *RateLimitError) Unwrap() error {
	return &e.APIError
}
