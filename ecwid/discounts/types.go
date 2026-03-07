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

// Promotion represents a discount rule (promotion) in Ecwid.
type Promotion struct {
	ID          int64           `json:"id,omitempty"`
	Name        string          `json:"name,omitempty"`
	Enabled     *bool           `json:"enabled,omitempty"`
	Conditions  json.RawMessage `json:"conditions,omitempty"`
	Rewards     json.RawMessage `json:"rewards,omitempty"`
	Priority    int             `json:"priority,omitempty"`
	Description string          `json:"description,omitempty"`
}

// PromotionSearchResult is the paginated response from the promotions search API.
type PromotionSearchResult struct {
	Total int         `json:"total"`
	Items []Promotion `json:"items"`
}

// --- Coupons ---

// Coupon represents a discount coupon in Ecwid.
type Coupon struct {
	ID                 int64    `json:"id,omitempty"`
	Name               string   `json:"name,omitempty"`
	Code               string   `json:"code,omitempty"`
	DiscountType       string   `json:"discountType,omitempty"`
	Status             string   `json:"status,omitempty"`
	Discount           float64  `json:"discount,omitempty"`
	LaunchDate         string   `json:"launchDate,omitempty"`
	ExpirationDate     string   `json:"expirationDate,omitempty"`
	TotalLimit         *float64 `json:"totalLimit,omitempty"`
	UsesLimit          string   `json:"usesLimit,omitempty"`
	RepeatPurchaseOnly *bool    `json:"repeatPurchaseOnly,omitempty"`
	CreationDate       string   `json:"creationDate,omitempty"`
	UpdateDate         string   `json:"updateDate,omitempty"`
	OrderCount         int      `json:"orderCount,omitempty"`
}

// CouponSearchOptions holds query parameters for searching coupons.
type CouponSearchOptions struct {
	Code         string
	DiscountType string
	Availability string
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
