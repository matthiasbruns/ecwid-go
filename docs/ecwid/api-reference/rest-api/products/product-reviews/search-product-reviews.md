# Search product reviews

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/reviews`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
GET /api/v3/1003/reviews HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
```

Response:

```json
{
  "total": 2,
  "count": 2,
  "offset": 0,
  "limit": 100,
  "items": [
    {
      "id": 712737671,
      "status": "published",
      "customerId": 8108179152,
      "productId": 123456,
      "orderId": "2D31G",
      "rating": 5,
      "review": "Just what I need",
      "reviewerInfo": {
        "name": "Abraham Smith",
        "email": "example_customer@example.com",
        "city": "New York",
        "orders": 2
      },
      "createDate": "2025-02-26 13:37:46 +0000",
      "updateDate": "2025-02-27 04:40:33 +0000",
      "createTimestamp": 1740562666,
      "updateTimestamp": 1740616833
    },
    {
      "id": 812738127,
      "status": "published",
      "customerId": 239921001,
      "productId": 123456,
      "orderId": "HI88G",
      "rating": 5,
      "review": "A really good product. Definitely worth buying!",
      "reviewerInfo": {
        "name": "John Doe",
        "email": "ec.apps@lightspeedhq.com",
        "city": "New York",
        "orders": 1
      },
      "createDate": "2024-12-02 10:00:00 +0000",
      "updateDate": "2024-12-02 10:00:00 +0000",
      "createTimestamp": 1733119200,
      "updateTimestamp": 1733119200
    }
  ]
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

### Query params

All query params are optional.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>status</td><td>boolean</td><td><p>Search term for the review status. <br><br>One of: </p><p><code>moderated</code> - Review is not yet published by the store owner.</p><p><code>published</code> - Review is published by the store owner.</p></td></tr><tr><td>rating</td><td>number</td><td>Search term for the review 5-star rating (from <code>1</code> to <code>5</code>).</td></tr><tr><td>orderId</td><td>string</td><td>Search by the order ID associated with the review.</td></tr><tr><td>productId</td><td>number</td><td>Search by the product ID associated with the review.</td></tr><tr><td>reviewId</td><td>number</td><td>Search by the internal review ID.</td></tr><tr><td>createdFrom</td><td>number/string</td><td>Review creation datetime (lower bound). Supported formats: UNIX timestamp, date/time. <br><br>For example: <code>1447804800</code>, <code>2023-01-15 19:27:50</code></td></tr><tr><td>createdTo</td><td>number/string</td><td>Review creation datetime (upper bound). Supported formats: UNIX timestamp, date/time. <br><br>For example: <code>1447804800</code>, <code>2023-01-15 19:27:50-</code></td></tr><tr><td>updatedFrom</td><td>number/string</td><td>Review latest update datetime (lower bound). Supported formats: UNIX timestamp, date/time. <br><br>For example: <code>1447804800</code>, <code>2023-01-15 19:27:50</code></td></tr><tr><td>updatedTo</td><td>number/string</td><td>Review latest update datetime (upper bound). Supported formats: UNIX timestamp, date/time. <br><br>For example: <code>1447804800</code>, <code>2023-01-15 19:27:50</code></td></tr><tr><td>sortBy</td><td>string</td><td>Define review sorting in the response. <br><br>One of: <br><code>DATE_CREATED_ASC</code><br><code>DATE_CREATED_DESC</code><br><code>RATING_ASC</code><br><code>RATING_DESC</code></td></tr><tr><td>keyword</td><td>string</td><td>Search term for the customer's review text (<code>review</code> field).</td></tr><tr><td>responseFields</td><td>string</td><td>Specify the exact fields to receive in response JSON. If not specified, the response JSON will have all available fields for the entity.<br><br>For example: <code>?responseFields=items(id,status,rating</code>)</td></tr></tbody></table>

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
| items  | array of objects [items](#items) | Detailed information about returned reviews.                                                 |

#### items

<table><thead><tr><th width="215">Field</th><th width="141">Type</th><th>Description</th></tr></thead><tbody><tr><td>id</td><td>number</td><td>Internal unique review ID.</td></tr><tr><td>status</td><td>string</td><td><p>Review publication status.<br><br>One of: </p><p><code>moderated</code> - Review is not yet published by the store owner.</p><p><code>published</code> - Review is published by the store owner.</p></td></tr><tr><td>rating</td><td>number</td><td>5-star rating of the review. (from <code>1</code> to <code>5</code>).</td></tr><tr><td>orderId</td><td>string</td><td>Order ID associated with the review.</td></tr><tr><td>productId</td><td>number</td><td>Product ID associated with the review.</td></tr><tr><td>customerId</td><td>number</td><td>Customer ID associated with the review.</td></tr><tr><td>review</td><td>string</td><td>Review text left by the customer.</td></tr><tr><td>reviewerInfo</td><td>object reviewerInfo</td><td>Details about the customer who placed the review.</td></tr><tr><td>createDate</td><td>string</td><td>Datetime when customer left the review in date format. <br><br>For example, <code>2024-05-26 13:37:46 +0000</code></td></tr><tr><td>updateDate</td><td>string</td><td>Datetime of the latest review update in date format.<br><br>For example, <code>2024-05-26 13:37:46 +0000</code></td></tr><tr><td>createTimestamp</td><td>number</td><td>Datetime when customer left the review in UNIX timestamp format. <br><br>For example, <code>1622036266926</code></td></tr><tr><td>updateTimestamp</td><td>number</td><td>Datetime of the latest review update in UNIX timestamp format. <br><br>For example, <code>1622036266926</code></td></tr></tbody></table>

#### reviewerInfo

| Field   | Type   | Description                                                               |
| ------- | ------ | ------------------------------------------------------------------------- |
| name    | string | Name of the customer who left the review.                                 |
| email   | string | Email of the customer who left the review.                                |
| country | string | Country specified by the customer in an order associated with the review. |
| city    | string | City specified by the customer in an order associated with the review.    |
| orders  | number | Amount of orders placed by the customer who left the review.              |
