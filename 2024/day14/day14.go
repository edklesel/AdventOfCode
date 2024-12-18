package main

import (
	"aoc/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func getPosition(row string) []int {
	position := strings.Split(strings.Split(strings.Split(row, " ")[0], "=")[1], ",")
	x, _ := strconv.Atoi(position[0])
	y, _ := strconv.Atoi(position[1])
	return []int{x, y}
}

func getVelocity(row string) []int {
	velocity := strings.Split(strings.Split(strings.Split(row, " ")[1], "=")[1], ",")
	x, _ := strconv.Atoi(velocity[0])
	y, _ := strconv.Atoi(velocity[1])
	return []int{x, y}
}

func countDistinct(robots [][]int) int {

	var distinct []string = make([]string, 0)
	for _, robot := range robots {

		robotStr := fmt.Sprintf("%s,%x", robot[0], robot[1])
		if !slices.Contains(distinct, robotStr) {
			distinct = append(distinct, robotStr)
		}

	}

	return len(distinct)
}

func scanQuadrants(robots [][]int, height int, width int) map[string]int {

	var tb, rl, quadrant string
	var quadrants map[string]int = map[string]int{
		"tr": 0,
		"br": 0,
		"tl": 0,
		"bl": 0,
	}

	midHeight := (height-1)/2
	midWidth  := (width-1)/2
	for _, robot := range robots {

		// If they're along the middle line, skip them
		if robot[0] == midWidth || robot[1] == midHeight {
			continue
		}

		if robot[0] < midWidth {
			rl = "l"
		} else {
			rl = "r"
		}

		if robot[1] < midHeight {
			tb = "t"
		} else {
			tb = "b"
		}

		quadrant = fmt.Sprintf("%s%s", tb, rl)
		quadrants[quadrant]++

	}

	return quadrants

}

func main() {

	input, _ := utils.ReadFile("input.txt")

	var height int = 103
	var width int = 101
	var total int = 0

	// Get the initial positions of the robots
	var robots [][]int = make([][]int, 0)
	for _, r := range(input) {
		robots = append(robots, getPosition(r))
	}

	var seconds int = 0
	for {
		seconds++

		// Update the robots position
		for i:=0; i < len(robots); i++ {

			v := getVelocity(input[i])

			// Update the X and Y positions
			robots[i][0] += v[0]
			robots[i][1] += v[1]

			// Move it to the other side of the grid
			for {
				if robots[i][0] >= width {
					robots[i][0] -= width
				} else if robots[i][0] < 0 {
					robots[i][0] += width
				} else {
					break
				}
			}
			for {
				if robots[i][1] >= height {
					robots[i][1] -= height
				} else if robots[i][1] < 0 {
					robots[i][1] += height
				} else {
					break
				}
			}

		}

		if seconds==100 {
			safety := 1
			for _, n := range(scanQuadrants(robots, height, width)) {
				safety *= n
			}
			total += safety

			fmt.Println("Part 1", total)
		}

		if len(robots) == countDistinct(robots) {
			fmt.Println("Part 2:", seconds)
				grid := make([][]string, 0)
				for j := range(height) {
					grid = append(grid, []string{})
					for range(width) {
						grid[j] = append(grid[j], ".")
					}
				}
				for _, robot := range(robots) {
					grid[robot[1]][robot[0]] = "X"
				}
				for j := range(height) {
					fmt.Println(strings.Join(grid[j], ""))
				}
			break
		}

	}

}