// Package customers provides access to the Ecwid customers API.
package customers

import "encoding/json"

// Customer represents a customer in Ecwid.
// Docs: https://docs.ecwid.com/api-reference/rest-api/customers/search-customers.md
type Customer struct {
	ID                int64            `json:"id"`
	Name              string           `json:"name,omitempty"`
	Email             string           `json:"email,omitempty"`
	Registered        string           `json:"registered,omitempty"`
	Updated           string           `json:"updated,omitempty"`
	TotalOrderCount   int              `json:"totalOrderCount,omitempty"`
	CustomerGroupID   int64            `json:"customerGroupId,omitempty"`
	CustomerGroupName string           `json:"customerGroupName,omitempty"`
	BillingPerson     *json.RawMessage `json:"billingPerson,omitempty"`
	ShippingAddresses *json.RawMessage `json:"shippingAddresses,omitempty"`
	Contacts          *json.RawMessage `json:"contacts,omitempty"`
	TaxExempt         *bool            `json:"taxExempt,omitempty"`
	TaxID             string           `json:"taxId,omitempty"`
	TaxIDValid        *bool            `json:"taxIdValid,omitempty"`
	B2BB2C            string           `json:"b2b_b2c,omitempty"`
	AcceptMarketing   *bool            `json:"acceptMarketing,omitempty"`
	Lang              string           `json:"lang,omitempty"`
	Stats             *json.RawMessage `json:"stats,omitempty"`
	PrivateAdminNotes string           `json:"privateAdminNotes,omitempty"`
	Favorites         *json.RawMessage `json:"favorites,omitempty"`
}

// SearchResult is the paginated response from the customers search API.
type SearchResult struct {
	Total            int        `json:"total"`
	Count            int        `json:"count"`
	Offset           int        `json:"offset"`
	Limit            int        `json:"limit"`
	Items            []Customer `json:"items"`
	AllCustomerCount int        `json:"allCustomerCount,omitempty"`
}

// SearchOptions holds query parameters for searching customers.
type SearchOptions struct {
	Keyword             string
	Name                string
	Email               string
	UseExactEmailMatch  *bool
	Phone               string
	City                string
	PostalCode          string
	StateOrProvinceCode string
	CountryCodes        string
	CompanyName         string
	AcceptMarketing     *bool
	Lang                string
	CustomerGroupIDs    string
	MinOrderCount       int
	MaxOrderCount       int
	MinSalesValue       *float64
	MaxSalesValue       *float64
	PurchasedProductIDs string
	B2BB2C              string
	TaxExempt           *bool
	CreatedFrom         string
	CreatedTo           string
	UpdatedFrom         string
	UpdatedTo           string
	SortBy              string
	Offset              int
	Limit               int
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
