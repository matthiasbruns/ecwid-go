# Search categories by path

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/categoriesByPath?path={path}&delimeter={delimeter}`

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_catalog`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

Most query params are optional, only two are required: `path` and `delimeter`.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>path</td><td>string</td><td>Full path separated by a delimiter. Spaces around the delimiter and empty path elements are ignored. <br><br><strong>Required param</strong></td></tr><tr><td>delimeter</td><td>string</td><td>A string of 1 or more characters used as a path delimiter. <br><br><strong>Required param</strong></td></tr><tr><td>keyword</td><td>string</td><td>Search term for category name and description.</td></tr><tr><td>parent</td><td>number</td><td>Parent category ID. If specified, the response will contain only categories that belong inside the specified one.</td></tr><tr><td>hidden_categories</td><td>boolean</td><td>Set <code>true</code> to include disabled categories.</td></tr><tr><td>baseUrl</td><td>string</td><td>Set base URL for URLs in response. <br><br>If not specified, Ecwid uses the store URL from general settings.</td></tr><tr><td>cleanURLs</td><td>boolean</td><td>Set <code>true</code> to force receiving clean URLs – catalog links without hashbang (<code>/#!/</code>). <br><br>If not specified, Ecwid checks URL settings automatically and responds with matching URLs.</td></tr><tr><td>slugsWithoutIds</td><td>boolean</td><td>Set <code>true</code>to receive category page links without IDs. <br><br>If not specified, Ecwid checks URL settings automatically and responds with matching URLs.</td></tr><tr><td>offset</td><td>number</td><td>Offset from the beginning of the returned items list. Used when the response contains more items than <code>limit</code> allows to receive in one request.<br><br>Usually used to receive all items in several requests with multiple of a hundred, for example:<br><br><code>?offset=0</code> for the first request,<br><code>?offset=100</code>, for the second request,<br><code>?offset=200</code>, for the third request, etc.</td></tr><tr><td>limit</td><td>number</td><td>Limit to the number of returned items. Maximum and default value (if not specified) is <code>100</code>.</td></tr><tr><td>lang</td><td>string</td><td>Language ISO code for translations in JSON response, e.g. <code>en</code>, <code>fr</code>. Translates fields like: <code>title</code>, <code>description</code>, <code>pickupInstruction</code>, <code>text</code>, etc.<br><br>If not specified, the default store language is used.</td></tr><tr><td>responseFields</td><td>string</td><td>Specify the exact fields to receive in response JSON. If not specified, the response JSON will have all available fields for the entity.<br><br>For example: <code>?responseFields=total,items(id,name,enabled)</code></td></tr></tbody></table>

Example of using `responseFields` param:

{% tabs %}
{% tab title="Request" %}

```
curl --location 'https://app.ecwid.com/api/v3/1003/categories?responseFields=total,items(id,name,enabled)' \
--header 'Authorization: Bearer secret_ab***cd'
```

{% endtab %}

{% tab title="Response" %}

```json
{
    "total": 1,
    "items": [
        {
            "id": 172786255,
            "name": "Best Toys",
            "enabled": true
        }
    ]
}
```

{% endtab %}
{% endtabs %}

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
| items  | array of objects [items](#items) | Detailed information about returned categories.                                              |

#### items

| Field                   | Type                                   | Description                                                                                                                                                                                                 |
| ----------------------- | -------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| id                      | number                                 | Internal unique category ID.                                                                                                                                                                                |
| parentId                | number                                 | ID of the parent category, if any.                                                                                                                                                                          |
| orderBy                 | number                                 | Sorting order of the category. Starts from `10` and increments by `10`.                                                                                                                                     |
| hdThumbnailUrl          | string                                 | Link to the category image resized to fit 800x800px container.                                                                                                                                              |
| thumbnailUrl            | string                                 | Link to the category image resized to fit 400x400px container.                                                                                                                                              |
| imageUrl                | string                                 | Link to the category image resized to fit 1200x1200px container.                                                                                                                                            |
| originalImageUrl        | string                                 | Link to the full-sized category image.                                                                                                                                                                      |
| imageExternalId         | string                                 | <p>Image ID for Lightspeed R-Series/X-Series image sync.<br><br><strong>Read-only</strong></p>                                                                                                              |
| name                    | string                                 | Category name visible on the storefront.                                                                                                                                                                    |
| nameTranslated          | object [translations](#translations)   | Available translations for the category name.                                                                                                                                                               |
| originalImage           | object [originalImage](#originalimage) | Details of the category image.                                                                                                                                                                              |
| thumbnail               | object [thumbnail](#thumbnail)         | Details of the category thumbnail.                                                                                                                                                                          |
| origin                  | string                                 | <p>Internal field that defines category origin inside Lightspeed. <br><br>One of:</p><p><code>LIGHTSPEED</code></p><p><code>ECWID</code></p><p><code>X-SERIES</code> <br><br><strong>Read-only</strong></p> |
| url                     | string                                 | Full URL of the category page on the streofront.                                                                                                                                                            |
| autogeneratedSlug       | string                                 | Autogenerated slug for the category page URL.                                                                                                                                                               |
| customSlug              | string                                 | Custom slug for the category page URL.                                                                                                                                                                      |
| productCount            | number                                 | Number of products in the category and its subcategories. When the product count in the category changes, `productCount` value will update in several minutes.                                              |
| enabledProductCount     | number                                 | Number of enabled products in the category (excluding any subcategories). Requires the `productIds=true` query param.                                                                                       |
| description             | string                                 | Category description in HTML format.                                                                                                                                                                        |
| descriptionTranslated   | object [translations](#translations)   | Available translations for the category description.                                                                                                                                                        |
| enabled                 | boolean                                | `true` if the category is enabled, `false` otherwise. Use `hidden_categories` in request to get disabled categories                                                                                         |
| productIds              | array of numbers                       | IDs of products assigned to the category as they appear in [Ecwid admin > Catalog > Categories](https://my.ecwid.com/#category). Requires `productIds=true` query param.                                    |
| seoTitle                | string                                 | SEO page title for web search results. Recommended length is under 55 characters.                                                                                                                           |
| seoTitleTranslated      | string                                 | SEO page title translations.                                                                                                                                                                                |
| seoDescription          | string                                 | SEO page description for web search results. Recommended length is under 160 characters.                                                                                                                    |
| seoDecriptionTranslated | string                                 | SEO page description translations.                                                                                                                                                                          |
| alt                     | object [alt](#alt)                     | Alt texts of a category image.                                                                                                                                                                              |
| externalReferenceId     | string                                 | <p>Internal field for Lightspeed X-Series connection.<br><br>This ID is unique for each category in one store.</p>                                                                                          |

#### originalImage

| Field  | Type    | Description  |
| ------ | ------- | ------------ |
| url    | string  | Image URL    |
| width  | integer | Image width  |
| height | integer | Image height |

#### thumbnail

| Field  | Type    | Description  |
| ------ | ------- | ------------ |
| url    | string  | Image URL    |
| width  | integer | Image width  |
| height | integer | Image height |

#### alt

<table><thead><tr><th>Field</th><th width="182">Type</th><th>Description</th></tr></thead><tbody><tr><td>main</td><td>string</td><td>Image description for the "alt" HTML attribute of the image.</td></tr><tr><td>translations</td><td>object <a href="#translations">translations</a></td><td>Available translations for the "alt" text.</td></tr></tbody></table>

#### translations

Object with text field translations in the `"lang": "text"` format, where the `"lang"` is an ISO 639-1 language code. For example:

```
{
    "en": "Sample text",
    "nl": "Voorbeeldtekst"
}
```

Translations are available for all active store languages. Only the default language translations are returned if no other translations are provided for the field. Find active store languages with <mark style="color:green;">`GET`</mark> `/profile` request > `languages` > `enabledLanguages`.
