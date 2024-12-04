package main

import (
	"aoc24/utils"
	"fmt"
	"strings"
)

func main() {

	input, _ := utils.ReadFile("input.txt")

	// Make the grid
	var grid [][]string = make([][]string, len(input))
	for i := 0; i < len(input); i++ {
		grid[i] = strings.Split(input[i], "")
	}
	var rowLength int = len(grid[0])
	var canLeft, canRight, canUp, canDown bool
	var found int = 0
	var debug bool = false

	// Part 1
	// Traverse the grid looking for Xs
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if grid[i][j] == "X" {

				canLeft = j >= 3
				canRight = j < rowLength-3
				canUp = i >= 3
				canDown = i < len(input)-3

				// Check up
				if canUp {
					if grid[i-1][j] == "M" && grid[i-2][j] == "A" && grid[i-3][j] == "S" {
						if debug {fmt.Printf("XMAS Found up from (%d,%d)\n",i,j)}
						found++
					}
				}

				// Check down
				if canDown {
					if grid[i+1][j] == "M" && grid[i+2][j] == "A" && grid[i+3][j] == "S" {
						if debug {fmt.Printf("XMAS Found down from (%d,%d)\n",i,j)}
						found++
					}
				}

				// Check Right
				if canRight {
					if grid[i][j+1] == "M" && grid[i][j+2] == "A" && grid[i][j+3] == "S" {
						if debug {fmt.Printf("XMAS Found right from (%d,%d)\n",i,j)}
						found++
					}
				}

				// Check Left
				if canLeft {
					if grid[i][j-1] == "M" && grid[i][j-2] == "A" && grid[i][j-3] == "S" {
						if debug {fmt.Printf("XMAS Found left from (%d,%d)\n",i,j)}
						found++
					}
				}
				
				// Check up and left
				if canUp && canLeft {
					if grid[i-1][j-1] == "M" && grid[i-2][j-2] == "A" && grid[i-3][j-3] == "S" {
						if debug {fmt.Printf("XMAS Found up and left from (%d,%d)\n",i,j)}
						found++
					}
				}
				
				// Check up and right
				if canUp && canRight {
					if grid[i-1][j+1] == "M" && grid[i-2][j+2] == "A" && grid[i-3][j+3] == "S" {
						if debug {fmt.Printf("XMAS Found up and right from (%d,%d)\n",i,j)}
						found++
					}
				}
				
				// Check down and right
				if canDown && canRight {
					if grid[i+1][j+1] == "M" && grid[i+2][j+2] == "A" && grid[i+3][j+3] == "S" {
						if debug {fmt.Printf("XMAS Found down and right from (%d,%d)\n",i,j)}
						found++
					}
				}
				
				// Check down and left
				if canDown && canLeft {
					if grid[i+1][j-1] == "M" && grid[i+2][j-2] == "A" && grid[i+3][j-3] == "S" {
						if debug {fmt.Printf("XMAS Found down and left from (%d,%d)\n",i,j)}
						found++
					}
				}

			}
		}
	}

	fmt.Println("Part 1:", found)

	// Part 2
	// Traverse the grid looking for As
	var masFound int = 0
	var masCheck int
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {

			masCheck = 0

			if grid[i][j] == "A" {

				// if debug {fmt.Printf("A found at (%d,%d)\n",i,j)}

				canLeft = j >= 1
				canRight = j < rowLength-1
				canUp = i >= 1
				canDown = i < len(input)-1

				// If we're along the edges, skip
				if !canLeft || !canRight || !canUp || !canDown {
					continue
				}

				// \
				if (grid[i-1][j+1] == "M" && grid[i+1][j-1] == "S") ||
					(grid[i-1][j+1] == "S" && grid[i+1][j-1] == "M") {
					masCheck++
				}

				// /
				if (grid[i-1][j-1] == "M" && grid[i+1][j+1] == "S") ||
					(grid[i-1][j-1] == "S" && grid[i+1][j+1] == "M") {
					masCheck++
				}

				if masCheck == 2 {
					if debug {fmt.Printf("X-MAS found at (%d,%d)\n",i,j)}
					masFound++
				}

			}
		}
		
	}

	fmt.Println("Part 2:",masFound)
	
}