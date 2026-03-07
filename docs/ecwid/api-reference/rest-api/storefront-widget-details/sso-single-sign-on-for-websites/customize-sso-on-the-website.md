# Customize SSO on the website

SSO workflow can be tailored to specifics of merchant's website. Check out the available customizations below.

{% hint style="info" %}
These customizations will work only when Ecwid is in SSO mode. That is, when the global variable ***ecwid\_sso\_profile*** is also defined at least as an empty string "".
{% endhint %}

### Add log in link to the store in SSO mode

By default, when SSO is used on a web page, Ecwid hides Sign in and Sign out links in the storefront assuming that the user logs in by means of your site login form.

Alternatively, you can customize the Sign in / Sign out links in a store and link them to the site login form to make login process more convenient for your customers.

```javascript
Ecwid.OnAPILoaded.add(function() {
    Ecwid.setSignInUrls({
        signInUrl: 'http://my.site.com/signin',
        signOutUrl: 'http://my.site.com/signout' // signOutUrl is optional
    });
});
```

This can be done by means of `Ecwid.setSignInUrls()` API extension, which accepts two parameters:

| Field      | Type   | Description                                                                                                                                                                         |
| ---------- | ------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| signInUrl  | string | <p>URL where your customer will be redirected to log in to your site<br><br><strong>Required</strong></p>                                                                           |
| signOutUrl | string | URL where your customer will be redirected, when they click the 'Log out' link in the store. This parameter is optional. If not specified, Ecwid will not show the ’sign out’ link. |

When you use `setSignInUrls()` extension, the customer will be redirected to the specified URL in the following situations:

* They click the 'Sign in' link in the product browser
* They open a secured page inside Product Browser (e.g. ’My account’ page)
* They go to checkout and your store settings require a user to be registered to place an order

### Custom Sign in / Sign out handlers

Ecwid SSO API provides a tool for advanced customization of the login functionality in Single Sign-on mode – `Ecwid.setSignInProvider()` method.

```javascript
Ecwid.OnAPILoaded.add(function() {
  Ecwid.setSignInProvider({
    addSignInLinkToPB: function() { return true; },
    signIn: function () { 
      alert('User is logging in!');
      // do something
    },

    canSignOut: function() { return true; },
    signOut: function () { 
      alert('User is signing out');
    }
  });
});
```

| Field             | Type                             | Description                                                                                                                                                                                                                                                                                   |
| ----------------- | -------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| addSignInLinkToPB | boolean                          | <p>Set it to return <code>true</code> if you need to show 'Sign in' link inside your store.<br><br><strong>Required</strong></p>                                                                                                                                                              |
| signIn            | function                         | <p>Specify function to be called when Ecwid is trying to authorize the customer, when either the 'Sign in' link is clicked or a secured page is opened inside Product Browser. This function does not accept any arguments nor should return any result.<br><br><strong>Required</strong></p> |
| canSignOut        | function returning true or false | <p>Set it to return <code>true</code> if you need to show 'Sign out' link inside your store.<br><br><strong>Required</strong></p>                                                                                                                                                             |
| signOut           | function                         | <p>This function is called when <code>canSignOut</code> returns true and the customer clicks the 'Sign out' link inside the store widget.<br><br><strong>Required</strong></p>                                                                                                                |

<br>
