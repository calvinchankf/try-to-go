package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("defer it called")
		jason()
		fmt.Println("defer it called 2")
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)
}

func jason() {

	var hashtable = make(map[string]string)
	hashtable["a"] = "calvi"
	hashtable["b"] = "calvin"

	// marshal
	jsonObj, _ := json.Marshal(hashtable)
	fmt.Println(string(jsonObj))

	// unmarshal
	var mapResult map[string]string
	if err := json.Unmarshal(jsonObj, &mapResult); err != nil {
		panic(err)
	}
	fmt.Println(mapResult)
}
