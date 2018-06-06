package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func main() {
	fmt.Printf("Hello, world.\n")

	excelFileName := "asset/combine.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("Hello, world.\n")
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
}
