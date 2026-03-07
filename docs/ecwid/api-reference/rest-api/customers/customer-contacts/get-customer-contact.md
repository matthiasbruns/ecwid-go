# Get customer contact

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/customers/{customerId}/contacts/{contactId}`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
GET /api/v3/1003/customers/177737165/contacts/113861381 HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
```

Response:

```json
{
  "id": 113861381,
  "contact": "ec.apps@lightspeedhq.com",
  "type": "EMAIL",
  "default": true,
  "orderBy": 0,
  "timestamp": "2025-03-26 11:07:12 +0000"
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_customers`

### Path params

All path params are required.

| Param      | Type   | Description                   |
| ---------- | ------ | ----------------------------- |
| storeId    | number | Ecwid store ID.               |
| customerId | number | Internal customer ID.         |
| contactId  | number | Internal customer contact ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

<table><thead><tr><th width="167.4375">Field</th><th width="109.3515625">Type</th><th>Description</th></tr></thead><tbody><tr><td>id</td><td>number</td><td>Internal ID of the customer contact, for example, <code>113861381</code>.</td></tr><tr><td>contact</td><td>string</td><td><p>Email or link to reach the contact. Examples:</p><ul><li><code>ec.apps@lightspeedhq.com</code> contact for <code>EMAIL</code> type.</li><li><code>https://www.facebook.com/myshop_page</code> contact for <code>FACEBOOK</code> type.</li></ul></td></tr><tr><td>handle</td><td>string</td><td>Contact identifier on social media. For example, for <code>FACEBOOK</code> type of contact, it's a page slug:<br><br><code>contact</code> field: <code>https://www.facebook.com/myshop_page</code> <br><code>handle</code> field: <code>myshop_page</code></td></tr><tr><td>note</td><td>string</td><td>Store owner's notes on the contact.</td></tr><tr><td>type</td><td>string</td><td><p>Contact type. Customer can have several contacts of the same type.<br><br>One of:</p><p><code>EMAIL</code>, <br><code>PHONE</code>,<br><code>FACEBOOK</code>,<br><code>INSTAGRAM</code>,<br><code>TWITTER</code>,<br><code>YOUTUBE</code>,<br><code>TIKTOK</code>,<br><code>PINTEREST</code>,<br><code>VK</code>,<br><code>FB_MESSENGER</code>,<br><code>WHATSAPP</code>,<br><code>TELEGRAM</code>,<br><code>VIBER</code>,<br><code>URL</code>,<br><code>OTHER</code>.</p></td></tr><tr><td>default</td><td>boolean</td><td>Defines if it's a default customer contact. Only one contact of the same type can be default.</td></tr><tr><td>orderBy</td><td>boolean</td><td>Sorting order for contacts on the customer details page. Starts with <code>0</code> and increments by <code>1</code>.</td></tr><tr><td>timestamp</td><td>string</td><td>Datetime when the customer contact was created.</td></tr></tbody></table>
