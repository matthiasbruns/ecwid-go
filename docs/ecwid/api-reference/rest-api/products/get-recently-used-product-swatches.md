# Get recently used product swatches

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/swatches`&#x20;

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_catalog`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

<table data-full-width="false"><thead><tr><th width="149.62109375">Name</th><th width="149.71484375">Type</th><th>Description</th></tr></thead><tbody><tr><td>colors</td><td>array of objects <a data-mention href="#colors">#colors</a></td><td>List of the <code>SWATCHES</code> product options recently used by the store owner. <br><br>It is formed from the products with the <code>SWATCHES</code> option and sorted by the update time.</td></tr></tbody></table>

#### colors

<table data-full-width="false"><thead><tr><th width="149.62109375">Name</th><th width="149.71484375">Type</th><th>Description</th></tr></thead><tbody><tr><td>name</td><td>string</td><td>Name of the <code>SWATCHES</code> product option. This name should describe the color, for example, "Red" or "Dark Green".</td></tr><tr><td>hexCode</td><td>string</td><td>HEX code that defines what color must be displayed when a user selects the related product <code>SWATCHES</code> option, for example: <code>["#fff000"]</code>.</td></tr><tr><td>translations</td><td>object <a href="#translations">translations</a></td><td>Available translations for the <code>SWATCHES</code> option name.</td></tr></tbody></table>

#### translations

Object with text field translations in the `"lang": "text"` format, where the `"lang"` is an ISO 639-1 language code. For example:

```
{
    "en": "Sample text",
    "nl": "Voorbeeldtekst"
}
```

Translations are available for all active store languages. Only the default language translations are returned if no other translations are provided for the field. Find active store languages with <mark style="color:green;">`GET`</mark> `/profile` request > `languages` > `enabledLanguages`.
