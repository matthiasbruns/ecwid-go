# Update store profile

<mark style="color:purple;">`PUT`</mark> `https://app.ecwid.com/api/v3/{storeId}/profile`

<details>

<summary>Request and response example</summary>

Request:

```http
PUT /api/v3/1003/profile HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
Content-Type: application/json
Cache-Control: no-cache

{
  "settings": {
    "closed": false,
    "storeDescription": "<p>Welcome to my store!</p>",
    "googleRemarketingEnabled": false,
    "googleAnalyticsId": "UA-654321-1",
    "fbPixelId": "12305215151521",
    "hideOutOfStockProductsInStorefront": false,
    "askCompanyName": true,
    "askConsentToTrackInStorefront": false,
    "askZipCode": true,
    "allowPreordersForOutOfStockProducts": true,
    "showPricePerUnit": false,
    "pinterestTagId": "1251515431215"
  },
  "company": {
    "companyName": "My Company, Inc.",
    "email": "store@example.com",
    "street": "144 West D Street",
    "city": "Encinitas",
    "countryCode": "US",
    "postalCode": "92024",
    "stateOrProvinceCode": "CA",
    "phone": "1(800)5555555"
  },
  "designSettings": {
    "product_list_image_size": "MEDIUM",
    "product_list_image_aspect_ratio": "SQUARE",
    "product_list_product_info_layout": "CENTER",
    "product_list_show_additional_image_on_hover": true,
    "product_list_title_behavior": "SHOW",
    "product_filters_opened_by_default_on_category_page": true
  },
  "socialLinksSettings": {
    "facebook": {
      "url": "https://facebook.com/0123456789012"
    }
  }
}
```

Response:

```json
{
    "updateCount": 1
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `update_store_profile`

### Path params

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Request JSON

A JSON object with the following fields:

<table data-full-width="false"><thead><tr><th width="225">Field</th><th>Type</th><th>Description</th></tr></thead><tbody><tr><td>generalInfo</td><td>Object <a href="#generalinfo">generalInfo</a></td><td>Basic data about Ecwid store: ID, website URL, website platform, Instant Site settings.</td></tr><tr><td>account</td><td>Object <a href="#account">account</a></td><td>Store owner's account details.</td></tr><tr><td>settings</td><td>Object <a href="#settings">settings</a></td><td>Store general settings.</td></tr><tr><td>mailNotifications</td><td>Object <a href="#mailnotifications">mailNotifications</a></td><td>Mail notifications settings.</td></tr><tr><td>phoneNotifications</td><td>Object <a href="#phonenotifications">phoneNotifications</a></td><td>Phone notifications settings.</td></tr><tr><td>company</td><td>Object <a href="#company">company</a></td><td>Information about physical store: company name, phone, address.</td></tr><tr><td>formatsAndUnits</td><td>Object <a href="#formatsandunits">formatsAndUnits</a></td><td>Store formats/untis settings.</td></tr><tr><td>languages</td><td>Object <a href="#languages">languages</a></td><td>Store language settings.</td></tr><tr><td>shipping</td><td>Object <a href="#shipping">shipping</a></td><td>Store shipping settings.</td></tr><tr><td>zones</td><td>Array <a href="#taxsettings">taxSettings</a></td><td>List of store destination zones.</td></tr><tr><td>taxes</td><td>Array <a href="#taxes">taxes</a></td><td>List of store taxes.</td></tr><tr><td>taxSettings</td><td>Object <a href="#taxsettings">taxSettings</a></td><td>Detailed settings for store taxes.</td></tr><tr><td>businessRegistrationID</td><td>Object <a href="#businessregistrationid">businessRegistrationID</a></td><td>Company registration ID, e.g. VAT reg number or company ID, which is set under Settings / Invoice in Control panel</td></tr><tr><td>legalPagesSettings</td><td>Object <a href="#legalpagessettingsdetails">legalPagesSettings</a></td><td>Legal pages settings for a store (<em>System Settings → General → Legal Pages</em>).</td></tr><tr><td>designSettings</td><td>Object <a href="#designsettings">designSettings</a></td><td>Design settings of an Ecwid store. Can be overriden by updating store profile or by customizing design via JS config in storefront.</td></tr><tr><td>productFiltersSettings</td><td>Object <a href="#productfilterssettings">productFiltersSettings</a></td><td>Settings for product filters in a store.</td></tr><tr><td>orderInvoiceSettings</td><td>Object <a href="#orderinvoicesettings">orderInvoiceSettings</a></td><td>Store settings for order invoices.</td></tr><tr><td>socialLinksSettings</td><td>Object <a href="#sociallinkssettings">socialLinksSettings</a></td><td>Store settings for social media accounts.</td></tr><tr><td>registrationAnswers</td><td>Object <a href="#registrationanswers">registrationAnswers</a></td><td>Merchants' answers provided while registering their Ecwid accounts.</td></tr><tr><td>tipsSettings</td><td>Object <a href="ref:get-store-profile#tipssettings">tipsSettings</a></td><td>Store settings for tips.</td></tr></tbody></table>

#### **generalInfo**

<table><thead><tr><th width="187">Field</th><th width="167">Type</th><th>Description</th></tr></thead><tbody><tr><td>storeUrl</td><td>string</td><td>Main website URL.</td></tr><tr><td>starterSite</td><td>Object <a href="#startersite">starterSite</a></td><td>Details of Ecwid Instant site for account. Learn more about <a href="https://support.ecwid.com/hc/en-us/articles/207100069-Instant-site">Instant site</a>.</td></tr><tr><td>websitePlatform</td><td>string</td><td>Website platform that store is added to. Possible values: <code>"wix"</code>, <code>"wordpress"</code>, <code>"iframe"</code>, <code>"joomla"</code>, <code>"yola"</code>, etc. Default is <code>"unknown"</code>.</td></tr></tbody></table>

#### **account**

| Field       | Type   | Description           |
| ----------- | ------ | --------------------- |
| accountName | string | Full store owner name |

#### **settings**

| Field                                       | Type                                                                     | Description                                                                                                                                                                                          |
| ------------------------------------------- | ------------------------------------------------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| closed                                      | boolean                                                                  | `true` if the store is closed for maintenance, `false` otherwise                                                                                                                                     |
| storeName                                   | string                                                                   | The store name displayed in Instant Site                                                                                                                                                             |
| storeDescription                            | string                                                                   | HTML description for the main store page – Store Front page                                                                                                                                          |
| googleRemarketingEnabled                    | boolean                                                                  | `true` if Remarketing with Google Analytics is enabled, `false` otherwise                                                                                                                            |
| googleAnalyticsId                           | string                                                                   | [Google Analytics ID](https://help.ecwid.com/customer/en/portal/articles/1170264-google-analytics) connected to a store                                                                              |
| fbPixelId                                   | string                                                                   | Your Facebook Pixel ID. This field is not returned if it is empty in the Ecwid Control Panel. [Learn more](https://support.ecwid.com/hc/en-us/articles/115004303345-Step-2-Implement-Facebook-pixel) |
| orderCommentsEnabled                        | boolean                                                                  | `true` if order comments feature is enabled, `false` otherwise                                                                                                                                       |
| orderCommentsCaption                        | string                                                                   | Caption for order comments field in storefront                                                                                                                                                       |
| orderCommentsCaptionTranslated              | Object [translations](#translations)                                     | Available translations for the caption for order comments field.                                                                                                                                     |
| orderCommentsRequired                       | boolean                                                                  | `true` if order comments are required to be filled, `false` otherwise                                                                                                                                |
| askZipCode                                  | boolean                                                                  | `true` if the zip code field is shown on the checkout ('Ask for a ZIP/postal code' in checkout settings is enabled), `false` otherwise                                                               |
| showPricePerUnit                            | boolean                                                                  | `true` if the "Show price per unit" option is turned on, otherwise `false`                                                                                                                           |
| hideOutOfStockProductsInStorefront          | boolean                                                                  | `true` if out of stock products are hidden in storefront, `false` otherwise. This setting is located in Ecwid Control Panel > Settings > General > Cart                                              |
| askCompanyName                              | boolean                                                                  | `true` if "Ask for the company name" in checkout settings is enabled, `false` otherwise                                                                                                              |
| favoritesEnabled                            | boolean                                                                  | `true` if favorites feature is enabled for storefront, `false` otherwise                                                                                                                             |
| productReviewsFeatureEnabled                | boolean                                                                  | `true` if product reviews feature is enabled in the store, `false` otherwise.                                                                                                                        |
| defaultProductSortOrder                     | string                                                                   | Default products sort order setting from *Settings > Cart & Checkout*. Possible values: `"DEFINED_BY_STORE_OWNER"`, `"ADDED_TIME_DESC"`, `"PRICE_ASC"`, `"PRICE_DESC"`, `"NAME_ASC"`, `"NAME_DESC"`  |
| abandonedSales                              | Object [abandonedSales](#abandonedsales)                                 | Abandoned sales settings                                                                                                                                                                             |
| salePrice                                   | Object [salePrice](#saleprice)                                           | Sale (compare to) price settings                                                                                                                                                                     |
| showAcceptMarketingCheckbox                 | boolean                                                                  | `true` if merchant shows the checkbox to accept marketing. `false` otherwise                                                                                                                         |
| acceptMarketingCheckboxDefaultValue         | boolean                                                                  | Default value for the checkbox at checkout to accept marketing                                                                                                                                       |
| acceptMarketingCheckboxCustomText           | string                                                                   | Custom text label for the checkbox to accept marketing at checkout                                                                                                                                   |
| acceptMarketingCheckboxCustomTextTranslated | Object [translations](#translations)                                     | Available translations for custom text label for the checkbox to accept marketing at checkout.                                                                                                       |
| askConsentToTrackInStorefront               | boolean                                                                  | `true` if merchant shows warning to accept cookies in storefront. `false` otherwise                                                                                                                  |
| snapPixelId                                 | string                                                                   | Snapchat pixel ID from your [Snapchat business account](https://ads.snapchat.com/)                                                                                                                   |
| pinterestTagId                              | string                                                                   | Pinterest Tag Id from your [Pinterest business account](https://ads.pinterest.com/)                                                                                                                  |
| googleTagId                                 | string                                                                   | Global site tag from your [Google Ads account](https://ads.google.com/intl/en_US/home/)                                                                                                              |
| googleEventId                               | string                                                                   | Event snippet from your [Google Ads account](https://ads.google.com/intl/en_US/home/)                                                                                                                |
| recurringSubscriptionsSettings              | Object [recurringSubscriptionsSettings](#recurringsubscriptionssettings) | Recurring subscription settings information.                                                                                                                                                         |
| allowPreordersForOutOfStockProducts         | boolean                                                                  | `true` if pre-orders for out of stock products are allowed, `false` otherwise.                                                                                                                       |
| linkUpEnabled                               | boolean                                                                  | `true` if [LinkUp integration](https://support.ecwid.com/hc/en-us/articles/8987228834460) is enabled, `false` otherwise                                                                              |

#### **mailNotifications**

| Field                     | Type                                                           | Description                                                                                                                                       |
| ------------------------- | -------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| adminNotificationEmails   | Array of strings                                               | Email addresses, which the store admin notifications are sent to                                                                                  |
| customerOrderMessages     | Object [customerOrderMessages](#customerordermessages)         | Settings for email notifications that are automatically sent to customers to confirm their orders and keep them informed about the order progress |
| adminMessages             | Object [adminMessages](#adminmessages)                         | Settings for email notifications that are automatically sent to the store owner and staff members                                                 |
| customerMarketingMessages | Object [customerMarketingMessages](#customermarketingmessages) | Settings for email notifications that are automatically sent to customers to engage them and increase store sales                                 |

#### **customerOrderMessages**

| Field                 | Type                                                                                | Description                                                                                                          |
| --------------------- | ----------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| orderConfirmation     | Object [MailNotificationsSettings](#mailnotificationssettings)                      | Settings for `Order confirmation` emails. Supported settings: `enabled`, `marketingBlockEnabled`, `discountCouponId` |
| orderStatusChanged    | Object [MailNotificationsSettings](ref:get-store-profile#mailnotificationssettings) | Settings for `Order status changed` emails. Supported settings: `enabled`                                            |
| orderIsReadyForPickup | Object [MailNotificationsSettings](ref:get-store-profile#mailnotificationssettings) | Settings for `Order is ready for pickup` emails. Supported settings: `enabled`                                       |
| downloadEgoods        | Object [MailNotificationsSettings](ref:get-store-profile#mailnotificationssettings) | Settings for `Download e-goods` emails. Supported settings: `enabled`                                                |
| orderShipped          | Object [MailNotificationsSettings](ref:get-store-profile#mailnotificationssettings) | Settings for `Order shipped` emails. Supported settings: `enabled`                                                   |

#### **adminMessages**

| Field                | Type                                                                                | Description                                                                 |
| -------------------- | ----------------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| newOrderPlaced       | Object [MailNotificationsSettings](ref:get-store-profile#mailnotificationssettings) | Settings for `New order placed` emails. Supported settings: `enabled`       |
| lowStockNotification | Object [MailNotificationsSettings](ref:get-store-profile#mailnotificationssettings) | Settings for `Low stock notification` emails. Supported settings: `enabled` |
| weeklyStatsReport    | Object [MailNotificationsSettings](ref:get-store-profile#mailnotificationssettings) | Settings for weekly stats reports. Supported settings: `enabled`            |

#### **customerMarketingMessages**

| Field                       | Type                                                           | Description                                                                                                          |
| --------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| abandonedCartRecovery       | Object [MailNotificationsSettings](#mailnotificationssettings) | Settings for `Order confirmation` emails. Supported settings: `enabled`, `marketingBlockEnabled`, `discountCouponId` |
| favoriteProductsReminder    | Object [MailNotificationsSettings](#mailnotificationssettings) | Settings for `Order status changed` emails. Supported settings: `enabled`, `discountCouponId`                        |
| feedbackRequest             | Object [MailNotificationsSettings](#mailnotificationssettings) | Settings for `Order is ready for pickup` emails. Supported settings: `enabled`, `discountCouponId`                   |
| customerLoyaltyAppreciation | Object [MailNotificationsSettings](#mailnotificationssettings) | Settings for `Order confirmation` emails. Supported settings: `enabled`, `discountCouponId`                          |
| inactiveCustomerReminder    | Object [MailNotificationsSettings](#mailnotificationssettings) | Settings for `Order status changed` emails. Supported settings: `enabled`, `discountCouponId`                        |
| purchaseAnniversary         | Object [MailNotificationsSettings](#mailnotificationssettings) | Settings for `Order is ready for pickup` emails. Supported settings: `enabled`, `discountCouponId`                   |

#### **MailNotificationsSettings**

| Field                 | Type    | Description                                                            |
| --------------------- | ------- | ---------------------------------------------------------------------- |
| enabled               | boolean | `true` if emails are enabled, `false` otherwise                        |
| marketingBlockEnabled | boolean | `true` if the marketing block for emails is enabled, `false` otherwise |
| discountCouponId      | number  | `id` of the discount coupon added to emails                            |

#### **phoneNotifications**

| Field                   | Type             | Description                                                                                                           |
| ----------------------- | ---------------- | --------------------------------------------------------------------------------------------------------------------- |
| adminNotificationPhones | Array of strings | Phone numbers that are used for store admin notifications, supports up to 100 phone numbers ***(for future usage)***. |

#### **recurringSubscriptionsSettings**

| Field                                    | Type    | Description                                                                             |
| ---------------------------------------- | ------- | --------------------------------------------------------------------------------------- |
| showRecurringSubscriptionsInControlPanel | boolean | `true` if recurring subscriptions feature is visible in admin panel, `false` otherwise. |

#### **company**

| Field               | Type   | Description                                                     |
| ------------------- | ------ | --------------------------------------------------------------- |
| companyName         | string | The company name displayed on the invoice                       |
| email               | string | Company (store administrator) email                             |
| street              | string | Company address. 1 or 2 lines separated by a new line character |
| city                | string | Company city                                                    |
| countryCode         | string | A two-letter ISO code of the country                            |
| postalCode          | string | Postal code or ZIP code                                         |
| stateOrProvinceCode | string | State code (e.g. `NY`) or a region name.                        |
| phone               | string | Company phone number                                            |

#### **formatsAndUnits**

| Field                          | Type                                   | Description                                                                                                                                                                                                                                                                                      |
| ------------------------------ | -------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| currency                       | string                                 | 3-letters code of the store currency (ISO 4217). Examples: `USD`, `CAD`                                                                                                                                                                                                                          |
| currencyPrefix                 | string                                 | Currency prefix (e.g. $)                                                                                                                                                                                                                                                                         |
| currencySuffix                 | string                                 | Currency suffix                                                                                                                                                                                                                                                                                  |
| currencyPrecision              | number                                 | Numbers of digits after decimal point in the store prices. E.g. `2` ($2.99) or `0` (¥500).                                                                                                                                                                                                       |
| currencyGroupSeparator         | string                                 | Price thousands separator. Supported values: space `" "`, dot `"."`, comma `","` or empty value `""`.                                                                                                                                                                                            |
| currencyDecimalSeparator       | string                                 | Price decimal separator. Possible values: `.` or `,`                                                                                                                                                                                                                                             |
| currencyTruncateZeroFractional | boolean                                | Hide zero fractional part of the prices in storefront. `true` or `false` .                                                                                                                                                                                                                       |
| currencyRate                   | number                                 | Currency rate in U.S. dollars, as set in the merchant control panel                                                                                                                                                                                                                              |
| weightUnit                     | string                                 | Weight unit. Supported values: `CARAT`, `GRAM`, `OUNCE`, `POUND`, `KILOGRAM`                                                                                                                                                                                                                     |
| weightPrecision                | number                                 | Numbers of digits after decimal point in weights displayed in the store                                                                                                                                                                                                                          |
| weightGroupSeparator           | string                                 | Weight thousands separator. Supported values: space `" "`, dot `"."`, comma `","` or empty value `""`                                                                                                                                                                                            |
| weightDecimalSeparator         | string                                 | Weight decimal separator. Possible values: `.` or `,`                                                                                                                                                                                                                                            |
| weightTruncateZeroFractional   | boolean                                | Hide zero fractional part of the weight values in storefront. `true` or `false` .                                                                                                                                                                                                                |
| dateFormat                     | string                                 | Date format. Only these formats are accepted: `"dd-MM-yyyy"`, `"dd/MM/yyyy"`, `"dd.MM.yyyy"`, `"MM-dd-yyyy"`, `"MM/dd/yyyy"`, `"yyyy/MM/dd"`, `"MMM d, yyyy"`, `"MMMM d, yyyy"`, `"EEE, MMM d, ''yy"`, `"EEE, MMMM d, yyyy"`                                                                     |
| timeFormat                     | string                                 | Time format. Only these formats are accepted: `"HH:mm:ss"`, `"HH:mm"`, `"hh.mm.ss a"`, `"hh:mm a"`                                                                                                                                                                                               |
| timezone                       | string                                 | Store timezone, e.g. `Europe/Moscow`                                                                                                                                                                                                                                                             |
| dimensionsUnit                 | string                                 | Product dimensions units. Supported values: `MM`, `CM`, `IN`, `YD`                                                                                                                                                                                                                               |
| orderNumberPrefix              | string                                 | <p>Prefix for the order ID. Max length: 20 symbols.<br><br>For example, if a prefix is "01\_", then order ID "XGX7J" becomes "01\_XGX7J" in all customer nofications and in Ecwid admin.</p>                                                                                                     |
| orderNumberSuffix              | string                                 | <p>Suffix for the order ID. Max length: 20 symbols.<br><br>For example, if a suffix is "\_25", then order ID "XGX7J" becomes "XGX7J\_25" in all customer nofications and in Ecwid admin.</p>                                                                                                     |
| orderNumberMinDigitsAmount     | number                                 | Minimum digits amount of an order number (can be 0-19 digits).                                                                                                                                                                                                                                   |
| orderNumberNextNumber          | number                                 | Next order number in a store (should be more than 0).                                                                                                                                                                                                                                            |
| addressFormat                  | Object [addressFormat](#addressformat) | Address format: `plain` and `multiline` formats. Displays the way address is written according to the requirements of the country set up in the profile settings. Supports the following variables: `%NAME%`, `%COMPANY_NAME%`, `%STREET%`, `%CITY%`, `%STATE_NAME% %POSTAL%`, `%COUNTRY_NAME%`. |

#### **addressFormat**

| Field     | Type   | Description                                   |
| --------- | ------ | --------------------------------------------- |
| plain     | string | Single line address format, with a delimiter. |
| multiline | string | Multiline address format.                     |

#### **languages**

| Field            | Type             | Description                                                                                               |
| ---------------- | ---------------- | --------------------------------------------------------------------------------------------------------- |
| enabledLanguages | Array of strings | A list of enabled languages in the storefront. First language code is the default language for the store. |
| defaultLanguage  | string           | ISO code of the default language in store                                                                 |

#### **shipping**

| Field           | Type                                      | Description                                                                                                                                   |
| --------------- | ----------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| handlingFee     | Object [handlingFee](#handlingfee)        | Handling fee settings                                                                                                                         |
| shippingOrigin  | Object [shippingOrigin](#shippingorigin)  | Shipping origin address. If matches company address, company address is returned. Available in read-only mode only                            |
| shippingOptions | Array [shippingOptions](#shippingoptions) | Details of each shipping option present in a store. **For public tokens enabled methods are returned** only. Available in read-only mode only |

#### **handlingFee**

| Field       | Type   | Description                                           |
| ----------- | ------ | ----------------------------------------------------- |
| name        | string | Handling fee name set by store admin. E.g. `Wrapping` |
| value       | number | Handling fee value                                    |
| description | string | Handling fee description for customer                 |

#### **shippingOrigin**

| Field               | Type   | Description                                                     |
| ------------------- | ------ | --------------------------------------------------------------- |
| companyName         | string | The company name displayed on the invoice                       |
| email               | string | Company (store administrator) email                             |
| street              | string | Company address. 1 or 2 lines separated by a new line character |
| city                | string | Company city                                                    |
| countryCode         | string | A two-letter ISO code of the country                            |
| postalCode          | string | Postal code or ZIP code                                         |
| stateOrProvinceCode | string | State code (e.g. `NY`) or a region name                         |
| phone               | string | Company phone number                                            |

#### **taxSettings**

<table><thead><tr><th width="188.359375">Field</th><th width="111.703125">Type</th><th>Description</th></tr></thead><tbody><tr><td>automaticTaxEnabled</td><td>boolean</td><td><code>true</code> if store taxes are calculated automatically, <code>else</code> otherwise. As seen in the <em>Ecwid Control Panel > Settings > Taxes > Automatic</em></td></tr><tr><td>taxes</td><td>Array <a href="#taxes">taxes</a></td><td>Manual tax settings for a store</td></tr><tr><td>pricesIncludeTax</td><td>boolean</td><td><code>true</code> if store has "gross prices" setting enabled. <code>false</code> if store has "net prices" setting enabled.</td></tr><tr><td>taxExemptBusiness</td><td>boolean</td><td>Defines if your business is tax-exempt under § 19 UStG. When <code>true</code>, it will display the “Tax exemption § 19 UStG” message to customers to explain the zero VAT rate.</td></tr><tr><td>ukVatRegistered</td><td>boolean</td><td>If <code>true</code> and order is sent from EU to UK - charges VAT for orders less than GBP 135.</td></tr><tr><td>euIossEnabled</td><td>boolean</td><td>If <code>true</code> and order is sent to EU - charges VAT for orders less than EUR 150. For Import One-Stop Shop (IOSS).</td></tr><tr><td>taxOnShippingCalculationScheme</td><td>string</td><td>Shipping tax calculation schemes. Default value: <code>AUTOMATIC</code>. Possible values: <code>AUTOMATIC</code>, <code>BASED_ON_PRODUCT_TAXES_PROPORTION_BY_PRICE</code>, <code>BASED_ON_PRODUCT_TAXES_PROPORTION_BY_WEIGHT</code>, <code>TAXED_SEPARATELY_FROM_PRODUCTS</code></td></tr></tbody></table>

#### **taxes**

| Field              | Type                  | Description                                                                                                                                        |
| ------------------ | --------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| id                 | number                | Unique internal ID of the tax                                                                                                                      |
| name               | string                | Displayed tax name                                                                                                                                 |
| enabled            | boolean               | Whether tax is enabled `true` / `false`                                                                                                            |
| includeInPrice     | boolean               | `true` if the tax rate is included in product prices. More details: [Taxes in Ecwid](http://help.ecwid.com/customer/portal/articles/1182159-taxes) |
| useShippingAddress | boolean               | `true` if the tax is calculated based on shipping address, `false` if billing address is used                                                      |
| taxShipping        | boolean               | `true` is the tax applies to subtotal+shipping cost . `false` if the tax is applied to subtotal only                                               |
| appliedByDefault   | boolean               | `true` if the tax is applied to all products. `false` is the tax is only applied to thos product that have this tax enabled                        |
| defaultTax         | number                | Tax value, in %, when none of the destination zones match                                                                                          |
| rules              | Array [rules](#rules) | Tax rates                                                                                                                                          |

#### **rules**

| Field  | Type   | Description                 |
| ------ | ------ | --------------------------- |
| zoneId | string | Destination zone ID         |
| tax    | number | Tax rate for this zone in % |

#### **zones**

| Field                | Type                               | Description                                                                                                                                                                                                               |
| -------------------- | ---------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| name                 | string                             | Zone displayed name.                                                                                                                                                                                                      |
| countryCodes         | Array of strings                   | Country codes this zone includes .                                                                                                                                                                                        |
| stateOrProvinceCodes | Array of strings                   | State or province codes the zone includes.                                                                                                                                                                                |
| postCodes            | Array of strings                   | Postcode (or zip code) templates this zone includes. More details: [Destination zones in Ecwid](http://help.ecwid.com/customer/portal/articles/1163922-destination-zones).                                                |
| geoPolygons          | Object [geoPolygons](#geopolygons) | Dot coordinates of the polygon (if destination zone is created using [Zone on Map](https://support.ecwid.com/hc/en-us/articles/207100279-Adding-and-managing-destination-zones#adding-a-shipping-zone-using-google-map)). |

#### **geoPolygons**

| Field           | Type            | Description                                                                                                                                                        |
| --------------- | --------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| <*COORDINATES*> | Array of arrays | Each array contains coordinates of a single dot of the polygon. (E.g. `[ [37.036539581171105, -95.66864041664617], [37.07548018723009, -95.6404782452158], ...]`). |

#### **businessRegistrationID**

| Field | Type   | Description                            |
| ----- | ------ | -------------------------------------- |
| name  | string | ID name, e.g. `Vat ID`, `P.IVA`, `ABN` |
| value | string | ID value                               |

#### **starterSite**

| Field          | Type   | Description                                                                          |
| -------------- | ------ | ------------------------------------------------------------------------------------ |
| ecwidSubdomain | string | Store subdomain on ecwid.com domain, e.g. `mysuperstore` in `mysuperstore.ecwid.com` |
| customDomain   | string | Custom Instant site domain, e.g. `www.mysuperstore.com`                              |

#### **legalPagesSettings**

| Field                           | Type                            | Description                                                            |
| ------------------------------- | ------------------------------- | ---------------------------------------------------------------------- |
| requireTermsAgreementAtCheckout | boolean                         | `true` if customers must agree to store's terms of service at checkout |
| legalPages                      | Array [legalPages](#legalpages) | Information about the legal pages set up in a store                    |

#### **legalPages**

| Field                 | Type                                 | Description                                                                                                                     |
| --------------------- | ------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------- |
| type                  | string                               | Legal page type. One of: `"LEGAL_INFO"`, `"SHIPPING_COST_PAYMENT_INFO"`, `"REVOCATION_TERMS"`, `"TERMS"`, `"PRIVACY_STATEMENT"` |
| enabled               | boolean                              | `true` if legal page is shown at checkout process, `false` otherwise                                                            |
| title                 | string                               | Legal page title                                                                                                                |
| titleTranslated       | Object [translations](#translations) | Available translations for legal page title.                                                                                    |
| display               | string                               | Legal page display mode – in a popup or on external URL. One of: `"INLINE"`, `"EXTERNAL_URL"`                                   |
| displayTranslated     | Object [translations](#translations) | Legal translated page display mode – in a popup or on external URL. One of: `"INLINE"`, `"EXTERNAL_URL"`                        |
| text                  | string                               | HTML contents of a legal page                                                                                                   |
| textTranslated        | Object [translations](#translations) | Available translations for legal page text.                                                                                     |
| externalUrl           | string                               | URL to external location of a legal page                                                                                        |
| externalUrlTranslated | Object [translations](#translations) | URL to external location of a translated legal page                                                                             |

#### **designSettings**

<table><thead><tr><th>Field</th><th width="169.27734375">Type</th><th>Description</th></tr></thead><tbody><tr><td>DESIGN_CONFIG_FIELD_NAME</td><td>string or boolean</td><td>Store design settings as seen in <a href="ref:customize-appearance">storefront design customization</a>. If a specific config field is not provided, it will not be changed</td></tr></tbody></table>

#### **productFiltersSettings**

| Field               | Type                                    | Description                                                             |
| ------------------- | --------------------------------------- | ----------------------------------------------------------------------- |
| enabledInStorefront | boolean                                 | `true` if product filters are enabled in storefront. `false` otherwise. |
| filterSections      | Array [filterSections](#filtersections) | Specific product filters                                                |

#### **filterSections**

<table><thead><tr><th width="199.6875">Field</th><th width="199.984375">Type</th><th>Description</th></tr></thead><tbody><tr><td>type</td><td>string</td><td>Type of specific product filter. Possible values: <code>IN_STOCK</code>, <code>ON_SALE</code>, <code>PRICE</code>, <code>CATEGORIES</code>, <code>SEARCH</code>, <code>SKU</code>, <code>OPTION</code>, <code>ATTRIBUTE</code>, <code>LOCATIONS</code>.</td></tr><tr><td>name</td><td>string</td><td>Name of the product field. Works only with <code>OPTION</code> and <code>ATTRIBUTE</code> filter types and is required for them.</td></tr><tr><td>displayComponent</td><td>string</td><td><p>Style of displaying <code>OPTION</code> filters on the storefront. </p><p></p><p>One of: </p><ul><li><code>CHECKBOXES</code> - Default checkboxes style.</li><li><code>BUTTON_GRID</code> - Grid with buttons that better suits product options like "Size".</li></ul></td></tr><tr><td>enabled</td><td>boolean</td><td><code>true</code> if specific product filter is enabled. <code>false</code> otherwise</td></tr></tbody></table>

#### **abandonedSales**

| Field                      | Type    | Description                                                                        |
| -------------------------- | ------- | ---------------------------------------------------------------------------------- |
| autoAbandonedSalesRecovery | boolean | `true` if abandoned sale recovery emails are sent automatically, `false` otherwise |

#### **salePrice**

| Field                   | Type                                 | Description                                                                                                                                 |
| ----------------------- | ------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------- |
| displayOnProductList    | boolean                              | `true` if sale price is displayed on product list and product details page. `false` if sale price is displayed on product details page only |
| oldPriceLabel           | string                               | Text label for sale price name                                                                                                              |
| oldPriceLabelTranslated | Object [translations](#translations) | Translations for sale price text labels                                                                                                     |
| displayDiscount         | string                               | Show discount in three modes: `"NONE"`, `"ABS"` and `"PERCENT`                                                                              |

#### **socialLinksSettings**

| Field     | Type                                | Description                     |
| --------- | ----------------------------------- | ------------------------------- |
| facebook  | Object [facebook](#sociallinksurl)  | Settings for the Facebook page  |
| instagram | Object [instagram](#sociallinksurl) | Settings for the Instagram page |
| twitter   | Object [twitter](#sociallinksurl)   | Settings for the Twitter page   |
| youtube   | Object [youtube](#sociallinksurl)   | Settings for the Youtube page   |
| vk        | Object [vk](#sociallinksurl)        | Settings for the VK page        |
| pinterest | Object [pinterest](#sociallinksurl) | Settings for the Pinterest page |

#### **socialLinksUrl**

| Field | Type   | Description                   |
| ----- | ------ | ----------------------------- |
| url   | string | URL for the social media page |

#### **orderInvoiceSettings**

| Field                                  | Type    | Description                                                                                                                                                     |
| -------------------------------------- | ------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| displayOrderInvoices                   | boolean | If `false`, Ecwid will disable printing and viewing order invoices for customer and store admin. If `true`, order invoices will be available to view and print. |
| attachInvoiceToOrderEmailNotifications | string  | Possible values: `"ATTACH_TO_ALL_EMAILS"`, `"DO_NOT_ATTACH"`.                                                                                                   |
| invoiceLogoUrl                         | string  | Invoice logo URL.                                                                                                                                               |

#### translations

Object with text field translations in the `"lang": "text"` format, where the `"lang"` is an ISO 639-1 language code. For example:

```
{
    "en": "Sample text",
    "nl": "Voorbeeldtekst"
}
```

Translations are available for all active store languages. Only the default language translations are returned if no other translations are provided for the field. Find active store languages with <mark style="color:green;">`GET`</mark> `/profile` request > `languages` > `enabledLanguages`.

#### **registrationAnswers**

| Field          | Type   | Description                                                                                                                                                                                                                                                                                                                                                |
| -------------- | ------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| alreadySelling | string | Answer to the question "Do you already have experience selling online?", supported values: `getting_started`, `offline_only`, `online_different`, `looking_around`                                                                                                                                                                                         |
| goods          | string | Answer to the question "What type of products will you be selling?", supported values: `apparel`, `art`, `auto`, `books`, `electronics`, `food_restaurant`, `food_ecommerce`, `gifts`, `hardware`, `health`, `home`, `jewelry`, `office`, `pet`, `services`, `sports`, `streaming`, `subscription_product`, `toys`, `tobacco`, `adult`, `notsure`, `other` |
| otherGoods     | string | Applicable if the field `goods` has value `other`. Merchant's text answer to the question "Your goods?"                                                                                                                                                                                                                                                    |
| forSomeone     | string | Answer to the question "Are you setting up a store for someone else?", supported values: `yes` or `no`                                                                                                                                                                                                                                                     |
| website        | string | Answer to the question "Do you already have a website?", supported values: `yes` or `no`                                                                                                                                                                                                                                                                   |
| platform       | string | Applicable if the previous answer is `yes`. Answer to the question "What website platform do you use?", supported values: `joomla`, `rapid_weaver`, `wordpress`, `wix`, `weebly`, `blogspot`, `drupal`, `custom_site`, `not_sure`, `other`                                                                                                                 |
| customPlatform | string | Applicable if the field `platform` has value `other`. Merchant's text answer to the question "Your platform?"                                                                                                                                                                                                                                              |
| useFor         | string | Answer to the question "What are you planning to use Ecwid for?"                                                                                                                                                                                                                                                                                           |
| shopEase       | string | Answer to the question "How would you like your shop to be?"                                                                                                                                                                                                                                                                                               |
| costAttitude   | string | Answer to the question "What are your budget preferences?"                                                                                                                                                                                                                                                                                                 |
| pos            | string | Answer to the question "What point-of-sale system are you using?"                                                                                                                                                                                                                                                                                          |
| salesChannels  | string | Answer to the question Where do you sell online?"                                                                                                                                                                                                                                                                                                          |
| ecom           | string | Answer to the question "What e-commerce platform do you use to sell?"                                                                                                                                                                                                                                                                                      |

#### **tipsSettings**

| Field              | Type                                 | Description                                                                                                                                                                                                |
| ------------------ | ------------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| enabled            | boolean                              | `true` if enabled, `false` otherwise                                                                                                                                                                       |
| type               | string                               | <p>Tip type that defines how its value is calculated. Supported values:<br><code>ABSOLUTE</code> - tip is added as a flat value<br><code>PERCENT</code> - tip is added as a percent of the order total</p> |
| options            | Array of numbers                     | Three number values, e.g. `[0, 5, 10]`. Each value defines tip amount.                                                                                                                                     |
| defaultOption      | number                               | Default tip amount. It must match with any value from the `options` array.                                                                                                                                 |
| customTipSettings  | object customTipSettings             | Custom tip settings ("Another amount" option)..                                                                                                                                                            |
| title              | string                               | Text displayed above the tip input field.                                                                                                                                                                  |
| subTitle           | string                               | Grayed-out text displayed upder the tip input field.                                                                                                                                                       |
| titleTranslated    | Object [translations](#translations) | Available translations for tip title.                                                                                                                                                                      |
| subtitleTranslated | Object [translations](#translations) | Available translations for tip subtitle.                                                                                                                                                                   |

#### customTipSettings

| Field   | Type    | Description                                                                                                            |
| ------- | ------- | ---------------------------------------------------------------------------------------------------------------------- |
| enabled | boolean | <p>Defines if customers can input custom tip amount on the storefront. <br><br><code>true</code> if it's possible.</p> |

### Response JSON

A JSON object with the following fields:

| Field       | Type   | Description                                                                                                                                                                                   |
| ----------- | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| updateCount | number | <p>The number of updated items that defines if the request was successful.<br><br>One of:</p><p><code>1</code> if the item was updated,</p><p><code>0</code> if the item was not updated.</p> |
