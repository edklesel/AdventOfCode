package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	tools "github.com/edklesel/AdventOfCode/2022/GO/tools"
)

var filename string = "day9.txt"

func snakeButNotSnake(input []string, nKnots int) (nTailPositions int) {

	var positions = [][]int{}
	var tPositions = []string{}
	var tPositionString string
	var dx int
	var dy int

	for i := 0; i < nKnots; i++ {
		positions = append(positions, []int{0, 0})
	}

	var dir string
	var steps int
	var increment int
	var coord int

	for _, row := range input {

		dir = strings.Split(row, " ")[0]
		steps, _ = strconv.Atoi(strings.Split(row, " ")[1])

		switch dir {
		case "R":
			increment = 1
			coord = 0
		case "L":
			increment = -1
			coord = 0
		case "U":
			increment = 1
			coord = 1
		case "D":
			increment = -1
			coord = 1
		}

		for step := 0; step < steps; step++ {

			// Move the head
			positions[0][coord] += increment

			// Move the knots
			for knot := 1; knot < nKnots; knot++ {

				if math.Pow(math.Abs(float64(positions[knot][0])-float64(positions[knot-1][0])), 2)+math.Pow(math.Abs(float64(positions[knot][1])-float64(positions[knot-1][1])), 2) > 2 {

					switch positions[knot-1][0] - positions[knot][0] {
					case 2:
						dx = 1
					case -2:
						dx = -1
					default:
						dx = 0
					}

					switch positions[knot-1][1] - positions[knot][1] {
					case 2:
						dy = 1
					case -2:
						dy = -1
					default:
						dy = 0
					}

					positions[knot][0] += positions[knot-1][0] - positions[knot][0] - dx
					positions[knot][1] += positions[knot-1][1] - positions[knot][1] - dy

				}

				if knot == nKnots-1 {

					tPositionString = fmt.Sprintf("(%d,%d)", positions[knot][0], positions[knot][1])

					if !tools.Find(tPositions, tPositionString) {
						tPositions = append(tPositions, tPositionString)
					}
				}

			}

		}

	}

	return len(tPositions)

}

func main() {

	input := tools.ReadInput(filename)

	fmt.Println("")
	fmt.Println("Part 1")
	fmt.Printf("Total unique tail positions: %d\n", snakeButNotSnake(input, 2))

	fmt.Println("")
	fmt.Println("Part 2")
	fmt.Printf("Total unique tail positions: %d\n", snakeButNotSnake(input, 10))

}
