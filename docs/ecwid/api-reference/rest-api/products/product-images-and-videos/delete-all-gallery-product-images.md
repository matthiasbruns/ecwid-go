# Delete all gallery product images

<mark style="color:red;">`DELETE`</mark> `https://app.ecwid.com/api/v3/{storeId}/products/{productId}/gallery`

### Required access scopes

Your app must have the following **access scopes** to make this request: `update_catalog`, `update_catalog_batch_delete`

### Path params

All path params are required.

| Param     | Type   | Description          |
| --------- | ------ | -------------------- |
| storeId   | number | Ecwid store ID.      |
| productId | number | Internal product ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Field       | Type   | Description                                                                                                                                                                                |
| ----------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| deleteCount | number | <p>The number of deleted items that defines if the request was successful.<br><br>One of:<br><code>1</code> or more if items were deleted,<br><code>0</code> if no items were deleted.</p> |
