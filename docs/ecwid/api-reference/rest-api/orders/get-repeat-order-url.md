# Get repeat order URL

Receive a URL that allows you to repeat the order with all the same details. After receiving the link, send it to the customer, so they can repeat the order without filling the cart or any of the checkout details.\
\
Read more about the repeated orders in [Help Center](https://support.ecwid.com/hc/en-us/articles/4404924373138-Taking-repeat-orders).&#x20;

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/orders/{orderId}/repeatURL`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
GET /api/v3/1003/orders/JJ5HH/repeatURL HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
```

Response:

```json
{
  "repeatOrderUrl":"https://example-store.ecwid.com/repeat-order?id=JJ5HH&token=v***y4&type=order&utm_source=repeat_order_button&utm_medium=api"
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_orders`

### Path params

All path params are required.

<table><thead><tr><th>Param</th><th width="170">Type</th><th>Description</th></tr></thead><tbody><tr><td>storeId</td><td>number</td><td>Ecwid store ID.</td></tr><tr><td>orderId</td><td>number</td><td>Order ID. Can contain prefixes and suffixes, for example: <code>EG4H2,J77J8,SALE-G01ZG</code></td></tr></tbody></table>

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Field          | Type   | Description                                                                                            |
| -------------- | ------ | ------------------------------------------------------------------------------------------------------ |
| repeatOrderUrl | string | Link that allows customers to repeat the order they made without filling the cart with products again. |
