# Get app subscription status

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/application`&#x20;

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_store_profile`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Field              | Type                | Description                                                                                                                       |
| ------------------ | ------------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| subscription       | object subscription | Subscription details for the application.                                                                                         |
| subscriptionStatus | string              | <p>Application status in Ecwid store. <br><br>One of: <br><code>ACTIVE</code><br><code>SUSPENDED</code><br><code>TRIAL</code></p> |

#### subscription

| Field          | Type   | Description                                                                                                                       |
| -------------- | ------ | --------------------------------------------------------------------------------------------------------------------------------- |
| startDate      | string | Datetime of the app subscription start.                                                                                           |
| expirationDate | string | Datetime when the app subscription ends.                                                                                          |
| status         | string | <p>Application status in Ecwid store. <br><br>One of: <br><code>ACTIVE</code><br><code>SUSPENDED</code><br><code>TRIAL</code></p> |
