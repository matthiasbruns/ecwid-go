# Update recurring subscription

<mark style="color:purple;">`PUT`</mark> `https://app.ecwid.com/api/v3/{storeId}/subscriptions/{subscriptionId}`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
PUT /api/v3/1003/subscriptions/66839 HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
Content-Type: application/json
Cache-Control: no-cache

{
  "nextCharge": "2021-08-16 12:53:40 +0000"
}
```

Response:

```json
{
  "updateCount": 1
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `update_subscriptions`

### Path params

All path params are required.

| Param          | Type   | Description               |
| -------------- | ------ | ------------------------- |
| storeId        | number | Ecwid store ID.           |
| subscriptionId | number | Internal subscription ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Request JSON

A JSON object with the following fields:

| Field          | Type                                     | Description                                                                                                                                               |
| -------------- | ---------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| chargeSettings | object [chargeSettings](#chargesettings) | Details about recurring charges set up for the subscription.                                                                                              |
| nextCharge     | string                                   | <p>Datetime of the next recurring charge for the subscription. </p><p><br>For example, <code>2024-05-26 13:37:46 +0000</code></p>                         |
| status         | string                                   | <p>One of: <br><code>ACTIVE</code><br><code>CANCELLED</code><br><code>LAST\_CHARGE\_FAILED</code> </p><p><code>REQUIRES\_PAYMENT\_CONFIRMATION</code></p> |

#### chargeSettings

| Field                  | Type   | Description                                                                                                                                                                                                                                                                                                                                                         |
| ---------------------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| recurringInterval      | string | <p>Search term for the time scale of subscription's recurring interval. </p><p><br>One of: <br><code>DAY</code><br><code>WEEK</code><br><code>MONTH</code><br><code>YEAR</code></p>                                                                                                                                                                                 |
| recurringIntervalCount | number | <p>Search term for the frequency of subscription's recurring charges. <br><br>One of: <br>for <code>DAY</code> - <code>1</code> (daily)<br>for <code>WEEK</code> - <code>1</code> (weekly) or <code>2</code> (biweekly)<br>for <code>MONTH</code> - <code>1</code> (monthly) or <code>3</code> (quarterly)<br>for <code>YEAR</code> - <code>1</code> (annually)</p> |

### Response JSON

A JSON object with the following fields:

| Field       | Type   | Description                                                                                                                                                                                   |
| ----------- | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| updateCount | number | <p>The number of updated items that defines if the request was successful.<br><br>One of:</p><p><code>1</code> if the item was updated,</p><p><code>0</code> if the item was not updated.</p> |
