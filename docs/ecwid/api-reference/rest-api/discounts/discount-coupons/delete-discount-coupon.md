# Delete discount coupon

<mark style="color:red;">`DELETE`</mark> `https://app.ecwid.com/api/v3/{storeId}/discount_coupons/{discountCouponId}`&#x20;

### Required access scopes

Your app must have the following **access scopes** to make this request: `update_discount_coupons`

### Path params

All path params are required.

| Param            | Type   | Description                  |
| ---------------- | ------ | ---------------------------- |
| storeId          | number | Ecwid store ID.              |
| discountCouponId | number | Internal discount coupon ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Field       | Type   | Description                                                                                                                                                                                   |
| ----------- | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| deleteCount | number | <p>The number of deleted items that defines if the request was successful.<br><br>One of:</p><p><code>1</code> if the item was deleted,</p><p><code>0</code> if the item was not deleted.</p> |
