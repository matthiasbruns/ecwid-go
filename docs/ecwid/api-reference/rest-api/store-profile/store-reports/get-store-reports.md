# Get store reports

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/reports/{reportType}`

### Access scopes

Your app must have the following **access scopes** to make this request: `read_store_stats`

### Path params

All path params are required.

| Param      | Type                             | Description                                                                                                     |
| ---------- | -------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| storeId    | number                           | Ecwid store ID.                                                                                                 |
| reportType | string [reportType](#reporttype) | Report type that defines what data will be received in response. Find the full list of available reports below. |

#### reportType

A list of available reports. Chart types include:

* Chart — Default line chart.
* Piechart — Sliced round chart.
* Table — Table with the data.
* Barchart — Vertical bar chart.

**Visitors:**

| reportType                   | Сhart type                                                |
| ---------------------------- | --------------------------------------------------------- |
| allTraffic                   | chart                                                     |
| siteUniqueNewVisitorsByGroup | table                                                     |
| newVsReturningVisitors       | piechart                                                  |
| visitorsByCities             | <p>piechart<br><br><strong>Not available yet</strong></p> |
| visitorsByCountry            | barchart OR piechart                                      |
| visitorsByDevice             | piechart                                                  |
| visitorsByLanguage           | piechart                                                  |
| siteEmailActivity            | piechart                                                  |
| siteAggregatedActivity       | piechart                                                  |
| siteGroupActivity            | piechart                                                  |
| siteSocialActivity           | piechart                                                  |
| siteContactWidgetActivity    | chart                                                     |
| sitePhoneActivity            | chart                                                     |
| siteAddressActivity          | chart                                                     |
| siteLocationMapActivity      | chart                                                     |

**Conversions:**

| reportType                  | Сhart type                                             |
| --------------------------- | ------------------------------------------------------ |
| salesFunnel                 | barchart                                               |
| topOfCategoriesByViews      | <p>table<br><br><strong>Not available yet</strong></p> |
| topOfProductsByViews        | <p>table<br><br><strong>Not available yet</strong></p> |
| topOfProductsByAddingToCart | <p>table<br><br><strong>Not available yet</strong></p> |

**Orders:**

| reportType                   | Сhart type                                             |
| ---------------------------- | ------------------------------------------------------ |
| allOrders                    | chart                                                  |
| newOrdersVsRepeatOrders      | piechart                                               |
| topOfProductsByOrders        | <p>table<br><br><strong>Not available yet</strong></p> |
| topOfCustomersByOrders       | <p>table<br><br><strong>Not available yet</strong></p> |
| topOfPaymentMethodsByOrders  | table                                                  |
| topOfShippingMethodsByOrders | table                                                  |
| topOfProductsByAvailability  | <p>table<br><br><strong>Not available yet</strong></p> |

**Finances:**

| reportType            | Сhart type                                             |
| --------------------- | ------------------------------------------------------ |
| allRevenue            | chart                                                  |
| allExpenses           | table                                                  |
| allProfit             | <p>chart<br><br><strong>Not available yet</strong></p> |
| topOfProductsByProfit | table                                                  |
| tips                  | <p>table<br><br><strong>Not available yet</strong></p> |

**Marketing:**

| reportType             | Сhart type                                                |
| ---------------------- | --------------------------------------------------------- |
| topOfMarketingSources  | piechart                                                  |
| abandonedCarts         | <p>table<br><br><strong>Not available yet</strong></p>    |
| automatedEmails        | <p>table<br><br><strong>Not available yet</strong></p>    |
| acceptMarketing        | <p>piechart<br><br><strong>Not available yet</strong></p> |
| mailchimpCampaigns     | <p>table<br><br><strong>Not available yet</strong></p>    |
| topOfCouponsByOrders   | <p>table<br><br><strong>Not available yet</strong></p>    |
| topOfDiscountsByOrders | <p>table<br><br><strong>Not available yet</strong></p>    |
| giftCards              | <p>piechart<br><br><strong>Not available yet</strong></p> |

### Query params

All query params are optional.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>startedFrom</td><td>number</td><td>Lower bound of a time interval for report generation. If not specified, report will be generated from the store creation date. <br><br>Supported value: <code>UNIX timestamp</code>, for example: <code>1591646400</code></td></tr><tr><td>endedAt</td><td>number</td><td>Upper bound of a time interval for report generation. If not specified, report will be generated up to the request date and time. <br><br>Supported value: <code>UNIX timestamp</code>, for example: <code>1591679000</code></td></tr><tr><td>timeScaleValue</td><td>string</td><td>Time scale of the chart in response. <br><br>Must be one of: <code>hour</code>, <code>day</code>, <code>week</code>, <code>month</code>, <code>year</code>.</td></tr><tr><td>comparePeriod</td><td>string</td><td><p>Period for comparing and calculating the period-over-period metrics in a received report. If not specified, no such metric will be added to the report.</p><p></p><p>Must be one of:</p><ul><li><code>noComparePeriod</code></li><li><code>similarPeriodInPreviousWeek</code></li><li><code>similarPeriodInPreviousMonth</code></li><li><code>similarPeriodInPreviousYear</code></li><li><code>previousPeriod</code> </li></ul></td></tr></tbody></table>

### Headers

The **Authorization** header is required.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

<table><thead><tr><th width="260.8359375">Name</th><th width="152.5078125">Type</th><th>Description</th></tr></thead><tbody><tr><td>reportType</td><td>string</td><td>Type of the received report.</td></tr><tr><td>startedFrom</td><td>number</td><td>Lower bound of time interval used for report generation. Only present if it was passed as a request query param.</td></tr><tr><td>endedAt</td><td>number</td><td>Upper bound of time interval used for report generation. Only present if it was passed as a request query param.</td></tr><tr><td>timeScaleValue</td><td>string</td><td>Time scale of the chart in response. Only present if it was passed as a request query param.</td></tr><tr><td>firstDayOfWeek</td><td>string</td><td>First day of the week used in the report. One of: <code>MONDAY</code>, <code>SUNDAY</code>.</td></tr><tr><td>comparePeriod</td><td>string</td><td><p>Compare period for the report. </p><p></p><p>One of: </p><p><code>PREVIOUS_PERIOD</code>, </p><p><code>SIMILAR_PERIOD_IN_PREVIOUS_WEEK</code>, </p><p><code>SIMILAR_PERIOD_IN_PREVIOUS_MONTH</code>, </p><p><code>SIMILAR_PERIOD_IN_PREVIOUS_YEAR</code>, </p><p><code>NO_COMPARE_PERIOD</code>.</p></td></tr><tr><td>aggregatedData</td><td>array of objects <a href="#aggregateddata">aggregatedData</a></td><td>Metric values aggregated for the set period.</td></tr><tr><td>dataset</td><td>array of objects <a href="#dataset">dataSet</a></td><td>Part of the report defined by the time/device/region/etc. that depends on the report type (except <code>table</code> chart type).</td></tr><tr><td>comparePeriodAggregatedData</td><td>object comparePeriodAggregatedData</td><td>Metric values aggregated for the specified compare period.<br><br>Requires any <code>comparePeriod</code> except for the <code>NO_COMPARE_PERIOD</code>.</td></tr><tr><td>comparePeriodDataset</td><td>object comparePeriodDataset</td><td>Partial report defined by the compare period and the report type.<br><br> (except <code>table</code> chart type).<br><br>Requires any <code>reportType</code> except for the <code>table</code> and any <code>comparePeriod</code> except for the <code>NO_COMPARE_PERIOD</code>.</td></tr></tbody></table>

#### aggregatedData

| Name      | Type   | Description              |
| --------- | ------ | ------------------------ |
| dataId    | string | ID of the passed metric. |
| dataValue | number | Metric value.            |

#### dataSet

<table><thead><tr><th width="231">Name</th><th width="145">Type</th><th>Description</th></tr></thead><tbody><tr><td>orderBy</td><td>number</td><td>Sorting number for the datasets in the report.</td></tr><tr><td>datapointId</td><td>string</td><td>Name of the graph, pie segment, or position at the top that describes the data inside. <br><br>Example 1: <code>"Mobile"</code> and <code>"Desktop"</code> are two IDs for the <code>visitorsByDeviceReport</code> report.<br><br>Example 2: <code>"June 1, 2024"</code> and <code>"June 2, 2024"</code> are IDs for the <code>allTraffic</code> report.</td></tr><tr><td>startTimeStamp</td><td>number</td><td>Starting datetime of the specific dataset.<br><br>Available only for the <code>chart</code> chart type.</td></tr><tr><td>endTimeStamp</td><td>number</td><td>Final datetime of the specific dataset.<br><br>Available only for the <code>chart</code> chart type.</td></tr><tr><td>percentage</td><td>number</td><td>Percentage of the "slice" in the report.<br><br>Available only for the <code>piechart</code> report type.</td></tr><tr><td>data</td><td>array of objects</td><td>Aggregated data for the specific dataset in the report. Matches its structure with the <code>aggregatedData</code> field.</td></tr></tbody></table>
