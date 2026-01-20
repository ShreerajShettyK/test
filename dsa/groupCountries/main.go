// q: A postal sorting system receives thousands of envelopes with city names on them.
// The system must group all envelopes that have the same letters in different orders together, so they can be processed by the same route.

// Task: Write a function that takes a list of city names and groups together all the names that are made of the same letters (case-insensitive).

// Example Input:

// ["listen", "silent", "enlist", "rat", "tar", "art"]

// output
// [["listen", "silent", "enlist"], ["rat", "tar", "art"]]

package main

import (
	"fmt"
	"sort"
	"strings"
)

func GroupCities(cities []string) [][]string {
	groups := make(map[string][]string) ///key will be string and value will be slice []strings

	for _, city := range cities {
		lowerCity := strings.ToLower(city)
		runes := []byte(lowerCity)
		//breaks down the city word to induvidual letters "bat" ---> "b","a","t"
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		sortedKey := string(runes)
		groups[sortedKey] = append(groups[sortedKey], city)
	}

	result := make([][]string, 0, len(groups))
	for key, group := range groups {
		fmt.Println("Key:", key)
		result = append(result, group)
	}

	return result
}

func main() {
	input := []string{"listen", "art", "silent", "enlist", "rat", "tar", "art", "dog", "nistel"}

	grouped := GroupCities(input)
	fmt.Println(grouped)
	// for _, group := range grouped {
	// 	fmt.Println(group)
	// }
}
