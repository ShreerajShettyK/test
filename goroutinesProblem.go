// 2 goroutines
// 1 print odd numbers 1,3,5,7
// 1 print even numbers 2,4,6,8
// sequentially goroutines should execute
// output should be like 1,2,3,4,5,6,7,8,9,10

package main

import (
	"fmt"
	"sync"
)

func printOddNumbers(oddCh, evenCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		<-oddCh
		fmt.Print(i, ",")
		evenCh <- true
	}

}

func printEvenNumbers(oddCh, evenCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<-evenCh
		if i == 10 {
			fmt.Print(i)
		} else {
			fmt.Print(i, ",")
			oddCh <- true
		}
	}
}

func main() {
	oddCh := make(chan bool)
	evenCh := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)

	go printOddNumbers(oddCh, evenCh, &wg)
	go printEvenNumbers(oddCh, evenCh, &wg)

	oddCh <- true

	wg.Wait()
}
