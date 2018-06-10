package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var concurCmd = &cobra.Command{
	Use:   "concur",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		a := []int{7, 2, 8, -9, 4, 0}

		c := make(chan int)
		go sum(a[:len(a)/2], c)

		d := make(chan int)
		go sum(a[len(a)/2:], d)
		x, y := <-c, <-d

		fmt.Println(x, y)
	},
}

func init() {
	rootCmd.AddCommand(concurCmd)
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	fmt.Println(total)
	c <- total // send total to c
}
