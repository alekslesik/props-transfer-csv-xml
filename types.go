package main

import "encoding/xml"

type AsdIblockProps struct {
	XMLName xml.Name `xml:"asd_iblock_props"`
	Text    string   `xml:",chardata"`
	Props   Props    `xml:"props"`
}

type Props struct {
	Text  string `xml:",chardata"`
	Prop []Prop `xml:"prop"`
}

type Prop struct {
	Text             string `xml:",chardata"`
	OldID            string `xml:"old_id"`
	Name             CDATA  `xml:"name"`
	Active           string `xml:"active"`
	Sort             string `xml:"sort"`
	Code             string `xml:"code"`
	DefaultValue     string `xml:"default_value"`
	PropertyType     string `xml:"property_type"`
	RowCount         string `xml:"row_count"`
	ColCount         string `xml:"col_count"`
	ListType         string `xml:"list_type"`
	Multiple         string `xml:"multiple"`
	XMLID            CDATA  `xml:"xml_id"`
	FileType         string `xml:"file_type"`
	MultipleCnt      string `xml:"multiple_cnt"`
	LinkIblockID     string `xml:"link_iblock_id"`
	WithDescription  string `xml:"with_description"`
	Searchable       string `xml:"searchable"`
	Filtrable        string `xml:"filtrable"`
	IsRequired       string `xml:"is_required"`
	Version          string `xml:"version"`
	UserType         string `xml:"user_type"`
	UserTypeSettings CDATA `xml:"user_type_settings"`
	Hint             string `xml:"hint"`
}

type CDATA struct {
	Value string `xml:",cdata"`
}
