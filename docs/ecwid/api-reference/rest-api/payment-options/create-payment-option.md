# Create payment option

<mark style="color:blue;">`POST`</mark> `https://app.ecwid.com/api/v3/{storeId}/profile/paymentOptions`&#x20;

### Required access scopes

Your app must have the following **access scopes** to make this request: `update_store_profile` , `add_payment_method`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Request JSON

A JSON object with the following fields:

<table><thead><tr><th width="245">Field</th><th width="217">Type</th><th>Description</th></tr></thead><tbody><tr><td>enabled</td><td>boolean</td><td><code>true</code> if payment method is enabled and visible at the checkout, <code>false</code> otherwise.</td></tr><tr><td>configured</td><td>boolean</td><td><p>Setup status of the payment method. <br></p><p>If <code>false</code>, highlights payment methods with the red border.</p></td></tr><tr><td>checkoutTitle</td><td>string</td><td>Name visible to customers at the checkout above the payment option name.</td></tr><tr><td>checkoutTitleTranslated</td><td>object <a href="#translations">translations</a></td><td>Available translations for payment method title.</td></tr><tr><td>checkoutDescription</td><td>string</td><td>Payment method description visible under the payment option name at the checkout.</td></tr><tr><td>paymentProcessorId</td><td>string</td><td>Payment processor ID in Ecwid</td></tr><tr><td>paymentProcessorTitle</td><td>string</td><td>Payment processor title. The same as <code>paymentModule</code> in order details in REST API</td></tr><tr><td>orderBy</td><td>number</td><td>Payment method position at checkout and in Ecwid Control Panel. The smaller the number, the higher the position is</td></tr><tr><td>appClientId</td><td>string</td><td>client_id value of payment application. <code>""</code> if not an application</td></tr><tr><td>paymentSurcharges</td><td>object <a href="#paymentsurcharges">paymentSurcharges</a></td><td>Payment method fee added to the order as a set amount or as a percentage of the order total</td></tr><tr><td>instructionsForCustomer</td><td>object <a href="#instructionsforcustomer">instructionsForCustomer</a></td><td>Payment instructions visible to customers at the checkout.</td></tr><tr><td>shippingSettings</td><td>object <a href="#shippingsettings">shippingSettings</a></td><td>Limit payment option by the list of shipping methods selected at the checkout.</td></tr></tbody></table>

#### paymentSurcharges

<table><thead><tr><th width="208">Field</th><th width="180">Type</th><th>Description</th></tr></thead><tbody><tr><td>type</td><td>string</td><td><p>Surcharge type that defines how it applies to the payment.<br><br>One of: </p><p><code>ABSOLUTE</code></p><p><code>PERCENT</code></p></td></tr><tr><td>value</td><td>number</td><td>Surcharge value.</td></tr></tbody></table>

#### instructionsForCustomer

<table><thead><tr><th width="213">Field</th><th width="174">Type</th><th>Description</th></tr></thead><tbody><tr><td>instructionsTitle</td><td>string</td><td>Name visible above the payment instructions block at the checkout.</td></tr><tr><td>instructions</td><td>string</td><td><p>Content inside the payment instructions block. </p><p></p><p>Supports HTML tags.</p></td></tr><tr><td>instructionsTranslated</td><td>object <a href="#translations">translations</a></td><td>Available translations for payment instructions content.<br><br>Supports HTML tags.</td></tr></tbody></table>

#### shippingSettings

<table><thead><tr><th width="236">Field</th><th width="158">Type</th><th>Description</th></tr></thead><tbody><tr><td>enabledShippingMethods</td><td>array of strings</td><td>List of shipping methods (internal shipping method IDs).<br><br>If specified, the payment option is only available when customers select the specified shipping method at the checkout.<br><br>It allows, for example, disabling online payments for pickup orders (leaving it available for deliveries). </td></tr></tbody></table>

#### translations

Object with text field translations in the `"lang": "text"` format, where the `"lang"` is an ISO 639-1 language code. For example:

```
{
    "en": "Sample text",
    "nl": "Voorbeeldtekst"
}
```

Translations are available for all active store languages. Only the default language translations are returned if no other translations are provided for the field. Find active store languages with <mark style="color:green;">`GET`</mark> `/profile` request > `languages` > `enabledLanguages`.

### Response JSON

A JSON object with the following fields:

| Field | Type   | Description                                |
| ----- | ------ | ------------------------------------------ |
| id    | number | Internal ID of the created payment method. |
