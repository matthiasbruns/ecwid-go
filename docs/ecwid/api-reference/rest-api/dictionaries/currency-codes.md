# Currency codes

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/currencies?lang=en`

### Required access scopes

Your app doesn't need any specific **access scopes** to make this request.

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

All query params are optional.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>lang</td><td>string</td><td>Language ISO code for translations in JSON response, e.g. <code>en</code>, <code>fr</code>.<br><br>If not specified, response will be in English.</td></tr></tbody></table>

### Headers

The **Authorization** header is optional. Request works the same way with or without access token.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON **array of objects** with the following fields:

| Field | Type   | Description                                   |
| ----- | ------ | --------------------------------------------- |
| code  | string | Three-digit currency code in ISO 4217 format. |
| name  | string | Currency name in the specified language.      |
