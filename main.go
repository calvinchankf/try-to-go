package main

import (
	"fmt"

	"github.com/calvinchankf/try-to-go/cmd"
)

func add(a int, b int) int {
	c := a + b
	return c
}

func main() {
	fmt.Println("hello world")
	result := add(1, 2)
	fmt.Printf("%d", result)

	cmd.Execute()
}
