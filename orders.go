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
	"net/url"
	"strings"
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
	UseComplWeight          string
	UseProductTaxRate       string
	BuyDate                 string
	ShippingMethod          string
	ShippingCosts           string
	PaymentMethodsSurcharge string
	NoFeedback              string
	NoShippingCalc          string
	ShippingGroup           string
	DoNotDisplayVat         string
	PaymentMethod           string
	VMemo                   string
	OrderInfo1              string
	OrderInfo2              string
	OrderInfo3              string
	VId                     string
	SoldCurrency            string
	SetPay                  string
	SetPayDate              string
	CheckVId                string
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
	Attribute        string
	ArticleMasterId  string
	ArticleTag1      string
	ArticleTag2      string
	ArticleTag3      string
	ArticleTag4      string
	ArticleTag5      string
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

	// Config new request
	c := Config{
		BaseUrl: "https://api.afterbuy.de/afterbuy/ShopInterfaceUTF8.aspx",
		Method:  "POST",
		Body:    nil,
		Header:  map[string]string{},
	}

	// Add header to config
	header := make(map[string]string)
	header["Content-Type"] = "application/x-www-form-urlencoded"
	c.Header = header

	// Parse the url from base url
	parse, err := url.Parse(c.BaseUrl)
	if err != nil {
		return AddOrderReturn{}, err
	}

	// Check the parameter
	parameter := parse.Query()

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
	parameter.Add("UseComplWeight", body.UseComplWeight)
	parameter.Add("UseProductTaxRate", body.UseProductTaxRate)
	parameter.Add("BuyDate", body.BuyDate)
	parameter.Add("Versandart", body.ShippingMethod)
	parameter.Add("Versandkosten", body.ShippingCosts)
	parameter.Add("ZahlartenAufschlag", body.PaymentMethodsSurcharge)
	parameter.Add("NoFeedback", body.NoFeedback)
	parameter.Add("NoVersandCalc", body.NoShippingCalc)
	parameter.Add("Versandgruppe", body.ShippingGroup)
	parameter.Add("MwStNichtAusweisen", body.DoNotDisplayVat)
	parameter.Add("Zahlart", body.PaymentMethod)
	parameter.Add("VMemo", body.VMemo)
	parameter.Add("OrderInfo1", body.OrderInfo1)
	parameter.Add("OrderInfo2", body.OrderInfo2)
	parameter.Add("OrderInfo3", body.OrderInfo3)
	parameter.Add("VID", body.VId)
	parameter.Add("SoldCurrency", body.SoldCurrency)
	parameter.Add("SetPay", body.SetPay)
	parameter.Add("SetPayDate", body.SetPayDate)
	parameter.Add("CheckVID", body.CheckVId)
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
		parameter.Add(fmt.Sprintf("Attribute_%d", index+1), value.Attribute)
		parameter.Add(fmt.Sprintf("ArtikelStammID_%d", index+1), value.ArticleMasterId)
		parameter.Add(fmt.Sprintf("ArtikelTag_%d_1", index+1), value.ArticleTag1)
		parameter.Add(fmt.Sprintf("ArtikelTag_%d_2", index+1), value.ArticleTag2)
		parameter.Add(fmt.Sprintf("ArtikelTag_%d_3", index+1), value.ArticleTag3)
		parameter.Add(fmt.Sprintf("ArtikelTag_%d_4", index+1), value.ArticleTag4)
		parameter.Add(fmt.Sprintf("ArtikelTag_%d_5", index+1), value.ArticleTag5)

	}

	// Save parameter to config struct
	c.Body = strings.NewReader(parameter.Encode())

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
