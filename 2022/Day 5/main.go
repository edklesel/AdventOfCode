package main

import (
	"fmt"
	"strconv"
	"strings"

	tools "github.com/edklesel/AdventOfCode/2022/GO/tools"
)

var filename string = "day5.txt"

func removeCrate(index int, arr []string) (returnArr []string) {

	for i := 0; i < len(arr); i++ {

		if i == index {
			continue
		}

		returnArr = append(returnArr, arr[i])

	}

	return returnArr

}

func stackCrate(crate string, stack []string) (newStack []string) {

	newStack = []string{crate}
	newStack = append(newStack, stack...)

	return newStack

}

func parseStacks(input []string) (stacks [][]string, instructions []string) {
	var crates []string
	var breakIndex int

	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			breakIndex = i
			break
		}
		crates = append(crates, input[i])
	}

	for i := breakIndex + 1; i < len(input); i++ {
		instructions = append(instructions, input[i])
	}

	var rows []string = strings.Split(crates[len(crates)-1], " ")
	numStacks, _ := strconv.Atoi(rows[len(rows)-2])

	for i := 0; i < numStacks; i++ {
		stacks = append(stacks, []string{})
	}

	for i := 0; i < len(crates)-1; i++ {

		for j := 0; j < len(crates[i]); j += 4 {

			crate := crates[i][j : j+3]

			if crate == "   " {
				continue
			}

			crate = strings.Split(crate, "")[1]

			stacks[j/4] = append(stacks[j/4], crate)

		}

	}

	return stacks, instructions
}

func main() {

	input := tools.ReadInput(filename)

	// Model the initial stacks
	var stacks [][]string
	var instructions []string

	stacks, instructions = parseStacks(input)

	// Part 1

	var oldStack int
	var newStack int
	var numCrates int

	for i := 0; i < len(instructions); i++ {

		numCrates, _ = strconv.Atoi(strings.Split(instructions[i], " ")[1])
		oldStack, _ = strconv.Atoi(strings.Split(instructions[i], " ")[3])
		oldStack -= 1
		newStack, _ = strconv.Atoi(strings.Split(instructions[i], " ")[5])
		newStack -= 1

		for j := 0; j < numCrates; j++ {

			stacks[newStack] = stackCrate(stacks[oldStack][0], stacks[newStack])
			stacks[oldStack] = removeCrate(0, stacks[oldStack])

		}
	}

	var output string
	for i := 0; i < len(stacks); i++ {

		output += stacks[i][0]

	}

	fmt.Println("Part 1")
	fmt.Printf("Top Crates on all stacks: %s\n", output)

	// Part 2

	stacks, instructions = parseStacks(input)
	var cratesToMove []string

	for i := 0; i < len(instructions); i++ {

		numCrates, _ = strconv.Atoi(strings.Split(instructions[i], " ")[1])
		oldStack, _ = strconv.Atoi(strings.Split(instructions[i], " ")[3])
		oldStack -= 1
		newStack, _ = strconv.Atoi(strings.Split(instructions[i], " ")[5])
		newStack -= 1

		cratesToMove = []string{}

		for j := 0; j < numCrates; j++ {

			cratesToMove = append(cratesToMove, stacks[oldStack][0])
			stacks[oldStack] = removeCrate(0, stacks[oldStack])

		}

		stacks[newStack] = append(cratesToMove, stacks[newStack]...)

	}

	output = ""
	for i := 0; i < len(stacks); i++ {

		output += stacks[i][0]

	}

	fmt.Println("")
	fmt.Println("Part 2")
	fmt.Printf("Top Crates on all stacks: %s\n", output)

}
