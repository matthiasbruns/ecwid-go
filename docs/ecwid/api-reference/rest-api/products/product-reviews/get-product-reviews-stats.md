# Get product reviews stats

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/reviews/filters_data`

<details>

<summary>Request and response example</summary>

Request:

```http
GET /api/v3/1003/reviews/filters_data HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
```

Response:

```json
{
  "allCount": 2,
  "moderatedCount": 0,
  "publishedCount": 2
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_reviews`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Field          | Type   | Description                                                  |
| -------------- | ------ | ------------------------------------------------------------ |
| allCount       | number | Total count of reviews in the store.                         |
| moderatedCount | number | Total count of reviews not yet published by the store owner. |
| publishedCount | number | Total count of reviews already published by the store owner. |
