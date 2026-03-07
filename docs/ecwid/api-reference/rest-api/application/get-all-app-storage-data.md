# Get all app storage data

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/storage`

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_store_profile`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON **array of objects** with the following fields:

| Field | Type   | Description                         |
| ----- | ------ | ----------------------------------- |
| key   | string | Key for the data stored by the app. |
| value | string | Data stored by the app.             |
