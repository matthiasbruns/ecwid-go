# Create customer group

<mark style="color:blue;">`POST`</mark> `https://app.ecwid.com/api/v3/{storeId}/customer_groups`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
POST /api/v3/1003/customer_groups HTTP/1.1
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
  "id": 9367001
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `create_customers`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Request JSON

A JSON object with the following fields:

| Field               | Type   | Description                                                                                         |
| ------------------- | ------ | --------------------------------------------------------------------------------------------------- |
| name                | string | <p>Customer group name visible to customers on the storefront.<br><br><strong>Required</strong></p> |
| externalReferenceId | string | External ID for syncing customer goups with other services.                                         |

### Response JSON

A JSON object with the following fields:

| Field | Type   | Description                                |
| ----- | ------ | ------------------------------------------ |
| id    | number | Internal ID of the created customer group. |
