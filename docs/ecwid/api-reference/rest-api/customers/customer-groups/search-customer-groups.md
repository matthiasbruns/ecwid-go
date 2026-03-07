# Search customer groups

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/customer_groups`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
GET /api/v3/1003/customer_groups HTTP/1.1
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
      "id": 0,
      "name": "General"
    },
    {
      "id": 9367001,
      "name": "VIP"
    }
  ]
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_customers`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

All query params are optional.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>customerGroupId</td><td>number</td><td>Search specific customer groups by listing their IDs. <br><br>Supports multiple values, for example: <code>13456, 35678, 57890</code></td></tr><tr><td>keyword</td><td>string</td><td>Search term for the customer group name.</td></tr><tr><td>offset</td><td>number</td><td>Maximum number of returned items. Default value: <code>100</code>. Maximum allowed value: <code>1000</code>.</td></tr><tr><td>limit</td><td>number</td><td>Limit to the number of returned items. Maximum and default value (if not specified) is <code>100</code>.</td></tr><tr><td>responseFields</td><td>string</td><td>Specify the exact fields to receive in response JSON. If not specified, the response JSON will have all available fields for the entity.<br>Example: <code>?responseFields=total,items(name)</code></td></tr></tbody></table>

Example of using `responseFields` param:

{% tabs %}
{% tab title="Request" %}

```
curl --location 'https://app.ecwid.com/api/v3/1003/customer_groups?responseFields=total,items(name)' \
--header 'Authorization: Bearer secret_ab***cd'
```

{% endtab %}

{% tab title="Response" %}

```json
{
    "total": 2,
    "items": [
        {
            "name": "General"
        },
        {
            "name": "VIP"
        }
    ]
}
```

{% endtab %}
{% endtabs %}

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
| items  | array of objects [items](#items) | Detailed information about returned customer groups.                                         |

#### items

| Field               | Type   | Description                                                 |
| ------------------- | ------ | ----------------------------------------------------------- |
| id                  | number | Unique internal ID of the customer group.                   |
| name                | string | Customer group name visible to customers on the storefront. |
| externalReferenceId | string | External ID for syncing customer goups with other services. |
