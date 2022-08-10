package goafterbuy_test

import (
	"github.com/jjideenschmiede/goafterbuy"
	"testing"
)

// TestStock is to test the stock function
func TestStock(t *testing.T) {

	// Define variables for request
	partnerToken := ""
	accountToken := ""

	// Define stock body
	body := goafterbuy.StockBody{
		Request: goafterbuy.StockRequest{
			AfterbuyGlobal: goafterbuy.AfterbuyGlobal{
				PartnerToken:  partnerToken,
				AccountToken:  accountToken,
				CallName:      "GetShopProducts",
				DetailLevel:   2,
				ErrorLanguage: "DE",
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
		t.Fatal(err)
	}

	// Check the results
	var results []string
	for _, value := range stock.Result.Products.Product {
		results = append(results, value.Name)
	}

	// Print output
	t.Logf("The stock were read. There are \"%d\" products. Here you can see the names of the products read out: %v.", len(results), results)

}
