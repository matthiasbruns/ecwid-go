// Package subscriptions provides access to the Ecwid recurring subscriptions API.
package subscriptions

import "encoding/json"

// Subscription represents a recurring subscription in Ecwid.
type Subscription struct {
	SubscriptionID  int64                `json:"subscriptionId"`
	CustomerID      int64                `json:"customerId"`
	Status          string               `json:"status"`
	StatusUpdated   string               `json:"statusUpdated,omitempty"`
	Created         string               `json:"created,omitempty"`
	Cancelled       string               `json:"cancelled,omitempty"`
	NextCharge      string               `json:"nextCharge,omitempty"`
	CreateTimestamp int64                `json:"createTimestamp,omitempty"`
	UpdateTimestamp int64                `json:"updateTimestamp,omitempty"`
	ChargeSettings  *ChargeSettings      `json:"chargeSettings,omitempty"`
	PaymentMethod   *SubscriptionPayment `json:"paymentMethod,omitempty"`
	OrderTemplate   json.RawMessage      `json:"orderTemplate,omitempty"`
	Orders          []SubscriptionOrder  `json:"orders,omitempty"`
}

// ChargeSettings defines the recurring charge interval for a subscription.
type ChargeSettings struct {
	RecurringInterval      string `json:"recurringInterval"`
	RecurringIntervalCount int    `json:"recurringIntervalCount"`
}

// SubscriptionPayment holds masked payment method details.
type SubscriptionPayment struct {
	CreditCardMaskedNumber string `json:"creditCardMaskedNumber,omitempty"`
	CreditCardBrand        string `json:"creditCardBrand,omitempty"`
}

// SubscriptionOrder is a brief reference to an order created by a subscription.
type SubscriptionOrder struct {
	ID         int64   `json:"id"`
	Total      float64 `json:"total"`
	CreateDate string  `json:"createDate"`
}

// SearchResult is the paginated response from the subscriptions search API.
type SearchResult struct {
	Total  int            `json:"total"`
	Count  int            `json:"count"`
	Offset int            `json:"offset"`
	Limit  int            `json:"limit"`
	Items  []Subscription `json:"items"`
}

// SearchOptions holds query parameters for searching subscriptions.
type SearchOptions struct {
	ID                int64
	CreatedFrom       string
	CreatedTo         string
	UpdatedFrom       string
	UpdatedTo         string
	NextChargeFrom    string
	NextChargeTo      string
	// Deprecated: Use NextChargeFrom instead.
	ChargeFrom        string
	// Deprecated: Use NextChargeTo instead.
	ChargeTo          string
	CancelledFrom     string
	CancelledTo       string
	CustomerID        int64
	ProductID         int64
	RecurringInterval string
	Status            string
	Offset            int
	Limit             int
}

// UpdateRequest holds fields for updating a subscription.
type UpdateRequest struct {
	ChargeSettings *ChargeSettings `json:"chargeSettings,omitempty"`
	NextCharge     string          `json:"nextCharge,omitempty"`
	Status         string          `json:"status,omitempty"`
}

// UpdateResult represents the response from an update operation.
type UpdateResult struct {
	UpdateCount int `json:"updateCount"`
}
