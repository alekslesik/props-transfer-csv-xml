package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mozillazg/go-unidecode"
)

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

	csvProps, err := getCsvArr(csvFile)
	if err != nil {
		fmt.Printf("%s: %s: csv records read error", op, err)
	}

	AsdIblockProps := new(AsdIblockProps)
	id := 88
	xmlId := 35

	for _, prop := range csvProps[0] {
		id = id + 1
		strID := strconv.Itoa(id)
		xmlId = xmlId + 1
		strXmlId := strconv.Itoa(xmlId)
		code := translit(prop)

		Prop := newXMLProp(strID, strXmlId, code, prop)
		AsdIblockProps.Props.Prop = append(AsdIblockProps.Props.Prop, Prop)

		output, err := xml.MarshalIndent(AsdIblockProps, "", "\t")
		if err != nil {
			fmt.Printf("%s: %s: MarshalIndent error", op, err)
		}

		XMLPath := "result.xml"
		err = os.WriteFile(XMLPath, output, 0755)
		if err != nil {
			fmt.Printf("%s: %s: WriteFile error", op, err)
		}
	}
}

// Get all records from csv file
func getCsvArr(file *os.File) ([][]string, error) {
	// create new reader
	reader := csv.NewReader(file)
	reader.Comma = rune(';')
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func newXMLProp(id, xmlId, code, name string) Prop {
	Prop := Prop{
		OldID:            id,
		Name:             CDATA{name},
		Active:           "Y",
		Sort:             "100",
		Code:             code,
		DefaultValue:     "",
		PropertyType:     "S",
		RowCount:         "1",
		ColCount:         "30",
		ListType:         "L",
		Multiple:         "N",
		XMLID:            CDATA{xmlId},
		FileType:         "",
		MultipleCnt:      "5",
		LinkIblockID:     "0",
		WithDescription:  "N",
		Searchable:       "N",
		Filtrable:        "N",
		IsRequired:       "N",
		Version:          "1",
		UserType:         "",
		UserTypeSettings: CDATA{"a:0:{}"},
		Hint:             "",
	}

	return Prop
}

// Translit to uppercase english, replace spaces by _
func translit(input string) string {
	translit := unidecode.Unidecode(input)
	translit = strings.ReplaceAll(translit, " ", "_")
	translit = strings.ToUpper(translit)

	return translit
}
