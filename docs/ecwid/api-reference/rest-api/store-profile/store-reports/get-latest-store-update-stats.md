# Get latest store update stats

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/latest-stats`&#x20;

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_store_profile`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

All query params are optional.

<table data-full-width="false"><thead><tr><th width="247">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>reviewsUpdatesRequired</td><td>boolean</td><td>Send <code>true</code> to get the datetime of the latest review update in the store.</td></tr><tr><td>domainsRequired</td><td>boolean</td><td>Send <code>true</code> to get stats about the latest store domain update.</td></tr><tr><td>subscriptionRequired</td><td>boolean</td><td>Send <code>true</code> to get stats about the latest store subscription updates.</td></tr><tr><td>productCountRequired</td><td>boolean</td><td>Send <code>true</code> to get the number of products in the store (excluding demo products).</td></tr><tr><td>categoryCountRequired</td><td>boolean</td><td>Send <code>true</code> to get the number of categories in the store.</td></tr></tbody></table>

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Name                   | Type   | Description                                                                                          |
| ---------------------- | ------ | ---------------------------------------------------------------------------------------------------- |
| productsUpdated        | string | Date of the latest changes in store catalog (products, categories), e.g. `2014-10-15 16:54:11 +0400` |
| ordersUpdated          | string | Date of the latest changes in store orders, e.g. `2014-10-15 16:54:11 +0400`                         |
| reviewsUpdated         | string | Date of the latest changes in product reviews, e.g. `2024-10-15 16:54:11 +0400`                      |
| domainsUpdated         | string | Date of the latest changes in store domains, e.g. `2021-04-27 13:13:55 +0000`                        |
| profileUpdated         | string | Date of the latest changes in store information, e.g. `2014-10-15 16:54:11 +0400`                    |
| categoriesUpdated      | string | Date of the latest changes in store categories, e.g. `2014-10-19 12:23:12 +0400`                     |
| discountCouponsUpdated | string | Date of the latest changes in store discount coupons, e.g. `2014-10-19 12:23:12 +0400`               |
| abandonedSalesUpdated  | string | Date of the latest changes to abandoned carts in a store, e.g. `2014-10-19 12:23:12 +0400`           |
| customersUpdated       | string | Date of the latest changes to customers in a store, e.g. `2014-10-19 12:23:12 +0400`                 |
| subscriptionsUpdated   | string | Date of the latest changes to subscriptions in a store, e.g. `2021-04-27 13:13:55 +0000`             |
| productCountRequired   | number | The number of all products in the store (demo excluded). See above to get it.                        |
| categoryCountRequired  | number | The number of all categories in the store. See above how to get it.                                  |
