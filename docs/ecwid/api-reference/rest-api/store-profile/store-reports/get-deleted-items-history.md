# Get deleted items history

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/{entity}/deleted`&#x20;

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_store_profile` , `read_orders`, `read_catalog`, `read_discount_coupons`, `read_customers`, `read_reviews`

### Path params

All path params are required.

| Param   | Type   | Description                                                                                                                                                                                  |
| ------- | ------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| storeId | number | Ecwid store ID.                                                                                                                                                                              |
| entity  | string | <p>Defines which item's history will be received. <br><br>Must be one of: <code>orders</code>, <code>products</code>, <code>coupons</code>, <code>customers</code>, <code>reviews</code></p> |

### Query params

All query params are optional.

<table data-full-width="false"><thead><tr><th width="133">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>from_date</td><td>string</td><td><p>Item deletion date lower bound. Supported formats:</p><ul><li>UNIX timestamp</li><li>yyyy-MM-dd HH:mm:ss Z</li><li>yyyy-MM-dd HH:mm:ss</li><li>yyyy-MM-dd</li></ul><p>Examples:</p><ul><li><code>1447804800</code></li><li><code>2015-04-22 18:48:38 -0500</code></li><li><code>2015-04-22</code> (matches 2015-04-22 00:00:00 UTC)</li></ul></td></tr><tr><td>to_date</td><td>string</td><td>Item deletion date upper bound.</td></tr><tr><td>offset</td><td>number</td><td>Offset from the beginning of the returned items list (for paging). Used when the number of returned items is higher than the <code>limit</code> of a single request.</td></tr><tr><td>limit</td><td>number</td><td>Maximum number of returned items. Default value: <code>100</code></td></tr></tbody></table>

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

<table><thead><tr><th width="118">Field</th><th width="158">Type</th><th>Description</th></tr></thead><tbody><tr><td>total</td><td>number</td><td>Total number of found items (might be more than the number of returned items).</td></tr><tr><td>count</td><td>number</td><td>Total number of items returned in the response.</td></tr><tr><td>offset</td><td>number</td><td>Offset from the beginning of the returned items list specified in the request.</td></tr><tr><td>limit</td><td>number</td><td>Maximum number of returned items specified in the request. Maximum allowed value: <code>100</code>. Default value: <code>10</code></td></tr><tr><td>items</td><td>array of objects <a href="#removeditem">removedItem</a></td><td>List of removed items where each "item" contains ID and a deletion date.</td></tr></tbody></table>

#### removedItem

<table><thead><tr><th width="186">Field</th><th width="157">Type</th><th>Description</th></tr></thead><tbody><tr><td>id</td><td>number</td><td>Internal item ID that depends on the request. For example, product ID, customer ID, etc.</td></tr><tr><td>date</td><td>string</td><td>Item deletion date.</td></tr></tbody></table>
