# Update order of categories

<mark style="color:purple;">`PUT`</mark> `https://app.ecwid.com/api/v3/{storeId}/categories/sort?parentCategory={categoryId}`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
PUT /api/v3/1003/categories/sort?parentCategory=0 HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
Content-Type: application/json
Cache-Control: no-cache

{
  "sortedIds": [
    172966754,
    172786255
  ]
}
```

Response:

```json
{
  "updateCount": 1
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_catalog` , `update_catalog`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

Request supports only one **required** query param.

| Param          | Type   | Description                                                                                                                                                                            |
| -------------- | ------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| parentCategory | number | <p>Internal ID of the parent category.<br><br>Use <code>0</code> value to sort categories inside the lop-level category named "Store front page".<br><br><strong>Required</strong></p> |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Request JSON

A JSON object with the following fields:

| Field     | Type             | Description                                                                                                                    |
| --------- | ---------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| sortedIds | array of numbers | List of categories inside their parent category order in the same way as they are listed in Ecwid admin and on the storefront. |

### Response JSON

A JSON object with the following fields:

| Field       | Type   | Description                                                                                                                                                                                   |
| ----------- | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| updateCount | number | <p>The number of updated items that defines if the request was successful.<br><br>One of:</p><p><code>1</code> if the item was updated,</p><p><code>0</code> if the item was not updated.</p> |
