package goafterbuy_test

import (
	"fmt"
	"github.com/jjideenschmiede/goafterbuy"
	"testing"
)

// TestAddOrder is to test the oder function
func TestAddOrder(t *testing.T) {

	// Define variables for request
	partnerId := ""
	partnerPass := ""
	userId := ""

	// Define order body
	body := goafterbuy.AddOrderBody{
		Action:                  "new",
		PartnerId:               partnerId,
		PartnerPass:             partnerPass,
		UserId:                  userId,
		ItemNo:                  "1",
		KUsername:               "Testing",
		KSalutation:             "Herr",
		KCompany:                "Test",
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
		ShippingMethod:          "Versandkostenfrei DHL",
		PaymentMethodsSurcharge: "",
		ShippingCosts:           "",
		NoShippingCalc:          "1",
		PaymentMethod:           "Kreditkarte",
		SoldCurrency:            "EUR",
		VMemo:                   "",
		OrderInfo1:              "",
		OrderInfo2:              "",
		OrderInfo3:              "",
		NoFeedback:              "0",
		SetPay:                  "1",
		CustomerIdentification:  "1",
		ArticleIdentification:   "0",
		StockType:               "shop",
		SetPayDate:              "",
		BuyDate:                 "",
		B1:                      "Abschicken",
		VId:                     "",
		CheckVId:                "0",
		Items: []goafterbuy.AddOrderBodyItem{
			{
				ArticleNo:        "80286342",
				AlternArticleNo1: "80286342",
				AlternArticleNo2: "#25457",
				ArticleName:      "Artikel 1",
				ArticleEPrice:    "26,95",
				ArticleVat:       "19",
				ArticleQuantity:  "1",
				ArticleLink:      "",
				Attribute:        "",
				ArticleMasterId:  "80286342",
				ArticleTag1:      "",
				ArticleTag2:      "",
				ArticleTag3:      "",
				ArticleTag4:      "",
				ArticleTag5:      "",
			},
		},
	}

	// Add order
	_, err := goafterbuy.AddOrder(body)
	if err != nil {
		fmt.Println(err)
	}

}
