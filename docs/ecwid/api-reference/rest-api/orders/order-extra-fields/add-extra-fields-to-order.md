# Add extra fields to order

<mark style="color:blue;">`POST`</mark> `https://app.ecwid.com/api/v3/{storeId}/orders/{orderId}/extraFields`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
POST /api/v3/1003/orders/JJ5HH/extraFields/tips HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
Content-Type: application/json
Cache-Control: no-cache

{
    "id": "reference_number",
    "value": "#334-3340-1",
    "customerInputType": "TEXT",
    "title": "Affiliate number",
    "orderDetailsDisplaySection": "billing_info",
    "showInNotifications": false,
    "showInInvoice": false
}
```

Response:

<pre class="language-json"><code class="lang-json">{
<strong>    "createCount": 1
</strong>}
</code></pre>

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `update_orders`

### Path params

All path params are required.

<table><thead><tr><th width="141">Param</th><th width="152">Type</th><th>Description</th></tr></thead><tbody><tr><td>storeId</td><td>number</td><td>Ecwid store ID.</td></tr><tr><td>orderId</td><td>string</td><td>Order ID. Can contain prefixes and suffixes, for example: <code>EG4H2,J77J8,SALE-G01ZG</code></td></tr></tbody></table>

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Request JSON

A JSON object with the following fields:

<table><thead><tr><th width="249">Field</th><th width="125">Type</th><th>Description</th></tr></thead><tbody><tr><td>id</td><td>string</td><td>Unique order extra field ID.<br><br><strong>Required</strong></td></tr><tr><td>value</td><td>string</td><td>Value saved to the order extra field.<br><br><strong>Required</strong></td></tr><tr><td>customerInputType</td><td>string</td><td><p>Extra field type.<br><br>One of:</p><p><code>TEXT</code> (default, if not specified)</p><p><code>TEXTAREA</code></p><p><code>SELECT</code></p><p><code>CHECKBOX</code></p><p><code>TOGGLE_BUTTON_GROUP</code></p><p><code>RADIO_BUTTONS</code></p><p><code>DATETIME</code></p><p><code>LABEL</code></p></td></tr><tr><td>title</td><td>string</td><td>Name visible at the checkout above the extra field.<br><br><strong>Required</strong></td></tr><tr><td>orderDetailsDisplaySection</td><td>string</td><td>Defines where on the order details page the extra field is shown to the store owner.<br><br>One of:<br><code>shipping_info</code> - Order shipping details.<br><code>billing_info</code> - Order payment details.<br><code>customer_info</code> - Details about the customer.<br><code>order_comments</code> - Order comments left by the customer.</td></tr><tr><td>orderBy</td><td>string</td><td>Number that defines the extra field position in Ecwid admin. <br><br>The smaller the number, the higher the position is. Starts with <code>"0"</code> and iterates by 1.</td></tr><tr><td>showInNotifications</td><td>boolean</td><td>Defines if extra field should be visible in order emails sent to the customer. Disabled by default (<code>false</code>).<br><br>The <code>orderDetailsDisplaySection</code> value defines where the extra field will appear.</td></tr><tr><td>showInInvoice</td><td>boolean</td><td>Defines if the extra field should be visible in order tax invoices.</td></tr></tbody></table>
