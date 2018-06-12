package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate called")
		migration()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func migration() {
	fmt.Println("migrate")

}
