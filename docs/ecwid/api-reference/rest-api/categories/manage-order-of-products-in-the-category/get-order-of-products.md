# Get order of products

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/products/sort?parentCategory={categoryId}`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
GET /api/v3/1003/products/sort?parentCategory=0 HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
```

Response:

```json
{
  "sortedIds": [
    689454040,
    692730761,
    724894174
  ]
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_catalog`&#x20;

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

Request supports only one **required** query param.

| Param          | Type   | Description                                                                                                                                                                                                |
| -------------- | ------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| parentCategory | number | <p>Internal ID of the parent category.<br><br>Use <code>0</code> value to receive a sorted list of categories inside the lop-level category named "Store front page".<br><br><strong>Required</strong></p> |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Field     | Type             | Description                                                                                                                    |
| --------- | ---------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| sortedIds | array of numbers | List of categories inside their parent category order in the same way as they are listed in Ecwid admin and on the storefront. |
