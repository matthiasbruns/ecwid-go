# Search staff accounts

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/staff`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```curl
curl --location 'https://app.ecwid.com/api/v3/1003/staff' \
--header 'Authorization: Bearer secret_ab***cd'
```

Response:

```json
{
    "staffList": [
        {
            "id": "p27632593",
            "name": "John Doe",
            "email": "john.doe@example.com",
            "staffScopes": [
                "REPORT_ACCESS",
                "SALES_MANAGEMENT",
                "MARKETING_MANAGEMENT"
            ],
            "inviteAccepted": true
        },
        {
            "id": "p3855016",
            "name": "API Team",
            "email": "ec.apps@lightspeedhq.com",
            "staffScopes": [
                "REPORT_ACCESS",
                "WEBSITE_MANAGEMENT",
                "STORE_MANAGEMENT",
                "CATALOG_MANAGEMENT",
                "SALES_CHANNELS_MANAGEMENT",
                "SALES_MANAGEMENT",
                "MARKETING_MANAGEMENT"
            ],
            "inviteAccepted": true
        }
    ]
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_staff`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

All query params are optional.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>responseFields</td><td>string</td><td>Specify the exact fields to receive in response JSON. If not specified, the response JSON will have all available fields for the entity.<br><br>For example: <code>?responseFields=staffList(name,email,inviteAccepted)</code></td></tr></tbody></table>

Example of using `responseFields` param:

{% tabs %}
{% tab title="Request" %}

```
curl --location 'https://app.ecwid.com/api/v3/1003/staff?responseFields=staffList(name,email,inviteAccepted)' \
--header 'Authorization: Bearer secret_ab***cd'
```

{% endtab %}

{% tab title="Response" %}

```json
{
    "staffList": [
        {
            "name": "John Doe",
            "email": "john.doe@example.com",
            "inviteAccepted": true
        },
        {
            "name": "API Team",
            "email": "ec.apps@lightspeedhq.com",
            "inviteAccepted": true
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

| Field     | Type                       | Description                                        |
| --------- | -------------------------- | -------------------------------------------------- |
| staffList | array of objects staffList | Details about staff accounts invited to the store. |

#### staffList

| Field          | Type             | Description                                                                                                                                                                                                                                                                                  |
| -------------- | ---------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| id             | string           | Internal unique ID of the staff account.                                                                                                                                                                                                                                                     |
| name           | string           | Staff account name.                                                                                                                                                                                                                                                                          |
| email          | string           | Staff account email.                                                                                                                                                                                                                                                                         |
| staffScopes    | array of strings | <p>Permissions enabled for the staff account. If empty, the account has all permissions. <br><br>Learn more about staff account permissions in <a href="https://support.ecwid.com/hc/en-us/articles/115005355089-Adding-and-managing-staff-accounts#-staff-permissions">Help Center</a>.</p> |
| inviteAccepted | boolean          | Defines if user accepted their invite to become staff account for the store.                                                                                                                                                                                                                 |
