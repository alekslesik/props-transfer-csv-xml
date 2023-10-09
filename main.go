package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"os"
)

type XMLProp struct {
	XMLName          xml.Name `xml:"prop"`
	Text             string   `xml:",chardata"`
	OldID            string   `xml:"old_id"`
	Name             string   `xml:"name"`
	Active           string   `xml:"active"`
	Sort             string   `xml:"sort"`
	Code             string   `xml:"code"`
	DefaultValue     string   `xml:"default_value"`
	PropertyType     string   `xml:"property_type"`
	RowCount         string   `xml:"row_count"`
	ColCount         string   `xml:"col_count"`
	ListType         string   `xml:"list_type"`
	Multiple         string   `xml:"multiple"`
	XMLID            string   `xml:"xml_id"`
	FileType         string   `xml:"file_type"`
	MultipleCnt      string   `xml:"multiple_cnt"`
	LinkIblockID     string   `xml:"link_iblock_id"`
	WithDescription  string   `xml:"with_description"`
	Searchable       string   `xml:"searchable"`
	Filtrable        string   `xml:"filtrable"`
	IsRequired       string   `xml:"is_required"`
	Version          string   `xml:"version"`
	UserType         string   `xml:"user_type"`
	UserTypeSettings string   `xml:"user_type_settings"`
	Hint             string   `xml:"hint"`
}

func main() {
	op := "main()"

	csvPath := "props.csv"

	// get properties from csv
	csvFile, err := os.OpenFile(csvPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("%s: %s: open csv file error", op, err)
		return
	}
	defer csvFile.Close()

	// create new reader
	reader := csv.NewReader(csvFile)
	reader.Comma = rune(';')

	records := getCsvArr(reader)

	for _, record := range records {
		for _, r := range record {
			fmt.Printf("record = %s", r)
		}
	}
}

// Get all records from csv file
func getCsvArr(reader *csv.Reader) ([][]string) {
	op := "main.getCsvArr()"
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("%s: %s: csv records read error", op, err)
		return nil
	}

	return records
}
