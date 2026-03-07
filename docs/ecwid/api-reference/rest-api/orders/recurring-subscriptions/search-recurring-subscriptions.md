# Search recurring subscriptions

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/subscriptions`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
GET /api/v3/1003/subscriptions HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
```

Response:

```json
{
  "total": 10,
  "count": 5,
  "offset": 0,
  "limit": 100,
  "items": [
    {
      "subscriptionId": 66839,
      "customerId": 43343272,
      "status": "CANCELLED",
      "statusUpdated": "2021-07-23 21:17:26 +0000",
      "created": "2021-06-16 12:53:40 +0000",
      "cancelled": "2021-07-23 21:17:26 +0000",
      "nextCharge": "2021-07-16 12:53:40 +0000",
      "createTimestamp": 1623848020508,
      "updateTimestamp": 1627075046478,
      "chargeSettings": {
        "recurringInterval": "MONTH",
        "recurringIntervalCount": 1
      },
      "paymentMethod": {
        "creditCardMaskedNumber": "*1111",
        "creditCardBrand": "visa"
      },
      "orderTemplate": {
        "id": "121",
        "email": "test@test.test",
        "additionalInfo": {
          "creditCard": "*1111 (12/2024)",
          "creditCardExpirationMonth": "12",
          "creditCardExpirationYear": "2024",
          "google_customer_id": "11111111.2222222222",
          "stripeCardId": "pm_1X2xXXX34XXxXxxXxxxXXXX",
          "stripeCreditCardBrand": "visa",
          "stripeCreditCardLast4Digit": "1111",
          "stripeCustomerId": "cus_XXXxxxX12xxxxX",
          "stripeFingerprint": "xxxxxXXX1X2xxxXX",
          "stripeLiveMode": "false"
        },
        "orderComments": "",
        "paymentMethod": "Credit or debit card",
        "paymentModule": "Stripe",
        "total": 26.84,
        "subtotal": 12,
        "usdTotal": 32.545878290056045,
        "tax": 4.84,
        "customerTaxExempt": false,
        "customerTaxId": "",
        "customerTaxIdValid": false,
        "reversedTaxApplied": false,
        "items": [
          {
            "id": 762397105,
            "productId": 11111111,
            "categoryId": 0,
            "price": 12,
            "productPrice": 12,
            "shipping": 10,
            "tax": 4.84,
            "fixedShippingRate": 0,
            "sku": "006789",
            "name": "Mug",
            "shortDescription": "This is the best product in the world!",
            "quantity": 1,
            "quantityInStock": 37,
            "weight": 0.2,
            "trackQuantity": false,
            "fixedShippingRateOnly": false,
            "digital": false,
            "productAvailable": true,
            "imageUrl": "https://d2j6dbq0eux0bg.cloudfront.net/images/11111/111111.jpg",
            "recurringChargeSettings": {
              "recurringInterval": "MONTH",
              "intervalCount": 1
            },
            "selectedOptions": [
              {
                "name": "Color",
                "type": "CHOICE",
                "value": "White",
                "valuesArray": [
                  "White"
                ],
                "selections": [
                  {
                    "selectionTitle": "White",
                    "selectionModifier": 0,
                    "selectionModifierType": "ABSOLUTE"
                  }
                ]
              }
            ],
            "taxes": [
              {
                "name": "VAT",
                "value": 22,
                "total": 4.84,
                "taxOnDiscountedSubtotal": 2.64,
                "taxOnShipping": 2.2,
                "includeInPrice": true
              }
            ],
            "dimensions": {
              "length": 21,
              "width": 21,
              "height": 21
            }
          }
        ],
        "billingPerson": {
          "name": "First Last",
          "companyName": "",
          "street": "Otto-Braun-Straße",
          "city": "Berlin",
          "countryCode": "DE",
          "countryName": "Germany",
          "postalCode": "10178",
          "stateOrProvinceCode": "BE",
          "stateOrProvinceName": "Berlin",
          "phone": "+491625555012"
        },
        "shippingPerson": {
          "name": "First Last",
          "companyName": "",
          "street": "Otto-Braun-Straße",
          "city": "Berlin",
          "countryCode": "DE",
          "countryName": "Germany",
          "postalCode": "10178",
          "stateOrProvinceCode": "BE",
          "stateOrProvinceName": "Berlin",
          "phone": "+11111111111"
        },
        "shippingOption": {
          "shippingMethodName": "Local delivery – City",
          "shippingRate": 10
        },
        "handlingFee": {},
        "pricesIncludeTax": false
      },
      "orders": [
        {
          "id": 266645518,
          "total": 26.84,
          "createDate": "2021-06-16 12:53:39 +0000"
        }
      ]
    }
  ]
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_subscriptions`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

All query params are optional.

| Name                   | Type   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                |
| ---------------------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| id                     | number | Search term for the internal subscription ID.                                                                                                                                                                                                                                                                                                                                                                                              |
| createdFrom            | string | <p>Datetime when subscription was created (lower bound). </p><p><br>For example, <code>2024-05-26 13:37:46 +0000</code></p>                                                                                                                                                                                                                                                                                                                |
| createdTo              | string | <p>Datetime when subscription was created (upper bound). </p><p><br>For example, <code>2024-05-26 13:37:46 +0000</code></p>                                                                                                                                                                                                                                                                                                                |
| cancelledFrom          | string | <p>Datetime when subscription was cancelled (lower bound). </p><p><br>For example, <code>2024-05-26 13:37:46 +0000</code></p>                                                                                                                                                                                                                                                                                                              |
| cancelledTo            | string | <p>Datetime when subscription was cancelled (upper bound). </p><p><br>For example, <code>2024-05-26 13:37:46 +0000</code></p>                                                                                                                                                                                                                                                                                                              |
| updatedFrom            | string | <p>Datetime when subscription was updated for the last time (lower bound). </p><p><br>For example, <code>2024-05-26 13:37:46 +0000</code></p>                                                                                                                                                                                                                                                                                              |
| updatedTo              | string | <p>Datetime when subscription was updated for the last time (upper bound). </p><p><br>For example, <code>2024-05-26 13:37:46 +0000</code></p>                                                                                                                                                                                                                                                                                              |
| customerId             | number | Search term for the internal customer ID.                                                                                                                                                                                                                                                                                                                                                                                                  |
| status                 | string | <p>Subscription status . <br><br>One of: <br><code>ACTIVE</code><br><code>CANCELLED</code><br><code>LAST\_CHARGE\_FAILED</code> </p><p><code>REQUIRES\_PAYMENT\_CONFIRMATION</code></p>                                                                                                                                                                                                                                                    |
| nextChargeFrom         | string | <p>Datetime of the next subscription charge (lower bound). </p><p><br>For example, <code>2024-05-26 13:37:46 +0000</code></p>                                                                                                                                                                                                                                                                                                              |
| nextChargeTo           | string | <p>Datetime of the next subscription charge (upper bound). </p><p><br>For example, <code>2024-05-26 13:37:46 +0000</code></p>                                                                                                                                                                                                                                                                                                              |
| recurringInterval      | string | <p>Search term for the time scale of subscription's recurring interval. </p><p><br>One of: <br><code>DAY</code><br><code>WEEK</code><br><code>MONTH</code><br><code>YEAR</code></p>                                                                                                                                                                                                                                                        |
| recurringIntervalCount | number | <p>Search term for the frequency of subscription's recurring charges. <br><br>One of: <br>for <code>DAY</code> - <code>1</code> (daily)<br>for <code>WEEK</code> - <code>1</code> (weekly) or <code>2</code> (biweekly)<br>for <code>MONTH</code> - <code>1</code> (monthly) or <code>3</code> (quarterly)<br>for <code>YEAR</code> - <code>1</code> (annually)</p>                                                                        |
| productId              | number | Search term for the ID of the subscription product.                                                                                                                                                                                                                                                                                                                                                                                        |
| email                  | string | Email of the customer who bought a subscription product.                                                                                                                                                                                                                                                                                                                                                                                   |
| orderId                | string | Search term for the order ID with the subscription product inside.                                                                                                                                                                                                                                                                                                                                                                         |
| orderTotal             | number | Subscription's order total.                                                                                                                                                                                                                                                                                                                                                                                                                |
| orderCreatedFrom       | string | <p>Datetime when the order with a subscription was placed (lower bound). </p><p><br>For example, <code>2024-05-26 13:37:46 +0000</code></p>                                                                                                                                                                                                                                                                                                |
| orderCreatedTo         | string | <p>atetime when the order with a subscription was placed (upper bound). </p><p><br>For example, <code>2024-05-26 13:37:46 +0000</code></p>                                                                                                                                                                                                                                                                                                 |
| limit                  | number | Limit to the number of returned items. Maximum and default value (if not specified) is `100`.                                                                                                                                                                                                                                                                                                                                              |
| offset                 | number | <p>Offset from the beginning of the returned items list. Used when the response contains more items than <code>limit</code> allows to receive in one request.<br><br>Usually used to receive all items in several requests with multiple of a hundred, for example:<br><br><code>?offset=0</code> for the first request,<br><code>?offset=100</code>, for the second request,<br><code>?offset=200</code>, for the third request, etc.</p> |
| sortBy                 | string | <p>Sorting order for the response. <br><br>One of:<br><code>DATE\_CREATED\_ASC</code><br><code>DATE\_CREATED\_DESC</code> (default) <br><code>ORDER\_COUNT\_ASC</code><br><code>ORDER\_COUNT\_DESC</code><br><code>NEXT\_CHARGE\_DATE\_ASC</code><br><code>NEXT\_CHARGE\_DATE\_DESC</code></p>                                                                                                                                             |

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
| items  | array of objects [items](#items) | Detailed information about returned subscriptions.                                           |

#### items

| Field           | Type                                     | Description                                                                                                                                               |
| --------------- | ---------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| subscriptionId  | number                                   | Internal subscription ID.                                                                                                                                 |
| customerId      | number                                   | Internal ID of the customer who bought subscription.                                                                                                      |
| status          | string                                   | <p>One of: <br><code>ACTIVE</code><br><code>CANCELLED</code><br><code>LAST\_CHARGE\_FAILED</code> </p><p><code>REQUIRES\_PAYMENT\_CONFIRMATION</code></p> |
| statusUpdated   | string                                   | <p>Datetime when the subscription status was updated for the last time. </p><p><br>For example, <code>2024-05-26 13:37:46 +0000</code></p>                |
| created         | string                                   | <p>Datetime when the order containing subscription was placed. </p><p><br>For example, <code>2024-05-26 13:37:46 +0000</code></p>                         |
| cancelled       | string                                   | <p>Datetime when the subscription was cancelled. </p><p><br>For example, <code>2024-05-26 13:37:46 +0000</code></p>                                       |
| nextCharge      | string                                   | <p>Datetime of the next recurring charge for the subscription. </p><p><br>For example, <code>2024-05-26 13:37:46 +0000</code></p>                         |
| createTimestamp | number                                   | <p>Datetime when the order containing subscription was placed in UNIX timestamp format.</p><p></p><p>For example, <code>1623848020508</code></p>          |
| updateTimestamp | number                                   | <p>Datetime when the subscription status was updated for the last time in UNIX timestamp format.</p><p></p><p>For example, <code>1623848020508</code></p> |
| chargeSettings  | object [chargeSettings](#chargesettings) | Details about recurring charges set up for the subscription.                                                                                              |
| paymentMethod   | object [paymentMethod](#paymentmethod)   | Details about the payment method used for the subscription.                                                                                               |
| orders          | object [orders](#orders)                 | Short details about the order containing the subscription: ID, total, creation date.                                                                      |
| orderTemplate   | object [orderTemplate](#ordertemplate)   | Detailed information about the order containing the subscription.                                                                                         |

#### chargeSettings

| Field                  | Type   | Description                                                                                                                                                                                                                                                                                                                                                         |
| ---------------------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| recurringInterval      | string | <p>Search term for the time scale of subscription's recurring interval. </p><p><br>One of: <br><code>DAY</code><br><code>WEEK</code><br><code>MONTH</code><br><code>YEAR</code></p>                                                                                                                                                                                 |
| recurringIntervalCount | number | <p>Search term for the frequency of subscription's recurring charges. <br><br>One of: <br>for <code>DAY</code> - <code>1</code> (daily)<br>for <code>WEEK</code> - <code>1</code> (weekly) or <code>2</code> (biweekly)<br>for <code>MONTH</code> - <code>1</code> (monthly) or <code>3</code> (quarterly)<br>for <code>YEAR</code> - <code>1</code> (annually)</p> |

#### paymentMethod

| Field                  | Type   | Description                             |
| ---------------------- | ------ | --------------------------------------- |
| creditCardMaskedNumber | string | Credit card masked number, e.g. \*1111. |
| creditCardBrand        | string | Credit card brand, e.g. Visa.           |

#### orders

| Field      | Type   | Description                                                   |
| ---------- | ------ | ------------------------------------------------------------- |
| id         | number | Internal order ID.                                            |
| total      | number | Order total .                                                 |
| createDate | string | Order creation date and time, e.g. 2021-06-16 12:53:39 +0000. |

#### orderTemplate

A JSON object with the order details in the same format as in the [GET order request](https://docs.ecwid.com/api-reference/rest-api/orders/get-order).
