# Create category

<mark style="color:blue;">`POST`</mark> `https://app.ecwid.com/api/v3/{storeId}/categories`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
POST /api/v3/1003/categories HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
Content-Type: application/json
Cache-Control: no-cache

{
  "name": "Lemons",
  "enabled": true,
  "orderBy": 10,
  "parentId": 9691094
}
```

Response:

```json
{
  "id": 10869029
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `create_catalog`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Request JSON

A JSON object with the following fields:

| Field                   | Type                                 | Description                                                                                                                                                              |
| ----------------------- | ------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| parentId                | number                               | <p>ID of the parent category.<br></p><p>If not specified, new category will be assigned with <code>parentId: 0</code> which is the main store category.</p>              |
| productIds              | array of numbers                     | IDs of products assigned to the category as they appear in [Ecwid admin > Catalog > Categories](https://my.ecwid.com/#category). Requires `productIds=true` query param. |
| orderBy                 | number                               | Sorting order of the category. Starts from `10` and increments by `10`.                                                                                                  |
| name                    | string                               | <p>Category name visible on the storefront.<br><br><strong>Required</strong></p>                                                                                         |
| nameTranslated          | object [translations](#translations) | Available translations for the category name.                                                                                                                            |
| description             | string                               | Category description in HTML format.                                                                                                                                     |
| descriptionTranslated   | object [translations](#translations) | Available translations for the category description.                                                                                                                     |
| seoTitle                | string                               | SEO page title for web search results. Recommended length is under 55 characters.                                                                                        |
| seoTitleTranslated      | string                               | Available translations for the SEO page title.                                                                                                                           |
| seoDescription          | string                               | SEO page description for web search results. Recommended length is under 160 characters.                                                                                 |
| seoDecriptionTranslated | string                               | Available translations for the SEO page description.                                                                                                                     |
| enabled                 | boolean                              | `true` if the category is enabled, `false` otherwise. Use `hidden_categories` in request to get disabled categories                                                      |
| customSlug              | string                               | Custom slug for the category page URL.                                                                                                                                   |
| externalReferenceId     | string                               | <p>Internal field for Lightspeed X-Series connection.<br><br>This ID is unique for each category in one store.</p>                                                       |

#### translations

Object with text field translations in the `"lang": "text"` format, where the `"lang"` is an ISO 639-1 language code. For example:

```
{
    "en": "Sample text",
    "nl": "Voorbeeldtekst"
}
```

Translations are available for all active store languages. Only the default language translations are returned if no other translations are provided for the field. Find active store languages with <mark style="color:green;">`GET`</mark> `/profile` request > `languages` > `enabledLanguages`.

### Response JSON

A JSON object with the following fields:

| Field | Type   | Description                          |
| ----- | ------ | ------------------------------------ |
| id    | number | Internal ID of the created category. |
