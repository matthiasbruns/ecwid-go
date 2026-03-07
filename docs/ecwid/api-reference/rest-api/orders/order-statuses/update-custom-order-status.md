# Update custom order status

<mark style="color:purple;">`PUT`</mark> `https://app.ecwid.com/api/v3/{storeId}/profile/order_status/{statusId}`

<details>

<summary>Request and response example</summary>

Request:

```http
PUT /api/v3/1003/profile/order_status/CUSTOM_PAYMENT_STATUS_1 HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
Content-Type: application/json
Cache-Control: no-cache

{
    "name": "Money on hold",
    "enabled": true,
    "sendNotificationWhenStatusIsAssigned": true
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

Your app must have the following **access scopes** to make this request: `update_store_profile`

### Path params

All path params are required.

<table><thead><tr><th width="154">Param</th><th width="127">Type</th><th>Description</th></tr></thead><tbody><tr><td>storeId</td><td>number</td><td>Ecwid store ID.</td></tr><tr><td>statusId</td><td>string</td><td><p>Status ID in the same format as in the <code>statusId</code> field. <strong>Case sensitive</strong>.</p><p></p><p>Supports only custom statuses:</p><p><code>CUSTOM_PAYMENT_STATUS_1</code></p><p><code>CUSTOM_PAYMENT_STATUS_2</code></p><p><code>CUSTOM_PAYMENT_STATUS_3</code></p><p><code>CUSTOM_FULFILLMENT_STATUS_1</code></p><p><code>CUSTOM_FULFILLMENT_STATUS_2</code></p><p><code>CUSTOM_FULFILLMENT_STATUS_3</code></p></td></tr></tbody></table>

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Request JSON

A JSON object with the following fields:

<table><thead><tr><th width="240.82421875">Field</th><th width="162.359375">Type</th><th>Description</th></tr></thead><tbody><tr><td>enabled</td><td>boolean</td><td>Set <code>true</code> to enable custom status in the store.</td></tr><tr><td>sendNotificationWhenStatusIsAssigned</td><td>boolean</td><td>Set <code>true</code> to enable automatic "Order status updated" email notifications for customers.</td></tr><tr><td>name</td><td>string</td><td>Name of the custom order status visible in Ecwid admin and emails.</td></tr><tr><td>nameTranslations</td><td>object <a href="#translations">translations</a></td><td>Available translations for the custom order status name.</td></tr></tbody></table>

#### translations

Object with text field translations in the `"lang": "text"` format, where the `"lang"` is an ISO 639-1 language code. For example:

```
{
    "en": "Sample text",
    "nl": "Voorbeeldtekst"
}
```

Translations are available for all active store languages. Only the default language translations are returned if no other translations are provided for the field. Find active store languages with <mark style="color:green;">`GET`</mark> `/profile` request > `languages` > `enabledLanguages`.

### Response JSON

A JSON object with the following fields:

| Field       | Type   | Description                                                                                                                                                                                   |
| ----------- | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| updateCount | number | <p>The number of updated items that defines if the request was successful.<br><br>One of:</p><p><code>1</code> if the item was updated,</p><p><code>0</code> if the item was not updated.</p> |
