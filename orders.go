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
	Action                  string
	PartnerId               string
	PartnerPass             string
	UserId                  string
	ItemNo                  string
	KUsername               string
	KSalutation             string
	KCompany                string
	KFirstName              string
	KLastName               string
	KStreet                 string
	KStreet2                string
	KZip                    string
	KLocation               string
	KState                  string
	KCountry                string
	KPhone                  string
	KFax                    string
	KEmail                  string
	DeliveryAddress         string
	KLSalutation            string
	KLCompany               string
	KLFirstName             string
	KLLastName              string
	KLStreet                string
	KLStreet2               string
	KLZip                   string
	KLLocation              string
	KLState                 string
	KLCountry               string
	KLPhone                 string
	Comment                 string
	ShippingGroup           string
	ShippingMethod          string
	PaymentMethodsSurcharge string
	ShippingCosts           string
	NoShippingCalc          string
	PaymentMethod           string
	SoldCurrency            string
	VMemo                   string
	OrderInfo1              string
	OrderInfo2              string
	OrderInfo3              string
	NoFeedback              string
	SetPay                  string
	CustomerIdentification  string
	ArticleIdentification   string
	StockType               string
	B1                      string
	Items                   []AddOrderBodyItem
}

type AddOrderBodyItem struct {
	ArticleNo        string
	AlternArticleNo1 string
	AlternArticleNo2 string
	ArticleName      string
	ArticleEPrice    string
	ArticleVat       string
	ArticleQuantity  string
	ArticleLink      string
	ArticleMasterId  string
}

// AddOrderReturn is to decode xml data
type AddOrderReturn struct {
	XmlName   xml.Name                  `xml:"result"`
	Success   int                       `xml:"success"`
	Action    string                    `xml:"Action"`
	Data      AddOrderReturnData        `xml:"data"`
	ErrorList []AddOrderReturnErrorList `xml:"errorlist"`
}

type AddOrderReturnData struct {
	XmlName    xml.Name `xml:"data"`
	AId        int      `xml:"AID"`
	UId        string   `xml:"UID"`
	KundenNr   int      `xml:"KundenNr"`
	EKundenNr  int      `xml:"EKundenNr"`
	CouponUsed int      `xml:"CouponUsed"`
}

type AddOrderReturnErrorList struct {
	Error string `xml:"error"`
}

// AddOrder is to add a new order via shop interface
func AddOrder(body AddOrderBody) (AddOrderReturn, error) {

	// Define request
	request, err := http.NewRequest(method, shopInterfaceBaseUrl, nil)
	if err != nil {
		return AddOrderReturn{}, err
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
	parameter.Add("KFirma", body.KCompany)
	parameter.Add("KVorname", body.KFirstName)
	parameter.Add("KNachname", body.KLastName)
	parameter.Add("KStrasse", body.KStreet)
	parameter.Add("KStrasse2", body.KStreet2)
	parameter.Add("KPLZ", body.KZip)
	parameter.Add("KOrt", body.KLocation)
	parameter.Add("KBundesland", body.KState)
	parameter.Add("KLand", body.KCountry)
	parameter.Add("KTelefon", body.KPhone)
	parameter.Add("Kfax", body.KFax)
	parameter.Add("Kemail", body.KEmail)
	parameter.Add("Lieferanschrift", body.DeliveryAddress)
	parameter.Add("KLanrede", body.KLSalutation)
	parameter.Add("KLFirma", body.KLCompany)
	parameter.Add("KLVorname", body.KLFirstName)
	parameter.Add("KLNachname", body.KLLastName)
	parameter.Add("KLStrasse", body.KLStreet)
	parameter.Add("KLStrasse2", body.KLStreet2)
	parameter.Add("KLPLZ", body.KLZip)
	parameter.Add("KLOrt", body.KLLocation)
	parameter.Add("KLBundesland", body.KLState)
	parameter.Add("KLLand", body.KLCountry)
	parameter.Add("KLTelefon", body.KLPhone)
	parameter.Add("Kommentar", body.Comment)
	parameter.Add("Versandgruppe", body.ShippingGroup)
	parameter.Add("Versandart", body.ShippingMethod)
	parameter.Add("ZahlartenAufschlag", body.PaymentMethodsSurcharge)
	parameter.Add("Versandkosten", body.ShippingCosts)
	parameter.Add("NoVersandCalc", body.NoShippingCalc)
	parameter.Add("Zahlart", body.PaymentMethod)
	parameter.Add("SoldCurrency", body.SoldCurrency)
	parameter.Add("VMemo", body.VMemo)
	parameter.Add("OrderInfo1", body.OrderInfo1)
	parameter.Add("OrderInfo2", body.OrderInfo2)
	parameter.Add("OrderInfo3", body.OrderInfo3)
	parameter.Add("NoFeedback", body.NoFeedback)
	parameter.Add("SetPay", body.SetPay)
	parameter.Add("Kundenerkennung", body.CustomerIdentification)
	parameter.Add("Artikelerkennung", body.ArticleIdentification)
	parameter.Add("Bestandart", body.StockType)
	parameter.Add("B1", body.B1)

	// Add items
	for index, value := range body.Items {

		// Add url parameter
		parameter.Add(fmt.Sprintf("Artikelnr_%d", index+1), value.ArticleNo)
		parameter.Add(fmt.Sprintf("AlternArtikelNr1_%d", index+1), value.AlternArticleNo1)
		parameter.Add(fmt.Sprintf("AlternArtikelNr2_%d", index+1), value.AlternArticleNo2)
		parameter.Add(fmt.Sprintf("Artikelname_%d", index+1), value.ArticleName)
		parameter.Add(fmt.Sprintf("ArtikelEpreis_%d", index+1), value.ArticleEPrice)
		parameter.Add(fmt.Sprintf("ArtikelMwSt_%d", index+1), value.ArticleVat)
		parameter.Add(fmt.Sprintf("ArtikelMenge_%d", index+1), value.ArticleQuantity)
		parameter.Add(fmt.Sprintf("ArtikelLink_%d", index+1), value.ArticleLink)
		parameter.Add(fmt.Sprintf("ArtikelStammID_%d", index+1), value.ArticleMasterId)

	}

	// Set new url
	request.URL.RawQuery = parameter.Encode()

	// Config new request
	c := Config{request, nil}

	// Send new request
	response, err := c.Send()
	if err != nil {
		return AddOrderReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Decode data
	var decode AddOrderReturn

	err = xml.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return AddOrderReturn{}, err
	}

	// Return data
	return decode, nil

}
