package discounts

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

// Service provides access to the Ecwid promotions and discount coupons API.
type Service struct {
	requester api.Requester
}

// NewService creates a new discounts service.
func NewService(requester api.Requester) *Service {
	return &Service{requester: requester}
}

// --- Promotions ---

// SearchPromotions returns a list of promotions (discount rules).
//
// API: GET /discount_rules
func (s *Service) SearchPromotions(ctx context.Context) (*PromotionSearchResult, error) {
	var result PromotionSearchResult
	if err := s.requester.Get(ctx, "/discount_rules", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPromotion returns a single promotion by ID.
//
// API: GET /discount_rules/{ruleId}
func (s *Service) GetPromotion(ctx context.Context, ruleID int64) (*Promotion, error) {
	if ruleID == 0 {
		return nil, errors.New("ruleID must not be zero")
	}

	path := fmt.Sprintf("/discount_rules/%d", ruleID)

	var result Promotion
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreatePromotion creates a new promotion.
//
// API: POST /discount_rules
func (s *Service) CreatePromotion(ctx context.Context, promo *Promotion) (*CreateResult, error) {
	var result CreateResult
	if err := s.requester.Post(ctx, "/discount_rules", promo, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdatePromotion updates an existing promotion.
//
// API: PUT /discount_rules/{ruleId}
func (s *Service) UpdatePromotion(ctx context.Context, ruleID int64, promo *Promotion) (*UpdateResult, error) {
	if ruleID == 0 {
		return nil, errors.New("ruleID must not be zero")
	}

	path := fmt.Sprintf("/discount_rules/%d", ruleID)

	var result UpdateResult
	if err := s.requester.Put(ctx, path, promo, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeletePromotion deletes a promotion.
//
// API: DELETE /discount_rules/{ruleId}
func (s *Service) DeletePromotion(ctx context.Context, ruleID int64) (*DeleteResult, error) {
	if ruleID == 0 {
		return nil, errors.New("ruleID must not be zero")
	}

	path := fmt.Sprintf("/discount_rules/%d", ruleID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- Coupons ---

// SearchCoupons returns a paginated list of discount coupons.
//
// API: GET /discount_coupons
func (s *Service) SearchCoupons(ctx context.Context, opts *CouponSearchOptions) (*CouponSearchResult, error) {
	q := url.Values{}
	if opts != nil {
		if opts.Code != "" {
			q.Set("code", opts.Code)
		}
		if opts.DiscountType != "" {
			q.Set("discount_type", opts.DiscountType)
		}
		if opts.Availability != "" {
			q.Set("availability", opts.Availability)
		}
		if opts.Limit > 0 {
			q.Set("limit", fmt.Sprintf("%d", opts.Limit))
		}
		if opts.Offset > 0 {
			q.Set("offset", fmt.Sprintf("%d", opts.Offset))
		}
	}

	var result CouponSearchResult
	if err := s.requester.Get(ctx, "/discount_coupons", q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCoupon returns a single coupon by ID.
//
// API: GET /discount_coupons/{couponId}
func (s *Service) GetCoupon(ctx context.Context, couponID int64) (*Coupon, error) {
	if couponID == 0 {
		return nil, errors.New("couponID must not be zero")
	}

	path := fmt.Sprintf("/discount_coupons/%d", couponID)

	var result Coupon
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateCoupon creates a new discount coupon.
//
// API: POST /discount_coupons
func (s *Service) CreateCoupon(ctx context.Context, coupon *Coupon) (*CreateResult, error) {
	var result CreateResult
	if err := s.requester.Post(ctx, "/discount_coupons", coupon, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateCoupon updates an existing coupon.
//
// API: PUT /discount_coupons/{couponId}
func (s *Service) UpdateCoupon(ctx context.Context, couponID int64, coupon *Coupon) (*UpdateResult, error) {
	if couponID == 0 {
		return nil, errors.New("couponID must not be zero")
	}

	path := fmt.Sprintf("/discount_coupons/%d", couponID)

	var result UpdateResult
	if err := s.requester.Put(ctx, path, coupon, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteCoupon deletes a discount coupon.
//
// API: DELETE /discount_coupons/{couponId}
func (s *Service) DeleteCoupon(ctx context.Context, couponID int64) (*DeleteResult, error) {
	if couponID == 0 {
		return nil, errors.New("couponID must not be zero")
	}

	path := fmt.Sprintf("/discount_coupons/%d", couponID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
