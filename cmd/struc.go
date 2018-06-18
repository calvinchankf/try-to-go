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

		var ptr *int
		if ptr != nil {
			fmt.Println("has")
		} else {
			fmt.Println("has not")
		}

		i := 1
		ptr = &i
		fmt.Println(*ptr)

		struc()
	},
}

func init() {
	rootCmd.AddCommand(strucCmd)
}

// class

type Car struct {
	name  *string
	price *float64
}

type Ferrari struct {
	Car
	formula *int
}

func (c Car) Name() *string {
	return c.name
}

func (c Car) Price() *float64 {
	return c.price
}

func (f Ferrari) Formula() *int {
	return f.formula
}

// try construct

func struc() {
	// a := "calvin"
	// var a string
	// a_ptr := &a

	ferrari := Ferrari{
		Car{
			returnStringAdd("calvin"),
			returnFloatAdd(1),
		},
		returnIntAdd(2),
	}
	fmt.Println("ff", *ferrari.name)
	fmt.Println("ff", ferrari.price)
	fmt.Println("ff", ferrari.formula)

	benzi := Ferrari{
		Car{
			nil,
			nil,
		},
		returnIntAdd(2),
	}
	fmt.Println("ff", benzi.name)
	fmt.Println("ff", benzi.price)
	fmt.Println("ff", benzi.formula)
}

func returnStringAdd(value string) *string {
	a := value
	return &a
}

func returnFloatAdd(value float64) *float64 {
	a := value
	return &a
}

func returnIntAdd(value int) *int {
	a := value
	return &a
}
