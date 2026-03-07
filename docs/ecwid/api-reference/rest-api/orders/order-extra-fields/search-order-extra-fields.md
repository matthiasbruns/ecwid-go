# Search order extra fields

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/orders/{orderId}/extraFields`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
GET /api/v3/1003/orders/JJ5HH/extraFields HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
```

Response:

```json
[
    {
        "id": "tips",
        "value": "3%",
        "customerInputType": "TOGGLE_BUTTON_GROUP",
        "title": "Support our pet shop with a small donation",
        "orderDetailsDisplaySection": "",
        "orderBy": "0"
    }
]
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_orders`

### Path params

All path params are required.

<table><thead><tr><th width="141">Param</th><th width="152">Type</th><th>Description</th></tr></thead><tbody><tr><td>storeId</td><td>number</td><td>Ecwid store ID.</td></tr><tr><td>orderId</td><td>string</td><td>Order ID. Can contain prefixes and suffixes, for example: <code>EG4H2,J77J8,SALE-G01ZG</code></td></tr></tbody></table>

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

<table><thead><tr><th width="268">Field</th><th width="133">Type</th><th>Description</th></tr></thead><tbody><tr><td>id</td><td>string</td><td>Unique order extra field ID.</td></tr><tr><td>value</td><td>string</td><td>Value saved to the order extra field.</td></tr><tr><td>customerInputType</td><td>string</td><td>Extra field type.<br><br>One of:<br><code>TEXT</code><br><code>TEXTAREA</code><br><code>SELECT</code><br><code>CHECKBOX</code><br><code>TOGGLE_BUTTON_GROUP</code><br><code>RADIO_BUTTONS</code><br><code>DATETIME</code><br><code>LABEL</code></td></tr><tr><td>title</td><td>string</td><td>Name visible at the checkout above the extra field.</td></tr><tr><td>orderDetailsDisplaySection</td><td>string</td><td>Defines where on the order details page the extra field is shown to the store owner.<br><br>One of:<br><code>shipping_info</code> - Order shipping details.<br><code>billing_info</code> - Order payment details.<br><code>customer_info</code> - Details about the customer.<br><code>order_comments</code> - Order comments left by the customer.</td></tr><tr><td>orderBy</td><td>string</td><td>Number that defines the extra field position in Ecwid admin. <br><br>The smaller the number, the higher the position is. Starts with <code>"0"</code> and iterates by 1.</td></tr><tr><td>showInNotifications</td><td>boolean</td><td>Defines if extra field should be visible in order emails sent to the customer. Disabled by default (<code>false</code>).<br><br>The <code>orderDetailsDisplaySection</code> value defines where the extra field will appear.</td></tr><tr><td>showInInvoice</td><td>boolean</td><td>Defines if the extra field should be visible in order tax invoices.</td></tr></tbody></table>
