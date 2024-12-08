package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"slices"
	"strings"
)

func hypotenuse(vec []int) (float64) {
	return math.Sqrt(float64(vec[0]*vec[0]) + float64(vec[1]*vec[1])) 
}

func getYMXC(pos1 []int, pos2 []int) (float64, float64) {

	// Here we want to get the y = mx + c form of the line connecting the points
	dx := float64(pos2[0]-pos1[0])
	dy := float64(pos2[1]-pos1[1])

	m := dy/dx

	c := float64(pos1[1]) - m*float64(pos1[0])

	return m, c

}

func main() {

	input, _ := utils.ReadFile("input.txt")

	// Create the map of Antinodes
	var antinodes [][]int
	for i := range(input) {
		antinodes = append(antinodes, []int{})
		for range(strings.Split(input[0], "")) {
			antinodes[i] = append(antinodes[i], 0)
		}
	}

	// Find all the antennas
	var antennas map[string][][]int = make(map[string][][]int, 0)
	for y, row := range(input) {
		chars := strings.Split(row, "")
		for x, char := range(chars) {

			// Skip blank spots
			if char == "." {
				continue
			} else {
				antennas[char] = append(antennas[char], []int{x, y})
			}

		}
	}

	// Loop through the spaces on the board
	var uniquePos []string = make([]string, 0)
	var uniquePos2 []string = make([]string, 0)

	for y := range(input) {
		for x := range(strings.Split(input[0], "")) {

			// Current location
			location := fmt.Sprintf("%d,%d", x, y)

			// Loop through all the pairs of antennas and find the antinodes
			for antennaType := range(antennas) {

				for _, a1 := range(antennas[antennaType]) {

					coord1 := fmt.Sprintf("%d,%d",a1[0],a1[1])

					// If the location is the same as the antenna,
					// it can't be an antinode
					if coord1 == location {
						continue
					}

					for _, a2 := range(antennas[antennaType]) {

						coord2 := fmt.Sprintf("%d,%d",a2[0],a2[1])

						// Also skip if the antennas
						// we're comparing are the same
						if coord1 == coord2 {
							continue
						}

						// If the location is on one of the coordinates, add it
						// as an antinode for Part 2
						if (location == coord1 || location == coord2) && !slices.Contains(uniquePos2, location) {
							uniquePos2 = append(uniquePos2, location)
						}

						// Get the distance between the antennas
						antennaDist := hypotenuse([]int{a1[0]-a2[0], a1[1]-a2[1]})

						// Get the vectors between the position and each antenna
						// and claculate how far they are from each antenna
						vec1 := []int{a1[0]-x, a1[1]-y}
						vec2 := []int{a2[0]-x, a2[1]-y}
						dist1 := hypotenuse(vec1)
						dist2 := hypotenuse(vec2)

						// Find the largest one, which we use to check how many multiples
						// of the atenna's distance away it is
						dist := math.Max(dist1, dist2)

						m1, c1 := getYMXC(a1, []int{x, y})
						m2, c2 := getYMXC(a2, []int{x, y})

						// if (c1 == c2) && (m1 == m2) {
						if (math.Abs(c1 - c2) <= 1e-9) && (math.Abs(m1 - m2) <= 1e-9) {
						
							// Part 1
							if (dist / antennaDist == 2) && !slices.Contains(uniquePos, location) {
								uniquePos = append(uniquePos, location)
							}
							
							// Part 2
							if !slices.Contains(uniquePos2, location) {
								uniquePos2 = append(uniquePos2, location)
							}
						}

					}
				}
			}

		}
	}

	fmt.Println("Part 1:", len(uniquePos))
	fmt.Println("Part 2:", len(uniquePos2))

}