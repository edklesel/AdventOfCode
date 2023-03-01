package main

import (
	"fmt"
	"strconv"
	"strings"

	tools "github.com/edklesel/AdventOfCode/2022/GO/tools"
)

var filename string = "day8.txt"

func checkDistance(dir string, x int, y int, input []string) (visible int, viewingDistance int) {

	var dmin int
	var dmax int
	var dint int
	var xmax int = len(input[0])
	var ymax int = len(input)
	var adjacentTreeHeight int
	var treeHeight int

	viewingDistance = 0

	treeHeight, _ = strconv.Atoi(strings.Split(input[y], "")[x])

	switch dir {
	case "l":
		dmin = -1
		dmax = -x - 1
		dint = -1
	case "r":
		dmin = 1
		dmax = xmax - x
		dint = 1
	case "u":
		dmin = -1
		dmax = -y - 1
		dint = -1
	case "d":
		dmin = 1
		dmax = ymax - y
		dint = 1
	default:
		panic("Please specify dir as one of l,r,u,d")
	}

	for d := dmin; d != dmax; d += dint {

		if dir == "l" || dir == "r" {
			adjacentTreeHeight, _ = strconv.Atoi(strings.Split(input[y], "")[x+d])
			// fmt.Printf(" (%d,%d) ", x+d, y)
		} else if dir == "u" || dir == "d" {
			adjacentTreeHeight, _ = strconv.Atoi(strings.Split(input[y+d], "")[x])
			// fmt.Printf(" (%d,%d) ", x, y+d)
		} else {
			panic("Please specify dir as one of l,r,u,d")
		}

		// fmt.Printf("%d (> %d?)\n", adjacentTreeHeight, treeHeight)
		viewingDistance++

		if adjacentTreeHeight >= treeHeight {
			// fmt.Println("Higher tree found")
			// fmt.Printf(" Direction: %s, Visible: %d, Visible Trees: %d\n", dir, visible, viewingDistance)
			return 0, viewingDistance
		}

	}
	// fmt.Printf(" Direction: %s, Visible: %d, Visible Trees: %d\n", dir, visible, viewingDistance)
	// fmt.Printf("Tree is visible from %s\n", dir)
	return 1, viewingDistance

}

func main() {

	input := tools.ReadInput(filename)

	var visibleDir int
	var visibleDirs int
	var noVisible int = 2*len(input) + 2*len(strings.Split(input[0], "")) - 4
	var noInvisible int = 0
	var viewingDistance int
	var viewingScore int = 0
	var viewingDistanceDir int

	for x := 1; x < len(input[0])-1; x++ {

		for y := 1; y < len(input)-1; y++ {

			// fmt.Printf("\n(%d,%d)\n", x, y)

			visibleDirs = 0
			viewingDistance = 1

			for _, dir := range [4]string{"l", "r", "u", "d"} {

				visibleDir, viewingDistanceDir = checkDistance(dir, x, y, input)

				visibleDirs += visibleDir
				viewingDistance *= viewingDistanceDir

			}

			if visibleDirs > 0 {
				noVisible++
			} else {
				// fmt.Printf("Tree at (%d,%d) is not visible.\n", x, y)
				noInvisible++
			}

			// fmt.Printf("Tree (%d,%d) has viewing score %d\n", x, y, viewingDistance)

			if viewingDistance > viewingScore {
				viewingScore = viewingDistance
			}

		}

	}

	fmt.Println("")
	fmt.Println("Part 1")
	fmt.Printf("Total number of visible trees: %d\n", noVisible)
	fmt.Printf("Total number of hidden trees: %d\n", noInvisible)

	fmt.Println("")
	fmt.Println("Part 2")
	fmt.Printf("Maximum viewing score: %d\n", viewingScore)

}
