package main

import (
	"fmt"
)

func sender(ch chan int) {
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	close(ch) // channel is closed by the sender
}

func main() {
	ch := make(chan int)

	go sender(ch)

	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("Channel closed, no more values")
			break
		}
		fmt.Println("Received:", value)
	}
}
