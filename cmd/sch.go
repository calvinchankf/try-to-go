package cmd

import (
	"fmt"

	"github.com/robfig/cron"
	"github.com/spf13/cobra"
)

var schCmd = &cobra.Command{
	Use:   "sch",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sch called")
		sch()
	},
}

func init() {
	rootCmd.AddCommand(schCmd)
}

func sch() {
	fmt.Println("schedule")

	c := cron.New()
	c.AddFunc("* * * * * *", func() { fmt.Println("Every sec") })
	c.Start()
	// c.Stop()
	select {}
}
