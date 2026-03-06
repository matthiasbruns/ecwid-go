package ecwid

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/matthiasbruns/ecwid-go/config"
)

// Client is the Ecwid API client. It is safe for concurrent use.
// The client is stateless — all configuration is immutable after creation.
type Client struct {
	cfg        config.Config
	httpClient *http.Client
	logger     *slog.Logger
	baseURL    string

	// Services — one per API domain.
	Products      *ProductService
	Orders        *OrderService
	Categories    *CategoryService
	Customers     *CustomerService
	Carts         *CartService
	Subscriptions *SubscriptionService
	Promotions    *PromotionService
	Coupons       *CouponService
	Profile       *ProfileService
	Reviews       *ReviewService
	Staff         *StaffService
	Domains       *DomainService
	Dictionaries  *DictionaryService
	Reports       *ReportService
}

// Option configures the Client.
type Option func(*Client)

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(c *http.Client) Option {
	return func(client *Client) {
		client.httpClient = c
	}
}

// WithLogger sets a custom slog logger.
func WithLogger(l *slog.Logger) Option {
	return func(client *Client) {
		client.logger = l
	}
}

// NewClient creates a new Ecwid API client with the given configuration.
//
// The client applies defaults to the config (BaseURL, Output, LogLevel)
// but does NOT validate it — call [config.Config.Validate] beforehand if needed.
func NewClient(cfg config.Config, opts ...Option) *Client {
	cfg = cfg.WithDefaults()

	c := &Client{
		cfg:        cfg,
		httpClient: http.DefaultClient,
		logger:     slog.Default(),
		baseURL:    strings.TrimRight(cfg.BaseURL, "/") + "/" + cfg.StoreID,
	}

	for _, opt := range opts {
		opt(c)
	}

	// Wrap transport with retry if configured.
	if cfg.MaxRetries > 0 {
		c.httpClient = &http.Client{
			Transport: &retryTransport{
				base:       c.httpClient.Transport,
				maxRetries: cfg.MaxRetries,
				logger:     c.logger,
			},
			Timeout: c.httpClient.Timeout,
		}
	}

	// Initialize services.
	c.Products = &ProductService{client: c}
	c.Orders = &OrderService{client: c}
	c.Categories = &CategoryService{client: c}
	c.Customers = &CustomerService{client: c}
	c.Carts = &CartService{client: c}
	c.Subscriptions = &SubscriptionService{client: c}
	c.Promotions = &PromotionService{client: c}
	c.Coupons = &CouponService{client: c}
	c.Profile = &ProfileService{client: c}
	c.Reviews = &ReviewService{client: c}
	c.Staff = &StaffService{client: c}
	c.Domains = &DomainService{client: c}
	c.Dictionaries = &DictionaryService{client: c}
	c.Reports = &ReportService{client: c}

	return c
}

// request creates an HTTP request with proper headers and authentication.
func (c *Client) request(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	u := c.baseURL + path

	req, err := http.NewRequestWithContext(ctx, method, u, body)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.cfg.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// do executes an HTTP request and decodes the JSON response into v.
// If v is nil, the response body is discarded.
func (c *Client) do(req *http.Request, v any) error {
	start := time.Now()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("execute request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	duration := time.Since(start)

	c.logger.Debug("ecwid request",
		"method", req.Method,
		"path", req.URL.Path,
		"status", resp.StatusCode,
		"duration", duration,
	)

	// Handle gzip response.
	var reader io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gr, err := gzip.NewReader(resp.Body)
		if err != nil {
			return fmt.Errorf("decompress gzip: %w", err)
		}
		defer func() { _ = gr.Close() }()
		reader = gr
	}

	// Handle error responses.
	if resp.StatusCode >= 400 {
		return c.handleErrorResponse(resp.StatusCode, resp.Header, reader)
	}

	// Decode successful response.
	if v != nil {
		if err := json.NewDecoder(reader).Decode(v); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}

	return nil
}

// handleErrorResponse parses an error response and returns the appropriate error type.
func (c *Client) handleErrorResponse(statusCode int, header http.Header, body io.Reader) error {
	data, _ := io.ReadAll(body)

	apiErr := &APIError{
		StatusCode: statusCode,
		Message:    http.StatusText(statusCode),
	}

	// Try to parse the error body.
	if len(data) > 0 {
		_ = json.Unmarshal(data, apiErr)
		apiErr.StatusCode = statusCode // Ensure status code is always set.
	}

	// Handle rate limiting.
	if statusCode == http.StatusTooManyRequests {
		rlErr := &RateLimitError{APIError: *apiErr}
		if after := header.Get("Retry-After"); after != "" {
			if secs, err := strconv.Atoi(after); err == nil {
				rlErr.RetryAfter = time.Duration(secs) * time.Second
			}
		}
		return rlErr
	}

	return apiErr
}

// get performs a GET request to the given path with optional query parameters.
func (c *Client) get(ctx context.Context, path string, params url.Values, v any) error {
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	req, err := c.request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return err
	}

	return c.do(req, v)
}

// post performs a POST request with a JSON body.
func (c *Client) post(ctx context.Context, path string, body any, v any) error {
	var reader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal request body: %w", err)
		}
		reader = strings.NewReader(string(data))
	}

	req, err := c.request(ctx, http.MethodPost, path, reader)
	if err != nil {
		return err
	}

	return c.do(req, v)
}

// put performs a PUT request with a JSON body.
func (c *Client) put(ctx context.Context, path string, body any, v any) error {
	var reader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal request body: %w", err)
		}
		reader = strings.NewReader(string(data))
	}

	req, err := c.request(ctx, http.MethodPut, path, reader)
	if err != nil {
		return err
	}

	return c.do(req, v)
}

// delete performs a DELETE request.
func (c *Client) delete(ctx context.Context, path string, v any) error {
	req, err := c.request(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	return c.do(req, v)
}

// TODO: upload method will be added when image/file upload endpoints are implemented.
