# Get customer group

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/customer_groups/{groupId}`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
GET /api/v3/1003/customer_groups/9367001 HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
```

Response:

```json
{
  "id": 9367001,
  "name": "VIP"
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_customers`

### Path params

All path params are required.

| Param   | Type   | Description                 |
| ------- | ------ | --------------------------- |
| storeId | number | Ecwid store ID.             |
| groupId | number | Internal customer group ID. |

### Query params

All query params are optional.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>responseFields</td><td>string</td><td>Specify the exact fields to receive in response JSON. If not specified, the response JSON will have all available fields for the entity.<br>Example: <code>?responseFields=name</code></td></tr></tbody></table>

Example of using `responseFields` param:

{% tabs %}
{% tab title="Request" %}

```
curl --location 'https://app.ecwid.com/api/v3/1003/customer_groups/0?responseFields=name' \
--header 'Authorization: Bearer secret_ab***cd'
```

{% endtab %}

{% tab title="Response" %}

```json
{
    "name": "General"
}
```

{% endtab %}
{% endtabs %}

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Field               | Type   | Description                                                 |
| ------------------- | ------ | ----------------------------------------------------------- |
| id                  | number | Unique internal ID of the customer group.                   |
| name                | string | Customer group name visible to customers on the storefront. |
| externalReferenceId | string | External ID for syncing customer goups with other services. |
