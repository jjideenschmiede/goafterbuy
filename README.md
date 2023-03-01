# goafterbuy

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jjideenschmiede/goafterbuy.svg)](https://golang.org/) [![Go](https://github.com/jjideenschmiede/goafterbuy/actions/workflows/go.yml/badge.svg)](https://github.com/jjideenschmiede/goafterbuy/actions/workflows/go.yml)
 [![Go Report Card](https://goreportcard.com/badge/github.com/jjideenschmiede/goafterbuy)](https://goreportcard.com/report/github.com/jjideenschmiede/goafterbuy) [![Go Reference](https://pkg.go.dev/badge/github.com/jjideenschmiede/goafterbuy.svg)](https://pkg.go.dev/github.com/jjideenschmiede/goafterbuy) ![Lines of code](https://img.shields.io/tokei/lines/github/jjideenschmiede/goafterbuy) [![Developed with <3](https://img.shields.io/badge/Developed%20with-%3C3-19ABFF)](https://jj-dev.de/)

```console
go get github.com/jjideenschmiede/goafterbuy
```

## How to use?

## Get products

If you want to read the afterbuy products, you can do it as follows. Some parameters from afterbuy are needed, like the PartnerId, PartnerPassword, UserId & UserPassword.

If you want to output multiple pages, then you need to adjust the details in goafterbuy.ProductsRequest. The first request will output values such as maximum number of pages. At this you can execute the functions with adjusted parameters.

[Here](https://xmldoku.afterbuy.de/dokued/) you can find an example from Afterbuy.

**Attention! Here the XML interface is used.**

```go
// Define products body
body := goafterbuy.ProductsBody{
    Request: goafterbuy.ProductsRequest{
        AfterbuyGlobal: goafterbuy.AfterbuyGlobal{
            PartnerId:       0,
            PartnerPassword: "",
            UserId:          "",
            UserPassword:    "",
            CallName:        "GetShopProducts",
            DetailLevel:     0,
            ErrorLanguage:   "DE",
        },
        DataFilter: &goafterbuy.ProductsRequestDataFilter{},
        MaxShopItems:                   100,
        SuppressBaseProductRelatedData: 0,
        PaginationEnabled:              1,
        PageNumber:                     0,
    },
}

// Add filter value for product id
body.Request.DataFilter.Filter = append(body.Request.DataFilter.Filter, goafterbuy.ProductsRequestFilter{
    FilterName: "ProductID",
	FilterValues: goafterbuy.ProductsRequestFilterValues{
        FilterValue: []string{"79341966", "79341972"},
    },
})

// Add filter value for level
body.Request.DataFilter.Filter = append(body.Request.DataFilter.Filter, goafterbuy.ProductsRequestFilter{
    FilterName: "Level",
    FilterValues: goafterbuy.ProductsRequestFilterValues{
        LevelFrom: 0,
        LevelTo:   99,
    },
})

// Get afterbuy products
products, err := goafterbuy.Products(body)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(products)
}
```

**For all requests described in this documentation, you can also verify with the PartnerToken and the AccountToken. The whole thing is then also deposited in goafterbuy.AfterbuyGlobal and looks like this:**

```go
// Define products body
body := goafterbuy.ProductsBody{
    Request: goafterbuy.ProductsRequest{
        AfterbuyGlobal: goafterbuy.AfterbuyGlobal{
            PartnerToken:  "",
            AccountToken:  "",
            CallName:      "GetShopProducts",
            DetailLevel:   0,
            ErrorLanguage: "DE",
        },
		
        // Rest of the request body
		
    },
}
```

## Update product

If you want to update a product, then you can use the following function.

```go
// Define products body
body := goafterbuy.UpdateProductsBody{
	Request: goafterbuy.UpdateProductsRequest{
		AfterbuyGlobal: goafterbuy.AfterbuyGlobal{
			PartnerToken:  partnerToken,
			AccountToken:  accountToken,
			CallName:      "UpdateShopProducts",
			ErrorLanguage: "DE",
		},
		Products: goafterbuy.UpdateProductsProducts{
			Product: []goafterbuy.UpdateProductsProduct{
				{
					ProductIdent: goafterbuy.UpdateProductsProductIdent{
						ProductInsert: 0,
						ProductID:     81865201,
					},
					Quantity: 187,
				},
			},
		},
	},
}

// Update product
product, err := goafterbuy.UpdateProducts(body)
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println(product)
}
```

## Get stock

If you want to read out a stock, then you can do this here via the ID of the product. Please note that you need the same data here as when reading out the products.

[Here](https://xmldoku.afterbuy.de/dokued/) you can find an example from Afterbuy.

**Attention! Here the XML interface is used.**

```go
// Define stock body
body := goafterbuy.StockBody{
    Request: goafterbuy.StockRequest{
        AfterbuyGlobal: goafterbuy.AfterbuyGlobal{
            PartnerId:       0,
            PartnerPassword: "",
            UserId:          "",
            UserPassword:    "",
            CallName:        "GetShopProducts",
            DetailLevel:     2,
            ErrorLanguage:   "DE",
        },
        Products: goafterbuy.StockProducts{
            Product: goafterbuy.StockProduct{
                ProductId: 0,
            },
        },
    },
}

// Get stock
stock, err := goafterbuy.Stock(body)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(stock)
}
```

## Get catalogs

To enable them to read the catalogs, you can use the following function.

**Attention! Here the XML interface is used.**

```go
// Define catalogs body
body := goafterbuy.CatalogsBody{
    Request: goafterbuy.CatalogsRequest{
        AfterbuyGlobal: goafterbuy.AfterbuyGlobal{
            PartnerId:       0,
            PartnerPassword: "",
            UserId:          "",
            UserPassword:    "",
            CallName:        "GetShopCatalogs",
            DetailLevel:     0,
            ErrorLanguage:   "DE",
        },
        MaxCatalogs: 200,
        DataFilter: &goafterbuy.CatalogsRequestDataFilter{
            Filter: []goafterbuy.CatalogsRequestFilter{
                {
                    FilterName: "RangeID",
                    FilterValues: goafterbuy.CatalogsRequestFilterValues{
                        ValueFrom: 0, 
                        ValueTo:   0,
                    },
                },
                {
                    FilterName: "Level",
                    FilterValues: goafterbuy.CatalogsRequestFilterValues{
                        FilterValue: []string{"0", "1", "2"},
                    },
                },
            },
        },
    },
}

// Get catalogs
catalogs, err := goafterbuy.Catalogs(body)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(catalogs)
}
```

## Get shipping services

To get the shipping services from afterbuy.

```go
// Define sold item request body
body := goafterbuy.ShippingServicesBody{
    Request: goafterbuy.ShippingServicesRequest{
        AfterbuyGlobal: goafterbuy.AfterbuyGlobal{
            PartnerId:       0,
            PartnerPassword: "",
            UserId:          "",
            UserPassword:    "",
            CallName:        "GetShippingServices",
            ErrorLanguage:   "DE",
        },
    },
}

// Get afterbuy shipping services
shippingServices, err := goafterbuy.ShippingServices(body)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(shippingServices)
}
```

## Get sold items

You can use this function to read out all purchase orders in a specific period.

```go
// Define sold items body
body := goafterbuy.SoldItemsBody{
    Request: goafterbuy.SoldItemsRequest{
        AfterbuyGlobal: goafterbuy.AfterbuyGlobal{
            PartnerId:       0,
            PartnerPassword: "",
            UserId:          "",
            UserPassword:    "",
            CallName:        "GetSoldItems",
            DetailLevel:     0,
            ErrorLanguage:   "DE",
        },
        DataFilter: goafterbuy.SoldItemsDataFilter{
            Filter: []goafterbuy.SoldItemsFilter{},
        },
        MaxSoldItems: 100,
        ReturnHiddenItems: 1,
    },
}

// Add filter
body.Request.DataFilter.Filter = append(body.Request.DataFilter.Filter, goafterbuy.SoldItemsFilter{
    FilterName: "DateFilter",
    FilterValues: goafterbuy.SoldItemsFilterValues{
        DateFrom:    "01.11.2021 00:00:00",
        DateTo:      "01.12.2021 23:59:59",
        FilterValue: "PayDate",
    },
}, goafterbuy.SoldItemsFilter{
    FilterName: "RangeID",
    FilterValues: goafterbuy.SoldItemsFilterValues{
        ValueFrom: 525914526,
    },
})

// Get afterbuy sold items
soldItems, err := goafterbuy.SoldItems(body)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(soldItems)
}
```

## Get sold item

If you want to read out an order based on the Id, then you can use the following function for this.

```go
// Define sold item request body
body := goafterbuy.SoldItemBody{
    Request: goafterbuy.SoldItemRequest{
        AfterbuyGlobal: goafterbuy.AfterbuyGlobal{
            PartnerId:       0,
            PartnerPassword: "",
            UserId:          "",
            UserPassword:    "",
            CallName:        "GetSoldItems",
            ErrorLanguage:   "DE",
        },
        DataFilter: goafterbuy.DataFilter{
            Filter: goafterbuy.DataFilterFilter{
                FilterName: "OrderID",
                FilterValues: goafterbuy.DataFilterFilterValues{
                    FilterValue: 558996468,
                },
            },
        },
        ReturnHiddenItems: 1,
    },
}

// Get sold item
soldItem, err := goafterbuy.SoldItem(body)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(soldItem.Result.Orders.Order.AdditionalInfo)
}
```

## Update sold item

If you want to update an order, you can do this using the following function. A palette of attributes can be sent with the order.

```go
// Define update sold item request body
body := goafterbuy.UpdateSoldItemBody{
    goafterbuy.UpdateSoldItemBodyRequest{
        AfterbuyGlobal: goafterbuy.AfterbuyGlobal{
            PartnerId:       0,
            PartnerPassword: "",
            UserId:          "",
            UserPassword:    "",
            CallName:        "UpdateSoldItems",
            ErrorLanguage:   "DE",
        },
        Orders: goafterbuy.UpdateSoldItemBodyOrders{
            goafterbuy.UpdateSoldItemBodyOrder{
                OrderId:        523371348,
                ItemId:         523371348,
                AdditionalInfo: "DE9034274324",
            },
        },
    },
}

// Get sold item
soldItem, err := goafterbuy.UpdateSoldItem(body)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(soldItem)
}
```

## Add order

If an order is to be returned to the system, then this can be done as follows. Please inform yourself [here](https://xmldoku.afterbuy.de/shopdoku/) how it works.

**Attention! Here the store interface UTF-8 is used.**

```go
// Define order body
body := AddOrderBody{
    Action:                  "new",
    PartnerId:               "",
    PartnerPass:             "",
    UserId:                  "",
    ItemNo:                  "1",
    KUsername:               "WooCommerce",
    KSalutation:             "Herr",
    KCompany:                "J&J Ideenschmiede GmbH",
    KFirstName:              "Jonas",
    KLastName:               "Kwiedor",
    KStreet:                 "Mercatorstraße 32a",
    KStreet2:                "",
    KZip:                    "21502",
    KLocation:               "Geesthacht",
    KState:                  "SH",
    KCountry:                "Deutschland",
    KPhone:                  "",
    KFax:                    "",
    KEmail:                  "jonas.kwiedor@jj-ideenschmiede.de",
    DeliveryAddress:         "0",
    KLSalutation:            "",
    KLCompany:               "",
    KLFirstName:             "",
    KLLastName:              "",
    KLStreet:                "",
    KLStreet2:               "",
    KLZip:                   "",
    KLLocation:              "",
    KLState:                 "",
    KLCountry:               "",
    KLPhone:                 "",
    Comment:                 "Das ist ein Kommentar.",
    ShippingGroup:           "",
    ShippingMethod:          "",
    PaymentMethodsSurcharge: "",
    ShippingCosts:           "0",
    NoShippingCalc:          "1",
    PaymentMethod:           "PayPal",
    SoldCurrency:            "EUR",
    VMemo:                   "Das ist eine Notiz",
    OrderInfo1:              "1234",
    OrderInfo2:              "5678",
    OrderInfo3:              "90",
    NoFeedback:              "0",
    SetPay:                  "1",
    CustomerIdentification:  "1",
    ArticleIdentification:   "0",
    StockType:               "shop",
    B1:                      "Abschicken",
    Items:                   []AddOrderBodyItem{},
}

body.Items = append(body.Items, goafterbuy.AddOrderBodyItem{
    ArticleNo:        "2131424124",
    AlternArticleNo1: "2131424124",
    AlternArticleNo2: "2131424124",
    ArticleName:      "Testartikel",
    ArticleEPrice:    "2,99",
    ArticleVat:       "0,99",
    ArticleQuantity:  "19",
    ArticleLink:      "http://jj-dev.de",
    Attribute:        "Größe:XL|Farbe:Blau",
    ArticleMasterId:  "2131424124",
    ArticleTag1:      "Stornierungsanfrage",
    ArticleTag2:      "",
    ArticleTag3:      "",
    ArticleTag4:      "",
    ArticleTag5:      "",
})

// Set afterbuy order
order, err := goafterbuy.AddOrder(body)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(order)
}
```
