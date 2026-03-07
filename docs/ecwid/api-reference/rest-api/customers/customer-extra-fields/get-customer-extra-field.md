# Get customer extra field

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/store_extrafields/customers/{extrafield_key}`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
GET /api/v3/1003/store_extrafields/customers/6r0fUGw HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
```

Response:

```json
{
  "key": "6r0fUGw",
  "title": "Offline card ID",
  "entityTypes": [
    "CUSTOMERS"
  ],
  "type": "TEXT",
  "shownOnOrderDetails": false,
  "linkedWithCheckoutField": false,
  "createdDate": "2025-03-26 10:24:22 +0000",
  "lastModifiedDate": "2025-03-26 10:24:22 +0000"
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_customers_extrafields`

### Path params

All path params are required.

| Param           | Type   | Description                                               |
| --------------- | ------ | --------------------------------------------------------- |
| storeId         | number | Ecwid store ID.                                           |
| extrafield\_key | string | Internal key (identificator) of the customer extra field. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

<table><thead><tr><th width="259">Field</th><th width="139">Type</th><th>Description</th></tr></thead><tbody><tr><td>key</td><td>string</td><td>Internal ID of the customer extra field.</td></tr><tr><td>title</td><td>string</td><td>Extra field visible at the checkout and in Ecwid admin.</td></tr><tr><td>entityTypes</td><td>string</td><td>Always <code>CUSTOMERS</code> for customer extra fields.</td></tr><tr><td>type</td><td>string</td><td><p>Extra field type that defines its functionality. </p><p><br>Two values are supported for customer extra fields:<br><code>text</code> - Single-line text input. Supports placeholders and pre-defined values (default).<br><code>datetime</code> - Customizable date and time picker in the form of a calendar widget on the checkout.</p></td></tr><tr><td>shownOnOrderDetails</td><td>boolean</td><td>Defines if the extra field should be visible on the order details page in Ecwid admin.</td></tr><tr><td>linkedWithCheckoutField</td><td>boolean</td><td>If <code>true</code>, this customer extra field is linked to a certain checkout extra field.</td></tr><tr><td>createdDate</td><td>string</td><td>Datetime when the extra fields was created</td></tr><tr><td>lastModifiedDate</td><td>string</td><td>Datetime of the latest extra field change.</td></tr></tbody></table>
