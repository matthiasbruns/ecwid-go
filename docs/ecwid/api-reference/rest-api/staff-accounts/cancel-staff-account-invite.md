# Cancel staff account invite

<mark style="color:red;">`DELETE`</mark> `https://app.ecwid.com/api/v3/{storeId}/staff/invite`&#x20;

### Required access scopes

Your app must have the following **access scopes** to make this request: `invite_staff`

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

| Field | Type   | Description                                                  |
| ----- | ------ | ------------------------------------------------------------ |
| email | string | <p>Staff account email.<br><br><strong>Required</strong></p> |

### Response JSON

A JSON object with the following fields:

| Field       | Type   | Description                                                                                                                            |
| ----------- | ------ | -------------------------------------------------------------------------------------------------------------------------------------- |
| deleteCount | number | <p>Defines if the invite was cancelled.<br><br>One of:<br><code>1</code> if the invite was cancelled,<br><code>0</code> if wasn't.</p> |
