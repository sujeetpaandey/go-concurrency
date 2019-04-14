package main

import (
	"fmt"
	"time"
)

func player(table chan int) {
	for {
		ball := <-table
		ball++
		time.Sleep(100 * time.Millisecond)
		table <- ball
		fmt.Printf("Channel Received Table Value as %v\n", table)
	}
}
