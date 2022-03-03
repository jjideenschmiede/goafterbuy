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

// CatalogsBody is to structure the xml data
type CatalogsBody struct {
	Request CatalogsRequest `xml:"Request"`
}

type CatalogsRequest struct {
	AfterbuyGlobal AfterbuyGlobal             `xml:"AfterbuyGlobal"`
	MaxCatalogs    int                        `xml:"MaxCatalogs"`
	DataFilter     *CatalogsRequestDataFilter `xml:"DataFilter"`
}

type CatalogsRequestDataFilter struct {
	Filter []CatalogsRequestFilter `xml:"Filter"`
}

type CatalogsRequestFilter struct {
	FilterName   string                      `xml:"FilterName"`
	FilterValues CatalogsRequestFilterValues `xml:"FilterValues"`
}

type CatalogsRequestFilterValues struct {
	LevelFrom   int      `xml:"LevelFrom,omitempty"`
	LevelTo     int      `xml:"LevelTo,omitempty"`
	ValueFrom   int      `xml:"ValueFrom,omitempty"`
	ValueTo     int      `xml:"ValueTo,omitempty"`
	DateFrom    string   `xml:"DateFrom,omitempty"`
	DateTo      string   `xml:"DateTo,omitempty"`
	FilterValue []string `xml:"FilterValue,omitempty"`
}

// CatalogsReturn is to decode the xml return data
type CatalogsReturn struct {
	XmlName    xml.Name             `xml:"Afterbuy"`
	CallStatus string               `xml:"CallStatus"`
	CallName   string               `xml:"CallName"`
	VersionId  int                  `xml:"VersionID"`
	Result     CatalogsReturnResult `xml:"Result"`
}

type CatalogsReturnResult struct {
	XmlName         xml.Name               `xml:"Result"`
	HasMoreCatalogs int                    `xml:"HasMoreCatalogs"`
	Catalogs        CatalogsReturnCatalogs `xml:"Catalogs"`
	LastCatalogId   int                    `xml:"LastCatalogID"`
}

type CatalogsReturnCatalogs struct {
	XmlName xml.Name                `xml:"Catalogs"`
	Catalog []CatalogsReturnCatalog `xml:"Catalog"`
}

type CatalogsReturnCatalog struct {
	XmlName        xml.Name `xml:"Catalog"`
	CatalogId      int      `xml:"CatalogID"`
	Name           string   `xml:"Name"`
	Description    string   `xml:"Description"`
	ParentId       int      `xml:"ParentID"`
	Level          int      `xml:"Level"`
	Position       int      `xml:"Position"`
	AdditionalText string   `xml:"AdditionalText"`
	Show           int      `xml:"Show"`
	URL            string   `xml:"URL"`
	Picture1       string   `xml:"Picture1"`
	Picture2       string   `xml:"Picture2"`
	TitlePicture   string   `xml:"TitlePicture"`
}

// Catalogs is to get all catalogs from ab interface
func Catalogs(body CatalogsBody) (CatalogsReturn, error) {

	// Convert body
	convert, err := xml.Marshal(body)
	if err != nil {
		return CatalogsReturn{}, err
	}

	// Config new request
	c := Config{nil, convert}

	// Send new request
	response, err := c.Send()
	if err != nil {
		return CatalogsReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Decode data
	var decode CatalogsReturn

	err = xml.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return CatalogsReturn{}, err
	}

	// Return data
	return decode, nil

}
