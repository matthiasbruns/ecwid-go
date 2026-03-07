# Dynamic loading for storefront widget

### Delayed widget initialization

Sometimes it is necessary to delay widget initialization while the website page finishes the initialization procedures. This is useful when the website is built dynamically using libraries such as **ReactJs**.

See the example of delayed initialization of different Ecwid widgets below.

```html
<div id="my-store-1003"></div>
<div id="my-categories-1003"></div>
<div id="my-search-1003"></div>
<div class="ec-cart-widget"></div>

<script>
  window.ecwid_script_defer = true;
  
  var script = document.createElement('script');
  script.charset = 'utf-8';
  script.type = 'text/javascript';
  script.src = 'https://app.ecwid.com/script.js?1003';

  document.getElementById('my-store-1003').appendChild(script);

  window._xnext_initialization_scripts = [
      // Storefront widget
      { 
        widgetType: 'ProductBrowser', id: 'my-store-1003', arg: [
          'id=my-store-1003', 'views=grid(1,60)', 'defaultCategoryId=172786255'
        ] 
      },
      // Horizontal categories widget
      { 
        widgetType: 'CategoriesV2', id: 'id=my-categories-1003', arg: [
          'id=my-categories-1003'
        ] 
      },
      // Search widget
      { 
        widgetType: 'SearchWidget', id: 'my-search-1003', arg: [
          'id=my-search-1003'
        ] 
      },
      // 'Buy now' button for product
      { 
        widgetType: 'SingleProduct', id: 'Product-1', arg: [
          'id=Product-1'
        ]
      }
  ];

// Initialize Minicart Widget. A div with class '.ec-cart-widget' must be present on a page  
  
  Ecwid.init();

</script>
```

Replace all `1003` values with your **store ID**.

You can additionally control some settings of the main store widget called **ProductBrowser**:

* Number of products on category pages: `views=grid(1,60)`, where number (60 means that Ecwid will load up to 60 products per page).
* Category opened by default: `defaultCategoryId=172786255`, where you can pass category ID or leave it at `0` (which means the store will open default first page).

### Dynamic embedding of Ecwid storefront widget

Sometimes you need to dynamically create and destroy the storefront widget on the website. For example, on websites that allow switching pages without actually reloading it.

{% hint style="info" %}
Dynamic embedding of the storefront widget is slower than the direct one. Only use it when you need the storefront embedded on the fly.
{% endhint %}

```html
<div id="my-store-1003"></div>
<script>
window.ecwid_script_defer = true;
window.ecwid_dynamic_widgets = true;

   if (typeof Ecwid != 'undefined') Ecwid.destroy(); 
   window._xnext_initialization_scripts = [{
        widgetType: 'ProductBrowser',
        id: 'my-store-1003',
        arg: ["id=productBrowser"]
      }];

  if (!document.getElementById('ecwid-script')) {
      var script = document.createElement('script');
      script.charset = 'utf-8';
      script.type = 'text/javascript';
      script.src = 'https://app.ecwid.com/script.js?1003';
      script.id = 'ecwid-script'
      document.body.appendChild(script);
    } else {
      ecwid_onBodyDone();
    }
</script>
```

Replace all `1003` values with your **store ID**.

You can additionally control some settings of the main store widget called **ProductBrowser**:

* Number of products on category pages: `views=grid(1,60)`, where number (60 means that Ecwid will load up to 60 products per page).
* Category opened by default: `defaultCategoryId=172786255`, where you can pass category ID or leave it at `0` (which means the store will open default first page).

Use `window.ecwid_dynamic_widgets` variable to enable dynamic widget creation in Ecwid. See the example above that shows how to create and destroy Ecwid widget through the JavaScript functions.

{% hint style="info" %}
Dynamic embedding supports only the main store widget called **ProductBrowser**. If you need to embed other widgets dynamically, please, use the code for deferred widget initialization.
{% endhint %}
