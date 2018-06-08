package cmd

import (
	"fmt"

	"github.com/robfig/cron"
	"github.com/spf13/cobra"
)

var schCmd = &cobra.Command{
	Use:   "sch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	c.AddFunc("5 * * * * *", func() { fmt.Println("Every sec") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()
	c.Stop()
}
