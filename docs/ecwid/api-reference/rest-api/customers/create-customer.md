# Create customer

<mark style="color:blue;">`POST`</mark> `https://app.ecwid.com/api/v3/{storeId}/customers`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
POST /api/v3/1003/customers HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
Content-Type: application/json
Cache-Control: no-cache

{
  "name": "Example customer",
  "email": "example_customer@example.com",
  "customerGroupId": 0,
  "customerGroupName": "General",
  "billingPerson": {
    "name": "Support team"
  },
  "taxExempt": false
}
```

Response:

```json
{
  "id": 188278629
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

<table><thead><tr><th width="197.37109375">Field</th><th width="174.0703125">Type</th><th>Description</th></tr></thead><tbody><tr><td>email</td><td>string</td><td>Customer's email.<br><br><strong>Required</strong></td></tr><tr><td>password</td><td>string</td><td>Customer's password for login.<br><br>Note: only available for stores with the legacy customer sign-in.</td></tr><tr><td>billingPerson</td><td>object <a href="#billingperson">billingPerson</a></td><td>Customer's billing name/address.</td></tr><tr><td>shippingAddresses</td><td>array of objects <a href="#shippingaddresses">shippingAddresses</a></td><td>List of saved shipping addresses for the customer.</td></tr><tr><td>contacts</td><td>array of objects <a href="#contacts">contacts</a></td><td>Customer's contact information: email, phone, social media links.</td></tr><tr><td>customerGroupId</td><td>number</td><td>ID of the customer group the customer is assigned to.</td></tr><tr><td>b2b_b2c</td><td>string</td><td>Defines business-to-customer relation. One of:<br><code>b2c</code> - Business-to-customer (default)<br><code>b2b</code> - Business-to-business</td></tr><tr><td>taxId</td><td>string</td><td>Customer's tax ID.</td></tr><tr><td>taxIdValid</td><td>boolean</td><td>Defines if customer's tax ID is valid.</td></tr><tr><td>taxExempt</td><td>boolean</td><td>Defines if customer is tax exempt. Requires a valid tax ID.<br><br>Read more about handling tax exempt customers in <a href="https://support.ecwid.com/hc/en-us/articles/213823045-Handling-tax-exempt-customers">Help Center</a>.</td></tr><tr><td>acceptMarketing</td><td>boolean</td><td>Defines if the customer has accepted email marketing. <br><br>If <code>true</code>, you can use customer's email for promotions.</td></tr><tr><td>lang</td><td>string</td><td>Customer's language code. Customers see storefront and emails in this language.<br><br>This language must be one of the translations enabled in the store.</td></tr><tr><td>privateAdminNotes</td><td>string</td><td>Personal notes about the customer. Visible only to the store owner.</td></tr></tbody></table>

#### billingPerson

| Field               | Type   | Description                                           |
| ------------------- | ------ | ----------------------------------------------------- |
| name                | string | Full name of the customer.                            |
| companyName         | string | Customer's company name.                              |
| street              | string | Address line 1 and address line 2, separated by `\n`. |
| city                | string | City.                                                 |
| countryCode         | string | Two-letter country code.                              |
| countryName         | string | Country name.                                         |
| postalCode          | string | Postal/ZIP code.                                      |
| stateOrProvinceCode | string | State/province code, for example, `NY`.               |
| phone               | string | Customer's phone number.                              |

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
| phone               | string | Customer's phone number.                                                 |
| addressFormatted    | string | Formatted full address. Includes street, city, state, and country names. |

#### contacts

<table><thead><tr><th width="167.4375">Field</th><th width="109.3515625">Type</th><th>Description</th></tr></thead><tbody><tr><td>id</td><td>number</td><td>Internal ID of the customer contact, for example, <code>113861381</code>.</td></tr><tr><td>contact</td><td>string</td><td><p>Email or link to reach the contact. Examples:</p><ul><li><code>ec.apps@lightspeedhq.com</code> contact for <code>EMAIL</code> type.</li><li><code>https://www.facebook.com/myshop_page</code> contact for <code>FACEBOOK</code> type.</li></ul></td></tr><tr><td>handle</td><td>string</td><td>Contact identifier on social media. For example, for <code>FACEBOOK</code> type of contact, it's a page slug:<br><br><code>contact</code> field: <code>https://www.facebook.com/myshop_page</code> <br><code>handle</code> field: <code>myshop_page</code></td></tr><tr><td>note</td><td>string</td><td>Store owner's notes on the contact.</td></tr><tr><td>type</td><td>string</td><td><p>Contact type. Customer can have several contacts of the same type.<br><br>One of:</p><p><code>EMAIL</code>, <br><code>PHONE</code>,<br><code>FACEBOOK</code>,<br><code>INSTAGRAM</code>,<br><code>TWITTER</code>,<br><code>YOUTUBE</code>,<br><code>TIKTOK</code>,<br><code>PINTEREST</code>,<br><code>VK</code>,<br><code>FB_MESSENGER</code>,<br><code>WHATSAPP</code>,<br><code>TELEGRAM</code>,<br><code>VIBER</code>,<br><code>URL</code>,<br><code>OTHER</code>.</p></td></tr><tr><td>default</td><td>boolean</td><td>Defines if it's a default customer contact. Only one contact of the same type can be default.</td></tr><tr><td>orderBy</td><td>boolean</td><td>Sorting order for contacts on the customer details page. Starts with <code>0</code> and increments by <code>1</code>.</td></tr><tr><td>timestamp</td><td>string</td><td>Datetime when the customer contact was created.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Field | Type   | Description                          |
| ----- | ------ | ------------------------------------ |
| id    | number | Internal ID of the created customer. |
