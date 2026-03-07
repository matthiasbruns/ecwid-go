# Get page slug and static code

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/storefront-widget-pages`&#x20;

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_store_profile` , `read_storefront_widget_pages`&#x20;

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

All query params are optional.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>slug</td><td>string</td><td>Specify page slug for search.</td></tr><tr><td>storeRootPage</td><td>boolean</td><td>Set <code>true</code> to receive root page slug. If <code>true</code>, <code>slug</code> field is ignored and could be omited from the request.</td></tr><tr><td>getStaticContent</td><td>boolean</td><td>Set <code>true</code> to receive static HTML code for found page.</td></tr><tr><td>lang</td><td>string</td><td><p>Set language for text labels. </p><p></p><p>If not specified, Ecwid will respond with default store language.</p></td></tr><tr><td>baseUrl</td><td>string</td><td>Set base store URL for links in response. <br><br>If not specified, Ecwid will use the main URL from store settings.</td></tr><tr><td>cleanUrls</td><td>boolean</td><td>Set <code>true</code> to force receiving clean URLs – catalog links without hashbang (<code>/#!/</code>). <br><br>If not specified, links will use store URL format.</td></tr><tr><td>slugsWithoutIds</td><td>boolean</td><td>Set <code>false</code> to force receiving product and category URLs with their IDs in static page code.<br><br>Requires <code>getStaticContent=true</code>.<br><br>If <code>true</code> or not specified, URLs in the static page code will return without IDs.</td></tr><tr><td>tplvar_*</td><td>string</td><td>Pass <code>ec.storefront.*</code> appearance option in the request to customize static page design. <br><br>If not specified, the response will use current design settings.</td></tr><tr><td>internationalPages</td><td>string</td><td>Set URL to receive store translations. Format: <code>international_pages[en]</code>. HTML is returned in <code>hrefLangHtml</code> field. <a href="https://support.google.com/webmasters/answer/189077?hl=en">Google specification</a></td></tr><tr><td>limit</td><td>number</td><td>Specify the exact fields to receive in response JSON. If not specified, the response JSON will have all available fields for the entity.<br>Example: <code>?responseFields=generalInfo(storeId,storeUrl)</code></td></tr></tbody></table>
