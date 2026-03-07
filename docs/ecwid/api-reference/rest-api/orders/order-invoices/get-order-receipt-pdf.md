# Get order receipt PDF

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/orders/{orderId}/invoice-pdf`&#x20;

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_orders`

### Path params

All path params are required.

<table><thead><tr><th width="150">Param</th><th width="126">Type</th><th>Description</th></tr></thead><tbody><tr><td>storeId</td><td>number</td><td>Ecwid store ID.</td></tr><tr><td>orderId</td><td>number</td><td>Order ID. Can contain prefixes and suffixes, for example: <code>EG4H2,J77J8,SALE-G01ZG</code></td></tr></tbody></table>

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response

Tax invoice file in PDF format.
