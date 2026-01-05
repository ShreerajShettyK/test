// package main

// import "fmt"

// func add(x int, y int) int {
// 	var val int
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("Iteration:", i)
// 		val = i
// 	}
// 	if x > y {
// 		fmt.Println("x greather than y")
// 	} else if x == y {
// 		fmt.Println("x equal to y")
// 	} else {
// 		fmt.Println("y greather than x")
// 	}
// 	return val
// }

// func addPtr(x *int, y int) int {
// 	*x += y
// 	return *x
// }

// func main() {
// 	// fmt.Println(add(2, 4))

// 	var num int = 43
// 	var myPtr *int = &num
// 	fmt.Println("Address of num:", myPtr)
// 	fmt.Println("Value of num:", *myPtr)

// 	fmt.Println(addPtr(myPtr, 7))

// 	fmt.Println("New Address of num:", myPtr)
// 	fmt.Println("New Value of num:", *myPtr)

// 	myArr := [7]int{1, 2, 3, 4, 5}
// 	for _, num := range myArr {
// 		fmt.Println(num)
// 	}
// }

package main

import "fmt"

// type Speaker interface {
// 	Speak()
// }

// //
// // 2️⃣ Base struct (reusable behavior)
// //
// type Animal struct {
// 	Name string
// }

// func (a Animal) Speak() {
// 	fmt.Println("Animal speaks")
// }

// //
// // 3️⃣ Derived struct using composition (embedding)
// //
// type Dog struct {
// 	Animal // embedded struct
// 	Breed  string
// }

// // This method shadows Animal.Speak()
// func (d Dog) Speak() {
// 	fmt.Println("Dog barks")
// }

// //
// // 4️⃣ Another struct using the same base
// //
// type Cat struct {
// 	Animal
// }

// func (c Cat) Speak() {
// 	fmt.Println("Cat meows")
// }

// //
// // 5️⃣ Main function
// //
// func main() {
// 	dog := Dog{
// 		Animal: Animal{Name: "Bruno"},
// 		Breed:  "Labrador",
// 	}

// 	cat := Cat{
// 		Animal: Animal{Name: "Kitty"},
// 	}

// 	// Direct method calls
// 	dog.Speak() // Dog barks
// 	dog.Animal.Speak()

// 	cat.Speak() // Cat meows

// 	fmt.Println()

// 	// Polymorphism using interface
// 	animals := []Speaker{dog, cat}

// 	for _, a := range animals {
// 		a.Speak() // Calls the correct implementation
// 	}
// }

// func main() {

// 	//diff in array and slice
// 	myarr := [5]int{10, 20, 30, 40, 50}

// 	mySlice := []int{1, 2, 3, 4}

// 	fmt.Println(myarr)
// 	fmt.Println(mySlice)

// 	// make(map[keytype]valuetype) map in golang
// 	myMap := make(map[int]string)
// 	myMap[1] = "Hello"
// }

func main() {
	ch := make(chan int)

	go func() {
		ch <- 42
		ch <- 26
		close(ch)
	}()
	for {
		val, ok := <-ch
		if !ok {
			return
		}
		println(val)
	}
}

func Add(x int, y int) int {
	var val int
	for i := 0; i < 10; i++ {
		fmt.Println("Iteration:", i)
		val = i
	}
	if x > y {
		fmt.Println("x greather than y")
	} else if x == y {
		fmt.Println("x equal to y")
	} else {
		fmt.Println("y greather than x")
	}
	return val
}
