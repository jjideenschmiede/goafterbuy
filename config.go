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
	"io"
	"net/http"
)

// AfterbuyGlobal is to define afterbuy global data
type AfterbuyGlobal struct {
	PartnerId       int    `xml:"PartnerID,omitempty"`
	PartnerPassword string `xml:"PartnerPassword,omitempty"`
	UserId          string `xml:"UserID,omitempty"`
	UserPassword    string `xml:"UserPassword,omitempty"`
	PartnerToken    string `xml:"PartnerToken,omitempty"`
	AccountToken    string `xml:"AccountToken,omitempty"`
	CallName        string `xml:"CallName,omitempty"`
	DetailLevel     int    `xml:"DetailLevel,omitempty"`
	ErrorLanguage   string `xml:"ErrorLanguage,omitempty"`
}

// Config is to define the request data
type Config struct {
	BaseUrl, Method string
	Body            io.Reader
	Header          map[string]string
}

// Send is to send a new request
func (c *Config) Send() (*http.Response, error) {

	// Define client
	client := &http.Client{}

	// Define request
	request, err := http.NewRequest(c.Method, c.BaseUrl, c.Body)
	if err != nil {
		return nil, err
	}

	// Add request header
	for index, value := range c.Header {
		request.Header.Add(index, value)
	}

	// Send request & get response
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Return data
	return response, nil

}
