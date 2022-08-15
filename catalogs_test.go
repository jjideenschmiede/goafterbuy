package goafterbuy_test

import (
	"github.com/jjideenschmiede/goafterbuy"
	"testing"
)

// TestCatalogs is to test the catalogs function
func TestCatalogs(t *testing.T) {

	// Define variables for request
	partnerToken := ""
	accountToken := ""

	// Define catalogs body
	body := goafterbuy.CatalogsBody{
		Request: goafterbuy.CatalogsRequest{
			AfterbuyGlobal: goafterbuy.AfterbuyGlobal{
				PartnerToken:  partnerToken,
				AccountToken:  accountToken,
				CallName:      "GetShopCatalogs",
				ErrorLanguage: "DE",
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
		t.Fatal(err)
	}

	// Check the results
	var results []string
	for _, value := range catalogs.Result.Catalogs.Catalog {
		results = append(results, value.Name)
	}

	// Print output
	t.Logf("The catalogs were read. There are \"%d\" catalogs. Here you can see the names of the catalogs read out: %v.", len(results), results)

}
