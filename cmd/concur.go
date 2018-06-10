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
		// pingPong()

		selec()
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
func pingPong() {
	var c chan string = make(chan string)

	// Using a channel like this synchronizes the two goroutines.
	// When pinger attempts to send a message on the channel it will wait until printer is ready to receive the message.
	// (this is known as blocking)
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

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

// select

func selec() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "every 1 sec"
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			c2 <- "every 3 sec"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
