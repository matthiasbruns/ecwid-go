# Get store profile

<mark style="color:green;">`GET`</mark> `https://app.ecwid.com/api/v3/{storeId}/profile`

<details>

<summary>Request and response example</summary>

Request:

```http
GET /api/v3/1003/profile HTTP/1.1
Authorization: Bearer secret_token
Host: app.ecwid.com
```

Response:

```json
{
    "generalInfo": {
        "storeId": 1003,
        "storeUrl": "https://store1003.company.site/products",
        "websitePlatform": "instantsite",
        "profileId": "p3855016",
        "starterSite": {
            "ecwidSubdomain": "demostore",
            "generatedUrl": "https://store1003.company.site",
            "storeLogoUrl": "https://d2j6dbq0eux0bg.cloudfront.net/images/15695068/1107529597.jpg",
            "ecwidSubdomainSuffix": "company.site",
            "slugsWithoutIdsEnabled": true
        }
    },
    "account": {
        "accountName": "API Team Store",
        "accountNickName": "Support team",
        "accountEmail": "ec.apps@lightspeedhq.com",
        "whiteLabel": false,
        "brandName": "Ecwid",
        "supportEmail": "support@ecwid.com",
        "suspended": false,
        "itunesSubscriptionsAvailable": false,
        "googlePlaySubscriptionsAvailable": true,
        "trackStorefrontStats": true,
        "availableFeatures": [
            "AFFILIATE",
            "BULK_INVOICE_PRINTING",
            "BULK_PRICES",
            "CDN",
            "COMBINATIONS",
            "COMPARE_TO_PRICE",
            "CSV_EXPORT",
            "CUSTOM_NOTIFICATIONS",
            "CUSTOMER_GROUPS",
            "DESTINATION_ZONES_LIMIT",
            "DISCOUNT_COUPONS",
            "DISCOUNTS",
            "EBAY_AMAZON_THIRD_PARTY",
            "GOOGLE_SHOPPING",
            "EDIT_INVOICE",
            "EGOODS",
            "EGOODS_LIMITS",
            "FACEBOOK_STORE",
            "FB_SYNC_PRODUCTS",
            "TRACKING_PIXELS",
            "AUTO_FACEBOOK_ADS",
            "GOOGLE_REMARKETING",
            "INVENTORY_TRACKING",
            "LOCAL_PICKUP",
            "MARKETPLACES",
            "ORDER_AMOUNT_LIMITS",
            "PREMIUM",
            "API",
            "SHIPMENT_TRACKING",
            "CUSTOM_DOMAIN",
            "CUSTOM_DOMAIN_STARTER_SITE",
            "HIDE_POWERED_BY_ECWID_BADGE_INSTANTSITE",
            "LINKUP_HIDE_BRANDING",
            "LINKUP_ACTIVATE_ANALYTICS",
            "STORE_STATS",
            "XERO",
            "PRODUCT_FILTERS",
            "FAVORITES",
            "HIDE_POWERED_BY_ECWID_NOTIFICATIONS",
            "VEND",
            "SQUARE_PAYMENTS",
            "HANDLING_FEE",
            "CHECKOUT",
            "ORDER_EDITOR",
            "PRIVATE_ADMIN_NOTES",
            "STRIPE_PAYMENTS",
            "STRIPE_STORED_CREDIT_CARD",
            "MOBILE",
            "ECWID_LIVE_CHAT",
            "ECWID_LIVE_CHAT_TRIAL",
            "AUTOMATIC_TAXES_US",
            "AUTOMATIC_TAXES_REST_OF_THE_WORLD",
            "VK_MARKET",
            "PRODUCT_DIMENSIONS",
            "PRINT_SHIPPING_LABELS",
            "SEO_FIELDS",
            "TAX_EXEMPT_CUSTOMERS",
            "CUSTOMER_TAX_ID",
            "FB_MESSENGER",
            "AUTO_ABANDONED_SALES_RECOVERY",
            "LEGACY_ABANDONED_SALES_VIEW",
            "MULTIADMIN",
            "SCHEDULED_PICKUP_OR_DELIVERY",
            "INSTAGRAM_SHOPPING",
            "STARTERSITE_SITEMAP",
            "INSTANT_SITE_VERIFICATION_CODE",
            "INSTANT_SITE_VERIFICATION_CODE_API",
            "INSTANT_SITE_CUSTOM_JS_CODE",
            "EXPORT_CONTACTS",
            "SHIPPING_PER_PRODUCT",
            "PRINTFUL",
            "CLOVER_PAY",
            "MULTILINGUAL_STORE",
            "GIFT_CARDS",
            "AUTO_YANDEX_ADS",
            "TRIGGERED_EMAILS",
            "MAILCHIMP",
            "WHOLESALE2B",
            "STOREFRONT_LABEL_EDITOR",
            "SET_DELIVERY_REGION_ON_MAP",
            "TIPS",
            "INSTANT_SITE_SALE_FEATURE",
            "WEBSITES_SALE_FEATURE",
            "BUY_BUTTONS_SALE_FEATURE",
            "WORDPRESS_SALE_FEATURE",
            "JOOMLA_SALE_FEATURE",
            "WIX_SALE_FEATURE",
            "WEEBLY_SALE_FEATURE",
            "SQUARESPACE_SALE_FEATURE",
            "RAPID_WEAVER_SALE_FEATURE",
            "PRODUCT_SUBTITLES_FEATURE",
            "PRODUCT_RIBBONS_FEATURE",
            "NAME_YOUR_PRICE_FEATURE",
            "TAX_INVOICES_FEATURE",
            "RECURRING_SUBSCRIPTION_FEATURE",
            "TIKTOK_SHOPS",
            "CHECKOUT_CUSTOM_FIELDS",
            "PRODUCT_DELIVERY_TIME_FEATURE",
            "COST_PRICE_FEATURE",
            "PREORDERS_FEATURE",
            "PRODUCT_PURCHASE_LIMITS_FEATURE",
            "DOMAIN_PURCHASE_FEATURE",
            "PAYMENT_METHOD_SURCHARGE_FEATURE",
            "PAYPAL_GUEST_CHECKOUT_FEATURE",
            "CUSTOM_ORDER_STATUSES",
            "STAFF_SCOPES",
            "VIDEO_EMBED_IN_GALLERY_FEATURE",
            "CUSTOM_REDIRECTS_FOR_INSTANT_SITES",
            "CUSTOM_URL_SLUGS_FOR_CATALOG_PAGES",
            "INVITE_STAFF_FEATURE",
            "ADVANCED_DISCOUNTS_FEATURE",
            "NEW_INSTANT_SITE_PAGES",
            "LOWEST_PRICE_FEATURE",
            "PRODUCT_REVIEWS_FEATURE",
            "BASIC_ECOMMERCE_FEATURE",
            "LINKUP_EXTERNAL_LINKS",
            "FACEBOOK_PIXEL",
            "ALL_TRAFFIC",
            "NEW_VS_RETURNING_VISITORS",
            "VISITORS_BY_DEVICE",
            "VISITORS_BY_LANGUAGE",
            "VISITORS_BY_COUNTRY",
            "ALL_ORDERS",
            "NEW_ORDERS_VS_REPEAT_ORDERS",
            "INVENTORY_REPORT",
            "TOP_OF_SHIPPING_METHODS_BY_ORDERS",
            "TOP_OF_PAYMENT_METHODS_BY_ORDERS",
            "ALL_REVENUE",
            "ALL_EXPENSES",
            "CONVERSION_REPORT",
            "ADD_TO_CART_CONVERSION",
            "CHECKOUT_SALES_FUNNEL",
            "TOP_OF_MARKETING_SOURCES",
            "CUSTOMERS_BY_MARKETING_CONSENT"
        ],
        "registrationDate": "2018-11-23 14:39:14 +0000",
        "paid": true,
        "limitsAndRestrictions": {
            "maxProductLimit": 2500
        }
    },
    "settings": {
        "closed": false,
        "storeName": "API Team Store",
        "storeDescription": "",
        "storeDescriptionTranslated": {
            "cs": "",
            "en": ""
        },
        "invoiceLogoUrl": "https://d2j6dbq0eux0bg.cloudfront.net/images/15695068/1107529597.jpg",
        "emailLogoUrl": "https://d2j6dbq0eux0bg.cloudfront.net/images/15695068/1107529597.jpg",
        "googleRemarketingEnabled": false,
        "googleAnalyticsId": "UA-132449871-2",
        "orderCommentsEnabled": false,
        "orderCommentsCaption": "Leave order comments here:",
        "orderCommentsCaptionTranslated": {
            "en": "Leave order comments here:"
        },
        "orderCommentsRequired": true,
        "hideOutOfStockProductsInStorefront": false,
        "askCompanyName": true,
        "favoritesEnabled": true,
        "defaultProductSortOrder": "DEFINED_BY_STORE_OWNER",
        "productSortOrderInCart": "TIME_ADDED_TO_CART_ASC",
        "abandonedSales": {
            "autoAbandonedSalesRecovery": true
        },
        "salePrice": {
            "displayOnProductList": true,
            "oldPriceLabel": "",
            "oldPriceLabelTranslated": {
                "cs": "",
                "en": ""
            },
            "displayDiscount": "PERCENT",
            "displayLowestPrice": true
        },
        "showAcceptMarketingCheckbox": false,
        "acceptMarketingCheckboxDefaultValue": true,
        "acceptMarketingCheckboxCustomText": "",
        "acceptMarketingCheckboxCustomTextTranslated": {
            "cs": "",
            "en": ""
        },
        "askConsentToTrackInStorefront": false,
        "wixExternalTrackingEnabled": false,
        "openBagOnAddition": false,
        "requirePhoneOnCheckout": true,
        "askZipCode": true,
        "showPricePerUnit": true,
        "askTaxId": false,
        "googleProductCategory": 412,
        "googleProductCategoryName": "Food, Beverages & Tobacco",
        "productCondition": "NEW",
        "recurringSubscriptionsSettings": {
            "showRecurringSubscriptionsInControlPanel": true,
            "supportedPaymentMethodsStatuses": {
                "supportedPaymentMethodsAreAvailable": true,
                "supportedPaymentMethodsAreConnected": false
            }
        },
        "allowPreordersForOutOfStockProducts": false,
        "showCostPriceInControlPanel": false,
        "rootCategorySeoTitle": "",
        "rootCategorySeoTitleTranslated": {
            "cs": "",
            "en": ""
        },
        "rootCategorySeoDescription": "",
        "rootCategorySeoDescriptionTranslated": {
            "cs": "",
            "en": ""
        },
        "showRepeatOrderButton": false,
        "productReviewsFeatureEnabled": false,
        "linkUpEnabled": true
    },
    "mailNotifications": {
        "adminNotificationEmails": [
            "ec.apps@lightspeedhq.com"
        ],
        "customerNotificationFromEmail": "ec.support@lightspeedhq.com",
        "customerOrderMessages": {
            "orderConfirmation": {
                "enabled": true,
                "marketingBlockEnabled": true
            },
            "orderStatusChanged": {
                "enabled": true
            },
            "orderIsReadyForPickup": {
                "enabled": true
            },
            "downloadEgoods": {
                "enabled": true
            },
            "orderShipped": {
                "enabled": true
            },
            "orderDelivered": {
                "enabled": false
            }
        },
        "adminMessages": {
            "newOrderPlaced": {
                "enabled": true
            },
            "lowStockNotification": {
                "enabled": true
            },
            "weeklyStatsReport": {
                "enabled": true
            }
        },
        "customerMarketingMessages": {
            "abandonedCartRecovery": {
                "enabled": true,
                "marketingBlockEnabled": false
            },
            "favoriteProductsReminder": {
                "enabled": false
            },
            "feedbackRequest": {
                "enabled": true
            },
            "customerLoyaltyAppreciation": {
                "enabled": false
            },
            "inactiveCustomerReminder": {
                "enabled": false
            },
            "purchaseAnniversary": {
                "enabled": true
            }
        }
    },
    "phoneNotifications": {
        "adminNotificationPhones": []
    },
    "company": {
        "companyName": "Ecwid API Team Store",
        "email": "ec.apps@lightspeedhq.com",
        "street": "Test st., 1",
        "city": "Tbilisi",
        "countryCode": "GE",
        "postalCode": "0164",
        "stateOrProvinceCode": "TB",
        "phone": "0123456789"
    },
    "formatsAndUnits": {
        "currency": "EUR",
        "currencyPrefix": "€",
        "currencySuffix": "",
        "currencyGroupSeparator": " ",
        "currencyDecimalSeparator": ",",
        "currencyPrecision": 2,
        "currencyTruncateZeroFractional": false,
        "currencyRate": 1.060951674,
        "weightUnit": "KILOGRAM",
        "weightGroupSeparator": " ",
        "weightDecimalSeparator": ".",
        "weightTruncateZeroFractional": false,
        "timeFormat": "HH:mm:ss",
        "dateFormat": "dd.MM.yyyy",
        "timezone": "Europe/London",
        "dimensionsUnit": "CM",
        "volumeUnit": "ML",
        "orderNumberPrefix": "",
        "orderNumberSuffix": "",
        "addressFormat": {
            "plain": "%NAME%, %COMPANY_NAME%, %STREET%, %POSTAL% %CITY% %STATE_NAME%, %COUNTRY_NAME%",
            "multiline": "%NAME%\n%COMPANY_NAME%\n%STREET%\n%POSTAL% %CITY% %STATE_NAME%\n%COUNTRY_NAME%"
        }
    },
    "languages": {
        "enabledLanguages": [
            "en",
            "cs"
        ],
        "facebookPreferredLocale": "en_US",
        "defaultLanguage": "en"
    },
    "shipping": {
        "handlingFee": {
            "value": 0
        },
        "shippingOrigin": {
            "street": "Industrieweg 9-C",
            "city": "Oirschot",
            "countryCode": "NL",
            "countryName": "Netherlands",
            "postalCode": "5688 DP",
            "stateOrProvinceCode": "NBR"
        },
        "shippingOptions": [
            {
                "id": "6589-1709547151586",
                "title": "Standard shipping",
                "titleTranslated": {
                    "cs": "",
                    "en": "Standard shipping"
                },
                "enabled": true,
                "orderby": 30,
                "fulfilmentType": "shipping",
                "ratesCalculationType": "flat",
                "destinationZone": {
                    "id": "WORLD",
                    "name": "WORLD"
                },
                "businessHours": "{\"THU\":[[\"00:00\",\"00:00\"]],\"TUE\":[[\"00:00\",\"00:00\"]],\"WED\":[[\"00:00\",\"00:00\"]],\"SAT\":[[\"00:00\",\"00:00\"]],\"FRI\":[[\"00:00\",\"00:00\"]],\"MON\":[[\"00:00\",\"00:00\"]],\"SUN\":[[\"00:00\",\"00:00\"]]}",
                "minimumOrderSubtotal": 0,
                "businessHoursLimitationType": "ALLOW_ORDERS_AND_DONT_INFORM_CUSTOMERS",
                "flatRate": {
                    "rateType": "ABSOLUTE",
                    "rate": 20
                },
                "carrier": "",
                "estimatedShippingTimeAtCheckoutSettings": {
                    "estimatedDeliveryDateAtCheckoutEnabled": true,
                    "estimatedTransitTimeInDays": [
                        2,
                        4
                    ],
                    "fulfillmentTimeInDays": [
                        3,
                        3
                    ],
                    "cutoffTimeForSameDayPacking": "16:00",
                    "shippingBusinessDays": [
                        "MON",
                        "TUE",
                        "WED",
                        "THU",
                        "FRI"
                    ],
                    "deliveryDays": [
                        "MON",
                        "TUE",
                        "WED",
                        "THU",
                        "FRI",
                        "SAT"
                    ]
                }
            },
            {
                "id": "3919-1640004025851",
                "title": "FREE Shipping",
                "titleTranslated": {
                    "cs": "",
                    "en": "FREE Shipping"
                },
                "enabled": true,
                "orderby": 40,
                "fulfilmentType": "shipping",
                "destinationZone": {
                    "id": "WORLD",
                    "name": "WORLD"
                },
                "businessHours": "{\"THU\":[[\"00:00\",\"00:00\"]],\"TUE\":[[\"00:00\",\"00:00\"]],\"WED\":[[\"00:00\",\"00:00\"]],\"SAT\":[[\"00:00\",\"00:00\"]],\"FRI\":[[\"00:00\",\"00:00\"]],\"MON\":[[\"00:00\",\"00:00\"]],\"SUN\":[[\"00:00\",\"00:00\"]]}",
                "minimumOrderSubtotal": 0,
                "businessHoursLimitationType": "ALLOW_ORDERS_AND_DONT_INFORM_CUSTOMERS",
                "estimatedShippingTimeAtCheckoutSettings": {
                    "estimatedDeliveryDateAtCheckoutEnabled": false,
                    "estimatedTransitTimeInDays": [
                        null,
                        null
                    ],
                    "fulfillmentTimeInDays": [
                        2,
                        2
                    ],
                    "cutoffTimeForSameDayPacking": "16:00",
                    "shippingBusinessDays": [
                        "MON",
                        "TUE",
                        "WED",
                        "THU",
                        "FRI"
                    ],
                    "deliveryDays": []
                }
            },
            {
                "id": "4959-1595934622523",
                "title": "Pickup",
                "titleTranslated": {
                    "cs": "",
                    "en": "Pickup"
                },
                "enabled": true,
                "orderby": 50,
                "fulfilmentType": "pickup",
                "destinationZone": {
                    "id": "WORLD",
                    "name": "WORLD"
                },
                "businessHours": "{\"MON\":[[\"00:00\",\"00:00\"]], \"TUE\":[[\"00:00\",\"00:00\"]], \"WED\":[[\"00:00\",\"00:00\"]], \"THU\":[[\"00:00\",\"00:00\"]], \"FRI\":[[\"00:00\",\"00:00\"]]}",
                "scheduled": false,
                "fulfillmentTimeInMinutes": 0,
                "pickupInstruction": "",
                "pickupInstructionTranslated": {
                    "cs": "",
                    "en": ""
                },
                "scheduledPickup": false,
                "pickupPreparationTimeHours": 0,
                "pickupBusinessHours": "{\"MON\":[[\"00:00\",\"00:00\"]], \"TUE\":[[\"00:00\",\"00:00\"]], \"WED\":[[\"00:00\",\"00:00\"]], \"THU\":[[\"00:00\",\"00:00\"]], \"FRI\":[[\"00:00\",\"00:00\"]]}",
                "flatRate": {
                    "rateType": "ABSOLUTE",
                    "rate": 0
                }
            }
        ]
    },
    "zones": [
        {
            "id": "2686-1580712192121",
            "name": "Весь мир",
            "countryCodes": [
                "AF",
                "AL",
                "DZ",
                "AS",
                "AD",
                "AO",
                "AI",
                "AQ",
                "AG",
                "AR",
                "AM",
                "AW",
                "AU",
                "AT",
                "AZ",
                "BS",
                "BH",
                "BD",
                "BB",
                "BY",
                "BE",
                "BZ",
                "BJ",
                "BM",
                "BT",
                "BO",
                "BQ",
                "BA",
                "BW",
                "BV",
                "BR",
                "IO",
                "BN",
                "BG",
                "BF",
                "BI",
                "KH",
                "CM",
                "CA",
                "CV",
                "KY",
                "CF",
                "TD",
                "CL",
                "CN",
                "CX",
                "CC",
                "CO",
                "KM",
                "CG",
                "CD",
                "CK",
                "CR",
                "HR",
                "CW",
                "CY",
                "CZ",
                "CI",
                "DK",
                "DJ",
                "DM",
                "DO",
                "EC",
                "EG",
                "SV",
                "GQ",
                "ER",
                "EE",
                "ET",
                "FK",
                "FO",
                "FJ",
                "FI",
                "FR",
                "GF",
                "PF",
                "TF",
                "GA",
                "GM",
                "GE",
                "DE",
                "GH",
                "GI",
                "GR",
                "GL",
                "GD",
                "GP",
                "GU",
                "GT",
                "GG",
                "GN",
                "GW",
                "GY",
                "HT",
                "HM",
                "VA",
                "HN",
                "HK",
                "HU",
                "IS",
                "IN",
                "ID",
                "IQ",
                "IE",
                "IM",
                "IL",
                "IT",
                "JM",
                "JP",
                "JE",
                "JO",
                "KZ",
                "KE",
                "KI",
                "KR",
                "KW",
                "KG",
                "LA",
                "LV",
                "LB",
                "LS",
                "LR",
                "LY",
                "LI",
                "LT",
                "LU",
                "MO",
                "MG",
                "MW",
                "MY",
                "MV",
                "ML",
                "MT",
                "MH",
                "MQ",
                "MR",
                "MU",
                "YT",
                "MX",
                "FM",
                "MD",
                "MC",
                "MN",
                "ME",
                "MS",
                "MA",
                "MZ",
                "MM",
                "NA",
                "NR",
                "NP",
                "NL",
                "NC",
                "NZ",
                "NI",
                "NE",
                "NG",
                "NU",
                "NF",
                "MK",
                "MP",
                "NO",
                "OM",
                "PK",
                "PW",
                "PS",
                "PA",
                "PG",
                "PY",
                "PE",
                "PH",
                "PN",
                "PL",
                "PT",
                "PR",
                "QA",
                "RO",
                "RU",
                "RW",
                "RE",
                "BL",
                "SH",
                "KN",
                "LC",
                "MF",
                "PM",
                "VC",
                "WS",
                "SM",
                "ST",
                "SA",
                "SN",
                "RS",
                "SC",
                "SL",
                "SG",
                "SX",
                "SK",
                "SI",
                "SB",
                "SO",
                "ZA",
                "GS",
                "ES",
                "IC",
                "LK",
                "SD",
                "SR",
                "SJ",
                "SZ",
                "SE",
                "CH",
                "TW",
                "TJ",
                "TZ",
                "TH",
                "TL",
                "TG",
                "TK",
                "TO",
                "TT",
                "TN",
                "TR",
                "TM",
                "TC",
                "TV",
                "UG",
                "UA",
                "AE",
                "GB",
                "US",
                "UM",
                "UY",
                "UZ",
                "VU",
                "VE",
                "VN",
                "VG",
                "VI",
                "WF",
                "EH",
                "YE",
                "ZM",
                "ZW",
                "AX"
            ]
        },
        {
            "id": "WORLD",
            "name": "WORLD"
        }
    ],
    "taxes": [
        {
            "id": 947976181,
            "name": "10% Tax",
            "enabled": true,
            "includeInPrice": true,
            "useShippingAddress": true,
            "taxShipping": false,
            "appliedByDefault": true,
            "defaultTax": 10,
            "rules": []
        }
    ],
    "taxSettings": {
        "automaticTaxEnabled": false,
        "euIossEnabled": false,
        "ukVatRegistered": false,
        "taxes": [
            {
                "id": 947976181,
                "name": "10% Tax",
                "enabled": true,
                "includeInPrice": true,
                "useShippingAddress": true,
                "taxShipping": false,
                "appliedByDefault": true,
                "defaultTax": 10,
                "rules": []
            }
        ],
        "pricesIncludeTax": true,
        "taxExemptBusiness": false,
        "b2b_b2c": "b2c",
        "electronicInvoiceFieldsAtCheckoutEnabled": false,
        "taxOnShippingCalculationScheme": "AUTOMATIC"
    },
    "payment": {
        "paymentOptions": [
            {
                "id": "13379-1606718590771",
                "enabled": true,
                "configured": true,
                "checkoutTitle": "Offline",
                "checkoutTitleTranslated": {
                    "cs": "",
                    "en": "Offline"
                },
                "checkoutDescription": "",
                "paymentProcessorId": "offline",
                "paymentProcessorTitle": "",
                "orderBy": 0,
                "appClientId": "",
                "appNamespace": "",
                "paymentSurcharges": {
                    "type": "ABSOLUTE",
                    "value": 0
                },
                "instructionsForCustomer": {
                    "instructionsTitle": "",
                    "instructions": "<p>TEST</p>",
                    "instructionsTranslated": {
                        "cs": "",
                        "en": "<p>TEST</p>"
                    }
                },
                "methods": []
            },
            {
                "id": "110084370-1655211194360",
                "enabled": true,
                "configured": true,
                "checkoutTitle": "Amazon Pay",
                "checkoutTitleTranslated": {
                    "cs": "",
                    "en": "Amazon Pay"
                },
                "checkoutDescription": "",
                "paymentProcessorId": "customPaymentApp",
                "paymentProcessorTitle": "CUSTOM_PAYMENT_APP-infiniteapps-dev",
                "orderBy": 10,
                "appClientId": "infiniteapps-dev",
                "appNamespace": "infiniteapps-dev",
                "instructionsForCustomer": {
                    "instructionsTitle": "",
                    "instructions": "",
                    "instructionsTranslated": {
                        "cs": "",
                        "en": ""
                    }
                },
                "methods": []
            },
            {
                "id": "217779581-1724151534488",
                "enabled": true,
                "configured": true,
                "checkoutTitle": "TEST PAYMENT",
                "checkoutTitleTranslated": {
                    "cs": "",
                    "en": "TEST PAYMENT"
                },
                "checkoutDescription": "",
                "paymentProcessorId": "customPaymentApp",
                "paymentProcessorTitle": "CUSTOM_PAYMENT_APP-custom-app-15695068-1",
                "orderBy": 20,
                "appClientId": "custom-app-15695068-1",
                "appNamespace": "custom-app-15695068-1",
                "instructionsForCustomer": {
                    "instructionsTitle": "",
                    "instructions": "",
                    "instructionsTranslated": {
                        "cs": "",
                        "en": ""
                    }
                },
                "methods": []
            }
        ],
        "applePay": {
            "enabled": false,
            "available": false
        },
        "applePayOptions": []
    },
    "featureToggles": [
        {
            "name": "ALLOW_EMPTY_AND_NON_UNIQUE_SKU",
            "visible": false,
            "enabled": false
        },
        {
            "name": "CONSECUTIVE_ORDER_IDS",
            "visible": false,
            "enabled": false
        },
        {
            "name": "INSTANT_SITE_V2",
            "visible": true,
            "enabled": true
        },
        {
            "name": "NEW_STARTERSITE",
            "visible": false,
            "enabled": true
        },
        {
            "name": "RANDOM_ORDER_IDS",
            "visible": false,
            "enabled": true
        },
        {
            "name": "REGIONAL_ADDRESS_RULES_ON_CHECKOUT",
            "visible": true,
            "enabled": true
        }
    ],
    "legalPagesSettings": {
        "requireTermsAgreementAtCheckout": false,
        "legalPages": []
    },
    "designSettings": {
        "enable_catalog_on_one_page": false,
        "product_list_image_size": "SMALL",
        "product_list_image_aspect_ratio": "LANDSCAPE_15",
        "product_list_image_position": "FIT",
        "product_list_category_image_aspect_ratio": "LANDSCAPE_1333",
        "product_list_category_image_position": "AUTO",
        "product_list_show_product_images": true,
        "product_list_product_info_layout": "CENTER",
        "product_list_show_frame": false,
        "product_list_show_additional_image_on_hover": false,
        "product_list_title_behavior": "SHOW",
        "product_list_price_behavior": "SHOW",
        "product_list_sku_behavior": "HIDE",
        "product_list_buybutton_behavior": "SHOW",
        "product_list_category_title_behavior": "SHOW_BELOW_IMAGE",
        "product_list_image_has_shadow": false,
        "show_signin_link": true,
        "show_signin_link_with_unified_account_page": false,
        "show_footer_menu": true,
        "show_breadcrumbs": true,
        "product_list_show_sort_viewas_options": true,
        "product_filters_position_search_page": "LEFT",
        "product_filters_position_category_page": "RIGHT",
        "product_filters_opened_by_default_on_category_page": true,
        "product_details_show_product_sku": true,
        "product_details_layout": "TWO_COLUMNS_SIDEBAR_ON_THE_RIGHT",
        "product_details_two_columns_with_right_sidebar_show_product_description_on_sidebar": false,
        "product_details_two_columns_with_left_sidebar_show_product_description_on_sidebar": true,
        "product_details_show_product_name": true,
        "product_details_show_breadcrumbs": true,
        "product_details_show_product_price": true,
        "product_details_show_sale_price": true,
        "product_details_show_tax": true,
        "product_details_show_attributes": true,
        "product_details_show_weight": false,
        "product_details_show_product_description": true,
        "product_details_show_delivery_time": true,
        "product_details_show_product_options": true,
        "product_details_show_wholesale_prices": true,
        "product_details_show_save_for_later": true,
        "product_details_show_share_buttons": true,
        "product_details_position_product_name": 200,
        "product_details_position_breadcrumbs": 100,
        "product_details_position_product_sku": 300,
        "product_details_position_product_price": 400,
        "product_details_position_product_options": 600,
        "product_details_position_buy_button": 700,
        "product_details_position_wholesale_prices": 800,
        "product_details_position_product_description": 2147483647,
        "product_details_position_save_for_later": 1000,
        "product_details_position_share_buttons": 1100,
        "product_details_show_price_per_unit": true,
        "product_details_show_qty": true,
        "product_details_show_in_stock_label": true,
        "product_details_show_number_of_items_in_stock": true,
        "product_details_gallery_layout": "IMAGE_SINGLE_THUMBNAILS_HORIZONTAL",
        "cart_widget_layout": "SMALL_ICON_COUNTER",
        "shopping_cart_show_qty_inputs_on_mobile": true,
        "enable_page_transitions": true,
        "product_list_subtitles_behavior": "SHOW",
        "product_details_show_subtitle": true,
        "product_details_position_subtitle": 500,
        "product_details_show_navigation_arrows": true,
        "product_details_show_product_photo_zoom": true,
        "product_details_show_breadcrumbs_position": "PRODUCT_DETAILS_SIDEBAR",
        "product_details_position_review_section": 950,
        "product_details_show_rating_section": true,
        "product_details_show_reviews_section": true,
        "product_list_rating_section_behavior": "SHOW",
        "show_numeric_rating_in_five_stars_view": true,
        "show_rating_section_in_single_star_view": false,
        "show_review_count_in_five_stars_view": true,
        "show_review_section_in_single_review_view": true
    },
    "productFiltersSettings": {
        "enabledInStorefront": true,
        "filterSections": [
            {
                "type": "IN_STOCK",
                "enabled": true
            },
            {
                "type": "ON_SALE",
                "enabled": true
            },
            {
                "type": "PRICE",
                "enabled": true
            },
            {
                "type": "CATEGORIES",
                "enabled": true
            },
            {
                "type": "SEARCH",
                "enabled": true
            },
            {
                "name": "Custom Attribute 1",
                "type": "ATTRIBUTE",
                "enabled": true
            }
        ]
    },
    "orderInvoiceSettings": {
        "invoiceLogoUrl": "https://d2j6dbq0eux0bg.cloudfront.net/images/15695068/1107529597.jpg",
        "attachInvoiceToOrderEmailNotifications": "DO_NOT_ATTACH"
    },
    "socialLinksSettings": {
        "facebook": {},
        "instagram": {},
        "twitter": {},
        "youtube": {},
        "vk": {},
        "pinterest": {}
    },
    "registrationAnswers": {
        "alreadySelling": "offline_only",
        "goods": "health",
        "forSomeone": "yes",
        "website": "no"
    },
    "tipsSettings": {
        "enabled": false,
        "type": "PERCENT",
        "options": [
            0,
            5,
            10
        ],
        "defaultOption": 0,
        "title": "Support us with a donation",
        "subtitle": "We appreciate all donations, and even the tiniest bit helps us continue what we’re doing.",
        "titleTranslated": {
            "cs": "",
            "en": "Support us with a donation"
        },
        "subtitleTranslated": {
            "cs": "",
            "en": "We appreciate all donations, and even the tiniest bit helps us continue what we’re doing."
        }
    },
    "taxInvoiceSettings": {
        "taxInvoiceLogoUrl": "https://d2j6dbq0eux0bg.cloudfront.net/images/15695068/1107529597.jpg",
        "attachTaxInvoiceToOrderEmailNotifications": "ATTACH_TO_ALL_EMAILS",
        "enableTaxInvoices": true,
        "generateInvoicesAutomatically": "ON_ORDER_PLACED",
        "taxInvoiceIdPrefix": "",
        "taxInvoiceIdSuffix": "",
        "taxInvoiceIdMinDigitsAmount": 2,
        "taxInvoiceIdNextNumber": 28
    }
}
```

</details>

### Required access scopes

Your app must have the following **access scopes** to make this request: `read_store_profile`

Additional scopes include `read_store_limits` granting access to [store limits](#limitsandrestrictions) and `read_store_profile_extended` granting access to [detailed store billing](#accountbilling) data.

### Path params

All path params are required.

| Param   | Type   | Description     |
| ------- | ------ | --------------- |
| storeId | number | Ecwid store ID. |

### Query params

All query params are optional.

<table data-full-width="false"><thead><tr><th width="187">Name</th><th width="97">Type</th><th>Description</th></tr></thead><tbody><tr><td>showExtendedInfo</td><td>boolean</td><td>Set <code>true</code> to receive additional store profile data including account/billing data. Requires <code>read_store_profile_extended</code> access scope.</td></tr><tr><td>lang</td><td>string</td><td>Language ISO code for translations in JSON response, e.g. <code>en</code>, <code>fr</code>. Translates fields like: <code>title</code>, <code>description</code>, <code>pickupInstruction</code>, <code>text</code>, etc.</td></tr><tr><td>getReportAvailabilityStatus</td><td>boolean</td><td>Set <code>true</code> to receive <code>account</code> > <code>reportAvailabilityDetails</code> field in response.</td></tr><tr><td>getStoreVertical</td><td>boolean</td><td>Set <code>true</code> to receive <code>account</code> > <code>storeVertical</code> field in response.</td></tr><tr><td>getFeaturesByPlans</td><td>boolean</td><td>Set <code>true</code> to receive <code>account</code> > <code>featuresByPlans</code> object in response.</td></tr><tr><td>responseFields</td><td>string</td><td>Specify the exact fields to receive in response JSON. If not specified, the response JSON will have all available fields for the entity.<br>Example: <code>?responseFields=generalInfo(storeId,storeUrl)</code></td></tr></tbody></table>

Example of using `responseFields` param:

{% tabs %}
{% tab title="Request" %}

```
curl --location 'https://app.ecwid.com/api/v3/1003/profile?responseFields=generalInfo(storeId,storeUrl)' \
--header 'Authorization: Bearer secret_ab***cd'
```

{% endtab %}

{% tab title="Response" %}

```json
{
    "generalInfo": {
        "storeId": 1003,
        "storeUrl": "https://store1003.company.site/"
    }
}
```

{% endtab %}
{% endtabs %}

### Headers

The **Authorization** header is required. Request works with **any access token**, though the public token receives limited data.

<table><thead><tr><th>Header</th><th width="252">Format</th><th>Description</th></tr></thead><tbody><tr><td>Authorization</td><td><code>Bearer secret_ab***cd</code></td><td>Access token of the application.</td></tr></tbody></table>

### Response JSON

A JSON object with the following fields:

<table data-full-width="false"><thead><tr><th width="201.94921875">Field</th><th width="245.859375">Type</th><th>Description</th></tr></thead><tbody><tr><td>generalInfo</td><td>Object <a href="#generalinfo">generalInfo</a></td><td>Basic data about Ecwid store: ID, website URL, website platform, Instant Site settings.</td></tr><tr><td>account</td><td>Object <a href="#account">account</a></td><td>Store owner's account details.</td></tr><tr><td>settings</td><td>Object <a href="#settings">settings</a></td><td>Store general settings.</td></tr><tr><td>mailNotifications</td><td>Object <a href="#mailnotifications">mailNotifications</a></td><td>Mail notifications settings.</td></tr><tr><td>phoneNotifications</td><td>Object <a href="#phonenotifications">phoneNotifications</a></td><td>Phone notifications settings.</td></tr><tr><td>company</td><td>Object <a href="#company">company</a></td><td>Information about physical store: company name, phone, address.</td></tr><tr><td>formatsAndUnits</td><td>Object <a href="#formatsandunits">formatsAndUnits</a></td><td>Store formats/untis settings.</td></tr><tr><td>languages</td><td>Object <a href="#languages">languages</a></td><td>Store language settings.</td></tr><tr><td>shipping</td><td>Object <a href="#shipping">shipping</a></td><td>Store shipping settings.</td></tr><tr><td>zones</td><td>Array <a href="#taxsettings">taxSettings</a></td><td>List of store destination zones.</td></tr><tr><td>taxes</td><td>Array <a href="#taxes">taxes</a></td><td>List of store taxes.</td></tr><tr><td>taxSettings</td><td>Object <a href="#taxsettings">taxSettings</a></td><td>Detailed settings for store taxes.</td></tr><tr><td>businessRegistrationID</td><td>Object <a href="#businessregistrationid">businessRegistrationID</a></td><td>Company registration ID, e.g. VAT reg number or company ID, which is set under Settings / Invoice in Control panel.</td></tr><tr><td>payment</td><td>Object <a href="#payment">payment</a></td><td>Store payment settings information.<br><strong>Read only</strong></td></tr><tr><td>featureToggles</td><td>Array <a href="#featuretoggles">featureToggles</a></td><td>Information about enabled/disabled new store features and their visibility in Ecwid Control Panel. Not provided via public token. Some of them are available in Ecwid JS API.<br><strong>Read only</strong></td></tr><tr><td>legalPagesSettings</td><td>Object <a href="#legalpagessettingsdetails">legalPagesSettings</a></td><td>Legal pages settings for a store (<em>System Settings → General → Legal Pages</em>).</td></tr><tr><td>designSettings</td><td>Object <a href="#designsettings">designSettings</a></td><td>Design settings of an Ecwid store. Can be overriden by updating store profile or by customizing design via JS config in storefront.</td></tr><tr><td>productFiltersSettings</td><td>Object <a href="#productfilterssettings">productFiltersSettings</a></td><td>Settings for product filters in a store.</td></tr><tr><td>verticalSettings</td><td>Object <a href="#verticalsettings">verticalSettings</a></td><td>Store vertical info.</td></tr><tr><td>fbMessengerSettings</td><td>Object <a href="#fbmessengersettings">fbMessengerSettings</a></td><td>Store settings for FB Messenger feature. <br><strong>Read only</strong></td></tr><tr><td>mailchimpSettings</td><td>Object <a href="#mailchimpsettings">mailchimpSettings</a></td><td>Store settings for Mailchimp integration. <br><strong>Read only</strong></td></tr><tr><td>orderInvoiceSettings</td><td>Object <a href="#orderinvoicesettings">orderInvoiceSettings</a></td><td>Store settings for order invoices.</td></tr><tr><td>socialLinksSettings</td><td>Object <a href="#sociallinkssettings">socialLinksSettings</a></td><td>Store settings for social media accounts.</td></tr><tr><td>registrationAnswers</td><td>Object <a href="#registrationanswers">registrationAnswers</a></td><td>Merchants' answers provided while registering their Ecwid accounts.</td></tr><tr><td>giftCardSettings</td><td>Object <a href="#giftcardsettings">giftCardSettings</a></td><td>Store settings for gift cards.<br><strong>Read only</strong></td></tr><tr><td>tipsSettings</td><td>Object <a href="ref:get-store-profile#tipssettings">tipsSettings</a></td><td>Store settings for tips.</td></tr><tr><td>accountBilling</td><td>Object <a href="#accountbilling">accountBilling</a></td><td>Store billing and plan info. Requires <code>read_store_profile_extended</code> access scope.<br><strong>Read only</strong></td></tr></tbody></table>

#### **generalInfo**

<table><thead><tr><th width="208.97265625">Field</th><th width="114.25390625">Type</th><th>Description</th></tr></thead><tbody><tr><td>storeId</td><td>number</td><td>Ecwid Store ID.</td></tr><tr><td>profileId</td><td>string</td><td>Internal profile ID in the store. <br><br>Profile ID is unique for every staff account in the store and is assigned to custom apps. For example, if a staff account creates a custom app, they'll see their profile ID and not the storo owner's ID. </td></tr><tr><td>storeUrl</td><td>string</td><td>Main website URL.</td></tr><tr><td>starterSite</td><td>Object <a href="#startersite">starterSite</a></td><td>Details of Ecwid Instant site for account. Learn more about <a href="https://support.ecwid.com/hc/en-us/articles/207100069-Instant-site">Instant site</a>.</td></tr><tr><td>websitePlatform</td><td>string</td><td>Website platform that store is added to. Possible values: <code>"wix"</code>, <code>"wordpress"</code>, <code>"iframe"</code>, <code>"joomla"</code>, <code>"yola"</code>, etc. Default is <code>"unknown"</code>.</td></tr><tr><td>storefrontUrlFormat</td><td>string</td><td><p>Format of the product browser links (URLs of catalog, account, cart, and checkout pages) on the storefront.<br><br>One of: </p><p><code>IFRAME</code> - Store is added to an external website through iframe window.<br><code>WIX</code> - Store works on a Wix website.<br><code>QUERY</code> - Store uses [deprecated] query-based URL format.<br><code>HASH</code> - Store uses old URL format with "hashbangs".<br><code>CLEAN</code> - Store uses modern clean URL format.<br><br><strong>Read-only</strong></p></td></tr><tr><td>storefrontUrlSlugFormat</td><td>string</td><td><p>Format of product and category page URLs on the storefront. Reflects state of the "slugs without IDs" feature in the store.<br><br>One of:<br><code>WITH_IDS</code> - Old URL format with category/product ID included in the URLs.<br><code>WITHOUT_IDS</code> - New URL format with category/product URLs without IDs.<br><br>For example: </p><ul><li>Product link with ID: </li></ul><p><code>{domain}/products/my_product-p123456</code> </p><ul><li>Product link without ID: </li></ul><p><code>{domain}/products/my_product</code><br><br><strong>Read-only</strong></p></td></tr></tbody></table>

#### **account**

| Field                 | Type                                                   | Description                                                                                                                                                                                                                                   |
| --------------------- | ------------------------------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| accountName           | string                                                 | Full store owner name                                                                                                                                                                                                                         |
| accountNickName       | string                                                 | Store owner nickname on the Ecwid forums                                                                                                                                                                                                      |
| accountEmail          | string                                                 | Store owner email. **Note:** Not WL-friendly. If you want to send notifications/emails to customers, use `adminNotificationEmails`field instead.                                                                                              |
| whiteLabel            | boolean                                                | `true` if Ecwid brand is not mentioned in merchant's interface, `false` otherwise. Read only                                                                                                                                                  |
| suspended             | boolean                                                | `true` if Ecwid account is suspended (prevents the storefront from showing any products or creating orders), `false` otherwise. Read only                                                                                                     |
| availableFeatures     | Array of strings                                       | List of the features available on the store's pricing plan                                                                                                                                                                                    |
| registrationDate      | string                                                 | The store registration date                                                                                                                                                                                                                   |
| limitsAndRestrictions | Object [limitsAndRestrictions](#limitsandrestrictions) | Store limits and restrictions, e.g. maximum number of available products. Requires `read_store_limits` access scope.                                                                                                                          |
| featuresByPlans       | Object `featuresByPlans`                               | <p>Map of feature name to the minimum plan name where the feature is available for this user.<br><br>Features not available on any plan for the current user are excluded.<br><br>Returned only when <code>getFeaturesByPlans=true</code></p> |

#### **limitsAndRestrictions**

| Field           | Type   | Description                                                                                                                                                                                    |
| --------------- | ------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| maxProductLimit | number | Maximum number of individual products available in the store, e.g. `2500`. This number doesn't include product options or variations. Requires `read_store_profile_extended` scope, read only. |

#### **settings**

| Field                                       | Type                                                                     | Description                                                                                                                                                                                                                                                                                       |
| ------------------------------------------- | ------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| closed                                      | boolean                                                                  | `true` if the store is closed for maintenance, `false` otherwise                                                                                                                                                                                                                                  |
| storeName                                   | string                                                                   | The store name displayed in Instant Site                                                                                                                                                                                                                                                          |
| storeDescription                            | string                                                                   | HTML description for the main store page – Store Front page                                                                                                                                                                                                                                       |
| invoiceLogoUrl                              | string                                                                   | Company logo displayed on the invoice                                                                                                                                                                                                                                                             |
| emailLogoUrl                                | string                                                                   | Company logo displayed in the store email notifications                                                                                                                                                                                                                                           |
| googleRemarketingEnabled                    | boolean                                                                  | `true` if Remarketing with Google Analytics is enabled, `false` otherwise                                                                                                                                                                                                                         |
| googleAnalyticsId                           | string                                                                   | [Google Analytics ID](https://help.ecwid.com/customer/en/portal/articles/1170264-google-analytics) connected to a store                                                                                                                                                                           |
| fbPixelId                                   | string                                                                   | Your Facebook Pixel ID. This field is not returned if it is empty in the Ecwid Control Panel. [Learn more](https://support.ecwid.com/hc/en-us/articles/115004303345-Step-2-Implement-Facebook-pixel)                                                                                              |
| orderCommentsEnabled                        | boolean                                                                  | `true` if order comments feature is enabled, `false` otherwise                                                                                                                                                                                                                                    |
| orderCommentsCaption                        | string                                                                   | Caption for order comments field in storefront                                                                                                                                                                                                                                                    |
| orderCommentsCaptionTranslated              | Object [translations](#translations)                                     | Available translations for the caption for order comments field.                                                                                                                                                                                                                                  |
| orderCommentsRequired                       | boolean                                                                  | `true` if order comments are required to be filled, `false` otherwise                                                                                                                                                                                                                             |
| askZipCode                                  | boolean                                                                  | `true` if the zip code field is shown on the checkout ('Ask for a ZIP/postal code' in checkout settings is enabled), `false` otherwise                                                                                                                                                            |
| showPricePerUnit                            | boolean                                                                  | `true` if the "Show price per unit" option is turned on, otherwise `false`                                                                                                                                                                                                                        |
| hideOutOfStockProductsInStorefront          | boolean                                                                  | `true` if out of stock products are hidden in storefront, `false` otherwise. This setting is located in Ecwid Control Panel > Settings > General > Cart                                                                                                                                           |
| askCompanyName                              | boolean                                                                  | `true` if "Ask for the company name" in checkout settings is enabled, `false` otherwise                                                                                                                                                                                                           |
| favoritesEnabled                            | boolean                                                                  | `true` if favorites feature is enabled for storefront, `false` otherwise                                                                                                                                                                                                                          |
| productReviewsFeatureEnabled                | boolean                                                                  | `true` if product reviews feature is enabled in the store, `false` otherwise.                                                                                                                                                                                                                     |
| defaultProductSortOrder                     | string                                                                   | <p>Default products sorting from <em>Settings > Cart & Checkout</em>. Possible values: <code>"DEFINED\_BY\_STORE\_OWNER"</code><br>, (default), <code>"ADDED\_TIME\_DESC"</code>, <code>"PRICE\_ASC"</code>, <code>"PRICE\_DESC"</code>, <code>"NAME\_ASC"</code>, <code>"NAME\_DESC"</code></p>  |
| defaultAllProductsViewSortOrder             | string                                                                   | <p>Default products sorting when the "one-page catalog" is enabled. <br><br>Possible values: <code>"DEFINED\_BY\_STORE\_OWNER"</code>, <code>"ADDED\_TIME\_DESC"</code>, <code>"PRICE\_ASC"</code>, <code>"PRICE\_DESC"</code>, <code>"NAME\_ASC"</code> (default), <code>"NAME\_DESC"</code></p> |
| abandonedSales                              | Object [abandonedSales](#abandonedsales)                                 | Abandoned sales settings                                                                                                                                                                                                                                                                          |
| salePrice                                   | Object [salePrice](#saleprice)                                           | Sale (compare to) price settings                                                                                                                                                                                                                                                                  |
| showAcceptMarketingCheckbox                 | boolean                                                                  | `true` if merchant shows the checkbox to accept marketing. `false` otherwise                                                                                                                                                                                                                      |
| acceptMarketingCheckboxDefaultValue         | boolean                                                                  | Default value for the checkbox at checkout to accept marketing                                                                                                                                                                                                                                    |
| acceptMarketingCheckboxCustomText           | string                                                                   | Custom text label for the checkbox to accept marketing at checkout                                                                                                                                                                                                                                |
| acceptMarketingCheckboxCustomTextTranslated | Object [translations](#translations)                                     | Available translations for custom text label for the checkbox to accept marketing at checkout.                                                                                                                                                                                                    |
| askConsentToTrackInStorefront               | boolean                                                                  | `true` if merchant shows warning to accept cookies in storefront. `false` otherwise                                                                                                                                                                                                               |
| snapPixelId                                 | string                                                                   | Snapchat pixel ID from your [Snapchat business account](https://ads.snapchat.com/)                                                                                                                                                                                                                |
| pinterestTagId                              | string                                                                   | Pinterest Tag Id from your [Pinterest business account](https://ads.pinterest.com/)                                                                                                                                                                                                               |
| googleTagId                                 | string                                                                   | Global site tag from your [Google Ads account](https://ads.google.com/intl/en_US/home/)                                                                                                                                                                                                           |
| googleEventId                               | string                                                                   | Event snippet from your [Google Ads account](https://ads.google.com/intl/en_US/home/)                                                                                                                                                                                                             |
| recurringSubscriptionsSettings              | Object [recurringSubscriptionsSettings](#recurringsubscriptionssettings) | Recurring subscription settings information.                                                                                                                                                                                                                                                      |
| allowPreordersForOutOfStockProducts         | boolean                                                                  | `true` if pre-orders for out of stock products are allowed, `false` otherwise.                                                                                                                                                                                                                    |
| linkUpEnabled                               | boolean                                                                  | `true` if [LinkUp integration](https://support.ecwid.com/hc/en-us/articles/8987228834460) is enabled, `false` otherwise                                                                                                                                                                           |

#### **mailNotifications**

| Field                         | Type                                                           | Description                                                                                                                                       |
| ----------------------------- | -------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| adminNotificationEmails       | Array of strings                                               | Email addresses, which the store admin notifications are sent to                                                                                  |
| customerNotificationFromEmail | string                                                         | <p>The email address used as the 'reply-to' field in the notifications to customers.<br><br><strong>Read-only</strong></p>                        |
| customerOrderMessages         | Object [customerOrderMessages](#customerordermessages)         | Settings for email notifications that are automatically sent to customers to confirm their orders and keep them informed about the order progress |
| adminMessages                 | Object [adminMessages](#adminmessages)                         | Settings for email notifications that are automatically sent to the store owner and staff members                                                 |
| customerMarketingMessages     | Object [customerMarketingMessages](#customermarketingmessages) | Settings for email notifications that are automatically sent to customers to engage them and increase store sales                                 |

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

| Field                                    | Type                                                                       | Description                                                                             |
| ---------------------------------------- | -------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| showRecurringSubscriptionsInControlPanel | boolean                                                                    | `true` if recurring subscriptions feature is visible in admin panel, `false` otherwise. |
| supportedPaymentMethodsStatuses          | Object [supportedPaymentMethodsStatuses](#supportedpaymentmethodsstatuses) | Supported payment methods statuses information.                                         |

#### **supportedPaymentMethodsStatuses**

| Field                               | Type    | Description                                                                                                       |
| ----------------------------------- | ------- | ----------------------------------------------------------------------------------------------------------------- |
| supportedPaymentMethodsAreAvailable | boolean | `true` if Stripe payment method can be set up in the store, `false` otherwise. Depends on a country. *Read only.* |
| supportedPaymentMethodsAreConnected | boolean | `true` if Stripe payment method it set up in the store, `false` otherwise. *Read only.*                           |

#### **company**

*System Settings → General → Store Profile*

| Field               | Type   | Description                                                            |
| ------------------- | ------ | ---------------------------------------------------------------------- |
| companyName         | string | The company name displayed on the invoice                              |
| email               | string | Company (store administrator) email                                    |
| street              | string | Company address. 1 or 2 lines separated by a new line character        |
| city                | string | Company city                                                           |
| countryCode         | string | A two-letter ISO code of the country                                   |
| postalCode          | string | Postal code or ZIP code                                                |
| stateOrProvinceCode | string | [State code](ref:state-codes-by-country) (e.g. `NY`) or a region name. |
| phone               | string | Company phone number                                                   |

#### **formatsAndUnits**

*System settings → General → Formats & Units.*

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

*System Settings → General → Languages*

| Field                   | Type             | Description                                                                                               |
| ----------------------- | ---------------- | --------------------------------------------------------------------------------------------------------- |
| enabledLanguages        | Array of strings | A list of enabled languages in the storefront. First language code is the default language for the store. |
| facebookPreferredLocale | string           | Language automatically chosen be default in Facebook storefront (if any)                                  |
| defaultLanguage         | string           | ISO code of the default language in store                                                                 |

#### **shipping**

*System Settings → Shipping*

| Field           | Type                                      | Description                                                                                                                                   |
| --------------- | ----------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| handlingFee     | Object [handlingFee](#handlingfee)        | Handling fee settings                                                                                                                         |
| shippingOrigin  | Object [shippingOrigin](#shippingorigin)  | Shipping origin address. If matches company address, company address is returned. Available in read-only mode only                            |
| shippingOptions | Array [shippingOptions](#shippingoptions) | Details of each shipping option present in a store. **For public tokens enabled methods are returned** only. Available in read-only mode only |

#### **handlingFee**

*System Settings → Shipping → Handling Fee*

| Field       | Type   | Description                                           |
| ----------- | ------ | ----------------------------------------------------- |
| name        | string | Handling fee name set by store admin. E.g. `Wrapping` |
| value       | number | Handling fee value                                    |
| description | string | Handling fee description for customer                 |

#### **shippingOrigin**

*Settings → Shipping & Pickup → Origin address*

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

#### **shippingOptions**

| Field                        | Type                                       | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| ---------------------------- | ------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| id                           | string                                     | Unique ID of shipping option                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| title                        | string                                     | Title of shipping option in store settings                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| titleTranslated              | Object [translations](#translations)       | Available translations for shipping option title.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| enabled                      | boolean                                    | `true` if shipping option is used at checkout to calculate shipping. `false` otherwise                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| orderby                      | number                                     | Sort position or shipping option at checkout and in store settings. The smaller the number, the higher the position                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| fulfilmentType               | string                                     | Fulfillment type. `"pickup"` for in-store pickup methods, `"delivery"` for local delivery methods, `"shipping"` for everything else                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| locationId                   | number                                     | ID of the pickup location for Lightspeed X-Series integration                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| minimumOrderSubtotal         | number                                     | Order subtotal before discounts. The delivery method won’t be available at checkout for orders below that amount. The field is displayed if the value is not 0                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| destinationZone              | Array [zones](#zones)                      | Destination zone set for shipping option. **Empty for public token**                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| deliveryTimeDays             | string                                     | Estimated delivery time in days. Currently, it is equal to the `description` value.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| description                  | string                                     | Shipping method description.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| descriptionTranslated        | Object [translations](#translations)       | Available translations for shipping option description.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| carrier                      | string                                     | Carrier used for shipping the order. Is provided for carrier-calculated shipping options                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| carrierMethods               | Array [carrierMethods](#carriermethods)    | Carrier-calculated shipping methods available for this shipping option                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| carrierSettings              | Object [carrierSettings](#carriersettings) | Carrier-calculated shipping option settings                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| ratesCalculationType         | string                                     | Rates calculation type. One of `"carrier-calculated"`, `"table"`, `"flat"`, `"app"`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| shippingCostMarkup           | number                                     | Shipping cost markup for carrier-calculated methods                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| flatRate                     | Object [flatRate](#flatrate)               | Flat rate details                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| ratesTable                   | Object [ratesTable](#ratestable)           | Custom table rates details                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| appClientId                  | string                                     | `client_id` value of the app (for custom shipping apps only)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| pickupInstruction            | string                                     | String of HTML code of instructions on in-store pickup                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| pickupInstructionTranslated  | Object [translations](#translations)       | Available translations for pickup instruction.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| scheduledPickup              | boolean                                    | `true` if pickup time is scheduled, `false` otehrwise. (*Ask for Pickup Date and Time at Checkout* option in pickup settings)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| pickupPreparationTimeHours   | number                                     | Amount of time required for store to prepare pickup (*Order Fulfillment Time* setting) **DEPRECATED**                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| fulfillmentTimeInMinutes     | number                                     | Amount of time (in minutes) required for store to prepare pickup or to deliver an order (*Order Fulfillment Time* setting)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| businessHours                | Object [businessHours](#businesshours)     | Available and scheduled times to pickup orders                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| pickupBusinessHours          | Object [businessHours](#businesshours)     | \[Deprecated] Available and scheduled times to pickup orders (duplicates `businessHours` field)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| businessHoursLimitationType  | string                                     | One of: `ALLOW_ORDERS_AND_INFORM_CUSTOMERS` - makes it possible to place an order using this delivery method at any time, but if delivery doesn't work at the moment when the order is being placed, a warning will be shown to a customer. `DISALLOW_ORDERS_AND_INFORM_CUSTOMERS` - makes it possible to place an order using this delivery method only during the operational hours. If delivery doesn't work when an order is placed, this delivery method will be shown at the checkout as a disabled one and will contain a note about when delivery will start working again. `ALLOW_ORDERS_AND_DONT_INFORM_CUSTOMERS` - makes it possible to place an order using this delivery method at any time. Works only for delivery methods with a schedule. |
| scheduled                    | boolean                                    | `true` if "Allow to select delivery date or time at checkout" or "Ask for Pickup Date and Time at Checkout" setting is enabled. `false` otherwise.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| scheduledTimePrecisionType   | string                                     | Format of how delivery date is chosen at the checkout - date or date and time. One of: `DATE`, `DATE_AND_TIME_SLOT`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| timeSlotLengthInMinutes      | number                                     | Length of the delivery time slot in minutes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| allowSameDayDelivery         | boolean                                    | `true` if same-day delivery is allowed. `false` otherwise.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| cutoffTimeForSameDayDelivery | string                                     | Orders placed after this time (in a 24-hour format) will be scheduled for delivery the next business day.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| availabilityPeriod           | string                                     | The merchant can specify the maximum possible delivery date for [local delivery and pickup shipping options](https://support.ecwid.com/hc/en-us/articles/115000252285-Order-pickup#setting-up-pickup-date-and-time) ("Allow choosing pickup date within"). Values: `THREE_DAYS`, `SEVEN_DAYS`, `ONE_MONTH`, `THREE_MONTHS`, `SIX_MONTHS`, `ONE_YEAR`, `UNLIMITED`.                                                                                                                                                                                                                                                                                                                                                                                          |
| blackoutDates                | Object [blackoutDates](#blackoutdates)     | Dates when the store doesn’t work, so customers can't choose these dates for local delivery. Each period of dates is a JSON object.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |

#### **carrierMethods**

| Field   | Type    | Description                          |
| ------- | ------- | ------------------------------------ |
| id      | string  | Carrier ID and specific method name  |
| name    | string  | Carrier method name                  |
| enabled | boolean | `true` if enabled, `false` otherwise |
| orderBy | number  | Position of that carrier method      |

#### **carrierSettings**

| Field                        | Type                                                         | Description                                                                          |
| ---------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| defaultCarrierAccountEnabled | boolean                                                      | `true` if default Ecwid account is enabled to calculate the rates. `false` otherwise |
| defaultPostageDimensions     | Object [defaultPostageDimensions](#defaultpostagedimensions) | Default postage dimensions for this shipping option                                  |

#### **defaultPostageDimensions**

| Field  | Type   | Description       |
| ------ | ------ | ----------------- |
| length | number | Length of postage |
| width  | number | Width of postage  |
| height | number | Height of postage |

#### **flatRate**

| Field    | Type   | Description                      |
| -------- | ------ | -------------------------------- |
| rateType | string | One of `"ABSOLUTE"`, `"PERCENT"` |
| rate     | number | Shipping rate                    |

#### **ratesTable**

| Field        | Type                   | Description                                                                                         |
| ------------ | ---------------------- | --------------------------------------------------------------------------------------------------- |
| tableBasedOn | string                 | What is this table rate based on. Possible values: `"subtotal"`, `"discountedSubtotal"`, `"weight"` |
| rates        | Object [rates](#rates) | Details of table rate                                                                               |

#### **rates**

| Field      | Type                             | Description                                       |
| ---------- | -------------------------------- | ------------------------------------------------- |
| conditions | Object [conditions](#conditions) | Conditions for this shipping rate in custom table |
| rate       | Object [rate](#rate)             | Table rate details                                |

#### **conditions**

| Field                  | Type   | Description                                |
| ---------------------- | ------ | ------------------------------------------ |
| weightFrom             | number | "Weight from" condition value              |
| weightTo               | number | "Weight to" condition value                |
| subtotalFrom           | number | "Subtotal from" condition value            |
| subtotalTo             | number | "Subtotal to" condition value              |
| discountedSubtotalFrom | number | "Discounted subtotal from" condition value |
| discountedSubtotalTo   | number | "Discounted subtotal from" condition value |

#### **rate**

| Field     | Type   | Description              |
| --------- | ------ | ------------------------ |
| perOrder  | number | Absolute per order rate  |
| percent   | number | Percent per order rate   |
| perItem   | number | Absolute per item rate   |
| perWeight | number | Absolute per weight rate |

#### **businessHours**

| Field | Type             | Description                                                                                             |
| ----- | ---------------- | ------------------------------------------------------------------------------------------------------- |
| MON   | Array time range | Array of time ranges in format `["FROM TIME", "TO TIME"]`. Ex: `['08:30', '13:30'], ['13:30', '19:00']` |
| TUE   | Array time range | Array of time ranges in format `["FROM TIME", "TO TIME"]`. Ex: `['08:30', '13:30'], ['13:30', '19:00']` |
| WED   | Array time range | Array of time ranges in format `["FROM TIME", "TO TIME"]`. Ex: `['08:30', '13:30'], ['13:30', '19:00']` |
| THU   | Array time range | Array of time ranges in format `["FROM TIME", "TO TIME"]`. Ex: `['08:30', '13:30'], ['13:30', '19:00']` |
| FRI   | Array time range | Array of time ranges in format `["FROM TIME", "TO TIME"]`. Ex: `['08:30', '13:30'], ['13:30', '19:00']` |
| SAT   | Array time range | Array of time ranges in format `["FROM TIME", "TO TIME"]`. Ex: `['08:30', '13:30'], ['13:30', '19:00']` |
| SUN   | Array time range | Array of time ranges in format `["FROM TIME", "TO TIME"]`. Ex: `['08:30', '13:30'], ['13:30', '19:00']` |

#### **taxSettings**

*System Settings → Taxes*

| Field                          | Type                  | Description                                                                                                                                                                                                               |
| ------------------------------ | --------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| automaticTaxEnabled            | boolean               | `true` if store taxes are calculated automatically, `else` otherwise. As seen in the *Ecwid Control Panel > Settings > Taxes > Automatic*                                                                                 |
| taxes                          | Array [taxes](#taxes) | Manual tax settings for a store                                                                                                                                                                                           |
| pricesIncludeTax               | boolean               | `true` if store has "gross prices" setting enabled. `false` if store has "net prices" setting enabled.                                                                                                                    |
| taxExemptBusiness              | boolean               | Defines if your business is tax-exempt under § 19 UStG. When `true`, it will display the “Tax exemption § 19 UStG” message to customers to explain the zero VAT rate.                                                     |
| ukVatRegistered                | boolean               | If `true` and order is sent from EU to UK - charges VAT for orders less than GBP 135.                                                                                                                                     |
| euIossEnabled                  | boolean               | If `true` and order is sent to EU - charges VAT for orders less than EUR 150. For Import One-Stop Shop (IOSS).                                                                                                            |
| taxOnShippingCalculationScheme | string                | Shipping tax calculation schemes. Default value: `AUTOMATIC`. Possible values: `AUTOMATIC`, `BASED_ON_PRODUCT_TAXES_PROPORTION_BY_PRICE`, `BASED_ON_PRODUCT_TAXES_PROPORTION_BY_WEIGHT`, `TAXED_SEPARATELY_FROM_PRODUCTS` |

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

*System Settings → Zones*

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

*System Settings → General → Instant site*

| Field          | Type   | Description                                                                          |
| -------------- | ------ | ------------------------------------------------------------------------------------ |
| ecwidSubdomain | string | Store subdomain on ecwid.com domain, e.g. `mysuperstore` in `mysuperstore.ecwid.com` |
| customDomain   | string | Custom Instant site domain, e.g. `www.mysuperstore.com`                              |
| generatedUrl   | string | Instant Site generated URL, e.g. `http://mysuperstore.ecwid.com/`                    |
| storeLogoUrl   | string | Instant Site logo URL                                                                |

#### **legalPagesSettings**

*System Settings → General → Legal Pages*

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

#### **payment**

| Field           | Type                                    | Description                                            |
| --------------- | --------------------------------------- | ------------------------------------------------------ |
| paymentOptions  | Array [paymentOptions](#paymentoptions) | Details about all payment methods set up in that store |
| applePay        | Object [applePay](#applepay)            | Details about Apple Pay setup in that store            |
| applePayOptions | Object [applePay](#applepay)            | Details about payment processors accepting Apple Pay   |

#### **paymentOptions**

| Field                   | Type                                                       | Description                                                                                                                                                                                                                                        |
| ----------------------- | ---------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| id                      | string                                                     | Payment method ID in a store                                                                                                                                                                                                                       |
| enabled                 | boolean                                                    | `true` if payment method is enabled and shown in storefront, `false` otherwise                                                                                                                                                                     |
| configured              | boolean                                                    | Contains the payment method setup status. Read-only for in-built payment methods (where `"appClientId": ""`). Can be set for payment applications and will affect the payment method list in a store dashboard, see <https://ecwid.d.pr/i/FpeCIb>. |
| checkoutTitle           | string                                                     | Payment method title at checkout                                                                                                                                                                                                                   |
| checkoutTitleTranslated | Object [translations](#translations)                       | Available translations for payment option title.                                                                                                                                                                                                   |
| checkoutDescription     | string                                                     | Payment method description at checkout (subtitle)                                                                                                                                                                                                  |
| paymentProcessorId      | string                                                     | Payment processor ID in Ecwid                                                                                                                                                                                                                      |
| paymentProcessorTitle   | string                                                     | Payment processor title. The same as `paymentModule` in order details in REST API                                                                                                                                                                  |
| orderBy                 | number                                                     | Payment method position at checkout and in Ecwid Control Panel. The smaller the number, the higher the position is                                                                                                                                 |
| appClientId             | string                                                     | client\_id value of payment application. `""` if not an application                                                                                                                                                                                |
| paymentSurcharges       | Array [paymentSurcharges](#paymentsurcharges)              | Payment method fee added to the order as a set amount or as a percentage of the order total                                                                                                                                                        |
| instructionsForCustomer | Object [instructionsForCustomer](#instructionsforcustomer) | Customer instructions details                                                                                                                                                                                                                      |
| shippingSettings        | Object [shippingSettings](#shippingsettings)               | Shipping settings of the payment option                                                                                                                                                                                                            |

#### **paymentSurcharges**

| Field | Type   | Description                             |
| ----- | ------ | --------------------------------------- |
| type  | string | Supported values: `ABSOLUTE`, `PERCENT` |
| value | number | Surcharge value                         |

#### **applePay**

| Field            | Type    | Description                                                                        |
| ---------------- | ------- | ---------------------------------------------------------------------------------- |
| enabled          | boolean | `true` if Apple Pay is enabled and shown in storefront, `false` otherwise          |
| available        | boolean | `true` if Stripe payment method is set up, `false` otherwise                       |
| gateway          | string  | Always `"stripe"`                                                                  |
| verificationFile | string  | <https://stripe.com/files/apple-pay/apple-developer-merchantid-domain-association> |

```json
"payment": {
        "paymentOptions": [
            ...
        ],
        "applePay": {
            "enabled": false,
            "available": true,
            "gateway": "stripe",
            "verificationFileUrl": "https://stripe.com/files/apple-pay/apple-developer-merchantid-domain-association"
        }
```

```json
"payment": {
        "paymentOptions": [
            ...
        ],
        "applePay": {
            "enabled": false,
            "available": false
        }
```

#### **instructionsForCustomer**

| Field                  | Type                                 | Description                                         |
| ---------------------- | ------------------------------------ | --------------------------------------------------- |
| instructionsTitle      | string                               | Payment instructions title                          |
| instructions           | string                               | Payment instructions content. Can contain HTML tags |
| instructionsTranslated | Object [translations](#translations) | Available translations for instructions.            |

#### **shippingSettings**

| Field                  | Type             | Description                                                                                                                         |
| ---------------------- | ---------------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| enabledShippingMethods | Array of strings | Contains IDs of shipping methods, if payment method is available for certain shipping methods only ("Payment Per Shipping" feature) |

#### **featureToggles**

| Field   | Type    | Description                                                                                                |
| ------- | ------- | ---------------------------------------------------------------------------------------------------------- |
| name    | string  | Feature name                                                                                               |
| visible | boolean | `true` if feature is shown to merchant in *Ecwid Control Panel > Settings > What's new*. `false` otherwise |
| enabled | boolean | `true` if feature is enabled and active in store                                                           |

#### **designSettings**

| Field                       | Type              | Description                                                                                                                                                      |
| --------------------------- | ----------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| DESIGN\_CONFIG\_FIELD\_NAME | string or boolean | Store design settings as seen in [storefront design customization](ref:customize-appearance). If a specific config field is not provided, it will not be changed |

#### **productFiltersSettings**

| Field               | Type                                               | Description                                                             |
| ------------------- | -------------------------------------------------- | ----------------------------------------------------------------------- |
| enabledInStorefront | boolean                                            | `true` if product filters are enabled in storefront. `false` otherwise. |
| filterSections      | array of objects [filterSections](#filtersections) | Specific product filters                                                |

#### **filterSections**

<table><thead><tr><th width="199.6875">Field</th><th width="199.984375">Type</th><th>Description</th></tr></thead><tbody><tr><td>type</td><td>string</td><td>Type of specific product filter. Possible values: <code>IN_STOCK</code>, <code>ON_SALE</code>, <code>PRICE</code>, <code>CATEGORIES</code>, <code>SEARCH</code>, <code>SKU</code>, <code>OPTION</code>, <code>ATTRIBUTE</code>, <code>LOCATIONS</code>.</td></tr><tr><td>name</td><td>string</td><td>Name of the product field. Works only with <code>OPTION</code> and <code>ATTRIBUTE</code> filter types and is required for them.</td></tr><tr><td>displayComponent</td><td>string</td><td><p>Style of displaying <code>OPTION</code> filters on the storefront. </p><p></p><p>One of: </p><ul><li><code>CHECKBOXES</code> - Default checkboxes style.</li><li><code>BUTTON_GRID</code> - Grid with buttons that better suits product options like "Size".</li></ul></td></tr><tr><td>enabled</td><td>boolean</td><td><code>true</code> if specific product filter is enabled. <code>false</code> otherwise</td></tr></tbody></table>

#### verticalSettings

| Field              | Type             | Description                                           |
| ------------------ | ---------------- | ----------------------------------------------------- |
| vertical           | string           | Primary vertical (e.g., "apparel", "food", "sports"). |
| primarySubvertical | string           | Primary subvertical (e.g., "clothing" for "apparel"). |
| subverticals       | array of strings | List of subverticals.                                 |
| productTypes       | array of strings | List of product types.                                |

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

#### **fbMessengerSettings**

| Field                           | Type    | Description                                     |
| ------------------------------- | ------- | ----------------------------------------------- |
| enabled                         | boolean | `true` if enabled, `false` otherwise            |
| fbMessengerPageId               | string  | Page ID of the connected page on Facebook       |
| fbMessengerThemeColor           | string  | Chat color theme for FB Messenger               |
| fbMessengerMessageUsButtonColor | string  | Color for the FB Messenger button in storefront |

#### **mailchimpSettings**

| Field  | Type   | Description                                                                               |
| ------ | ------ | ----------------------------------------------------------------------------------------- |
| script | string | JS script for the Mailchimp integration, e.g. `"<script id="mcjs">!function...</script>"` |

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

#### **giftCardSettings**

| Field           | Type                        | Description                                                                                            |
| --------------- | --------------------------- | ------------------------------------------------------------------------------------------------------ |
| products        | Array [products](#products) | Basic information of gift card products in a store                                                     |
| displayLocation | string                      | Display location for gift cards on storefront: `"CATALOG_AND_FOOTER"` and `"CATALOG"`. Read only field |

#### **products**

| Field | Type   | Description                      |
| ----- | ------ | -------------------------------- |
| id    | number | Product ID                       |
| name  | string | Gift card product name           |
| url   | string | Gift card product URL in a store |

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

#### **accountBilling**

| Field                                | Type    | Description                                                                                                                  |
| ------------------------------------ | ------- | ---------------------------------------------------------------------------------------------------------------------------- |
| channelid                            | string  | Store channel id info                                                                                                        |
| pricingPlanId                        | string  | Store plan id                                                                                                                |
| pricingPlanName                      | string  | Store plan name                                                                                                              |
| pricingPlanPeriod                    | string  | Store plan billing period                                                                                                    |
| nextRecurringChargeDate              | string  | Next charge date                                                                                                             |
| nextRecurringChargeAmount            | string  | Next charge amount                                                                                                           |
| nextRecurringChargeCurrency          | string  | Charge currency                                                                                                              |
| nextRecurringChargeAmountFormatted   | string  | Next charge in the amount and currency format                                                                                |
| inGracePeriod                        | boolean | `true` if store is on the 'grace' period, `false` otherwise                                                                  |
| willDowngradeAt                      | string  | Date when the 'grace' period will expire in UTC. `null` if an account is not on the 'grace' period                           |
| internalBilling                      | boolean | Store billing. `true` if the store is on the internal billing, `false` otherwise                                             |
| paymentMethod                        | string  | Store billing payment method                                                                                                 |
| billingPageVisibleInCP               | boolean | Store’s ‘Billing and Plans’ page in the Control Panel. `true` if the page is visible in the Control Panel, `false` otherwise |
| itunesSubscriptionAvailableOnChannel | boolean | Store’s interface for iTunes subscription. `true` if the interface is available, `false` otherwise                           |

#### **blackoutDates**

| Field            | Type    | Description                                                         |
| ---------------- | ------- | ------------------------------------------------------------------- |
| fromDate         | string  | Starting date of the period, e.g. `2022-04-28`.                     |
| toDate           | string  | The end date of the period, e.g. `2022-04-30`.                      |
| repeatedAnnually | boolean | Specifies whether the period repeats in the following years or not. |
