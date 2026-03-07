# Update order extra field

<mark style="color:purple;">`PUT`</mark> `https://app.ecwid.com/api/v3/{storeId}/orders/{orderId}/extraFields/{extraFieldId}`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
PUT /api/v3/1003/orders/JJ5HH/extraFields/tips HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
Content-Type: application/json
Cache-Control: no-cache

{
    "value": "4%",
    "showInInvoice": true,
    "showInNotifications": false
}
```

Response:

<pre class="language-json"><code class="lang-json">{
<strong>    "updateCount": 1
</strong>}
</code></pre>

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `update_orders`

### Path params

All path params are required.

<table><thead><tr><th width="141">Param</th><th width="152">Type</th><th>Description</th></tr></thead><tbody><tr><td>storeId</td><td>number</td><td>Ecwid store ID.</td></tr><tr><td>orderId</td><td>string</td><td>Order ID. Can contain prefixes and suffixes, for example: <code>EG4H2,J77J8,SALE-G01ZG</code></td></tr><tr><td>extraFieldId</td><td>string</td><td>ID of the extra field.</td></tr></tbody></table>

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Request JSON

A JSON object with the following fields:

<table><thead><tr><th width="255">Field</th><th width="104">Type</th><th>Description</th></tr></thead><tbody><tr><td>value</td><td>string</td><td>Value for the order extra field.</td></tr><tr><td>title</td><td>string</td><td>Name visible at the checkout above the extra field.</td></tr><tr><td>orderDetailsDisplaySection</td><td>string</td><td><p>Defines where on the order details page the extra field is shown to the store owner.<br><br>One of:</p><p><code>shipping_info</code> - Order shipping details.</p><p><code>billing_info</code> - Order payment details.</p><p><code>customer_info</code> - Details about the customer.</p><p><code>order_comments</code> - Order comments left by the customer.</p></td></tr><tr><td>orderBy</td><td>string</td><td>Number that defines the extra field position in Ecwid admin. <br><br>The smaller the number, the higher the position is. Starts with <code>"0"</code> and iterates by 1.</td></tr><tr><td>showInNotifications</td><td>boolean</td><td>Defines if extra field should be visible in order emails sent to the customer. Disabled by default (<code>false</code>).<br><br>The <code>orderDetailsDisplaySection</code> value defines where the extra field will appear.</td></tr><tr><td>showInInvoice</td><td>boolean</td><td>Defines if the extra field should be visible in order tax invoices.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Field       | Type   | Description                                                                                                                                                                                   |
| ----------- | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| updateCount | number | <p>The number of updated items that defines if the request was successful.<br><br>One of:</p><p><code>1</code> if the item was updated,</p><p><code>0</code> if the item was not updated.</p> |
