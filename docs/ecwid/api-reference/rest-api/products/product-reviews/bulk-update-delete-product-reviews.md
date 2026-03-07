# Bulk update/delete product reviews

<mark style="color:purple;">`PUT`</mark> `https://app.ecwid.com/api/v3/{storeId}/reviews/mass_update`

<details>

<summary>Request and response example</summary>

Request:

```http
PUT /api/v3/1003/reviews/mass_update HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
Content-Type: application/json
Cache-Control: no-cache

{
  "selectMode": "ALL",
  "delete": false,
  "newStatus": "published"
}
```

Response:

```json
{
  "updateCount": 2
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `update_reviews`

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

<table><thead><tr><th width="159.7578125">Field</th><th width="183.33984375">Type</th><th>Description</th></tr></thead><tbody><tr><td>selectMode</td><td>string</td><td><p>Filtering behavior for finding product reviews to update/delete.</p><p></p><p>One of:</p><ul><li><code>SELECTED</code> - Request only updates the product reviews whose IDs are specified in the <code>reviewIds</code> array.</li><li><code>ALL_FILTERED</code> - Request updates the product reviews matching criteria in the <code>currentFilters</code> object. </li><li><code>ALL</code> - Request updates all reviews in the store without any filtering.</li></ul><p></p><p><strong>Required</strong></p></td></tr><tr><td>delete</td><td>boolean</td><td><p>Working mode for the request. It can either delete product reviews from the store or update their status (published/unpublished).</p><p></p><p>One of:</p><ul><li><code>true</code> - Request deletes product reviews.</li><li><code>false</code> - Request updates product reviews' status.</li></ul><p>Defaut value is <code>false</code>.</p></td></tr><tr><td>newStatus</td><td>string</td><td><p>Set new review status.<br><br>One of: </p><ul><li><code>MODERATED</code> - Product reviews become unpublished. Such reviews are not visible on the storefront.</li><li><code>PUBLISHED</code> - Product reviews become. Such reviews are visible on the storefront.</li></ul><p>Requires <code>"delete": false</code>. <br></p><p><strong>Case sensitive</strong></p></td></tr><tr><td>reviewIds</td><td>array of numbers</td><td>Specify the list of review IDs for the update as an array.<br><br>Requires <code>"selectMode": "SELECTED"</code>. Otherwise, request ingores this field.</td></tr><tr><td>currentFilters</td><td>object <a href="#currentfilters">currentFilters</a></td><td>Specify search criteria to find product reviews for the update.<br><br>Requires <code>"selectMode": "ALL_FILTERED"</code>. Otherwise, request ingores this field.</td></tr></tbody></table>

#### currentFilters

<table><thead><tr><th width="144.84375">Field</th><th width="157.83203125">Type</th><th>Description</th></tr></thead><tbody><tr><td>reviewId</td><td>array of numbers</td><td><p>Find reviews by their IDs. </p><p>For example, <code>[76259972, 97266752]</code>.</p></td></tr><tr><td>productId</td><td>array of numbers</td><td><p>Find reviews by product IDs. </p><p>For example, <code>[689454040, 692730761]</code>.</p></td></tr><tr><td>orderId</td><td>array of numbers</td><td><p>Find reviews by their assigned order IDs. Requires internal order IDs.</p><p>For example, <code>[82163452, 144937920]</code>.</p></td></tr><tr><td>rating</td><td>array of numbers</td><td>Find product reviews by their rating from 1 to 5. Supports multiple rating values.<br>For example, <code>[1, 2, 3]</code>.</td></tr><tr><td>status</td><td>string</td><td><p>Search reviews by their status. <br><br>One of: </p><ul><li><code>MODERATED</code> - Find unpublished product reviews invisible on the storefront.</li><li><code>PUBLISHED</code> - Find published product reviews visible on the storefront.</li></ul><p><strong>Case sensitive</strong></p></td></tr><tr><td>searchKeyword</td><td>string</td><td>Find product reviews by searching specific words in the review text left by customers.</td></tr><tr><td>createdFrom</td><td>string</td><td>Find reviews by their creation time (upper bound).<br>For example, <code>1744013600</code></td></tr><tr><td>createdTo</td><td>string</td><td>ind reviews by their creation time (lower bound).<br>For example, <code>1742110000</code></td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Field       | Type   | Description                                                                                                                                                                                   |
| ----------- | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| updateCount | number | <p>The number of updated items that defines if the request was successful.<br><br>One of:</p><p><code>1</code> if the item was updated,</p><p><code>0</code> if the item was not updated.</p> |
