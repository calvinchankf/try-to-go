package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var strucCmd = &cobra.Command{
	Use:   "struct",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("struc called")
		struc()
	},
}

func init() {
	rootCmd.AddCommand(strucCmd)
}

// class

type Car struct {
	name  string
	price float64
}

type Ferrari struct {
	Car
	formula int
}

func (c Car) Name() string {
	return c.name
}

func (c Car) Price() float64 {
	return c.price
}

func (f Ferrari) Formula() int {
	return f.formula
}

// try construct

func struc() {
	ferrari := Ferrari{Car{"ferra", 1000000}, 1}
	fmt.Println("ff", ferrari.name)
	fmt.Println("ff", ferrari.price)
	fmt.Println("ff", ferrari.formula)
}
