// unbuffered
// 1️⃣ Unbuffered channel (sender & receiver must meet)

// 👉 Rule:

// Sender blocks until a receiver is ready

// Receiver blocks until a sender sends
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	ch := make(chan int)
// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("Sender sending data")
// 		ch <- 50
// 		fmt.Println("Sender sent data successfully")
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		val := <-ch
// 		fmt.Println("Receiver received data:", val)
// 	}()

// 	wg.Wait()
// }

// 2️⃣ Buffered channel (decoupled sender & receiver)

// 👉 Rule:

// Sender blocks only when buffer is full

// Receiver blocks only when buffer is empty

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)

	ch <- 10
	ch <- 20
	fmt.Println("Sender: Sent 2 values to the buffered channel (Now it is full and sender blocks)")

	go func() {
		val := <-ch
		fmt.Println("Receiver: Received value", val)
		ch <- 30
		fmt.Println("Sender: Sent another value to the buffered channel")
	}()

	fmt.Println("Receiver: Received value(outside main)", <-ch)
	fmt.Println("Receiver: Received value", <-ch)
}
