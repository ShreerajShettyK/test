package main

import (
	"errorHandling/practise"
	"fmt"
)

// func main() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered from panic:", r)
// 		}
// 		fmt.Println("Defered function executed")
// 	}()

// 	panic("Something went wrong!")

// }

func main() {
	defer recoverFunc()

	practise.Add(3, 5)

}

func recoverFunc() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
	fmt.Println("Defered function executed")

}
