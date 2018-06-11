package cmd

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
	"github.com/spf13/cobra"
)

var sch2Cmd = &cobra.Command{
	Use:   "sch2",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sch2Cmd called")
		sch2()
	},
}

func init() {
	rootCmd.AddCommand(sch2Cmd)
}

func task() {
	fmt.Println("I am runnning task.")
}

func taskWithParams(a int, b string) {
	fmt.Println(b, a)
}

func sch2() {
	fmt.Println("sch2")

	// Do jobs with params
	gocron.Every(2).Seconds().Do(taskWithParams, 2, "hello")

	// Do jobs without params
	gocron.Every(1).Seconds().Do(task)

	<-gocron.Start()

	// also , you can create a your new scheduler,
	// to run two scheduler concurrently
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(taskWithParams, 3, "hi")
	<-s.Start()
}
