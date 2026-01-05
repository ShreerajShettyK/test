package main

import "sync"

var (
	counter int
	wg      sync.WaitGroup
	mu      sync.Mutex
)

func incrementCounter() {
	defer wg.Done()
	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementCounter()
	}

	wg.Wait()
	println("Final Counter:", counter)

}

// without mutex
// package main

// import "sync"

// var (
// 	counter int
// 	wg      sync.WaitGroup
// 	// mu      sync.Mutex
// )

// func incrementCounter() {
// 	defer wg.Done()
// 	// mu.Lock()
// 	counter++
// 	// mu.Unlock()
// }

// func main() {
// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go incrementCounter()
// 	}

// 	wg.Wait()
// 	println("Final Counter:", counter)

// }
