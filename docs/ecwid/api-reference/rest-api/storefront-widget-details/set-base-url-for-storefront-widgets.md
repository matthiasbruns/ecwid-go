# Set base URL for storefront widgets

Ecwid storefront widgets can be added to several of your website pages. For example:

`example.com/store` - Main store page.

`example.com/sale` - Secondary page with one category (products currently on sale).

`example.com/gift-card` - Secondary page with one product widget (gift card).

Such pages won't be "linked" by default. They'll have separate carts and the `/gift-card` page will have a popup checkout. However, there is an easy way of connecting all three pages to one checkout by adding a line of code to all of the secondary pages (`example.com/sale` and `example.com/gift-card`).

The script must be added right after the widget integration code on your website:

```html
// 
// Ecwid integration code
// 
<script>
    Ecwid.setStorefrontBaseUrl('example.com/store');
</script>
```

This way, all 3 pages will work through the checkout on the `/store` page. When customers click on the cart icons/buttons on secondary pages, they'll be redirected to the checkout on the main store page.
