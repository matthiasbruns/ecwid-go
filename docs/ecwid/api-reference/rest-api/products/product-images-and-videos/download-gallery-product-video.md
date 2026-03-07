# Download gallery product video

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/products/{productId}/gallery/video/{galleryVideoId}`

### Required access scopes

Your app must have the following **access scopes** to make this request: `update_catalog`

### Path params

All path params are required.

| Param          | Type   | Description                               |
| -------------- | ------ | ----------------------------------------- |
| storeId        | number | Ecwid store ID.                           |
| productId      | number | Internal product ID.                      |
| galleryVideoId | number | Internal video ID in the product gallery. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Field          | Type    | Description                                             |
| -------------- | ------- | ------------------------------------------------------- |
| id             | number  | Internal video ID                                       |
| videoCoverId   | number  | ID of the cover image for the video (from 'images')     |
| url            | boolean | URL to the video file                                   |
| embedHtml      | string  | Embedded code for the video file                        |
| providerName   | string  | Video hosting provider name (could be empty)            |
| title          | string  | Video title (could be empty)                            |
| image160pxUrl  | string  | URL of the video cover image resized to fit 160x160px   |
| image400pxUrl  | string  | URL of the video cover image resized to fit 400x400px   |
| image800pxUrl  | string  | URL of the video cover image resized to fit 800x800px   |
| image1500pxUrl | string  | URL of the video cover image resized to fit 1500x1500px |
