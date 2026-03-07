# Static store pages

**Static store pages** allow generating a lightweight HTML snapshot of the home page or specific category/product pages. You can display that snapshot to users almost immediately and load a full-functioning Ecwid store in the background. This way you can dramatically increase the loading speed of the Ecwid storefront on your website.

If your website is based on Wix or WordPress site builders, use our official integrations: [WordPress plugin](https://support.ecwid.com/hc/en-us/articles/207101259-Adding-your-Ecwid-store-to-WordPress-site), [Wix app](https://support.ecwid.com/hc/en-us/articles/115005874885-Adding-your-Ecwid-store-to-Wix-site). These integrations have static store pages enabled out of the box.

And if you build a storefront on another CMS or a custom website, use our **Storefront SDK** and **Static code endpoints** to set up static pages for your website.

### How to use static pages

To start working with static pages, you need access to the HTML and JS code of your website. First, you'll need to receive the static page code with REST API request and display it on the website, then load a dynamic full-functioning storefront and switch to it.

#### Enable Storefront SDK on your website

Storefront SDK contains several pre-built methods for managing and switching static pages on the storefront with simple JavaScript calls. To enable it on your website, add a JavaScript SDK file to the website `<head>`.

Latest version: `https://djqizrxa6f10j.cloudfront.net/ec-sdk/storefront/2.2.0/storefront.min.js`.

Code example:

```html
<head>
    <script type='text/javascript' src='https://djqizrxa6f10j.cloudfront.net/ec-sdk/storefront/2.2.0/storefront.min.js'></script>
</head>
```

#### Receive static code for the page

When a user opens your home, category, or product page, request a static HTML code for that page through REST API. While GET requests for static code don't require access tokens, we still recommend using backend calls to avoid CORS errors. For example, using a separate PHP file for REST API calls.

Find endpoint descriptions by the following links:

* [Home page](https://docs.ecwid.com/api-reference/rest-api/storefront-widget-details/static-store-pages/static-code-for-home-page) endpoint.
* [Category page](https://docs.ecwid.com/api-reference/rest-api/storefront-widget-details/static-store-pages/static-code-for-category-page) endpoint.
* [Product page](https://docs.ecwid.com/api-reference/rest-api/storefront-widget-details/static-store-pages/static-code-for-product-page) endpoint.

Code examples:

```php
<?php

// Send cURL request

$curl = curl_init();

curl_setopt_array($curl, array(
  CURLOPT_URL => 'https://storefront.ecwid.com/home-page/STOREID/static-code?limit=60',
  CURLOPT_RETURNTRANSFER => true,
  CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
  CURLOPT_CUSTOMREQUEST => 'GET',
));

$response = curl_exec($curl);

curl_close($curl);

// Echo the response

echo $response;
?>
```

```javascript
// Import jQuery .js file to the website to enable Ajax requests

$.ajax({
    
    // Call PHP file that requests static code for the home page

    url: 'call.php',

    // Parse response on successful call

    success: function (data) {
        var cssFiles = JSON.parse(data).cssFiles;
        var htmlCode = JSON.parse(data).htmlCode;
        var jsCode = JSON.parse(data).jsCode;

        // Add received JavaScript store code to <head>

        var jsCodeAppended = document.createElement("script");
        jsCodeAppended.innerHTML = jsCode;
        jsCodeAppended.type = "text/javascript";
        document.head.appendChild(jsCodeAppended);

        // Add received CSS stylesheet file to <head>

        var cssFilesAppended = document.createElement("link");
        cssFilesAppended.rel = 'stylesheet';
        cssFilesAppended.href = cssFiles;
        document.head.appendChild(cssFilesAppended);

        // Add received HTML store code to a static code contained in <body>

        $(document).ready(function () {
            $("#static-ec-store-container").html(htmlCode);
        });
    }
});
```

In the example above, we call a PHP file that receives static code from REST API. On success, the script parses the `$response` and adds JS and CSS codes to the `<head>` and HTML code to the `<div id="dynamic-ec-store-container">` inside website `<body>`.

As a result, the static page is displayed on the website.

#### Start loading dynamic code

With the static code loaded and displayed to website visitors, you can start loading the dynamic storefront. Load the code in a hidden `<div>` on the same page. As a result, there will be two `<div>`s with the store page code: static storefront (visible) and dynamic storefront (hidden).

There are two ways of loading dynamic code: **default loading** and **lazy loading**. We recommend using lazy loading as it should give better website performance scores.

Default loading code example:

```html
<script data-cfasync="false" data-no-optimize="1" type="text/javascript">
    ec.storefront.staticPages.staticStorefrontEnabled = true;
    ec.storefront.staticPages.staticContainerID = 'static-ec-store-container';
    ec.storefront.staticPages.dynamicContainerID = 'dynamic-ec-store-container';
    ec.storefront.staticPages.autoSwitchStaticToDynamicWhenReady = true;
</script>

<div id="dynamic-ec-store-container">
    <div>
        <script data-cfasync="false" type="text/javascript"
            src="https://app.ecwid.com/script.js?<store_id>&data_platform=code&data_date=2021-12-29"
            charset="utf-8"></script>
        <script
            type="text/javascript"> xProductBrowser("categoriesPerRow=3", "views=grid(20,3) list(60) table(60)", "categoryView=grid", "searchView=list");</script>
    </div>
</div>
```

Lazy loading code example:

```html
<script data-cfasync="false" data-no-optimize="1" type="text/javascript">
    ec.storefront.staticPages.staticStorefrontEnabled = true;
    ec.storefront.staticPages.staticContainerID = 'static-ec-store-container';
    ec.storefront.staticPages.dynamicContainerID = 'dynamic-ec-store-container';
    ec.storefront.staticPages.autoSwitchStaticToDynamicWhenReady = true;
    ec.storefront.staticPages.lazyLoading = {
        scriptJsLink: 'https://app.ecwid.com/script.js?STOREID&data_platform=code',
        xProductBrowserArguments: ["categoriesPerRow=3", "views=grid(20,3) list(60) table(60)", "categoryView=grid", "searchView=list", "id=dynamic-ec-store"]
    }
</script>

<div id="dynamic-ec-store-container">
    <div id="dynamic-ec-store">

    </div>
</div>

<script>
    window.ec.storefront.staticPages.forceDynamicLoadingIfRequired();
</script>
```

In both cases, static page code gets automatically replaced with a dynamic storefront when it's ready. If you want to call the switch manually, replace the `ec.storefront.staticPages.autoSwitchStaticToDynamicWhenReady = true;` config with `false` and call `StaticPageLoader.switchToDynamicMode();` when you are ready to switch.

Check out [home page example](https://d35z3p2poghz10.cloudfront.net/apps/ecwid-static-pages/examples/static-page-demo.htm) and its source code to get to know the static code feature better.

### Summary

Let's combine the examples above into a project with two files: a PHP file making REST API requests and an HMTL page code. Replace `STOREID` with your store ID in both files to test it locally:

```php
<?php

$curl = curl_init();

curl_setopt_array($curl, array(
  CURLOPT_URL => 'https://storefront.ecwid.com/home-page/STOREID/static-code?limit=60',
  CURLOPT_RETURNTRANSFER => true,
  CURLOPT_ENCODING => '',
  CURLOPT_MAXREDIRS => 10,
  CURLOPT_TIMEOUT => 0,
  CURLOPT_FOLLOWLOCATION => true,
  CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
  CURLOPT_CUSTOMREQUEST => 'GET',
));

$response = curl_exec($curl);

curl_close($curl);
echo $response;
?>
```

```html
<!DOCTYPE html>
<html lang="en-EN">

<head>
    <script type='text/javascript' src='https://djqizrxa6f10j.cloudfront.net/ec-sdk/storefront/2.2.0/storefront.min.js'></script>
    <script type='text/javascript' src="https://ajax.googleapis.com/ajax/libs/jquery/3.7.1/jquery.min.js"></script>
    <script>
        $.ajax({
            url: 'test_.php',
            success: function (data) {
                var cssFiles = JSON.parse(data).cssFiles;
                var htmlCode = JSON.parse(data).htmlCode;
                var jsCode = JSON.parse(data).jsCode;

                var jsCodeAppended = document.createElement("script");
                jsCodeAppended.innerHTML = jsCode;
                jsCodeAppended.type = "text/javascript";
                document.head.appendChild(jsCodeAppended);

                var cssFilesAppended = document.createElement("link");
                cssFilesAppended.rel = 'stylesheet';
                cssFilesAppended.href = cssFiles;
                document.head.appendChild(cssFilesAppended);

                $(document).ready(function () {
                    $("#static-ec-store-container").html(htmlCode);
                });
            }
        });
    </script>
</head>

<body>
    <script data-cfasync="false" data-no-optimize="1" type="text/javascript">
        ec.storefront.staticPages.staticStorefrontEnabled = true;
        ec.storefront.staticPages.staticContainerID = 'static-ec-store-container';
        ec.storefront.staticPages.dynamicContainerID = 'dynamic-ec-store-container';
        ec.storefront.staticPages.autoSwitchStaticToDynamicWhenReady = true;
        ec.storefront.staticPages.lazyLoading = {
            scriptJsLink: 'https://app.ecwid.com/script.js?STOREID&data_platform=code',
            xProductBrowserArguments: ["categoriesPerRow=3", "views=grid(20,3) list(60) table(60)", "categoryView=grid", "searchView=list", "id=dynamic-ec-store"]
        }
    </script>

    <div id="dynamic-ec-store-container">
        <div id="dynamic-ec-store"></div>
    </div>

    <div id="static-ec-store-container">

    </div>

    <script>
        window.ec.storefront.staticPages.forceDynamicLoadingIfRequired();
    </script>
</body>

</html>
```

This example loads static code for the store home page, then loads dynamic code for the same page and switches to the dynamic code once it's ready.
