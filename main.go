package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
)

func main() {
	sheetName := os.Getenv("SHEET_NAME")
	if sheetName == "" {
		log.Fatalln("SHEET_NAME must be given")
	}

	file, err := excelize.OpenFile("kelimeler.xlsx")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()

	rows, err := file.GetRows(os.Getenv("SHEET_NAME"))
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, row := range rows {
		dto := fulfillWordInformation(row)
		fmt.Printf("%+v\n", dto)
	}
}

func fulfillWordInformation(row []string) EnglishWordDTO {
	var dto EnglishWordDTO

	if len(row) == 1 {
		dto.Word = row[0]
	} else if len(row) == 2 {
		dto.Word = row[0]
		dto.Meaning = row[1]
	} else if len(row) == 3 {
		dto.Word = row[0]
		dto.Meaning = row[1]
		dto.ExampleSentence = row[2]
	}

	return dto
}
