# SSO (Single-Sign-On) for websites

Single Sign-On (SSO) allows customer to login to Ecwid store if they are already logged in to your website. If a merchant already has a customer base on their website, customers may find it rather inconvenient that they have to log into the website and the store separately.

Ecwid’s Single Sign-On feature allows those customers to sign into a merchant’s website and use the entire Ecwid store without having to log in again.

#### Access scopes

Requires the following access scopes: `customize_storefront` and `create_customers`

### How it works

There are two situations that can happen: customer is logged into your website or not:

#### Logged in user

When user is logged into your website, it should pass an authentication information to Ecwid on a page with Ecwid widgets: what user to log in and any other additional info.

If Ecwid detects that information and it is valid – the user will be automatically logged into the account your website provided. If that user doesn't exist, Ecwid will create account for them.

After the login, Ecwid will hide any sign in links and behave as if the user is logged into their customer account in the Ecwid storefront.

#### Logged out user

When user is not logged into your website, Ecwid will show sign in links, which can lead to custom URLs – login functionality of your website. Once the login is done, everything works the same as described for logged in users.
