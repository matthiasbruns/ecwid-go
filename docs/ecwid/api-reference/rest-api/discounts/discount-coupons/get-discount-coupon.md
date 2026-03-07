# Get discount coupon

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/discount_coupons/{discountCouponId}`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
GET /api/v3/1003/discount_coupons/162428889 HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
```

Response:

```json
{
  "id": 162428889,
  "name": "10% OFF",
  "code": "DISC",
  "discountType": "ABS",
  "status": "ACTIVE",
  "discount": 10,
  "launchDate": "2022-07-28 23:00:00 +0000",
  "usesLimit": "UNLIMITED",
  "repeatCustomerOnly": false,
  "applicationLimit": "UNLIMITED",
  "creationDate": "2022-07-29 15:22:35 +0000",
  "updateDate": "2024-05-01 05:26:28 +0000",
  "orderCount": 1
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_discount_coupons`

### Path params

All path params are required.

| Param            | Type   | Description                  |
| ---------------- | ------ | ---------------------------- |
| storeId          | number | Ecwid store ID.              |
| discountCouponId | number | Internal discount coupon ID. |

### Query params

All query params are optional.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>responseFields</td><td>string</td><td>Specify the exact fields to receive in response JSON. If not specified, the response JSON will have all available fields for the entity.<br><br>For example: <code>?responseFields=id,name,status</code></td></tr></tbody></table>

Example of using `responseFields` param:

{% tabs %}
{% tab title="Request" %}

```
curl --location 'https://app.ecwid.com/api/v3/1003/discount_coupons/DISC?responseFields=id,name,status' \
--header 'Authorization: Bearer secret_ab***cd'
```

{% endtab %}

{% tab title="Response" %}

```json
{
    "id": 162428889,
    "name": "Summer Promo",
    "status": "ACTIVE"
}
```

{% endtab %}
{% endtabs %}

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

<table><thead><tr><th width="203">Field</th><th width="205">Type</th><th>Description</th></tr></thead><tbody><tr><td>id</td><td>number</td><td>Internal unique coupon ID.</td></tr><tr><td>name</td><td>string</td><td>Discount coupon name visible on the storefront.</td></tr><tr><td>code</td><td>string</td><td>Discount coupon code used for its activation at the checkout.</td></tr><tr><td>discountType</td><td>string</td><td>Discount type<br><br>One of: <br><code>ABS</code><br><code>PERCENT</code><br><code>SHIPPING</code><br><code>ABS_AND_SHIPPING</code><br><code>PERCENT_AND_SHIPPING</code></td></tr><tr><td>status</td><td>string</td><td>Current state of the discount coupon.<br><br>One of: <br><code>ACTIVE</code><br><code>PAUSED</code><br><code>EXPIRED</code><br><code>USEDUP</code></td></tr><tr><td>discount</td><td>number</td><td>Discount coupon value.</td></tr><tr><td>launchDate</td><td>string</td><td>The date of coupon launch, for example, <code>2014-06-06 08:00:00 +0400</code>. <br><br>Any date provided will be corrected to the UTC +0 timezone.</td></tr><tr><td>expirationDate</td><td>string</td><td>Coupon expiration date, e.g. <code>2014-06-06 08:00:00 +0400</code>.<br><br>Any date provided will be corrected to the UTC +0 timezone.</td></tr><tr><td>totalLimit</td><td>number</td><td>Minimum order subtotal for the discount coupon to be applied.</td></tr><tr><td>usesLimit</td><td>string</td><td>Number of uses limitation.<br><br>One of: <br><code>UNLIMITED</code><br><code>ONCEPERCUSTOMER</code><br><code>SINGLE</code></td></tr><tr><td>applicationLimit</td><td>string</td><td><p>User application limit for the discount coupon. </p><p><br>One of: <br><code>UNLIMITED</code> - no user application limits.<br><code>NEW_CUSTOMER_ONLY</code> - discount coupon can be applied only by customers without placed orders.<br><code>REPEAT_CUSTOMER_ONLY</code> - discount coupon can be applied only by customers who placed orders in the store before.</p></td></tr><tr><td>creationDate</td><td>string</td><td>Coupon creation date. Format example: <code>2023-06-29 11:36:55 +0000</code></td></tr><tr><td>updateDate</td><td>string</td><td>Coupon update date. Format example: <code>2023-06-29 11:36:55 +0000</code></td></tr><tr><td>orderCount</td><td>number</td><td>Count of orders where the discount coupon was used.</td></tr><tr><td>catalogLimit</td><td>object <a href="#cataloglimit">catalogLimit</a></td><td>Product and category limitations for the discount coupon. <br><br>If empty, discount coupon can be applied to all products and categories available on the storefront.</td></tr><tr><td>shippingLimit</td><td>object <a href="#shippinglimit">shippingLimit</a></td><td>Shipping method limitations for the discount coupon. <br><br>If empty, discount coupon can be applied to any shipping method available at the checkout.</td></tr></tbody></table>

#### catalogLimit

<table><thead><tr><th width="246">Field</th><th width="185">Type</th><th>Description</th></tr></thead><tbody><tr><td>products</td><td>array of numbers</td><td>List of product IDs the discount coupon can be applied to, for example, <code>[123456,234567]</code></td></tr><tr><td>categories</td><td>array of numbers</td><td>List of category IDs the discount coupon can be applied to, for example, <code>[0,87253552,765257901]</code></td></tr></tbody></table>

#### shippingLimit

<table><thead><tr><th width="187">Field</th><th width="185">Type</th><th>Description</th></tr></thead><tbody><tr><td>shippingMethods</td><td>array of strings</td><td>List of shipping method IDs the discount coupon can be applied to, for example, <code>["18765-8651899366181"]</code></td></tr></tbody></table>
