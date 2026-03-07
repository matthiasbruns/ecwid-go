// Package api provides the internal HTTP transport for the Ecwid API client.
// This package is internal — consumers of the ecwid module cannot import it directly.
package api

import (
	"context"
	"net/url"
)

// Requester defines the HTTP operations available to domain services.
// Domain packages use this interface without knowing the transport details.
type Requester interface {
	Get(ctx context.Context, path string, params url.Values, v any) error
	Post(ctx context.Context, path string, body, v any) error
	Put(ctx context.Context, path string, body, v any) error
	Delete(ctx context.Context, path string, v any) error
}
