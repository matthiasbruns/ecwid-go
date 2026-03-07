// Package staff provides access to the Ecwid staff members API.
package staff

// StaffMember represents a staff member in Ecwid.
type StaffMember struct {
	ID          int64       `json:"id,omitempty"`
	Email       string      `json:"email,omitempty"`
	FirstName   string      `json:"firstName,omitempty"`
	LastName    string      `json:"lastName,omitempty"`
	Role        string      `json:"role,omitempty"`
	Permissions Permissions `json:"permissions,omitempty"`
}

// Permissions describes a staff member's access levels.
type Permissions struct {
	CanManageOrders    bool `json:"canManageOrders,omitempty"`
	CanManageProducts  bool `json:"canManageProducts,omitempty"`
	CanManageCustomers bool `json:"canManageCustomers,omitempty"`
	CanViewReports     bool `json:"canViewReports,omitempty"`
	CanManageSettings  bool `json:"canManageSettings,omitempty"`
}

// ListResult is the paginated response from the staff list API.
type ListResult struct {
	Total int           `json:"total"`
	Items []StaffMember `json:"items"`
}

// CreateResult represents the response from creating a staff member.
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
