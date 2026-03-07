// Package staff provides access to the Ecwid staff accounts API.
package staff

// StaffAccount represents a staff account in Ecwid.
type StaffAccount struct {
	ID             string   `json:"id,omitempty"`
	Name           string   `json:"name,omitempty"`
	Email          string   `json:"email,omitempty"`
	StaffScopes    []string `json:"staffScopes,omitempty"`
	InviteAccepted bool     `json:"inviteAccepted,omitempty"`
}

// ListResult is the response from the staff search API.
type ListResult struct {
	StaffList []StaffAccount `json:"staffList"`
}

// GetResult is the response from getting a single staff account.
// Note: Get returns only email and staffScopes, not the full StaffAccount.
type GetResult struct {
	Email       string   `json:"email,omitempty"`
	StaffScopes []string `json:"staffScopes,omitempty"`
}

// CreateRequest holds fields for creating (inviting) a staff account.
type CreateRequest struct {
	Email       string   `json:"email"`
	StaffScopes []string `json:"staffScopes,omitempty"`
}

// CreateResult represents the response from creating a staff account.
type CreateResult struct {
	Success bool `json:"success"`
}

// UpdateRequest holds fields for updating a staff account.
type UpdateRequest struct {
	Email       string   `json:"email,omitempty"`
	StaffScopes []string `json:"staffScopes,omitempty"`
}

// UpdateResult represents the response from an update operation.
type UpdateResult struct {
	UpdateCount int `json:"updateCount"`
}

// DeleteResult represents the response from a delete operation.
type DeleteResult struct {
	DeleteCount int `json:"deleteCount"`
}
