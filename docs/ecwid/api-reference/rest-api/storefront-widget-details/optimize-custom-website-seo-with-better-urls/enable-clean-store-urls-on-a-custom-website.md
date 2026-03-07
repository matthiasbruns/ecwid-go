# Enable Clean Store URLs on a custom website

**Clean Store URLs** feature removes the hashbang (`/#!/`) and IDs (`/p/123456`) parts from product and category page URLs. \
\
For example, instead of URLs looking like: \
`https://example.com/store/#!/product-name/p/123456`,\
\
They will look like: \
`https://example.com/store/product-name`

Read more about [**URL features for custom websites**](https://docs.ecwid.com/api-reference/rest-api/storefront-widget-details/optimize-custom-website-seo-with-better-urls)

### Requirements for Enabling Clean Store URLs

To enable Clean Store URLs on a custom website, you must have:

* Access to its server rewrite rules
* HTML code of the store page.

The rewrite rules will depend on the server software your website uses:

* **Apache**: Use the `.htaccess` file. Refer to the guide on the [**Apache website**](https://httpd.apache.org/docs/2.4/howto/htaccess.html).
* **Nginx**: Modify the server configuration file. Refer to the guide on the [**Nginx website**](https://blog.nginx.org/blog/creating-nginx-rewrite-rules).
* **Microsoft IIS**: Use the URL rewrite module. See the guide on the [**Microsoft Learn portal**](https://learn.microsoft.com/en-us/iis/extensions/url-rewrite-module/creating-rewrite-rules-for-the-url-rewrite-module).

To simplify the setup, you can use our template, which includes the Apache rewrite rules, and the HTML code to integrate the Ecwid store into your page. The template is specifically designed for **Apache-based** web servers.

The template works with both **.htaccess** and **shop.html** files running from the root folder on the web server:

```
- [web server]
  - .htaccess
  - shop.html
```

Depending on the page the Ecwid store works from, you need to choose one of the two template codes:

* Ecwid store on the home page. \
  Example product URL: `https://example.com/product-name`
* Ecwid store on one of the website subpages. \
  Example product URL: `https://example.com/shop/product-name`

### Step 1. Add server rewrite rules

To create rewrite rules, you need to specify path to the HTML file with Ecwid store (`shop.html`) on the web server and the website page Ecwid works from (`RewriteBase`) in the **.htaccess** file template.

Example code (Apache-based servers):

{% tabs %}
{% tab title="Ecwid store works from the home page" %}

```apacheconf
<IfModule mod_rewrite.c>
    RewriteEngine On
    RewriteBase /
    RewriteRule ^.+$ shop.html
    RewriteRule ^cart$ shop.html
    RewriteRule ^search.*$ shop.html
    RewriteRule ^checkout/.+$ shop.html
    RewriteRule ^account/.+$ shop.html
    RewriteRule ^pages/.+$ shop.html
    RewriteRule ^signIn.*$ shop.html
    RewriteRule ^signOut.*$ shop.html
    RewriteRule ^resetPassword.*$ shop.html
    RewriteRule ^checkoutAB.*$ shop.html
    RewriteRule ^downloadError.*$ shop.html
    RewriteRule ^checkoutResult.*$ shop.html
    RewriteRule ^checkoutWait.*$ shop.html
    RewriteRule ^orderFailure.*$ shop.html
    RewriteRule ^checkoutCC.*$ shop.html
    RewriteRule ^checkoutEC.*$ shop.html
    RewriteRule ^checkoutAC.*$ shop.html
    RewriteRule ^FBAutofillCheckout.*$ shop.html
    RewriteRule ^pay.*$ shop.html
    RewriteRule ^repeat-order.*$ shop.html
    RewriteRule ^subscribe.*$ shop.html
    RewriteRule ^unsubscribe.*$ shop.html
</IfModule>
```

{% endtab %}

{% tab title="Ecwid store works from the subpage" %}

```apacheconf
<IfModule mod_rewrite.c>
    RewriteEngine On
    RewriteBase /shop/
    RewriteRule ^.+$ shop.html
</IfModule>
```

{% endtab %}
{% endtabs %}

### Step 2. Add the Ecwid integration HTML code

To create the HTML integration code that will add Ecwid store and enable Clean URLs, you need to replace all three placeholders with your store ID (`STOREID`) and the website subpage Ecwid works from (`window.ec.config.baseUrl`) in the **HTML** file template:

{% tabs %}
{% tab title="Ecwid store works from the home page" %}

```html
<html>

<head></head>

<body>
    <div id="my-store-STOREID"></div>
    <script>
        window.ec = window.ec || {};
        window.ec.config = window.ec.config || {};
        window.ec.config.storefrontUrls = window.ec.config.storefrontUrls || {};

        window.ec.config.storefrontUrls.cleanUrls = true;
        window.ec.config.storefrontUrls.slugsWithoutIds = true;
    </script>

    <script data-cfasync="false" type="text/javascript"
        src="https://app.ecwid.com/script.js?STOREID&data_platform=code" charset="utf-8"></script>
    <script type="text/javascript">
        xProductBrowser("categoriesPerRow=3", "views=grid(20,3) list(60) table(60)", "categoryView=grid", "searchView=list", "id=my-store-STOREID");
    </script>
</body>

</html>
```

{% endtab %}

{% tab title="Ecwid store works from the subpage" %}

```html
<html>

<head></head>

<body>
    <div id="my-store-STOREID"></div>

    <script>
        window.ec = window.ec || {};
        window.ec.config = window.ec.config || {};
        window.ec.config.storefrontUrls = window.ec.config.storefrontUrls || {};

        window.ec.config.storefrontUrls.cleanUrls = true;
        window.ec.config.storefrontUrls.slugsWithoutIds = true;

        window.ec.config.baseUrl = '/shop/';
    </script>

    <script data-cfasync="false" type="text/javascript"
        src="https://app.ecwid.com/script.js?STOREID&data_platform=code" charset="utf-8"></script>
    <script type="text/javascript">
        xProductBrowser("categoriesPerRow=3", "views=grid(20,3) list(60) table(60)", "categoryView=grid", "searchView=list", "id=my-store-STOREID");
    </script>
</body>

</html>
```

{% endtab %}
{% endtabs %}

Once you've added both server rewrite rules and HTML integration code, URLs added by the Ecwid store on your custom website will automatically be converted to the **Clean Store URLs** format.
