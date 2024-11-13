//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2022 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of goafterbuy.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package goafterbuy

import (
	"bytes"
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
	ErrorList           ProductsReturnErrorList        `xml:"ErrorList"`
}

type ProductsReturnProducts struct {
	XmlName xml.Name                `xml:"Products"`
	Product []ProductsReturnProduct `xml:"Product"`
}

type ProductsReturnProduct struct {
	XmlName                            xml.Name                                  `xml:"Product"`
	ProductShopOption                  int                                       `xml:"ProductShopOption"`
	ProductId                          int                                       `xml:"ProductID"`
	Anr                                int                                       `xml:"Anr"`
	Ean                                string                                    `xml:"EAN"`
	Name                               string                                    `xml:"Name"`
	CanonicalUrl                       string                                    `xml:"CanonicalUrl"`
	SeoName                            string                                    `xml:"SeoName"`
	ModDate                            string                                    `xml:"ModDate"`
	BaseProductFlag                    int                                       `xml:"BaseProductFlag"`
	BaseProductShowMode                int                                       `xml:"BaseProductShowMode"`
	BaseProducts                       ProductsReturnBaseProducts                `xml:"BaseProducts"`
	ShortDescription                   string                                    `xml:"ShortDescription"`
	Memo                               string                                    `xml:"Memo"`
	GoogleBaseLabels                   string                                    `xml:"GoogleBaseLabels"`
	HeaderId                           int                                       `xml:"HeaderID"`
	HeaderDescriptionName              string                                    `xml:"HeaderDescriptionName"`
	HeaderDescriptionValue             string                                    `xml:"HeaderDescriptionValue"`
	Description                        string                                    `xml:"Description"`
	FooterId                           int                                       `xml:"FooterID"`
	FooterDescriptionName              string                                    `xml:"FooterDescriptionName"`
	FooterDescriptionValue             string                                    `xml:"FooterDescriptionValue"`
	Keywords                           string                                    `xml:"Keywords"`
	AvailableShop                      int                                       `xml:"AvailableShop"`
	Quantity                           int                                       `xml:"Quantity"`
	AuctionQuantity                    int                                       `xml:"AuctionQuantity"`
	Stock                              int                                       `xml:"Stock"`
	Discontinued                       int                                       `xml:"Discontinued"`
	MergeStock                         int                                       `xml:"MergeStock"`
	UnitOfQuantity                     string                                    `xml:"UnitOfQuantity"`
	BasepriceFactor                    string                                    `xml:"BasepriceFactor"`
	MinimumStock                       int                                       `xml:"MinimumStock"`
	MinimumOrderQuantity               int                                       `xml:"MinimumOrderQuantity"`
	FullFilmentQuantity                int                                       `xml:"FullFilmentQuantity"`
	FullFilmentImport                  interface{}                               `xml:"FullFilmentImport"`
	SellingPrice                       string                                    `xml:"SellingPrice"`
	BuyingPrice                        string                                    `xml:"BuyingPrice"`
	DealerPrice                        string                                    `xml:"DealerPrice"`
	AdditionalPrices                   ProductsReturnAdditionalPrices            `xml:"AdditionalPrices"`
	Level                              int                                       `xml:"Level"`
	Position                           int                                       `xml:"Position"`
	TitleReplace                       int                                       `xml:"TitleReplace"`
	ScaledDiscounts                    ProductsReturnScaledDiscounts             `xml:"ScaledDiscounts"`
	TaxRate                            string                                    `xml:"TaxRate"`
	Weight                             string                                    `xml:"Weight"`
	SearchAlias                        string                                    `xml:"SearchAlias"`
	Froogle                            int                                       `xml:"Froogle"`
	GoogleBaseShipping                 string                                    `xml:"GoogleBaseShipping"`
	Kelkoo                             int                                       `xml:"Kelkoo"`
	ShippingGroup                      string                                    `xml:"ShippingGroup"`
	ShopShippingGroup                  string                                    `xml:"ShopShippingGroup"`
	SearchEngineShipping               string                                    `xml:"SearchEngineShipping"`
	CrossCatalogId                     int                                       `xml:"CrossCatalogID"`
	Features                           ProductsReturnFeatures                    `xml:"Features"`
	FreeValue1                         string                                    `xml:"FreeValue1"`
	FreeValue2                         string                                    `xml:"FreeValue2"`
	FreeValue3                         string                                    `xml:"FreeValue3"`
	FreeValue4                         string                                    `xml:"FreeValue4"`
	FreeValue5                         string                                    `xml:"FreeValue5"`
	FreeValue6                         string                                    `xml:"FreeValue6"`
	FreeValue7                         string                                    `xml:"FreeValue7"`
	FreeValue8                         string                                    `xml:"FreeValue8"`
	FreeValue9                         string                                    `xml:"FreeValue9"`
	FreeValue10                        string                                    `xml:"FreeValue10"`
	DeliveryTime                       string                                    `xml:"DeliveryTime"`
	Stocklocation1                     string                                    `xml:"Stocklocation_1"`
	Stocklocation2                     string                                    `xml:"Stocklocation_2"`
	Stocklocation3                     string                                    `xml:"Stocklocation_3"`
	Stocklocation4                     string                                    `xml:"Stocklocation_4"`
	ImageSmallUrl                      string                                    `xml:"ImageSmallURL"`
	ImageLargeUrl                      string                                    `xml:"ImageLargeURL"`
	ProductPictures                    ProductsReturnProductPictures             `xml:"ProductPictures"`
	AmazonStandardProductIdType        string                                    `xml:"AmazonStandardProductIDType"`
	AmazonStandardProductIdValue       string                                    `xml:"AmazonStandardProductIDValue"`
	ManufacturerStandardProductIdType  string                                    `xml:"ManufacturerStandardProductIDType"`
	ManufacturerStandardProductIdValue string                                    `xml:"ManufacturerStandardProductIDValue"`
	ProductBrand                       string                                    `xml:"ProductBrand"`
	CustomsTariffNumber                string                                    `xml:"CustomsTariffNumber"`
	CountryOfOrigin                    string                                    `xml:"CountryOfOrigin"`
	LastSale                           string                                    `xml:"LastSale"`
	AdditionalDescriptionFields        ProductsReturnAdditionalDescriptionFields `xml:"AdditionalDescriptionFields"`
	Catalogs                           ProductsReturnCatalogs                    `xml:"Catalogs"`
	Attributes                         ProductsReturnAttributes                  `xml:"Attributes"`
	PartsFitment                       ProductsReturnPartsFitment                `xml:"PartsFitment"`
	GoogleProductCategory              string                                    `xml:"GoogleProductCategory"`
	Facebook                           int                                       `xml:"Facebook"`
	ManufacturerPartNumber             string                                    `xml:"ManufacturerPartNumber"`
	Condition                          int                                       `xml:"Condition"`
	AgeGroup                           int                                       `xml:"AgeGroup"`
	Gender                             int                                       `xml:"Gender"`
	Material                           string                                    `xml:"Material"`
	Pattern                            string                                    `xml:"Pattern"`
	CustomLabel0                       string                                    `xml:"CustomLabel0"`
	CustomLabel1                       string                                    `xml:"CustomLabel1"`
	CustomLabel2                       string                                    `xml:"CustomLabel2"`
	CustomLabel3                       string                                    `xml:"CustomLabel3"`
	CustomLabel4                       string                                    `xml:"CustomLabel4"`
	MultiLanguage                      ProductsReturnMultiLanguage               `xml:"MultiLanguage"`
	EconomicOperator                   ProductsReturnEconomicOperator            `xml:"EconomicOperator"`
}

type ProductsReturnBaseProducts struct {
	XmlName     xml.Name                    `xml:"BaseProducts"`
	BaseProduct []ProductsReturnBaseProduct `xml:"BaseProduct"`
}

type ProductsReturnBaseProduct struct {
	XmlName                  xml.Name                               `xml:"BaseProduct"`
	BaseProductId            int                                    `xml:"BaseProductID"`
	BaseProductType          int                                    `xml:"BaseProductType"`
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

type ProductsReturnFeatures struct {
	XmlName xml.Name                `xml:"Features"`
	Feature []ProductsReturnFeature `xml:"Feature"`
}

type ProductsReturnFeature struct {
	XmlName xml.Name `xml:"Feature"`
	Id      int      `xml:"ID"`
	Name    string   `xml:"Name"`
	Value   string   `xml:"Value"`
}

type ProductsReturnProductPictures struct {
	XmlName        xml.Name                       `xml:"ProductPictures"`
	ProductPicture []ProductsReturnProductPicture `xml:"ProductPicture"`
}

type ProductsReturnProductPicture struct {
	XmlName xml.Name             `xml:"ProductPicture"`
	Nr      int                  `xml:"Nr"`
	Typ     int                  `xml:"Typ"`
	Url     string               `xml:"Url"`
	AltText string               `xml:"AltText"`
	Childs  ProductsReturnChilds `xml:"Childs"`
}

type ProductsReturnChilds struct {
	XmlName        xml.Name                             `xml:"Childs"`
	ProductPicture []ProductsReturnChildsProductPicture `xml:"ProductPicture"`
}

type ProductsReturnChildsProductPicture struct {
	XmlName xml.Name `xml:"ProductPicture"`
	Nr      int      `xml:"Nr"`
	Typ     int      `xml:"Typ"`
	Url     string   `xml:"Url"`
	AltText string   `xml:"AltText"`
}

type ProductsReturnAdditionalDescriptionFields struct {
	XmlName                    xml.Name                                   `xml:"AdditionalDescriptionFields"`
	AdditionalDescriptionField []ProductsReturnAdditionalDescriptionField `xml:"AdditionalDescriptionField"`
}

type ProductsReturnAdditionalDescriptionField struct {
	XmlName      xml.Name `xml:"AdditionalDescriptionField"`
	FieldId      int      `xml:"FieldID"`
	FieldName    string   `xml:"FieldName"`
	FieldLabel   string   `xml:"FieldLabel"`
	FieldContent string   `xml:"FieldContent"`
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

type ProductsReturnPartsFitment struct {
	XmlName         xml.Name                        `xml:"ProductsReturnPartsFitment"`
	PartsProperties []ProductsReturnPartsProperties `xml:"PartsProperties"`
}

type ProductsReturnPartsProperties struct {
	XmlName       xml.Name                      `xml:"PartsProperties"`
	PartsProperty []ProductsReturnPartsProperty `xml:"PartsProperty"`
}

type ProductsReturnPartsProperty struct {
	XmlName       xml.Name `xml:"PartsProperty"`
	PropertyName  string   `xml:"PropertyName"`
	PropertyValue string   `xml:"PropertyValue"`
}

type ProductsReturnPaginationResult struct {
	XmlName              xml.Name `xml:"PaginationResult"`
	TotalNumberOfEntries int      `xml:"TotalNumberOfEntries"`
	TotalNumberOfPages   int      `xml:"TotalNumberOfPages"`
	ItemsPerPage         int      `xml:"ItemsPerPage"`
	PageNumber           int      `xml:"PageNumber"`
}

type ProductsReturnEconomicOperator struct {
	XmlName         xml.Name `xml:"EconomicOperator"`
	DatabaseId      int      `xml:"DatabaseId"`
	ContactType     int      `xml:"ContactType"`
	Company         string   `xml:"Company"`
	Street1         string   `xml:"Street1"`
	Street2         string   `xml:"Street2"`
	PostalCode      string   `xml:"PostalCode"`
	City            string   `xml:"City"`
	StateOrProvince string   `xml:"StateOrProvince"`
	Country         string   `xml:"Country"`
	Email           string   `xml:"Email"`
	Phone           string   `xml:"Phone"`
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

type ProductsReturnErrorList struct {
	XmlName xml.Name              `xml:"ErrorList"`
	Error   []ProductsReturnError `xml:"Error"`
}

type ProductsReturnError struct {
	XmlName              xml.Name `xml:"Error"`
	ErrorCode            int      `xml:"ErrorCode"`
	ErrorDescription     string   `xml:"ErrorDescription"`
	ErrorLongDescription string   `xml:"ErrorLongDescription"`
}

// UpdateProductsBody is to send the body data
type UpdateProductsBody struct {
	Request UpdateProductsRequest `xml:"Request"`
}

type UpdateProductsRequest struct {
	AfterbuyGlobal AfterbuyGlobal         `xml:"AfterbuyGlobal"`
	Products       UpdateProductsProducts `xml:"Products"`
}

type UpdateProductsProducts struct {
	XmlName xml.Name                `xml:"Products"`
	Product []UpdateProductsProduct `xml:"Product"`
}

type UpdateProductsProduct struct {
	XmlName         xml.Name                   `xml:"Product"`
	ProductIdent    UpdateProductsProductIdent `xml:"ProductIdent"`
	Quantity        int                        `xml:"Quantity"`
	AuctionQuantity int                        `xml:"AuctionQuantity"`
	Stock           int                        `xml:"Stock"`
	SellingPrice    string                     `xml:"SellingPrice"`
	BuyingPrice     string                     `xml:"BuyingPrice"`
	DealerPrice     string                     `xml:"DealerPrice"`
}

type UpdateProductsProductIdent struct {
	XmlName       xml.Name `xml:"ProductIdent"`
	ProductInsert int      `xml:"ProductInsert"`
	ProductID     int      `xml:"ProductID"`
	Anr           int      `xml:"Anr"`
	EAN           string   `xml:"EAN"`
}

// UpdateProductsReturn is to decode the xml data
type UpdateProductsReturn struct {
	XmlName    xml.Name                   `xml:"Afterbuy"`
	CallStatus string                     `xml:"CallStatus"`
	CallName   string                     `xml:"CallName"`
	VersionID  int                        `xml:"VersionID"`
	Result     UpdateProductsReturnResult `xml:"Result"`
}

type UpdateProductsReturnResult struct {
	XmlName     xml.Name                        `xml:"Result"`
	WarningList UpdateProductsReturnWarningList `xml:"WarningList"`
}

type UpdateProductsReturnWarningList struct {
	XmlName xml.Name                      `xml:"WarningList"`
	Warning []UpdateProductsReturnWarning `xml:"Warning"`
}

type UpdateProductsReturnWarning struct {
	XmlName                xml.Name `xml:"Warning"`
	WarningCode            int      `xml:"WarningCode"`
	WarningDescription     string   `xml:"WarningDescription"`
	WarningLongDescription string   `xml:"WarningLongDescription"`
	UserProductId          string   `xml:"UserProductId"`
}

// Products are to get all products from ab interface
func Products(body ProductsBody) (ProductsReturn, error) {

	// Convert body
	convert, err := xml.Marshal(body)
	if err != nil {
		return ProductsReturn{}, err
	}

	// Set config for request
	c := Config{
		BaseUrl: "https://api.afterbuy.de/afterbuy/ABInterface.aspx",
		Method:  "GET",
		Body:    bytes.NewBuffer(convert),
		Header:  map[string]string{},
	}

	// Add header to config
	header := make(map[string]string)
	header["Content-Type"] = "application/xml"
	c.Header = header

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

// UpdateProducts are to update a product from ab interface
func UpdateProducts(body UpdateProductsBody) (UpdateProductsReturn, error) {

	// Convert body
	convert, err := xml.Marshal(body)
	if err != nil {
		return UpdateProductsReturn{}, err
	}

	// Set config for request
	c := Config{
		BaseUrl: "https://api.afterbuy.de/afterbuy/ABInterface.aspx",
		Method:  "GET",
		Body:    bytes.NewBuffer(convert),
		Header:  map[string]string{},
	}

	// Add header to config
	header := make(map[string]string)
	header["Content-Type"] = "application/xml"
	c.Header = header

	// Send new request
	response, err := c.Send()
	if err != nil {
		return UpdateProductsReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Decode data
	var decode UpdateProductsReturn

	err = xml.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return UpdateProductsReturn{}, err
	}

	// Return data
	return decode, nil

}
