package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tealeg/xlsx"
)

// excelCmd represents the excel command
var excelCmd = &cobra.Command{
	Use:   "excel",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("excel called")
		excel()
	},
}

func init() {
	rootCmd.AddCommand(excelCmd)
}

func excel() {
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

// func Excel() {
// 	fmt.Printf("Hello, world.\n")

// 	excelFileName := "asset/combine.xlsx"
// 	xlFile, err := xlsx.OpenFile(excelFileName)
// 	if err != nil {
// 		fmt.Printf("Hello, world.\n")
// 	}
// 	for _, sheet := range xlFile.Sheets {
// 		for _, row := range sheet.Rows {
// 			for _, cell := range row.Cells {
// 				text := cell.String()
// 				fmt.Printf("%s\n", text)
// 			}
// 		}
// 	}
// }
