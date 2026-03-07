# Search checkout extra fields

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/profile/extrafields`&#x20;

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_store_profile`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

<table><thead><tr><th>Field</th><th width="252">Type</th><th>Description</th></tr></thead><tbody><tr><td>items</td><td>array of objects <a href="#items">items</a></td><td>Detailed information about checkout extra fields saved in the store settings.</td></tr></tbody></table>

#### items

<table><thead><tr><th width="259">Field</th><th width="139">Type</th><th>Description</th></tr></thead><tbody><tr><td>key</td><td>string</td><td>Internal ID of the checkout extra field. <br><br>Extra fields created through Ecwid admin or REST API get their key automatically, for example, <code>ti3hyrr</code></td></tr><tr><td>title</td><td>string</td><td>Title visible at the checkout above the extra field. Also appears in Ecwid admin and customer notifications (if added).</td></tr><tr><td>titleTranslated</td><td>object <a href="#translated">translated</a></td><td>Available translations for the extra field title. </td></tr><tr><td>type</td><td>string</td><td><p>Extra field type that defines its functionality. </p><p><br>One of:<br><code>text</code> - Single-line text input. Supports placeholders and pre-defined values (default).<br><code>textarea</code> - Multiple-line text input. Supports placeholders and pre-defined values.<br><code>select</code> - Drop-down list, where customers can choose only one value. Requires <code>options</code> array.<br><code>radio_buttons</code> - A group of radio buttons, where customers can choose only one value. Requires <code>options</code> array.<br><code>checkbox</code> - A group of checkboxes, where customers can choose multiple values. Requires <code>options</code> array.<br><code>toggle_button_group</code> - A group of buttons, where customers can choose only one value. Requires <code>options</code> array.<br><code>datetime</code> - Customizable date and time picker in the form of a calendar widget on the checkout.<br><code>empty</code> - Non-editable text visible at the checkout and unavailable in customer notifications and Ecwid admin.</p></td></tr><tr><td>textPlaceholder</td><td>string</td><td><p>Placeholder text for <code>text</code> and <code>textarea</code> field types. Does not affect the actual extra field value.</p><p><br>If not specified, the extra field won't have a placeholder.</p></td></tr><tr><td>textPlaceholderTranslated</td><td>object <a href="#translated">translated</a></td><td>Available translations for the text placeholder.</td></tr><tr><td>available</td><td>boolean</td><td><p>Set <code>false</code> to disable the extra field.</p><p><br>If not specified, the extra field is enabled.</p></td></tr><tr><td>required</td><td>boolean</td><td><p>Set <code>true</code>, to make the extra field required at the checkout. This way, customers won't be able to continue checkout until they type/select a value in the field. </p><p></p><p>If not specified, the extra field is not required at the checkout.</p></td></tr><tr><td>checkoutDisplaySection</td><td>string</td><td><p>Defines at which checkout step extra field shows to customers.</p><p><br>One of:<br><code>email</code> - First checkout step where customers enter their email and apply discounts.<br><code>shipping_address</code> - Second checkout step where customers enter their shipping address.<br><code>pickup_details</code> - Second checkout step where customers choose when and where they want to pick up an order.<br><code>shipping_methods</code> - Third checkout step where customers choose a shipping method.<br><code>pickup_methods</code> - Third checkout step where customers choose a pickup method.<br><code>payment_details</code> - Final checkout step where customers choose payment method and proceed to the payment.<br></p><p>If not specified, the extra field won't appear at the checkout, allowing you to store hidden order data or apply hidden extra charges to the order.</p></td></tr><tr><td>orderDetailsDisplaySection</td><td>string</td><td><p>Defines at which section the extra field shows in customer notifications and Ecwid admin. </p><p><br>One of:<br><code>shipping_info</code> - Shipping/pickup details.<br><code>billing_info</code> - Billing details.<br><code>customer_info</code> - Customer details.<br><code>order_comments</code> - Order comments left by a customer.<br><code>hidden</code> - Extra field is hidden from customer notifications and Ecwid admin.</p><p><br>If not specified, the extra field appears in the <code>order_comments</code> section.</p></td></tr><tr><td>showInInvoice</td><td>boolean</td><td><p>Set <code>true</code> to display the extra field in order tax invoices generated by Ecwid. </p><p></p><p>If not specified, the extra field will not be displayed in order tax invoices.</p></td></tr><tr><td>showInNotifications</td><td>boolean</td><td><p>Set <code>true</code> to display the extra field in customer emails generated by Ecwid. If <code>true</code>, requires <code>orderDetailsDisplaySection</code>.</p><p></p><p>If not specified, the extra field will not be displayed in order tax invoices.</p></td></tr><tr><td>value</td><td>string</td><td><p>Default value for the extra field. Any changes by customers or your scripts on the storefront override it.<br></p><p>If not specified, the extra field won't have any default value.</p></td></tr><tr><td>valueTranslated</td><td>object <a href="#translated">translated</a></td><td>Available translations for the extra field value.</td></tr><tr><td>subtitle</td><td>string</td><td>Subtitle visible under the extra field name at the checkout.<br><br>If not specified, the extra field won't have a subtitle.</td></tr><tr><td>subtitleTranslated</td><td>object <a href="#translated">translated</a></td><td>Available translations for the extra field subtitle.</td></tr><tr><td>tip</td><td>string</td><td><p>Extra field tip for <code>text</code> and <code>textarea</code> field types visible inside the input at the checkout.<br></p><p>If not specified, the extra field won't have a tip.</p></td></tr><tr><td>tipTranslated</td><td>object <a href="#translated">translated</a></td><td>Available translations for the extra field tip.</td></tr><tr><td>showForPaymentMethodIds</td><td>array of strings</td><td><p>Make extra field available to only some payment methods by setting a list of payment method IDs. <br><br>For example, you can collect additional data only for the specific online payment method by adding it alone to the list: <code>"showForPaymentMethodIds": ["4959-2345934622523"]</code>.</p><p></p><p>If not specified, the extra field is not limited by any payment methods.</p></td></tr><tr><td>showForShippingMethodIds</td><td>array of strings</td><td><p>Make extra field available to only some shipping methods by setting a list of shipping method IDs. <br><br>For example, you can collect additional data only for the store pickup method by adding it alone to the list: <code>"showForShippingMethodIds": ["4959-2345934622523"]</code>.</p><p></p><p>If not specified, the extra field is not limited by any shipping methods.</p></td></tr><tr><td>showForCountry</td><td>array of strings</td><td><p>Make extra field available only for some countries by setting a list of country codes. Ecwid will automatically check customer's address details to display the field.<br><br>For example, <code>"showForCountry": ["BE", "US"]</code>.<br></p><p>If not specified, the extra field is not limited by the customer's address details.</p></td></tr><tr><td>options</td><td>array of objects <a href="#options">options</a></td><td><p>Defines a list of available options for <code>select</code>, <code>checkbox</code>, <code>radio_buttons</code>, and<code>toggle_button_group</code> field types.<br></p><p>If not specified, the field type is transformed to <code>text</code>.</p></td></tr><tr><td>datepickerOptions</td><td>object <a href="#datepickeroptions">datepickerOptions</a></td><td><p>Defines settings and conditions for the datepicker at the checkout. For example, limited working hours or disable weekends for the calendar.</p><p><br>Affects only <code>datetime</code> extra field type.</p></td></tr><tr><td>overrides</td><td>object <a href="#translated">translated</a></td><td>Defines overrides for extra fields. This setting allows setting up logic and conditional changes for extra fields on the checkout without actual JavaScript code.</td></tr><tr><td>orderBy</td><td>number</td><td>Sorting order for checkout extra fields on the order details page in Ecwid admin and in customer emails.<br><br>Does not affect extra field's order at the checkout.</td></tr><tr><td>shownOnOrderDetails</td><td>boolean</td><td>Defines if the extra field should be visible on the order details page in Ecwid admin.</td></tr><tr><td>saveToCustomerProfile</td><td>boolean</td><td>Defines if the extra field should also be saved to customer extra fields.</td></tr><tr><td>cpField</td><td>boolean</td><td>Defines if the extra field was created from Ecwid admin (<code>true</code>).</td></tr></tbody></table>

#### options

<table><thead><tr><th>Field</th><th width="192">Type</th><th>Description</th></tr></thead><tbody><tr><td>title</td><td>string</td><td>Option name visible at the checkout.</td></tr><tr><td>titleTranslated</td><td>object <a href="#translated">translated</a></td><td>Available translations for the option name.</td></tr><tr><td>subtitle</td><td>string</td><td>Subtitle visble under the option name at the checkout.</td></tr><tr><td>subtitleTranslated</td><td>oobject <a href="#translated">translated</a></td><td>Available translations for the option subtitle.</td></tr><tr><td>surcharge</td><td>number</td><td>Tip/surcharge value added to the order total cost.</td></tr><tr><td>surchargeType</td><td>string</td><td>Defines if a surcharge applies a fixed sum or a percentage increase to the cart total.<br><br>One of:<br><code>ABSOLUTE</code> - Adds a fixed sum.<br><code>PERCENT</code> - Adds percentage increase.</td></tr><tr><td>surchargeTaxable</td><td>boolean</td><td>Defines if surcharge/tip is taxable.</td></tr><tr><td>showZeroSurchargeInTotal</td><td>boolean</td><td><p>Defines if zero-sum surcharge shows in order total breakdown. </p><p><br>Set <code>false</code> to hide zero-sum surcharges from customers at the checkout.</p></td></tr><tr><td>surchargeShortName</td><td>object <a href="#surchargeshortname">surchargeShortName</a></td><td>Defines how the surcharge shows in the "Shopping cart" block where the order total is calculated.</td></tr></tbody></table>

#### surchargeShortName

<table><thead><tr><th width="263">Field</th><th width="112">Type</th><th>Description</th></tr></thead><tbody><tr><td>name</td><td>string</td><td>Name of the surcharge for the "Shopping cart" block.</td></tr><tr><td>nameTranslated</td><td>object <a href="#translated">translated</a></td><td>Available translations for the surcharge short name.</td></tr><tr><td>showSurchargePercentValue</td><td>boolean</td><td>Defines if surcharge percentage value shows next to <code>name</code>. Default: <code>true</code>. <br><br>Works only with <code>PERCENT</code> value in <code>surchargeType</code>.</td></tr></tbody></table>

#### datepickerOptions

| Field                     | Type                                                           | Description                                                                                                                                               |
| ------------------------- | -------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| minDate                   | date                                                           | The earliest selectable date. Users can choose dates starting from this minimum date. Use a valid Date object.                                            |
| maxDate                   | date                                                           | The latest selectable date. Users cannot choose dates beyond this maximum date. Use a valid Date object.                                                  |
| showtime                  | boolean                                                        | Defines if datepicker allows users to select the time. If `true`, time selection is enabled.                                                              |
| incrementMinuteBy         | number                                                         | Increments for datepicker time frames (in minutes). For example, `30` means that customers can only choose from `17:00, 17:30, 18:00...` on the checkout. |
| limitAvailableHoursWeekly | object [limitAvailableHoursWeekly](#limitavailablehoursweekly) | Set working hours for each day of the week. If this setting is added but some days are missing, Ecwid counts such days as disabled.                       |
| disallowDates             | array [disallowDates](#disallowdates)                          | Disables specific ranges of date and time. For example, you can make the datepicker required and disable or limit holiday working hours.                  |

#### limitAvailableHoursWeekly

Limit working hours by setting daily time ranges when the store is open. You can add several time ranges to one day if you have a lunch hour, for example.

Format: JSON object with one or several weekdays, each having an array of time ranges (also in the array format.

<details>

<summary>Code example</summary>

```json
{
    "MON": [
        ["08: 30", "13: 30"],
        ["14: 00", "17: 30"]
    ],
    "TUE": [
        ["08: 30", "17: 30"]
    ],
    "WED": [
        ["08: 30", "13: 30"],
        ["14: 00", "17: 30"]
    ],
    "THU": [
        ["08: 30", "17: 30"]
    ],
    "FRI": [
        ["08: 30", "13: 30"],
        ["14: 00", "17: 30"]
    ],
    "SAT": [
        ["08: 30", "17: 30"]
    ],
    "SUN": [
        ["10: 00", "14: 00"]
    ]
}
```

</details>

#### disallowDates

Disable specific dates or date ranges to close the store for holidays, for example.

Format: JSON array with one or several arrays, each having a date or datetime range formatted to the `"YYYY-MM-DD HH:MM:SS", "YYYY-MM-DD HH:MM:SS"` string.

<details>

<summary>Code example</summary>

```json
{
    "disallowDates": [
        // Disallow placing/shipping orders for the specific day after 3 PM.
        ["2024-04-25 15:00:00", "2024-04-25 23:59:59"],
        // Close the store for 3 days.
        ["2024-04-26 00:00", "2024-04-29 00:00"]
    ]
}
```

</details>

#### overrides

| Field            | Type                                           | Description                                                                                                                                                    |
| ---------------- | ---------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| conditions       | object [conditions](#conditions)               | Conditions for the override. When conditions are met, override happens automatically.                                                                          |
| fieldsToOverride | object [datePickerOptions](#datepickeroptions) | <p>Set up the <code>datePickerOptions</code> object with a new logic inside. <br><br>When conditions are met, it will override defult datepicker behavior.</p> |

#### conditions

| Field          | Type   | Description                                                        |
| -------------- | ------ | ------------------------------------------------------------------ |
| shippingMethod | string | Selected shipping method name, for example. `'Pickup at North st'` |

#### translated

Object with text field translations in the `"lang": "text"` format, where the `"lang"` is an ISO 639-1 language code. For example:

```
{
    "en": "Sample text",
    "nl": "Voorbeeldtekst"
}
```

Translations are available for all active store languages. Only the default language translations are returned if no other translations are provided for the field. Find active store languages with <mark style="color:green;">`GET`</mark> `/profile` request > `languages` > `enabledLanguages`.
