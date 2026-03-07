# Create customer extra field

<mark style="color:blue;">`POST`</mark> `https://app.ecwid.com/api/v3/{storeId}/store_extrafields/customers`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
POST /api/v3/1003/store_extrafields/customers HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
Content-Type: application/json
Cache-Control: no-cache

{
  "title": "Offline card ID",
  "type": "TEXT",
  "shownOnOrderDetails": false
}
```

Response:

```json
{
  "createCount": 1,
  "key": "6r0fUGw"
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `update_customers_extrafields`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Request JSON

A JSON object with the following fields:

<table><thead><tr><th width="259">Field</th><th width="139">Type</th><th>Description</th></tr></thead><tbody><tr><td>title</td><td>string</td><td>Extra field visible at the checkout and in Ecwid admin.</td></tr><tr><td>type</td><td>string</td><td><p>Extra field type that defines its functionality. </p><p><br>Two values are supported for customer extra fields:<br><code>text</code> - Single-line text input. Supports placeholders and pre-defined values (default).<br><code>datetime</code> - Customizable date and time picker in the form of a calendar widget on the checkout.</p></td></tr><tr><td>shownOnOrderDetails</td><td>boolean</td><td>Defines if the extra field should be visible on the order details page in Ecwid admin.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Field       | Type   | Description                                                                                                                                                                                   |
| ----------- | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| createCount | number | <p>The number of created items that defines if the request was successful.<br><br>One of:</p><p><code>1</code> if the item was created,</p><p><code>0</code> if the item was not created.</p> |
| key         | string | Internal key (identificator) of the created customer extra field.                                                                                                                             |
