// Package customers provides access to the Ecwid customers API.
package customers

import "encoding/json"

// Customer represents a customer in Ecwid.
type Customer struct {
	ID                  int64           `json:"id"`
	Email               string          `json:"email,omitempty"`
	Registered          string          `json:"registered,omitempty"`
	Updated             string          `json:"updated,omitempty"`
	Name                string          `json:"name,omitempty"`
	FirstName           string          `json:"firstName,omitempty"`
	LastName            string          `json:"lastName,omitempty"`
	City                string          `json:"city,omitempty"`
	Street              string          `json:"street,omitempty"`
	CountryCode         string          `json:"countryCode,omitempty"`
	CountryName         string          `json:"countryName,omitempty"`
	PostalCode          string          `json:"postalCode,omitempty"`
	StateOrProvinceCode string          `json:"stateOrProvinceCode,omitempty"`
	StateOrProvinceName string          `json:"stateOrProvinceName,omitempty"`
	Phone               string          `json:"phone,omitempty"`
	TotalOrderCount     int             `json:"totalOrderCount,omitempty"`
	BillingPerson       json.RawMessage `json:"billingPerson,omitempty"`
	ShippingAddresses   json.RawMessage `json:"shippingAddresses,omitempty"`
	TaxID               string          `json:"taxId,omitempty"`
	TaxIDValid          *bool           `json:"taxIdValid,omitempty"`
	AcceptMarketing     *bool           `json:"acceptMarketing,omitempty"`
}

// SearchResult is the paginated response from the customers search API.
type SearchResult struct {
	Total  int        `json:"total"`
	Count  int        `json:"count"`
	Offset int        `json:"offset"`
	Limit  int        `json:"limit"`
	Items  []Customer `json:"items"`
}

// SearchOptions holds query parameters for searching customers.
type SearchOptions struct {
	Keyword       string
	Email         string
	Name          string
	MinOrderCount int
	MaxOrderCount int
	SortBy        string
	Offset        int
	Limit         int
	CreatedFrom   string
	CreatedTo     string
	UpdatedFrom   string
	UpdatedTo     string
}

// OrdersResult is the paginated response from the customer orders endpoint.
type OrdersResult struct {
	Total  int             `json:"total"`
	Count  int             `json:"count"`
	Offset int             `json:"offset"`
	Limit  int             `json:"limit"`
	Items  json.RawMessage `json:"items"`
}

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
