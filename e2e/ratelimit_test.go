package e2e

import (
	"errors"
	"net/http"
	"sync"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid"
)

func TestRateLimit_Handling(t *testing.T) {
	requireClient(t)
	// Fire many concurrent requests to try to trigger a 429.
	// The client is configured with MaxRetries: 3, so rate-limited
	// requests should be retried automatically.
	//
	// If all requests succeed, the retry logic handled it.
	// If we get a RateLimitError, that's also valid — it means
	// retries were exhausted and the error is properly typed.

	ctx := testContext(t)
	const concurrency = 20

	var (
		wg           sync.WaitGroup
		mu           sync.Mutex
		rateLimitHit bool
		otherErr     error
	)

	wg.Add(concurrency)
	for range concurrency {
		go func() {
			defer wg.Done()
			_, err := testClient.Products.Search(ctx, nil)
			if err != nil {
				mu.Lock()
				defer mu.Unlock()
				var rlErr *ecwid.RateLimitError
				var apiErr *ecwid.APIError
				switch {
				case errors.As(err, &rlErr):
					rateLimitHit = true
					t.Logf("rate limit hit: retry after %s", rlErr.RetryAfter)
				case errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusForbidden:
					// Token lacks scope — ignore.
				default:
					if otherErr == nil {
						otherErr = err
					}
				}
			}
		}()
	}
	wg.Wait()

	if otherErr != nil {
		t.Fatalf("unexpected error: %v", otherErr)
	}

	if rateLimitHit {
		t.Log("rate limit was hit and properly returned as *ecwid.RateLimitError")
	} else {
		t.Log("no rate limit hit — all requests succeeded (retries may have handled 429s)")
	}
}
