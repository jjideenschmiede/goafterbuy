//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of goafterbuy.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package goafterbuy

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// AddOrderBody is to structure the order body data
type AddOrderBody struct {
	Action                 string
	PartnerId              string
	PartnerPass            string
	UserId                 string
	ItemNo                 string
	KUsername              string
	KSalutation            string
	KFirstName             string
	KLastName              string
	KStreet                string
	KZip                   string
	KLocation              string
	KCountry               string
	KPhone                 string
	KEmail                 string
	ShippingGroup          string
	ShippingMethod         string
	ShippingCosts          string
	NoShippingCalc         string
	SoldCurrency           string
	NoFeedback             string
	DeliveryAddress        string
	SetPay                 string
	CustomerIdentification string
	ArticleIdentification  string
	StockType              string
	B1                     string
	Items                  []AddOrderBodyItem
}

type AddOrderBodyItem struct {
	ArticleNo       string
	AlternArticleNo string
	ArticleName     string
	ArticleEPrice   string
	ArticleVat      string
	ArticleQuantity string
	ArticleMasterId string
}

// AddOrderReturn is to decode xml data
type AddOrderReturn struct {
	XmlName xml.Name           `xml:"result"`
	Success int                `xml:"success"`
	Action  string             `xml:"Action"`
	Data    AddOrderReturnData `xml:"data"`
}

type AddOrderReturnData struct {
	XmlName    xml.Name `xml:"data"`
	AId        int      `xml:"AID"`
	UId        string   `xml:"UID"`
	KundenNr   int      `xml:"KundenNr"`
	EKundenNr  int      `xml:"EKundenNr"`
	CouponUsed int      `xml:"CouponUsed"`
}

// AddOrder is to add a new order via shop interface
func AddOrder(body *AddOrderBody) (*AddOrderReturn, error) {

	// Define request
	request, err := http.NewRequest(method, shopInterfaceBaseUrl, nil)
	if err != nil {
		return nil, err
	}

	// Set url parameter
	parameter := request.URL.Query()

	parameter.Add("Action", body.Action)
	parameter.Add("Partnerid", body.PartnerId)
	parameter.Add("PartnerPass", body.PartnerPass)
	parameter.Add("UserID", body.UserId)
	parameter.Add("PosAnz", body.ItemNo)
	parameter.Add("Kbenutzername", body.KUsername)
	parameter.Add("Kanrede", body.KSalutation)
	parameter.Add("KVorname", body.KFirstName)
	parameter.Add("KNachname", body.KLastName)
	parameter.Add("KStrasse", body.KStreet)
	parameter.Add("KPLZ", body.KZip)
	parameter.Add("KOrt", body.KLocation)
	parameter.Add("KLand", body.KCountry)
	parameter.Add("KTelefon", body.KPhone)
	parameter.Add("Kemail", body.KEmail)
	parameter.Add("Versandgruppe", body.ShippingGroup)
	parameter.Add("Versandart", body.ShippingMethod)
	parameter.Add("Versandkosten", body.ShippingCosts)
	parameter.Add("NoVersandCalc", body.NoShippingCalc)
	parameter.Add("SoldCurrency", body.SoldCurrency)
	parameter.Add("NoFeedback", body.NoFeedback)
	parameter.Add("Lieferanschrift", body.DeliveryAddress)
	parameter.Add("SetPay", body.SetPay)
	parameter.Add("Kundenerkennung", body.CustomerIdentification)
	parameter.Add("Artikelerkennung", body.ArticleIdentification)
	parameter.Add("Bestandart", body.StockType)
	parameter.Add("B1", body.B1)

	// Add items
	for index, value := range body.Items {

		// Add url parameter
		parameter.Add(fmt.Sprintf("Artikelnr_%d", index+1), value.ArticleNo)
		parameter.Add(fmt.Sprintf("AlternArtikelNr1_%d", index+1), value.AlternArticleNo)
		parameter.Add(fmt.Sprintf("Artikelname_%d", index+1), value.ArticleName)
		parameter.Add(fmt.Sprintf("ArtikelEpreis_%d", index+1), value.ArticleEPrice)
		parameter.Add(fmt.Sprintf("ArtikelMwSt_%d", index+1), value.ArticleVat)
		parameter.Add(fmt.Sprintf("ArtikelMenge_%d", index+1), value.ArticleQuantity)
		parameter.Add(fmt.Sprintf("ArtikelStammID_%d", index+1), value.ArticleMasterId)

	}

	// Set new url
	request.URL.RawQuery = parameter.Encode()

	// Config new request
	c := Config{request, nil}

	// Send new request
	response, err := c.Send()
	if err != nil {
		return nil, err
	}

	// Close request body
	defer response.Body.Close()

	// Decode data
	var decode AddOrderReturn

	err = xml.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return nil, err
	}

	// Return data
	return &decode, nil

}
