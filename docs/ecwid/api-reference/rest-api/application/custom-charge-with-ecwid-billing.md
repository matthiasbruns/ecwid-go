# Custom charge with Ecwid billing

<mark style="color:blue;">`POST`</mark> `https://app.ecwid.com/api/v3/{storeId}/billing/transactions`&#x20;

### Required access scopes

Your app must have the following **access scopes** to make this request: `charge`&#x20;

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

Some query params are **required**.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>amount</td><td>number</td><td>Charge amount that can be rounded to two decimal points, for example, <code>24.99</code>.<br><br>Maximum charge amount <em>in one request</em> is equal to <code>500</code> USD.<br>Maximum charge amount <em>in one day for one store</em> is equal to <code>5000</code> USD.<br><br><strong>Required</strong></td></tr><tr><td>currency</td><td>string</td><td><p>Charge currency. You can charge in any available currency, even if a store uses another one.</p><p><br>One of: <code>USD</code>, <code>EUR</code>, <code>MXN</code>, <code>INR</code>, <code>GBP</code>, <code>AUD</code>. <br><br><strong>Required</strong></p></td></tr><tr><td>description</td><td>string</td><td>Describe what functionality is covered by the charge. Limit: 255 characters.<br><br><strong>Required</strong></td></tr><tr><td>idempotencyKey</td><td>string</td><td>Generate a unique UUID key to guarantee there is no double charge.<br><br><strong>Required</strong></td></tr><tr><td>metadata</td><td>json</td><td>A JSON object for reference. For example, pass an ID of a bought product.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Name                | Type    | Description                                                                                                                                                                                                 |
| ------------------- | ------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| transactionId       | string  | Internal ID of the transaction. Use it to identify a specific charge in case of any issues.                                                                                                                 |
| idempotencyKeyInUse | boolean | <p>Defines if this was a duplicate request and Ecwid didn't execute it (<code>true</code>).<br><br>If <code>false</code>, the billing system accepted UUID and successfully completed the transaction. </p> |
