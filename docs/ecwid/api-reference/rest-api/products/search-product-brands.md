# Search product brands

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/brands`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
GET /api/v3/1003/brands HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
```

Response:

```json
{
  "total": 4,
  "limit": 0,
  "count": 4,
  "offset": 0,
  "items": [
    {
      "name": "BRANDING3",
      "productsFilteredByBrandUrl": "https://example.company.site/products/search?attribute_Brand=BRANDING3"
    },
    {
      "name": "Branding4",
      "productsFilteredByBrandUrl": "https://example.company.site/products/search?attribute_Brand=Branding4"
    },
    {
      "name": "HIDDEN ONE",
      "productsFilteredByBrandUrl": "https://example.company.site/products/search?attribute_Brand=HIDDEN+ONE"
    },
    {
      "name": "branding2",
      "productsFilteredByBrandUrl": "https://example.company.site/products/search?attribute_Brand=branding2"
    }
  ]
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_brands`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

All query params are optional.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>limit</td><td>number</td><td>Limit to the number of returned items. Maximum and default value (if not specified) is <code>100</code>.</td></tr><tr><td>offset</td><td>number</td><td>Offset from the beginning of the returned items list. Used when the response contains more items than <code>limit</code> allows to receive in one request.<br><br>Usually used to receive all items in several requests with multiple of a hundred, for example:<br><br><code>?offset=0</code> for the first request,<br><code>?offset=100</code>, for the second request,<br><code>?offset=200</code>, for the third request, etc.</td></tr><tr><td>sortBy</td><td>string</td><td><p>Sorting order for the results. <br><br>One of:</p><p><code>PRODUCT_COUNT_DESC</code> (default)</p><p><code>PRODUCT_COUNT_ASC</code></p><p><code>NAME_DESC</code></p><p><code>NAME_ASC</code></p></td></tr><tr><td>hidden_brands</td><td>boolean</td><td>Defines if the response should contain brands of disabled products. Set <code>true</code> to get such brands. <br><br>Default value: <code>false</code></td></tr><tr><td>lang</td><td>string</td><td>Language ISO code for translations in JSON response, e.g. <code>en</code>, <code>fr</code>. Translates fields like: <code>title</code>, <code>description</code>, <code>pickupInstruction</code>, <code>text</code>, etc.<br><br>If not specified, the default store language is used.</td></tr><tr><td>baseUrl</td><td>string</td><td>Set base URL for URLs in response. <br><br>If not specified, Ecwid will use the main URL from store settings.</td></tr><tr><td>cleanURLs</td><td>boolean</td><td>Set <code>true</code> to force receiving clean URLs – catalog links without hashbang (<code>/#!/</code>). <br><br>By default Ecwid checks if this setting is enabled for the store and responds with matching URLs.</td></tr></tbody></table>

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Field  | Type                             | Description                                                                                  |
| ------ | -------------------------------- | -------------------------------------------------------------------------------------------- |
| total  | number                           | Total number of found items (might be more than the number of returned items).               |
| count  | number                           | Total number of items returned in the response.                                              |
| offset | number                           | Offset from the beginning of the returned items list specified in the request.               |
| limit  | number                           | Maximum number of returned items specified in the request. Maximum and default value: `100`. |
| items  | array of objects [items](#items) | Detailed information about returned brands.                                                  |

#### items

<table><thead><tr><th width="264">Field</th><th>Type</th><th>Description</th></tr></thead><tbody><tr><td>name</td><td>string</td><td>Brand name visible on the storefront.</td></tr><tr><td>nameTranslated</td><td>object <a href="#translations">translations</a></td><td>Available translations for the brand name.</td></tr><tr><td>productsFilteredByBrandUrl</td><td>string</td><td>Link to your website's search page with applied brand filter.<br><br>Page will show all products related to the brand.</td></tr></tbody></table>

#### translations

Object with text field translations in the `"lang": "text"` format, where the `"lang"` is an ISO 639-1 language code. For example:

```
{
    "en": "Sample text",
    "nl": "Voorbeeldtekst"
}
```

Translations are available for all active store languages. Only the default language translations are returned if no other translations are provided for the field. Find active store languages with <mark style="color:green;">`GET`</mark> `/profile` request > `languages` > `enabledLanguages`.
