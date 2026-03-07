// Package discounts provides access to the Ecwid promotions and discount coupons API.
package discounts

import "encoding/json"

// --- Shared result types ---

// CreateResult represents the response from a create operation.
type CreateResult struct {
	ID int64 `json:"id"`
}

// UpdateResult represents the response from an update operation.
type UpdateResult struct {
	UpdateCount int `json:"updateCount"`
}

// DeleteResult represents the response from a delete operation.
type DeleteResult struct {
	DeleteCount int `json:"deleteCount"`
}

// --- Promotions ---

// Promotion represents a promotion in Ecwid.
// Docs: https://docs.ecwid.com/api-reference/rest-api/discounts/promotions/search-promotions.md
type Promotion struct {
	ID                  int64            `json:"id,omitempty"`
	Name                string           `json:"name,omitempty"`
	Enabled             *bool            `json:"enabled,omitempty"`
	DiscountBase        string           `json:"discountBase,omitempty"`
	DiscountType        string           `json:"discountType,omitempty"`
	Amount              float64          `json:"amount,omitempty"`
	Triggers            *json.RawMessage `json:"triggers,omitempty"`
	Targets             *json.RawMessage `json:"targets,omitempty"`
	ExternalReferenceID string           `json:"externalReferenceId,omitempty"`
}

// PromotionSearchResult is the paginated response from the promotions search API.
type PromotionSearchResult struct {
	Total  int         `json:"total"`
	Count  int         `json:"count"`
	Offset int         `json:"offset"`
	Limit  int         `json:"limit"`
	Items  []Promotion `json:"items"`
}

// --- Coupons ---

// Coupon represents a discount coupon in Ecwid.
// Docs: https://docs.ecwid.com/api-reference/rest-api/discounts/discount-coupons/get-discount-coupon.md
type Coupon struct {
	ID                 int64            `json:"id,omitempty"`
	Name               string           `json:"name,omitempty"`
	Code               string           `json:"code,omitempty"`
	DiscountType       string           `json:"discountType,omitempty"`
	Status             string           `json:"status,omitempty"`
	Discount           float64          `json:"discount,omitempty"`
	LaunchDate         string           `json:"launchDate,omitempty"`
	ExpirationDate     string           `json:"expirationDate,omitempty"`
	TotalLimit         *float64         `json:"totalLimit,omitempty"`
	UsesLimit          string           `json:"usesLimit,omitempty"`
	ApplicationLimit   string           `json:"applicationLimit,omitempty"`
	RepeatCustomerOnly *bool            `json:"repeatCustomerOnly,omitempty"`
	CreationDate       string           `json:"creationDate,omitempty"`
	UpdateDate         string           `json:"updateDate,omitempty"`
	OrderCount         int              `json:"orderCount,omitempty"`
	CatalogLimit       *json.RawMessage `json:"catalogLimit,omitempty"`
	ShippingLimit      *json.RawMessage `json:"shippingLimit,omitempty"`
}

// CouponSearchOptions holds query parameters for searching coupons.
type CouponSearchOptions struct {
	Code         string
	DiscountType string
	Availability string
	CreatedFrom  string
	CreatedTo    string
	UpdatedFrom  string
	UpdatedTo    string
	Limit        int
	Offset       int
}

// CouponSearchResult is the paginated response from the coupons search API.
type CouponSearchResult struct {
	Total  int      `json:"total"`
	Count  int      `json:"count"`
	Offset int      `json:"offset"`
	Limit  int      `json:"limit"`
	Items  []Coupon `json:"items"`
}
