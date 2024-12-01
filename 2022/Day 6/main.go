package main

import (
	"fmt"
	"strings"

	tools "github.com/edklesel/AdventOfCode/2022/GO/tools"
)

var filename string = "day6.txt"

func checkDuplicates(arr []string) bool {

	var chars = []string{}

	for _, i := range arr {

		for _, char := range chars {

			if i == char {

				return true

			}

		}

		chars = append(chars, i)

	}

	return false

}

func findMarker(sequence string, chars int) (index int, marker string) {

	var _marker = []string{}

	for i, char := range strings.Split(sequence, "") {

		// Create the marker
		_marker = append(_marker, char)

		if len(_marker) > chars {

			_marker = append([]string{}, _marker[1:]...)

		}

		// Search the marker for duplicates
		if !checkDuplicates(_marker) && len(_marker) == chars {
			index = i + 1
			break
		}
	}

	return index, strings.Join(_marker[len(_marker)-chars:], "")

}

func main() {

	input := tools.ReadInput(filename)

	var sequence string = input[0]
	var index int
	var marker string

	index, marker = findMarker(sequence, 4)

	fmt.Println("Part 1")
	fmt.Printf("Marker found at character %d\n", index)
	fmt.Printf("Marker: %s\n", marker)

	index, marker = findMarker(sequence, 14)

	fmt.Println("Part 2")
	fmt.Printf("Marker found at character %d\n", index)
	fmt.Printf("Marker: %s\n", marker)

}
