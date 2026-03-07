# Create product variation

<mark style="color:blue;">`POST`</mark> `https://app.ecwid.com/api/v3/{storeId}/products/{productId}/combinations`&#x20;

{% hint style="info" %}
If specified options do not yet exist, they will be created automatically as a dropdown list (`SELECT` option type).
{% endhint %}

<details>

<summary>Request and response example</summary>

Request:

```http
POST /api/v3/1003/products/692730761/combinations HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
Content-Type: application/json
Cache-Control: no-cache

{
    "options": [
        {
            "name": "Size",
            "value": "L",
        },
        {
            "name": "Color",
            "value": "Red",
        }
    ],
    "price": 10,
    "costPrice": 8,
    "quantity": 10,
    "inStock": true,
    "unlimited": false,
    "weight": 0.5,
    "compareToPrice": 15,
    "sku": "combination-sku",
    "attributes": [
        {
            "id": 9998010,
            "name": "UPC",
            "value": "0435943543594395",
            "show": "DESCR",
            "type": "UPC"
        }
    ],
    "isShippingRequired": true
}
```

Response:

{% code fullWidth="true" %}

```json
{
    "id": 397662764
}
```

{% endcode %}

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `create_catalog`

### Path params

All path params are required.

| Param     | Type   | Description          |
| --------- | ------ | -------------------- |
| storeId   | number | Ecwid store ID.      |
| productId | number | Internal product ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Request JSON

A JSON object with the following fields:

| Field                                | Type                                                 | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| ------------------------------------ | ---------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| options                              | array of objects [options](#options)                 | <p>Set of selected product option values that identify this variation.<br><br><strong>Required</strong></p>                                                                                                                                                                                                                                                                                                                                                  |
| sku                                  | string                                               | <p>Variation SKU.<br><br>If empty, variation inherits the base product's SKU.</p>                                                                                                                                                                                                                                                                                                                                                                            |
| thumbnailUrl                         | string                                               | Link to the variation image resized to fit 400x400px container.                                                                                                                                                                                                                                                                                                                                                                                              |
| imageUrl                             | string                                               | Link to the variation image resized to fit 1200x1200px container.                                                                                                                                                                                                                                                                                                                                                                                            |
| smallThumbnailUrl                    | string                                               | Link to the variation image resized to fit 160x160px container.                                                                                                                                                                                                                                                                                                                                                                                              |
| hdThumbnailUrl                       | string                                               | Link to the variation image resized to fit 800x800px container.                                                                                                                                                                                                                                                                                                                                                                                              |
| originalImageUrl                     | string                                               | Link to the full-sized variation image.                                                                                                                                                                                                                                                                                                                                                                                                                      |
| instock                              | boolean                                              | Defines if the variation is in stock (`quantity` is more than `0`).                                                                                                                                                                                                                                                                                                                                                                                          |
| quantity                             | number                                               | <p>Number of variation items in stock. </p><p></p><p>If the variation has unlimited stock (<code>unlimited</code> is <code>true</code>), this field is not returned.</p>                                                                                                                                                                                                                                                                                     |
| unlimited                            | boolean                                              | Defines if the variation has unlimited stock.                                                                                                                                                                                                                                                                                                                                                                                                                |
| price                                | number                                               | Base variation price without any modifiers.                                                                                                                                                                                                                                                                                                                                                                                                                  |
| defaultDisplayedPrice                | number                                               | <p>Variation price as it's shown on the storefront for logged out customers with default location (store location).</p><p><br>Pre-selected product options or variations modify the price.<br></p><p><strong>Includes taxes</strong></p>                                                                                                                                                                                                                     |
| defaultDisplayedPriceFormatted       | string                                               | <p>Formatted variant (curency symbol and delimeter settings) of <code>defaultDisplayedPrice</code> based on the store's format settings.<br><br>For example, <code>€11,00</code></p>                                                                                                                                                                                                                                                                         |
| lowestPrice                          | number                                               | Variation's lowest price for EU store.                                                                                                                                                                                                                                                                                                                                                                                                                       |
| lowestPriceSettings                  | object [lowestPriceSettings](#lowestpricesettings)   | <p>Variation's lowest price settings contain only one field: <code>lowestPriceEnabled</code> <br><br>It defines if the lowest price is enabled for the variation.</p>                                                                                                                                                                                                                                                                                        |
| defaultDisplayedLowestPrice          | number                                               | <p>Variation lowest price as it's shown on the storefront for logged out customers with default location (store location).<br><br><strong>Includes taxes</strong></p>                                                                                                                                                                                                                                                                                        |
| defaultDisplayedLowestPriceFormatted | string                                               | <p>Formatted variant (curency symbol and delimeter settings) of <code>defaultDisplayedLowestPrice</code> based on the store's format settings.<br><br>For example, <code>€11,00</code></p>                                                                                                                                                                                                                                                                   |
| dimensions                           | object [dimensions](#dimensions)                     | Variation's dimensions.                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| wholesalePrices                      | array of objects [wholesalePrices](#wholesaleprices) | Sorted list of wholesale price tiers specific to the variation: "minimum quantity = price" pairs.                                                                                                                                                                                                                                                                                                                                                            |
| weight                               | number                                               | Variation's weight for calculating shipping costs.                                                                                                                                                                                                                                                                                                                                                                                                           |
| volume                               | number                                               | Variation volume for calculations shipping costs, fractional number, `0` by default.                                                                                                                                                                                                                                                                                                                                                                         |
| warningLimit                         | number                                               | Minimum amount of variation in stock to trigger an automated "low stock" email notification for the store owner.                                                                                                                                                                                                                                                                                                                                             |
| attributes                           | array of objects [attributes](#attributes)           | List of variation attributes and their values.                                                                                                                                                                                                                                                                                                                                                                                                               |
| compareToPrice                       | number                                               | Pre-sale price for the variation.                                                                                                                                                                                                                                                                                                                                                                                                                            |
| minPurchaseQuantity                  | number                                               | <p>Sets minimum product purchase quantity. <br><br>Default value is <code>null</code>.</p>                                                                                                                                                                                                                                                                                                                                                                   |
| maxPurchaseQuantity                  | number                                               | <p>Sets maximum product purchase quantity. <br><br>Default value is <code>null</code>.</p>                                                                                                                                                                                                                                                                                                                                                                   |
| outOfStockVisibilityBehaviour        | boolean                                              | <p>Defines if a variation is visible and/or can be pre-ordered when out-of-stock. <br><br>Requires enabled pre-orders on the store level: <code>allowPreordersForOutOfStockProducts</code> setting in <code>/profile</code> endpoint.<br><br>Supported values:<br><code>SHOW</code> - Show out-of-stock variation, but adding it to the cart is disabled.<br><code>ALLOW\_PREORDER</code> - Show out-of-stock variation and allow adding it to the cart.</p> |
| alt                                  | object [alt](#alt)                                   | Image description for the "alt" HTML attribute and its translations.                                                                                                                                                                                                                                                                                                                                                                                         |

#### options

| Field           | Type                                 | Description                                          |
| --------------- | ------------------------------------ | ---------------------------------------------------- |
| name            | string                               | Name of the selected option.                         |
| nameTranslated  | object [translations](#translations) | Available translations for the product option name.  |
| value           | string                               | Value of the selected option.                        |
| valueTranslated | object [translations](#translations) | Available translations for the product option value. |

#### dimensions

<table><thead><tr><th width="250">Field</th><th>Type</th><th>Description</th></tr></thead><tbody><tr><td>length</td><td>number</td><td>Length of a product for calculating shipping costs.</td></tr><tr><td>width</td><td>number</td><td>Width of a product for calculating shipping costs.</td></tr><tr><td>height</td><td>number</td><td>Height of a product for calculating shipping costs.</td></tr></tbody></table>

#### lowestPriceSettings

<table><thead><tr><th>Field</th><th width="142">Type</th><th>Description</th></tr></thead><tbody><tr><td>lowestPriceEnabled</td><td>boolean</td><td>Defines if the lowest price is enabled for the product and shown on the storefront.</td></tr><tr><td>manualLowestPrice</td><td>number</td><td>Manually entered lowest price for the last 30 days before any discounts or taxes applied.</td></tr><tr><td>defaultDisplayedManualLowestPrice</td><td>number</td><td><code>manualLowestPrice</code> with taxes applied.</td></tr><tr><td>defaultDisplayedManualLowestPriceFormatted</td><td>string</td><td>Formatted display of <code>defaultDisplayedManualLowestPrice</code> using store format settings.</td></tr><tr><td>automaticLowestPrice</td><td>number</td><td>Automatically calculated lowest price for the last 30 days before any discounts or taxes applied. <br><br><strong>Read-only</strong></td></tr><tr><td>defaultDisplayedAutomaticLowestPrice</td><td>number</td><td><code>automaticLowestPrice</code> with taxes applied. <br><br><strong>Read-only</strong></td></tr><tr><td>defaultDisplayedAutomaticLowestPriceFormatted</td><td>string</td><td>Formatted display of <code>defaultDisplayedAutomaticLowestPrice</code> using store format settings. <br><br><strong>Read-only</strong></td></tr></tbody></table>

#### attributes

<table><thead><tr><th>Field</th><th width="185">Type</th><th>Description</th></tr></thead><tbody><tr><td>id</td><td>number</td><td>Internal attribute ID. </td></tr><tr><td>name</td><td>string</td><td>Attribute name visible on the storefront.</td></tr><tr><td>nameTranslated</td><td>object <a href="#translations">translations</a></td><td>Available translations for the attribute name.</td></tr><tr><td>value</td><td>string</td><td>Value of the attribute for this product.</td></tr><tr><td>valueTranslated</td><td>object <a href="#translations">translations</a></td><td>Available translations for the attribute value.</td></tr><tr><td>type</td><td>string</td><td>Attribute type. There are user-defined attributes, general attributes and attributes pre-defined by Ecwid, for example, "price per unit". <br><br>One of:<br><code>CUSTOM</code><br><code>UPC</code><br><code>BRAND</code><br><code>GENDER</code><br><code>AGE_GROUP</code><br><code>COLOR</code><br><code>SIZE</code><br><code>PRICE_PER_UNIT</code><br><code>UNITS_IN_PRODUCT</code></td></tr><tr><td>show</td><td>string</td><td>Defines if an attribute is visible on a product page. <br><br>One of:<br><code>NOTSHOW</code> - Not visible.<br><code>DESCR</code> - Visible under the product description.<br><code>PRICE</code> - Visible under the product price</td></tr></tbody></table>

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
