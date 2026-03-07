# Get tax invoices for order

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/orders/{orderId}/invoices`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
GET /api/v3/1003/orders/JJ5HH/invoices HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
```

Response:

```json
{
    "invoices": [
        {
            "internalId": 1002002,
            "id": "AB000308",
            "created": "2024-11-07 12:09:45 +0000",
            "link": "https://app.ecwid.com/download_tax_invoice?ownerid=JJ5HH&invoice_id=1002002&access_key=a***5",
            "type": "SALE"
        }
        {
            "internalId": 1002004,
            "id": "AB000310",
            "created": "2024-11-17 02:13:10 +0000",
            "link": "https://app.ecwid.com/download_tax_invoice?ownerid=JJ5HH&invoice_id=1002004&access_key=G***c",
            "type": "FULL_CANCEL"
        }
    ]
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_orders` , `read_invoices`

### Path params

All path params are required.

<table><thead><tr><th width="150">Param</th><th width="126">Type</th><th>Description</th></tr></thead><tbody><tr><td>storeId</td><td>number</td><td>Ecwid store ID.</td></tr><tr><td>orderId</td><td>number</td><td>Order ID. Can contain prefixes and suffixes, for example: <code>EG4H2,J77J8,SALE-G01ZG</code></td></tr></tbody></table>

### Query params

All query params are optional.

<table><thead><tr><th width="172">Param</th><th width="111">Type</th><th>Description</th></tr></thead><tbody><tr><td>responseFields</td><td>string</td><td>Limit JSON response by specific fields. If specified, all missing fields will be removed from the response body.<br>Example: <code>?responseFields=invoices(id,link)</code></td></tr></tbody></table>

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Field    | Type                        | Description                                       |
| -------- | --------------------------- | ------------------------------------------------- |
| invoices | array [invoices](#invoices) | Details about tax invoices created for the order. |

#### invoices

<table><thead><tr><th width="168">Field</th><th>Type</th><th>Description</th></tr></thead><tbody><tr><td>internalId</td><td>number</td><td>Internal ID of the tax invoice.</td></tr><tr><td>id</td><td>string</td><td>Public invoice ID visible in the tax invoice.</td></tr><tr><td>created</td><td>string</td><td>Datetime of tax invoice creation in UTC +0.</td></tr><tr><td>link</td><td>string</td><td>Download link for the tax invoice in PDF format.</td></tr><tr><td>type</td><td>string</td><td><p>Tax invoice type. </p><p><br>One of:<br><code>SALE</code> - Regular sale invoice.<br><code>FULL_CANCEL</code> - Full refund invoice.</p></td></tr></tbody></table>
