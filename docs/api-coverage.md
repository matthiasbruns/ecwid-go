# Ecwid API Coverage Report

> Auto-generated comparison of `ecwid-go` implementation against the [Ecwid REST API Reference](https://docs.ecwid.com/api-reference).

## Overview

The client exposes **11 services** via `ecwid.Client`:

| Service | Package | Coverage |
|---|---|---|
| Profile | `ecwid/profile` | Partial |
| Reports | `ecwid/reports` | Partial |
| Categories | `ecwid/categories` | Partial |
| Customers | `ecwid/customers` | Partial |
| Discounts | `ecwid/discounts` | Complete |
| Domains | `ecwid/domains` | Complete |
| Dictionaries | `ecwid/dictionaries` | Complete |
| Staff | `ecwid/staff` | Partial |
| Reviews | `ecwid/reviews` | Partial |
| Carts | `ecwid/carts` | Complete (types incomplete) |
| Subscriptions | `ecwid/subscriptions` | Complete (types incomplete) |

**Entirely missing API areas:** Orders (core), Products, Application, Batch Requests, Shipping Options, Payment Options, Checkout Extra Fields, Storefront Widget Details, Order Invoices, Order Statuses, Order Extra Fields.

---

## Previously Reported Bugs (now resolved)

| Issue | Status | Location |
|---|---|---|
| ~~Subscriptions `SearchOptions` used `chargeFrom`/`chargeTo` but API expects `nextChargeFrom`/`nextChargeTo`~~ | ✅ Fixed | `ecwid/subscriptions/service.go` |
| ~~Carts `PlaceResult` missing `id`, `vendorOrderNumber`, `cartId` fields from API response~~ | ✅ Fixed | `ecwid/carts/types.go` |
| ~~Subscriptions `Get`/`Update` methods lack zero-value validation for `subscriptionID`~~ | ✅ Fixed | `ecwid/subscriptions/service.go` |

---

## Detailed Coverage by Section

### Store Profile (`ecwid/profile`)

| Endpoint | Method | Path | Status |
|---|---|---|---|
| Get store profile | GET | `/profile` | ✅ Implemented (many sub-struct fields missing) |
| Update store profile | PUT | `/profile` | ✅ Implemented (several updatable sections missing from request type) |
| Upload store logo image | PUT | `/profile/{logo}` | ❌ Missing |
| Delete store logo image | DELETE | `/profile/{logo}` | ❌ Missing |

**Missing fields in Profile types:**
- `GeneralInfo`: `websitePlatform`, `profileId`
- `Account`: `brandName`, `supportEmail`, `suspended`, `registrationDate`, `paid`, `limitsAndRestrictions`, and more
- `Settings`: 30+ fields missing including `storeDescriptionTranslated`, `recurringSubscriptionsSettings`, SEO fields, tracking pixels, etc.
- `TaxSettings`: `euIossEnabled`, `ukVatRegistered`, `pricesIncludeTax`, and more
- Top-level: `taxes`, `phoneNotifications`, `orderInvoiceSettings`, `socialLinksSettings`, `tipsSettings`
- `UpdateRequest`: missing `account`, `shipping`, `zones`, `taxes`, `designSettings`, and more

---

### Store Reports (`ecwid/reports`)

| Endpoint | Method | Path | Status |
|---|---|---|---|
| Get store reports | GET | `/reports/{reportType}` | ✅ Implemented (missing `comparePeriod` response fields) |
| Get latest store update stats | GET | `/latest-stats` | ✅ Complete |
| Get deleted items history | GET | `/{entity}/deleted` | ❌ Missing |

---

### Orders — ❌ NOT IMPLEMENTED

No `orders` package exists. **15 endpoints missing:**

| Endpoint | Method | Path |
|---|---|---|
| Search orders | GET | `/orders` |
| Get order | GET | `/orders/{id}` |
| Get last order | GET | `/orders/lastOrder` |
| Calculate order details | POST | `/orders/calculate` |
| Update order | PUT | `/orders/{id}` |
| Create order | POST | `/orders` |
| Delete order | DELETE | `/orders/{id}` |
| Get repeat order URL | GET | `/orders/{id}/repeat` |
| Get order receipt PDF | GET | `/orders/{id}/receipt` |
| Get tax invoices for order | GET | `/orders/{id}/taxInvoices` |
| Generate tax invoice for order | POST | `/orders/{id}/taxInvoices` |
| Search order statuses | GET | `/orders/statuses` |
| Get order status | GET | `/orders/statuses/{id}` |
| Update custom order status | PUT | `/orders/statuses/{id}` |
| Search order extra fields | GET | `/orders/extraFields` |
| Update order extra field | PUT | `/orders/extraFields/{id}` |
| Add extra fields to order | POST | `/orders/extraFields` |
| Delete order extra field | DELETE | `/orders/extraFields/{id}` |

> Note: Abandoned carts and recurring subscriptions (sub-sections of Orders in the docs) ARE implemented via `carts` and `subscriptions` packages.

---

### Products — ❌ NOT IMPLEMENTED (except Reviews)

No `products` package exists. **45 endpoints missing:**

**Core (10):** Search, Get, Update, Create, Delete, Delete all, Adjust stock, Get filters, Get swatches, Search brands

**Product Images and Videos (14):** Upload/delete main image (sync+async), Upload/delete gallery image (sync+async), Delete all gallery images, Upload/delete main video, Upload/delete gallery video, Upload cover for gallery video, Download gallery video, Bulk update

**Product Files (5):** Upload, Download, Delete, Delete all, Change description

**Product Variations (10):** Search, Get, Update, Create, Delete, Delete all, Adjust stock, Upload image (sync+async), Delete image

**Product Types and Attributes (5):** Search, Get, Update, Create, Delete

**Product Reviews (1):** Get product reviews stats (`GET /reviews/filters_data`)

> Note: 4 of 5 product review endpoints ARE implemented via the `reviews` package.

---

### Categories (`ecwid/categories`)

| Endpoint | Method | Path | Status |
|---|---|---|---|
| Search categories | GET | `/categories` | ✅ Implemented |
| Get category | GET | `/categories/{id}` | ✅ Implemented (missing optional query params: `baseUrl`, `cleanURLs`, `slugsWithoutIds`, `lang`) |
| Create category | POST | `/categories` | ✅ Implemented |
| Update category | PUT | `/categories/{id}` | ✅ Implemented |
| Delete category | DELETE | `/categories/{id}` | ✅ Implemented |
| Search categories by path | GET | `/categories` (path variant) | ❌ Missing |
| Upload category image | POST | `/categories/{id}/image` | ❌ Missing |
| Upload category image (async) | POST | `/categories/{id}/image/async` | ❌ Missing |
| Delete category image | DELETE | `/categories/{id}/image` | ❌ Missing |
| Get order of categories | GET | `/categories/sort` | ❌ Missing |
| Update order of categories | PUT | `/categories/sort` | ❌ Missing |
| Get order of products | GET | `/products/sort` | ⚠️ Partial (read only) |
| Update order of products | PUT | `/products/sort` | ❌ Missing |
| Assign products to category | POST | | ❌ Missing |
| Unassign products from category | DELETE | | ❌ Missing |

**Missing Category struct fields:** `imageExternalId`, `origin`, `seoTitle`, `seoTitleTranslated`, `seoDescription`, `seoDecriptionTranslated`, `alt`, `externalReferenceId`

---

### Customers (`ecwid/customers`)

| Endpoint | Method | Path | Status |
|---|---|---|---|
| Search customers | GET | `/customers` | ✅ Implemented |
| Get customer | GET | `/customers/{id}` | ✅ Implemented |
| Create customer | POST | `/customers` | ✅ Implemented |
| Update customer | PUT | `/customers/{id}` | ✅ Implemented |
| Delete customer | DELETE | `/customers/{id}` | ✅ Implemented |
| Search customer groups | GET | `/customer_groups` | ❌ Missing |
| Get customer group | GET | `/customer_groups/{id}` | ❌ Missing |
| Create customer group | POST | `/customer_groups` | ❌ Missing |
| Update customer group | PUT | `/customer_groups/{id}` | ❌ Missing |
| Delete customer group | DELETE | `/customer_groups/{id}` | ❌ Missing |
| Search customer contacts | GET | `/customers/{id}/contacts` | ❌ Missing |
| Get customer contact | GET | `/customers/{id}/contacts/{contactId}` | ❌ Missing |
| Create customer contact | POST | `/customers/{id}/contacts` | ❌ Missing |
| Update customer contact | PUT | `/customers/{id}/contacts/{contactId}` | ❌ Missing |
| Delete customer contact | DELETE | `/customers/{id}/contacts/{contactId}` | ❌ Missing |
| Search customer extra fields | GET | `/customer_extra_fields` | ❌ Missing |
| Get customer extra field | GET | `/customer_extra_fields/{id}` | ❌ Missing |
| Create customer extra field | POST | `/customer_extra_fields` | ❌ Missing |
| Update customer extra field | PUT | `/customer_extra_fields/{id}` | ❌ Missing |
| Delete customer extra field | DELETE | `/customer_extra_fields/{id}` | ❌ Missing |

**Missing Customer struct fields:** `fiscalCode`, `electronicInvoicePecEmail`, `electronicInvoiceSdiCode`, `password` (write-only)

**Weak typing:** `billingPerson`, `shippingAddresses`, `contacts`, `stats`, `favorites` are `*json.RawMessage` instead of typed structs.

---

### Discounts (`ecwid/discounts`)

| Endpoint | Method | Path | Status |
|---|---|---|---|
| Search promotions | GET | `/promotions` | ✅ Implemented |
| Create promotion | POST | `/promotions` | ✅ Implemented |
| Update promotion | PUT | `/promotions/{id}` | ✅ Implemented |
| Delete promotion | DELETE | `/promotions/{id}` | ✅ Implemented |
| Search discount coupons | GET | `/discount_coupons` | ✅ Implemented |
| Get discount coupon | GET | `/discount_coupons/{id}` | ✅ Implemented |
| Create discount coupon | POST | `/discount_coupons` | ✅ Implemented |
| Update discount coupon | PUT | `/discount_coupons/{id}` | ✅ Implemented |
| Delete discount coupon | DELETE | `/discount_coupons/{id}` | ✅ Implemented |

**Type note:** `Promotion.Triggers`, `Promotion.Targets`, `Coupon.CatalogLimit`, `Coupon.ShippingLimit` use `*json.RawMessage` instead of typed structs.

---

### Domains (`ecwid/domains`)

| Endpoint | Method | Path | Status |
|---|---|---|---|
| Get domains | GET | `/domains` | ✅ Implemented |
| Purchase domain | POST | `/domains/purchase` | ✅ Implemented |

**Complete.**

---

### Dictionaries (`ecwid/dictionaries`)

| Endpoint | Method | Path | Status |
|---|---|---|---|
| Country codes | GET | `/countries` | ✅ Implemented |
| Currency codes | GET | `/currencies` | ✅ Implemented |
| Currency codes by country | GET | `/currencyByCountry` | ✅ Implemented |
| State codes by country | GET | `/states` | ✅ Implemented |
| Tax classes by country | GET | `/taxClasses` | ✅ Implemented |

**Complete.** Minor: `TaxClass.Localization` uses `map[string]string` instead of a typed struct.

---

### Staff Accounts (`ecwid/staff`)

| Endpoint | Method | Path | Status |
|---|---|---|---|
| Search staff accounts | GET | `/staff` | ✅ Implemented |
| Get staff account | GET | `/staff/{id}` | ✅ Implemented |
| Create staff account | POST | `/staff` | ✅ Implemented |
| Update staff account | PUT | `/staff/{id}` | ✅ Implemented |
| Delete staff account | DELETE | `/staff/{id}` | ✅ Implemented |
| Get staff account scopes | GET | `/profile/staffScopes` | ❌ Missing |
| Resend staff account invite | POST | `/staff/invite` | ❌ Missing |
| Cancel staff account invite | DELETE | `/staff/invite` | ❌ Missing |

---

### Product Reviews (`ecwid/reviews`)

| Endpoint | Method | Path | Status |
|---|---|---|---|
| Search product reviews | GET | `/reviews` | ✅ Implemented |
| Update product review status | PUT | `/reviews/{id}` | ✅ Implemented |
| Delete product review | DELETE | `/reviews/{id}` | ✅ Implemented |
| Bulk update/delete product reviews | PUT | `/reviews/mass_update` | ✅ Implemented |
| Get product reviews stats | GET | `/reviews/filters_data` | ❌ Missing |

---

### Abandoned Carts (`ecwid/carts`)

| Endpoint | Method | Path | Status |
|---|---|---|---|
| Search abandoned carts | GET | `/carts` | ✅ Implemented |
| Get abandoned cart | GET | `/carts/{id}` | ✅ Implemented |
| Update abandoned cart | PUT | `/carts/{id}` | ✅ Implemented |
| Convert abandoned cart to order | POST | `/carts/{id}/place` | ✅ Implemented |

**All endpoints implemented**, but types are significantly incomplete:
- `Cart` struct missing many fields (`customerTaxExempt`, `customSurcharges`, `couponDiscount`, `handlingFee`, `predictedPackage`, etc.)
- `Cart` uses `json.RawMessage` for `Items`, `BillingPerson`, `ShippingPerson`, `ShippingOption`, `DiscountCoupon`, `DiscountInfo`
- `SearchOptions` missing `showHidden`, `email` query params
- `UpdateRequest` only has `Hidden` — missing `taxesOnShipping`, `b2b_b2c`, and more
- ~~`PlaceResult` missing `id`, `vendorOrderNumber`, `cartId` fields~~ (fixed)

---

### Recurring Subscriptions (`ecwid/subscriptions`)

| Endpoint | Method | Path | Status |
|---|---|---|---|
| Search recurring subscriptions | GET | `/subscriptions` | ✅ Implemented |
| Get recurring subscription | GET | `/subscriptions/{id}` | ✅ Implemented |
| Update recurring subscription | PUT | `/subscriptions/{id}` | ✅ Implemented |

**All endpoints implemented**, but:
- ~~`SearchOptions` used `chargeFrom`/`chargeTo` query param names~~ (fixed — now uses `nextChargeFrom`/`nextChargeTo`)
- `SearchOptions` missing: `recurringIntervalCount`, `email`, `orderId`, `orderTotal`, `orderCreatedFrom`, `orderCreatedTo`, `sortBy`

---

## Entirely Unimplemented API Areas

| API Section | Endpoints | Priority |
|---|---|---|
| **Orders** (core CRUD) | 8 endpoints | HIGH |
| **Products** (core CRUD + stock) | 10 endpoints | HIGH |
| **Product Images/Videos** | 14 endpoints | MEDIUM |
| **Product Variations** | 10 endpoints | MEDIUM |
| **Product Files** | 5 endpoints | LOW |
| **Product Types/Attributes** | 5 endpoints | LOW |
| **Order Invoices** | 3 endpoints | MEDIUM |
| **Order Statuses** | 3 endpoints | MEDIUM |
| **Order Extra Fields** | 4 endpoints | MEDIUM |
| **Customer Groups** | 5 endpoints | MEDIUM |
| **Customer Contacts** | 5 endpoints | LOW |
| **Customer Extra Fields** | 5 endpoints | LOW |
| **Application (App Storage/Billing)** | 9 endpoints | LOW |
| **Shipping Options** | 4 endpoints | MEDIUM |
| **Payment Options** | 5 endpoints | MEDIUM |
| **Checkout Extra Fields** | 7 endpoints | LOW |
| **Storefront Widget Details** | 14 endpoints | LOW |
| **Batch Requests** | 1 endpoint | LOW |
| **Category Images** | 3 endpoints | LOW |
| **Category Ordering** | 4 endpoints | LOW |

---

## Summary Statistics

| Metric | Count |
|---|---|
| **Total API endpoints (documented)** | ~185 |
| **Endpoints implemented** | ~46 |
| **Endpoints missing** | ~139 |
| **Coverage** | ~25% |
| **Services fully complete** | 3 (Domains, Dictionaries, Discounts) |
| **Services partially complete** | 8 |
| **API areas with zero implementation** | 11+ |
| **Bugs found (all resolved)** | 3 |
