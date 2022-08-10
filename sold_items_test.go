package goafterbuy_test

import (
	"github.com/jjideenschmiede/goafterbuy"
	"testing"
)

// TestSoldItems is to test the sold items function
func TestSoldItems(t *testing.T) {

	// Define variables for request
	partnerToken := ""
	accountToken := ""

	// Define sold item body
	body := goafterbuy.SoldItemsBody{
		Request: goafterbuy.SoldItemsRequest{
			AfterbuyGlobal: goafterbuy.AfterbuyGlobal{
				PartnerToken:  partnerToken,
				AccountToken:  accountToken,
				CallName:      "GetSoldItems",
				ErrorLanguage: "DE",
			},
			DataFilter: goafterbuy.SoldItemsDataFilter{
				Filter: []goafterbuy.SoldItemsFilter{
					{
						FilterName: "DateFilter",
						FilterValues: goafterbuy.SoldItemsFilterValues{
							DateFrom:    "01.01.2022 00:00:00",
							DateTo:      "01.12.2022 23:59:59",
							FilterValue: "PayDate",
						},
					},
				},
			},
			MaxSoldItems:      100,
			ReturnHiddenItems: 1,
		},
	}

	// Check sold items
	soldItems, err := goafterbuy.SoldItems(body)
	if err != nil {
		t.Fatal(err)
	}

	// Check the results
	var results []int
	for _, value := range soldItems.Result.Orders.Order {
		results = append(results, value.OrderId)
	}

	// Print output
	t.Logf("The sold items were read. There are \"%d\" sold items. Here you can see the names of the sold items read out: %v.", len(results), results)

}

// TestSoldItem is to test the sold item function
func TestSoldItem(t *testing.T) {

	// Define variables for request
	partnerToken := ""
	accountToken := ""

	// Define sold item body
	body := goafterbuy.SoldItemBody{
		Request: goafterbuy.SoldItemRequest{
			AfterbuyGlobal: goafterbuy.AfterbuyGlobal{
				PartnerToken:  partnerToken,
				AccountToken:  accountToken,
				CallName:      "GetSoldItems",
				ErrorLanguage: "DE",
			},
			DataFilter: goafterbuy.SoldItemDataFilter{
				Filter: goafterbuy.SoldItemFilter{
					FilterName: "OrderID",
					FilterValues: goafterbuy.SoldItemFilterValues{
						FilterValue: 543236883,
					},
				},
			},
			ReturnHiddenItems: 1,
		},
	}

	// Check sold item
	soldItem, err := goafterbuy.SoldItem(body)
	if err != nil {
		t.Fatal(err)
	}

	// Print output
	t.Logf("The sold item were read. Here you can see the of the sold item read out: %d.", soldItem.Result.Orders.Order.OrderId)

}

// TestUpdateSoldItem is to test the update sold item function
func TestUpdateSoldItem(t *testing.T) {

	// Define variables for request
	partnerToken := ""
	accountToken := ""

	// Define update sold item body
	body := goafterbuy.UpdateSoldItemBody{
		Request: goafterbuy.UpdateSoldItemBodyRequest{
			AfterbuyGlobal: goafterbuy.AfterbuyGlobal{
				PartnerToken:  partnerToken,
				AccountToken:  accountToken,
				CallName:      "UpdateSoldItems",
				ErrorLanguage: "DE",
			},
			Orders: goafterbuy.UpdateSoldItemBodyOrders{
				Order: goafterbuy.UpdateSoldItemBodyOrder{
					OrderId:        542554727,
					AdditionalInfo: "DE9034274324",
				},
			},
		},
	}

	// Update sold item
	_, err := goafterbuy.UpdateSoldItem(body)
	if err != nil {
		t.Fatal(err)
	}

}
