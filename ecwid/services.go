package ecwid

// ProductService handles product-related API endpoints.
// Endpoints: /products, /products/{id}, variations, inventory, images, gallery, files, sort, filters, brands, classes, swatches.
type ProductService struct {
	client *Client
}

// OrderService handles order-related API endpoints.
// Endpoints: /orders, /orders/{id}, /orders/last, /orders/deleted, extra fields, invoices, calculate.
type OrderService struct {
	client *Client
}

// CategoryService handles category-related API endpoints.
// Endpoints: /categories, /categoriesByPath, /categories/{id}, sort, images, assign/unassign products.
type CategoryService struct {
	client *Client
}

// CustomerService handles customer-related API endpoints.
// Endpoints: /customers, /customers/{id}, contacts, extra fields, customer groups, deleted.
type CustomerService struct {
	client *Client
}

// CartService handles abandoned cart API endpoints.
// Endpoints: /carts, /carts/{id}, /carts/{id}/place.
type CartService struct {
	client *Client
}

// SubscriptionService handles recurring subscription API endpoints.
// Endpoints: /subscriptions, /subscriptions/{id}.
type SubscriptionService struct {
	client *Client
}

// PromotionService handles promotion API endpoints.
// Endpoints: /promotions, /promotions/{id}.
type PromotionService struct {
	client *Client
}

// CouponService handles discount coupon API endpoints.
// Endpoints: /discount_coupons, /discount_coupons/{id}, /coupons/deleted.
type CouponService struct {
	client *Client
}

// ProfileService handles store profile API endpoints.
// Endpoints: /profile, staffScopes, order_statuses, extrafields, logos, shipping/payment options.
type ProfileService struct {
	client *Client
}

// ReviewService handles product review API endpoints.
// Endpoints: /reviews, /reviews/filters_data, /reviews/deleted, /reviews/{id}, mass_update.
type ReviewService struct {
	client *Client
}

// StaffService handles staff account API endpoints.
// Endpoints: /staff, /staff/{id}.
type StaffService struct {
	client *Client
}

// DomainService handles domain API endpoints.
// Endpoints: /domains, /domains/search, /domains/purchase, verification email, reset password.
type DomainService struct {
	client *Client
}

// DictionaryService handles dictionary API endpoints (read-only).
// Endpoints: /countries, /currencies, /currencyByCountry, /states, /taxClasses.
type DictionaryService struct {
	client *Client
}

// ReportService handles store report and stats API endpoints.
// Endpoints: /reports/{type}, /latest-stats.
type ReportService struct {
	client *Client
}
