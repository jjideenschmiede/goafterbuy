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
	"bytes"
	"net/http"
)

const (
	abInterfaceBaseUrl   = "https://api.afterbuy.de/afterbuy/ABInterface.aspx"
	shopInterfaceBaseUrl = "https://api.afterbuy.de/afterbuy/ShopInterfaceUTF8.aspx"
	method               = "GET"
)

// AfterbuyGlobal is to define afterbuy global data
type AfterbuyGlobal struct {
	PartnerId       int    `xml:"PartnerID,omitempty"`
	PartnerPassword string `xml:"PartnerPassword,omitempty"`
	UserId          string `xml:"UserID,omitempty"`
	UserPassword    string `xml:"UserPassword,omitempty"`
	CallName        string `xml:"CallName,omitempty"`
	DetailLevel     int    `xml:"DetailLevel,omitempty"`
	ErrorLanguage   string `xml:"ErrorLanguage,omitempty"`
}

// Config is to define the request data
type Config struct {
	Request *http.Request
	Body    []byte
}

// Send is to send a new request
func (c *Config) Send() (*http.Response, error) {

	// Define client
	client := &http.Client{}

	// Define request
	request, err := http.NewRequest(method, abInterfaceBaseUrl, bytes.NewBuffer(c.Body))
	if err != nil {
		return nil, err
	}

	// Add new request request
	if c.Request != nil {
		request = c.Request
	}

	// Send request & get response
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Return data
	return response, nil

}
