# Search order statuses

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/profile/order_statuses`

<details>

<summary>Request and response example</summary>

Request:

```http
GET /api/v3/1003/profile/order_statuses HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
```

Response:

```json
{
    "items": [
        {
            "statusId": "AWAITING_PAYMENT",
            "orderStatusType": "PAYMENT_STATUS",
            "defaultStatus": true,
            "enabled": true,
            "sendNotificationWhenStatusIsAssigned": true
        },
        {
            "statusId": "PAID",
            "orderStatusType": "PAYMENT_STATUS",
            "defaultStatus": true,
            "enabled": true,
            "sendNotificationWhenStatusIsAssigned": true
        },
        {
            "statusId": "CANCELLED",
            "orderStatusType": "PAYMENT_STATUS",
            "defaultStatus": true,
            "enabled": true,
            "sendNotificationWhenStatusIsAssigned": true
        },
        {
            "statusId": "REFUNDED",
            "orderStatusType": "PAYMENT_STATUS",
            "defaultStatus": true,
            "enabled": true,
            "sendNotificationWhenStatusIsAssigned": true
        },
        {
            "statusId": "PARTIALLY_REFUNDED",
            "orderStatusType": "PAYMENT_STATUS",
            "defaultStatus": true,
            "enabled": true,
            "sendNotificationWhenStatusIsAssigned": true
        },
        {
            "statusId": "INCOMPLETE",
            "orderStatusType": "PAYMENT_STATUS",
            "defaultStatus": true,
            "enabled": true,
            "sendNotificationWhenStatusIsAssigned": true
        },
        {
            "statusId": "CUSTOM_PAYMENT_STATUS_1",
            "orderStatusType": "PAYMENT_STATUS",
            "enabled": false,
            "sendNotificationWhenStatusIsAssigned": true
        },
        {
            "statusId": "CUSTOM_PAYMENT_STATUS_2",
            "orderStatusType": "PAYMENT_STATUS",
            "enabled": false,
            "sendNotificationWhenStatusIsAssigned": true
        },
        {
            "statusId": "CUSTOM_PAYMENT_STATUS_3",
            "orderStatusType": "PAYMENT_STATUS",
            "enabled": false,
            "sendNotificationWhenStatusIsAssigned": true
        },
        {
            "statusId": "AWAITING_PROCESSING",
            "orderStatusType": "FULFILLMENT_STATUS",
            "defaultStatus": true,
            "enabled": true,
            "sendNotificationWhenStatusIsAssigned": true
        },
        {
            "statusId": "PROCESSING",
            "orderStatusType": "FULFILLMENT_STATUS",
            "defaultStatus": true,
            "enabled": true,
            "sendNotificationWhenStatusIsAssigned": true
        },
        {
            "statusId": "SHIPPED",
            "orderStatusType": "FULFILLMENT_STATUS",
            "defaultStatus": true,
            "enabled": true,
            "sendNotificationWhenStatusIsAssigned": true
        },
        {
            "statusId": "DELIVERED",
            "orderStatusType": "FULFILLMENT_STATUS",
            "defaultStatus": true,
            "enabled": true,
            "sendNotificationWhenStatusIsAssigned": true
        },
        {
            "statusId": "WILL_NOT_DELIVER",
            "orderStatusType": "FULFILLMENT_STATUS",
            "defaultStatus": true,
            "enabled": true,
            "sendNotificationWhenStatusIsAssigned": true
        },
        {
            "statusId": "RETURNED",
            "orderStatusType": "FULFILLMENT_STATUS",
            "defaultStatus": true,
            "enabled": true,
            "sendNotificationWhenStatusIsAssigned": true
        },
        {
            "statusId": "READY_FOR_PICKUP",
            "orderStatusType": "FULFILLMENT_STATUS",
            "defaultStatus": true,
            "enabled": true,
            "sendNotificationWhenStatusIsAssigned": true
        },
        {
            "statusId": "OUT_FOR_DELIVERY",
            "orderStatusType": "FULFILLMENT_STATUS",
            "defaultStatus": true,
            "enabled": true,
            "sendNotificationWhenStatusIsAssigned": true
        },
        {
            "statusId": "CUSTOM_FULFILLMENT_STATUS_1",
            "orderStatusType": "FULFILLMENT_STATUS",
            "enabled": false,
            "sendNotificationWhenStatusIsAssigned": true
        },
        {
            "statusId": "CUSTOM_FULFILLMENT_STATUS_2",
            "orderStatusType": "FULFILLMENT_STATUS",
            "enabled": false,
            "sendNotificationWhenStatusIsAssigned": true
        },
        {
            "statusId": "CUSTOM_FULFILLMENT_STATUS_3",
            "orderStatusType": "FULFILLMENT_STATUS",
            "enabled": false,
            "sendNotificationWhenStatusIsAssigned": true
        }
    ]
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_store_profile`

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

<table><thead><tr><th width="169">Field</th><th>Type</th><th>Description</th></tr></thead><tbody><tr><td>items</td><td>array of objects <a href="#items">items</a></td><td>Details about all order statuses.</td></tr></tbody></table>

#### items

<table><thead><tr><th width="193.26953125">Field</th><th width="114.42578125">Type</th><th>Description</th></tr></thead><tbody><tr><td>statusId</td><td>string</td><td><p>One of:</p><p><code>AWAITING_PAYMENT</code></p><p><code>PAID</code></p><p><code>CANCELLED</code></p><p><code>REFUNDED</code></p><p><code>PARTIALLY_REFUNDED</code></p><p><code>INCOMPLETE</code></p><p><code>CUSTOM_PAYMENT_STATUS_1</code></p><p><code>CUSTOM_PAYMENT_STATUS_2</code></p><p><code>CUSTOM_PAYMENT_STATUS_3</code></p><p><code>AWAITING_PROCESSING</code></p><p><code>PROCESSING</code></p><p><code>SHIPPED</code></p><p><code>DELIVERED</code></p><p><code>WILL_NOT_DELIVER</code></p><p><code>RETURNED</code></p><p><code>READY_FOR_PICKUP</code></p><p><code>OUT_FOR_DELIVERY</code></p><p><code>CUSTOM_FULFILLMENT_STATUS_1</code></p><p><code>CUSTOM_FULFILLMENT_STATUS_2</code></p><p><code>CUSTOM_FULFILLMENT_STATUS_3</code></p><p></p><p>Read more about order statuses in <a href="https://support.ecwid.com/hc/en-us/articles/207806235-Order-details-and-statuses-overview#-understanding-order-statuses">Help Center</a>.</p></td></tr><tr><td>orderStatusType</td><td>string</td><td><p>Defines if it's a payment or shipping status.<br><br>One of:</p><p><code>PAYMENT_STATUS</code></p><p><code>FULFILLMENT_STATUS</code></p></td></tr><tr><td>defaultStatus</td><td>boolean</td><td><p>Defines if it's a built-in status (true) or a custom one (<code>false</code>).<br><br>Custom order statuses:</p><p><code>CUSTOM_PAYMENT_STATUS_1</code></p><p><code>CUSTOM_PAYMENT_STATUS_2</code></p><p><code>CUSTOM_PAYMENT_STATUS_3</code></p><p><code>CUSTOM_FULFILLMENT_STATUS_1</code></p><p><code>CUSTOM_FULFILLMENT_STATUS_2</code></p><p><code>CUSTOM_FULFILLMENT_STATUS_3</code></p></td></tr><tr><td>enabled</td><td>boolean</td><td>Defines if the status is enabled and can be used in the store.</td></tr><tr><td>sendNotificationWhenStatusIsAssigned</td><td>boolean</td><td>Defines if the "Order status updated" email should be automatically sent to customers.</td></tr><tr><td>name</td><td>string</td><td>Name of the order status visible in Ecwid admin and emails.<br><br><strong>Available only for custom order statuses</strong> </td></tr><tr><td>nameTranslated</td><td>object <a href="#translations">translations</a></td><td>Available translations for the order status name.<br><br><strong>Available only for custom order statuses</strong></td></tr><tr><td>lastNameChangeDate</td><td>string</td><td>Datetime of the latest name change for custom order statuses in UTC +0.<br><br>For example, <code>2023-01-01 12:00:00 +0000</code><br><br><strong>Available only for custom order statuses</strong></td></tr></tbody></table>

#### translations

Object with text field translations in the `"lang": "text"` format, where the `"lang"` is an ISO 639-1 language code. For example:

```
{
    "en": "Sample text",
    "nl": "Voorbeeldtekst"
}
```

Translations are available for all active store languages. Only the default language translations are returned if no other translations are provided for the field. Find active store languages with <mark style="color:green;">`GET`</mark> `/profile` request > `languages` > `enabledLanguages`.
