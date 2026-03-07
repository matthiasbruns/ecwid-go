# Change product file description

<mark style="color:purple;">`PUT`</mark> `https://app.ecwid.com/api/v3/{storeId}/products/{productId}/files/{fileId}`&#x20;

### Required access scopes

Your app must have the following **access scopes** to make this request: `update_catalog`

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

### Request JSON

A JSON object with the following fields:

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>description</td><td>string</td><td>Text description visible to clients who bought the product with this file.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

<table><thead><tr><th width="184.54296875">Field</th><th width="162.77734375">Type</th><th>Description</th></tr></thead><tbody><tr><td>updateCount</td><td>number</td><td><p>The number of updated items that defines if the request was successful.<br><br>One of:</p><p><code>1</code> if the item was updated,</p><p><code>0</code> if the item was not updated.</p></td></tr></tbody></table>
