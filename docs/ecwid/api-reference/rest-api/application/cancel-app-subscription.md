# Cancel app subscription

<mark style="color:red;">`DELETE`</mark> `https://app.ecwid.com/api/v3/{storeId}/application?namespace={namespace}`&#x20;

{% hint style="info" %}
Request cancels app subscription, revokes all its tokens, and uninstalls the app from the store.
{% endhint %}

### Required access scopes

Your app must have the following **access scopes** to make this request: `update_store_profile`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

Request works with one required query param.

| Param     | Type   | Description        |
| --------- | ------ | ------------------ |
| namespace | string | App's `client_id`. |
