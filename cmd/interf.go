package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var interfCmd = &cobra.Command{
	Use:   "interface",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("struc called")
		interf()
	},
}

func init() {
	rootCmd.AddCommand(interfCmd)
}

// class

type Human struct {
	name string
}

type Tricker struct {
	Human
	combos []string
}

func (h Human) Name() string {
	return h.name
}

func (t Tricker) Combos() []string {
	return t.combos
}

// interface

type HumanInterf interface {
	Eat()
	Sleep()
}

type TrickerInterf interface {
	Training()
}

// implement interface

func (h Human) Eat() {
	fmt.Println("eat eat eat")
}

func (h Human) Sleep() {
	fmt.Println("sleep sleep sleep")
}

func (t Tricker) Training() {
	fmt.Println("train train train")
}

// try construct

func interf() {

	peter := Human{"calvin"}
	fmt.Println(peter.name)

	peter.Eat()
	peter.Sleep()
	// peter.Training() // it will fail 👍🏻

	calvin := Tricker{Human{"calvin"}, []string{"b9", "chartwheel", "frontflip"}}
	fmt.Println(calvin.name)
	fmt.Println(calvin.combos)

	calvin.Training()
}
