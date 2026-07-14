// Package billing provides access to the Ecwid application billing API.
package billing

// ChargeRequest holds the parameters for charging a store through the
// application billing API.
//
// IdempotencyKey makes the charge safe to retry: repeating a request with the
// same key does not charge the store twice.
type ChargeRequest struct {
	IdempotencyKey string         `json:"idempotencyKey"`
	Amount         float64        `json:"amount"`
	Currency       string         `json:"currency"`
	Description    string         `json:"description,omitempty"`
	Metadata       map[string]any `json:"metadata,omitempty"`
}

// ChargeResult represents the response from a successful charge.
type ChargeResult struct {
	TransactionID       string `json:"transactionId"`
	IdempotencyKeyInUse bool   `json:"idempotencyKeyInUse"`
}
