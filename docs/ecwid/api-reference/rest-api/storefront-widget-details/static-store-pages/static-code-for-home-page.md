# Static code for home page

<mark style="color:green;">`GET`</mark> \
`https://storefront.ecwid.com/home-page/{storeId}/static-code`&#x20;

### Required access scopes

Request does not require any **access scopes**.

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

All query params are optional.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>lang</td><td>string</td><td><p>Set language for text labels. </p><p></p><p>If not specified, Ecwid will respond with default store language.</p></td></tr><tr><td>baseUrl</td><td>string</td><td>Set base store URL for links in response. <br><br>If not specified, Ecwid will use the main URL from store settings.</td></tr><tr><td>cleanUrls</td><td>boolean</td><td>Set <code>true</code> to force receiving clean URLs – catalog links without hashbang (<code>/#!/</code>). <br><br>If not specified, links will use store URL format.</td></tr><tr><td>tplvar_*</td><td>string</td><td>Pass <code>ec.storefront.*</code> design configs in the request to customize the static page looks. <a data-mention href="https://app.gitbook.com/s/aRJpOy0U8IpbjUfcox4D/store-configuration-settings/design-configs">Design configs</a><br><br>If not specified, API responds with the current store design settings.</td></tr><tr><td>internationalPages</td><td>string</td><td>Set URL to receive store translations. Format: <code>international_pages[en]</code>. HTML is returned in <code>hrefLangHtml</code> field. <a href="https://support.google.com/webmasters/answer/189077?hl=en">Google specification</a></td></tr><tr><td>limit</td><td>number</td><td>Specify the exact fields to receive in response JSON. If not specified, the response JSON will have all available fields for the entity.<br>Example: <code>?responseFields=generalInfo(storeId,storeUrl)</code></td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

| Name                | Type             | Description                                                                                                                                                                                                |
| ------------------- | ---------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| cssFiles            | Array of strings | List of CSS URLs for the page design to display properly                                                                                                                                                   |
| htmlCode            | string           | HTML code for the static page                                                                                                                                                                              |
| metaDescriptionHtml | string           | HTML code for the meta title and description                                                                                                                                                               |
| canonicalUrl        | string           | Canonical URL for this page                                                                                                                                                                                |
| ogTagsHtml          | string           | HTML code for Open Graph tags                                                                                                                                                                              |
| jsonLDHtml          | string           | HTML code for JSON-LD product description                                                                                                                                                                  |
| hrefLangHtml        | string           | `hreflang` HTML tag for translated versions of your website. Returned if `international_pages` request parameter is set. [Google specification](https://support.google.com/webmasters/answer/189077?hl=en) |
| lastUpdated         | number           | UNIX timestamp for when the page was generated                                                                                                                                                             |
