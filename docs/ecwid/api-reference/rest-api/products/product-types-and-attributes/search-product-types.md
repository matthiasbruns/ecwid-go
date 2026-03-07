# Search product types

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/classes`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
GET /api/v3/1003/classes HTTP/1.1
Host: app.ecwid.com
Authorization: Bearer secret_token
```

Response:

{% code fullWidth="true" %}

```json
[
  {
    "id": 0,
    "attributes": [
      {
        "id": 139165261,
        "name": "Units in product",
        "type": "UNITS_IN_PRODUCT",
        "show": "DESCR"
      },
      {
        "id": 82991001,
        "name": "Price per unit",
        "type": "PRICE_PER_UNIT",
        "show": "PRICE"
      },
      {
        "id": 201437969,
        "name": "UPC",
        "type": "UPC",
        "show": "DESCR"
      },
      {
        "id": 201437970,
        "name": "Brand",
        "type": "BRAND",
        "show": "DESCR"
      }
    ]
  }
]
```

{% endcode %}

</details>

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

A JSON array of objects with the following fields:

<table><thead><tr><th width="179">Field</th><th width="229">Type</th><th>Description</th></tr></thead><tbody><tr><td>id</td><td>number</td><td>Internal unique ID of the product type. <br><br>By default, all products get the "General" type which ID is <code>0</code>.</td></tr><tr><td>name</td><td>string</td><td>Product type name. Empty for the "General" type.</td></tr><tr><td>googleTaxonomy</td><td>string</td><td>Google taxonomy associated with the type.</td></tr><tr><td>attributes</td><td>array of objects <a href="#attributes">attributes</a></td><td>Product attributes assigned to this product type.</td></tr></tbody></table>

#### attributes

<table><thead><tr><th width="181">Field</th><th width="183">Type</th><th>Description</th></tr></thead><tbody><tr><td>id</td><td>number</td><td>Internal unique ID of the product attribute.</td></tr><tr><td>name</td><td>string</td><td>Attribute title visible on the storefront if it has some value for the product.</td></tr><tr><td>nameTranslated</td><td>object <a href="#translations">translations</a></td><td>Available translations for product attribute name.</td></tr><tr><td>type</td><td>string</td><td><p>Type of the attribute that defines its functionality. </p><p></p><p>One of:</p><p><code>CUSTOM</code></p><p><code>UPC</code></p><p><code>BRAND</code></p><p><code>GENDER</code></p><p><code>AGE_GROUP</code></p><p><code>COLOR</code></p><p><code>SIZE</code> </p><p><code>TAGS</code></p><p><code>PRICE_PER_UNIT</code></p><p><code>UNITS_IN_PRODUCT</code><br><br>Attributes of type <code>PRICE_PER_UNIT</code> and <code>UNITS_IN_PRODUCT</code> are only available if the price-per-unit feature is enabled.</p></td></tr><tr><td>show</td><td>string</td><td>Defines if an attribute is visible on product pages. <br><br>One of: <br><code>NOTSHOW</code><br><code>DESCR</code><br><code>PRICE</code>. <br><br>The value <code>PRICE</code> = <code>DESCR</code>. Request can use <a href="ref:authentication-basics#access-tokens">public tokens</a>, but <code>NOTSHOW</code> attributes won't be returned in that case.</td></tr></tbody></table>

#### translations

Object with text field translations in the `"lang": "text"` format, where the `"lang"` is an ISO 639-1 language code. For example:

```
{
    "en": "Sample text",
    "nl": "Voorbeeldtekst"
}
```

Translations are available for all active store languages. Only the default language translations are returned if no other translations are provided for the field. Find active store languages with <mark style="color:green;">`GET`</mark> `/profile` request > `languages` > `enabledLanguages`.
