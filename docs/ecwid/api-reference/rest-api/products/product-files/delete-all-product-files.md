# Delete all product files

<mark style="color:red;">`DELETE`</mark> `https://app.ecwid.com/api/v3/{storeId}/products/{productId}/files`&#x20;

### Required access scopes

Your app must have the following **access scopes** to make this request: `update_catalog`

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

A JSON object with the following fields:

<table><thead><tr><th width="239.5625">Field</th><th width="165.1328125">Type</th><th>Description</th></tr></thead><tbody><tr><td>deleteCount</td><td>number</td><td><p>The number of deleted items that defines if the request was successful.<br><br>One of:</p><p><code>1</code> or more if items were deleted,</p><p><code>0</code> if no items were deleted.</p></td></tr></tbody></table>
