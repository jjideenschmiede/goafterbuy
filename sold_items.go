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

import "encoding/xml"

// SoldItemsBody is to send the body data
type SoldItemsBody struct {
	Request SoldItemsRequest `xml:"Request"`
}

type SoldItemsRequest struct {
	AfterbuyGlobal AfterbuyGlobal      `xml:"AfterbuyGlobal"`
	DataFilter     SoldItemsDataFilter `xml:"DataFilter"`
	MaxSoldItems   int                 `xml:"MaxSoldItems"`
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
	FilterValue string `xml:"FilterValue,omitempty"`
	ValueFrom   int    `xml:"ValueFrom,omitempty"`
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
	InvoiceNumber                 int                                  `xml:"InvoiceNumber"`
	OrderId                       int                                  `xml:"OrderID"`
	EbayAccount                   string                               `xml:"EbayAccount"`
	AmazonAccount                 string                               `xml:"AmazonAccount"`
	AlternativeItemNumber1        string                               `xml:"AlternativeItemNumber1"`
	Anr                           int                                  `xml:"Anr"`
	FeedbackDate                  string                               `xml:"FeedbackDate"`
	AdditionalInfo                string                               `xml:"AdditionalInfo"`
	TrackingLink                  string                               `xml:"TrackingLink"`
	FeedbackLink                  string                               `xml:"FeedbackLink"`
	OrderDate                     string                               `xml:"OrderDate"`
	OrderIdAlt                    interface{}                          `xml:"OrderIDAlt"`
	PaymentInfo                   SoldItemsReturnPaymentInfo           `xml:"PaymentInfo"`
	IsCheckoutConfirmedByCustomer int                                  `xml:"IsCheckoutConfirmedByCustomer"`
	BuyerInfo                     SoldItemsReturnBuyerInfo             `xml:"BuyerInfo"`
	ContainseBayPlusTransaction   bool                                 `xml:"ContainseBayPlusTransaction"`
	SoldItems                     SoldItemsReturnSoldItems             `xml:"SoldItems"`
	ShippingInfo                  SoldItemsReturnShippingInfo          `xml:"ShippingInfo"`
	OrderOriginalCurrency         SoldItemsReturnOrderOriginalCurrency `xml:"OrderOriginalCurrency"`
}

type SoldItemsReturnPaymentInfo struct {
	XmlName         xml.Name `xml:"PaymentInfo"`
	PaymentId       string   `xml:"PaymentID"`
	PaymentMethod   string   `xml:"PaymentMethod"`
	PaymentFunction string   `xml:"PaymentFunction"`
	PaymentDate     string   `xml:"PaymentDate"`
	AlreadyPaid     string   `xml:"AlreadyPaid"`
	FullAmount      string   `xml:"FullAmount"`
	InvoiceDate     string   `xml:"InvoiceDate"`
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
	TaxCollectedBy          int                                 `xml:"tax_collected_by"`
	ItemWeight              string                              `xml:"item_weight"`
	ItemModDate             string                              `xml:"ItemModDate"`
	ItemPlatformName        string                              `xml:"ItemPlatformName"`
	ItemLink                string                              `xml:"ItemLink"`
	EBayFeedbackCompleted   int                                 `xml:"eBayFeedbackCompleted"`
	EBayFeedbackReceived    int                                 `xml:"eBayFeedbackReceived"`
	ShopProductDetails      SoldItemsReturnShopProductDetails   `xml:"ShopProductDetails"`
}

type SoldItemsReturnItemOriginalCurrency struct {
	XmlName       xml.Name `xml:"ItemOriginalCurrency"`
	ItemPrice     string   `xml:"ItemPrice"`
	ItemPriceCode string   `xml:"ItemPriceCode"`
	ItemShipping  string   `xml:"ItemShipping"`
}

type SoldItemsReturnShopProductDetails struct {
	ProductId       int                            `xml:"product_id"`
	Name            string                         `xml:"Name"`
	Ean             string                         `xml:"EAN"`
	Anr             int                            `xml:"Anr"`
	BasepriceFactor int                            `xml:"BasepriceFactor"`
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

type SoldItemsReturnShippingInfo struct {
	XmlName                xml.Name `xml:"ShippingInfo"`
	ShippingMethod         string   `xml:"ShippingMethod"`
	ShippingCost           string   `xml:"ShippingCost"`
	ShippingAdditionalCost string   `xml:"ShippingAdditionalCost"`
	ShippingTotalCost      string   `xml:"ShippingTotalCost"`
	ShippingTaxRate        string   `xml:"ShippingTaxRate"`
	DeliveryDate           string   `xml:"DeliveryDate"`
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
	AfterbuyGlobal AfterbuyGlobal     `xml:"AfterbuyGlobal"`
	DataFilter     SoldItemDataFilter `xml:"DataFilter"`
}

type SoldItemDataFilter struct {
	Filter SoldItemDataFilterFilter `xml:"Filter,omitempty"`
}

type SoldItemDataFilterFilter struct {
	FilterName   string                         `xml:"FilterName,omitempty"`
	FilterValues SoldItemDataFilterFilterValues `xml:"FilterValues,omitempty"`
}

type SoldItemDataFilterFilterValues struct {
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
	InvoiceNumber                 int                                 `xml:"InvoiceNumber"`
	OrderId                       int                                 `xml:"OrderID"`
	EbayAccount                   string                              `xml:"EbayAccount"`
	AmazonAccount                 string                              `xml:"AmazonAccount"`
	AlternativeItemNumber1        string                              `xml:"AlternativeItemNumber1"`
	Anr                           int                                 `xml:"Anr"`
	FeedbackDate                  string                              `xml:"FeedbackDate"`
	AdditionalInfo                string                              `xml:"AdditionalInfo"`
	TrackingLink                  string                              `xml:"TrackingLink"`
	FeedbackLink                  string                              `xml:"FeedbackLink"`
	OrderDate                     string                              `xml:"OrderDate"`
	OrderIdAlt                    interface{}                         `xml:"OrderIDAlt"`
	PaymentInfo                   SoldItemReturnPaymentInfo           `xml:"PaymentInfo"`
	IsCheckoutConfirmedByCustomer int                                 `xml:"IsCheckoutConfirmedByCustomer"`
	BuyerInfo                     SoldItemReturnBuyerInfo             `xml:"BuyerInfo"`
	ContainseBayPlusTransaction   bool                                `xml:"ContainseBayPlusTransaction"`
	SoldItems                     SoldItemReturnSoldItems             `xml:"SoldItems"`
	ShippingInfo                  SoldItemReturnShippingInfo          `xml:"ShippingInfo"`
	OrderOriginalCurrency         SoldItemReturnOrderOriginalCurrency `xml:"OrderOriginalCurrency"`
}

type SoldItemReturnPaymentInfo struct {
	XmlName         xml.Name `xml:"PaymentInfo"`
	PaymentId       string   `xml:"PaymentID"`
	PaymentMethod   string   `xml:"PaymentMethod"`
	PaymentFunction string   `xml:"PaymentFunction"`
	PaymentDate     string   `xml:"PaymentDate"`
	AlreadyPaid     string   `xml:"AlreadyPaid"`
	FullAmount      string   `xml:"FullAmount"`
	InvoiceDate     string   `xml:"InvoiceDate"`
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
	BasepriceFactor int                           `xml:"BasepriceFactor"`
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

type SoldItemReturnShippingInfo struct {
	XmlName                xml.Name `xml:"ShippingInfo"`
	ShippingMethod         string   `xml:"ShippingMethod"`
	ShippingCost           string   `xml:"ShippingCost"`
	ShippingAdditionalCost string   `xml:"ShippingAdditionalCost"`
	ShippingTotalCost      string   `xml:"ShippingTotalCost"`
	ShippingTaxRate        string   `xml:"ShippingTaxRate"`
	DeliveryDate           string   `xml:"DeliveryDate"`
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
	PaymentMethod          string `xml:"PaymentMethod,omitempty"`
	PaymentDate            string `xml:"PaymentDate,omitempty"`
	AlreadyPaid            string `xml:"AlreadyPaid,omitempty"`
	PaymentAadditionalCost string `xml:"PaymentAadditionalCost,omitempty"`
}

type UpdateSoldItemBodyShippingInfo struct {
	ShippingMethod   string `xml:"ShippingMethod,omitempty"`
	ShippingGroup    string `xml:"ShippingGroup,omitempty"`
	ShippingCost     string `xml:"ShippingCost,omitempty"`
	DeliveryDate     string `xml:"DeliveryDate,omitempty"`
	EBayShippingCost string `xml:"eBayShippingCost,omitempty"`
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

	// Config new request
	c := Config{nil, convert}

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

	// Config new request
	c := Config{nil, convert}

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

	// Config new request
	c := Config{nil, convert}

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
