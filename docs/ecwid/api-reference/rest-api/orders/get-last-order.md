# Get last order

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/orders/last`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
GET /api/v3/1003/orders HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
```

Response:

```json
{
    "id": "K8XTQ",
    "internalId": 516483487,
    "refundedAmount": 0,
    "subtotal": 10,
    "subtotalWithoutTax": 9.09,
    "total": 10,
    "totalWithoutTax": 9.09,
    "giftCardRedemption": 0,
    "totalBeforeGiftCardRedemption": 10,
    "giftCardDoubleSpending": false,
    "email": "ec.apps@lightspeedhq.com",
    "paymentMethod": "Offline",
    "tax": 0.91,
    "customerTaxExempt": false,
    "customerTaxId": "",
    "customerTaxIdValid": false,
    "b2b_b2c": "b2c",
    "reversedTaxApplied": false,
    "customerRequestedInvoice": false,
    "customerFiscalCode": "",
    "electronicInvoicePecEmail": "",
    "electronicInvoiceSdiCode": "",
    "ipAddress": "80.241.251.10",
    "couponDiscount": 0,
    "paymentStatus": "PAID",
    "fulfillmentStatus": "DELIVERED",
    "orderNumber": 516483487,
    "vendorOrderNumber": "K8XTQ",
    "publicUid": "K8XTQ",
    "refererUrl": "https://store1003.company.site/products/",
    "orderComments": "",
    "volumeDiscount": 0,
    "customerId": 270919557,
    "membershipBasedDiscount": 0,
    "totalAndMembershipBasedDiscount": 0,
    "customDiscount": [],
    "customSurcharges": [],
    "discount": 0,
    "usdTotal": 11.067309160854485,
    "globalReferer": "https://my.ecwid.com/",
    "createDate": "2024-09-04 07:22:09 +0000",
    "updateDate": "2024-09-04 07:26:24 +0000",
    "createTimestamp": 1725434529,
    "updateTimestamp": 1725434784,
    "items": [
        {
            "id": 1828115520,
            "productId": 692730761,
            "categoryId": 0,
            "price": 10,
            "priceWithoutTax": 9.09,
            "productPrice": 10,
            "sku": "123123",
            "quantity": 1,
            "shortDescription": "",
            "shortDescriptionTranslated": {
                "cs": "",
                "en": ""
            },
            "tax": 0.91,
            "shipping": 0,
            "quantityInStock": 0,
            "name": "Toys",
            "nameTranslated": {
                "cs": "",
                "en": "Toys"
            },
            "isShippingRequired": false,
            "trackQuantity": false,
            "fixedShippingRateOnly": false,
            "imageUrl": "https://d2j6dbq0eux0bg.cloudfront.net/images/1003/4519591921.jpg",
            "smallThumbnailUrl": "https://d2j6dbq0eux0bg.cloudfront.net/images/1003/4519591920.jpg",
            "hdThumbnailUrl": "https://d2j6dbq0eux0bg.cloudfront.net/images/1003/4519591922.jpg",
            "fixedShippingRate": 0,
            "digital": true,
            "productAvailable": true,
            "couponApplied": false,
            "files": [
                {
                    "productFileId": 96178524,
                    "maxDownloads": 0,
                    "remainingDownloads": 0,
                    "expire": "2024-09-07 07:23:43 +0000",
                    "name": "images.jpeg",
                    "description": "",
                    "size": 12006,
                    "adminUrl": "https://app.ecwid.com/api/v3/1003/products/692730761/files/96178524",
                    "customerUrl": "https://app.ecwid.com/download/1003/71248dca454148b792e5975c87793a872aec44b184c246d7b953ee2e08990703/images.jpeg"
                }
            ],
            "taxes": [
                {
                    "name": "10% Tax",
                    "value": 10,
                    "total": 0.91,
                    "taxOnDiscountedSubtotal": 0.91,
                    "taxOnShipping": 0,
                    "includeInPrice": true,
                    "sourceTaxRateId": 947976181,
                    "sourceTaxRateType": "MANUAL",
                    "taxClassName": "Standard rate",
                    "taxClassCode": "default"
                }
            ],
            "dimensions": {
                "length": 0,
                "width": 0,
                "height": 0
            },
            "discounts": [],
            "discountsAllowed": true,
            "taxable": true,
            "giftCard": false,
            "recurringTaxIds": [],
            "isCustomerSetPrice": false,
            "attributes": []
        }
    ],
    "refunds": [],
    "billingPerson": {
        "name": "API Team",
        "firstName": "Support",
        "lastName": "Team",
        "companyName": "Ecwid API Team",
        "street": "Anhalter Strasse 98",
        "city": "Niederhausen An Der Appel",
        "countryCode": "ES",
        "countryName": "Spain",
        "postalCode": "08003",
        "stateOrProvinceCode": "AL",
        "stateOrProvinceName": "Alava",
        "phone": "0123456789"
    },
    "predictedPackage": [],
    "shippingLabelAvailableForShipment": false,
    "shipments": [],
    "additionalInfo": {
        "google_customer_id": "1904890913.1707375888"
    },
    "paymentParams": {},
    "extraFields": {},
    "ticket": -455011247,
    "hidden": false,
    "taxesOnShipping": [],
    "disableAllCustomerNotifications": false,
    "externalFulfillment": false,
    "utmDataSets": [],
    "invoices": [
        {
            "internalId": 117352011,
            "id": "27",
            "created": "2024-09-04 07:22:09 +0000",
            "link": "https://app.ecwid.com/download_tax_invoice?ownerid=1003&invoice_id=117352011&access_key=z7clyftZJZj",
            "type": "SALE"
        }
    ],
    "pricesIncludeTax": true,
    "loyalty": {
        "customerEnabled": false,
        "earned": 0,
        "balance": 0
    },
    "customerUserAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36"
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_orders`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

All query params are optional.

<table><thead><tr><th width="165">Param</th><th width="124">Type</th><th>Description</th></tr></thead><tbody><tr><td>responseFields</td><td>string</td><td>Limit JSON response by specific fields. If specified, all missing fields will be removed from the response body.<br>For example: <code>?responseFields=id,email,total</code></td></tr></tbody></table>

Example of using `responseFields` param:

{% tabs %}
{% tab title="Request" %}

```
curl --location 'https://app.ecwid.com/api/v3/1003/orders/last?responseFields=id,email,total' \
--header 'Authorization: Bearer secret_ab***cd'
```

{% endtab %}

{% tab title="Response" %}

```json
{
    "id": "K8XTQ",
    "email": "ec.apps@lightspeedhq.com",
    "total": 10
}
```

{% endtab %}
{% endtabs %}

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

<table><thead><tr><th width="225">Field</th><th width="189">Type</th><th>Description</th></tr></thead><tbody><tr><td>id</td><td>string</td><td>Unique order identificator with prefix and suffix defined by the store admin. For example, order ID <code>MYSTORE-X8UYE</code> contains <code>MYSTORE-</code> prefix.<br><br>Order ID is shown to customers in any notifications and to the store owner in Ecwid admin and notifications.</td></tr><tr><td>subtotal</td><td>number</td><td>Cost of all products in the order (item's <code>price</code> x <code>quantity</code>) before any cost modifiers such as discounts, taxes, fees, etc. are applied.</td></tr><tr><td>subtotalWithoutTax</td><td>number</td><td>Order subtotal without taxes included in price (GROSS) when <code>pricesIncludeTax</code> is <code>true</code>.</td></tr><tr><td>total</td><td>number</td><td>Order total cost with all cost modifiers: shipping costs, taxes, fees, and discounts.</td></tr><tr><td>totalWithoutTax</td><td>number</td><td>Order total without taxes. Calculates as <code>total</code> - <code>tax</code>.</td></tr><tr><td>refundedAmount</td><td>number</td><td>Sum of all refunds applied to the order.</td></tr><tr><td>giftCardRedemption</td><td>number</td><td>Amount deducted from the <a href="https://support.ecwid.com/hc/en-us/articles/360002011419">Gift Card</a> balance and applied to order total.</td></tr><tr><td>totalBeforeGiftCardRedemption</td><td>number</td><td>Order total before the Gift Card was applied.</td></tr><tr><td>email</td><td>string</td><td>Customer's email address.</td></tr><tr><td>paymentModule</td><td>string</td><td>Payment processor used to pay for the order online.<br><br>Only available to online payment integrations build by Ecwid team.</td></tr><tr><td>paymentMethod</td><td>string</td><td>Name of the payment method customer chosen at the checkout.</td></tr><tr><td>tax</td><td>number</td><td>Sum of all taxes applied to products and shipping.<br><br>If the order is modified after being placed, this value is <strong>not</strong> recalculated automatically.</td></tr><tr><td>customerTaxExempt</td><td>boolean</td><td>Defines if the customer is tax exempt. Requires valid tax ID.<br><br>Read more about handling tax exempt customers in <a href="https://support.ecwid.com/hc/en-us/articles/213823045-Handling-tax-exempt-customers">Help Center</a>.</td></tr><tr><td>customerTaxId</td><td>string</td><td>Tax ID entered by the customer at the checkout.</td></tr><tr><td>customerTaxIdValid</td><td>boolean</td><td>Defines if customer's tax ID is valid for tax exemption.</td></tr><tr><td>reversedTaxApplied</td><td>boolean</td><td>Defines if order tax was reversed (set to 0). Requires valid tax ID.</td></tr><tr><td>b2b_b2c</td><td>string</td><td><p>Order type. One of:</p><p><code>b2b</code> - business-to-business</p><p><code>b2c</code> - business-to-consumer </p></td></tr><tr><td>customerRequestedInvoice</td><td>boolean</td><td>Defines if customer requested an invoice.</td></tr><tr><td>customerFiscalCode</td><td>string</td><td>Fiscale code of the customer.</td></tr><tr><td>electronicInvoicePecEmail</td><td>string</td><td>PEC email for order invoices.</td></tr><tr><td>electronicInvoiceSdiCode</td><td>string</td><td>SDI code for order invoices.</td></tr><tr><td>ipAddress</td><td>string</td><td>Customer's IP address detected at the checkout.</td></tr><tr><td>paymentStatus</td><td>string</td><td>Order payment status. Supported values: <code>AWAITING_PAYMENT</code>, <code>PAID</code>, <code>CANCELLED</code>, <code>REFUNDED</code>, <code>PARTIALLY_REFUNDED</code>, <code>INCOMPLETE</code>, <code>CUSTOM_PAYMENT_STATUS_1</code>, <code>CUSTOM_PAYMENT_STATUS_2</code>, <code>CUSTOM_PAYMENT_STATUS_3</code>.<br><br>Read more about order statuses in <a href="https://support.ecwid.com/hc/en-us/articles/207806235-Order-details-and-statuses-overview#-understanding-order-statuses"><strong>Help Center</strong></a>.</td></tr><tr><td>fulfillmentStatus</td><td>string</td><td>Order fulfillment status. Supported values: <code>AWAITING_PROCESSING</code>, <code>PROCESSING</code>, <code>SHIPPED</code>, <code>DELIVERED</code>, <code>WILL_NOT_DELIVER</code>, <code>RETURNED</code>, <code>READY_FOR_PICKUP</code>, <code>OUT_FOR_DELIVERY</code>, <code>CUSTOM_FULFILLMENT_STATUS_1</code>, <code>CUSTOM_FULFILLMENT_STATUS_2</code>, <code>CUSTOM_FULFILLMENT_STATUS_3</code>.<br><br>Read more about order statuses in <a href="https://support.ecwid.com/hc/en-us/articles/207806235-Order-details-and-statuses-overview#-understanding-order-statuses"><strong>Help Center</strong></a>.</td></tr><tr><td>refererUrl</td><td>string</td><td>URL of the page when order was placed without page slugs (hash <code>#</code> part).</td></tr><tr><td>orderComments</td><td>string</td><td>Order comments, left by a customer at the checkout.</td></tr><tr><td>volumeDiscount</td><td>number</td><td>Sum of applied <strong>promotions</strong> based on subtotal. Included in the <code>discount</code> field.</td></tr><tr><td>membershipBasedDiscount</td><td>number</td><td>Sum of applied <strong>promotions</strong> based on customer group. Included in the <code>discount</code> field.</td></tr><tr><td>totalAndMembershipBasedDiscount</td><td>number</td><td>Sum of applied <strong>promotions</strong> based on both subtotal and customer group. Included in the <code>discount</code> field.</td></tr><tr><td>customDiscount</td><td>array of numbers</td><td>List of absolute discounts added by applications.</td></tr><tr><td>discount</td><td>number</td><td>Total order discount. Includes both promotions and discount coupons. <br><br>Calculated as the sum of the<code>couponDiscount</code> and <code>totalAndMembershipBasedDiscount</code> fields.</td></tr><tr><td>couponDiscount</td><td>number</td><td><p>Discount value from applied <strong>discount coupon</strong>, e.g. <code>10</code>.</p><p>Total order discount is the sum of the<code>couponDiscount</code> and <code>discount</code> fields.</p></td></tr><tr><td>discountInfo</td><td>array <a href="#discounts">discounts</a></td><td>Detailed information about applied <strong>promotions</strong>.</td></tr><tr><td>discountCoupon</td><td>object <a href="#discountcoupon">discountCoupon</a></td><td>Detailed information about applied <strong>discount coupons</strong>.</td></tr><tr><td>customerId</td><td>number</td><td>Unique internal ID assigned to the customer.</td></tr><tr><td>customSurcharges</td><td>array of objects <a href="#customsurcharges">customSurcharges</a></td><td>Information about surcharges applied to the order.</td></tr><tr><td>usdTotal</td><td>number</td><td>Order total converted from the store's currency to USD.</td></tr><tr><td>globalReferer</td><td>string</td><td>URL that the customer came to the store from</td></tr><tr><td>createDate</td><td>string</td><td>The datetime when the order was placed, for example <code>2014-06-06 18:57:19 +0000</code></td></tr><tr><td>updateDate</td><td>string</td><td>The datetime of the latest order update. This includes all changes made from Ecwid admin or API. For example, <code>2014-06-06 18:57:19 +0000</code></td></tr><tr><td>createTimestamp</td><td>number</td><td>The datetime when the order was placed in UNIX timestamp, for example <code>1427268654</code></td></tr><tr><td>updateTimestamp</td><td>number</td><td>The datetime of the latest order update in UNIX timestamp. This includes all changes made from Ecwid admin or API. For example, <code>1427268654</code></td></tr><tr><td>customerGroup</td><td>string</td><td>Name of the group the customer belongs to (if any).<br><br>Read more about <a href="https://support.ecwid.com/hc/en-us/articles/207807105-Customer-groups">customer groups</a>.</td></tr><tr><td>customerGroupId</td><td>number</td><td>ID of the group the customer belongs to.</td></tr><tr><td>items</td><td>array <a href="#items">items</a></td><td>Detailed information about products in the order.</td></tr><tr><td>refunds</td><td>array <a href="#refunds">refunds</a></td><td>Details about refunds made to order.</td></tr><tr><td>shippingPerson</td><td>object <a href="#shippingperson">shippingPerson</a></td><td>Name and shipping address details left by customer at the checkout.</td></tr><tr><td>billingPerson</td><td>object <a href="#billingperson">billingPerson</a></td><td>Name and billing address details left by customer at the checkout.</td></tr><tr><td>shippingOption</td><td>object <a href="#shippingoption">shippingOption</a></td><td>Details about the shipping option customer selected at the checkout.</td></tr><tr><td>handlingFee</td><td>object <a href="#handlingfee">handlingFee</a></td><td>Details about fees applied to order.</td></tr><tr><td>predictedPackages</td><td>object <a href="#predictedpackages">predictedPackages</a></td><td>Minimum total dimensions and weight of a single shipping package that will be enough to carry all products added to the order.</td></tr><tr><td>shippingLabelAvailableForShipment</td><td>boolean</td><td>Defines if the store owner can buy a shipping label through Ecwid for the order.</td></tr><tr><td>shipments</td><td>array <a href="#shipments">shipments</a></td><td>Detailed information about purchased shipping label.</td></tr><tr><td>additionalInfo</td><td>object</td><td>Internal order information for Ecwid services.</td></tr><tr><td>paymentParams</td><td>object</td><td>Internal payment parameters for Ecwid services.</td></tr><tr><td>extraFields</td><td>object <a href="#extrafields">extraFields</a></td><td>Names and values of custom checkout fields applied to the order.</td></tr><tr><td>orderExtraFields</td><td>array <a href="#orderextrafields">orderExtraFields</a></td><td>Additional optional information about the order's extra fields. Along with the value of the field, it contains technical information, such as id, type, etc. of the field. Total storage of extra fields cannot exceed 8Kb.</td></tr><tr><td>hidden</td><td>boolean</td><td>Defines if the order is hidden from Ecwid admin. Applies to unsfinished orders only.</td></tr><tr><td>trackingNumber</td><td>string</td><td>Shipping tracking code.</td></tr><tr><td>paymentMessage</td><td>string</td><td>Error message sent by the online payment method. Only appears if a customer had issues with paying for the order online. When order becomes paid, <code>paymentMessage</code> is cleared</td></tr><tr><td>externalTransactionId</td><td>string</td><td>Transaction ID saved to the order details by the payment system. For example, PayPal transaction ID.</td></tr><tr><td>affiliateId</td><td>string</td><td>If a store has several storefronts, this ID is used to track from which one the order came from.<br><br>Read more on setting up affiliate IDs in <a href="https://support.ecwid.com/hc/en-us/articles/207100679-How-to-track-which-storefront-an-order-came-from#add-special-id-to-the-integration-code">Help Center</a>.</td></tr><tr><td>creditCardStatus</td><td>object <a href="#creditcardstatus">creditCardStatus</a></td><td>Saves verification messages if customer paid for the order with a credit card.</td></tr><tr><td>privateAdminNotes</td><td>string</td><td>Private note added to the order by store owner.</td></tr><tr><td>pickupTime</td><td>string</td><td>Order pickup time in the store date format (UTC +0 timezone), for example: <code>2017-10-17 05:00:00 +0000</code></td></tr><tr><td>taxesOnShipping</td><td>array <a href="#taxesonshipping">taxesOnShipping</a></td><td>Taxes applied to shipping 'as is'. <code>null</code> for old orders, <code>[]</code> for orders with taxes applied to subtotal only. Are not recalculated if order is updated later manually. Is calculated like: <code>(shippingRate + handlingFee)*(taxValue/100)</code></td></tr><tr><td>acceptMarketing</td><td>boolean</td><td>Defines if customer has accepted email marketing at the checkout. If <code>true</code> or <code>null</code>, you can use their email for promotions.</td></tr><tr><td>refererId</td><td>string</td><td>Referer identifier. Can be set in storefront via JS or by creating / updating an order with REST API</td></tr><tr><td>disableAllCustomerNotifications</td><td>boolean</td><td>Defines if the customer should receive any email notifications:<br> <code>true</code> - no notifications are sent to the customer. If <code>false</code> - email notifications are sent to customer according to store mail notification settings. <br><br>This setting does not affect email notifications to the store owner.</td></tr><tr><td>externalFulfillment</td><td>boolean</td><td><p>Defines if the order is fulfilled with an external system and should not be managed through Ecwid:<br><code>true</code> - Ecwid will hide fulfillment status change feature and ability to set tracking number within Ecwid admin.</p><p><code>false</code> - store owner can manage order fulfillment within Ecwid admin (default value)</p></td></tr><tr><td>externalOrderId</td><td>string</td><td>Order ID in an external system where order is fulfilled. <br><br>Requires <code>externalFulfillment</code>  to be <code>true</code>.</td></tr><tr><td>invoices</td><td>array <a href="#invoices">invoices</a></td><td>Tax invoices generated for the order. <br><br><strong>Read-only</strong></td></tr><tr><td>pricesIncludeTax</td><td>boolean</td><td><p>Defines if taxes are included to product prices (GROSS or NET prices):<br> <code>true</code> - the tax rate is included in product prices (GROSS). <br><code>false</code> - the tax rate is not included in product prices (NET).<br></p><p>Read more about setting up taxes in <a href="https://support.ecwid.com/hc/en-us/articles/207099899-Setting-up-taxes">Help Center</a>.</p></td></tr><tr><td>paymentSubtype</td><td>string</td><td>Internal field for Ecwid services.</td></tr><tr><td>utmData</td><td>array <a href="#utmdata">utmData</a></td><td>UTM tags saved for the order. <br><br>Read more about using UTM tags in orders in <a href="https://support.ecwid.com/hc/en-us/articles/4545287177372">Help Center</a>.</td></tr><tr><td>utmDataSets</td><td>array <a href="#utmdatasets">utmDataSets</a></td><td>Detailed information about UTM tags saved for the order. Contains more information than the <code>utmData</code> field.</td></tr><tr><td>lang</td><td>string</td><td><p>Defines a list of available languages or a single language for customer notifications. Must match one of the active store translations.<br></p><p>List of active store languages is available in the <mark style="color:green;"><code>GET</code></mark> <code>/profile</code>  request> <code>languages</code> > <code>enabledLanguages</code> field.</p></td></tr><tr><td>customerUserAgent</td><td>string</td><td>Details about the customer's device and platform used to place an order based on the <a href="https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/User-Agent">User-Agent</a> browser data.</td></tr><tr><td>externalOrderData</td><td>object <a data-mention href="#externalorderdata">#externalorderdata</a></td><td>Details for orders created or managed externally, for example, by other Lightspeed products.</td></tr></tbody></table>

#### items

| Field                      | Type                                                       | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| -------------------------- | ---------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| id                         | number                                                     | Order item ID unique for this order. Can be used to manage ordered items.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| productId                  | number                                                     | Internal product ID. Can be used to find full product details with the <mark style="color:green;">`GET`</mark> `/products` request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| subscriptionId             | number                                                     | ID of the subscription available at Ecwid admin > My Sales > Subscriptions.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| recurringChargeSettings    | object [recurringChargeSettings](#recurringchargesettings) | Details about subscription charge intervals.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| categoryId                 | number                                                     | <p>ID of the category this product belongs to or was added from. <br><br>Returns <code>-1</code> if the product was added to the cart via the Buy Now button.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| price                      | number                                                     | <p>Price of product in the order with some modifiers (doesn't include discount modifiers).<br>Calculation: <code>productPrice</code> + <code>wholesalePrices</code> + price modifiers from selected <code>options</code>.</p>                                                                                                                                                                                                                                                                                                                                                                                       |
| priceWithoutTax            | number                                                     | Price of product in the order without taxes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| productPrice               | number                                                     | Basic product price without any modifiers: options markups, discounts, taxes, fees.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| costPrice                  | number                                                     | Purchase price of the product in the specific order used for reports and profit calculations.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| weight                     | number                                                     | Weight of the product.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| sku                        | string                                                     | <p>Product SKU. <br><br>If the chosen options match a variation, this will be a variation SKU.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| quantity                   | number                                                     | Quantity of the product in the order.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| shortDescription           | string                                                     | Product description truncated to 120 characters.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| shortDescriptionTranslated | object [translations](#translations)                       | Available translations for product short description.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| tax                        | number                                                     | Total tax applied to the product.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| shipping                   | number                                                     | Partial shipping costs specific to the product.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| quantityInStock            | number                                                     | Number of products in stock in the store before placing the order.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| name                       | string                                                     | Name of the product.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| nameTranslated             | object [translations](#translations)                       | Available translations for the product name.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| isShippingRequired         | boolean                                                    | Defines if the product requires shipping.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| trackQuantity              | boolean                                                    | Defines if low stock notifications to the store owner are enabled.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| imageUrl                   | string                                                     | Link to the main product image.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| fixedShippingRateOnly      | boolean                                                    | <p>Defines if the product has a unique fixed shipping rate. <br><br>If <code>true</code>, shipping costs won't calculate for the product and <code>fixedShippingRate</code> value will be used instead.</p>                                                                                                                                                                                                                                                                                                                                                                                                         |
| fixedShippingRate          | number                                                     | <p>Fixed shipping costs for the product. <br><br>Affects shipping costs only if <code>fixedShippingRateOnly</code> is <code>true</code>.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| couponApplied              | boolean                                                    | Defines if the product has a discount coupon applied.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| selectedOptions            | array [selectedOptions](#selectedoptions)                  | Product options values selected by the customer at the checkout.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| taxes                      | array [taxes](#taxes)                                      | Detailed information about taxes applied to the product in this order.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| combinationId              | number                                                     | <p>ID of a product variation whos options mathes with values chosen by the customer at the checkout.<br><br>Read more on product variations in <a href="https://support.ecwid.com/hc/en-us/articles/207100299-Product-variations">Help Center.</a></p>                                                                                                                                                                                                                                                                                                                                                              |
| digital                    | boolean                                                    | <p>Defines if the product has any downloadable files attached.<br><br>Read more on digital products in <a href="https://support.ecwid.com/hc/en-us/articles/207100559-Digital-products">Help Center</a>.</p>                                                                                                                                                                                                                                                                                                                                                                                                        |
| files                      | array of objects [files](#files)                           | Details about downloadable files attached to the product.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| dimensions                 | object [dimensions](#dimensions)                           | Details about product dimensions used for shipping costs calculations.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| couponAmount               | number                                                     | <p>Discount applied to the product from discount coupon. </p><p><br>If the order is manually updated after being placed, this field is not recalculated automatically.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| discounts                  | array [discounts](#discounts)                              | **Promotions** applied to the specific product in the order.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| taxesOnShipping            | array [taxesOnShipping](#taxesonshipping)                  | <p>Taxes applied to shipping costs for the product with the calculation formula of: <code>(shippingRate + handlingFee)\*(taxValue/100)</code></p><p><br>If the order is manually updated after being placed, this field is not recalculated automatically.</p>                                                                                                                                                                                                                                                                                                                                                      |
| isCustomerSetPrice         | boolean                                                    | <p>If <code>true</code>, customer set a custom product price using the "<a href="https://support.ecwid.com/hc/en-us/articles/360018423259-Pay-What-You-Want-pricing">Pay What You Want</a>" feature. <br><br>In this case, both the product <code>price</code> and <code>selectedPrice</code> -> <code>value</code> fields contain the price set by a customer.<br>If <code>false</code>, customer didn't choose the custom price. Therefore, the <code>selectedPrice</code> -> <code>value</code> field will be absent and the <code>price</code> field contains default product price set by the store owner.</p> |
| selectedPrice              | object selectedPrice > value                               | <p>If <code>isCustomerSetPrice</code> is <code>true</code>, this field contains the "<a href="https://support.ecwid.com/hc/en-us/articles/360018423259-Pay-What-You-Want-pricing">Pay What You Want</a>" price set by a customer at the checkout.</p><p>Example with the PWYW price set to 100:<br><code>"selectedPrice": { "value": 100 }</code></p>                                                                                                                                                                                                                                                               |
| isPreorder                 | boolean                                                    | <p>Defines if the product was pre-ordered in this order.<br><br>Read more about accepting pre-orders in <a href="https://support.ecwid.com/hc/en-us/articles/5135873315100-Accepting-pre-orders-in-your-Ecwid-store">Help Center</a>.</p>                                                                                                                                                                                                                                                                                                                                                                           |
| attributes                 | array of objects [attributes](#attributes)                 | <p>Details about product attributes.<br><br>Read more on product attributes in <a href="https://support.ecwid.com/hc/en-us/articles/207807495-Product-types-and-attributes">Help Center</a>.</p>                                                                                                                                                                                                                                                                                                                                                                                                                    |

#### attributes

| Field | Type   | Description                    |
| ----- | ------ | ------------------------------ |
| name  | string | Name of the product attribute. |
| value | string | Attribute value.               |

#### taxes

| Field                   | Type   | Description                                                                                                                                                   |
| ----------------------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| name                    | string | Name of the tax visible to customers at the checkout and in order invoices.                                                                                   |
| value                   | number | Tax value in percent.                                                                                                                                         |
| total                   | number | Tax amount applied to the product.                                                                                                                            |
| taxOnDiscountedSubtotal | number | Tax applied to product price (`price`) after all discounts.                                                                                                   |
| taxOnShipping           | number | Tax applied to the shipping costs of the product.                                                                                                             |
| sourceTaxRateId         | number | Tax rate ID. For manual taxes the value is copied from tax ID, for all other cases the value is `0`.                                                          |
| sourceTaxRateType       | string | <p>Type of tax rate.<br><br>One of <code>AUTO</code>, <code>MANUAL</code>, <code>CUSTOM</code> (if tax is changed via API), <code>LEGACY</code>.</p>          |
| taxType                 | string | <p>Type of detailed tax for USA. <br><br>One of: <code>STATE</code>, <code>COUNTY</code>, <code>CITY</code>, <code>SPECIAL\_DISTRICT</code></p>               |
| taxClassCode            | string | <p>Tax classification code applied to product. <br><br>See: <a href="ref:country-codes">Tax classes by country</a></p>                                        |
| taxClassName            | string | <p>Name of the tax classification code applied to product. Available only in English. <br><br>See: <a href="ref:country-codes">Tax classes by country</a></p> |

#### files

| Field              | Type   | Description                                                                                                                                                                                                   |
| ------------------ | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| productFileId      | number | Internal unique file ID.                                                                                                                                                                                      |
| maxDownloads       | number | <p>Maximum number of allowed file downloads. <br><br>Read more on digital products in <a href="https://support.ecwid.com/hc/en-us/articles/207100559-Digital-products">Help Center</a>.</p>                   |
| remainingDownloads | number | Remaining number of download attempts for the file.                                                                                                                                                           |
| expire             | string | Date/time of the customer download link expiration.                                                                                                                                                           |
| name               | string | File name visible to the customer.                                                                                                                                                                            |
| description        | string | File description visible to the customer.                                                                                                                                                                     |
| size               | number | File size in bytes (64-bit integer).                                                                                                                                                                          |
| adminUrl           | string | <p>Link to the file download for the store owner. <br><br><strong>Keep caution</strong>: the link contains the API access token. Never share it and do not display the link in publically available code.</p> |
| customerUrl        | string | File download link sent to the customer after the order was paid.                                                                                                                                             |

#### selectedOptions

| Field       | Type             | Description                                                                                                                                                                                                                                                                                                                                                                                       |
| ----------- | ---------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| name        | string           | Name of the product option.                                                                                                                                                                                                                                                                                                                                                                       |
| type        | string           | <p>Type of the product option that defines its functionality. <br><br>One of:<br><code>CHOICE -</code> Dropdown, radio button, or size. Allows selecting only one value from the list.<br><code>CHOICES -</code> Checkbox. Allows selecting multiple values.<br><code>TEXT -</code> Text input or area.<br><code>DATE -</code> Datetime selector.<br><code>FILES -</code> Upload file option.</p> |
| value       | string           | <p>Selected/entered value for the option as <code>string</code>. <br><br>For <code>CHOICES</code> type, provides a string with all selected values separated by a comma.</p>                                                                                                                                                                                                                      |
| valuesArray | array            | <p>Selected/entered value for the option as <code>array</code>. </p><p></p><p>For the <code>CHOICES</code> type, provides an array with all selected values.</p>                                                                                                                                                                                                                                  |
| files       | array of objects | <p>Detailed information about files attached to the selected option.<br><br>Available only if the option type is <code>FILES.</code></p>                                                                                                                                                                                                                                                          |
| selections  | array of objects | <p>Details of selected product options. <br><br>If sent in "Update order" request, other fields will be recalculated based on information from <code>selections</code>.</p>                                                                                                                                                                                                                       |
| hexCodes    | array of strings | <p>List of HEX codes.</p><p>Defines what color must be displayed when user changes color in the <code>SWATCHES</code> option, for example: <code>\["#fff000"]</code>. <br><br>Requires <code>useImageAsSwatchSelector</code> to be <code>true</code>.</p>                                                                                                                                         |
| changedTime | number           | UNIX timestamp of the latest change in the product option                                                                                                                                                                                                                                                                                                                                         |

#### filesAttached

| Field | Type   | Description                                                       |
| ----- | ------ | ----------------------------------------------------------------- |
| id    | number | ID of the file uploaded through `FILES` type option.              |
| name  | string | Name of the file uploaded through `FILES` type option.            |
| size  | number | Size (in bytes) of the file uploaded through `FILES` type option. |
| url   | string | Download link of the file uploaded through `FILES` type option.   |

#### selections

| Field                 | Type   | Description                                                                                                                                                                                          |
| --------------------- | ------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| selectionTitle        | string | Name of the selected option value.                                                                                                                                                                   |
| selectionModifier     | number | <p>Price modifier of the selected option value. <br><br>Value can be negative, for example, <code>-10</code> if it decreases the product price.</p>                                                  |
| selectionModifierType | string | <p>Price modifier type.<br><br>One of: <br><code>PERCENT</code> - Price modifier applies as a percent from the product price.<br><code>ABSOLUTE</code> - Price modifier applies as a flat value.</p> |

#### recurringChargeSettings

<table><thead><tr><th>Field</th><th width="132">Type</th><th>Description</th></tr></thead><tbody><tr><td>recurringInterval</td><td>string</td><td><p>Subscription charge interval. </p><p><br>One of: <code>day</code>, <code>week</code>, <code>month</code>, <code>year</code>.</p></td></tr><tr><td>recurringIntervalCount</td><td>number</td><td>Charge interval count that depends on the <code>recurringInterval</code>. For example <code>3</code> - once per 3 months, if <code>recurringInterval</code> is <code>month.</code></td></tr><tr><td>subscriptionPriceWithSignUpFee</td><td>number</td><td>Total product cost including the first subscription payment.</td></tr><tr><td>signUpFee</td><td>number</td><td>Fees imposed on the first payment.</td></tr></tbody></table>

#### dimensions

| Field  | Type   | Description         |
| ------ | ------ | ------------------- |
| length | number | Length of a product |
| width  | number | Width of a product  |
| height | number | Height of a product |

#### shippingPerson

| Field               | Type   | Description                                           |
| ------------------- | ------ | ----------------------------------------------------- |
| name                | string | Full name of the customer.                            |
| companyName         | string | Customer's company name.                              |
| street              | string | Address line 1 and address line 2, separated by `\n`. |
| city                | string | City.                                                 |
| countryCode         | string | Two-letter country code.                              |
| countryName         | string | Country name.                                         |
| postalCode          | string | Postal/ZIP code.                                      |
| stateOrProvinceCode | string | State/province code, for example, `NY`.               |
| stateOrProvinceName | string | State/province name.                                  |
| phone               | string | Customer's phone number.                              |

#### billingPerson

| Field               | Type   | Description                                           |
| ------------------- | ------ | ----------------------------------------------------- |
| name                | string | Full name of the customer.                            |
| companyName         | string | Customer's company name.                              |
| street              | string | Address line 1 and address line 2, separated by `\n`. |
| city                | string | City.                                                 |
| countryCode         | string | Two-letter country code.                              |
| countryName         | string | Country name.                                         |
| postalCode          | string | Postal/ZIP code.                                      |
| stateOrProvinceCode | string | State/province code, for example, `NY`.               |
| stateOrProvinceName | string | State/province name.                                  |
| phone               | string | Customer's phone number.                              |

#### customSurcharges

| Field                 | Type                  | Description                                                                                                                                                                                |
| --------------------- | --------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| id                    | string                | <p>Surcharge ID. <br><br>If not specified default value: <code>Custom Surcharge</code></p>                                                                                                 |
| value                 | number                | Surcharge value.                                                                                                                                                                           |
| type                  | string                | <p>Surcharges type.<br><br>One of: <br><code>"PERCENT"</code> - Surcharge applies as a percent from the product price.<br><code>"ABSOLUTE"</code> - Surcharge applies as a flat value.</p> |
| total                 | number                | Total value of the surcharge.                                                                                                                                                              |
| totalWithoutTax       | number                | Total value of the surcharge without taxes.                                                                                                                                                |
| description           | string                | Surcharge description defined by the store owner.                                                                                                                                          |
| descriptionTranslated | string                | Available translations for the surcharge description.                                                                                                                                      |
| taxable               | boolean               | Defines if taxes apply to the surcharge.                                                                                                                                                   |
| taxes                 | array [taxes](#taxes) | Details about taxes applied to the surcharge.                                                                                                                                              |

#### discounts

| Field        | Type                                           | Description                                          |
| ------------ | ---------------------------------------------- | ---------------------------------------------------- |
| discountInfo | array of objects [discountInfo](#discountinfo) | Details about **promotions** applied to the product. |
| total        | number                                         | Sum of **promotions** applied to the order.          |

#### discountCoupon

| Field            | Type                                 | Description                                                                                                                                                                          |
| ---------------- | ------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| id               | number                               | Internal discount coupon ID.                                                                                                                                                         |
| name             | string                               | Name of the discount coupon visible in Ecwid admin.                                                                                                                                  |
| code             | string                               | Discount coupon code.                                                                                                                                                                |
| discountType     | string                               | <p>Discount type.<br><br>One of: <br><code>ABS</code><br><code>PERCENT</code><br><code>SHIPPING</code><br><code>ABS\_AND\_SHIPPING</code><br><code>PERCENT\_AND\_SHIPPING</code></p> |
| status           | string                               | <p>Discount coupon state.<br><br>One of:<br><code>ACTIVE</code><br><code>PAUSED</code><br><code>EXPIRED</code><br><code>USEDUP</code></p>                                            |
| discount         | number                               | Discount value applied to the order total.                                                                                                                                           |
| launchDate       | string                               | The date of coupon launch, for example, `2014-06-06 08:00:00 +0000`.                                                                                                                 |
| expirationDate   | string                               | Coupon expiration date, for example, `2014-06-06 08:00:00 +0000`.                                                                                                                    |
| totalLimit       | number                               | The minimum order subtotal the coupon applies to.                                                                                                                                    |
| usesLimit        | string                               | Number of uses limitation: `UNLIMITED`, `ONCEPERCUSTOMER`, `SINGLE`                                                                                                                  |
| applicationLimit | string                               | <p>Application limit for discount coupons.<br><br>One of:<br><code>UNLIMITED</code><br><code>NEW\_CUSTOMER\_ONLY</code><br><code>REPEAT\_CUSTOMER\_ONLY</code></p>                   |
| creationDate     | string                               | Discount coupon creation date.                                                                                                                                                       |
| updateDate       | string                               | Date of the last discount coupon update.                                                                                                                                             |
| orderCount       | number                               | Amount of orders where the discount coupon was used previously.                                                                                                                      |
| catalogLimit     | object [catalogLimit](#cataloglimit) | Products and categories the coupon can be applied to                                                                                                                                 |

#### catalogLimit

| Field      | Type             | Description                                        |
| ---------- | ---------------- | -------------------------------------------------- |
| products   | array of numbers | List of product IDs the coupon can be applied to.  |
| categories | array of numbers | List of category IDs the coupon can be applied to. |

#### shippingOption

| Field                   | Type                                                     | Description                                                                                                                                                                                                                                                                           |
| ----------------------- | -------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| shippingCarrierName     | string                                                   | If an order is fulfilled with a native shipping carrier integration or a shipping app, this field holds carrier's name.                                                                                                                                                               |
| shippingMethodName      | string                                                   | Name of the shipping option visible at the checkout.                                                                                                                                                                                                                                  |
| shippingMethodId        | string                                                   | Internal shipping method ID.                                                                                                                                                                                                                                                          |
| shippingRate            | number                                                   | Shipping rate for the order.                                                                                                                                                                                                                                                          |
| shippingRateWithoutTax  | number                                                   | Shipping rate without taxes.                                                                                                                                                                                                                                                          |
| discountedShippingRate  | number                                                   | <p>Shipping rate with applied shipping discounts.<br><br><strong>Read-only</strong></p>                                                                                                                                                                                               |
| estimatedTransitTime    | number/string                                            | <p>Delivery time estimation.<br><br>Depending on the store settings it can be a number, for example, <code>5</code> or a string – <code>4-9 days</code>.<br><br>The string value is equal to the <code>description</code> field in the <code>Get shipping options</code> request.</p> |
| isPickup                | boolean                                                  | Defines if this is a store pickup method.                                                                                                                                                                                                                                             |
| pickupInstruction       | string                                                   | Instructions for customer on how to pickup the order.                                                                                                                                                                                                                                 |
| fulfillmentType         | string                                                   | <p>Shipping type.<br><br>One of:</p><p><code>shipping</code></p><p><code>pickup</code></p><p><code>deliver</code></p>                                                                                                                                                                 |
| timeSlotLengthInMinutes | number                                                   | Length of the delivery time slot in minutes.                                                                                                                                                                                                                                          |
| discount                | object [shippingOptionDiscount](#shippingoptiondiscount) | DIscount applied to the shipping method.                                                                                                                                                                                                                                              |

#### handlingFee

| Field       | Type   | Description                                                    |
| ----------- | ------ | -------------------------------------------------------------- |
| name        | string | Handling fee name set by store admin, for example, `Wrapping`. |
| value       | number | Handling fee flat value.                                       |
| description | string | Handling fee's description for customers.                      |

#### predictedPackages

| Name          | Type   | Description                                                           |
| ------------- | ------ | --------------------------------------------------------------------- |
| height        | number | Height of a predicted package.                                        |
| width         | number | Width of a predicted package.                                         |
| length        | number | Length of a predicted package.                                        |
| weight        | number | Total weight of a predicted package.                                  |
| declaredValue | number | Declared value of a predicted package (subtotal of items in package). |

#### shipments

| Field           | Type                                       | Description                                                                        |
| --------------- | ------------------------------------------ | ---------------------------------------------------------------------------------- |
| id              | string                                     | ID of the purchased shipping label.                                                |
| created         | date                                       | The date/time of shipping label purchase, for example, `2020-04-23 19:13:43 +0000` |
| shipTo          | object [shippingPerson](#shippingperson)   | Name and address of the person entered in shipping information.                    |
| shipFrom        | object [shipFrom](#shipfrom)               | Shipping origin address. If matches company address, company address is returned.  |
| parcel          | object [parcel](#parcel)                   | Information about the selected package to ship items to customer.                  |
| shippingService | object [shippingService](#shippingservice) | Selected shipping service.                                                         |
| tracking        | object [tracking](#tracking)               | Tracking details provided by shipping service.                                     |
| shippingLabel   | object [shippingLabel](#shippinglabel)     | Shipping label details.                                                            |

#### shipFrom

| Field               | Type   | Description                                                                                        |
| ------------------- | ------ | -------------------------------------------------------------------------------------------------- |
| companyName         | string | Store owner's company name.                                                                        |
| email               | string | Store owner's email.                                                                               |
| street              | string | Store's address in 1 or 2 lines format. If two address lines provided, they are separated by `\n`. |
| city                | string | City where the store is located.                                                                   |
| countryCode         | string | Two-letter country code.                                                                           |
| countryName         | string | Country where the store is located.                                                                |
| postalCode          | string | Postal/ZIP code for the store's location.                                                          |
| stateOrProvinceCode | string | State/province code, for example, `NY`.                                                            |
| stateOrProvinceName | string | State/province name.                                                                               |
| phone               | string | Store's phone number.                                                                              |

#### parcel

| Field         | Type   | Description                                                                                                                                                       |
| ------------- | ------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| weight        | number | Total weight of the labeled package.                                                                                                                              |
| weightUnit    | string | <p>Weight unit of the package.<br><br>One of:<br><code>CARAT</code><br><code>GRAM</code><br><code>OUNCE</code><br><code>POUND</code><br><code>KILOGRAM</code></p> |
| width         | number | Width of the labeled package.                                                                                                                                     |
| height        | number | Height of the labeled package.                                                                                                                                    |
| length        | number | Length of the labeled package.                                                                                                                                    |
| dimensionUnit | string | <p>Dimension unit of the package.<br><br>One of:<br><code>MM</code><br><code>CM</code><br><code>IN</code><br><code>YD</code></p>                                  |

#### shippingService

| Field              | Type   | Description                                                                                                                                                                                                                                 |
| ------------------ | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| carrier            | string | <p>Carrier used for shipping the order. Available only for integrations build by Ecwid team.<br><br>One of:<br><code>USPS</code><br><code>UPS</code><br><code>FEDEX</code><br><code>CANADA\_POST</code><br><code>AUSTRALIA\_POST</code></p> |
| carrierName        | string | Name of shipping option in store settings.                                                                                                                                                                                                  |
| carrierServiceName | string | Specific carrier's name visible at the checkout.                                                                                                                                                                                            |
| carrierServiceCode | string | Internal carrier code.                                                                                                                                                                                                                      |

#### tracking

| Field            | Type   | Description                                       |
| ---------------- | ------ | ------------------------------------------------- |
| tracking\_number | string | Tracking number provided by the shipping service. |
| tracking\_url    | string | Link to the delivery tracking  page.              |
| estimatedDays    | number | Estimated delivery time in days.                  |

#### shippingLabel

| Field      | Type   | Description                       |
| ---------- | ------ | --------------------------------- |
| label\_url | string | Link for download shipping label. |

#### discountInfo

<table><thead><tr><th width="191.83984375">Field</th><th width="163.10546875">Type</th><th>Description</th></tr></thead><tbody><tr><td>id</td><td>number</td><td>Internal discount ID.</td></tr><tr><td>value</td><td>number</td><td>Discount value.</td></tr><tr><td>type</td><td>string</td><td><p>Discount type.<br><br>One of:</p><ul><li><code>ABS</code></li><li><code>PERCENT</code></li></ul></td></tr><tr><td>base</td><td>string</td><td><p>Discount base. <br><br>One of: </p><ul><li><code>SUBTOTAL</code>  - Discount is based on order subtotal.</li><li><code>ITEM</code>   - Discount is only applied to certain products in the order.</li><li><code>SHIPPING</code>   - Discount is only applied to order shipping costs.</li><li><code>ON_MEMBERSHIP</code>   - Discount is only applied if customer belongs to a certain customer group.</li><li><code>ON_TOTAL_AND_MEMBERSHIP</code> - Discount is applied to </li><li><code>CUSTOM</code>  - Discount is created by an app with a custom logic.</li></ul></td></tr><tr><td>orderTotal</td><td>number</td><td>Minimum order subtotal the discount applies to.</td></tr><tr><td>membershipId</td><td>number</td><td>Customer group ID to which the discount is limited.</td></tr><tr><td>description</td><td>string</td><td>Description of a discount visible at the checkout. Available only for discounts with <code>CUSTOM</code> base.</td></tr><tr><td>appliesToItems</td><td>array of numbers</td><td>List of product IDs to which the discount can be applied.</td></tr><tr><td>appliesToOrderItems</td><td>array of objects </td><td>List of internal order item IDs, which defines a list of products the discount is applied in this specific order.</td></tr></tbody></table>

#### creditCardStatus

| Field      | Type   | Description                                                     |
| ---------- | ------ | --------------------------------------------------------------- |
| avsMessage | string | Address verification status returned by the payment system.     |
| cvvMessage | string | Credit card verification status returned by the payment system. |

#### extraFields

| Field                                         | Type   | Description                                                                                             |
| --------------------------------------------- | ------ | ------------------------------------------------------------------------------------------------------- |
| ecwid\_order\_delivery\_time\_interval\_start | string | Start of the delivery date/datetime interval.                                                           |
| ecwid\_order\_delivery\_time\_interval\_end   | string | End of the delivery date/datetime interval.                                                             |
| ecwid\_order\_delivery\_time\_display\_format | string | <p>Format of the delivery date chosen.<br><br>One of:<br><code>DATE</code><br><code>DATETIME</code></p> |

#### orderExtraFields

| Field                      | Type   | Description                                                                                                                                                                                                                                                                                |
| -------------------------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| id                         | string | Internal ID defined for the checkout extra field.                                                                                                                                                                                                                                          |
| value                      | string | Extra field value. Length cannot exceed 255 characters.                                                                                                                                                                                                                                    |
| customerInputType          | string | One of: `""`,`"TEXT"`, `"SELECT"`, `"DATETIME"`                                                                                                                                                                                                                                            |
| title                      | string | Extra field title visible at the checkout.                                                                                                                                                                                                                                                 |
| orderDetailsDisplaySection | string | <p>Defines a place where the field is visible to the store admin on the order details page. <br><br>One of:<br><code>shipping\_info</code> </p><p><code>billing\_info</code></p><p><code>customer\_info</code></p><p><code>order\_comments</code><br><br>Empty if the field is hidden.</p> |
| orderBy                    | string | Extra field position. Use it to sort fields within the same `orderDetailsDisplaySection`                                                                                                                                                                                                   |

#### refunds

| Field  | Type   | Description                                                                                                                                                                                                                                   |
| ------ | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| date   | string | The datetime of a refund, for example, `2014-06-06 18:57:19 +0000`                                                                                                                                                                            |
| source | string | <p>What action triggered refund. <br><br>One of:<br><code>CP</code> - changed by the store owner in Ecwid admin<br><code>API</code> - changed by an app through API<br><code>External</code> - refund made from payment processor website</p> |
| reason | string | A text reason for a refund. 256 characters max.                                                                                                                                                                                               |
| amount | number | Amount of this specific refund (not total amount refunded for order. see `redundedAmount` field)                                                                                                                                              |

#### utmData

| Field    | Type   | Description                                                 |
| -------- | ------ | ----------------------------------------------------------- |
| source   | string | Traffic source that indicates where the customer come from. |
| campaign | string | Saves the name of the advertising campaign if there is one. |
| medium   | string | Type of traffic that indicates customers reach the website. |

#### utmDataSets

| Field     | Type   | Description                                                 |
| --------- | ------ | ----------------------------------------------------------- |
| timestamp | string | Datetime of saving UTM data into the local browser storage. |
| source    | string | Traffic source that indicates where the customer come from. |
| campaign  | string | Saves the name of the advertising campaign if there is one. |
| medium    | string | Type of traffic that indicates customers reach the website. |

#### invoices

| Field      | Type   | Description                                                                                                                              |
| ---------- | ------ | ---------------------------------------------------------------------------------------------------------------------------------------- |
| internalId | number | Internal ID of the order invoice.                                                                                                        |
| id         | string | Public ID showed in the invoice.                                                                                                         |
| created    | string | Datetime of invoice creation in UTC +0.                                                                                                  |
| link       | string | Download link for the invoice in PDF format.                                                                                             |
| type       | string | <p>Invoice type. <br>One of:<code>A</code><br><code>SALE</code> - regular invoice<br><code>FULL\_CANCEL</code> - full refund invoice</p> |

#### taxesOnShipping

| Field | Type   | Description                                 |
| ----- | ------ | ------------------------------------------- |
| name  | string | Name of the tax applied to shipping costs.  |
| value | number | Value of the tax applied to shipping costs. |
| total | number | Total of taxes applied to shipping costs.   |

#### externalOrderData

| Field                  | Type    | Description                                                     |
| ---------------------- | ------- | --------------------------------------------------------------- |
| externalOrderId        | string  | Order ID from the external platform.                            |
| externalFulfillment    | boolean | Defines if the order uses external fulfillment.                 |
| refererChannel         | string  | Channel where the order is referred from.                       |
| refererId              | string  | Unique referrer ID for the order.                               |
| platformSpecificFields | string  | Stringified map with external fields in the "key:value" format. |

#### translations

Object with text field translations in the `"lang": "text"` format, where the `"lang"` is an ISO 639-1 language code. For example:

```
{
    "en": "Sample text",
    "nl": "Voorbeeldtekst"
}
```

Translations are available for all active store languages. Only the default language translations are returned if no other translations are provided for the field. Find active store languages with <mark style="color:green;">`GET`</mark> `/profile` request > `languages` > `enabledLanguages`.
