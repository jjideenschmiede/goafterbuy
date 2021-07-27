//**********************************************************
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of goafterbuy.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package goafterbuy

import (
	"encoding/xml"
)

// ProductsBody is to send the body data
type ProductsBody struct {
	Request ProductsRequest `xml:"Request"`
}

type ProductsRequest struct {
	AfterbuyGlobal                 AfterbuyGlobal `xml:"AfterbuyGlobal"`
	MaxShopItems                   int            `xml:"MaxShopItems"`
	SuppressBaseProductRelatedData int            `xml:"SuppressBaseProductRelatedData"`
	PaginationEnabled              int            `xml:"PaginationEnabled"`
	PageNumber                     int            `xml:"PageNumber"`
}

// ProductsReturn is to decode the xml data
type ProductsReturn struct {
	XmlName    xml.Name             `xml:"Afterbuy"`
	CallStatus string               `xml:"CallStatus"`
	CallName   string               `xml:"CallName"`
	VersionId  int                  `xml:"VersionID"`
	Result     ProductsReturnResult `xml:"Result"`
}

type ProductsReturnResult struct {
	XmlName             xml.Name                       `xml:"Result"`
	HasMoreProducts     int                            `xml:"HasMoreProducts"`
	Products            ProductsReturnProducts         `xml:"Products"`
	ShippingServiceList string                         `xml:"ShippingServiceList"`
	LastProductId       int                            `xml:"LastProductID"`
	PaginationResult    ProductsReturnPaginationResult `xml:"PaginationResult"`
}

type ProductsReturnProducts struct {
	XmlName xml.Name                `xml:"Products"`
	Product []ProductsReturnProduct `xml:"Product"`
}
type ProductsReturnProduct struct {
	XmlName                xml.Name                      `xml:"Product"`
	ProductShopOption      int                           `xml:"ProductShopOption"`
	ProductId              int                           `xml:"ProductID"`
	Anr                    int                           `xml:"Anr"`
	Ean                    string                        `xml:"EAN"`
	Name                   string                        `xml:"Name"`
	SeoName                string                        `xml:"SeoName"`
	ModDate                string                        `xml:"ModDate"`
	ShortDescription       string                        `xml:"ShortDescription"`
	Memo                   string                        `xml:"Memo"`
	GoogleBaseLabels       string                        `xml:"GoogleBaseLabels"`
	Description            string                        `xml:"Description"`
	Keywords               string                        `xml:"Keywords"`
	AvailableShop          int                           `xml:"AvailableShop"`
	Quantity               int                           `xml:"Quantity"`
	AuctionQuantity        int                           `xml:"AuctionQuantity"`
	Stock                  int                           `xml:"Stock"`
	Discontinued           int                           `xml:"Discontinued"`
	MergeStock             int                           `xml:"MergeStock"`
	UnitOfQuantity         string                        `xml:"UnitOfQuantity"`
	BasepriceFactor        int                           `xml:"BasepriceFactor"`
	MinimumStock           int                           `xml:"MinimumStock"`
	MinimumOrderQuantity   int                           `xml:"MinimumOrderQuantity"`
	FullFilmentQuantity    int                           `xml:"FullFilmentQuantity"`
	FullFilmentImport      interface{}                   `xml:"FullFilmentImport"`
	SellingPrice           string                        `xml:"SellingPrice"`
	BuyingPrice            string                        `xml:"BuyingPrice"`
	DealerPrice            string                        `xml:"DealerPrice"`
	Level                  int                           `xml:"Level"`
	Position               int                           `xml:"Position"`
	TitleReplace           int                           `xml:"TitleReplace"`
	ScaledDiscounts        ProductsReturnScaledDiscounts `xml:"ScaledDiscounts"`
	TaxRate                int                           `xml:"TaxRate"`
	Weight                 int                           `xml:"Weight"`
	SearchAlias            string                        `xml:"SearchAlias"`
	Froogle                int                           `xml:"Froogle"`
	GoogleBaseShipping     string                        `xml:"GoogleBaseShipping"`
	Kelkoo                 int                           `xml:"Kelkoo"`
	ShippingGroup          string                        `xml:"ShippingGroup"`
	ShopShippingGroup      string                        `xml:"ShopShippingGroup"`
	SearchEngineShipping   string                        `xml:"SearchEngineShipping"`
	CrossCatalogId         int                           `xml:"CrossCatalogID"`
	FreeValue1             string                        `xml:"FreeValue1"`
	FreeValue2             string                        `xml:"FreeValue2"`
	FreeValue3             string                        `xml:"FreeValue3"`
	FreeValue4             string                        `xml:"FreeValue4"`
	FreeValue5             string                        `xml:"FreeValue5"`
	DeliveryTime           string                        `xml:"DeliveryTime"`
	Stocklocation1         string                        `xml:"Stocklocation_1"`
	Stocklocation2         string                        `xml:"Stocklocation_2"`
	Stocklocation3         string                        `xml:"Stocklocation_3"`
	Stocklocation4         string                        `xml:"Stocklocation_4"`
	ImageSmallUrl          string                        `xml:"image_small_url"`
	ImageLargeUrl          string                        `xml:"ImageLargeURL"`
	ProductBrand           string                        `xml:"ProductBrand"`
	CustomsTariffNumber    string                        `xml:"CustomsTariffNumber"`
	CountryOfOrigin        string                        `xml:"CountryOfOrigin"`
	LastSale               string                        `xml:"LastSale"`
	Facebook               int                           `xml:"Facebook"`
	ManufacturerPartNumber string                        `xml:"ManufacturerPartNumber"`
	Condition              int                           `xml:"Condition"`
	AgeGroup               int                           `xml:"AgeGroup"`
	Gender                 int                           `xml:"Gender"`
	Material               string                        `xml:"Material"`
	Pattern                string                        `xml:"Pattern"`
	CustomLabel0           string                        `xml:"CustomLabel0"`
	CustomLabel1           string                        `xml:"CustomLabel1"`
	CustomLabel2           string                        `xml:"CustomLabel2"`
	CustomLabel3           string                        `xml:"CustomLabel3"`
	CustomLabel4           string                        `xml:"CustomLabel4"`
	MultiLanguage          interface{}                   `xml:"MultiLanguage"`
}

type ProductsReturnScaledDiscounts struct {
	XmlName        xml.Name                       `xml:"ScaledDiscounts"`
	ScaledDiscount []ProductsReturnScaledDiscount `xml:"ScaledDiscount"`
}

type ProductsReturnScaledDiscount struct {
	XmlName        xml.Name `xml:"ScaledDiscount"`
	ScaledQuantity int      `xml:"ScaledQuantity"`
	ScaledPrice    string   `xml:"ScaledPrice"`
	ScaledDPrice   string   `xml:"ScaledDPrice"`
}

type ProductsReturnPaginationResult struct {
	XmlName              xml.Name `xml:"PaginationResult"`
	TotalNumberOfEntries int      `xml:"TotalNumberOfEntries"`
	TotalNumberOfPages   int      `xml:"TotalNumberOfPages"`
	ItemsPerPage         int      `xml:"ItemsPerPage"`
	PageNumber           int      `xml:"PageNumber"`
}

// Products is to get all products from afterbuy
func Products(body *ProductsBody) (*ProductsReturn, error) {

	// Convert body
	convert, err := xml.Marshal(body)
	if err != nil {
		return nil, err
	}

	// Config new request
	r := Config{convert}

	// Send new request
	response, err := r.Send()
	if err != nil {
		return nil, err
	}

	// Close request body
	defer response.Body.Close()

	// Decode data
	var decode ProductsReturn

	err = xml.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return nil, err
	}

	// Return data
	return &decode, nil

}
