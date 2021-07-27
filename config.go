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
	PartnerId       int    `xml:"PartnerID"`
	PartnerPassword string `xml:"PartnerPassword"`
	UserId          string `xml:"UserID"`
	UserPassword    string `xml:"UserPassword"`
	CallName        string `xml:"CallName"`
	DetailLevel     int    `xml:"DetailLevel"`
	ErrorLanguage   string `xml:"ErrorLanguage"`
}

// Config is to define the request data
type Config struct {
	Request *http.Request
	Body    []byte
}

// Send is to send a new request
func (r *Config) Send() (*http.Response, error) {

	// Define client
	client := &http.Client{}

	// Define request
	request, err := http.NewRequest(method, abInterfaceBaseUrl, bytes.NewBuffer(r.Body))
	if err != nil {
		return nil, err
	}

	// Add new request request
	if r.Request != nil {
		request = r.Request
	}

	// Send request & get response
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Return data
	return response, nil

}
