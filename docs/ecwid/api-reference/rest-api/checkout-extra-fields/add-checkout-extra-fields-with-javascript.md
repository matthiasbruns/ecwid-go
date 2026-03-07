# Add checkout extra fields with JavaScript

With JavaScript, you have full control and flexibility over **checkout extra fields**. For example, you can call extra fields based on certain store/cart conditions or set up a datepicker tuned for your store's working hours and delivery schedule.

### Pre-requirements

To set up checkout extra fields on the storefront and manage their content, you need:

1. **Custom application**. You need a private app to be installed in your store. Without the app, it's impossible to ensure that your JS file always loads and executes on the storefront. [**Setup an app**](https://app.gitbook.com/s/uOzT5egoVTAjMJwRuMQT/get-started)
2. **Access scope**. Your custom app requires the following access scope: `customize_storefront`. [**Request an update**](https://app.gitbook.com/s/uOzT5egoVTAjMJwRuMQT/contact-ecwid-api-support-team)
3. **Endpoint**. The app requires one endpoint: `customJsUrl` leading to your JavaScript file. It is required for using the extra fields feature, and additionally grants access to Storefront JS API. [**Request an update**](https://app.gitbook.com/s/uOzT5egoVTAjMJwRuMQT/contact-ecwid-api-support-team)

### Set up basic checkout extra fields

Let's start with a basic text input in the shipping address form at checkout. This would be an optional request for a package sign. If customers type something in that field and place an order, you'll get the field in the `extraFields` field in order details. Optionally, you can make the field visible in order details for customers and in Ecwid admin.

<details>

<summary>Code example of the basic extra field</summary>

```javascript
// Initialize extra fields
window.ec = window.ec || {};
ec.order = ec.order || {};
ec.order.extraFields = ec.order.extraFields || {};

// New optional question 'How should we sign the package?'
// visible on the shipping address page
ec.order.extraFields.wrapping_box_signature = {
    'title': 'How should we sign the package?',
    'textPlaceholder': 'Package sign',
    'type': 'text',
    'tip': 'We will put a label on a box so the recipient knows who it is from',
    'required': false,
    'checkoutDisplaySection': 'shipping_address'
};

window.Ecwid && Ecwid.refreshConfig();
```

</details>

{% hint style="info" %}
The `wrapping_box_signature` in the example above is an internal ID for the extra field called `key` which can be used later on in REST API requests.\
\
You can create multiple checkout extra fields from the same JS file. However, every extra field must have its own unique `key`.
{% endhint %}

Checkout extra fields can be created from any store page with `Ecwid` object, for example, from category pages. However, you need to ensure that the function creating extra fields is not called at the checkout, as it can rewrite the customer's input.

### Learn checkout extra field configuration

Ecwid offers flexible configuration settings for checkout customizations with extra fields. Build a unique UX with different field types, translations, and overrides.

Full description of available config settings for checkout extra fields:

<table><thead><tr><th width="261">Field</th><th width="163">Type</th><th>Description</th></tr></thead><tbody><tr><td>title</td><td>string</td><td><p>Title visible at the checkout above the extra field. Also appears in Ecwid admin and customer notifications (if added).<br></p><p><strong>Required</strong></p></td></tr><tr><td>titleTranslated</td><td>object <a href="#translated">translated</a></td><td>Available translations for the extra field title. </td></tr><tr><td>type</td><td>string</td><td><p>Extra field type that defines its functionality. </p><p><br>One of:<br><code>text</code> - Single-line text input. Supports placeholders and pre-defined values (default).<br><code>textarea</code> - Multiple-line text input. Supports placeholders and pre-defined values.<br><code>select</code> - Drop-down list, where customers can choose only one value. Requires <code>options</code> array.<br><code>radio_buttons</code> - A group of radio buttons, where customers can choose only one value. Requires <code>options</code> array.<br><code>checkbox</code> - A group of checkboxes, where customers can choose multiple values. Requires <code>options</code> array.<br><code>toggle_button_group</code> - A group of buttons, where customers can choose only one value. Requires <code>options</code> array.<br><code>datetime</code> - Customizable date and time picker in the form of a calendar widget on the checkout.<br><code>empty</code> - Non-editable text visible at the checkout and unavailable in customer notifications and Ecwid admin.<br><br><strong>Required</strong></p></td></tr><tr><td>checkoutDisplaySection</td><td>string</td><td><p>Defines at which checkout step extra field shows to customers.</p><p><br>One of:<br><code>email</code> - First checkout step where customers enter their email and apply discounts.<br><code>shipping_address</code> - Second checkout step where customers enter their shipping address.<br><code>pickup_details</code> - Second checkout step where customers choose when and where they want to pick up an order.<br><code>shipping_methods</code> - Third checkout step where customers choose a shipping method.<br><code>pickup_methods</code> - Third checkout step where customers choose a pickup method.<br><code>payment_details</code> - Final checkout step where customers choose payment method and proceed to the payment.<br></p><p>If not specified, the extra field won't appear at the checkout, allowing you to store hidden order data or apply hidden extra charges to the order.</p></td></tr><tr><td>orderDetailsDisplaySection</td><td>string</td><td><p>Defines at which section the extra field shows in customer notifications and Ecwid admin. </p><p><br>One of:<br><code>shipping_info</code> - Shipping/pickup details.<br><code>billing_info</code> - Billing details.<br><code>customer_info</code> - Customer details.<br><code>order_comments</code> - Order comments left by a customer.<br><code>hidden</code> - Extra field is hidden from customer notifications and Ecwid admin.</p><p><br>If not specified, the extra field appears in the <code>order_comments</code> section.</p></td></tr><tr><td>textPlaceholder</td><td>string</td><td><p>Placeholder text for <code>text</code> and <code>textarea</code> field types. Does not affect the actual extra field value.</p><p><br>If not specified, the extra field won't have a placeholder.</p></td></tr><tr><td>textPlaceholderTranslated</td><td>object <a href="#translated">translated</a></td><td>Available translations for the text placeholder.</td></tr><tr><td>value</td><td>string</td><td><p>Default value for the extra field. Any changes by customers or your scripts on the storefront override it.<br></p><p>If not specified, the extra field won't have any default value.</p></td></tr><tr><td>valueTranslated</td><td>object <a href="#translated">translated</a></td><td>Available translations for the extra field value.</td></tr><tr><td>subtitle</td><td>string</td><td>Subtitle visible under the extra field name at the checkout.<br><br>If not specified, the extra field won't have a subtitle.</td></tr><tr><td>subtitleTranslated</td><td>object <a href="#translated">translated</a></td><td>Available translations for the extra field subtitle.</td></tr><tr><td>tip</td><td>string</td><td><p>Extra field tip for <code>text</code> and <code>textarea</code> field types visible inside the input at the checkout.<br></p><p>If not specified, the extra field won't have a tip.</p></td></tr><tr><td>tipTranslated</td><td>object <a href="#translated">translated</a></td><td>Available translations for the extra field tip.</td></tr><tr><td>available</td><td>boolean</td><td><p>Set <code>false</code> to disable the extra field.</p><p><br>If not specified, the extra field is enabled.</p></td></tr><tr><td>showInInvoice</td><td>boolean</td><td><p>Set <code>true</code> to display the extra field in order tax invoices generated by Ecwid. </p><p></p><p>If not specified, the extra field will not be displayed in order tax invoices.</p></td></tr><tr><td>showInNotifications</td><td>boolean</td><td><p>Set <code>true</code> to display the extra field in customer emails generated by Ecwid. If <code>true</code>, requires <code>orderDetailsDisplaySection</code>.</p><p></p><p>If not specified, the extra field will not be displayed in order tax invoices.</p></td></tr><tr><td>required</td><td>boolean</td><td><p>Set <code>true</code>, to make the extra field required at the checkout. This way, customers won't be able to continue checkout until they type/select a value in the field. </p><p></p><p>If not specified, the extra field is not required at the checkout.</p></td></tr><tr><td>showForPaymentMethodIds</td><td>array of strings</td><td><p>Make extra field available to only some payment methods by setting a list of payment method IDs. <br><br>For example, you can collect additional data only for the specific online payment method by adding it alone to the list: <code>"showForPaymentMethodIds": ["4959-2345934622523"]</code>.</p><p></p><p>If not specified, the extra field is not limited by any payment methods.</p></td></tr><tr><td>showForShippingMethodIds</td><td>array of strings</td><td><p>Make extra field available to only some shipping methods by setting a list of shipping method IDs. <br><br>For example, you can collect additional data only for the store pickup method by adding it alone to the list: <code>"showForShippingMethodIds": ["4959-2345934622523"]</code>.</p><p></p><p>If not specified, the extra field is not limited by any shipping methods.</p></td></tr><tr><td>showForCountry</td><td>array of strings</td><td><p>Make extra field available only for some countries by setting a list of country codes. Ecwid will automatically check customer's address details to display the field.<br><br>For example, <code>"showForCountry": ["BE", "US"]</code>.<br></p><p>If not specified, the extra field is not limited by the customer's address details.</p></td></tr><tr><td>options</td><td>array of objects <a href="#collect-tips-with-extra-fields">options</a></td><td><p>Defines a list of available options for <code>select</code>, <code>checkbox</code>, <code>radio_buttons</code>, and<code>toggle_button_group</code> field types.<br></p><p>If not specified, the field type is transformed to <code>text</code>.</p></td></tr><tr><td>datePickerOptions</td><td>object <a href="#datepickeroptions">datePickerOptions</a></td><td><p>Defines settings and conditions for the datepicker at the checkout. For example, limited working hours or disable weekends for the calendar.</p><p><br>Affects only <code>datetime</code> extra field type.</p></td></tr><tr><td>overrides</td><td>object <a href="#overrides">overrides</a></td><td>Defines overrides for extra fields. This setting allows setting up logic and conditional changes for extra fields on the checkout without actual JavaScript code.</td></tr></tbody></table>

#### translated

Object with text field translations in the `"lang": "text"` format, where the `"lang"` is an ISO 639-1 language code. For example:

```
{
    "en": "Sample text",
    "nl": "Voorbeeldtekst"
}
```

Translations are available for all active store languages. Only the default language translations are returned if no other translations are provided for the field. Find active store languages with <mark style="color:green;">`GET`</mark> `/profile` request > `languages` > `enabledLanguages`.

### Collect tips with extra fields

The `select`, `checkbox`, `radio_buttons`, and`toggle_button_group` field types allow you to set up a drop-down list, block of radio buttons, checkboxes, or buttons for customers. Use it to ask customers for additional delivery details or to add a "Leave a tip" block to the checkout.

The `options` field contains an array of objects, where each option can have the following fields:

<table><thead><tr><th>Field</th><th width="192">Type</th><th>Description</th></tr></thead><tbody><tr><td>title</td><td>string</td><td>Option name visible at the checkout.</td></tr><tr><td>titleTranslated</td><td>object <a href="#translated">translated</a></td><td>Available translations for the option name.</td></tr><tr><td>subtitle</td><td>string</td><td>Subtitle visble under the option name at the checkout.</td></tr><tr><td>subtitleTranslated</td><td>object <a href="#translated">translated</a></td><td>Available translations for the option subtitle.</td></tr><tr><td>surcharge</td><td>number</td><td>Tip/surcharge value added to the order total cost.</td></tr><tr><td>surchargeType</td><td>string</td><td>Defines if a surcharge applies a fixed sum or a percentage increase to the cart total.<br><br>One of:<br><code>ABSOLUTE</code> - Adds a fixed sum.<br><code>PERCENT</code> - Adds percentage increase.</td></tr><tr><td>surchargeTaxable</td><td>boolean</td><td>Defines if surcharge/tip is taxable.</td></tr><tr><td>showZeroSurchargeInTotal</td><td>boolean</td><td><p>Defines if zero-sum surcharge shows in order total breakdown. </p><p><br>Set <code>false</code> to hide zero-sum surcharges from customers at the checkout.</p></td></tr><tr><td>surchargeShortName</td><td>object <a href="#surchargeshortname">surchargeShortName</a></td><td>Defines how the surcharge shows in the "Shopping cart" block where the order total is calculated.</td></tr></tbody></table>

#### surchargeShortName

<table><thead><tr><th width="263">Field</th><th width="112">Type</th><th>Description</th></tr></thead><tbody><tr><td>name</td><td>string</td><td>Name of the surcharge for the "Shopping cart" block.</td></tr><tr><td>nameTranslated</td><td>object <a href="#translated">translated</a></td><td>Available translations for the surcharge short name.</td></tr><tr><td>showSurchargePercentValue</td><td>boolean</td><td>Defines if surcharge percentage value shows next to <code>name</code>. Default: <code>true</code>. <br><br>Works only with <code>PERCENT</code> value in <code>surchargeType</code>.</td></tr></tbody></table>

<details>

<summary>Code example for setting up tips</summary>

```javascript
window.ec = window.ec || {};
ec.order = ec.order || {};
ec.order.extraFields = ec.order.extraFields || {};

ec.order.extraFields.tips = {
    'title': 'Tips',
    'type': 'toggleButtonGroup',
    'options': [
    { 
      'title': 'No tips',
    },
    {
      'title': '5%',
      'subtitle': 'Tip 5% from your order total',  
      'surcharge': 5
    },
    {
      'title': '10%',
      'subtitle': 'Tip 10% from your order total',
      'surcharge': 10
    }
  ],
  'surchargeType': 'PERCENT',
  'surchargeShortName': {
    'name': 'Tips',
    'showSurchargePercentValue': true,
    'nameTranslated': {
      'en': 'Tips',
      'nl': 'Tips'
    }
  },
  'showZeroSurchargeInTotal' : false,
  'required': true,
  'checkoutDisplaySection': 'payment_details'
};

window.Ecwid && Ecwid.refreshConfig();
```

</details>

### Apply hidden surcharges with extra fields

With extra fields, you can apply a hidden surcharge to the order, for example, as an extra processing cost for some shipping methods.

To do so, you need to create an extra field with only 4 attributes and no others:

* **`value`**  with an internal name for the surcharge (customers won't be able to see it).
* **`surchargeShortName`** with `showSurchargePercentValue` set to `false`.
* **`options`** with only one option that has a `surcharge` value.
* **`surchargeType`**  that defines if the surcharge is calculated as a flat value (`absolute`) or percentage from the order total (`percent`).

<details>

<summary>Code example for setting up a hidden 5% surcharge</summary>

```javascript
// Initialize extra fields
window.ec = window.ec || {};
ec.order = ec.order || {};
ec.order.extraFields = ec.order.extraFields || {};
// Set order surcharge
ec.order.extraFields.surcharge = {
    'value': 'Custom charge',
    "options": [
    { 
        "title": "Custom charge",
        "surcharge": 5
    },
    ],
    "surchargeShortName": {
        "name": "Surcharge",
         "showSurchargePercentValue": false
    },
    'surchargeType': 'PERCENT'
}
window.Ecwid && Ecwid.refreshConfig();
```

This code will apply a 5% surcharge at the checkout invisible to the customers.

</details>

### Advanced datepicker settings

You can customize datepicker settings with the `datePickerOptions` setting: set up your working hours, limit shipping dates, disable weekdays and holidays, etc.

With `overrides`, you can further customize datepicker settings based on the selected shipping/pickup method. For example, set up different schedules for pickup points.

#### datePickerOptions

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

Format: JSON object with one or several weekdays, each having an array of time ranges (also in the array format).

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

| Field            | Type                                           | Description                                                                                                                                       |
| ---------------- | ---------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| conditions       | object [conditions](#conditions)               | Conditions for the override. When conditions are met, override happens automatically.                                                             |
| fieldsToOverride | object [datePickerOptions](#datepickeroptions) | <p>Set up the datePickerOptions object with a new logic inside. <br><br>When conditions are met, it will override defult datepicker behavior.</p> |

#### conditions

| Field          | Type   | Description                                                        |
| -------------- | ------ | ------------------------------------------------------------------ |
| shippingMethod | string | Selected shipping method name, for example. `'Pickup at North st'` |

<details>

<summary>Code example for advanced datepicker settings and overrides</summary>

```javascript
window.ec = window.ec || {};
ec.order = ec.order || {};
ec.order.extraFields = ec.order.extraFields || {};

// Add pickup time selection for customer
ec.order.extraFields.ecwid_pickup_time = {
  'title': '_msg_ShippingDetails.pickup.customer_header',
  'required': true,
  'type': 'datetime',
  'checkoutDisplaySection': 'pickup_details',
  'orderDetailsDisplaySection': 'order_comments',
  'datePickerOptions': {
    minDate: new Date(new Date().getTime() + 2*60*60*1000), // Add 2h order preparation time
    maxDate: new Date(2020, 12, 31),
    showTime: true,
    autoClose: false,
    use24hour: true,
    incrementMinuteBy: 30,
    limitAvailableHoursWeekly: {
      'MON': [
        ['08:30', '13:30'],
        ['14:00', '17:30']
      ],
      'TUE': [
        ['14:00', '17:30']
      ],
      'WED': [
        ['01:00', '13:30']
      ],
      'THU': [
        ['14:00', '23:30']
      ],
      'FRI': [
        ['14:00', '17:30']
      ]
    }
  },
  'overrides': [
    {
      'conditions': {
        'shippingMethod': 'Pickup at North st'
      },
      'fieldsToOverride': {
        'datePickerOptions': {
          minDate: new Date(new Date().getTime() + 2*60*60*1000),
          maxDate: new Date(2024, 12, 31),
          showTime: true,
          autoClose: false,
          use24hour: true,
          incrementMinuteBy: 30,
          limitAvailableHoursWeekly: {
            'MON': [
              ['08:30', '13:30'],
              ['14:00', '17:30']
            ],
            'TUE': [
              ['14:00', '17:30']
            ]
          },

          // Disallow specific dates
          'disallowDates': [
            // Disallow same-day pickup after 3PM
            ['2024-04-25 15:00:00', '2024-04-25 23:59:59'],

            // Disallow specific time interval (e.g. if you're booked at that time)
            ['2024-04-26 08:30', '2024-04-26 10:00']
          ]
        }
      }
    },
    {
      'conditions': {
        'shippingMethod': 'Pickup at East st'
      },
      'fieldsToOverride': {
        'datePickerOptions': {
          minDate: new Date(new Date().getTime() + 2*60*60*1000),
          maxDate: new Date(2024, 12, 31),
          showTime: true,
          autoClose: false,
          use24hour: true,
          incrementMinuteBy: 30,
          limitAvailableHoursWeekly: {
            SAT: [
              ['08:30', '13:30'],
              ['14:00', '17:30']
            ],
            SUN: [
              ['14:00', '17:30']
            ]
          }
        }
      }
    },
    {
      'conditions': {
        'shippingMethod': 'Pickup at West st'
      },
      'fieldsToOverride': {
        'available': false
      }
    }
  ]
};

Ecwid.refreshConfig && Ecwid.refreshConfig();
```

</details>

### Data size limits of extra fields

When using checkout extra fields, do not exceed the following limits:

* The length of a specific setting in the JavaScript configuration of an extra field must not exceed **255 characters**.
* The total data size saved as extra fields for one order must not exceed **8Kb**.

### Show extra fields in invoices and emails

All extra fields have the `showInNotifications` setting allowing to show extra fields in the customer emails **automatically**. You can access this setting through API or the Ecwid admin if the field was created there.

You can also display extra fields in customer notifications **manually**. This approach is better if email templates in your store are customized. To add a field manually, make sure the field has a `title` and `orderDetailsDisplaySection` value (any value other than `hidden`), then add new code in the email editor using the example below:

<details>

<summary>Code example for adding an extra field to invoices/emails</summary>

```xml
<#list order.extraFields as extraField>
    <#if extraField.title?has_content && extraField.orderDisplaySection?has_content>
        ${extraField.title}: ${extraField.value}
    </#if>
</#list>
```

</details>

Read more on the [customization of customer emails](https://support.ecwid.com/hc/en-us/articles/4988505141148-Advanced-customization-of-email-templates-in-Ecwid).

### Manage checkout extra fields with REST API

With the JavaScript code, you can dynamically call checkout extra fields. However, if you need the same extra field to always appear at the checkout, it is easier to set up with REST API once.

Find a list of available methods below.

<table data-view="cards"><thead><tr><th data-type="content-ref"></th><th data-hidden></th></tr></thead><tbody><tr><td><a href="manage-checkout-extra-fields-with-rest-api/search-checkout-extra-fields">search-checkout-extra-fields</a></td><td></td></tr><tr><td><a href="manage-checkout-extra-fields-with-rest-api/get-checkout-extra-field">get-checkout-extra-field</a></td><td></td></tr><tr><td><a href="manage-checkout-extra-fields-with-rest-api/update-checkout-extra-field">update-checkout-extra-field</a></td><td></td></tr><tr><td><a href="manage-checkout-extra-fields-with-rest-api/create-checkout-extra-field">create-checkout-extra-field</a></td><td></td></tr><tr><td><a href="manage-checkout-extra-fields-with-rest-api/delete-checkout-extra-field">delete-checkout-extra-field</a></td><td></td></tr></tbody></table>
