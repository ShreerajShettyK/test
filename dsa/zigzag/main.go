// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
// (you may want to display this pattern in a fixed font for better legibility)

// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"

// Write the code that will take a string and make this conversion given a number of rows:

// string convert(string s, int numRows);

// Example 1:

// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"

package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	rows := make([]strings.Builder, numRows)
	// this will create a slice like below
	//rows = [
	//   Builder{}, // row 0
	//   Builder{}, // row 1
	//   Builder{}, // row 2
	// ]
	currentRow := 0
	goingDown := false

	for _, ch := range s { // PAYPALISHIRING
		rows[currentRow].WriteRune(ch)

		if currentRow == 0 || currentRow == numRows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			currentRow++
		} else {
			currentRow--
		}
	}
	var result strings.Builder
	for i := 0; i < numRows; i++ {
		// fmt.Println(rows[i].String())
		result.WriteString(rows[i].String())
	}

	return result.String()
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
