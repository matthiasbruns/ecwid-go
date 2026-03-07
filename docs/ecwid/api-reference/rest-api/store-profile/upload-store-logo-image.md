# Upload store logo image

<mark style="color:purple;">`PUT`</mark> `https://app.ecwid.com/api/v3/{storeId}/profile/{logo}?externalUrl={externalUrl}`

### Required access scopes

Requires the following access scope: `update_store_profile`

### Path params

All path params are required.

| Param   | Type   | Description                                                                                                                                                                                                                                         |
| ------- | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| storeId | number | Ecwid store ID.                                                                                                                                                                                                                                     |
| logo    | string | <p>Type of logo you want to update. Must be one of:<br><code>logo</code>  - Main store logo visible on the storefront.<br><code>invoicelogo</code> - Logo image for invoices/receipts.<br><code>emaillogo</code> - Logo for other store emails.</p> |

### Query params

All query params are optional.

| Param       | Type   | Description                              |
| ----------- | ------ | ---------------------------------------- |
| externalUrl | string | Image URL available for public download. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Field | Type   | Description        |
| ----- | ------ | ------------------ |
| id    | number | Internal image ID. |
