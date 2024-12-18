package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 10000000000000

func solve(input []string, part int) int {

	var total int = 0

	for i:=0; i<len(input); i+=4 {

		// Parse the input
		buttonA := input[i]
		buttonAXY := strings.Split(buttonA, ": ")[1]
		_buttonAX := strings.Split(buttonAXY, ", ")[0]
		buttonAX, _  := strconv.Atoi(strings.Split(_buttonAX, "+")[1])
		_buttonAY := strings.Split(buttonAXY, ", ")[1]
		buttonAY, _  := strconv.Atoi(strings.Split(_buttonAY, "+")[1])

		buttonB := input[i+1]

		buttonBXY := strings.Split(buttonB, ": ")[1]
		_buttonBX := strings.Split(buttonBXY, ", ")[0]
		buttonBX, _  := strconv.Atoi(strings.Split(_buttonBX, "+")[1])
		_buttonBY := strings.Split(buttonBXY, ", ")[1]
		buttonBY, _  := strconv.Atoi(strings.Split(_buttonBY, "+")[1])

		prize   := input[i+2]
		prizeXY := strings.Split(prize, ": ")[1]
		_prizeX := strings.Split(prizeXY, ", ")[0]
		prizeX, _ := strconv.Atoi(strings.Split(_prizeX, "=")[1])
		_prizeY := strings.Split(prizeXY, ", ")[1]
		prizeY, _ := strconv.Atoi(strings.Split(_prizeY, "=")[1])

		if part == 2 {
			prizeX += 1e13
			prizeY += 1e13
		}

		// Solve the simultaneous equations
		// a*AX + b*BX = PX
		// a*AY + b*BY = PY

		// Solution 1

		// *AX, AY
		// a*AX*AY + b*BX*AY = PX*AY
		// a*AY*AX + b*BY*AX = PY*AX

		// Solve for a*AX*AY
		// PX*AY - b*BX*AY = PY*AX - b*BY*AX
		// b(BX*AY - BY*AX) = PX*AY - PY*AX

		// b = (PX*AY - PY*AX)/(BX*AY - BY*AX)
		// a = (PX - b*BX)/AX


		b := float64(prizeX*buttonAY - prizeY*buttonAX)/float64(buttonBX*buttonAY - buttonBY*buttonAX)
		a := (float64(prizeX) - b*float64(buttonBX))/float64(buttonAX)
		// If a or b are smaller than 0 or not an integer then
		// set this to the max value it can be
		if a<0 || b<0 || a != math.Floor(a) || b != math.Floor(b) {
			continue
		// Otherwise if we're in Part 1 and a or b are larger than 100
		} else if part == 1 && (a>100 || b>100) {
			continue
		// Else we meet the constraints
		} else {
			total += int(3*a + b)
		}

	}

	return total
}

func main() {

	input, _ := utils.ReadFile("input.txt")


	fmt.Println("Part 1:", solve(input, 1))
	fmt.Println("Part 2:", solve(input, 2))

}