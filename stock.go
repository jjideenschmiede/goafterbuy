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

// StockBody is to send the body data
type StockBody struct {
	Request StockRequest `xml:"Request"`
}

type StockRequest struct {
	AfterbuyGlobal AfterbuyGlobal `xml:"AfterbuyGlobal"`
	Products       StockProducts  `xml:"Products"`
}

type StockProducts struct {
	Product StockProduct `xml:"Product"`
}

type StockProduct struct {
	ProductId int `xml:"ProductID"`
}

// StockReturn is to decode xml return
type StockReturn struct {
	XmlName    xml.Name          `xml:"Afterbuy"`
	CallStatus string            `xml:"CallStatus"`
	CallName   string            `xml:"CallName"`
	VersionId  int               `xml:"VersionID"`
	Result     StockReturnResult `xml:"Result"`
}

type StockReturnResult struct {
	XmlName     xml.Name               `xml:"Result"`
	Products    StockReturnProducts    `xml:"Products"`
	WarningList StockReturnWarningList `xml:"WarningList,omitempty"`
}

type StockReturnProducts struct {
	XmlName xml.Name             `xml:"Products"`
	Product []StockReturnProduct `xml:"Product"`
}

type StockReturnProduct struct {
	ProductId           int    `xml:"ProductID"`
	Name                string `xml:"Name"`
	Anr                 int    `xml:"Anr"`
	Ean                 string `xml:"EAN"`
	AuctionQuantity     int    `xml:"AuctionQuantity"`
	Quantity            int    `xml:"Quantity"`
	FullFilmentQuantity int    `xml:"FullFilmentQuantity"`
	MinimumStock        int    `xml:"MinimumStock"`
	Level               int    `xml:"Level"`
	Discontinued        int    `xml:"Discontinued"`
	Stock               int    `xml:"Stock"`
	MergeStock          int    `xml:"MergeStock"`
	AvailableShop       int    `xml:"AvailableShop"`
}

type StockReturnWarningList struct {
	XmlName xml.Name             `xml:"WarningList"`
	Warning []StockReturnWarning `xml:"Warning"`
}

type StockReturnWarning struct {
	WarningCode            int    `xml:"WarningCode"`
	WarningDescription     string `xml:"WarningDescription"`
	WarningLongDescription string `xml:"WarningLongDescription"`
	ProductId              string `xml:"ProductID"`
}

// Stock is to get a stock of an product from ab interface
func Stock(body *StockBody) (*StockReturn, error) {

	// Convert body
	convert, err := xml.Marshal(body)
	if err != nil {
		return nil, err
	}

	// Config new request
	c := Config{nil, convert}

	// Send new request
	response, err := c.Send()
	if err != nil {
		return nil, err
	}

	// Close request body
	defer response.Body.Close()

	// Decode data
	var decode StockReturn

	err = xml.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return nil, err
	}

	// Return data
	return &decode, nil

}
