package goafterbuy_test

import (
	"github.com/jjideenschmiede/goafterbuy"
	"testing"
)

// TestShippingServices is to test the shipping services function
func TestShippingServices(t *testing.T) {

	// Define variables for request
	partnerToken := ""
	accountToken := ""

	// Define products body
	body := goafterbuy.ShippingServicesBody{
		Request: goafterbuy.ShippingServicesRequest{
			AfterbuyGlobal: goafterbuy.AfterbuyGlobal{
				PartnerToken:  partnerToken,
				AccountToken:  accountToken,
				CallName:      "GetShippingServices",
				ErrorLanguage: "DE",
			},
		},
	}

	// Get shipping services
	shippingServices, err := goafterbuy.ShippingServices(body)
	if err != nil {
		t.Fatal(err)
	}

	// Check the results
	var results []string
	for _, value := range shippingServices.Result.ShippingServices.ShippingService {
		results = append(results, value.Name)
	}

	// Print output
	t.Logf("The shipping services were read. There are \"%d\" shipping services. Here you can see the names of the shipping services read out: %v.", len(results), results)

}
