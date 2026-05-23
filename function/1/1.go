package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "worker 1 finished"
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch1 <- "worker 2 finished"
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	}
}
