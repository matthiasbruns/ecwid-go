// Package ecwid provides a Go client for the Ecwid REST API.
//
// The client is stateless — credentials are passed via [config.Config] and the
// [Client] holds only immutable configuration and an [net/http.Client]. It is
// safe for concurrent use.
//
// # Quick Start
//
//	cfg := config.Config{
//	    StoreID: "12345",
//	    Token:   "secret_abc",
//	}
//	client := ecwid.NewClient(cfg)
//
// Service methods will be available on client.Products, client.Orders, etc.
// once endpoint implementations are added.
//
// # Error Handling
//
// All API errors are returned as [*APIError], which includes the HTTP status code,
// error code, and message from the Ecwid API. Rate limit errors (429) are returned
// as [*RateLimitError] with the Retry-After duration.
//
//	var apiErr *ecwid.APIError
//	if errors.As(err, &apiErr) && apiErr.StatusCode == 404 {
//	    // handle not found
//	}
package ecwid
