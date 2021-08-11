# goafterbuy

## Install

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
body := &goafterbuy.ProductsBody{
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
body := &goafterbuy.StockBody{
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

## Add order

If an order is to be returned to the system, then this can be done as follows. Please inform yourself [here](https://xmldoku.afterbuy.de/shopdoku/) how it works.

**Attention! Here the store interface UTF-8 is used.**

```go
// Define order body
body := &goafterbuy.AddOrderBody{
    Action:                 "",
    PartnerId:              "",
    PartnerPass:            "",
    UserId:                 "",
    ItemNo:                 "",
    KUsername:              "",
    KSalutation:            "",
    KFirstName:             "",
    KLastName:              "",
    KStreet:                "",
    KZip:                   "",
    KLocation:              "",
    KCountry:               "",
    KPhone:                 "",
    KEmail:                 "",
    ShippingGroup:          "",
    ShippingMethod:         "",
    ShippingCosts:          "",
    NoShippingCalc:         "",
    SoldCurrency:           "",
    NoFeedback:             "",
    DeliveryAddress:        "",
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
