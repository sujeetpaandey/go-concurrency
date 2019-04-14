package main

import (
	"fmt"
	"time"
)

func main() {
	// single channel, single goroutine, one write, one read pattern
	//hello()

	// timers pattern
	//timers()

	// ping pong patters
	pingPong()

}

func hello() {
	// create new channel of type int
	ch := make(chan int)

	// start new anonymous goroutine
	go func() {
		// send 42 to channel
		ch <- 42
	}()
	// read from channel
	fmt.Printf("You receive %v", <-ch)
}

func timers() {
	for i := 0; i < 24; i++ {
		c := timer(1 * time.Second)
		fmt.Printf("Channel Received %v, iteration %d \n", <-c, i)
	}
}

func pingPong() {
	var Ball int
	table := make(chan int)
	go player(table)
	go player(table)

	table <- Ball
	time.Sleep(1 * time.Second)
	<-table
}
