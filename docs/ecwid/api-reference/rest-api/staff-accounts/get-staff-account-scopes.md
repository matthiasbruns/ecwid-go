# Get staff account scopes

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/profile/staffScopes`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```curl
curl --location 'https://app.ecwid.com/api/v3/1003/profile/staffScopes' \
--header 'Authorization: Bearer secret_ab***cd'
```

Response:

```json
{
    "staffScopes": [
        "SALES_MANAGEMENT",
        "CATALOG_MANAGEMENT",
        "MARKETING_MANAGEMENT",
        "REPORT_ACCESS",
        "WEBSITE_MANAGEMENT",
        "SALES_CHANNELS_MANAGEMENT",
        "STORE_MANAGEMENT"
    ]
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_store_profile`, `read_staff`

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

| Field       | Type             | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| ----------- | ---------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| staffScopes | array of strings | <p>List of all permissions that can be given to staff accounts in the store.<br><br>Supported values:<br><code>SALES\_MANAGEMENT</code> - Access to managing orders.<br><code>CATALOG\_MANAGEMENT</code> - Managing catalog (products/variations/categories).<br><code>MARKETING\_MANAGEMENT</code>- Managing marketing tools/SEO.<br><code>REPORT\_ACCESS</code> - Access to stats/reports.<br><code>WEBSITE\_MANAGEMENT</code> - Access to website settings (Instant Site Editor).<br><code>SALES\_CHANNELS\_MANAGEMENT</code> - Access to <a href="https://support.ecwid.com/hc/en-us/sections/360004414640-Sales-channels">sales channels settings</a>.<br><code>STORE\_MANAGEMENT</code> Full Ecwid admin access including the <a href="https://my.ecwid.com/#develop-apps">#develop-apps</a> page.</p> |
