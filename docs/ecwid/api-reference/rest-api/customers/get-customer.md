# Get customer

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/customers/{customerId}`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
GET /api/v3/1003/customers/177737165 HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
```

Response:

```json
{
  "id": 177737165,
  "name": "Support team",
  "email": "ec.apps@lightspeedhq.com",
  "registered": "2021-12-21 06:05:58 +0000",
  "updated": "2024-06-04 21:15:10 +0000",
  "customerGroupId": 0,
  "customerGroupName": "General",
  "billingPerson": {
    "name": "Support team",
    "firstName": "Support",
    "lastName": "team"
  },
  "shippingAddresses": [],
  "contacts": [
    {
      "id": 113861381,
      "contact": "ec.apps@lightspeedhq.com",
      "type": "EMAIL",
      "default": true,
      "orderBy": 0,
      "timestamp": "2024-06-04 21:15:10 +0000"
    }
  ],
  "taxExempt": false,
  "taxId": "",
  "taxIdValid": true,
  "b2b_b2c": "b2c",
  "fiscalCode": "",
  "electronicInvoicePecEmail": "",
  "electronicInvoiceSdiCode": "",
  "acceptMarketing": false,
  "stats": {
    "numberOfOrders": 0,
    "salesValue": 0,
    "averageOrderValue": 0
  },
  "privateAdminNotes": "",
  "favorites": []
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_customers`&#x20;

### Path params

All path params are required.

| Param      | Type   | Description           |
| ---------- | ------ | --------------------- |
| storeId    | number | Ecwid store ID.       |
| customerId | number | Internal customer ID. |

### Query params

All query params are optional.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>responseFields</td><td>string</td><td>Specify the exact fields to receive in response JSON. If not specified, the response JSON will have all available fields for the entity.<br>Example: <code>?responseFields=id,name,email</code></td></tr></tbody></table>

Example of using `responseFields` param:

{% tabs %}
{% tab title="Request" %}

```
curl --location 'https://app.ecwid.com/api/v3/1003/customers/177737165?responseFields=id,name,email' \
--header 'Authorization: Bearer secret_ab***cd'
```

{% endtab %}

{% tab title="Response" %}

```json
{
    "id": 177737165,
    "name": "Support team",
    "email": "ec.apps@lightspeedhq.com"
}
```

{% endtab %}
{% endtabs %}

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

<table><thead><tr><th width="206.64453125">Field</th><th width="170.765625">Type</th><th>Description</th></tr></thead><tbody><tr><td>id</td><td>number</td><td>Unique internal customer ID.</td></tr><tr><td>email</td><td>string</td><td>Customer's email.</td></tr><tr><td>registered</td><td>string</td><td>Customer's registration datetime, for example, <code>2014-06-06 18:57:19 +0400</code></td></tr><tr><td>updated</td><td>string</td><td>Datetime of the latest update of customer's details, for example, <code>2014-06-06 18:57:19 +0400</code></td></tr><tr><td>billingPerson</td><td>object <a href="#billingperson">billingPerson</a></td><td>Customer's billing name/address.</td></tr><tr><td>shippingAddresses</td><td>array of objects <a href="#shippingaddresses">shippingAddresses</a></td><td>List of saved shipping addresses for the customer.</td></tr><tr><td>contacts</td><td>array of objects <a href="#contacts">contacts</a></td><td>Customer's contact information: email, phone, social media links.</td></tr><tr><td>customerGroupName</td><td>string</td><td>Name of the customer group the customer is assigned to.</td></tr><tr><td>customerGroupId</td><td>number</td><td>ID of the customer group the customer is assigned to.</td></tr><tr><td>b2b_b2c</td><td>string</td><td>Defines business-to-customer relation. One of:<br><code>b2c</code> - Business-to-customer (default)<br><code>b2b</code> - Business-to-business</td></tr><tr><td>taxId</td><td>string</td><td>Customer's tax ID.</td></tr><tr><td>taxIdValid</td><td>boolean</td><td>Defines if customer's tax ID is valid.</td></tr><tr><td>taxExempt</td><td>boolean</td><td>Defines if customer is tax exempt. Requires a valid tax ID.<br><br>Read more about handling tax exempt customers in <a href="https://support.ecwid.com/hc/en-us/articles/213823045-Handling-tax-exempt-customers">Help Center</a>.</td></tr><tr><td>acceptMarketing</td><td>boolean</td><td>Defines if the customer has accepted email marketing. <br><br>If <code>true</code>, you can use customer's email for promotions.</td></tr><tr><td>lang</td><td>string</td><td>Customer's language code. Customers see storefront and emails in this language.<br><br>This language must be one of the translations enabled in the store.</td></tr><tr><td>stats</td><td>object <a href="#stats">stats</a></td><td>Customer's sales stats: number of orders, total revenue, first order date, etc.</td></tr><tr><td>privateAdminNotes</td><td>string</td><td>Personal notes about the customer. Visible only to the store owner.</td></tr><tr><td>favorites</td><td>array of objects <a href="#favorites">favorites</a></td><td>List of customer's favorite products.</td></tr></tbody></table>

#### billingPerson

| Field               | Type   | Description                                                                          |
| ------------------- | ------ | ------------------------------------------------------------------------------------ |
| name                | string | Full name of the customer.                                                           |
| firstName           | string | First name of the customer. Only shows when the `name` field has at least two words. |
| lastName            | string | Last name of the customer. Only shows when the `name` field has at least two words.  |
| companyName         | string | Customer's company name.                                                             |
| street              | string | Address line 1 and address line 2, separated by `\n`.                                |
| city                | string | City.                                                                                |
| countryCode         | string | Two-letter country code.                                                             |
| countryName         | string | Country name.                                                                        |
| postalCode          | string | Postal/ZIP code.                                                                     |
| stateOrProvinceCode | string | State/province code, for example, `NY`.                                              |
| stateOrProvinceName | string | State/province name.                                                                 |
| phone               | string | Customer's phone number.                                                             |

#### shippingAddresses

| Field               | Type   | Description                                                              |
| ------------------- | ------ | ------------------------------------------------------------------------ |
| id                  | number | Internal ID of the saved address.                                        |
| name                | string | Full name of the customer.                                               |
| companyName         | string | Customer's company name.                                                 |
| street              | string | Address line 1 and address line 2, separated by `\n`.                    |
| city                | string | City.                                                                    |
| countryCode         | string | Two-letter country code.                                                 |
| countryName         | string | Country name.                                                            |
| postalCode          | string | Postal/ZIP code.                                                         |
| stateOrProvinceCode | string | State/province code, for example, `NY`.                                  |
| stateOrProvinceName | string | State/province name.                                                     |
| phone               | string | Customer's phone number.                                                 |
| addressFormatted    | string | Formatted full address. Includes street, city, state, and country names. |

#### contacts

<table><thead><tr><th width="167.4375">Field</th><th width="109.3515625">Type</th><th>Description</th></tr></thead><tbody><tr><td>contact</td><td>string</td><td><p>Email or link to reach the contact. Examples:</p><ul><li><code>ec.apps@lightspeedhq.com</code> contact for <code>EMAIL</code> type.</li><li><code>https://www.facebook.com/myshop_page</code> contact for <code>FACEBOOK</code> type.</li></ul><p><strong>Required</strong></p></td></tr><tr><td>handle</td><td>string</td><td>Contact identifier on social media. For example, for <code>FACEBOOK</code> type of contact, it's a page slug:<br><br><code>contact</code> field: <code>https://www.facebook.com/myshop_page</code> <br><code>handle</code> field: <code>myshop_page</code></td></tr><tr><td>note</td><td>string</td><td>Store owner's notes on the contact.</td></tr><tr><td>type</td><td>string</td><td><p>Contact type. Customer can have several contacts of the same type.<br><br>One of:</p><p><code>EMAIL</code>, <br><code>PHONE</code>,<br><code>FACEBOOK</code>,<br><code>INSTAGRAM</code>,<br><code>TWITTER</code>,<br><code>YOUTUBE</code>,<br><code>TIKTOK</code>,<br><code>PINTEREST</code>,<br><code>VK</code>,<br><code>FB_MESSENGER</code>,<br><code>WHATSAPP</code>,<br><code>TELEGRAM</code>,<br><code>VIBER</code>,<br><code>URL</code>,<br><code>OTHER</code>.<br><br><strong>Required</strong></p></td></tr><tr><td>default</td><td>boolean</td><td>Defines if it's a default customer contact. Only one contact of the same type can be default.</td></tr><tr><td>orderBy</td><td>boolean</td><td>Sorting order for contacts on the customer details page. Starts with <code>0</code> and increments by <code>1</code>.</td></tr><tr><td>timestamp</td><td>string</td><td>Datetime when the customer contact was created.</td></tr></tbody></table>

#### stats

| Field             | Type   | Description                                     |
| ----------------- | ------ | ----------------------------------------------- |
| numberOfOrders    | number | Count of customer's orders in the store.        |
| salesValue        | number | Total cost of orders placed by the customer.    |
| averageOrderValue | number | Average total of orders placed by the customer. |
| firstOrderDate    | string | Date the customer placed their first order.     |
| lastOrderDate     | string | Date the customer placed their last order.      |

#### favorites

| Field          | Type   | Description                                                                             |
| -------------- | ------ | --------------------------------------------------------------------------------------- |
| productId      | number | Internal ID of the favorited product, for example, `689454040`                          |
| addedTimestamp | string | Datetime when the product was added to favorites, favorites `2024-09-11 06:43:02 +0000` |
