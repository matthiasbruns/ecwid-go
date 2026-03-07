# Update product type and attributes

<mark style="color:purple;">`PUT`</mark> `https://app.ecwid.com/api/v3/{storeId}/classes/{classId}`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
PUT /api/v3/1003/classes/4208002 HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
Content-Type: application/json
Cache-Control: no-cache

{
    "name": "New Class Name", 
    "attributes": [
        {
            "id": 5508062, 
            "name": "New attribute name",
            "type": "CUSTOM",
            "show": "DESCR"            
        },
        {
            "name": "Model ID",
            "type": "CUSTOM",
            "show": "DESCR"            
        }
    ]
}
```

Response:

{% code fullWidth="true" %}

```json
{
    "updateCount": 1
}
```

{% endcode %}

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `update_catalog`

### Path params

All path params are required.

| Param   | Type   | Description               |
| ------- | ------ | ------------------------- |
| storeId | number | Ecwid store ID.           |
| classId | number | Internal product type ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Request JSON

A JSON object with the following fields:

<table><thead><tr><th width="179">Field</th><th width="229">Type</th><th>Description</th></tr></thead><tbody><tr><td>name</td><td>string</td><td>Product type name. Empty for the "General" type.</td></tr><tr><td>attributes</td><td>array of objects <a href="#attributes">attributes</a></td><td>Product attributes assigned to this product type.<br><br><strong>Note:</strong> to add new product attributes, send both new and all existing attributes in the request. Otherwise, the request will delete all attributes currently assigned to the product type.</td></tr></tbody></table>

#### attributes

<table><thead><tr><th width="181">Field</th><th width="183">Type</th><th>Description</th></tr></thead><tbody><tr><td>id</td><td>number</td><td>Internal unique ID of the product attribute.<br><br>Specify attribute ID to update the existing attribute. If ID is not specified, the attribute is considered new.</td></tr><tr><td>name</td><td>string</td><td>Attribute title visible. Product attribute with an empty name field will also be returned</td></tr><tr><td>nameTranslated</td><td>object <a href="#translations">translations</a></td><td>Available translations for product attribute name</td></tr><tr><td>type</td><td>string</td><td><p>Type of the attribute that defines its functionality. </p><p></p><p>One of:</p><p><code>CUSTOM</code></p><p><code>UPC</code></p><p><code>BRAND</code></p><p><code>GENDER</code></p><p><code>AGE_GROUP</code></p><p><code>COLOR</code></p><p><code>SIZE</code> </p><p><code>TAGS</code></p><p><code>PRICE_PER_UNIT</code></p><p><code>UNITS_IN_PRODUCT</code><br><br>Attributes of type <code>PRICE_PER_UNIT</code> and <code>UNITS_IN_PRODUCT</code> are only available if the price-per-unit feature is enabled.</p></td></tr><tr><td>show</td><td>string</td><td>Defines if an attribute is visible on a product page. Supported values: <code>NOTSHOW</code>, <code>DESCR</code>, <code>PRICE</code>. The value <code>PRICE</code> = <code>DESCR</code>. For <a href="ref:authentication-basics#access-tokens">public tokens</a>, <code>NOTSHOW</code> attributes are not returned</td></tr></tbody></table>

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

| Field       | Type   | Description                                                                                                                                                                                   |
| ----------- | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| updateCount | number | <p>The number of updated items that defines if the request was successful.<br><br>One of:</p><p><code>1</code> if the item was updated,</p><p><code>0</code> if the item was not updated.</p> |
