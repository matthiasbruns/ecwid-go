# Upload cover for gallery video

<mark style="color:purple;">`PUT`</mark> `https://app.ecwid.com/api/v3/{storeId}/products/{productId}/gallery/video/{galleryVideoId}`

### Required access scopes

Your app must have the following **access scopes** to make this request: `update_catalog`

### Path params

All path params are required.

| Param          | Type   | Description                               |
| -------------- | ------ | ----------------------------------------- |
| storeId        | number | Ecwid store ID.                           |
| productId      | number | Internal product ID.                      |
| galleryVideoId | number | Internal video ID in the product gallery. |

### Query params

All query params are optional.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>externalUrl</td><td>string</td><td>HTTPS link to the image file that will be uploaded to the store.<br><br>Alternatively, you can send the image as binary data in the request body.</td></tr></tbody></table>

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Field | Type   | Description                        |
| ----- | ------ | ---------------------------------- |
| id    | number | Internal ID of the uploaded image. |
