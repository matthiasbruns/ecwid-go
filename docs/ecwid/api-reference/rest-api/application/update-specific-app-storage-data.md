# Update specific app storage data

<mark style="color:purple;">`PUT`</mark> `https://app.ecwid.com/api/v3/{storeId}/storage/{key}`

### Required access scopes

Your app must have the following **access scopes** to make this request: `update_store_profile`

### Path params

All path params are required.

| Param   | Type   | Description                                     |
| ------- | ------ | ----------------------------------------------- |
| storeId | number | Ecwid store ID.                                 |
| key     | string | **Key** for updating specific app storage data. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Request JSON

A JSON object with the following fields:

| Field | Type   | Description                                                                                                                                       |
| ----- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| value | string | <p>Data for saving in the app storage. <br><br>Size limit: <br>Private keys – <strong>1MB</strong></p><p>Public keys - <strong>256KB</strong></p> |

### Response JSON

A JSON object with the following fields:

| Field   | Type    | Description                            |
| ------- | ------- | -------------------------------------- |
| success | boolean | Defines if the request was successful. |
