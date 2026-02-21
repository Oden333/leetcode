package main

import (
	"fmt"
	"strings"
)

func main2() {
	var n int
	fmt.Scan(&n)

	var text string
	fmt.Scan(&text)

	minLength, maxLength := MinMaxLength(text)

	fmt.Printf("%d %d\n", minLength, maxLength)
}

func MinMaxLength(text string) (int, int) {

	lines := strings.Split(text, "#")

	minLength := len(lines[0])
	maxLength := len(lines[0])

	for _, line := range lines {
		length := len(line)
		if length < minLength {
			minLength = length
		}
		if length > maxLength {
			maxLength = length
		}
	}

	return minLength, maxLength
}
