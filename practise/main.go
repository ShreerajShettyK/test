package main

import "fmt"

// 2. Valid Parentheses (Easy)

// Check if brackets are properly balanced: ()[]{}
// Perfect for practicing stacks (I see you have a stacks folder!)
// LeetCode #20

type stack struct {
	items []rune
}

func (s *stack) push(i rune) {
	s.items = append(s.items, i)
}

func (s *stack) pop() rune {
	if len(s.items) == 0 {
		return 0
	}
	lastIndex := len(s.items) - 1
	deletedChar := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return deletedChar
}

func (s *stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *stack) peek() rune {
	if len(s.items) == 0 {
		return 0
	}
	return s.items[len(s.items)-1]
}

func isValid(str string) bool {
	// Map of closing brackets to their matching opening brackets
	matches := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	var s stack

	for _, char := range str {
		// Check if it's a closing bracket
		if openingBracket, isClosing := matches[char]; isClosing {
			// If stack is empty or top doesn't match, invalid
			if s.isEmpty() || s.peek() != openingBracket {
				return false
			}
			s.pop()
		} else {
			// It's an opening bracket, push it
			s.push(char)
		}
	}

	// Valid only if all brackets are matched (stack is empty)
	return s.isEmpty()
}

func main() {
	testCases := []string{
		")(",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
		"",
		"(((",
		")))",
	}

	for _, test := range testCases {
		result := isValid(test)
		fmt.Printf("Input: %-10s -> Valid: %v\n", `"`+test+`"`, result)
	}
}
