package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	input, _ := utils.ReadFile("input.txt")

	var gameId int
	var gamePossible bool
	var total, total2 int = 0, 0

	var maxCubes map[string]int = map[string]int{
		"red": 12,
		"green": 13,
		"blue": 14,
	}
	var minColour map[string]int

	for i, game := range(input) {

		gameId = i+1
		gamePossible = true
		minColour = map[string]int{
			"red": 0,
			"green": 0,
			"blue": 0,
		}

		_hands := strings.Replace(game, fmt.Sprintf("Game %d: ", gameId), "", 1)
		hands := strings.Split(_hands, "; ")

		for _, hand := range(hands) {
			cubes := strings.Split(hand, ", ")

			for _, cube := range(cubes) {

				count, _ := strconv.Atoi(strings.Split(cube, " ")[0])
				colour := strings.Split(cube, " ")[1]

				// Check for Part 1
				if count > maxCubes[colour] {
					gamePossible = false
				}

				// Check for Part 2
				if count > minColour[colour] {
					minColour[colour] = count
				}
			}
		}

		// Check for Part 1
		if gamePossible {
			total += gameId
		}

		// Check for Part 2
		total2 += (minColour["red"] * minColour["green"] * minColour["blue"])
	}

	fmt.Println("Part 1:", total)
	fmt.Println("Part 2:", total2)
}