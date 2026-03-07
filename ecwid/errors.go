package ecwid

import "github.com/matthiasbruns/ecwid-go/ecwid/internal/api"

// APIError represents an error response from the Ecwid API.
// Use type assertion to inspect error details:
//
//	var apiErr *ecwid.APIError
//	if errors.As(err, &apiErr) {
//	    fmt.Println(apiErr.StatusCode, apiErr.Message)
//	}
type APIError = api.APIError

// RateLimitError is returned when the API responds with 429 Too Many Requests.
// It embeds [APIError] and adds the RetryAfter duration.
type RateLimitError = api.RateLimitError
