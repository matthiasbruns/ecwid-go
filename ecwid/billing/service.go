package billing

import (
	"context"
	"errors"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

// Service provides access to the Ecwid application billing API.
type Service struct {
	requester api.Requester
}

// NewService creates a new billing service.
func NewService(requester api.Requester) *Service {
	return &Service{requester: requester}
}

// Charge charges the store through the application billing API.
//
// On a failed charge the API responds with a non-2xx status (e.g. 402 with an
// errorMessage such as CHARGE_LIMIT_EXCEEDED or CHARGE_DECLINED). Those surface
// as an *ecwid.APIError with StatusCode and Message populated, so callers can
// branch on the specific failure reason.
//
// API: POST /billing/transactions
// Required scope: allow_charge (application billing)
func (s *Service) Charge(ctx context.Context, req *ChargeRequest) (*ChargeResult, error) {
	if req == nil {
		return nil, errors.New("charge request must not be nil")
	}

	var result ChargeResult
	if err := s.requester.Post(ctx, "/billing/transactions", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
