// Package profile provides access to the Ecwid store profile API.
package profile

import "encoding/json"

// Profile represents the store profile returned by the Ecwid API.
type Profile struct {
	GeneralInfo            *GeneralInfo          `json:"generalInfo,omitempty"`
	Account                *Account              `json:"account,omitempty"`
	Settings               *Settings             `json:"settings,omitempty"`
	MailSettings           *MailSettings         `json:"mailNotifications,omitempty"`
	Company                *Company              `json:"company,omitempty"`
	FormatsAndUnits        *FormatsAndUnits      `json:"formatsAndUnits,omitempty"`
	Languages              *Languages            `json:"languages,omitempty"`
	Shipping               *Shipping             `json:"shipping,omitempty"`
	TaxSettings            *TaxSettings          `json:"taxSettings,omitempty"`
	Zones                  []Zone                `json:"zones,omitempty"`
	BusinessRegistrationID *BusinessRegistration `json:"businessRegistrationID,omitempty"`
	LegalPagesSettings     *LegalPagesSettings   `json:"legalPagesSettings,omitempty"`
	Payment                json.RawMessage       `json:"payment,omitempty"`
	FeatureToggles         json.RawMessage       `json:"featureToggles,omitempty"`
	DesignSettings         json.RawMessage       `json:"designSettings,omitempty"`
	ProductFiltersSettings json.RawMessage       `json:"productFiltersSettings,omitempty"`
}

// GeneralInfo holds basic store information.
type GeneralInfo struct {
	StoreID        int64        `json:"storeId"`
	StoreURL       string       `json:"storeUrl,omitempty"`
	StarterSite    *StarterSite `json:"starterSite,omitempty"`
	InstantSiteURL string       `json:"instantSiteUrl,omitempty"`
}

// StarterSite holds the Ecwid Instant Site configuration.
type StarterSite struct {
	EcwidSubdomain string `json:"ecwidSubdomain,omitempty"`
	CustomDomain   string `json:"customDomain,omitempty"`
	GeneratedURL   string `json:"generatedUrl,omitempty"`
	StoreLogoURL   string `json:"storeLogoUrl,omitempty"`
}

// Account holds store account details.
type Account struct {
	AccountName       string   `json:"accountName,omitempty"`
	AccountNickName   string   `json:"accountNickName,omitempty"`
	AccountEmail      string   `json:"accountEmail,omitempty"`
	AvailableFeatures []string `json:"availableFeatures,omitempty"`
	WhiteLabel        *bool    `json:"whiteLabel,omitempty"`
}

// Settings holds general store settings.
type Settings struct {
	Closed                             *bool                   `json:"closed,omitempty"`
	StoreName                          string                  `json:"storeName,omitempty"`
	StoreDescription                   string                  `json:"storeDescription,omitempty"`
	GoogleRemarketingEnabled           *bool                   `json:"googleRemarketingEnabled,omitempty"`
	GoogleAnalyticsID                  string                  `json:"googleAnalyticsId,omitempty"`
	FbPixelID                          string                  `json:"fbPixelId,omitempty"`
	OrderCommentsEnabled               *bool                   `json:"orderCommentsEnabled,omitempty"`
	OrderCommentsCaption               string                  `json:"orderCommentsCaption,omitempty"`
	OrderCommentsRequired              *bool                   `json:"orderCommentsRequired,omitempty"`
	HideOutOfStockProductsInStorefront *bool                   `json:"hideOutOfStockProductsInStorefront,omitempty"`
	AskCompanyName                     *bool                   `json:"askCompanyName,omitempty"`
	FavoritesEnabled                   *bool                   `json:"favoritesEnabled,omitempty"`
	AbandonedSales                     *AbandonedSalesSettings `json:"abandonedSales,omitempty"`
	SalePrice                          *SalePriceSettings      `json:"salePrice,omitempty"`
}

// AbandonedSalesSettings holds abandoned cart recovery settings.
type AbandonedSalesSettings struct {
	AutoAbandonedSalesRecovery *bool `json:"autoAbandonedSalesRecovery,omitempty"`
}

// SalePriceSettings holds sale price display settings.
type SalePriceSettings struct {
	DisplayOnProductList *bool  `json:"displayOnProductList,omitempty"`
	OldPriceLabel        string `json:"oldPriceLabel,omitempty"`
	DisplayDiscount      string `json:"displayDiscount,omitempty"`
}

// MailSettings holds mail notification settings.
type MailSettings struct {
	AdminNotificationEmails       []string `json:"adminNotificationEmails,omitempty"`
	CustomerNotificationFromEmail string   `json:"customerNotificationFromEmail,omitempty"`
}

// Company holds store company information.
type Company struct {
	CompanyName         string `json:"companyName,omitempty"`
	Email               string `json:"email,omitempty"`
	Street              string `json:"street,omitempty"`
	City                string `json:"city,omitempty"`
	CountryCode         string `json:"countryCode,omitempty"`
	PostalCode          string `json:"postalCode,omitempty"`
	StateOrProvinceCode string `json:"stateOrProvinceCode,omitempty"`
	Phone               string `json:"phone,omitempty"`
}

// FormatsAndUnits holds formatting settings.
type FormatsAndUnits struct {
	Currency                       string  `json:"currency,omitempty"`
	CurrencyPrefix                 string  `json:"currencyPrefix,omitempty"`
	CurrencySuffix                 string  `json:"currencySuffix,omitempty"`
	CurrencyGroupSeparator         string  `json:"currencyGroupSeparator,omitempty"`
	CurrencyDecimalSeparator       string  `json:"currencyDecimalSeparator,omitempty"`
	CurrencyPrecision              int     `json:"currencyPrecision,omitempty"`
	CurrencyTruncateZeroFractional *bool   `json:"currencyTruncateZeroFractional,omitempty"`
	CurrencyRate                   float64 `json:"currencyRate,omitempty"`
	WeightUnit                     string  `json:"weightUnit,omitempty"`
	WeightPrecision                int     `json:"weightPrecision,omitempty"`
	WeightGroupSeparator           string  `json:"weightGroupSeparator,omitempty"`
	WeightDecimalSeparator         string  `json:"weightDecimalSeparator,omitempty"`
	WeightTruncateZeroFractional   *bool   `json:"weightTruncateZeroFractional,omitempty"`
	DateFormat                     string  `json:"dateFormat,omitempty"`
	TimeFormat                     string  `json:"timeFormat,omitempty"`
	Timezone                       string  `json:"timezone,omitempty"`
	DimensionsUnit                 string  `json:"dimensionsUnit,omitempty"`
	OrderNumberPrefix              string  `json:"orderNumberPrefix,omitempty"`
	OrderNumberSuffix              string  `json:"orderNumberSuffix,omitempty"`
}

// Languages holds store language settings.
type Languages struct {
	EnabledLanguages []string `json:"enabledLanguages,omitempty"`
	DefaultLanguage  string   `json:"defaultLanguage,omitempty"`
	FbMessengerLang  string   `json:"facebookPreferredLocale,omitempty"`
}

// Shipping holds store shipping settings.
type Shipping struct {
	HandlingFee *HandlingFee `json:"handlingFee,omitempty"`
}

// HandlingFee represents a handling fee configuration.
type HandlingFee struct {
	Name        string  `json:"name,omitempty"`
	Value       float64 `json:"value,omitempty"`
	Description string  `json:"description,omitempty"`
}

// TaxSettings holds store tax settings.
type TaxSettings struct {
	AutomaticTaxEnabled *bool `json:"automaticTaxEnabled,omitempty"`
	Taxes               []Tax `json:"taxes,omitempty"`
}

// Tax represents a tax configuration.
type Tax struct {
	ID                 int64           `json:"id"`
	Name               string          `json:"name,omitempty"`
	Enabled            *bool           `json:"enabled,omitempty"`
	IncludeInPrice     *bool           `json:"includeInPrice,omitempty"`
	UseShippingAddress *bool           `json:"useShippingAddress,omitempty"`
	TaxShipping        *bool           `json:"taxShipping,omitempty"`
	AppliedByDefault   *bool           `json:"appliedByDefault,omitempty"`
	DefaultTax         float64         `json:"defaultTax,omitempty"`
	Rules              json.RawMessage `json:"rules,omitempty"`
}

// Zone represents a destination zone.
type Zone struct {
	ID                   string          `json:"id,omitempty"`
	Name                 string          `json:"name,omitempty"`
	CountryCodes         []string        `json:"countryCodes,omitempty"`
	StateOrProvinceCodes []string        `json:"stateOrProvinceCodes,omitempty"`
	Postcode             json.RawMessage `json:"postcode,omitempty"`
}

// BusinessRegistration holds business registration ID info.
type BusinessRegistration struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// LegalPagesSettings holds legal page configurations.
type LegalPagesSettings struct {
	RequireTermsAgreementAtCheckout *bool       `json:"requireTermsAgreementAtCheckout,omitempty"`
	LegalPages                      []LegalPage `json:"legalPages,omitempty"`
}

// LegalPage represents a single legal page.
type LegalPage struct {
	Type        string `json:"type,omitempty"`
	Enabled     *bool  `json:"enabled,omitempty"`
	Title       string `json:"title,omitempty"`
	Display     string `json:"display,omitempty"`
	Text        string `json:"text,omitempty"`
	ExternalURL string `json:"externalUrl,omitempty"`
}

// UpdateRequest holds fields for updating the store profile.
// All fields are optional — only set fields are sent.
type UpdateRequest struct {
	GeneralInfo     *GeneralInfo     `json:"generalInfo,omitempty"`
	Settings        *Settings        `json:"settings,omitempty"`
	MailSettings    *MailSettings    `json:"mailNotifications,omitempty"`
	Company         *Company         `json:"company,omitempty"`
	FormatsAndUnits *FormatsAndUnits `json:"formatsAndUnits,omitempty"`
	Languages       *Languages       `json:"languages,omitempty"`
	TaxSettings     *TaxSettings     `json:"taxSettings,omitempty"`
}

// UpdateResult represents the response from an update operation.
type UpdateResult struct {
	UpdateCount int `json:"updateCount"`
}
