# Delete all products

<mark style="color:red;">`DELETE`</mark> `https://app.ecwid.com/api/v3/{storeId}/products`&#x20;

### Required access scopes

Your app must have the following **access scopes** to make this request: `update_catalog_batch_delete`

### Path params

All path params are required.

| Param     | Type   | Description          |
| --------- | ------ | -------------------- |
| storeId   | number | Ecwid store ID.      |
| productId | number | Internal product ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response

`202 Accepted` HTTP status code.
