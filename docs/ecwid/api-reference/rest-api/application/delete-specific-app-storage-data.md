# Delete specific app storage data

<mark style="color:red;">`DELETE`</mark> `https://app.ecwid.com/api/v3/{storeId}/storage/{key}`

### Required access scopes

Your app must have the following **access scopes** to make this request: `update_store_profile`

### Path params

All path params are required.

| Param   | Type   | Description                                          |
| ------- | ------ | ---------------------------------------------------- |
| storeId | number | Ecwid store ID.                                      |
| key     | string | **Key** for deleting specific data from app storage. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Field   | Type    | Description                            |
| ------- | ------- | -------------------------------------- |
| success | boolean | Defines if the request was successful. |
