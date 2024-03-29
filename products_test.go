package goafterbuy_test

import (
	"github.com/jjideenschmiede/goafterbuy"
	"testing"
)

// TestProducts is to test the products function
func TestProducts(t *testing.T) {

	// Define variables for request
	partnerToken := ""
	accountToken := ""

	// Define products body
	body := goafterbuy.ProductsBody{
		Request: goafterbuy.ProductsRequest{
			AfterbuyGlobal: goafterbuy.AfterbuyGlobal{
				PartnerToken:  partnerToken,
				AccountToken:  accountToken,
				CallName:      "GetShopProducts",
				ErrorLanguage: "DE",
			},
			DataFilter:                     nil,
			MaxShopItems:                   250,
			SuppressBaseProductRelatedData: 0,
			PaginationEnabled:              1,
			PageNumber:                     0,
		},
	}

	// Get products
	products, err := goafterbuy.Products(body)
	if err != nil {
		t.Fatal(err)
	}

	// Check the results
	var results []string
	for _, value := range products.Result.Products.Product {
		results = append(results, value.Name)
	}

	// Print output
	t.Logf("The products were read. There are \"%d\" products. Here you can see the names of the products read out: %v.", len(results), results)

}

// TestUpdateProducts is to test the products update function
func TestUpdateProducts(t *testing.T) {

	// Define variables for request
	partnerToken := ""
	accountToken := ""

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

	// Update products
	product, err := goafterbuy.UpdateProducts(body)
	if err != nil {
		t.Fatal(err)
	}

	// Print output
	t.Logf("The product were updated. Here is the result: %v.", product)

}
