package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HTTPClient implements [Requester] using the standard library HTTP client.
type HTTPClient struct {
	httpClient *http.Client
	logger     *slog.Logger
	baseURL    string
	token      string
}

// HTTPClientConfig holds the configuration needed to create an HTTPClient.
type HTTPClientConfig struct {
	BaseURL    string
	StoreID    string
	Token      string
	MaxRetries int
	HTTPClient *http.Client
	Logger     *slog.Logger
}

// NewHTTPClient creates a new HTTPClient with the given configuration.
func NewHTTPClient(cfg HTTPClientConfig) *HTTPClient {
	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	logger := cfg.Logger
	if logger == nil {
		logger = slog.Default()
	}

	// Wrap transport with retry if configured.
	if cfg.MaxRetries > 0 {
		wrapped := *httpClient
		wrapped.Transport = &retryTransport{
			base:       httpClient.Transport,
			maxRetries: cfg.MaxRetries,
			logger:     logger,
		}
		httpClient = &wrapped
	}

	return &HTTPClient{
		httpClient: httpClient,
		logger:     logger,
		baseURL:    strings.TrimRight(cfg.BaseURL, "/") + "/" + cfg.StoreID,
		token:      cfg.Token,
	}
}

// Get performs a GET request to the given path with optional query parameters.
func (c *HTTPClient) Get(ctx context.Context, path string, params url.Values, v any) error {
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	req, err := c.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return err
	}

	return c.do(req, v)
}

// Post performs a POST request with a JSON body.
func (c *HTTPClient) Post(ctx context.Context, path string, body, v any) error {
	reader, err := marshalBody(body)
	if err != nil {
		return err
	}

	req, err := c.newRequest(ctx, http.MethodPost, path, reader)
	if err != nil {
		return err
	}

	return c.do(req, v)
}

// Put performs a PUT request with a JSON body.
func (c *HTTPClient) Put(ctx context.Context, path string, body, v any) error {
	reader, err := marshalBody(body)
	if err != nil {
		return err
	}

	req, err := c.newRequest(ctx, http.MethodPut, path, reader)
	if err != nil {
		return err
	}

	return c.do(req, v)
}

// Delete performs a DELETE request.
func (c *HTTPClient) Delete(ctx context.Context, path string, v any) error {
	req, err := c.newRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	return c.do(req, v)
}

// newRequest creates an HTTP request with proper headers and authentication.
func (c *HTTPClient) newRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	u := c.baseURL + path

	req, err := http.NewRequestWithContext(ctx, method, u, body)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// do executes an HTTP request and decodes the JSON response into v.
func (c *HTTPClient) do(req *http.Request, v any) error {
	start := time.Now()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("execute request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	c.logger.Debug("ecwid request",
		"method", req.Method,
		"path", req.URL.Path,
		"status", resp.StatusCode,
		"duration", time.Since(start),
	)

	var reader io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gr, err := gzip.NewReader(resp.Body)
		if err != nil {
			return fmt.Errorf("decompress gzip: %w", err)
		}
		defer func() { _ = gr.Close() }()
		reader = gr
	}

	if resp.StatusCode >= 400 {
		return handleErrorResponse(resp.StatusCode, resp.Header, reader)
	}

	if v != nil {
		if err := json.NewDecoder(reader).Decode(v); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	} else {
		_, _ = io.Copy(io.Discard, reader)
	}

	return nil
}

// handleErrorResponse parses an error response and returns the appropriate error type.
func handleErrorResponse(statusCode int, header http.Header, body io.Reader) error {
	data, _ := io.ReadAll(body)

	apiErr := &APIError{
		StatusCode: statusCode,
		Message:    http.StatusText(statusCode),
	}

	if len(data) > 0 {
		_ = json.Unmarshal(data, apiErr)
		apiErr.StatusCode = statusCode
	}

	if statusCode == http.StatusTooManyRequests {
		rlErr := &RateLimitError{APIError: *apiErr}
		if after := header.Get("Retry-After"); after != "" {
			rlErr.RetryAfter = parseRetryAfter(after)
		}
		return rlErr
	}

	return apiErr
}

// marshalBody marshals a body value to JSON, returning nil reader for nil body.
func marshalBody(body any) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}
	data, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("marshal request body: %w", err)
	}
	return bytes.NewReader(data), nil
}
