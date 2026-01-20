package main

import "fmt"

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())

// 	go func(ctx context.Context) {
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				fmt.Println("GOroutine stoppped:", ctx.Err())
// 				return
// 				// ctx.Done() returns a channel :This channel is closed when cancel() is called
// 			default:
// 				fmt.Println("Goroutine working")
// 				time.Sleep(100 * time.Millisecond)
// 			}
// 		}
// 	}(ctx)

// 	time.Sleep(500 * time.Millisecond)
// 	cancel()
// 	time.Sleep(500 * time.Millisecond)
// }

////// new guaranteed ordered using channels and goroutines

var x = 10

func a(done chan bool) {
	x = 20
	done <- true
}

func b(done chan bool, complete chan bool) {
	<-done
	fmt.Println(x)
	complete <- true
}

func main() {
	done := make(chan bool)
	complete := make(chan bool)
	go a(done)
	go b(done, complete)
	<-complete // Wait for b to finish
}
