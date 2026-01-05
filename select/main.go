package main

import "sync"

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go func() {
// 		ch1 <- "Hello from channel 1"
// 	}()

// 	go func() {
// 		ch2 <- "Hello from channel 2"
// 	}()

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case msg1 := <-ch1:
// 			println("message got from ch1:", msg1)
// 		case msg2 := <-ch2:
// 			println("message got from ch2:", msg2)
// 		}
// 	}
// }

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		ch1 <- "Hello from channel 1"
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ch2 <- "Hello from channel 2"
		ack := <-ch2
		println("ch2 goroutine received:", ack)
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			println("message got from ch1:", msg1)
		case msg2 := <-ch2:
			println("message got from ch2:", msg2)
			ch2 <- "acknowledged from channel 2"
		}
	}

	wg.Wait()
}
