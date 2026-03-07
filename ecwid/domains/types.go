// Package domains provides access to the Ecwid store domains API.
package domains

// DomainsResult is the response from the domains search API.
type DomainsResult struct {
	InstantSiteDomain *InstantSiteDomain `json:"instantSiteDomain,omitempty"`
	PurchasedDomains  []PurchasedDomain  `json:"purchasedDomains,omitempty"`
}

// InstantSiteDomain holds the Ecwid Instant Site domain details.
type InstantSiteDomain struct {
	PrimaryInstantSiteDomain       string `json:"primaryInstantSiteDomain,omitempty"`
	PrimaryInstantSiteDomainStatus string `json:"primaryInstantSiteDomainStatus,omitempty"`
	EcwidSubdomain                 string `json:"ecwidSubdomain,omitempty"`
	InstantSiteIPAddress           string `json:"instantSiteIpAddress,omitempty"`
	InstantSiteURL                 string `json:"instantSiteUrl,omitempty"`
}

// PurchasedDomain represents a domain purchased through Ecwid.
type PurchasedDomain struct {
	ID                      int64  `json:"id"`
	Name                    string `json:"name,omitempty"`
	Status                  string `json:"status,omitempty"`
	ConnectedToInstantSite  bool   `json:"connectedToInstantSite,omitempty"`
	PrimaryDomain           bool   `json:"primaryDomain,omitempty"`
	RedirectToPrimaryDomain bool   `json:"redirectToPrimaryDomain,omitempty"`
}

// PurchaseRequest holds contact info for purchasing a domain.
type PurchaseRequest struct {
	DomainName          string `json:"domainName"`
	FirstName           string `json:"firstName,omitempty"`
	LastName            string `json:"lastName,omitempty"`
	Email               string `json:"email,omitempty"`
	Street              string `json:"street,omitempty"`
	City                string `json:"city,omitempty"`
	CountryCode         string `json:"countryCode,omitempty"`
	PostalCode          string `json:"postalCode,omitempty"`
	StateOrProvinceCode string `json:"stateOrProvinceCode,omitempty"`
	Phone               string `json:"phone,omitempty"`
	CompanyName         string `json:"companyName,omitempty"`
}

// PurchaseResult represents the response from purchasing a domain.
type PurchaseResult struct {
	ID                      int64  `json:"id"`
	Name                    string `json:"name,omitempty"`
	Status                  string `json:"status,omitempty"`
	ConnectedToInstantSite  bool   `json:"connectedToInstantSite,omitempty"`
	PrimaryDomain           bool   `json:"primaryDomain,omitempty"`
	RedirectToPrimaryDomain bool   `json:"redirectToPrimaryDomain,omitempty"`
}
