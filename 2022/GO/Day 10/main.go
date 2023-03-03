package main

import (
	"fmt"
	"strconv"
	"strings"

	tools "github.com/edklesel/AdventOfCode/2022/GO/tools"
)

var filename string = "day10.txt"

func sum(arr []int) (result int) {

	result = 0

	for _, val := range arr {
		result += val
	}

	return result

}

func checkInterval(cycleValues []int, registerX int, nCycles int, nCycleIntervals []int) []int {
	for _, interval := range nCycleIntervals {
		if interval == nCycles {
			return append(cycleValues, registerX*nCycles)
		}
	}

	return cycleValues
}

func addPixel(registerX int, nCycles int, crtOutput []string) []string {

	if len(crtOutput) < (nCycles/40)+1 {
		crtOutput = append(crtOutput, "")
	}

	if nCycles%40 == registerX || nCycles%40 == registerX+1 || nCycles%40 == registerX+2 {
		crtOutput[(nCycles-1)/40] += "#"
	} else {
		crtOutput[(nCycles-1)/40] += "."
	}

	return crtOutput

}

func main() {

	input := tools.ReadInput(filename)

	var registerX int = 1
	var incV int
	var nCycles int = 0
	var nCycleIntervals = []int{20, 60, 100, 140, 180, 220}
	var cycleValues = []int{}
	var crtOutput = []string{}

	for _, row := range input {

		nCycles += 1
		cycleValues = checkInterval(cycleValues, registerX, nCycles, nCycleIntervals)
		crtOutput = addPixel(registerX, nCycles, crtOutput)

		switch strings.Split(row, " ")[0] {
		case "noop":
		case "addx":

			nCycles += 1
			cycleValues = checkInterval(cycleValues, registerX, nCycles, nCycleIntervals)
			crtOutput = addPixel(registerX, nCycles, crtOutput)

			incV, _ = strconv.Atoi(strings.Split(row, " ")[1])
			registerX += incV

		}

	}

	fmt.Println("")
	fmt.Println("Part 1")
	fmt.Printf("Sum of register values at 20th cycles: %d\n", sum(cycleValues))

	fmt.Println("")
	fmt.Println("Part 2")
	fmt.Println("")
	fmt.Println(strings.Join(crtOutput, "\n"))

}
