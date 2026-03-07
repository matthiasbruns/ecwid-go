# Unassign products from the category

<mark style="color:blue;">`POST`</mark> `https://app.ecwid.com/api/v3/{storeId}/categories/{categoryId}/unassignProducts`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
POST /api/v3/1003/categories/9691094/unassignProducts HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
Content-Type: application/json
Cache-Control: no-cache

{
  "productIds": [
    37208345
  ]
}
```

Response:

```json
{
  "deleteCount": 1
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `update_catalog`

### Path params

All path params are required.

| Param      | Type   | Description           |
| ---------- | ------ | --------------------- |
| storeId    | number | Ecwid store ID.       |
| categoryId | number | Internal category ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Request JSON

A JSON object with the following fields:

| Field      | Type             | Description                                                                                                                                                                                                                                                                          |
| ---------- | ---------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| productIds | array of numbers | <p>List of products that will be unassigned to the category.<br><br><strong>Note 1:</strong> Unassigned products are not deleted from the store.<br><br><strong>Note 2:</strong> Unassigned product without any categories are automatically assigned to the top-level category.</p> |

### Response JSON

A JSON object with the following fields:

| Field       | Type   | Description                                                                                                                                                             |
| ----------- | ------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| deleteCount | number | <p>Defines if the request was successful.<br><br>One of:</p><p><code>1</code> if any products were unassigned,</p><p><code>0</code> if no products were unassigned.</p> |
