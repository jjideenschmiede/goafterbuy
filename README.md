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

```go
// Define products body
body := &goafterbuy.ProductsBody{
    goafterbuy.ProductsRequest{
        goafterbuy.AfterbuyGlobal{
            "PartnerId",
            "PartnerPassword",
            "UserId",
            "UserPassword",
            "GetShopProducts",
            0,
            "DE",
        },
        250,
        0,
        1,
        1,
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

```go
// Define stock body
body := &goafterbuy.StockBody{
    goafterbuy.StockRequest{
        goafterbuy.AfterbuyGlobal{
            "PartnerId",
            "PartnerPassword",
            "UserId",
            "UserPassword",
            "GetStockInfo",
            2,
            "DE",
        },
        goafterbuy.StockProducts{
            goafterbuy.StockProduct{
                ProductID,
            },
        },
    },
}

// Get stock
stock, err := Stock(body)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(stock)
}
```