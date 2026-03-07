// Package carts provides access to the Ecwid abandoned carts API.
package carts

import "encoding/json"

// Cart represents an abandoned cart in Ecwid.
type Cart struct {
	CartID          string          `json:"cartId"`
	Tax             float64         `json:"tax,omitempty"`
	Subtotal        float64         `json:"subtotal,omitempty"`
	Total           float64         `json:"total,omitempty"`
	UsdTotal        float64         `json:"usdTotal,omitempty"`
	PaymentMethod   string          `json:"paymentMethod,omitempty"`
	RefererURL      string          `json:"refererUrl,omitempty"`
	GlobalReferer   string          `json:"globalReferer,omitempty"`
	CreateDate      string          `json:"createDate,omitempty"`
	UpdateDate      string          `json:"updateDate,omitempty"`
	CreateTimestamp int64           `json:"createTimestamp,omitempty"`
	UpdateTimestamp int64           `json:"updateTimestamp,omitempty"`
	Hidden          bool            `json:"hidden,omitempty"`
	OrderComments   string          `json:"orderComments,omitempty"`
	Email           string          `json:"email,omitempty"`
	IPAddress       string          `json:"ipAddress,omitempty"`
	CustomerID      int64           `json:"customerId,omitempty"`
	CustomerGroupID int64           `json:"customerGroupId,omitempty"`
	CustomerGroup   string          `json:"customerGroup,omitempty"`
	Items           json.RawMessage `json:"items,omitempty"`
	BillingPerson   json.RawMessage `json:"billingPerson,omitempty"`
	ShippingPerson  json.RawMessage `json:"shippingPerson,omitempty"`
	ShippingOption  json.RawMessage `json:"shippingOption,omitempty"`
	DiscountCoupon  json.RawMessage `json:"discountCoupon,omitempty"`
	DiscountInfo    json.RawMessage `json:"discountInfo,omitempty"`
}

// SearchResult is the paginated response from the carts search API.
type SearchResult struct {
	Total  int    `json:"total"`
	Count  int    `json:"count"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Items  []Cart `json:"items"`
}

// SearchOptions holds query parameters for searching abandoned carts.
type SearchOptions struct {
	CreatedFrom string
	CreatedTo   string
	UpdatedFrom string
	UpdatedTo   string
	CustomerID  int64
	TotalFrom   *float64
	TotalTo     *float64
	Offset      int
	Limit       int
}

// UpdateRequest holds fields for updating an abandoned cart.
type UpdateRequest struct {
	Hidden *bool `json:"hidden,omitempty"`
}

// UpdateResult represents the response from an update operation.
type UpdateResult struct {
	UpdateCount int `json:"updateCount"`
}

// PlaceResult represents the response from converting a cart to an order.
type PlaceResult struct {
	ID                string `json:"id"`
	OrderNumber       int64  `json:"orderNumber"`
	VendorOrderNumber string `json:"vendorOrderNumber"`
	CartID            string `json:"cartId"`
}
