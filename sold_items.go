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

// SoldItemsBody is to send the body data
type SoldItemsBody struct {
	Request SoldItemsRequest `xml:"Request"`
}

type SoldItemsRequest struct {
	AfterbuyGlobal    AfterbuyGlobal      `xml:"AfterbuyGlobal"`
	DataFilter        SoldItemsDataFilter `xml:"DataFilter"`
	MaxSoldItems      int                 `xml:"MaxSoldItems"`
	ReturnHiddenItems int                 `xml:"ReturnHiddenItems"`
}

type SoldItemsDataFilter struct {
	Filter []SoldItemsFilter `xml:"Filter"`
}

type SoldItemsFilter struct {
	FilterName   string                `xml:"FilterName"`
	FilterValues SoldItemsFilterValues `xml:"FilterValues"`
}

type SoldItemsFilterValues struct {
	DateFrom    string `xml:"DateFrom,omitempty"`
	DateTo      string `xml:"DateTo,omitempty"`
	ValueFrom   int    `xml:"ValueFrom,omitempty"`
	ValueTo     int    `xml:"ValueTo,omitempty"`
	FilterValue string `xml:"FilterValue,omitempty"`
}

// SoldItemsReturn is to decode the xml data
type SoldItemsReturn struct {
	XmlName    xml.Name              `xml:"Afterbuy"`
	CallStatus string                `xml:"CallStatus"`
	CallName   string                `xml:"CallName"`
	VersionId  int                   `xml:"VersionID"`
	Result     SoldItemsReturnResult `xml:"Result"`
}

type SoldItemsReturnResult struct {
	XmlName      xml.Name              `xml:"Result"`
	HasMoreItems int                   `xml:"HasMoreItems"`
	Orders       SoldItemsReturnOrders `xml:"Orders"`
	OrdersCount  int                   `xml:"OrdersCount"`
	LastOrderId  int                   `xml:"LastOrderID"`
	ItemsCount   int                   `xml:"ItemsCount"`
}

type SoldItemsReturnOrders struct {
	XmlName xml.Name               `xml:"Orders"`
	Order   []SoldItemsReturnOrder `xml:"Order"`
}

type SoldItemsReturnOrder struct {
	XmlName                       xml.Name                             `xml:"Order"`
	InvoiceNumber                 string                               `xml:"InvoiceNumber"`
	OrderId                       int                                  `xml:"OrderID"`
	EbayAccount                   string                               `xml:"EbayAccount"`
	AmazonAccount                 string                               `xml:"AmazonAccount"`
	AlternativeItemNumber1        string                               `xml:"AlternativeItemNumber1"`
	Anr                           int                                  `xml:"Anr"`
	FeedbackDate                  string                               `xml:"FeedbackDate"`
	AdditionalInfo                string                               `xml:"AdditionalInfo"`
	TrackingLink                  string                               `xml:"TrackingLink"`
	Memo                          string                               `xml:"Memo"`
	FeedbackLink                  string                               `xml:"FeedbackLink"`
	OrderDate                     string                               `xml:"OrderDate"`
	OrderIdAlt                    interface{}                          `xml:"OrderIDAlt"`
	PaymentInfo                   SoldItemsReturnPaymentInfo           `xml:"PaymentInfo"`
	IsCheckoutConfirmedByCustomer int                                  `xml:"IsCheckoutConfirmedByCustomer"`
	BuyerInfo                     SoldItemsReturnBuyerInfo             `xml:"BuyerInfo"`
	ContainseBayPlusTransaction   bool                                 `xml:"ContainseBayPlusTransaction"`
	SoldItems                     SoldItemsReturnSoldItems             `xml:"SoldItems"`
	ShippingInfo                  SoldItemsReturnShippingInfo          `xml:"ShippingInfo"`
	VorgangsInfo                  SoldItemsReturnVorgangsInfo          `xml:"VorgangsInfo"`
	OrderOriginalCurrency         SoldItemsReturnOrderOriginalCurrency `xml:"OrderOriginalCurrency"`
}

type SoldItemsReturnPaymentInfo struct {
	XmlName              xml.Name `xml:"PaymentInfo"`
	PaymentId            string   `xml:"PaymentID"`
	PaymentMethod        string   `xml:"PaymentMethod"`
	PaymentFunction      string   `xml:"PaymentFunction"`
	PaymentTransactionId string   `xml:"PaymentTransactionID"`
	PaymentStatus        string   `xml:"PaymentStatus"`
	PaymentDate          string   `xml:"PaymentDate"`
	AlreadyPaid          string   `xml:"AlreadyPaid"`
	FullAmount           string   `xml:"FullAmount"`
	InvoiceDate          string   `xml:"InvoiceDate"`
}

type SoldItemsReturnBuyerInfo struct {
	XmlName         xml.Name                       `xml:"BuyerInfo"`
	BillingAddress  SoldItemsReturnBillingAddress  `xml:"BillingAddress"`
	ShippingAddress SoldItemsReturnShippingAddress `xml:"ShippingAddress"`
}

type SoldItemsReturnBillingAddress struct {
	XmlName           xml.Name `xml:"BillingAddress"`
	AfterbuyUserId    int      `xml:"AfterbuyUserID"`
	AfterbuyUserIdAlt int      `xml:"AfterbuyUserIDAlt"`
	UserIdPlattform   string   `xml:"UserIDPlattform"`
	FirstName         string   `xml:"FirstName"`
	LastName          string   `xml:"LastName"`
	Title             string   `xml:"Title"`
	Company           string   `xml:"Company"`
	Street            string   `xml:"Street"`
	Street2           string   `xml:"Street2"`
	PostalCode        string   `xml:"PostalCode"`
	City              string   `xml:"City"`
	StateOrProvince   string   `xml:"StateOrProvince"`
	Country           string   `xml:"Country"`
	CountryISO        string   `xml:"CountryISO"`
	Phone             string   `xml:"Phone"`
	Fax               string   `xml:"Fax"`
	Mail              string   `xml:"Mail"`
	IsMerchant        int      `xml:"IsMerchant"`
	TaxIdNumber       string   `xml:"TaxIDNumber"`
}

type SoldItemsReturnShippingAddress struct {
	XmlName         xml.Name `xml:"ShippingAddress"`
	FirstName       string   `xml:"FirstName"`
	LastName        string   `xml:"LastName"`
	Company         string   `xml:"Company"`
	Street          string   `xml:"Street"`
	Street2         string   `xml:"Street2"`
	PostalCode      string   `xml:"PostalCode"`
	City            string   `xml:"City"`
	StateOrProvince string   `xml:"StateOrProvince"`
	Phone           string   `xml:"Phone"`
	Country         string   `xml:"Country"`
	CountryISO      string   `xml:"CountryISO"`
	TaxIDNumber     string   `xml:"TaxIDNumber"`
}

type SoldItemsReturnSoldItems struct {
	XmlName      xml.Name                  `xml:"SoldItems"`
	SoldItem     []SoldItemsReturnSoldItem `xml:"SoldItem"`
	ItemsInOrder int                       `xml:"ItemsInOrder"`
}

type SoldItemsReturnSoldItem struct {
	XmlName                 xml.Name                            `xml:"SoldItem"`
	ItemDetailsDone         int                                 `xml:"ItemDetailsDone"`
	ItemId                  int                                 `xml:"ItemID"`
	AlternativeItemNumber1  string                              `xml:"AlternativeItemNumber1"`
	AlternativeItemNumber   string                              `xml:"AlternativeItemNumber"`
	Anr                     int                                 `xml:"Anr"`
	FulfillmentServiceLevel int                                 `xml:"FulfillmentServiceLevel"`
	IsAmazonInvoiced        bool                                `xml:"IsAmazonInvoiced"`
	PlatformSpecificOrderId string                              `xml:"PlatformSpecificOrderId"`
	EBayTransactiondD       int                                 `xml:"eBayTransactionID"`
	EBayPlusTransaction     bool                                `xml:"eBayPlusTransaction"`
	InternalItemType        int                                 `xml:"InternalItemType"`
	UserDefinedFlag         int                                 `xml:"UserDefinedFlag"`
	ItemTitle               string                              `xml:"ItemTitle"`
	ItemQuantity            int                                 `xml:"ItemQuantity"`
	ItemPrice               string                              `xml:"ItemPrice"`
	ItemOriginalCurrency    SoldItemsReturnItemOriginalCurrency `xml:"ItemOriginalCurrency"`
	ItemEndDate             string                              `xml:"ItemEndDate"`
	TaxRate                 string                              `xml:"TaxRate"`
	TaxCollectedBy          int                                 `xml:"TaxCollectedBy"`
	ItemWeight              string                              `xml:"ItemWeight"`
	ItemModDate             string                              `xml:"ItemModDate"`
	ItemPlatformName        string                              `xml:"ItemPlatformName"`
	ItemLink                string                              `xml:"ItemLink"`
	EBayFeedbackCompleted   int                                 `xml:"eBayFeedbackCompleted"`
	EBayFeedbackReceived    int                                 `xml:"eBayFeedbackReceived"`
	ShopProductDetails      SoldItemsReturnShopProductDetails   `xml:"ShopProductDetails"`
	SoldItemAttributes      SoldItemsReturnSoldItemAttributes   `xml:"SoldItemAttributes"`
	Tags                    SoldItemReturnTags                  `xml:"Tags"`
}

type SoldItemsReturnItemOriginalCurrency struct {
	XmlName       xml.Name `xml:"ItemOriginalCurrency"`
	ItemPrice     string   `xml:"ItemPrice"`
	ItemPriceCode string   `xml:"ItemPriceCode"`
	ItemShipping  string   `xml:"ItemShipping"`
}

type SoldItemsReturnShopProductDetails struct {
	XmlName         xml.Name                       `xml:"ShopProductDetails"`
	ProductId       int                            `xml:"ProductID"`
	Name            string                         `xml:"Name"`
	Ean             string                         `xml:"EAN"`
	Anr             int                            `xml:"Anr"`
	BasepriceFactor string                         `xml:"BasepriceFactor"`
	BaseProductData SoldItemsReturnBaseProductData `xml:"BaseProductData"`
}

type SoldItemsReturnBaseProductData struct {
	XmlName         xml.Name                    `xml:"BaseProductData"`
	BaseProductType int                         `xml:"BaseProductType"`
	ChildProduct    SoldItemsReturnChildProduct `xml:"ChildProduct"`
}

type SoldItemsReturnChildProduct struct {
	XmlName          xml.Name `xml:"ChildProduct"`
	ProductId        int      `xml:"ProductID"`
	ProductANr       int      `xml:"ProductANr"`
	ProductName      string   `xml:"ProductName"`
	ProductQuantity  int      `xml:"ProductQuantity"`
	ProductVat       int      `xml:"ProductVAT"`
	ProductWeight    string   `xml:"ProductWeight"`
	ProductUnitPrice string   `xml:"ProductUnitPrice"`
}

type SoldItemsReturnSoldItemAttributes struct {
	XmlName           xml.Name                           `xml:"SoldItemAttributes"`
	SoldItemAttribute []SoldItemsReturnSoldItemAttribute `xml:"SoldItemAttribute"`
}

type SoldItemsReturnSoldItemAttribute struct {
	XmlName           xml.Name `xml:"SoldItemAttribute"`
	AttributeName     string   `xml:"AttributeName"`
	AttributeValue    string   `xml:"AttributeValue"`
	AttributePosition int      `xml:"AttributePosition"`
}

type SoldItemsReturnTags struct {
	XmlName xml.Name `xml:"Tags"`
	Tag     []string `xml:"Tag"`
}

type SoldItemsReturnShippingInfo struct {
	XmlName                xml.Name                    `xml:"ShippingInfo"`
	ShippingMethod         string                      `xml:"ShippingMethod"`
	ShippingCost           string                      `xml:"ShippingCost"`
	ShippingAdditionalCost string                      `xml:"ShippingAdditionalCost"`
	ShippingTotalCost      string                      `xml:"ShippingTotalCost"`
	ShippingTaxRate        string                      `xml:"ShippingTaxRate"`
	DeliveryDate           string                      `xml:"DeliveryDate"`
	ParcelLabels           SoldItemsReturnParcelLabels `xml:"ParcelLabels"`
}

type SoldItemsReturnParcelLabels struct {
	XmlName     xml.Name                     `xml:"ParcelLabels"`
	ParcelLabel []SoldItemsReturnParcelLabel `xml:"ParcelLabel"`
}

type SoldItemsReturnParcelLabel struct {
	XmlName           xml.Name `xml:"ParcelLabel"`
	ItemID            int      `xml:"ItemID"`
	PackageNumber     int      `xml:"PackageNumber"`
	ParcelLabelNumber string   `xml:"ParcelLabelNumber"`
}

type SoldItemsReturnVorgangsInfo struct {
	XmlName       xml.Name `xml:"VorgangsInfo"`
	VorgangsInfo1 string   `xml:"VorgangsInfo1"`
	VorgangsInfo2 string   `xml:"VorgangsInfo2"`
	VorgangsInfo3 string   `xml:"VorgangsInfo3"`
}

type SoldItemsReturnOrderOriginalCurrency struct {
	XmlName                 xml.Name `xml:"OriginalCurrency"`
	EbayShippingAmount      string   `xml:"EbayShippingAmount"`
	ShippingAmount          string   `xml:"ShippingAmount"`
	PaymentSurcharge        string   `xml:"PaymentSurcharge"`
	PaymentSurchargePerCent string   `xml:"PaymentSurchargePerCent"`
	InvoiceAmount           string   `xml:"InvoiceAmount"`
	ExchangeRate            string   `xml:"ExchangeRate"`
	PayedAmount             string   `xml:"PayedAmount"`
}

// SoldItemBody is to send the body data
type SoldItemBody struct {
	Request SoldItemRequest `xml:"Request"`
}

type SoldItemRequest struct {
	AfterbuyGlobal    AfterbuyGlobal     `xml:"AfterbuyGlobal"`
	DataFilter        SoldItemDataFilter `xml:"DataFilter"`
	ReturnHiddenItems int                `xml:"ReturnHiddenItems"`
}

type SoldItemDataFilter struct {
	Filter SoldItemFilter `xml:"Filter,omitempty"`
}

type SoldItemFilter struct {
	FilterName   string               `xml:"FilterName,omitempty"`
	FilterValues SoldItemFilterValues `xml:"FilterValues,omitempty"`
}

type SoldItemFilterValues struct {
	FilterValue int `xml:"FilterValue,omitempty"`
}

// SoldItemReturn is to decode the xml data
type SoldItemReturn struct {
	XmlName    xml.Name             `xml:"Afterbuy"`
	CallStatus string               `xml:"CallStatus"`
	CallName   string               `xml:"CallName"`
	VersionId  int                  `xml:"VersionID"`
	Result     SoldItemReturnResult `xml:"Result"`
}

type SoldItemReturnResult struct {
	XmlName      xml.Name             `xml:"Result"`
	HasMoreItems int                  `xml:"HasMoreItems"`
	Orders       SoldItemReturnOrders `xml:"Orders"`
	OrdersCount  int                  `xml:"OrdersCount"`
	LastOrderId  int                  `xml:"LastOrderID"`
	ItemsCount   int                  `xml:"ItemsCount"`
}

type SoldItemReturnOrders struct {
	XmlName xml.Name            `xml:"Orders"`
	Order   SoldItemReturnOrder `xml:"Order"`
}

type SoldItemReturnOrder struct {
	XmlName                       xml.Name                            `xml:"Order"`
	InvoiceNumber                 string                              `xml:"InvoiceNumber"`
	OrderId                       int                                 `xml:"OrderID"`
	EbayAccount                   string                              `xml:"EbayAccount"`
	AmazonAccount                 string                              `xml:"AmazonAccount"`
	AlternativeItemNumber1        string                              `xml:"AlternativeItemNumber1"`
	Anr                           int                                 `xml:"Anr"`
	FeedbackDate                  string                              `xml:"FeedbackDate"`
	AdditionalInfo                string                              `xml:"AdditionalInfo"`
	TrackingLink                  string                              `xml:"TrackingLink"`
	Memo                          string                              `xml:"Memo"`
	FeedbackLink                  string                              `xml:"FeedbackLink"`
	OrderDate                     string                              `xml:"OrderDate"`
	OrderIdAlt                    interface{}                         `xml:"OrderIDAlt"`
	PaymentInfo                   SoldItemReturnPaymentInfo           `xml:"PaymentInfo"`
	IsCheckoutConfirmedByCustomer int                                 `xml:"IsCheckoutConfirmedByCustomer"`
	BuyerInfo                     SoldItemReturnBuyerInfo             `xml:"BuyerInfo"`
	ContainseBayPlusTransaction   bool                                `xml:"ContainseBayPlusTransaction"`
	SoldItems                     SoldItemReturnSoldItems             `xml:"SoldItems"`
	ShippingInfo                  SoldItemReturnShippingInfo          `xml:"ShippingInfo"`
	VorgangsInfo                  SoldItemReturnVorgangsInfo          `xml:"VorgangsInfo"`
	OrderOriginalCurrency         SoldItemReturnOrderOriginalCurrency `xml:"OrderOriginalCurrency"`
}

type SoldItemReturnPaymentInfo struct {
	XmlName              xml.Name `xml:"PaymentInfo"`
	PaymentId            string   `xml:"PaymentID"`
	PaymentMethod        string   `xml:"PaymentMethod"`
	PaymentFunction      string   `xml:"PaymentFunction"`
	PaymentTransactionId string   `xml:"PaymentTransactionID"`
	PaymentStatus        string   `xml:"PaymentStatus"`
	PaymentDate          string   `xml:"PaymentDate"`
	AlreadyPaid          string   `xml:"AlreadyPaid"`
	FullAmount           string   `xml:"FullAmount"`
	InvoiceDate          string   `xml:"InvoiceDate"`
}

type SoldItemReturnBuyerInfo struct {
	XmlName         xml.Name                      `xml:"BuyerInfo"`
	BillingAddress  SoldItemReturnBillingAddress  `xml:"BillingAddress"`
	ShippingAddress SoldItemReturnShippingAddress `xml:"ShippingAddress"`
}

type SoldItemReturnBillingAddress struct {
	XmlName           xml.Name `xml:"BillingAddress"`
	AfterbuyUserId    int      `xml:"AfterbuyUserID"`
	AfterbuyUserIdAlt int      `xml:"AfterbuyUserIDAlt"`
	UserIdPlattform   string   `xml:"UserIDPlattform"`
	FirstName         string   `xml:"FirstName"`
	LastName          string   `xml:"LastName"`
	Title             string   `xml:"Title"`
	Company           string   `xml:"Company"`
	Street            string   `xml:"Street"`
	Street2           string   `xml:"Street2"`
	PostalCode        string   `xml:"PostalCode"`
	City              string   `xml:"City"`
	StateOrProvince   string   `xml:"StateOrProvince"`
	Country           string   `xml:"Country"`
	CountryISO        string   `xml:"CountryISO"`
	Phone             string   `xml:"Phone"`
	Fax               string   `xml:"Fax"`
	Mail              string   `xml:"Mail"`
	IsMerchant        int      `xml:"IsMerchant"`
	TaxIdNumber       string   `xml:"TaxIDNumber"`
}

type SoldItemReturnShippingAddress struct {
	XmlName         xml.Name `xml:"ShippingAddress"`
	FirstName       string   `xml:"FirstName"`
	LastName        string   `xml:"LastName"`
	Company         string   `xml:"Company"`
	Street          string   `xml:"Street"`
	Street2         string   `xml:"Street2"`
	PostalCode      string   `xml:"PostalCode"`
	City            string   `xml:"City"`
	StateOrProvince string   `xml:"StateOrProvince"`
	Phone           string   `xml:"Phone"`
	Country         string   `xml:"Country"`
	CountryISO      string   `xml:"CountryISO"`
	TaxIDNumber     string   `xml:"TaxIDNumber"`
}

type SoldItemReturnSoldItems struct {
	XmlName      xml.Name                 `xml:"SoldItems"`
	SoldItem     []SoldItemReturnSoldItem `xml:"SoldItem"`
	ItemsInOrder int                      `xml:"ItemsInOrder"`
}

type SoldItemReturnSoldItem struct {
	XmlName                 xml.Name                           `xml:"SoldItem"`
	ItemDetailsDone         int                                `xml:"ItemDetailsDone"`
	ItemId                  int                                `xml:"ItemID"`
	AlternativeItemNumber1  string                             `xml:"AlternativeItemNumber1"`
	AlternativeItemNumber   string                             `xml:"AlternativeItemNumber"`
	Anr                     int                                `xml:"Anr"`
	FulfillmentServiceLevel int                                `xml:"FulfillmentServiceLevel"`
	IsAmazonInvoiced        bool                               `xml:"IsAmazonInvoiced"`
	PlatformSpecificOrderId string                             `xml:"PlatformSpecificOrderId"`
	EBayTransactionId       int                                `xml:"eBayTransactionID"`
	EBayPlusTransaction     bool                               `xml:"eBayPlusTransaction"`
	InternalItemType        int                                `xml:"InternalItemType"`
	UserDefinedFlag         int                                `xml:"UserDefinedFlag"`
	ItemTitle               string                             `xml:"ItemTitle"`
	ItemQuantity            int                                `xml:"ItemQuantity"`
	ItemPrice               string                             `xml:"ItemPrice"`
	ItemOriginalCurrency    SoldItemReturnItemOriginalCurrency `xml:"ItemOriginalCurrency"`
	ItemEndDate             string                             `xml:"ItemEndDate"`
	TaxRate                 string                             `xml:"TaxRate"`
	TaxCollectedBy          int                                `xml:"TaxCollectedBy"`
	ItemWeight              string                             `xml:"ItemWeight"`
	ItemModDate             string                             `xml:"ItemModDate"`
	ItemPlatformName        string                             `xml:"ItemPlatformName"`
	ItemLink                string                             `xml:"ItemLink"`
	EBayFeedbackCompleted   int                                `xml:"eBayFeedbackCompleted"`
	EBayFeedbackReceived    int                                `xml:"eBayFeedbackReceived"`
	ShopProductDetails      SoldItemReturnShopProductDetails   `xml:"ShopProductDetails"`
	SoldItemAttributes      SoldItemReturnSoldItemAttributes   `xml:"SoldItemAttributes"`
	Tags                    SoldItemReturnTags                 `xml:"Tags"`
}

type SoldItemReturnItemOriginalCurrency struct {
	XmlName       xml.Name `xml:"ItemOriginalCurrency"`
	ItemPrice     string   `xml:"ItemPrice"`
	ItemPriceCode string   `xml:"ItemPriceCode"`
	ItemShipping  string   `xml:"ItemShipping"`
}

type SoldItemReturnShopProductDetails struct {
	XmlName         xml.Name                      `xml:"ShopProductDetails"`
	ProductId       int                           `xml:"ProductID"`
	Name            string                        `xml:"Name"`
	Ean             string                        `xml:"EAN"`
	Anr             int                           `xml:"Anr"`
	BasepriceFactor string                        `xml:"BasepriceFactor"`
	BaseProductData SoldItemReturnBaseProductData `xml:"BaseProductData"`
}

type SoldItemReturnBaseProductData struct {
	XmlName         xml.Name                   `xml:"BaseProductData"`
	BaseProductType int                        `xml:"BaseProductType"`
	ChildProduct    SoldItemReturnChildProduct `xml:"ChildProduct"`
}

type SoldItemReturnChildProduct struct {
	XmlName          xml.Name `xml:"ChildProduct"`
	ProductId        int      `xml:"ProductID"`
	ProductANr       int      `xml:"ProductANr"`
	ProductName      string   `xml:"ProductName"`
	ProductQuantity  int      `xml:"ProductQuantity"`
	ProductVat       int      `xml:"ProductVAT"`
	ProductWeight    string   `xml:"ProductWeight"`
	ProductUnitPrice string   `xml:"ProductUnitPrice"`
}

type SoldItemReturnSoldItemAttributes struct {
	XmlName           xml.Name                          `xml:"SoldItemAttributes"`
	SoldItemAttribute []SoldItemReturnSoldItemAttribute `xml:"SoldItemAttribute"`
}

type SoldItemReturnSoldItemAttribute struct {
	XmlName           xml.Name `xml:"SoldItemAttribute"`
	AttributeName     string   `xml:"AttributeName"`
	AttributeValue    string   `xml:"AttributeValue"`
	AttributePosition int      `xml:"AttributePosition"`
}

type SoldItemReturnTags struct {
	XmlName xml.Name `xml:"Tags"`
	Tag     []string `xml:"Tag"`
}

type SoldItemReturnShippingInfo struct {
	XmlName                xml.Name                   `xml:"ShippingInfo"`
	ShippingMethod         string                     `xml:"ShippingMethod"`
	ShippingCost           string                     `xml:"ShippingCost"`
	ShippingAdditionalCost string                     `xml:"ShippingAdditionalCost"`
	ShippingTotalCost      string                     `xml:"ShippingTotalCost"`
	ShippingTaxRate        string                     `xml:"ShippingTaxRate"`
	DeliveryDate           string                     `xml:"DeliveryDate"`
	ParcelLabels           SoldItemReturnParcelLabels `xml:"ParcelLabels"`
}

type SoldItemReturnParcelLabels struct {
	XmlName     xml.Name                    `xml:"ParcelLabels"`
	ParcelLabel []SoldItemReturnParcelLabel `xml:"ParcelLabel"`
}

type SoldItemReturnParcelLabel struct {
	XmlName           xml.Name `xml:"ParcelLabel"`
	ItemID            int      `xml:"ItemID"`
	PackageNumber     int      `xml:"PackageNumber"`
	ParcelLabelNumber string   `xml:"ParcelLabelNumber"`
}

type SoldItemReturnVorgangsInfo struct {
	XmlName       xml.Name `xml:"VorgangsInfo"`
	VorgangsInfo1 string   `xml:"VorgangsInfo1"`
	VorgangsInfo2 string   `xml:"VorgangsInfo2"`
	VorgangsInfo3 string   `xml:"VorgangsInfo3"`
}

type SoldItemReturnOrderOriginalCurrency struct {
	XmlName                 xml.Name `xml:"OriginalCurrency"`
	EbayShippingAmount      string   `xml:"EbayShippingAmount"`
	ShippingAmount          string   `xml:"ShippingAmount"`
	PaymentSurcharge        string   `xml:"PaymentSurcharge"`
	PaymentSurchargePerCent string   `xml:"PaymentSurchargePerCent"`
	InvoiceAmount           string   `xml:"InvoiceAmount"`
	ExchangeRate            string   `xml:"ExchangeRate"`
	PayedAmount             string   `xml:"PayedAmount"`
}

// UpdateSoldItemBody is to send the body data
type UpdateSoldItemBody struct {
	Request UpdateSoldItemBodyRequest `xml:"Request"`
}

type UpdateSoldItemBodyRequest struct {
	AfterbuyGlobal AfterbuyGlobal           `xml:"AfterbuyGlobal"`
	Orders         UpdateSoldItemBodyOrders `xml:"Orders"`
}

type UpdateSoldItemBodyOrders struct {
	Order UpdateSoldItemBodyOrder `xml:"Order"`
}

type UpdateSoldItemBodyOrder struct {
	OrderId          int                             `xml:"OrderID,omitempty"`
	ItemId           int                             `xml:"ItemID,omitempty"`
	AdditionalInfo   string                          `xml:"AdditionalInfo,omitempty"`
	MailDate         string                          `xml:"MailDate,omitempty"`
	ReminderMailDate string                          `xml:"ReminderMailDate,omitempty"`
	UserComment      string                          `xml:"UserComment,omitempty"`
	OrderMemo        string                          `xml:"OrderMemo,omitempty"`
	InvoiceMemo      string                          `xml:"InvoiceMemo,omitempty"`
	InvoiceNumber    int                             `xml:"InvoiceNumber,omitempty"`
	OrderExported    int                             `xml:"OrderExported,omitempty"`
	InvoiceDate      string                          `xml:"InvoiceDate,omitempty"`
	HideOrder        int                             `xml:"HideOrder,omitempty"`
	Reminder1Date    string                          `xml:"Reminder1Date,omitempty"`
	Reminder2Date    string                          `xml:"Reminder2Date,omitempty"`
	XmlDate          string                          `xml:"XmlDate,omitempty"`
	BuyerInfo        *UpdateSoldItemBodyBuyerInfo    `xml:"BuyerInfo,omitempty"`
	PaymentInfo      *UpdateSoldItemBodyPaymentInfo  `xml:"PaymentInfo,omitempty"`
	ShippingInfo     *UpdateSoldItemBodyShippingInfo `xml:"ShippingInfo,omitempty"`
	VorgangsInfo     *UpdateSoldItemBodyVorgangsInfo `xml:"VorgangsInfo,omitempty"`
	Tags             []*UpdateSoldItemBodyTags       `xml:"Tags,omitempty"`
}

type UpdateSoldItemBodyBuyerInfo struct {
	ShippingAddress UpdateSoldItemBodyShippingAddress `xml:"ShippingAddress,omitempty"`
}

type UpdateSoldItemBodyShippingAddress struct {
	FirstName  string `xml:"FirstName,omitempty"`
	LastName   string `xml:"LastName,omitempty"`
	Company    string `xml:"Company,omitempty"`
	Street     string `xml:"Street,omitempty"`
	PostalCode string `xml:"PostalCode,omitempty"`
	City       string `xml:"City,omitempty"`
	Country    string `xml:"Country,omitempty"`
}

type UpdateSoldItemBodyPaymentInfo struct {
	PaymentId            string `xml:"PaymentID,omitempty"`
	PaymentMethod        string `xml:"PaymentMethod,omitempty"`
	PaymentFunction      string `xml:"PaymentFunction,omitempty"`
	PaymentTransactionId string `xml:"PaymentTransactionID,omitempty"`
	PaymentStatus        string `xml:"PaymentStatus,omitempty"`
	PaymentDate          string `xml:"PaymentDate,omitempty"`
	AlreadyPaid          string `xml:"AlreadyPaid,omitempty"`
	FullAmount           string `xml:"FullAmount,omitempty"`
	InvoiceDate          string `xml:"InvoiceDate,omitempty"`
}

type UpdateSoldItemBodyShippingInfo struct {
	ShippingMethod   string `xml:"ShippingMethod,omitempty"`
	ShippingGroup    string `xml:"ShippingGroup,omitempty"`
	ShippingCost     string `xml:"ShippingCost,omitempty"`
	DeliveryDate     string `xml:"DeliveryDate,omitempty"`
	EBayShippingCost string `xml:"eBayShippingCost,omitempty"`
}

type UpdateSoldItemBodyVorgangsInfo struct {
	VorgangsInfo1 string `xml:"VorgangsInfo1,omitempty"`
	VorgangsInfo2 string `xml:"VorgangsInfo2,omitempty"`
	VorgangsInfo3 string `xml:"VorgangsInfo3,omitempty"`
}

type UpdateSoldItemBodyTags struct {
	Tag string `xml:"Tag,omitempty"`
}

// UpdateSoldItemReturn is to decode the xml return
type UpdateSoldItemReturn struct {
	XmlName    xml.Name                   `xml:"Afterbuy"`
	CallStatus string                     `xml:"CallStatus"`
	CallName   string                     `xml:"CallName"`
	VersionId  int                        `xml:"VersionID"`
	Result     UpdateSoldItemReturnResult `xml:"Result,omitempty"`
}

type UpdateSoldItemReturnResult struct {
	XmlName   xml.Name                      `xml:"Result"`
	ErrorList UpdateSoldItemReturnErrorList `xml:"ErrorList"`
}

type UpdateSoldItemReturnErrorList struct {
	XmlName xml.Name                    `xml:"ErrorList"`
	Error   []UpdateSoldItemReturnError `xml:"Error"`
}

type UpdateSoldItemReturnError struct {
	XmlName              xml.Name `xml:"Error"`
	ErrorCode            int      `xml:"ErrorCode"`
	ErrorDescription     string   `xml:"ErrorDescription"`
	ErrorLongDescription string   `xml:"ErrorLongDescription"`
}

// SoldItems is to get all sold items
func SoldItems(body SoldItemsBody) (SoldItemsReturn, error) {

	// Convert body
	convert, err := xml.Marshal(body)
	if err != nil {
		return SoldItemsReturn{}, err
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
		return SoldItemsReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Decode data
	var decode SoldItemsReturn

	err = xml.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return SoldItemsReturn{}, err
	}

	// Return data
	return decode, nil

}

// SoldItem is to get a sold item by id
func SoldItem(body SoldItemBody) (SoldItemReturn, error) {

	// Convert body
	convert, err := xml.Marshal(body)
	if err != nil {
		return SoldItemReturn{}, err
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
		return SoldItemReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Decode data
	var decode SoldItemReturn

	err = xml.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return SoldItemReturn{}, err
	}

	// Return data
	return decode, nil

}

// UpdateSoldItem is to update a sold item
func UpdateSoldItem(body UpdateSoldItemBody) (UpdateSoldItemReturn, error) {

	// Convert body
	convert, err := xml.Marshal(body)
	if err != nil {
		return UpdateSoldItemReturn{}, err
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
		return UpdateSoldItemReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Decode data
	var decode UpdateSoldItemReturn

	err = xml.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return UpdateSoldItemReturn{}, err
	}

	// Return data
	return decode, nil

}
