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
            CallName:        "",
            DetailLevel:     0,
            ErrorLanguage:   "DE",
        },
        MaxShopItems:                   100,
        SuppressBaseProductRelatedData: 0,
        PaginationEnabled:              1,
        PageNumber:                     0,
    },
}

// Get afterbuy products
products, err := goafterbuy.Products(body)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(products)
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
            CallName:        "",
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
    },
    MaxCatalogs: 200,
    DataFilter: goafterbuy.CatalogsRequestDataFilter{
        Filter: goafterbuy.CatalogsRequestFilter{
            FilterName:   "RangeID",
            FilterValues: goafterbuy.CatalogsRequestFilterValues{
                ValueFrom: 0,
                ValueTo: 0,
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
  body := SoldItemsBody{
      Request: SoldItemsRequest{
          AfterbuyGlobal: AfterbuyGlobal{
              PartnerId:       0,
              PartnerPassword: "",
              UserId:          "",
              UserPassword:    "",
              CallName:        "GetSoldItems",
              DetailLevel:     0,
              ErrorLanguage:   "DE",
          },
          DataFilter: SoldItemsDataFilter{
              Filter: []SoldItemsFilter{},
          },
          MaxSoldItems: 100,
      },
  }

  // Add filter
  body.Request.DataFilter.Filter = append(body.Request.DataFilter.Filter, SoldItemsFilter{
      FilterName: "DateFilter",
      FilterValues: SoldItemsFilterValues{
          DateFrom:    "01.11.2021 00:00:00",
          DateTo:      "01.12.2021 23:59:59",
          FilterValue: "PayDate",
      },
  }, SoldItemsFilter{
      FilterName: "RangeID",
      FilterValues: SoldItemsFilterValues{
          ValueFrom: "525914526",
      },
  })

  // Get afterbuy sold items
  soldItems, err := SoldItems(body)
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
                AdditionalInfo: "Kleiner Pimmeldachs",
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
body := goafterbuy.AddOrderBody{
    Action:                 "",
    PartnerId:              "",
    PartnerPass:            "",
    UserId:                 "",
    ItemNo:                 "",
    KUsername:              "",
    KSalutation:            "",
    KCompany:               "",
    KFirstName:             "",
    KLastName:              "",
    KStreet:                "",
    KStreet2:               "",
    KZip:                   "",
    KLocation:              "",
    KState:                 "",
    KCountry:               "",
    KPhone:                 "",
    KFax:                   "",
    KEmail:                 "",
    DeliveryAddress:        "",
    KLSalutation:           "",
    KLCompany:              "",
    KLFirstName:            "",
    KLLastName:             "",
    KLStreet:               "",
    KLStreet2:              "",
    KLZip:                  "",
    KLLocation:             "",
    KLState:                "",
    KLCountry:              "",
    KLPhone:                "",
    ShippingGroup:          "",
    ShippingMethod:         "",
    ShippingCosts:          "",
    NoShippingCalc:         "",
    SoldCurrency:           "",
    NoFeedback:             "",
    SetPay:                 "",
    CustomerIdentification: "",
    ArticleIdentification:  "",
    StockType:              "",
    B1:                     "",
    Items:                  []goafterbuy.AddOrderBodyItem{},
}

body.Items = append(body.Items, goafterbuy.AddOrderBodyItem{
    ArticleNo:       "",
    AlternArticleNo: "",
    ArticleName:     "",
    ArticleEPrice:   "",
    ArticleVat:      "",
    ArticleQuantity: "",
    ArticleMasterId: "",
})

// Set afterbuy order
order, err := goafterbuy.AddOrder(body)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(order)
}
```
