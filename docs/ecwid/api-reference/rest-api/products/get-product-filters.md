# Get product filters

<mark style="color:blue;">`POST`</mark> `https://app.ecwid.com/api/v3/{storeId}/products/filters`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
POST /api/v3/1003/products/filters HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
Content-Type: application/json

{
   "params": {
      "enabled": "true",
      "filterFacetLimit": "200",
      "filterFields": "price,inventory,onsale,categories,attribute_Brand,option_Size",
      "includeProductsFromSubcategories": "true",
      "lang": "en"
   }
}
```

Response:

```json
{
  "productCount": 2,
  "filters": {
    "price": {
      "minValue": 27.5,
      "maxValue": 77,
      "status": "SUCCESS"
    },
    "inventory": {
      "values": [
        {
          "id": "instock",
          "title": "In stock",
          "productCount": 2
        }
      ],
      "status": "SUCCESS"
    },
    "onsale": {
      "values": [
        {
          "id": "notonsale",
          "title": "Regular price",
          "productCount": 2
        }
      ],
      "status": "SUCCESS"
    },
    "categories": {
      "values": [
        {
          "id": 172966754,
          "title": "Not toys",
          "productCount": 1
        },
        {
          "id": 172786255,
          "title": "Toys",
          "productCount": 1
        }
      ],
      "status": "SUCCESS",
      "sortingOrder": "PRODUCTS_COUNT_DESC"
    },
    "attribute_Brand": {
      "title": "Brand",
      "values": [
        {
          "title": "LIGHTSPEED",
          "productCount": 1
        },
        {
          "title": "Test_Brand",
          "productCount": 1
        }
      ],
      "status": "SUCCESS",
      "sortingOrder": "FILTER_VALUE_ASC"
    },
    "option_Size": {
      "title": "Size",
      "values": [
        {
          "title": "24",
          "productCount": 1
        },
        {
          "title": "28",
          "productCount": 1
        },
        {
          "title": "32",
          "productCount": 1
        }
      ],
      "status": "SUCCESS",
      "sortingOrder": "FILTER_VALUE_ASC"
    }
  }
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_catalog`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

Request **requires** one query param.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>filterFields</td><td>string</td><td>Comma-separated list of filters for Ecwid to return.<br><br>Supported filters: <code>"price"</code>,<code>"inventory"</code>,<code>"onsale"</code>,<code>"categories"</code>, <code>"option_{optionName}"</code>, <code>"attribute_{attributeName}"</code>. Example: <code>"price,inventory,option_Size,attribute_Brand,categories"</code>. <br><br>If an option or attribute has a comma or backslash in its name, escape it with a backslash: <code>"\"</code>. I.e. option name <code>"Color, size"</code> will transform to <code>"option_Color\, size"</code> when used in query param.</td></tr></tbody></table>

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Request JSON

A JSON object with the following fields:

<table><thead><tr><th>Field</th><th width="252">Type</th><th>Description</th></tr></thead><tbody><tr><td>params</td><td>object <a data-mention href="#params">#params</a></td><td>Filtering params for the request.</td></tr></tbody></table>

#### params

<table><thead><tr><th width="120.36328125">Field</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>filterFields</td><td>string</td><td><p>String that defines all applied filtering params, for example: <code>"price,inventory,option_Size"</code>.<br><br>Full list of supported filters: </p><ul><li><code>"price"</code>  - Filter products by price.</li><li><code>"inventory"</code>  - Filter products by their stock.</li><li><code>"enabled"</code> - Filter products by their availability on the storefront.</li><li><code>"onsale"</code>  - Filter products that are currently on sale.</li><li><code>"categories"</code> - Filter products by their categories.</li><li><code>"option_{optionName}"</code> - Filter products by option name and values.</li><li><code>"attribute_{attributeName}"</code>  - Filter products by attribute name and values.</li></ul><p>If an option/attribute has a comma or a backslash in its name, escape it with a backslash  <code>\</code>. <br><br>For example, an option named <code>"Color, size"</code> transforms to <code>"option_Color\, size"</code> in the request.</p></td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

<table><thead><tr><th width="200.39453125">Field</th><th width="200.08203125">Type</th><th>Description</th></tr></thead><tbody><tr><td>productCount</td><td>number</td><td>Total number of products matching any specified filters</td></tr><tr><td>filters</td><td>object <a data-mention href="#filters">#filters</a></td><td>List of found product filters.</td></tr></tbody></table>

#### filters

<table><thead><tr><th width="200.37890625">Field</th><th width="199.65625">Type</th><th>Description</th></tr></thead><tbody><tr><td>price</td><td>object <a data-mention href="#price">#price</a></td><td>Price filters.</td></tr><tr><td>inventory</td><td>object <a data-mention href="#inventory">#inventory</a></td><td>Inventory filters – number of in stock and out-of-stock products.</td></tr><tr><td>onsale</td><td>object <a data-mention href="#onsale">#onsale</a></td><td>Filter for currently discounted products.</td></tr><tr><td>categories</td><td>object <a data-mention href="#categories">#categories</a></td><td>Filter for product categories.</td></tr><tr><td>attribute_{attrName}</td><td>object <a data-mention href="#attribute_-attrname">#attribute_-attrname</a></td><td>Filter for product attributes. <br><br>Response can contain multiple attribute filters. Field name for each attribute filter includes the attribute name.</td></tr><tr><td>option_{optionName}</td><td>object <a data-mention href="#option_-optionname">#option_-optionname</a></td><td>Filter for product options. <br><br>Response can contain multiple option filters. Field name for each option filter includes the attribute name.</td></tr></tbody></table>

#### price

<table><thead><tr><th width="200.39453125">Field</th><th width="200.08203125">Type</th><th>Description</th></tr></thead><tbody><tr><td>minValue</td><td>number</td><td>Minimal product price in the store for applying filters.</td></tr><tr><td>minValue</td><td>number</td><td>Maximum product price in the store for applying filters.</td></tr><tr><td>status</td><td>string</td><td>If <code>SUCCESS</code>, the price filter can be applied on the storefront.</td></tr></tbody></table>

#### inventory

<table><thead><tr><th width="200.39453125">Field</th><th width="200.08203125">Type</th><th>Description</th></tr></thead><tbody><tr><td>values</td><td>array of objects values</td><td>Found invontory filters.</td></tr><tr><td>success</td><td>string</td><td>If <code>SUCCESS</code>, found filters can be applied on the storefront.</td></tr></tbody></table>

Objects inside the `values` array list found inventory filters in the following format:

```json
{
  "id": "instock",
  "title": "In stock",
  "productCount": 2
}
```

#### onsale

<table><thead><tr><th width="200.39453125">Field</th><th width="200.08203125">Type</th><th>Description</th></tr></thead><tbody><tr><td>values</td><td>array of objects values</td><td>Found sale filters.</td></tr><tr><td>success</td><td>string</td><td>If <code>SUCCESS</code>, found filters can be applied on the storefront.</td></tr></tbody></table>

Objects inside the `values` array list found sale filters in the following format:

```json
{
  "id": "notonsale",
  "title": "Regular price",
  "productCount": 2
}
```

#### categories

<table><thead><tr><th width="200.39453125">Field</th><th width="200.08203125">Type</th><th>Description</th></tr></thead><tbody><tr><td>values</td><td>array of objects values</td><td>Found category filters.</td></tr><tr><td>success</td><td>string</td><td>If <code>SUCCESS</code>, found filters can be applied on the storefront.</td></tr></tbody></table>

Objects inside the `values` array list found category filters in the following format:

```json
{
  "id": 172786255,
  "title": "Toys",
  "productCount": 1
}
```

#### attribute\_{attrName}

<table><thead><tr><th width="200.39453125">Field</th><th width="200.08203125">Type</th><th>Description</th></tr></thead><tbody><tr><td>title</td><td>string</td><td>Title of the found attribute filter.</td></tr><tr><td>values</td><td>array of objects values</td><td>Found attribute filters.</td></tr><tr><td>success</td><td>string</td><td>If <code>SUCCESS</code>, found filters can be applied on the storefront.</td></tr><tr><td>sortingOrder</td><td>string</td><td>Default sorting order for the filter, for example, <code>"FILTER_VALUE_ASC"</code></td></tr></tbody></table>

Objects inside the `values` array list found attribute filters in the following format:

```json
"values": [
  {
    "title": "LIGHTSPEED",
    "productCount": 1
  },
  {
    "title": "Test_Brand",
    "productCount": 1
  }
]
```

#### option\_{optionName}

<table><thead><tr><th width="200.39453125">Field</th><th width="200.08203125">Type</th><th>Description</th></tr></thead><tbody><tr><td>title</td><td>string</td><td>Title of the found option filter.</td></tr><tr><td>values</td><td>array of objects values</td><td>Found option filters.</td></tr><tr><td>success</td><td>string</td><td>If <code>SUCCESS</code>, found filters can be applied on the storefront.</td></tr><tr><td>sortingOrder</td><td>string</td><td>Default sorting order for the filter, for example, <code>"FILTER_VALUE_ASC"</code></td></tr></tbody></table>

Objects inside the `values` array list found option filters in the following format:

```json
"values": [
  {
    "title": "24",
    "productCount": 1
  },
  {
    "title": "28",
    "productCount": 1
  },
  {
    "title": "32",
    "productCount": 1
  }
]
```
