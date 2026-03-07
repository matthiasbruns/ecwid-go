# Country codes

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/countries?lang=en&withStates=true`

### Required access scopes

Your app doesn't need any specific **access scopes** to make this request.

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

All query params are optional.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>withStates</td><td>boolean</td><td>If <code>true</code>, adds <code>states</code> field for each country in the response.</td></tr><tr><td>lang</td><td>string</td><td>Language ISO code for translations in JSON response, e.g. <code>en</code>, <code>fr</code>.<br><br>If not specified, response will be in English.</td></tr></tbody></table>

### Headers

The **Authorization** header is optional. Request works the same way with or without access token.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON **array of objects** with the following fields:

| Field  | Type          | Description                                     |
| ------ | ------------- | ----------------------------------------------- |
| code   | string        | Two-digit country code in ISO 639-1.            |
| name   | string        | Country name in the specified language.         |
| states | object states | Details about states that beong to the country. |

#### states

| Field | Type   | Description                           |
| ----- | ------ | ------------------------------------- |
| code  | string | Two-digit state code in ISO 639-1.    |
| name  | string | State name in the specified language. |
