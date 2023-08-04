//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2023 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
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

// MailTemplatesBody is to structure the xml data
type MailTemplatesBody struct {
	Request MailTemplatesRequest `xml:"Request"`
}

type MailTemplatesRequest struct {
	AfterbuyGlobal AfterbuyGlobal `xml:"AfterbuyGlobal"`
}

// MailTemplatesReturn is to decode the xml return data
type MailTemplatesReturn struct {
	XmlName    xml.Name                  `xml:"Afterbuy"`
	CallStatus string                    `xml:"CallStatus"`
	CallName   string                    `xml:"CallName"`
	Result     MailTemplatesReturnResult `xml:"Result"`
}

type MailTemplatesReturnResult struct {
	XmlName       xml.Name                         `xml:"Result"`
	MailTemplates MailTemplatesReturnMailTemplates `xml:"MailTemplates"`
}

type MailTemplatesReturnMailTemplates struct {
	XmlName      xml.Name                          `xml:"MailTemplates"`
	MailTemplate []MailTemplatesReturnMailTemplate `xml:"MailTemplate"`
}

type MailTemplatesReturnMailTemplate struct {
	XmlName         xml.Name `xml:"MailTemplate"`
	TemplateId      int      `xml:"TemplateID"`
	TemplateName    string   `xml:"TemplateName"`
	TemplateSubject string   `xml:"TemplateSubject"`
	TemplateHtml    int      `xml:"TemplateHtml"`
	TemplateText    string   `xml:"TemplateText"`
}

// MailTemplates is to get all mail templates
func MailTemplates(body MailTemplatesBody) (MailTemplatesReturn, error) {

	// Convert body
	convert, err := xml.Marshal(body)
	if err != nil {
		return MailTemplatesReturn{}, err
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
		return MailTemplatesReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Decode data
	var decode MailTemplatesReturn

	err = xml.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return MailTemplatesReturn{}, err
	}

	// Return data
	return decode, nil

}
