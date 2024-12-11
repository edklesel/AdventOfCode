package main

import (
	"aoc/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func step(trail [][]string, x int, y int, peaks *[]string, trails *int) {

	var gradCurrent, gradNext int

	// fmt.Println(x, y)

	gradCurrent, _ = strconv.Atoi(trail[y][x])

	canUp := y != 0
	canDown := y != len(trail)-1
	canLeft := x != 0
	canRight := x != len(trail[0])-1

	// Check up
	if canUp {
		gradNext, _ = strconv.Atoi(trail[y-1][x])
		coordNext := fmt.Sprintf("%d,%d", x, y-1)
		if gradCurrent == 8 && gradNext == 9 {
			if !slices.Contains(*peaks, coordNext) {
				*peaks = append(*peaks, coordNext)
			}
			*trails ++
		} else if gradNext == gradCurrent + 1 {
			step(trail, x, y-1, peaks, trails)
		}
	}

	// Check down
	if canDown {
		gradNext, _ = strconv.Atoi(trail[y+1][x])
		coordNext := fmt.Sprintf("%d,%d", x, y+1)
		if gradCurrent == 8 && gradNext == 9 {
			if !slices.Contains(*peaks, coordNext) {
				*peaks = append(*peaks, coordNext)
			}
			*trails ++
		} else if gradNext == gradCurrent + 1 {
			step(trail, x, y+1, peaks, trails)
		}
	}

	// Check left
	if canLeft {
		gradNext, _ = strconv.Atoi(trail[y][x-1])
		coordNext := fmt.Sprintf("%d,%d", x-1, y)
		if gradCurrent == 8 && gradNext == 9 {
			if !slices.Contains(*peaks, coordNext) {
				*peaks = append(*peaks, coordNext)
			}
			*trails ++
		} else if gradNext == gradCurrent + 1 {
			step(trail, x-1, y, peaks, trails)
		}
	}

	// Check Right
	if canRight {
		gradNext, _ = strconv.Atoi(trail[y][x+1])
		coordNext := fmt.Sprintf("%d,%d", x+1, y)
		if gradCurrent == 8 && gradNext == 9 {
			if !slices.Contains(*peaks, coordNext) {
				*peaks = append(*peaks, coordNext)
			}
			*trails ++
		} else if gradNext == gradCurrent + 1 {
			step(trail, x+1, y, peaks, trails)
		}
	}


}


func main() {

	input, _ := utils.ReadFile("input.txt")

	var trail [][]string = make([][]string, 0)
	var trailHeadsPotential [][2]int = make([][2]int, 0)

	// Find the trail heads
	for _, row := range(input) {
		trail = append(trail, strings.Split(row, ""))
	}

	for y := range(trail) {
		for x := range(trail[y]) {
			if trail[y][x] == "0" {
				trailHeadsPotential = append(trailHeadsPotential, [2]int{x, y})
			}
		}
	}

	// Loop through the trail heads
	var peaks []string
	var trails int
	var trailHeads map[string]int = make(map[string]int)
	var uniquePaths map[string]int = make(map[string]int)
	for _, trailHead := range(trailHeadsPotential) {

		peaks = make([]string, 0)
		trails = 0
		trailHeadLocation := fmt.Sprintf("%d,%d", trailHead[0], trailHead[1])
		step(trail, trailHead[0], trailHead[1], &peaks, &trails)

		if len(peaks) > 0 {
			trailHeads[trailHeadLocation] = len(peaks)
			uniquePaths[trailHeadLocation] = trails
		}

	}

	var total int = 0
	for _, score := range(trailHeads) {
		total += score
	}
	fmt.Println("Part 1:", total)

	total = 0
	for _, rating := range(uniquePaths) {
		total += rating
	}
	fmt.Println("Part 2:", total)

	
}