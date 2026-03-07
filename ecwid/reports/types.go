// Package reports provides access to the Ecwid store reports and stats API.
package reports

import "encoding/json"

// ReportOptions holds optional parameters for the GetReport request.
type ReportOptions struct {
	// StartedFrom is the lower bound UNIX timestamp for the report period.
	StartedFrom int64
	// EndedAt is the upper bound UNIX timestamp for the report period.
	EndedAt int64
	// TimeScaleValue is the chart time scale: "hour", "day", "week", "month", "year".
	TimeScaleValue string
	// ComparePeriod defines the comparison period.
	ComparePeriod string
}

// Report represents the response from the store reports API.
// Dataset and AggregatedData use [json.RawMessage] because the shape varies by report type.
type Report struct {
	ReportType     string          `json:"reportType"`
	StartedFrom    int64           `json:"startedFrom,omitempty"`
	EndedAt        int64           `json:"endedAt,omitempty"`
	TimeScaleValue string          `json:"timeScaleValue,omitempty"`
	FirstDayOfWeek string          `json:"firstDayOfWeek,omitempty"`
	ComparePeriod  string          `json:"comparePeriod,omitempty"`
	AggregatedData json.RawMessage `json:"aggregatedData,omitempty"`
	Dataset        json.RawMessage `json:"dataset,omitempty"`
}

// LatestStatsOptions holds optional parameters for the LatestStats request.
type LatestStatsOptions struct {
	ReviewsUpdatesRequired bool
	DomainsRequired        bool
	SubscriptionRequired   bool
	ProductCountRequired   bool
	CategoryCountRequired  bool
}

// LatestStats represents the response from the latest-stats API.
type LatestStats struct {
	ProductsUpdated        string `json:"productsUpdated,omitempty"`
	OrdersUpdated          string `json:"ordersUpdated,omitempty"`
	ReviewsUpdated         string `json:"reviewsUpdated,omitempty"`
	DomainsUpdated         string `json:"domainsUpdated,omitempty"`
	ProfileUpdated         string `json:"profileUpdated,omitempty"`
	CategoriesUpdated      string `json:"categoriesUpdated,omitempty"`
	DiscountCouponsUpdated string `json:"discountCouponsUpdated,omitempty"`
	AbandonedSalesUpdated  string `json:"abandonedSalesUpdated,omitempty"`
	CustomersUpdated       string `json:"customersUpdated,omitempty"`
	SubscriptionsUpdated   string `json:"subscriptionsUpdated,omitempty"`
	ProductCountRequired   int    `json:"productCountRequired,omitempty"`
	CategoryCountRequired  int    `json:"categoryCountRequired,omitempty"`
}
