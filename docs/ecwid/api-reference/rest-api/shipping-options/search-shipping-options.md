# Search shipping options

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/profile/shippingOptions`&#x20;

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_store_profile`

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

All query params are optional.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>lang</td><td>string</td><td>Language ISO code for translations in JSON response, e.g. <code>en</code>, <code>fr</code>. Translates fields like: <code>title</code>, <code>description</code>, <code>pickupInstruction</code>, <code>text</code>, etc.</td></tr></tbody></table>

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

<table><thead><tr><th width="261">Field</th><th width="181">Type</th><th>Description</th></tr></thead><tbody><tr><td>id</td><td>string</td><td>Internal ID of the shipping option.</td></tr><tr><td>title</td><td>string</td><td>Name of shipping option visible at the checkout.</td></tr><tr><td>titleTranslated</td><td>object <a href="#translations">translations</a></td><td>Available translations for the shipping option name.</td></tr><tr><td>enabled</td><td>boolean</td><td>Defines if the shipping option is available at the checkout (<code>true</code>).</td></tr><tr><td>orderby</td><td>number</td><td>Sort position or shipping option at checkout and in store settings. <br><br>Starts with <code>10</code> and iterates by <code>10</code>.</td></tr><tr><td>fulfilmentType</td><td>string</td><td><p>Fulfillment type.<br><br>One of: </p><p><code>pickup</code> – In-store pickup method.</p><p><code>delivery</code> – Local delivery method. </p><p><code>shipping</code> – Shipping method with externally calculated rates.</p></td></tr><tr><td>minimumOrderSubtotal</td><td>number</td><td>Order subtotal before discounts. The delivery method won’t be available at checkout for orders below that amount. The field is displayed if the value is not 0</td></tr><tr><td>destinationZone</td><td>object <a href="#destinationzone">destinationZone</a></td><td>Destination zone set for shipping option.</td></tr><tr><td>deliveryTimeDays</td><td>string</td><td><p>Estimated delivery time in days, for example: <code>4-6</code></p><p></p><p>Equal to the <code>description</code> value.</p></td></tr><tr><td>description</td><td>string</td><td>Shipping method description visible at the checkout.</td></tr><tr><td>descriptionTranslated</td><td>object <a href="#translations">translations</a></td><td>Available translations for the shipping option description.</td></tr><tr><td>carrier</td><td>string</td><td>Internal field. Name of the carrier for Ecwid-built shipping methods.</td></tr><tr><td>carrierMethods</td><td>array <a href="#carriermethods">carrierMethods</a></td><td>Carrier-calculated shipping methods available for this shipping option.</td></tr><tr><td>carrierSettings</td><td>object <a href="#carriersettings">carrierSettings</a></td><td>Shipping carrier settings.</td></tr><tr><td>ratesCalculationType</td><td>string</td><td><p>Rates calculation type. <br><br>One of:</p><p><code>carrier-calculated</code></p><p><code>table</code></p><p><code>flat</code></p><p><code>app</code></p></td></tr><tr><td>shippingCostMarkup</td><td>number</td><td>Shipping cost markup for carrier-calculated methods.</td></tr><tr><td>flatRate</td><td>object <a href="#flatrate">flatRate</a></td><td><p>Flat rate details. </p><p></p><p>Only available if <code>ratesCalculationType</code> is <code>flat</code>.</p></td></tr><tr><td>ratesTable</td><td>object <a href="#ratestable">ratesTable</a></td><td>Custom table rates details.<br><br>Only available if <code>ratesCalculationType</code> is <code>table</code>.</td></tr><tr><td>appClientId</td><td>string</td><td><p><code>client_id</code> value of the partent app.</p><p></p><p>Only available for shipping methods added by apps.</p></td></tr><tr><td>pickupInstruction</td><td>string</td><td>Pickup instructions in HTML format.<br><br>Only available if <code>fulfilmentType</code> is <code>pickup</code>.</td></tr><tr><td>pickupInstructionTranslated</td><td>object <a href="#translations">translations</a></td><td>Available translations for the pickup instructions.</td></tr><tr><td>pickupPrecisionType</td><td>string</td><td><p>Defines the availability of choosing a time in the datepicker for pickup methods at the checkout. <br><br>One of: </p><p><code>DATE</code> – Customers can only choose a pickup date.</p><p><code>DATE_AND_TIME_SLOT</code> - Customers can choose a pickup date and time.</p></td></tr><tr><td>scheduledPickup</td><td>boolean</td><td>Defines if pickup time is scheduled (<code>true</code>).</td></tr><tr><td>fulfillmentTimeInMinutes</td><td>number</td><td>Minimum amount of order preparation time (in minutes).<br><br>Order preparation time is reserved at the checkout. For example, if <code>fulfillmentTimeInMinutes</code> is <code>60</code>, then customers cannot pick any time less than an hour from the current time for delivery/pickup at the checkout.</td></tr><tr><td>businessHours</td><td>object <a href="#businesshours">businessHours</a></td><td>Limitation for available pickup times by days of the week.</td></tr><tr><td>businessHoursLimitationType</td><td>string</td><td>One of: <code>ALLOW_ORDERS_AND_INFORM_CUSTOMERS</code> - makes it possible to place an order using this delivery method at any time, but if delivery doesn't work at the moment when the order is being placed, a warning will be shown to a customer. <code>DISALLOW_ORDERS_AND_INFORM_CUSTOMERS</code> - makes it possible to place an order using this delivery method only during the operational hours. If delivery doesn't work when an order is placed, this delivery method will be shown at the checkout as a disabled one and will contain a note about when delivery will start working again. <code>ALLOW_ORDERS_AND_DONT_INFORM_CUSTOMERS</code> - makes it possible to place an order using this delivery method at any time. Works only for delivery methods with a schedule.</td></tr><tr><td>scheduled</td><td>boolean</td><td><code>true</code> if "Allow to select delivery date or time at checkout" or "Ask for Pickup Date and Time at Checkout" setting is enabled. <code>false</code> otherwise.</td></tr><tr><td>scheduledTimePrecisionType</td><td>string</td><td><p>Defines the availability of choosing a time in the datepicker for delivery methods at the checkout. <br><br>One of: </p><p><code>DATE</code> – Customers can only choose a delivery date.</p><p><code>DATE_AND_TIME_SLOT</code> - Customers can choose a delivery date and time.</p></td></tr><tr><td>timeSlotLengthInMinutes</td><td>number</td><td>Length of the delivery time slot in minutes.</td></tr><tr><td>allowSameDayDelivery</td><td>boolean</td><td><code>true</code> if same-day delivery is allowed. <code>false</code> otherwise.</td></tr><tr><td>cutoffTimeForSameDayDelivery</td><td>string</td><td>Orders placed after this time (in a 24-hour format) will be scheduled for delivery the next business day.</td></tr><tr><td>availabilityPeriod</td><td>string</td><td><p>Maximum possible delivery date for <a href="https://support.ecwid.com/hc/en-us/articles/115000252285-Order-pickup#-setting-up-pickup-date-and-time">local delivery and pickup shipping options</a>.</p><p><br>One of: </p><p><code>TWO_DAYS</code></p><p><code>THREE_DAYS</code></p><p><code>SEVEN_DAYS</code></p><p><code>ONE_MONTH</code></p><p><code>THREE_MONTHS</code></p><p><code>SIX_MONTHS</code></p><p><code>ONE_YEAR</code></p><p><code>UNLIMITED</code> (default)</p></td></tr><tr><td>blackoutDates</td><td>array of objects <a href="#blackoutdates">blackoutDates</a></td><td>Dates when the store doesn’t work, so customers can't choose these dates for local delivery. Each period of dates is a JSON object.</td></tr><tr><td>estimatedShippingTimeAtCheckoutSettings</td><td>object <a href="#estimatedshippingtimeatcheckoutsettings">estimatedShippingTimeAtCheckoutSettings</a></td><td>Information about estimated shipping time shown at checkout. te at checkout" section).</td></tr></tbody></table>

#### destinationZone

| Field                | Type                              | Description                                                                                                                                                                                                                                                                          |
| -------------------- | --------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| name                 | string                            | Zone displayed name.                                                                                                                                                                                                                                                                 |
| countryCodes         | array of strings                  | Country codes this zone includes .                                                                                                                                                                                                                                                   |
| stateOrProvinceCodes | array of strings                  | State or province codes the zone includes. Format: \[country code]-\[state code] Example: shipping zone for Alabama, Arizona, Alaska in United states looks like: `["US-AL","US-AK","US-AZ"]`. Please refer to [Dictionaries](ref:dictionaries) to get right country and state codes |
| postCodes            | array of strings                  | Postcode (or zip code) templates this zone includes. More details: [Destination zones in Ecwid](http://help.ecwid.com/customer/portal/articles/1163922-destination-zones).                                                                                                           |
| geoPolygons          | array [geoPolygons](#geopolygons) | Dot coordinates of the polygon (if destination zone is created using [Zone on Map](https://support.ecwid.com/hc/en-us/articles/207100279-Adding-and-managing-destination-zones#adding-a-shipping-zone-using-google-map)).                                                            |

#### geoPolygons

<table><thead><tr><th width="183">Field</th><th width="142">Type</th><th>Description</th></tr></thead><tbody><tr><td>&#x3C;<em>COORDINATES</em>></td><td>Array of arrays of strings</td><td>Each array contains coordinates of a single dot of the polygon. For example, <code>[ [37.0365395, -95.66864041664617], [37.0754801, -95.6404782452158], ...]</code>).</td></tr></tbody></table>

#### carrierMethods

| Field   | Type    | Description                          |
| ------- | ------- | ------------------------------------ |
| id      | string  | Carrier ID and specific method name  |
| name    | string  | Carrier method name                  |
| enabled | boolean | `true` if enabled, `false` otherwise |
| orderBy | number  | Position of that carrier method      |

#### carrierSettings

| Field                        | Type                                                         | Description                                                                          |
| ---------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| defaultCarrierAccountEnabled | boolean                                                      | `true` if default Ecwid account is enabled to calculate the rates. `false` otherwise |
| defaultPostageDimensions     | object [defaultPostageDimensions](#defaultpostagedimensions) | Default postage dimensions for this shipping option                                  |

#### defaultPostageDimensions

| Field  | Type   | Description       |
| ------ | ------ | ----------------- |
| length | number | Length of postage |
| width  | number | Width of postage  |
| height | number | Height of postage |

#### flatRate

| Field    | Type   | Description                      |
| -------- | ------ | -------------------------------- |
| rateType | string | One of `"ABSOLUTE"`, `"PERCENT"` |
| rate     | number | Shipping rate                    |

#### ratesTable

| Field        | Type                  | Description                                                                                                                                                                                                                                                                                                                                                            |
| ------------ | --------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| tableBasedOn | string                | <p>Defines how rates are applied.</p><p></p><p>One of: </p><p><code>subtotal</code> - Shipping rates are based on order subtotal before any discounts are applied.</p><p><code>discountedSubtotal</code> - Shipping rates are based on order subtotal after all discounts are applied.</p><p><code>weight</code> - Shipping rates are based on order total weight.</p> |
| rates        | array [rates](#rates) | Details of shipping rates table.                                                                                                                                                                                                                                                                                                                                       |

#### rates

| Field      | Type                             | Description                                                                                                                                           |
| ---------- | -------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| conditions | object [conditions](#conditions) | Conditions for this shipping rate in custom table.                                                                                                    |
| rate       | object [rate](#rate)             | <p>Details of shipping rates applied on certain conditions.<br><br>It is possible to combine different rates for more complex rate calculations. </p> |

#### conditions

| Field                  | Type   | Description                                |
| ---------------------- | ------ | ------------------------------------------ |
| weightFrom             | number | "Weight from" condition value              |
| weightTo               | number | "Weight to" condition value                |
| subtotalFrom           | number | "Subtotal from" condition value            |
| subtotalTo             | number | "Subtotal to" condition value              |
| discountedSubtotalFrom | number | "Discounted subtotal from" condition value |
| discountedSubtotalTo   | number | "Discounted subtotal from" condition value |

#### rate

| Field     | Type   | Description                                                                                                                                                                                                   |
| --------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| perOrder  | number | Absolute rate applied to the whole order.                                                                                                                                                                     |
| percent   | number | Percent rate applied to the whole order.                                                                                                                                                                      |
| perItem   | number | <p>Absolute rate based on the total quantity of products in the order.<br><br>For example, if <code>perItem</code> is <code>15</code>, then an order with 4 products will have a rate of <code>60</code>.</p> |
| perWeight | number | <p>Absolute rate based on the total weight of products in the order.<br><br>For example, if <code>perWeight</code> is <code>10</code>, then an order weightning 5kg will have a rate of <code>50</code>.</p>  |

#### businessHours

Limit available hours when customers can pickup their orders in your store.

Format: String made from a JSON `object` with one or several weekdays, each having an `array` of time ranges (also in the `array` format).

{% hint style="info" %}
Quotes must be escaped
{% endhint %}

Code example:&#x20;

```json
"{\"MON\":[[\"07:00\",\"19:00\"]], \"TUE\":[[\"07:00\",\"13:00\"],[\"13:30\",\"19:00\"]]}"
```

#### blackoutDates

| Field            | Type    | Description                                                         |
| ---------------- | ------- | ------------------------------------------------------------------- |
| fromDate         | string  | Starting date of the period, e.g. `2022-04-28`.                     |
| toDate           | string  | The end date of the period, e.g. `2022-04-30`.                      |
| repeatedAnnually | boolean | Specifies whether the period repeats in the following years or not. |

#### estimatedShippingTimeAtCheckoutSettings

<table><thead><tr><th width="209">Field</th><th width="162">Type</th><th>Description</th></tr></thead><tbody><tr><td>estimatedDeliveryDateAtCheckoutEnabled</td><td>boolean</td><td><code>true</code> if the estimated delivery time is shown at checkout, otherwise <code>false</code>.</td></tr><tr><td>estimatedTransitTimeInDays</td><td>array of numbers</td><td>How long it usually takes for the package to be delivered at the shipping address after being handed in to the shipping company. Field uses<code>[from, to]</code> format. The same number can be used, e.g.: <code>[2, 2]</code>. For an approximate time, use different values, e.g.: <code>[2, 6]</code>.</td></tr><tr><td>fulfillmentTimeInDays</td><td>array of numbers</td><td><p>How many days it usually takes you to prepare an order for shipment. That time will be taken into account when calculating the delivery date for customers. Field uses <code>[from, to]</code> format. <br><br>Precise delivery time example: <code>[2, 2]</code></p><p>Approximate delivery time example: <code>[2, 6]</code>.</p></td></tr><tr><td>cutoffTimeForSameDayPacking</td><td>string</td><td>Local time to pack orders received past this time on the next day in a 24-hour format.<br><br>For example: <code>"13:00"</code> makes orders placed after 13:00 to be scheduled for packing and shipping the next business day.</td></tr><tr><td>shippingBusinessDays</td><td>array of strings</td><td>Days of the week when your orders can be delivered. These days will be taken into account when calculating and displaying the approximate delivery date for customers at checkout.<br><br>Format:<br><code>[ "MON", "TUE", "WED", "THU", "FRI", "SUN", "SAT" ]</code></td></tr><tr><td>deliveryDays</td><td>array of strings</td><td>Days of the week when you pack orders for shipment. Your schedule will be taken into account when calculating the delivery date for customers.<br><br>Format:<br><code>[ "MON", "TUE", "WED", "THU", "FRI", "SUN", "SAT" ]</code></td></tr></tbody></table>

#### translations

Object with text field translations in the `"lang": "text"` format, where the `"lang"` is an ISO 639-1 language code. For example:

```
{
    "en": "Sample text",
    "nl": "Voorbeeldtekst"
}
```

Translations are available for all active store languages. Only the default language translations are returned if no other translations are provided for the field. Find active store languages with <mark style="color:green;">`GET`</mark> `/profile` request > `languages` > `enabledLanguages`.
