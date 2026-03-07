# Search discount coupons

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/discount_coupons`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
GET /api/v3/1003/discount_coupons HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
```

Response:

```json
{
  "total": 2,
  "count": 2,
  "offset": 0,
  "limit": 100,
  "items": [
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
    },
    {
      "id": 224219782,
      "name": "Test Coupon",
      "code": "DISC2",
      "discountType": "PERCENT",
      "status": "ACTIVE",
      "discount": 10,
      "launchDate": "2024-09-01 23:00:00 +0000",
      "usesLimit": "UNLIMITED",
      "repeatCustomerOnly": false,
      "applicationLimit": "UNLIMITED",
      "creationDate": "2024-09-02 07:55:21 +0000",
      "updateDate": "2024-09-02 08:33:26 +0000",
      "orderCount": 3
    }
  ]
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_discount_coupons`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

All query params are optional.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>offset</td><td>number</td><td>Offset from the beginning of the returned items list. Used when the response contains more items than <code>limit</code> allows to receive in one request.<br><br>Usually used to receive all items in several requests with multiple of a hundred, for example:<br><br><code>?offset=0</code> for the first request,<br><code>?offset=100</code>, for the second request,<br><code>?offset=200</code>, for the third request, etc.</td></tr><tr><td>limit</td><td>number</td><td>Limit to the number of returned items. Maximum and default value (if not specified) is <code>100</code>.</td></tr><tr><td>code</td><td>string</td><td>Search term for the discount coupon code.</td></tr><tr><td>discount_type</td><td>string</td><td><p>Search term for the discount coupon type. </p><p><br>One of:<br><code>ABS</code><br><code>PERCENT</code><br><code>SHIPPING</code><br><code>ABS_AND_SHIPPING</code><br><code>PERCENT_AND_SHIPPING</code></p></td></tr><tr><td>availability</td><td>string</td><td>Search term for the current state of the discount coupon. <br><br>One of: <br><code>ACTIVE</code><br><code>PAUSED</code><br><code>EXPIRED</code><br><code>USEDUP</code></td></tr><tr><td>createdFrom</td><td>number/string</td><td>Coupon creation date/time (lower bound). Supported formats: UNIX timestamp, date/time. Examples: <code>1447804800</code>, <code>2023-01-15 19:27:50</code>.</td></tr><tr><td>createdTo</td><td>number/string</td><td>Coupon creation date/time (upper bound). Supported formats: UNIX timestamp, date/time. Examples: <code>1447804800</code>, <code>2023-01-15 19:27:50</code>.</td></tr><tr><td>updatedFrom</td><td>number/string</td><td>Coupon last update date/time (lower bound). Supported formats: UNIX timestamp, date/time. Examples: <code>1447804800</code>, <code>2023-01-15 19:27:50</code>.</td></tr><tr><td>updatedTo</td><td>number/string</td><td>Coupon last update date/time (upper bound). Supported formats: UNIX timestamp, date/time. Examples: <code>1447804800</code>, <code>2023-01-15 19:27:50</code>.</td></tr><tr><td>responseFields</td><td>string</td><td>Specify the exact fields to receive in response JSON. If not specified, the response JSON will have all available fields for the entity.<br><br>For example: <code>?responseFields=items(id,name,status)</code></td></tr></tbody></table>

Example of using `responseFields` param:

{% tabs %}
{% tab title="Request" %}

```
curl --location 'https://app.ecwid.com/api/v3/1003/discount_coupons?responseFields=items(id,name,status)' \
--header 'Authorization: Bearer secret_ab***cd'
```

{% endtab %}

{% tab title="Response" %}

```json
{
    "items": [
        {
            "id": 162428889,
            "name": "Summer Promo",
            "status": "ACTIVE"
        },
        {
            "id": 224219782,
            "name": "Test Coupon",
            "status": "ACTIVE"
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

| Field  | Type                             | Description                                                                                  |
| ------ | -------------------------------- | -------------------------------------------------------------------------------------------- |
| total  | number                           | Total number of found items (might be more than the number of returned items).               |
| count  | number                           | Total number of items returned in the response.                                              |
| offset | number                           | Offset from the beginning of the returned items list specified in the request.               |
| limit  | number                           | Maximum number of returned items specified in the request. Maximum and default value: `100`. |
| items  | array of objects [items](#items) | Detailed information about returned discount coupons.                                        |

#### items

| Field            | Type                                   | Description                                                                                                                                                                                                                                                                                                                                                                       |
| ---------------- | -------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| id               | number                                 | Internal unique coupon ID.                                                                                                                                                                                                                                                                                                                                                        |
| name             | string                                 | Discount coupon name visible on the storefront.                                                                                                                                                                                                                                                                                                                                   |
| code             | string                                 | Discount coupon code used for its activation at the checkout.                                                                                                                                                                                                                                                                                                                     |
| discountType     | string                                 | <p>Discount type<br><br>One of: <br><code>ABS</code><br><code>PERCENT</code><br><code>SHIPPING</code><br><code>ABS\_AND\_SHIPPING</code><br><code>PERCENT\_AND\_SHIPPING</code></p>                                                                                                                                                                                               |
| status           | string                                 | <p>Current state of the discount coupon.<br><br>One of: <br><code>ACTIVE</code><br><code>PAUSED</code><br><code>EXPIRED</code><br><code>USEDUP</code></p>                                                                                                                                                                                                                         |
| discount         | number                                 | Discount coupon value.                                                                                                                                                                                                                                                                                                                                                            |
| launchDate       | string                                 | <p>The date of coupon launch, for example, <code>2014-06-06 08:00:00 +0400</code>. <br><br>Any date provided will be corrected to the UTC +0 timezone.</p>                                                                                                                                                                                                                        |
| expirationDate   | string                                 | <p>Coupon expiration date, e.g. <code>2014-06-06 08:00:00 +0400</code>.<br><br>Any date provided will be corrected to the UTC +0 timezone.</p>                                                                                                                                                                                                                                    |
| totalLimit       | number                                 | Minimum order subtotal for the discount coupon to be applied.                                                                                                                                                                                                                                                                                                                     |
| usesLimit        | string                                 | <p>Number of uses limitation.<br><br>One of: <br><code>UNLIMITED</code><br><code>ONCEPERCUSTOMER</code><br><code>SINGLE</code></p>                                                                                                                                                                                                                                                |
| applicationLimit | string                                 | <p>User application limit for the discount coupon. </p><p><br>One of: <br><code>UNLIMITED</code> - no user application limits.<br><code>NEW\_CUSTOMER\_ONLY</code> - discount coupon can be applied only by customers without placed orders.<br><code>REPEAT\_CUSTOMER\_ONLY</code> - discount coupon can be applied only by customers who placed orders in the store before.</p> |
| creationDate     | string                                 | Coupon creation date. Format example: `2023-06-29 11:36:55 +0000`                                                                                                                                                                                                                                                                                                                 |
| updateDate       | string                                 | Coupon update date. Format example: `2023-06-29 11:36:55 +0000`                                                                                                                                                                                                                                                                                                                   |
| orderCount       | number                                 | Count of orders where the discount coupon was used.                                                                                                                                                                                                                                                                                                                               |
| catalogLimit     | object [catalogLimit](#cataloglimit)   | <p>Product and category limitations for the discount coupon. <br><br>If empty, discount coupon can be applied to all products and categories available on the storefront.</p>                                                                                                                                                                                                     |
| shippingLimit    | object [shippingLimit](#shippinglimit) | <p>Shipping method limitations for the discount coupon. <br><br>If empty, discount coupon can be applied to any shipping method available at the checkout.</p>                                                                                                                                                                                                                    |

#### catalogLimit

<table><thead><tr><th width="246">Field</th><th width="185">Type</th><th>Description</th></tr></thead><tbody><tr><td>products</td><td>array of numbers</td><td>List of product IDs the discount coupon can be applied to, for example, <code>[123456,234567]</code></td></tr><tr><td>categories</td><td>array of numbers</td><td>List of category IDs the discount coupon can be applied to, for example, <code>[0,87253552,765257901]</code></td></tr></tbody></table>

#### shippingLimit

<table><thead><tr><th width="187">Field</th><th width="185">Type</th><th>Description</th></tr></thead><tbody><tr><td>shippingMethods</td><td>array of strings</td><td>List of shipping method IDs the discount coupon can be applied to, for example, <code>["18765-8651899366181"]</code></td></tr></tbody></table>
