# Update abandoned cart

<mark style="color:purple;">`PUT`</mark> `https://app.ecwid.com/api/v3/{storeId}/carts/{cartId}`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
PUT /api/v3/1003/carts/6626E60A-A6F9-4CD5-8230-43D5F162E0CD HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
Content-Type: application/json
Cache-Control: no-cache

{
  "hidden": true,
  "taxesOnShipping": [
    {
      "name": "Tax X",
      "value": 20,
      "total": 2.86
    }
  ]
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

Your app must have the following **access scopes** to make this request: `update_orders`

### Path params

All path params are required.

<table><thead><tr><th width="240.30859375">Param</th><th width="220.45703125">Type</th><th>Description</th></tr></thead><tbody><tr><td>storeId</td><td>number</td><td>Ecwid store ID.</td></tr><tr><td>cartId</td><td>string</td><td>Internal cart ID, for example, <code>0B299518-FB54-491A-A6E0-5B6BA6E20868</code>.</td></tr></tbody></table>

### Headers

The **Authorization** header is required.

<table><thead><tr><th width="240.19921875">Header</th><th width="220.03515625">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Request JSON

A JSON object with the following fields:

<table><thead><tr><th width="240.23046875">Field</th><th width="220.4609375">Type</th><th>Description</th></tr></thead><tbody><tr><td>hidden</td><td>boolean</td><td>Determines if the order is hidden (removed from the list). Applies to abandoned carts only</td></tr><tr><td>taxesOnShipping</td><td>array of objects <a href="#taxesonshipping">taxesOnShipping</a></td><td>Taxes applied to shipping. <code>null</code> for old orders, <code>[]</code> for orders with taxes applied to subtotal only. Are not recalculated if cart is updated later manually. Is calculated like: <code>(shippingRate + handlingFee)*(taxValue/100)</code></td></tr><tr><td>b2b_b2c</td><td>string</td><td>Order type: business-to-business (<code>b2b</code>) or business-to-consumer (<code>b2c</code>)</td></tr><tr><td>customerRequestedInvoice</td><td>boolean</td><td><code>true</code> if customer requested an invoice, <code>false</code> otherwise</td></tr><tr><td>customerFiscalCode</td><td>string</td><td>Fiscale code of a customer</td></tr><tr><td>electronicInvoicePecEmail</td><td>string</td><td>PEC email for invoices</td></tr><tr><td>electronicInvoiceSdiCode</td><td>string</td><td>SDI code for invoices</td></tr></tbody></table>

#### taxesOnShipping

<table><thead><tr><th width="240.140625">Field</th><th width="214.9453125">Type</th><th>Description</th></tr></thead><tbody><tr><td>name</td><td>string</td><td>Tax name</td></tr><tr><td>value</td><td>number</td><td>Tax value in store settings, applied to destination zone</td></tr><tr><td>total</td><td>number</td><td>Tax total applied to shipping</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

<table><thead><tr><th width="239.75390625">Field</th><th width="220.125">Type</th><th>Description</th></tr></thead><tbody><tr><td>updateCount</td><td>number</td><td><p>The number of updated items that defines if the request was successful.<br><br>One of:</p><p><code>1</code> if the item was updated,</p><p><code>0</code> if the item was not updated.</p></td></tr></tbody></table>
