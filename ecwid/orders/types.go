// Package orders provides access to the Ecwid orders API.
package orders

import "encoding/json"

// Order represents an order in Ecwid.
type Order struct {
	// Identifiers
	ID                string `json:"id"`
	InternalID        int64  `json:"internalId,omitempty"`
	OrderNumber       int64  `json:"orderNumber,omitempty"`
	VendorOrderNumber string `json:"vendorOrderNumber,omitempty"`
	PublicUID         string `json:"publicUid,omitempty"`

	// Timestamps
	CreateDate      string `json:"createDate,omitempty"`
	UpdateDate      string `json:"updateDate,omitempty"`
	CreateTimestamp int64  `json:"createTimestamp,omitempty"`
	UpdateTimestamp int64  `json:"updateTimestamp,omitempty"`

	// Financials
	Subtotal                        float64 `json:"subtotal,omitempty"`
	SubtotalWithoutTax              float64 `json:"subtotalWithoutTax,omitempty"`
	Total                           float64 `json:"total,omitempty"`
	TotalWithoutTax                 float64 `json:"totalWithoutTax,omitempty"`
	Tax                             float64 `json:"tax,omitempty"`
	RefundedAmount                  float64 `json:"refundedAmount,omitempty"`
	UsdTotal                        float64 `json:"usdTotal,omitempty"`
	GiftCardRedemption              float64 `json:"giftCardRedemption,omitempty"`
	TotalBeforeGiftCardRedemption   float64 `json:"totalBeforeGiftCardRedemption,omitempty"`
	CouponDiscount                  float64 `json:"couponDiscount,omitempty"`
	VolumeDiscount                  float64 `json:"volumeDiscount,omitempty"`
	MembershipBasedDiscount         float64 `json:"membershipBasedDiscount,omitempty"`
	TotalAndMembershipBasedDiscount float64 `json:"totalAndMembershipBasedDiscount,omitempty"`
	Discount                        float64 `json:"discount,omitempty"`

	// Status
	PaymentStatus  string `json:"paymentStatus,omitempty"`
	PaymentMethod  string `json:"paymentMethod,omitempty"`
	PaymentModule  string `json:"paymentModule,omitempty"`
	PaymentMessage string `json:"paymentMessage,omitempty"`

	FulfillmentStatus     string `json:"fulfillmentStatus,omitempty"`
	ExternalTransactionID string `json:"externalTransactionId,omitempty"`

	// Customer
	Email              string `json:"email,omitempty"`
	CustomerID         int64  `json:"customerId,omitempty"`
	CustomerGroup      string `json:"customerGroup,omitempty"`
	CustomerGroupID    int64  `json:"customerGroupId,omitempty"`
	AcceptMarketing    *bool  `json:"acceptMarketing,omitempty"`
	CustomerUserAgent  string `json:"customerUserAgent,omitempty"`
	CustomerTaxExempt  *bool  `json:"customerTaxExempt,omitempty"`
	CustomerTaxID      string `json:"customerTaxId,omitempty"`
	CustomerTaxIDValid *bool  `json:"customerTaxIdValid,omitempty"`
	B2BB2C             string `json:"b2b_b2c,omitempty"`

	// Addresses
	BillingPerson  json.RawMessage `json:"billingPerson,omitempty"`
	ShippingPerson json.RawMessage `json:"shippingPerson,omitempty"`

	// Shipping
	ShippingOption    json.RawMessage `json:"shippingOption,omitempty"`
	TrackingNumber    string          `json:"trackingNumber,omitempty"`
	HandlingFee       json.RawMessage `json:"handlingFee,omitempty"`
	PredictedPackages json.RawMessage `json:"predictedPackages,omitempty"`
	Shipments         json.RawMessage `json:"shipments,omitempty"`

	// Items
	Items json.RawMessage `json:"items,omitempty"`

	// Discounts
	CustomDiscount json.RawMessage `json:"customDiscount,omitempty"`
	DiscountInfo   json.RawMessage `json:"discountInfo,omitempty"`
	DiscountCoupon json.RawMessage `json:"discountCoupon,omitempty"`

	// Surcharges
	CustomSurcharges json.RawMessage `json:"customSurcharges,omitempty"`

	// Tax
	PricesIncludeTax   *bool           `json:"pricesIncludeTax,omitempty"`
	TaxesOnShipping    json.RawMessage `json:"taxesOnShipping,omitempty"`
	ReversedTaxApplied *bool           `json:"reversedTaxApplied,omitempty"`

	// Notes
	OrderComments    string `json:"orderComments,omitempty"`
	PrivateAdminNotes string `json:"privateAdminNotes,omitempty"`

	// Tracking/Referral
	RefererURL  string          `json:"refererUrl,omitempty"`
	GlobalReferer string        `json:"globalReferer,omitempty"`
	AffiliateID string          `json:"affiliateId,omitempty"`
	RefererId   int64           `json:"refererId,omitempty"`
	IPAddress   string          `json:"ipAddress,omitempty"`
	UtmData     json.RawMessage `json:"utmData,omitempty"`
	UtmDataSets json.RawMessage `json:"utmDataSets,omitempty"`

	// External
	ExternalFulfillment *bool           `json:"externalFulfillment,omitempty"`
	ExternalOrderID     string          `json:"externalOrderId,omitempty"`
	ExternalOrderData   json.RawMessage `json:"externalOrderData,omitempty"`

	// Other
	Hidden           *bool           `json:"hidden,omitempty"`
	ExtraFields      json.RawMessage `json:"extraFields,omitempty"`
	OrderExtraFields json.RawMessage `json:"orderExtraFields,omitempty"`
	Invoices         json.RawMessage `json:"invoices,omitempty"`
	Refunds          json.RawMessage `json:"refunds,omitempty"`
	CreditCardStatus json.RawMessage `json:"creditCardStatus,omitempty"`
	PickupTime       string          `json:"pickupTime,omitempty"`
	Lang             string          `json:"lang,omitempty"`
	AdditionalInfo   json.RawMessage `json:"additionalInfo,omitempty"`
	PaymentParams    json.RawMessage `json:"paymentParams,omitempty"`
}

// SearchResult is the paginated response from the orders search API.
type SearchResult struct {
	Total  int     `json:"total"`
	Count  int     `json:"count"`
	Offset int     `json:"offset"`
	Limit  int     `json:"limit"`
	Items  []Order `json:"items"`
}

// SearchOptions holds query parameters for searching orders.
type SearchOptions struct {
	IDs                   string
	Keywords              string
	Email                 string
	CustomerID            int64
	ProductID             string
	TotalFrom             *float64
	TotalTo               *float64
	CreatedFrom           string
	CreatedTo             string
	UpdatedFrom           string
	UpdatedTo             string
	PickupTimeFrom        string
	PickupTimeTo          string
	ShippingMethod        string
	FulfillmentStatus     string
	PaymentMethod         string
	PaymentModule         string
	PaymentStatus         string
	AcceptMarketing       *bool
	ContainsPreorderItems *bool
	CouponCode            string
	SubscriptionID        int64
	RefererId             int64
	ResponseFields        string
	Offset                int
	Limit                 int
}

// CreateRequest holds fields for creating a new order.
type CreateRequest struct {
	Subtotal          float64         `json:"subtotal"`
	Total             float64         `json:"total"`
	Email             string          `json:"email"`
	FulfillmentStatus string          `json:"fulfillmentStatus"`
	PaymentStatus     string          `json:"paymentStatus"`
	PaymentMethod     string          `json:"paymentMethod,omitempty"`
	PaymentModule     string          `json:"paymentModule,omitempty"`
	Tax               float64         `json:"tax,omitempty"`
	Items             json.RawMessage `json:"items,omitempty"`
	BillingPerson     json.RawMessage `json:"billingPerson,omitempty"`
	ShippingPerson    json.RawMessage `json:"shippingPerson,omitempty"`
	ShippingOption    json.RawMessage `json:"shippingOption,omitempty"`
	HandlingFee       json.RawMessage `json:"handlingFee,omitempty"`
	DiscountCoupon    json.RawMessage `json:"discountCoupon,omitempty"`
	DiscountInfo      json.RawMessage `json:"discountInfo,omitempty"`
	CustomSurcharges  json.RawMessage `json:"customSurcharges,omitempty"`
	ExtraFields       json.RawMessage `json:"extraFields,omitempty"`
	OrderExtraFields  json.RawMessage `json:"orderExtraFields,omitempty"`

	CustomerID         int64   `json:"customerId,omitempty"`
	CustomerGroup      string  `json:"customerGroup,omitempty"`
	CustomerGroupID    int64   `json:"customerGroupId,omitempty"`
	CustomerTaxExempt  *bool   `json:"customerTaxExempt,omitempty"`
	CustomerTaxID      string  `json:"customerTaxId,omitempty"`
	CustomerTaxIDValid *bool   `json:"customerTaxIdValid,omitempty"`
	B2BB2C             string  `json:"b2b_b2c,omitempty"`
	AcceptMarketing    *bool   `json:"acceptMarketing,omitempty"`
	IPAddress          string  `json:"ipAddress,omitempty"`
	RefererURL         string  `json:"refererUrl,omitempty"`
	GlobalReferer      string  `json:"globalReferer,omitempty"`
	RefererId          int64   `json:"refererId,omitempty"`
	OrderComments      string  `json:"orderComments,omitempty"`
	PrivateAdminNotes  string  `json:"privateAdminNotes,omitempty"`
	TrackingNumber     string  `json:"trackingNumber,omitempty"`
	CreateDate         string  `json:"createDate,omitempty"`
	Hidden             *bool   `json:"hidden,omitempty"`
	Discount           float64 `json:"discount,omitempty"`
	CouponDiscount     float64 `json:"couponDiscount,omitempty"`
	VolumeDiscount     float64 `json:"volumeDiscount,omitempty"`

	GiftCardRedemption            float64 `json:"giftCardRedemption,omitempty"`
	TotalBeforeGiftCardRedemption float64 `json:"totalBeforeGiftCardRedemption,omitempty"`
	PricesIncludeTax              *bool   `json:"pricesIncludeTax,omitempty"`

	Lang                             string          `json:"lang,omitempty"`
	ExternalTransactionID            string          `json:"externalTransactionId,omitempty"`
	ExternalFulfillment              *bool           `json:"externalFulfillment,omitempty"`
	ExternalOrderID                  string          `json:"externalOrderId,omitempty"`
	UtmData                          json.RawMessage `json:"utmData,omitempty"`
	UtmDataSets                      json.RawMessage `json:"utmDataSets,omitempty"`
	DisableAllCustomerNotifications  *bool           `json:"disableAllCustomerNotifications,omitempty"`
}

// CreateResult represents the response from creating an order.
type CreateResult struct {
	ID      int64  `json:"id"`
	OrderID string `json:"orderid"`
}

// UpdateRequest holds fields for updating an order. All fields are optional.
type UpdateRequest struct {
	Subtotal          *float64        `json:"subtotal,omitempty"`
	Total             *float64        `json:"total,omitempty"`
	Email             string          `json:"email,omitempty"`
	FulfillmentStatus string          `json:"fulfillmentStatus,omitempty"`
	PaymentStatus     string          `json:"paymentStatus,omitempty"`
	PaymentMethod     string          `json:"paymentMethod,omitempty"`
	PaymentModule     string          `json:"paymentModule,omitempty"`
	Tax               *float64        `json:"tax,omitempty"`
	Items             json.RawMessage `json:"items,omitempty"`
	BillingPerson     json.RawMessage `json:"billingPerson,omitempty"`
	ShippingPerson    json.RawMessage `json:"shippingPerson,omitempty"`
	ShippingOption    json.RawMessage `json:"shippingOption,omitempty"`
	HandlingFee       json.RawMessage `json:"handlingFee,omitempty"`
	DiscountCoupon    json.RawMessage `json:"discountCoupon,omitempty"`
	DiscountInfo      json.RawMessage `json:"discountInfo,omitempty"`
	CustomSurcharges  json.RawMessage `json:"customSurcharges,omitempty"`
	ExtraFields       json.RawMessage `json:"extraFields,omitempty"`
	OrderExtraFields  json.RawMessage `json:"orderExtraFields,omitempty"`

	CustomerID         int64   `json:"customerId,omitempty"`
	CustomerGroup      string  `json:"customerGroup,omitempty"`
	CustomerGroupID    int64   `json:"customerGroupId,omitempty"`
	CustomerTaxExempt  *bool   `json:"customerTaxExempt,omitempty"`
	CustomerTaxID      string  `json:"customerTaxId,omitempty"`
	CustomerTaxIDValid *bool   `json:"customerTaxIdValid,omitempty"`
	B2BB2C             string  `json:"b2b_b2c,omitempty"`
	AcceptMarketing    *bool   `json:"acceptMarketing,omitempty"`
	IPAddress          string  `json:"ipAddress,omitempty"`
	RefererURL         string  `json:"refererUrl,omitempty"`
	GlobalReferer      string  `json:"globalReferer,omitempty"`
	RefererId          int64   `json:"refererId,omitempty"`
	OrderComments      string  `json:"orderComments,omitempty"`
	PrivateAdminNotes  string  `json:"privateAdminNotes,omitempty"`
	TrackingNumber     string  `json:"trackingNumber,omitempty"`
	Hidden             *bool   `json:"hidden,omitempty"`
	Discount           *float64 `json:"discount,omitempty"`
	CouponDiscount     *float64 `json:"couponDiscount,omitempty"`
	VolumeDiscount     *float64 `json:"volumeDiscount,omitempty"`

	GiftCardRedemption            *float64 `json:"giftCardRedemption,omitempty"`
	TotalBeforeGiftCardRedemption *float64 `json:"totalBeforeGiftCardRedemption,omitempty"`
	PricesIncludeTax              *bool    `json:"pricesIncludeTax,omitempty"`

	Lang                             string          `json:"lang,omitempty"`
	ExternalTransactionID            string          `json:"externalTransactionId,omitempty"`
	ExternalFulfillment              *bool           `json:"externalFulfillment,omitempty"`
	ExternalOrderID                  string          `json:"externalOrderId,omitempty"`
	UtmData                          json.RawMessage `json:"utmData,omitempty"`
	UtmDataSets                      json.RawMessage `json:"utmDataSets,omitempty"`
	DisableAllCustomerNotifications  *bool           `json:"disableAllCustomerNotifications,omitempty"`
}

// UpdateResult represents the response from an update operation.
type UpdateResult struct {
	UpdateCount int `json:"updateCount"`
}

// DeleteResult represents the response from a delete operation.
type DeleteResult struct {
	DeleteCount int `json:"deleteCount"`
}

// DeletedOrdersResult is the paginated response for deleted orders history.
type DeletedOrdersResult struct {
	Total  int            `json:"total"`
	Count  int            `json:"count"`
	Offset int            `json:"offset"`
	Limit  int            `json:"limit"`
	Items  []DeletedOrder `json:"items"`
}

// DeletedOrder represents a deleted order reference.
type DeletedOrder struct {
	ID   int64  `json:"id"`
	Date string `json:"date"`
}

// DeletedOrdersOptions holds query parameters for the deleted orders endpoint.
type DeletedOrdersOptions struct {
	FromDate string
	ToDate   string
	Offset   int
	Limit    int
}

// RepeatOrderURLResult represents the response from the repeat order URL endpoint.
type RepeatOrderURLResult struct {
	RepeatOrderURL string `json:"repeatOrderUrl"`
}

// Invoice represents a tax invoice for an order.
type Invoice struct {
	InternalID int64  `json:"internalId"`
	ID         string `json:"id"`
	Created    string `json:"created"`
	Link       string `json:"link"`
	Type       string `json:"type"`
}

// InvoicesResult represents the response from the invoices endpoint.
type InvoicesResult struct {
	Invoices []Invoice `json:"invoices"`
}

// CreateInvoiceResult represents the response from creating an invoice.
type CreateInvoiceResult struct {
	ID int64 `json:"id"`
}

// ExtraField represents an order extra field.
type ExtraField struct {
	ID                         string `json:"id"`
	Value                      string `json:"value"`
	CustomerInputType          string `json:"customerInputType,omitempty"`
	Title                      string `json:"title"`
	OrderDetailsDisplaySection string `json:"orderDetailsDisplaySection,omitempty"`
	OrderBy                    string `json:"orderBy,omitempty"`
	ShowInNotifications        *bool  `json:"showInNotifications,omitempty"`
	ShowInInvoice              *bool  `json:"showInInvoice,omitempty"`
}

// CreateExtraFieldResult represents the response from creating an extra field.
type CreateExtraFieldResult struct {
	CreateCount int `json:"createCount"`
}

// CalculateRequest holds fields for calculating order details.
type CalculateRequest struct {
	Items            json.RawMessage `json:"items"`
	BillingPerson    json.RawMessage `json:"billingPerson,omitempty"`
	ShippingPerson   json.RawMessage `json:"shippingPerson,omitempty"`
	CustomSurcharges json.RawMessage `json:"customSurcharges,omitempty"`
	DiscountInfo     json.RawMessage `json:"discountInfo,omitempty"`
	GiftCardCode     string          `json:"giftCardCode,omitempty"`
	HandlingFee      json.RawMessage `json:"handlingFee,omitempty"`
	ShippingOption   json.RawMessage `json:"shippingOption,omitempty"`
	DiscountCoupon   json.RawMessage `json:"discountCoupon,omitempty"`
	Email            string          `json:"email,omitempty"`
	IPAddress        string          `json:"ipAddress,omitempty"`
	CustomerID       int64           `json:"customerId,omitempty"`
	CustomerTaxExempt *bool          `json:"customerTaxExempt,omitempty"`
}
