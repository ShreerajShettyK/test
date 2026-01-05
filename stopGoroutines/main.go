package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("GOroutine stoppped:", ctx.Err())
				return
				// ctx.Done() returns a channel :This channel is closed when cancel() is called
			default:
				fmt.Println("Goroutine working")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(500 * time.Millisecond)
	cancel()
	time.Sleep(500 * time.Millisecond)
}
