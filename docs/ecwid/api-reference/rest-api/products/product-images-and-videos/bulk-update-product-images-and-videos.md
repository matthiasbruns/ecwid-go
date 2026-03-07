# Bulk update product images and videos

<mark style="color:purple;">`PUT`</mark> `https://app.ecwid.com/api/v3/{storeId}/products/{productId}/media`&#x20;

### Required access scopes

Your app must have the following **access scopes** to make this request: `update_catalog`

### Path params

All path params are required.

| Param     | Type   | Description          |
| --------- | ------ | -------------------- |
| storeId   | number | Ecwid store ID.      |
| productId | number | Internal product ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Request JSON

A JSON object with the following fields:

| Field        | Type                          | Description                                                                                                            |
| ------------ | ----------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| mainMedia    | object mainMedia              | <p>Link to the main product image or video. <br><br>Empty or <code>null</code> value commits no changes.</p>           |
| galleryMedia | array of objects galleryMedia | <p>List of links to product gallery images and videos.<br><br>If both lists are empty the gallery will be deleted.</p> |

#### mainMedia

| Field    | Type     | Description                                                                                                                                           |
| -------- | -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| videoUrl | imageUrl | <p>Link to the main product video. <br><br>Overrides current main product image/video.<br><br>If missing, the main product video will be deleted.</p> |
| imageUrl | imageUrl | <p>Link to the main product image. <br><br>Overrides current main product image/video.<br><br>If missing, the main product image will be deleted.</p> |

#### galleryMedia

| Field    | Type     | Description                                                                                                                                                                                                                                      |
| -------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| videoUrl | imageUrl | <p>Link to the gallery product video. <br><br><code>galleryMedia</code> array overrides current gallery of product images and videos.<br><br>If <code>galleryMedia</code> is missing, all gallery product images and videos will be deleted.</p> |
| imageUrl | imageUrl | <p>Link to the gallery product image. <br><br><code>galleryMedia</code> array overrides current gallery of product images and videos.<br><br>If <code>galleryMedia</code> is missing, all gallery product images and videos will be deleted.</p> |
