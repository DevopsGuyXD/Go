package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {

	fmt.Println("Welcome to the Excel program")

	f := excelize.NewFile()

	index, err := f.NewSheet("Sheet1"); if err != nil{ log.Fatal(err) }

	f.SetCellValue("Sheet1", "A1", "Hello")
	f.SetCellValue("Sheet1", "B1", "World")

	f.SetActiveSheet(index)

	err = f.SaveAs("testFile.xlsx"); if err != nil{ log.Fatal(err) }
}