package main

import (
	"fmt"
	"sort"
	"strings"

	tools "github.com/edklesel/AdventOfCode/2022/GO/tools"
)

var filename string = "day3.txt"

func priority(char string) int {

	var priority int
	var alphabet string = "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < len(alphabet); i++ {
		if strings.Split(alphabet, "")[i] == strings.ToLower(char) {
			priority = i + 1
			break
		}
	}

	if char == strings.ToUpper(char) {
		priority += 26
	}

	return priority

}

func main() {

	input := tools.ReadInput(filename)
	var totalpriority int = 0

	for i := 0; i < len(input); i++ {

		var duplicate string

		// Split the backpack up
		var items string = input[i]
		var comp1 string = items[0 : len(items)/2]
		var comp2 string = items[len(items)/2:]

		for j := 0; j < len(comp1); j++ {
			if tools.Find(strings.Split(comp2, ""), strings.Split(comp1, "")[j]) {
				duplicate = strings.Split(comp1, "")[j]
				break
			}
		}

		totalpriority += priority(duplicate)

	}

	fmt.Println("Part 1")
	fmt.Printf("Total Priority: %d\n", totalpriority)

	i := 0
	totalpriority = 0
	for i < len(input) {

		var backpacks [3]string
		backpacks[0] = input[i]
		backpacks[1] = input[i+1]
		backpacks[2] = input[i+2]

		sort.Slice(backpacks[:], func(i, j int) bool {
			return len(backpacks[i]) > len(backpacks[j])
		})

		var char string
		for j := 0; j < len(backpacks[0]); j++ {
			if tools.Find(strings.Split(backpacks[1], ""), strings.Split(backpacks[0], "")[j]) &&
				tools.Find(strings.Split(backpacks[2], ""), strings.Split(backpacks[0], "")[j]) {
				char = strings.Split(backpacks[0], "")[j]
			}
		}

		totalpriority += priority(char)

		i += 3
	}

	fmt.Println("Part 2")
	fmt.Printf("Total Priority: %d\n", totalpriority)

}
