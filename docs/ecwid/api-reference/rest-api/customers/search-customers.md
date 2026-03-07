# Search customers

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/customers`&#x20;

<details>

<summary>Request and response example</summary>

Request:

```http
GET /api/v3/1003/customers HTTP/1.1
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
      "id": 177737165,
      "name": "Support team",
      "email": "ec.apps@lightspeedhq.com",
      "registered": "2021-12-21 06:05:58 +0000",
      "updated": "2024-06-04 21:15:10 +0000",
      "totalOrderCount": 0,
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
    },
    {
      "id": 277137633,
      "name": "Ecwid test acc",
      "email": "ecwid.test@ecwid-test.com",
      "registered": "2024-08-21 07:52:09 +0000",
      "updated": "2024-08-21 07:52:09 +0000",
      "totalOrderCount": 0,
      "customerGroupId": 0,
      "customerGroupName": "General",
      "contacts": [
        {
          "id": 176135453,
          "contact": "00000000000",
          "type": "PHONE",
          "default": true,
          "orderBy": 0,
          "timestamp": "2024-08-21 07:52:09 +0000"
        },
        {
          "id": 176135452,
          "contact": "ecwid.test@ecwid-test.com",
          "type": "EMAIL",
          "default": true,
          "orderBy": 0,
          "timestamp": "2024-08-21 07:52:09 +0000"
        }
      ],
      "taxExempt": false,
      "taxId": "",
      "taxIdValid": true,
      "b2b_b2c": "b2c",
      "fiscalCode": "",
      "electronicInvoicePecEmail": "",
      "electronicInvoiceSdiCode": "",
      "lang": "en",
      "stats": {
        "numberOfOrders": 1,
        "salesValue": 97,
        "averageOrderValue": 97,
        "firstOrderDate": "2024-08-21 07:52:09 +0000",
        "lastOrderDate": "2024-08-21 07:52:09 +0000"
      },
      "privateAdminNotes": "",
      "favorites": []
    }
  ],
  "allCustomerCount": 2
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_customers`&#x20;

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

All query params are optional.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>keyword</td><td>string</td><td>Search term for both customer name and email. <br><br>Special characters must be URI-encoded.</td></tr><tr><td>name</td><td>string</td><td>Search term for customer's name – <code>billingPerson.name</code> field.</td></tr><tr><td>email</td><td>string</td><td>Search term for customer's email.</td></tr><tr><td>useExactEmailMatch</td><td>boolean</td><td>If <code>true</code>, search for exact email match. Requires <code>email</code> query param to work.</td></tr><tr><td>phone</td><td>string</td><td>Search term for customer's phone number or its part in both <code>shippingAddress</code> and <code>contacts</code>. Checks all saved addresses and contacts.</td></tr><tr><td>city</td><td>string</td><td>Search term for customer's city in <code>shippingAddress</code>. Checks all saved addresses.</td></tr><tr><td>postalCode</td><td>string</td><td>Search term for customer's ZIP code in <code>shippingAddress</code>. Checks all saved addresses.</td></tr><tr><td>stateOrProvinceCode</td><td>string</td><td>Search term for customer's two-digit state code in <code>shippingAddress</code>. Checks all saved addresses.</td></tr><tr><td>countryCodes</td><td>string</td><td>Search term for customer's country codes in <code>shippingAddress</code>. Checks all saved addresses.</td></tr><tr><td>companyName</td><td>string</td><td>Search term for customer's company name in <code>shippingAddress</code>. Checks all saved addresses.</td></tr><tr><td>acceptMarketing</td><td>boolean</td><td>Set <code>true</code> to <strong>only</strong> find customers who accepted receiving marketing emails. <br><br>Set <code>false</code> to <strong>only</strong> find customers who rejected such emails.<br><br>Do not add param to the request to receive both.</td></tr><tr><td>lang</td><td>string</td><td>Search term for customer's languages used for making orders.</td></tr><tr><td>customerGroupIds</td><td>string</td><td>Search term for customer group ID. Supports multiple values, for example, <code>123456,234567</code>.</td></tr><tr><td>minOrderCount</td><td>number</td><td>Search by the minimum number of customer's orders.</td></tr><tr><td>maxOrderCount</td><td>number</td><td>Search by the maximum number of customer's orders.</td></tr><tr><td>minSalesValue</td><td>number</td><td>Search by the minimum total order value of customer's orders.</td></tr><tr><td>maxSalesValue</td><td>number</td><td>Search by the maximum total order value of customer's orders.</td></tr><tr><td>purchasedProductIds</td><td>string</td><td>Search term for product IDs in customer's orders.</td></tr><tr><td>b2b_b2c</td><td>string</td><td>Defines business-to-customer relation. One of:<br><code>b2c</code> - Business-to-customer (default)<br><code>b2b</code> - Business-to-business</td></tr><tr><td>taxExempt</td><td>boolean</td><td>Set <code>true</code> to <strong>only</strong> find tax-exempt customers. <br><br>Set <code>false</code> to <strong>only</strong> find customers without tax exemption.<br><br>Do not add param to the request to receive both.</td></tr><tr><td>createdFrom</td><td>number/string</td><td>Datetime when a customer registered in the store or placed their first order without registration (lower bound). Supported formats: UNIX timestamp, datetime. <br><br>Examples: <code>1447804800</code>, <code>2023-01-15 19:27:50</code></td></tr><tr><td>createdTo</td><td>number/string</td><td>Datetime when a customer registered in the store or placed their first order without registration (upper bound). Supported formats: UNIX timestamp, datetime. <br><br>Examples: <code>1447804800</code>, <code>2023-01-15 19:27:50</code></td></tr><tr><td>updatedFrom</td><td>number/string</td><td>Datetime of the latest update of customer's details (lower bound. Supported formats: UNIX timestamp, datetime. <br><br>Examples: <code>1447804800</code>, <code>2023-01-15 19:27:50</code></td></tr><tr><td>updatedTo</td><td>number/string</td><td>Datetime of the latest update of customer's details (upper bound. Supported formats: UNIX timestamp, datetime. <br><br>Examples: <code>1447804800</code>, <code>2023-01-15 19:27:50</code></td></tr><tr><td>sortBy</td><td>string</td><td>Sorting order for the results. <br><br>One of: <br><code>NAME_ASC</code><br><code>NAME_DESC</code><br><code>EMAIL_ASC</code><br><code>EMAIL_DESC</code><br><code>ORDER_COUNT_ASC</code><br><code>ORDER_COUNT_DESC</code><br><code>REGISTERED_DATE_DESC</code><br><code>REGISTERED_DATE_ASC</code><br><code>UPDATED_DATE_DESC</code><br><code>UPDATED_DATE_ASC</code><br><code>SALES_VALUE_ASC</code><br><code>SALES_VALUE_DESC</code><br><code>FIRST_ORDER_DATE_ASC</code><br><code>FIRST_ORDER_DATE_DESC</code><br><code>LAST_ORDER_DATE_ASC</code><br><code>LAST_ORDER_DATE_DESC</code></td></tr><tr><td>offset</td><td>number</td><td><p>Offset from the beginning of the returned items list. Used when the response contains more items than <code>limit</code> allows to receive in one request.<br><br>Usually used to receive all items in several requests with multiple of a hundred, for example:<br><code>?offset=0</code> for the first request,</p><p><code>?offset=100</code>, for the second request,</p><p><code>?offset=200</code>, for the third request, etc.</p></td></tr><tr><td>limit</td><td>number</td><td>Limit to the number of returned items. Maximum and default value (if not specified) is <code>100</code>.</td></tr><tr><td>responseFields</td><td>string</td><td>Specify the exact fields to receive in response JSON. If not specified, the response JSON will have all available fields for the entity.<br>Example: <code>?responseFields=total,items(id,name,email)</code></td></tr></tbody></table>

Example of using `responseFields` param:

{% tabs %}
{% tab title="Request" %}

```
curl --location 'https://app.ecwid.com/api/v3/1003/customers?responseFields=total,items(id,name,email)' \
--header 'Authorization: Bearer secret_ab***cd'
```

{% endtab %}

{% tab title="Response" %}

```json
{
    "total": 2,
    "items": [
        {
            "id": 177737165,
            "name": "Support team",
            "email": "ec.apps@lightspeedhq.com"
        },
        {
            "id": 277137633,
            "name": "Ecwid test acc",
            "email": "ecwid.test@ecwid-test.com"
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

| Field            | Type                             | Description                                                                                  |
| ---------------- | -------------------------------- | -------------------------------------------------------------------------------------------- |
| total            | number                           | Total number of found items (might be more than the number of returned items).               |
| count            | number                           | Total number of items returned in the response.                                              |
| offset           | number                           | Offset from the beginning of the returned items list specified in the request.               |
| limit            | number                           | Maximum number of returned items specified in the request. Maximum and default value: `100`. |
| items            | array of objects [items](#items) | Detailed information about returned customers.                                               |
| allCustomerCount | number                           | Total count of unique customers in the store.                                                |

#### items

<table><thead><tr><th width="198.36328125">Field</th><th width="167.0625">Type</th><th>Description</th></tr></thead><tbody><tr><td>id</td><td>number</td><td>Unique internal customer ID.</td></tr><tr><td>email</td><td>string</td><td>Customer's email.</td></tr><tr><td>name</td><td>string</td><td>Customer's name. Duplicates <code>billingPerson.name</code> field.</td></tr><tr><td>totalOrderCount</td><td>number</td><td>Total count of orders placed by the customer.</td></tr><tr><td>registered</td><td>string</td><td>Customer's registration datetime, for example, <code>2014-06-06 18:57:19 +0400</code></td></tr><tr><td>updated</td><td>string</td><td>Datetime of the latest update of customer's details, for example, <code>2014-06-06 18:57:19 +0400</code></td></tr><tr><td>billingPerson</td><td>object <a href="#billingperson">billingPerson</a></td><td>Customer's billing name/address.</td></tr><tr><td>shippingAddresses</td><td>array of objects <a href="#shippingaddresses">shippingAddresses</a></td><td>List of saved shipping addresses for the customer.</td></tr><tr><td>contacts</td><td>array of objects <a href="#contacts">contacts</a></td><td>Customer's contact information: email, phone, social media links.</td></tr><tr><td>customerGroupName</td><td>string</td><td>Name of the customer group the customer is assigned to.</td></tr><tr><td>customerGroupId</td><td>number</td><td>ID of the customer group the customer is assigned to.</td></tr><tr><td>taxId</td><td>string</td><td>Customer's tax ID.</td></tr><tr><td>taxIdValid</td><td>boolean</td><td>Defines if customer's tax ID is valid.</td></tr><tr><td>taxExempt</td><td>boolean</td><td>Defines if customer is tax exempt. Requires a valid tax ID.<br><br>Read more about handling tax exempt customers in <a href="https://support.ecwid.com/hc/en-us/articles/213823045-Handling-tax-exempt-customers">Help Center</a>.</td></tr><tr><td>acceptMarketing</td><td>boolean</td><td>Defines if the customer has accepted email marketing. <br><br>If <code>true</code>, you can use customer's email for promotions.</td></tr><tr><td>lang</td><td>string</td><td>Customer's language code. Customers see storefront and emails in this language.<br><br>This language must be one of the translations enabled in the store.</td></tr><tr><td>stats</td><td>object <a href="#stats">stats</a></td><td>Customer's sales stats: number of orders, total revenue, first order date, etc.</td></tr><tr><td>privateAdminNotes</td><td>string</td><td>Personal notes about the customer. Visible only to the store owner.</td></tr><tr><td>favorites</td><td>array of objects <a href="#favorites">favorites</a></td><td>List of customer's favorite products.</td></tr></tbody></table>

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
