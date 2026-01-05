package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() (i int) {
	lastIndex := len(s.items) - 1
	toDelete := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return toDelete
}

func main() {
	var s Stack
	s.Push(100)
	s.Push(200)
	s.Push(300)

	fmt.Println(s.items)

	s.Pop()
	s.Pop()

	fmt.Println(s.items)
}
