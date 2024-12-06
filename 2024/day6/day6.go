package main

import (
	"aoc/utils"
	"fmt"
	"slices"
	"strings"
)

func guardWalk(guardMap [][]string) (bool, []string) {

	var guardDir, guardNewDir, guardPosition, guardPositionDirection string
	var guardX, guardY, guardNewX, guardNewY int
	var guardPositions []string = make([]string, 0)
	var guardPositionsDirections []string = make([]string, 0)

	// Find the Guard's initial position
	for i := range(guardMap) {
		for j := range(guardMap[i]) {
			if strings.Contains("^v<>", guardMap[i][j]) {
				guardX = j
				guardY = i
				guardDir= guardMap[i][j]
			}
		}
	}


	// Begin the game
	for {


		guardPosition = fmt.Sprintf("%d,%d", guardX, guardY)
		guardPositionDirection = fmt.Sprintf("%d,%d,%s", guardX, guardY, guardDir)

		if !slices.Contains(guardPositions, guardPosition) {
			guardPositions = append(guardPositions, guardPosition)
		}

		if slices.Contains(guardPositionsDirections, guardPositionDirection) {
			// fmt.Println("Guard got stuck in a loop!")
			return true, guardPositions
		} else {
			guardPositionsDirections = append(guardPositionsDirections, guardPositionDirection)
		}

		guardNewY, guardNewX = guardY, guardX

		// Work out the next position of the guard
		// and his new direction
		switch guardDir {
			case "^":
				guardNewY = guardY - 1
				guardNewDir = ">"
			case "v":
				guardNewY = guardY + 1
				guardNewDir = "<"
			case "<":
				guardNewX = guardX - 1
				guardNewDir = "^"
			case ">":
				guardNewX = guardX + 1
				guardNewDir = "v"
		}

		// First check if he'll be going out of bounds
		if guardNewX < 0 || guardNewX == len(guardMap[0]) || guardNewY < 0 || guardNewY == len(guardMap) {
			break

		// Next, check if he hits somthing, in which case
		// don't move him but turn him instead.
		} else if guardMap[guardNewY][guardNewX] == "#" {
			guardDir = guardNewDir

		// Otherwise, move him
		} else {
			guardX = guardNewX
			guardY = guardNewY
		}

	}

	return false, guardPositions

}

func parseMap(input []string) [][]string {

	// Create the map
	var guardMap [][]string = make([][]string, len(input), len(input[0]))
	for i, row := range(input) {
		guardMap[i] = strings.Split(row, "")
	}

	return guardMap

}

func main() {

	input, _ := utils.ReadFile("input.txt")

	// Part 1
	_, positions := guardWalk(parseMap(input))
	fmt.Println("Part 1:", len(positions))

	// Part 2
	var guardMapNew [][]string = make([][]string, len(input), len(input[0]))
	var loops int = 0
	for y := range(input) {
		for x := range(strings.Split(input[0], "")) {

			// We can skip Positions where the Guard never goes in the first place
			if !slices.Contains(positions, fmt.Sprintf("%d,%d", x, y)) {
				continue
			}

			guardMapNew = parseMap(input)
			if !strings.Contains("^<>v#", guardMapNew[y][x]) {
				guardMapNew[y][x] = "#"
			}

			if loop, _ := guardWalk(guardMapNew); loop {
				loops++
			}
		}
	}

	fmt.Println("Part 2:", loops)

}