package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func main() {

	input, _ := utils.ReadFile("test.txt")

	var debug bool = false

	// Find the grid
	var grid [][]string = make([][]string, 0)
	var paths []string
	for i := range(input) {
		if input[i] == "" {
			continue
		} else if strings.Split(input[i], "")[0] == "#" {
			grid = append(grid, strings.Split(input[i], ""))
		} else if strings.Split(input[i], "")[0] != "" {
			paths = append(paths, strings.Split(input[i], "")...)
		}
	}

	// Find the robots starting position and all the boxes
	var robot []int
	for i := range(grid) {
		for j := range(grid[i]) {
			if grid[i][j] == "@" {
				robot = []int{j, i}
			}
		}
	}

	var canMove bool
	var newX, newY, dx, dy int
	for _, path := range(paths) {

		if debug {
			for _, row := range grid {
				fmt.Println(strings.Join(row, ""))
			}
			fmt.Println()
			fmt.Println(path)
		}

		canMove = false
		switch path {
			case "<":
				newX = robot[0]-1
				newY = robot[1]
				dx = -1
				dy = 0
			case ">":
				newX = robot[0]+1
				newY = robot[1]
				dx = 1
				dy = 0
			case "^":
				newX = robot[0]
				newY = robot[1]-1
				dx = 0
				dy = -1
			case "v":
				newX = robot[0]
				newY = robot[1]+1
				dx = 0
				dy = 1
			}

		nextGrid := grid[newY][newX]
		// If he's walking into a wall, skip this step
		if nextGrid == "#" {
			continue
		// If the next step is a box, see if there's a gap on the other eise of the box
		} else if nextGrid == "O" {
			for i:=1; i<len(grid[i]); i++ {
				// If there's a gap, move the box to the gap and leave the space blank
				gridStep := grid[newY + dy*i][newX + dx*i]
				if gridStep == "." {
					grid[newY + dy*i][newX + dx*i] = "O"
					grid[newY][newX] = "."
					canMove = true
					break
				} else if gridStep == "#" {
					canMove = false
					break
				} else if gridStep == "O" {
					continue
				}
			}
		} else if nextGrid == "." {
			canMove = true
		}

		if canMove {
			grid[robot[1]][robot[0]] = "."
			robot = []int{newX, newY}
			grid[robot[1]][robot[0]] = "@"
		}

	}

	if debug {
		for _, row := range grid {
			fmt.Println(strings.Join(row, ""))
		}
	}

	// Now get the GPS coordinates
	var total int = 0
	for i := range grid {
		for j := range grid[i] {

			if grid[i][j] != "O" {
				continue
			}

			total += 100*i + j

		}
	}

	fmt.Println("Part 1:", total)

}