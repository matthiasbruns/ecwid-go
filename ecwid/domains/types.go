// Package domains provides access to the Ecwid store domain settings API.
package domains

// DomainSettings represents the store domain configuration.
type DomainSettings struct {
	CustomDomain        string `json:"customDomain,omitempty"`
	RegisteredDomain    string `json:"registeredDomain,omitempty"`
	SSLEnabled          bool   `json:"sslEnabled,omitempty"`
	RedirectToCanonical bool   `json:"redirectToCanonical,omitempty"`
	StoreFrontURL       string `json:"storeFrontUrl,omitempty"`
}

// DomainTemplate represents an available domain template.
type DomainTemplate struct {
	Name        string `json:"name,omitempty"`
	Price       string `json:"price,omitempty"`
	Description string `json:"description,omitempty"`
}

// WhoisResult represents the result of a domain availability check.
type WhoisResult struct {
	Available bool   `json:"available"`
	Domain    string `json:"domain,omitempty"`
}

// PurchaseRequest holds fields for purchasing a domain.
type PurchaseRequest struct {
	DomainName string `json:"domainName"`
}

// PurchaseResult represents the response from a domain purchase.
type PurchaseResult struct {
	Status string `json:"status,omitempty"`
}

// UpdateResult represents the response from an update operation.
type UpdateResult struct {
	UpdateCount int `json:"updateCount"`
}

// DeleteResult represents the response from a delete operation.
type DeleteResult struct {
	DeleteCount int `json:"deleteCount"`
}
