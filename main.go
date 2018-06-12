package main

import (
	"fmt"

	"github.com/calvinchankf/try-to-go/cmd"
)

func main() {
	// golang int64 = postgres biginit in range
	a := 9223372036854775807
	fmt.Println("hello world", a)
	cmd.Execute()
}
