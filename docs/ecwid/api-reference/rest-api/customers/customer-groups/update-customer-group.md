# Update customer group

<mark style="color:purple;">`PUT`</mark> `https://app.ecwid.com/api/v3/{storeId}/customer_groups/{groupId}`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
PUT /api/v3/1003/customer_groups/9367001 HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
Content-Type: application/json
Cache-Control: no-cache

{
  "name": "VIP Customers"
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

Your app must have the following **access scopes** to make this request: `update_customers`

### Path params

All path params are required.

| Param   | Type   | Description                 |
| ------- | ------ | --------------------------- |
| storeId | number | Ecwid store ID.             |
| groupId | number | Internal customer group ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Request JSON

A JSON object with the following fields:

| Field               | Type   | Description                                                 |
| ------------------- | ------ | ----------------------------------------------------------- |
| name                | string | Customer group name visible to customers on the storefront. |
| externalReferenceId | string | External ID for syncing customer goups with other services. |

### Response JSON

A JSON object with the following fields:

| Field       | Type   | Description                                                                                                                                                                                   |
| ----------- | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| updateCount | number | <p>The number of updated items that defines if the request was successful.<br><br>One of:</p><p><code>1</code> if the item was updated,</p><p><code>0</code> if the item was not updated.</p> |
