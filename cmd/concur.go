package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var concurCmd = &cobra.Command{
	Use:   "concur",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		// trySum()
		var c chan string = make(chan string)

		// Using a channel like this synchronizes the two goroutines.
		// When pinger attempts to send a message on the channel it will wait until printer is ready to receive the message.
		// (this is known as blocking)
		go pinger(c)
		go ponger(c)
		go printer(c)

		var input string
		fmt.Scanln(&input)
	},
}

func init() {
	rootCmd.AddCommand(concurCmd)
}

// goroutine 101

func trySum() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)

	d := make(chan int)
	go sum(a[len(a)/2:], d)
	x, y := <-c, <-d

	fmt.Println(x, y)
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	fmt.Println(total)
	c <- total // send total to c
}

// goroutine 102

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
