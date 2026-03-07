# Currency codes by country

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/currencyByCountry?countryCode=US&lang=en`

### Required access scopes

Your app doesn't need any specific **access scopes** to make this request.

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

Some query params are required.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>lang</td><td>string</td><td>Language ISO code for translations in JSON response, e.g. <code>en</code>, <code>fr</code>.<br><br>If not specified, response will be in English.</td></tr><tr><td>countryCode</td><td>string</td><td>Country code in ISO 639-1.<br><br><strong>Required</strong></td></tr></tbody></table>

### Headers

The **Authorization** header is optional. Request works the same way with or without access token.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON **array of objects** with the following fields:

| Field         | Type   | Description                                     |
| ------------- | ------ | ----------------------------------------------- |
| countryCode   | string | Two-digit country code in ISO 639-1 format.     |
| currencyCode  | string | Three-digit currency code in ISO 4217 format.   |
| currencyName  | string | Currency name in specified language             |
| prefix        | string | Currency prefix symbol                          |
| suffix        | string | Currency suffix symbol                          |
| decimalPlaces | string | Quantity of decimal places for currency subunit |
