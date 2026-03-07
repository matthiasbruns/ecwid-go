# Update staff account

<mark style="color:purple;">`PUT`</mark> `https://app.ecwid.com/api/v3/{storeId}/staff/{staffAccountId}`&#x20;

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_staff` , `update_staff`

### Path params

All path params are required.

| Param          | Type   | Description                |
| -------------- | ------ | -------------------------- |
| storeId        | number | Ecwid store ID.            |
| staffAccountId | string | Internal staff account ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Request JSON

A JSON object with the following fields:

| Field       | Type             | Description                                                                                                                                                                                                                                                                                  |
| ----------- | ---------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| email       | string           | Staff account email.                                                                                                                                                                                                                                                                         |
| staffScopes | array of strings | <p>Permissions enabled for the staff account. If empty, the account has all permissions. <br><br>Learn more about staff account permissions in <a href="https://support.ecwid.com/hc/en-us/articles/115005355089-Adding-and-managing-staff-accounts#-staff-permissions">Help Center</a>.</p> |

### Response JSON

A JSON object with the following fields:

| Field       | Type   | Description                                                                                                                                                                             |
| ----------- | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| updateCount | number | <p>The number of updated items that defines if the request was successful.<br><br>One of:<br><code>1</code> if the item was updated,<br><code>0</code> if the item was not updated.</p> |
