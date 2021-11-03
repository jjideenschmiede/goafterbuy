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

// ShippingServicesBody is to send the body data
type ShippingServicesBody struct {
	Request ShippingServicesRequest `xml:"Request"`
}

type ShippingServicesRequest struct {
	AfterbuyGlobal AfterbuyGlobal `xml:"AfterbuyGlobal"`
}

// ShippingServicesReturn is to decode the xml data
type ShippingServicesReturn struct {
	XmlName    xml.Name               `xml:"Afterbuy"`
	CallStatus string                 `xml:"CallStatus"`
	CallName   string                 `xml:"CallName"`
	VersionId  int                    `xml:"VersionID"`
	Result     ShippingServicesResult `xml:"Result"`
}

type ShippingServicesResult struct {
	XmlName          xml.Name                         `xml:"Result"`
	ShippingServices ShippingServicesShippingServices `xml:"ShippingServices"`
}

type ShippingServicesShippingServices struct {
	XmlName              xml.Name                          `xml:"ShippingServices"`
	DefaultShippingGroup string                            `xml:"DefaultShippingGroup"`
	ShippingService      []ShippingServicesShippingService `xml:"ShippingService"`
}

type ShippingServicesShippingService struct {
	XmlName         xml.Name                        `xml:"ShippingService"`
	Name            string                          `xml:"Name"`
	DisplayArea     string                          `xml:"DisplayArea"`
	GroupPrio       int                             `xml:"GroupPrio"`
	ShippingMethods ShippingServicesShippingMethods `xml:"ShippingMethods"`
}

type ShippingServicesShippingMethods struct {
	XmlName        xml.Name                         `xml:"ShippingMethods"`
	ShippingMethod []ShippingServicesShippingMethod `xml:"ShippingMethod"`
}

type ShippingServicesShippingMethod struct {
	XmlName               xml.Name                          `xml:"ShippingMethod"`
	ShippingMethodId      int                               `xml:"ShippingMethodID"`
	Name                  string                            `xml:"Name"`
	CountryGroup          string                            `xml:"CountryGroup"`
	CountryGroupCountries string                            `xml:"CountryGroupCountries"`
	Level                 int                               `xml:"Level"`
	TaxRate               int                               `xml:"TaxRate"`
	Priority              int                               `xml:"Priority"`
	PriceFrom             string                            `xml:"PriceFrom"`
	PriceTo               string                            `xml:"PriceTo"`
	IslandAdditionalCosts string                            `xml:"IslandAdditionalCosts"`
	FreeShippingPriceFrom string                            `xml:"FreeShippingPriceFrom"`
	AdditionalItemCosts   string                            `xml:"AdditionalItemCosts"`
	WeightDefinitions     ShippingServicesWeightDefinitions `xml:"WeightDefinitions"`
}

type ShippingServicesWeightDefinitions struct {
	WeightFrom string `xml:"WeightFrom"`
	WeightTo   string `xml:"WeightTo"`
	Price      string `xml:"Price"`
}

// ShippingServices is to check all shipping services
func ShippingServices(body ShippingServicesBody) (ShippingServicesReturn, error) {

	// Convert body
	convert, err := xml.Marshal(body)
	if err != nil {
		return ShippingServicesReturn{}, err
	}

	// Config new request
	c := Config{nil, convert}

	// Send new request
	response, err := c.Send()
	if err != nil {
		return ShippingServicesReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Decode data
	var decode ShippingServicesReturn

	err = xml.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ShippingServicesReturn{}, err
	}

	// Return data
	return decode, nil

}
