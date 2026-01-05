package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() (i int) {
	toDelete := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
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
