package reports

import (
	"context"
	"fmt"
	"net/url"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

// Service provides access to the Ecwid store reports API.
type Service struct {
	requester api.Requester
}

// NewService creates a new reports service.
func NewService(requester api.Requester) *Service {
	return &Service{requester: requester}
}

// GetReport returns store stats for a specific report type.
//
// API: GET /reports/{reportType}
// Required scope: read_store_stats
func (s *Service) GetReport(ctx context.Context, reportType string, opts *ReportOptions) (*Report, error) {
	q := url.Values{}
	if opts != nil {
		if opts.StartedFrom > 0 {
			q.Set("startedFrom", fmt.Sprintf("%d", opts.StartedFrom))
		}
		if opts.EndedAt > 0 {
			q.Set("endedAt", fmt.Sprintf("%d", opts.EndedAt))
		}
		if opts.TimeScaleValue != "" {
			q.Set("timeScaleValue", opts.TimeScaleValue)
		}
		if opts.ComparePeriod != "" {
			q.Set("comparePeriod", opts.ComparePeriod)
		}
	}

	var result Report
	if err := s.requester.Get(ctx, "/reports/"+reportType, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// LatestStats returns the latest update timestamps for various store entities.
//
// API: GET /latest-stats
// Required scope: read_store_profile
func (s *Service) LatestStats(ctx context.Context, opts *LatestStatsOptions) (*LatestStats, error) {
	q := url.Values{}
	if opts != nil {
		if opts.ReviewsUpdatesRequired {
			q.Set("reviewsUpdatesRequired", "true")
		}
		if opts.DomainsRequired {
			q.Set("domainsRequired", "true")
		}
		if opts.SubscriptionRequired {
			q.Set("subscriptionRequired", "true")
		}
		if opts.ProductCountRequired {
			q.Set("productCountRequired", "true")
		}
		if opts.CategoryCountRequired {
			q.Set("categoryCountRequired", "true")
		}
	}

	var result LatestStats
	if err := s.requester.Get(ctx, "/latest-stats", q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
