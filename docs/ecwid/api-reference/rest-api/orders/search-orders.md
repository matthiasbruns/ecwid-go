# Search orders

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/orders`&#x20;

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
  "total": 23,
  "count": 1,
  "offset": 0,
  "limit": 1,
  "items": [
    {
      "id": "EBJFT",
      "internalId": 492512057,
      "refundedAmount": 0,
      "subtotal": 500,
      "subtotalWithoutTax": 500,
      "total": 600,
      "totalWithoutTax": 590,
      "giftCardRedemption": 0,
      "totalBeforeGiftCardRedemption": 600,
      "giftCardDoubleSpending": false,
      "email": "",
      "tax": 10,
      "customerTaxExempt": false,
      "customerTaxIdValid": true,
      "b2b_b2c": "b2c",
      "reversedTaxApplied": false,
      "customerRequestedInvoice": false,
      "customerFiscalCode": "",
      "electronicInvoicePecEmail": "",
      "electronicInvoiceSdiCode": "",
      "couponDiscount": 10,
      "paymentStatus": "PAID",
      "fulfillmentStatus": "SHIPPED",
      "orderNumber": 492512057,
      "vendorOrderNumber": "EBJFT",
      "publicUid": "EBJFT",
      "volumeDiscount": 0,
      "membershipBasedDiscount": 0,
      "totalAndMembershipBasedDiscount": 0,
      "customSurcharges": [],
      "discount": 0,
      "usdTotal": 642.404477130936,
      "createDate": "2024-05-01 05:26:28 +0000",
      "updateDate": "2024-05-01 05:26:28 +0000",
      "createTimestamp": 1714541188,
      "updateTimestamp": 1714541188,
      "discountCoupon": {
        "id": 215189589,
        "name": "Test Coupon",
        "code": "DISC",
        "discountType": "ABS",
        "status": "ACTIVE",
        "discount": 10,
        "launchDate": "2024-04-30 23:00:00 +0000",
        "usesLimit": "UNLIMITED",
        "repeatCustomerOnly": false,
        "applicationLimit": "UNLIMITED",
        "creationDate": "2024-05-01 05:26:28 +0000",
        "updateDate": "2024-05-01 05:26:28 +0000",
        "orderCount": 0
      },
      "items": [
        {
          "id": 1741253497,
          "productId": 439710255,
          "price": 500,
          "priceWithoutTax": 500,
          "productPrice": 0,
          "sku": "000001",
          "quantity": 1,
          "shortDescriptionTranslated": {
            "ru": "",
            "en": ""
          },
          "tax": 0,
          "shipping": 0,
          "quantityInStock": 0,
          "name": "Pizza",
          "nameTranslated": {
            "ru": "",
            "en": "Pizza"
          },
          "isShippingRequired": true,
          "weight": 0,
          "trackQuantity": false,
          "fixedShippingRateOnly": false,
          "imageUrl": "https://d2j6dbq0eux0bg.cloudfront.net/images/15695068/2870741131.jpg",
          "smallThumbnailUrl": "https://d2j6dbq0eux0bg.cloudfront.net/images/15695068/2870741133.jpg",
          "hdThumbnailUrl": "https://d2j6dbq0eux0bg.cloudfront.net/images/15695068/2870741134.jpg",
          "fixedShippingRate": 0,
          "digital": false,
          "productAvailable": true,
          "couponApplied": false,
          "files": [
            {
              "productFileId": 92603033,
              "maxDownloads": 0,
              "remainingDownloads": 0,
              "expire": "2024-05-04 05:26:28 +0000",
              "name": "header.png",
              "description": "",
              "size": 99304,
              "adminUrl": "https://app.ecwid.com/api/v3/15695068/products/439710255/files/92603033",
              "customerUrl": "https://app.ecwid.com/download/15695068/d23f0e4ae9368716687f51b9d53820e2/header.png"
            },
            {
              "productFileId": 92603034,
              "maxDownloads": 0,
              "remainingDownloads": 0,
              "expire": "2024-05-04 05:26:28 +0000",
              "name": "logo.png",
              "description": "",
              "size": 9487,
              "adminUrl": "https://app.ecwid.com/api/v3/15695068/products/439710255/files/92603034",
              "customerUrl": "https://app.ecwid.com/download/15695068/dac5865171876f936e90391236228a94/logo.png"
            },
            {
              "productFileId": 92603035,
              "maxDownloads": 0,
              "remainingDownloads": 0,
              "expire": "2024-05-04 05:26:28 +0000",
              "name": "screen1.png",
              "description": "",
              "size": 56497,
              "adminUrl": "https://app.ecwid.com/api/v3/15695068/products/439710255/files/92603035",
              "customerUrl": "https://app.ecwid.com/download/15695068/6cb52be701b4181b733d42dfe1306e18/screen1.png"
            }
          ],
          "taxable": true,
          "isCustomerSetPrice": false,
          "attributes": []
        }
      ],
      "refunds": [],
      "shippingOption": {
        "shippingMethodId": "customShippingId",
        "shippingMethodName": "Shipping",
        "shippingRate": 110,
        "shippingRateWithoutTax": 110,
        "isPickup": false,
        "fulfillmentType": "SHIPPING",
        "isShippingLimit": false
      },
      "predictedPackage": [],
      "shippingLabelAvailableForShipment": false,
      "shipments": [],
      "additionalInfo": {

      },
      "paymentParams": {

      },
      "extraFields": {

      },
      "ticket": -160802399,
      "hidden": false,
      "taxesOnShipping": [
        {
          "name": "Custom tax",
          "value": 10,
          "total": 10
        }
      ],
      "disableAllCustomerNotifications": false,
      "externalFulfillment": false,
      "utmDataSets": [],
      "invoices": [],
      "pricesIncludeTax": false
    }
  ]
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

<table data-full-width="false"><thead><tr><th width="204.296875">Name</th><th width="138.41015625">Type</th><th>Description</th></tr></thead><tbody><tr><td>ids</td><td>string</td><td>List of order identifiers. Works with order ID, internal order ID, order prefixes, and suffixes. <br><br>Supports multiple values, for example: <code>EG4H2,J77J8,SALE-G01ZG</code></td></tr><tr><td>offset</td><td>number</td><td><p>Offset from the beginning of the returned items list. Used when the response contains more items than <code>limit</code> allows to receive in one request.<br><br>Usually used to receive all items in several requests with multiple of a hundred, for example:<br><code>?offset=0</code> for the first request,</p><p><code>?offset=100</code>, for the second request,</p><p><code>?offset=200</code>, for the third request, etc.</p></td></tr><tr><td>limit</td><td>number</td><td>Limit to the number of returned items. Maximum and default value (if not specified) is <code>100</code>.</td></tr><tr><td>keywords</td><td>string</td><td>Search term that supports: order ID, external transaction ID, billing and shipping address, customer email, shipping tracking code, item SKUs, names, selected options, and private admin notes. <br><br>Any special characters must be URI-encoded.</td></tr><tr><td>email</td><td>string</td><td>Search term for customer email.</td></tr><tr><td>customerId</td><td>number</td><td>Search term for customer's internal ID.</td></tr><tr><td>productId</td><td>number/string</td><td>Search term for IDs of products in order. Supports multiple values separated by comma, for example: <code>10031004,86427531</code>.</td></tr><tr><td>totalFrom</td><td>number</td><td>Search term for minimum order total.</td></tr><tr><td>totalTo</td><td>number</td><td>Search term for maximum order total.</td></tr><tr><td>createdFrom</td><td>number/string</td><td>Order placement datetime (lower bound). Supported formats: UNIX timestamp, date/time. Examples: <code>1447804800</code>, <code>2023-01-15 19:27:50</code></td></tr><tr><td>createdTo</td><td>number/string</td><td>Order placement datetime (upper bound). Supported formats: UNIX timestamp, date/time. Examples: <code>1447804800</code>, <code>2023-01-15 21:30:00</code></td></tr><tr><td>updatedFrom</td><td>number/string</td><td>Order latest update datetime (lower bound). Supported formats: UNIX timestamp, date/time. Examples: <code>1447804800</code>, <code>2023-01-15 19:27:50</code></td></tr><tr><td>updatedTo</td><td>number/string</td><td>Order latest update date/time (upper bound). Supported formats: UNIX timestamp, date/time. Examples: <code>1447804800</code>, <code>2023-01-15 21:30:00</code></td></tr><tr><td>pickupTimeFrom</td><td>number/string</td><td>Order pickup datetime (lower bound). Supported formats: UNIX timestamp, date/time. Examples: <code>1447804800</code>, <code>2023-01-15 19:27:50</code></td></tr><tr><td>pickupTimeTo</td><td>number/string</td><td>Order pickup datetime (upper bound). Supported formats: UNIX timestamp, date/time. Examples: <code>1447804800</code>, <code>2023-01-15 21:30:00</code></td></tr><tr><td>shippingMethod</td><td>string</td><td>Search term for the shipping method name selected on the checkout.</td></tr><tr><td>fulfillmentStatus</td><td>string</td><td>Order shipping status. Supports multiple values separated by a comma. <br><br>Supported values: <code>AWAITING_PROCESSING</code>, <code>PROCESSING</code>, <code>SHIPPED</code>, <code>DELIVERED</code>, <code>WILL_NOT_DELIVER</code>, <code>RETURNED</code>, <code>READY_FOR_PICKUP</code>, <code>OUT_FOR_DELIVERY</code>, <code>CUSTOM_FULFILLMENT_STATUS_1</code>, <code>CUSTOM_FULFILLMENT_STATUS_2</code>, <code>CUSTOM_FULFILLMENT_STATUS_3</code>.<br><br>Read more about order statuses in <a href="https://support.ecwid.com/hc/en-us/articles/207806235-Order-details-and-statuses-overview#-understanding-order-statuses"><strong>Help Center</strong></a>.</td></tr><tr><td>paymentMethod</td><td>string</td><td>Search term for the payment method name selected on the checkout.</td></tr><tr><td>paymentModule</td><td>string</td><td>Search term for the payment module selected on the checkout.<br><br>Payment module contains the name of the internal payment app (built by Ecwid dev team) or a custom one in the <code>"paymentModule":"CUSTOM_PAYMENT_APP-client_id"</code> format. </td></tr><tr><td>paymentStatus</td><td>string</td><td>Order payment status. Supports multiple values separated by a comma. <br><br>Supported values: <code>AWAITING_PAYMENT</code>, <code>PAID</code>, <code>CANCELLED</code>, <code>REFUNDED</code>, <code>PARTIALLY_REFUNDED</code>, <code>INCOMPLETE</code>, <code>CUSTOM_PAYMENT_STATUS_1</code>, <code>CUSTOM_PAYMENT_STATUS_2</code>, <code>CUSTOM_PAYMENT_STATUS_3</code>.<br><br>Read more about order statuses in <a href="https://support.ecwid.com/hc/en-us/articles/207806235-Order-details-and-statuses-overview#-understanding-order-statuses"><strong>Help Center</strong></a>.</td></tr><tr><td>acceptMarketing</td><td>boolean</td><td>Set <code>true</code> to find orders where customer has accepted email marketing.</td></tr><tr><td>containsPreorderItems</td><td>boolean</td><td>Set <code>true</code> to find orders with pre-order products (out-of-stock products available to purchase).</td></tr><tr><td>couponCode</td><td>string</td><td>Search term for discount coupon <code>code</code> applied to the order.</td></tr><tr><td>subscriptionId</td><td>number</td><td>Search term for ID of subscription assigned to the order.</td></tr><tr><td>refererId</td><td>number</td><td>Search term for ID of order referer.</td></tr><tr><td>responseFields</td><td>string</td><td>Specify the exact fields to receive in response JSON. If not specified, the response JSON will have all available fields for the entity.<br>For example: <code>?responseFields=total,items(id,email,total)</code></td></tr></tbody></table>

Example of using `responseFields` param:

{% tabs %}
{% tab title="Request" %}

```
curl --location 'https://app.ecwid.com/api/v3/1003/orders?responseFields=total,items(id,email,total)' \
--header 'Authorization: Bearer secret_ab***cd'
```

{% endtab %}

{% tab title="Response" %}

```json
{
    "total": 3,
    "items": [
        {
            "id": "K8XTQ",
            "email": "artem.kudryashov@lightspeedhq.com",
            "total": 10
        },
        {
            "id": "1HOTQ",
            "email": "artem.kudryashov@lightspeedhq.com",
            "total": 101
        },
        {
            "id": "CYQLB",
            "email": "artem.kudryashov@lightspeedhq.com",
            "total": 109.1
        }
    ]
}
```

{% endtab %}
{% endtabs %}

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Field  | Type                                       | Description                                                                                  |
| ------ | ------------------------------------------ | -------------------------------------------------------------------------------------------- |
| total  | number                                     | Total number of found items (might be more than the number of returned items).               |
| count  | number                                     | Total number of items returned in the response.                                              |
| offset | number                                     | Offset from the beginning of the returned items list specified in the request.               |
| limit  | number                                     | Maximum number of returned items specified in the request. Maximum and default value: `100`. |
| items  | array of objects [orderItems](#orderitems) | Detailed information about returned orders.                                                  |

#### orderItems

| Field                             | Type                                                      | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| --------------------------------- | --------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| id                                | string                                                    | <p>Unique order identificator with prefix and suffix defined by the store admin. For example, order ID <code>MYSTORE-X8UYE</code> contains <code>MYSTORE-</code> prefix.<br><br>Order ID is shown to customers in any notifications and to the store owner in Ecwid admin and notifications.</p>                                                                                                                                                                                                                                                                                                                                           |
| subtotal                          | number                                                    | Cost of all products in the order (item's `price` x `quantity`) before any cost modifiers such as discounts, taxes, fees, etc. are applied.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| subtotalWithoutTax                | number                                                    | Order subtotal without taxes included in price (GROSS) when `pricesIncludeTax` is `true`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| total                             | number                                                    | Order total cost with all cost modifiers: shipping costs, taxes, fees, and discounts.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| totalWithoutTax                   | number                                                    | Order total without taxes. Calculates as `total` - `tax`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| refundedAmount                    | number                                                    | Sum of all refunds applied to the order.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| giftCardRedemption                | number                                                    | Amount deducted from the [Gift Card](https://support.ecwid.com/hc/en-us/articles/360002011419) balance and applied to order total.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| totalBeforeGiftCardRedemption     | number                                                    | Order total before the Gift Card was applied.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| email                             | string                                                    | Customer's email address.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| paymentModule                     | string                                                    | <p>Payment processor used to pay for the order online.<br><br>Only available to online payment integrations build by Ecwid team.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| paymentMethod                     | string                                                    | Name of the payment method customer chosen at the checkout.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| tax                               | number                                                    | <p>Sum of all taxes applied to products and shipping.<br><br>If the order is modified after being placed, this value is <strong>not</strong> recalculated automatically.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| customerTaxExempt                 | boolean                                                   | <p>Defines if the customer is tax exempt. Requires valid tax ID.<br><br>Read more about handling tax exempt customers in <a href="https://support.ecwid.com/hc/en-us/articles/213823045-Handling-tax-exempt-customers">Help Center</a>.</p>                                                                                                                                                                                                                                                                                                                                                                                                |
| customerTaxId                     | string                                                    | Tax ID entered by the customer at the checkout.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| customerTaxIdValid                | boolean                                                   | Defines if customer's tax ID is valid for tax exemption.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| reversedTaxApplied                | boolean                                                   | Defines if order tax was reversed (set to 0). Requires valid tax ID.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| b2b\_b2c                          | string                                                    | <p>Order type. One of:</p><p><code>b2b</code> - business-to-business</p><p><code>b2c</code> - business-to-consumer </p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| customerRequestedInvoice          | boolean                                                   | Defines if customer requested an invoice.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| customerFiscalCode                | string                                                    | Fiscale code of the customer.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| electronicInvoicePecEmail         | string                                                    | PEC email for order invoices.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| electronicInvoiceSdiCode          | string                                                    | SDI code for order invoices.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| ipAddress                         | string                                                    | Customer's IP address detected at the checkout.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| paymentStatus                     | string                                                    | <p>Order payment status. Supported values: <code>AWAITING\_PAYMENT</code>, <code>PAID</code>, <code>CANCELLED</code>, <code>REFUNDED</code>, <code>PARTIALLY\_REFUNDED</code>, <code>INCOMPLETE</code>, <code>CUSTOM\_PAYMENT\_STATUS\_1</code>, <code>CUSTOM\_PAYMENT\_STATUS\_2</code>, <code>CUSTOM\_PAYMENT\_STATUS\_3</code>.<br><br>Read more about order statuses in <a href="https://support.ecwid.com/hc/en-us/articles/207806235-Order-details-and-statuses-overview#-understanding-order-statuses"><strong>Help Center</strong></a>.</p>                                                                                        |
| fulfillmentStatus                 | string                                                    | <p>Order fulfillment status. Supported values: <code>AWAITING\_PROCESSING</code>, <code>PROCESSING</code>, <code>SHIPPED</code>, <code>DELIVERED</code>, <code>WILL\_NOT\_DELIVER</code>, <code>RETURNED</code>, <code>READY\_FOR\_PICKUP</code>, <code>OUT\_FOR\_DELIVERY</code>, <code>CUSTOM\_FULFILLMENT\_STATUS\_1</code>, <code>CUSTOM\_FULFILLMENT\_STATUS\_2</code>, <code>CUSTOM\_FULFILLMENT\_STATUS\_3</code>.<br><br>Read more about order statuses in <a href="https://support.ecwid.com/hc/en-us/articles/207806235-Order-details-and-statuses-overview#-understanding-order-statuses"><strong>Help Center</strong></a>.</p> |
| refererUrl                        | string                                                    | URL of the page when order was placed without page slugs (hash `#` part).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| orderComments                     | string                                                    | Order comments, left by a customer at the checkout.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| volumeDiscount                    | number                                                    | Sum of applied **promotions** based on subtotal. Included in the `discount` field.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| membershipBasedDiscount           | number                                                    | Sum of applied **promotions** based on customer group. Included in the `discount` field.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| totalAndMembershipBasedDiscount   | number                                                    | Sum of applied **promotions** based on both subtotal and customer group. Included in the `discount` field.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| customDiscount                    | array of numbers                                          | List of absolute discounts added by applications.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| discount                          | number                                                    | <p>Total order discount. Includes both promotions and discount coupons. <br><br>Calculated as the sum of the<code>couponDiscount</code> and <code>totalAndMembershipBasedDiscount</code> fields.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| couponDiscount                    | number                                                    | <p>Discount value from applied <strong>discount coupon</strong>, e.g. <code>10</code>.</p><p>Total order discount is the sum of the<code>couponDiscount</code> and <code>discount</code> fields.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| discountInfo                      | array [discounts](#discounts)                             | Detailed information about applied **promotions**.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| discountCoupon                    | object [discountCoupon](#discountcoupon)                  | Detailed information about applied **discount coupons**.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| customerId                        | number                                                    | Unique internal ID assigned to the customer.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| customSurcharges                  | array of objects [customSurcharges](#customsurcharges)    | Information about surcharges applied to the order.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| usdTotal                          | number                                                    | Order total converted from the store's currency to USD.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| globalReferer                     | string                                                    | URL that the customer came to the store from                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| createDate                        | string                                                    | The datetime when the order was placed, for example `2014-06-06 18:57:19 +0000`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| updateDate                        | string                                                    | The datetime of the latest order update. This includes all changes made from Ecwid admin or API. For example, `2014-06-06 18:57:19 +0000`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| createTimestamp                   | number                                                    | The datetime when the order was placed in UNIX timestamp, for example `1427268654`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| updateTimestamp                   | number                                                    | The datetime of the latest order update in UNIX timestamp. This includes all changes made from Ecwid admin or API. For example, `1427268654`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| customerGroup                     | string                                                    | <p>Name of the group the customer belongs to (if any).<br><br>Read more about <a href="https://support.ecwid.com/hc/en-us/articles/207807105-Customer-groups">customer groups</a>.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| customerGroupId                   | number                                                    | ID of the group the customer belongs to.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| items                             | array [items](#items)                                     | Detailed information about products in the order.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| refunds                           | array [refunds](#refunds)                                 | Details about refunds made to order.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| shippingPerson                    | object [shippingPerson](#shippingperson)                  | Name and shipping address details left by customer at the checkout.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| billingPerson                     | object [billingPerson](#billingperson)                    | Name and billing address details left by customer at the checkout.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| shippingOption                    | object [shippingOption](#shippingoption)                  | Details about the shipping option customer selected at the checkout.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| handlingFee                       | object [handlingFee](#handlingfee)                        | Details about fees applied to order.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| predictedPackages                 | object [predictedPackages](#predictedpackages)            | Minimum total dimensions and weight of a single shipping package that will be enough to carry all products added to the order.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| shippingLabelAvailableForShipment | boolean                                                   | Defines if the store owner can buy a shipping label through Ecwid for the order.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| shipments                         | array [shipments](#shipments)                             | Detailed information about purchased shipping label.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| additionalInfo                    | object                                                    | Internal order information for Ecwid services.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| paymentParams                     | object                                                    | Internal payment parameters for Ecwid services.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| extraFields                       | object [extraFields](#extrafields)                        | Names and values of custom checkout fields applied to the order.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| orderExtraFields                  | array [orderExtraFields](#orderextrafields)               | Additional optional information about the order's extra fields. Along with the value of the field, it contains technical information, such as id, type, etc. of the field. Total storage of extra fields cannot exceed 8Kb.                                                                                                                                                                                                                                                                                                                                                                                                                |
| hidden                            | boolean                                                   | Defines if the order is hidden from Ecwid admin. Applies to unsfinished orders only.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| trackingNumber                    | string                                                    | Shipping tracking code.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| paymentMessage                    | string                                                    | Error message sent by the online payment method. Only appears if a customer had issues with paying for the order online. When order becomes paid, `paymentMessage` is cleared                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| externalTransactionId             | string                                                    | Transaction ID saved to the order details by the payment system. For example, PayPal transaction ID.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| affiliateId                       | string                                                    | <p>If a store has several storefronts, this ID is used to track from which one the order came from.<br><br>Read more on setting up affiliate IDs in <a href="https://support.ecwid.com/hc/en-us/articles/207100679-How-to-track-which-storefront-an-order-came-from#add-special-id-to-the-integration-code">Help Center</a>.</p>                                                                                                                                                                                                                                                                                                           |
| creditCardStatus                  | object [creditCardStatus](#creditcardstatus)              | Saves verification messages if customer paid for the order with a credit card.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| privateAdminNotes                 | string                                                    | Private note added to the order by store owner.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| pickupTime                        | string                                                    | Order pickup time in the store date format (UTC +0 timezone), for example: `2017-10-17 05:00:00 +0000`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| taxesOnShipping                   | array [taxesOnShipping](#taxesonshipping)                 | Taxes applied to shipping 'as is'. `null` for old orders, `[]` for orders with taxes applied to subtotal only. Are not recalculated if order is updated later manually. Is calculated like: `(shippingRate + handlingFee)*(taxValue/100)`                                                                                                                                                                                                                                                                                                                                                                                                  |
| acceptMarketing                   | boolean                                                   | Defines if customer has accepted email marketing at the checkout. If `true` or `null`, you can use their email for promotions.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| refererId                         | string                                                    | Referer identifier. Can be set in storefront via JS or by creating / updating an order with REST API                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| disableAllCustomerNotifications   | boolean                                                   | <p>Defines if the customer should receive any email notifications:<br> <code>true</code> - no notifications are sent to the customer. If <code>false</code> - email notifications are sent to customer according to store mail notification settings. <br><br>This setting does not affect email notifications to the store owner.</p>                                                                                                                                                                                                                                                                                                     |
| externalFulfillment               | boolean                                                   | <p>Defines if the order is fulfilled with an external system and should not be managed through Ecwid:<br><code>true</code> - Ecwid will hide fulfillment status change feature and ability to set tracking number within Ecwid admin.</p><p><code>false</code> - store owner can manage order fulfillment within Ecwid admin (default value)</p>                                                                                                                                                                                                                                                                                           |
| externalOrderId                   | string                                                    | <p>Order ID in an external system where order is fulfilled. <br><br>Requires <code>externalFulfillment</code>  to be <code>true</code>.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| invoices                          | array [invoices](#invoices)                               | <p>Tax invoices generated for the order. <br><br><strong>Read-only</strong></p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| pricesIncludeTax                  | boolean                                                   | <p>Defines if taxes are included to product prices (GROSS or NET prices):<br> <code>true</code> - the tax rate is included in product prices (GROSS). <br><code>false</code> - the tax rate is not included in product prices (NET).<br></p><p>Read more about setting up taxes in <a href="https://support.ecwid.com/hc/en-us/articles/207099899-Setting-up-taxes">Help Center</a>.</p>                                                                                                                                                                                                                                                   |
| paymentSubtype                    | string                                                    | Internal field for Ecwid services.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| utmData                           | array [utmData](#utmdata)                                 | <p>UTM tags saved for the order. <br><br>Read more about using UTM tags in orders in <a href="https://support.ecwid.com/hc/en-us/articles/4545287177372">Help Center</a>.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| utmDataSets                       | array [utmDataSets](#utmdatasets)                         | Detailed information about UTM tags saved for the order. Contains more information than the `utmData` field.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| lang                              | string                                                    | <p>Defines a list of available languages or a single language for customer notifications. Must match one of the active store translations.<br></p><p>List of active store languages is available in the <mark style="color:green;"><code>GET</code></mark> <code>/profile</code>  request> <code>languages</code> > <code>enabledLanguages</code> field.</p>                                                                                                                                                                                                                                                                               |
| customerUserAgent                 | string                                                    | Details about the customer's device and platform used to place an order based on the [User-Agent](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/User-Agent) browser data.                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| externalOrderData                 | object [#externalorderdata](#externalorderdata "mention") | Details for orders created or managed externally, for example, by other Lightspeed products.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |

#### items

<table><thead><tr><th width="246">Field</th><th>Type</th><th>Description</th></tr></thead><tbody><tr><td>id</td><td>number</td><td>Order item ID unique for this order. Can be used to manage ordered items.</td></tr><tr><td>productId</td><td>number</td><td>Internal product ID. Can be used to find full product details with the <mark style="color:green;"><code>GET</code></mark> <code>/products</code> request.</td></tr><tr><td>subscriptionId</td><td>number</td><td>ID of the subscription available at Ecwid admin > My Sales > Subscriptions.</td></tr><tr><td>recurringChargeSettings</td><td>object <a href="#recurringchargesettings">recurringChargeSettings</a></td><td>Details about subscription charge intervals.</td></tr><tr><td>categoryId</td><td>number</td><td>ID of the category this product belongs to or was added from. <br><br>Returns <code>-1</code> if the product was added to the cart via the Buy Now button.</td></tr><tr><td>price</td><td>number</td><td>Price of product in the order with some modifiers (doesn't include discount modifiers).<br>Calculation: <code>productPrice</code> + <code>wholesalePrices</code> + price modifiers from selected <code>options</code>.</td></tr><tr><td>priceWithoutTax</td><td>number</td><td>Price of product in the order without taxes.</td></tr><tr><td>productPrice</td><td>number</td><td>Basic product price without any modifiers: options markups, discounts, taxes, fees.</td></tr><tr><td>costPrice</td><td>number</td><td>Purchase price of the product in the specific order used for reports and profit calculations.</td></tr><tr><td>weight</td><td>number</td><td>Weight of the product.</td></tr><tr><td>sku</td><td>string</td><td>Product SKU. <br><br>If the chosen options match a variation, this will be a variation SKU.</td></tr><tr><td>quantity</td><td>number</td><td>Quantity of the product in the order.</td></tr><tr><td>shortDescription</td><td>string</td><td>Product description truncated to 120 characters.</td></tr><tr><td>shortDescriptionTranslated</td><td>object <a href="#translations">translations</a></td><td>Available translations for product short description.</td></tr><tr><td>tax</td><td>number</td><td>Total tax applied to the product.</td></tr><tr><td>shipping</td><td>number</td><td>Partial shipping costs specific to the product.</td></tr><tr><td>quantityInStock</td><td>number</td><td>Number of products in stock in the store before placing the order.</td></tr><tr><td>name</td><td>string</td><td>Name of the product.</td></tr><tr><td>nameTranslated</td><td>object <a href="#translations">translations</a></td><td>Available translations for the product name.</td></tr><tr><td>isShippingRequired</td><td>boolean</td><td>Defines if the product requires shipping.</td></tr><tr><td>trackQuantity</td><td>boolean</td><td>Defines if low stock notifications to the store owner are enabled.</td></tr><tr><td>imageUrl</td><td>string</td><td>Link to the main product image.</td></tr><tr><td>fixedShippingRateOnly</td><td>boolean</td><td>Defines if the product has a unique fixed shipping rate. <br><br>If <code>true</code>, shipping costs won't calculate for the product and <code>fixedShippingRate</code> value will be used instead.</td></tr><tr><td>fixedShippingRate</td><td>number</td><td>Fixed shipping costs for the product. <br><br>Affects shipping costs only if <code>fixedShippingRateOnly</code> is <code>true</code>.</td></tr><tr><td>couponApplied</td><td>boolean</td><td>Defines if the product has a discount coupon applied.</td></tr><tr><td>selectedOptions</td><td>array <a href="#selectedoptions">selectedOptions</a></td><td>Product options values selected by the customer at the checkout.</td></tr><tr><td>taxes</td><td>array <a href="#taxes">taxes</a></td><td>Detailed information about taxes applied to the product in this order.</td></tr><tr><td>combinationId</td><td>number</td><td>ID of a product variation whos options mathes with values chosen by the customer at the checkout.<br><br>Read more on product variations in <a href="https://support.ecwid.com/hc/en-us/articles/207100299-Product-variations">Help Center.</a></td></tr><tr><td>digital</td><td>boolean</td><td>Defines if the product has any downloadable files attached.<br><br>Read more on digital products in <a href="https://support.ecwid.com/hc/en-us/articles/207100559-Digital-products">Help Center</a>.</td></tr><tr><td>files</td><td>array of objects <a href="#files">files</a></td><td>Details about downloadable files attached to the product.</td></tr><tr><td>dimensions</td><td>object <a href="#dimensions">dimensions</a></td><td>Details about product dimensions used for shipping costs calculations.</td></tr><tr><td>couponAmount</td><td>number</td><td><p>Discount applied to the product from discount coupon. </p><p><br>If the order is manually updated after being placed, this field is not recalculated automatically.</p></td></tr><tr><td>discounts</td><td>array <a href="#discounts">discounts</a></td><td><strong>Promotions</strong> applied to the specific product in the order.</td></tr><tr><td>taxesOnShipping</td><td>array <a href="#taxesonshipping">taxesOnShipping</a></td><td><p>Taxes applied to shipping costs for the product with the calculation formula of: <code>(shippingRate + handlingFee)*(taxValue/100)</code></p><p><br>If the order is manually updated after being placed, this field is not recalculated automatically.</p></td></tr><tr><td>taxAlreadyDeductedFromPrice</td><td>boolean</td><td>Indicates whether the taxes have been already deducted from the price.</td></tr><tr><td>isCustomerSetPrice</td><td>boolean</td><td>If <code>true</code>, customer set a custom product price using the "<a href="https://support.ecwid.com/hc/en-us/articles/360018423259-Pay-What-You-Want-pricing">Pay What You Want</a>" feature. <br><br>In this case, both the product <code>price</code> and <code>selectedPrice</code> -> <code>value</code> fields contain the price set by a customer.<br>If <code>false</code>, customer didn't choose the custom price. Therefore, the <code>selectedPrice</code> -> <code>value</code> field will be absent and the <code>price</code> field contains default product price set by the store owner.</td></tr><tr><td>selectedPrice</td><td>object selectedPrice > value</td><td><p>If <code>isCustomerSetPrice</code> is <code>true</code>, this field contains the "<a href="https://support.ecwid.com/hc/en-us/articles/360018423259-Pay-What-You-Want-pricing">Pay What You Want</a>" price set by a customer at the checkout.</p><p>Example with the PWYW price set to 100:<br><code>"selectedPrice": { "value": 100 }</code></p></td></tr><tr><td>isPreorder</td><td>boolean</td><td>Defines if the product was pre-ordered in this order.<br><br>Read more about accepting pre-orders in <a href="https://support.ecwid.com/hc/en-us/articles/5135873315100-Accepting-pre-orders-in-your-Ecwid-store">Help Center</a>.</td></tr><tr><td>attributes</td><td>array of objects <a href="#attributes">attributes</a></td><td>Details about product attributes.<br><br>Read more on product attributes in <a href="https://support.ecwid.com/hc/en-us/articles/207807495-Product-types-and-attributes">Help Center</a>.</td></tr></tbody></table>

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
| changedTime | number           | UNIX timestamp of the latest change in the product option.                                                                                                                                                                                                                                                                                                                                        |

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

#### shippingOptionDiscount

| Field                     | Type             | Description                                                                         |
| ------------------------- | ---------------- | ----------------------------------------------------------------------------------- |
| id                        | string           | Internal discount ID                                                                |
| value                     | number           | Discount value.                                                                     |
| type                      | string           | Discount type: `PERCENT` or `ABSOLUTE`.                                             |
| base                      | string           | Base from which the discount is calculated. For example, `SHIPPING`.                |
| orderTotal                | number           | Order total cost before the discount is applied.                                    |
| description               | string           | Discount name/description.                                                          |
| appliesToItems            | array of numbers | Internal IDs of items discount is applied to.                                       |
| appliesToOrderItems       | array of numbers | Numbered IDs of items in the order (`0`, `1`, `2`, etc.) discount is applied to.    |
| triggeredByOrderItems     | array of numbers | Internal IDs of items that triggered the discount.                                  |
| appliesToShippingMethodId | string           | Id of the shipping method discount applies to. For example, `"6589-1709547151586"`. |

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
