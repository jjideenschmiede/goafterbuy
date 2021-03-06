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
)

// ProductsBody is to send the body data
type ProductsBody struct {
	Request ProductsRequest `xml:"Request"`
}

type ProductsRequest struct {
	AfterbuyGlobal                 AfterbuyGlobal             `xml:"AfterbuyGlobal"`
	DataFilter                     *ProductsRequestDataFilter `xml:"DataFilter"`
	MaxShopItems                   int                        `xml:"MaxShopItems"`
	SuppressBaseProductRelatedData int                        `xml:"SuppressBaseProductRelatedData"`
	PaginationEnabled              int                        `xml:"PaginationEnabled"`
	PageNumber                     int                        `xml:"PageNumber"`
}

type ProductsRequestDataFilter struct {
	Filter []ProductsRequestFilter `xml:"Filter"`
}

type ProductsRequestFilter struct {
	FilterName   string                      `xml:"FilterName"`
	FilterValues ProductsRequestFilterValues `xml:"FilterValues"`
}

type ProductsRequestFilterValues struct {
	LevelFrom   int      `xml:"LevelFrom,omitempty"`
	LevelTo     int      `xml:"LevelTo,omitempty"`
	ValueFrom   int      `xml:"ValueFrom,omitempty"`
	ValueTo     int      `xml:"ValueTo,omitempty"`
	DateFrom    string   `xml:"DateFrom,omitempty"`
	DateTo      string   `xml:"DateTo,omitempty"`
	FilterValue []string `xml:"FilterValue,omitempty"`
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
	XmlName                            xml.Name                       `xml:"Product"`
	ProductShopOption                  int                            `xml:"ProductShopOption"`
	ProductId                          int                            `xml:"ProductID"`
	Anr                                int                            `xml:"Anr"`
	Ean                                string                         `xml:"EAN"`
	Name                               string                         `xml:"Name"`
	CanonicalUrl                       string                         `xml:"CanonicalUrl"`
	SeoName                            string                         `xml:"SeoName"`
	ModDate                            string                         `xml:"ModDate"`
	BaseProductFlag                    int                            `xml:"BaseProductFlag"`
	BaseProductShowMode                int                            `xml:"BaseProductShowMode"`
	BaseProducts                       ProductsReturnBaseProducts     `xml:"BaseProducts"`
	ShortDescription                   string                         `xml:"ShortDescription"`
	Memo                               string                         `xml:"Memo"`
	GoogleBaseLabels                   string                         `xml:"GoogleBaseLabels"`
	Description                        string                         `xml:"Description"`
	Keywords                           string                         `xml:"Keywords"`
	AvailableShop                      int                            `xml:"AvailableShop"`
	Quantity                           int                            `xml:"Quantity"`
	AuctionQuantity                    int                            `xml:"AuctionQuantity"`
	Stock                              int                            `xml:"Stock"`
	Discontinued                       int                            `xml:"Discontinued"`
	MergeStock                         int                            `xml:"MergeStock"`
	UnitOfQuantity                     string                         `xml:"UnitOfQuantity"`
	BasepriceFactor                    string                         `xml:"BasepriceFactor"`
	MinimumStock                       int                            `xml:"MinimumStock"`
	MinimumOrderQuantity               int                            `xml:"MinimumOrderQuantity"`
	FullFilmentQuantity                int                            `xml:"FullFilmentQuantity"`
	FullFilmentImport                  interface{}                    `xml:"FullFilmentImport"`
	SellingPrice                       string                         `xml:"SellingPrice"`
	BuyingPrice                        string                         `xml:"BuyingPrice"`
	DealerPrice                        string                         `xml:"DealerPrice"`
	AdditionalPrices                   ProductsReturnAdditionalPrices `xml:"AdditionalPrices"`
	Level                              int                            `xml:"Level"`
	Position                           int                            `xml:"Position"`
	TitleReplace                       int                            `xml:"TitleReplace"`
	ScaledDiscounts                    ProductsReturnScaledDiscounts  `xml:"ScaledDiscounts"`
	TaxRate                            string                         `xml:"TaxRate"`
	Weight                             string                         `xml:"Weight"`
	SearchAlias                        string                         `xml:"SearchAlias"`
	Froogle                            int                            `xml:"Froogle"`
	GoogleBaseShipping                 string                         `xml:"GoogleBaseShipping"`
	Kelkoo                             int                            `xml:"Kelkoo"`
	ShippingGroup                      string                         `xml:"ShippingGroup"`
	ShopShippingGroup                  string                         `xml:"ShopShippingGroup"`
	SearchEngineShipping               string                         `xml:"SearchEngineShipping"`
	CrossCatalogId                     int                            `xml:"CrossCatalogID"`
	FreeValue1                         string                         `xml:"FreeValue1"`
	FreeValue2                         string                         `xml:"FreeValue2"`
	FreeValue3                         string                         `xml:"FreeValue3"`
	FreeValue4                         string                         `xml:"FreeValue4"`
	FreeValue5                         string                         `xml:"FreeValue5"`
	FreeValue6                         string                         `xml:"FreeValue6"`
	FreeValue7                         string                         `xml:"FreeValue7"`
	FreeValue8                         string                         `xml:"FreeValue8"`
	FreeValue9                         string                         `xml:"FreeValue9"`
	FreeValue10                        string                         `xml:"FreeValue10"`
	DeliveryTime                       string                         `xml:"DeliveryTime"`
	Stocklocation1                     string                         `xml:"Stocklocation_1"`
	Stocklocation2                     string                         `xml:"Stocklocation_2"`
	Stocklocation3                     string                         `xml:"Stocklocation_3"`
	Stocklocation4                     string                         `xml:"Stocklocation_4"`
	ImageSmallUrl                      string                         `xml:"image_small_url"`
	ImageLargeUrl                      string                         `xml:"ImageLargeURL"`
	ProductPictures                    ProductsReturnProductPictures  `xml:"ProductPictures"`
	AmazonStandardProductIdType        string                         `xml:"AmazonStandardProductIDType"`
	AmazonStandardProductIdValue       string                         `xml:"AmazonStandardProductIDValue"`
	ManufacturerStandardProductIdType  string                         `xml:"ManufacturerStandardProductIDType"`
	ManufacturerStandardProductIdValue string                         `xml:"ManufacturerStandardProductIDValue"`
	ProductBrand                       string                         `xml:"ProductBrand"`
	CustomsTariffNumber                string                         `xml:"CustomsTariffNumber"`
	CountryOfOrigin                    string                         `xml:"CountryOfOrigin"`
	LastSale                           string                         `xml:"LastSale"`
	Catalogs                           ProductsReturnCatalogs         `xml:"Catalogs"`
	Attributes                         ProductsReturnAttributes       `xml:"Attributes"`
	Facebook                           int                            `xml:"Facebook"`
	ManufacturerPartNumber             string                         `xml:"ManufacturerPartNumber"`
	Condition                          int                            `xml:"Condition"`
	AgeGroup                           int                            `xml:"AgeGroup"`
	Gender                             int                            `xml:"Gender"`
	Material                           string                         `xml:"Material"`
	Pattern                            string                         `xml:"Pattern"`
	CustomLabel0                       string                         `xml:"CustomLabel0"`
	CustomLabel1                       string                         `xml:"CustomLabel1"`
	CustomLabel2                       string                         `xml:"CustomLabel2"`
	CustomLabel3                       string                         `xml:"CustomLabel3"`
	CustomLabel4                       string                         `xml:"CustomLabel4"`
	MultiLanguage                      ProductsReturnMultiLanguage    `xml:"MultiLanguage"`
}

type ProductsReturnBaseProducts struct {
	XmlName     xml.Name                    `xml:"BaseProducts"`
	BaseProduct []ProductsReturnBaseProduct `xml:"BaseProduct"`
}

type ProductsReturnBaseProduct struct {
	XmlName                  xml.Name                               `xml:"BaseProduct"`
	BaseProductId            int                                    `xml:"BaseProductID"`
	BaseProductsRelationData ProductsReturnBaseProductsRelationData `xml:"BaseProductsRelationData"`
}

type ProductsReturnBaseProductsRelationData struct {
	XmlName           xml.Name                          `xml:"BaseProductsRelationData"`
	Postion           int                               `xml:"Position"`
	VariationLabel    string                            `xml:"VariationLabel"`
	DefaultProduct    int                               `xml:"DefaultProduct"`
	EbayVariationData []ProductsReturnEbayVariationData `xml:"eBayVariationData"`
}

type ProductsReturnEbayVariationData struct {
	XmlName               xml.Name `xml:"eBayVariationData"`
	EbayVariationName     string   `xml:"eBayVariationName"`
	EbayVariationValue    string   `xml:"eBayVariationValue"`
	EbayVariationPosition int      `xml:"eBayVariationPosition"`
	EbayVariationUrls     string   `xml:"eBayVariationUrls"`
}

type ProductsReturnAdditionalPrices struct {
	XmlName         xml.Name                        `xml:"AdditionalPrices"`
	AdditionalPrice []ProductsReturnAdditionalPrice `xml:"AdditionalPrice"`
}

type ProductsReturnAdditionalPrice struct {
	XmlName      xml.Name `xml:"AdditionalPrice"`
	DefinitionId int      `xml:"DefinitionId"`
	Name         string   `xml:"Name"`
	Value        string   `xml:"Value"`
	Pretax       bool     `xml:"Pretax"`
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

type ProductsReturnProductPictures struct {
	XmlName        xml.Name                       `xml:"ProductPictures"`
	ProductPicture []ProductsReturnProductPicture `xml:"ProductPicture"`
}

type ProductsReturnProductPicture struct {
	XmlName xml.Name `xml:"ProductPicture"`
	Nr      int      `xml:"Nr"`
	Type    int      `xml:"Type"`
	Url     string   `xml:"Url"`
	AltText string   `xml:"AltText"`
}

type ProductsReturnCatalogs struct {
	XmlName   xml.Name `xml:"Catalogs"`
	CatalogId []int    `xml:"CatalogID"`
}

type ProductsReturnAttributes struct {
	XmlName  xml.Name                 `xml:"Attributes"`
	Attribut []ProductsReturnAttribut `xml:"Attribut"`
}

type ProductsReturnAttribut struct {
	XmlName          xml.Name `xml:"Attribut"`
	AttributName     string   `xml:"AttributName"`
	AttributValue    string   `xml:"AttributValue"`
	AttributType     int      `xml:"AttributType"`
	AttributRequired int      `xml:"AttributRequired"`
	AttributPosition int      `xml:"AttributPosition"`
}

type ProductsReturnPaginationResult struct {
	XmlName              xml.Name `xml:"PaginationResult"`
	TotalNumberOfEntries int      `xml:"TotalNumberOfEntries"`
	TotalNumberOfPages   int      `xml:"TotalNumberOfPages"`
	ItemsPerPage         int      `xml:"ItemsPerPage"`
	PageNumber           int      `xml:"PageNumber"`
}

type ProductsReturnMultiLanguage struct {
	XmlName xml.Name                  `xml:"MultiLanguage"`
	At      ProductsReturnLanguageAt  `xml:"AT"`
	B       ProductsReturnLanguageB   `xml:"B"`
	Bg      ProductsReturnLanguageBg  `xml:"BG"`
	Ch      ProductsReturnLanguageCh  `xml:"CH"`
	Cy      ProductsReturnLanguageCy  `xml:"CY"`
	Cz      ProductsReturnLanguageCz  `xml:"CZ"`
	De      ProductsReturnLanguageDe  `xml:"DE"`
	Dk      ProductsReturnLanguageDe  `xml:"DK"`
	E       ProductsReturnLanguageE   `xml:"E"`
	Fin     ProductsReturnLanguageFin `xml:"FIN"`
	Fr      ProductsReturnLanguageFr  `xml:"FR"`
	Gb      ProductsReturnLanguageGb  `xml:"GB"`
	Gr      ProductsReturnLanguageGr  `xml:"GR"`
	H       ProductsReturnLanguageH   `xml:"H"`
	Hr      ProductsReturnLanguageHr  `xml:"HR"`
	I       ProductsReturnLanguageI   `xml:"I"`
	Irl     ProductsReturnLanguageI   `xml:"IRL"`
	Is      ProductsReturnLanguageIs  `xml:"IS"`
	J       ProductsReturnLanguageJ   `xml:"J"`
	Kp      ProductsReturnLanguageKp  `xml:"KP"`
	Lt      ProductsReturnLanguageLt  `xml:"LT"`
	Lu      ProductsReturnLanguageLu  `xml:"LU"`
	Lv      ProductsReturnLanguageLv  `xml:"LV"`
	Mt      ProductsReturnLanguageMt  `xml:"MT"`
	N       ProductsReturnLanguageN   `xml:"N"`
	Nl      ProductsReturnLanguageNl  `xml:"NL"`
	P       ProductsReturnLanguageP   `xml:"P"`
	Pl      ProductsReturnLanguagePl  `xml:"PL"`
	Ro      ProductsReturnLanguageRo  `xml:"RO"`
	Ru      ProductsReturnLanguageRu  `xml:"RU"`
	S       ProductsReturnLanguageS   `xml:"S"`
	Sk      ProductsReturnLanguageSk  `xml:"SK"`
	Slo     ProductsReturnLanguageSlo `xml:"SLO"`
	Tr      ProductsReturnLanguageTr  `xml:"TR"`
	Us      ProductsReturnLanguageUs  `xml:"US"`
}

type ProductsReturnLanguageAt struct {
	XmlName          xml.Name `xml:"AT"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageB struct {
	XmlName          xml.Name `xml:"B"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageBg struct {
	XmlName          xml.Name `xml:"BG"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageCh struct {
	XmlName          xml.Name `xml:"CH"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageCy struct {
	XmlName          xml.Name `xml:"CY"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageCz struct {
	XmlName          xml.Name `xml:"CZ"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageDe struct {
	XmlName          xml.Name `xml:"DE"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageDk struct {
	XmlName          xml.Name `xml:"DK"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageE struct {
	XmlName          xml.Name `xml:"E"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageFin struct {
	XmlName          xml.Name `xml:"FIN"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageFr struct {
	XmlName          xml.Name `xml:"FR"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageGb struct {
	XmlName          xml.Name `xml:"GB"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageGr struct {
	XmlName          xml.Name `xml:"GR"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageH struct {
	XmlName          xml.Name `xml:"H"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageHr struct {
	XmlName          xml.Name `xml:"HR"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageI struct {
	XmlName          xml.Name `xml:"I"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageIrl struct {
	XmlName          xml.Name `xml:"IRL"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageIs struct {
	XmlName          xml.Name `xml:"IS"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageJ struct {
	XmlName          xml.Name `xml:"J"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageKp struct {
	XmlName          xml.Name `xml:"KP"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageLt struct {
	XmlName          xml.Name `xml:"LT"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageLu struct {
	XmlName          xml.Name `xml:"LU"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageLv struct {
	XmlName          xml.Name `xml:"LV"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageMt struct {
	XmlName          xml.Name `xml:"MT"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageN struct {
	XmlName          xml.Name `xml:"N"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageNl struct {
	XmlName          xml.Name `xml:"NL"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageP struct {
	XmlName          xml.Name `xml:"P"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguagePl struct {
	XmlName          xml.Name `xml:"PL"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageRo struct {
	XmlName          xml.Name `xml:"RO"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageRu struct {
	XmlName          xml.Name `xml:"RU"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageS struct {
	XmlName          xml.Name `xml:"S"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageSk struct {
	XmlName          xml.Name `xml:"SK"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageSlo struct {
	XmlName          xml.Name `xml:"SLO"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageTr struct {
	XmlName          xml.Name `xml:"TR"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

type ProductsReturnLanguageUs struct {
	XmlName          xml.Name `xml:"US"`
	Alias            string   `xml:"Alias"`
	Beschreibung     string   `xml:"Beschreibung"`
	Kurzbeschreibung string   `xml:"Kurzbeschreibung"`
	Name             string   `xml:"Name"`
	Ust              string   `xml:"ust"`
}

// Products are to get all products from ab interface
func Products(body ProductsBody) (ProductsReturn, error) {

	// Convert body
	convert, err := xml.Marshal(body)
	if err != nil {
		return ProductsReturn{}, err
	}

	// Config new request
	c := Config{nil, convert}

	// Send new request
	response, err := c.Send()
	if err != nil {
		return ProductsReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Decode data
	var decode ProductsReturn

	err = xml.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductsReturn{}, err
	}

	// Return data
	return decode, nil

}
