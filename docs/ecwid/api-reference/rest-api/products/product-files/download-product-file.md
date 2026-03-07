# Download product file

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/products/{productId}/files/{fileId}`&#x20;

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_catalog`

### Path params

All path params are required.

| Param     | Type   | Description               |
| --------- | ------ | ------------------------- |
| storeId   | number | Ecwid store ID.           |
| productId | number | Internal product ID.      |
| fileId    | number | Internal product file ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response

Product file in the same format it was uploaded.
