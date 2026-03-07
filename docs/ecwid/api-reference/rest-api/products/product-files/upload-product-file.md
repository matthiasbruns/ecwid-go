# Upload product file

<mark style="color:blue;">`POST`</mark> `https://app.ecwid.com/api/v3/{storeId}/products/{productId}/files`&#x20;

### Required access scopes

Your app must have the following **access scopes** to make this request: `create_catalog`

### Path params

All path params are required.

| Param     | Type   | Description          |
| --------- | ------ | -------------------- |
| storeId   | number | Ecwid store ID.      |
| productId | number | Internal product ID. |

### Query params

Some query params are required.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>fileName</td><td>string</td><td>Name of the file visible to clients who bought the product with this file.<br><br><strong>Required</strong></td></tr><tr><td>description</td><td>string</td><td>Text description visible to clients who bought the product with this file.</td></tr><tr><td>externalUrl</td><td>string</td><td>HTTPS link to the file that must be accessible by Ecwid servers. After uploading the file, Ecwid generates its own link for customers, so you may delete the file from your server.<br><br><strong>Required</strong></td></tr></tbody></table>

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Field | Type   | Description                       |
| ----- | ------ | --------------------------------- |
| id    | number | Internal ID of the uploaded file. |
