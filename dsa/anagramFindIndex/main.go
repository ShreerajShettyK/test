// q: Given two strings s and p, return all start indices of p's anagrams in s.

// Example Input:
// s = "cbaebabacd", p = "abc"

// Output:
// [0, 6]
// Explanation: "cba" at index 0, "bac" at index 6

package main

import (
	"fmt"
	"sort"
)

func findIndices(s, p string) []int {
	var indices []int
	if len(p) > len(s) || len(s) == 0 || len(p) == 0 {
		fmt.Println("Cant determine pls change ur inputs")
		return indices
	}

	brokenP := []byte(p)
	sort.Slice(brokenP, func(i, j int) bool {
		return brokenP[i] < brokenP[j]
	})

	pSorted := string(brokenP)

	for i := 0; i <= len(s)-len(p); i++ {
		window := s[i : i+len(p)]
		fmt.Println(window)
		subStr := []byte(window)
		sort.Slice(subStr, func(i, j int) bool {
			return subStr[i] < subStr[j]
		})
		if string(subStr) == pSorted {
			indices = append(indices, i)
		}
	}

	return indices
}

func main() {
	s := "cbaebabacdabc"
	p := "abc"
	indices := findIndices(s, p)
	fmt.Println(indices)
}
