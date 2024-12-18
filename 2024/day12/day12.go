package main

import (
	"aoc/utils"
	"fmt"
	"slices"
	"strings"
)

func coord(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func appendUnique(dst []string, src ...string) []string {

	for _, i := range src {
		if !slices.Contains(dst, i) {
			dst = append(dst, i)
		}
	}

	return dst

}

func dfs(x int, y int, garden [][]string, coordinates []string, perimiter int, corners int) ([]string, int, int) {

	var coordinate string = coord(x,y)

	coordinates = append(coordinates, coordinate)
	// fmt.Println(coordinates)

	var canUp, canDown, canLeft, canRight bool
	canUp = y > 0
	canDown = y < len(garden)-1
	canLeft = x > 0
	canRight = x < len(garden[y])-1

	var letter string = garden[y][x]
	var _coordinates []string

	// Check if we're at a corner, because nCorners = nEdges

	// Left edge
	if !canLeft {
		// If we can't go up then we're at an edge corner
		if !canUp {
			corners++
		// Otherwise if the letter above is not the same, we're at a corner
		} else if garden[y-1][x] != letter {
			corners++
		}

		if !canDown {
			corners++
		} else if garden[y+1][x] != letter {
			corners++
		}
	}

	// Right edge
	if !canRight {
		if !canUp {
			corners++
		} else if garden[y-1][x] != letter {
			corners++
		}

		if !canDown {
			corners++
		} else if garden[y+1][x] != letter {
			corners++
		}
	}

	// Top edge
	if !canUp {
		if !canRight {
			// Already checked this
		} else if garden[y][x+1] != letter {
			corners++
		}

		if !canLeft {
			// Already checked this
		} else if garden[y][x-1] != letter {
			corners++
		}
	}

	// Bottom edge
	if !canDown {
		if !canRight {
			// Already checked this
		} else if garden[y][x+1] != letter {
			corners++
		}

		if !canLeft {
			// Already checked this
		} else if garden[y][x-1] != letter {
			corners++
		}
	}

	// Non-edge convex corners
	// Up Right
	if canUp && canRight {
		if garden[y][x+1] != letter && garden[y-1][x] != letter {
			corners++
		}
	}
	// Up Left
	if canUp && canLeft {
		if garden[y][x-1] != letter && garden[y-1][x] != letter {
			corners++
		}
	}
	// Down Right
	if canDown && canRight {
		if garden[y][x+1] != letter && garden[y+1][x] != letter {
			corners++
		}
	}
	// Down Left
	if canDown && canLeft {
		if garden[y][x-1] != letter && garden[y+1][x] != letter {
			corners++
		}
	}

	// Concave corners
	// Up Right
	if canUp && canRight {
		if garden[y][x+1] == letter && garden[y-1][x] == letter && garden[y-1][x+1] != letter {
			corners++
		}
	}
	// Up Left
	if canUp && canLeft {
		if garden[y][x-1] == letter && garden[y-1][x] == letter && garden[y-1][x-1] != letter {
			corners++
		}
	}
	// Down Right
	if canDown && canRight {
		if garden[y][x+1] == letter && garden[y+1][x] == letter && garden[y+1][x+1] != letter {
			corners++
		}
	}
	// Down Left
	if canDown && canLeft {
		if garden[y][x-1] == letter && garden[y+1][x] == letter && garden[y+1][x-1] != letter {
			corners++
		}
	}


	// Now continue on to the full scan

	// Up
	if canUp {
		// If it's a different letter then add a perimiter
		if garden[y-1][x] != letter {
			perimiter++
		// Otherwise if it hasn't already been scanned, scan it
		} else if !slices.Contains(coordinates, coord(x,y-1)) {
			_coordinates, perimiter, corners = dfs(x, y-1, garden, coordinates, perimiter, corners)
			coordinates = appendUnique(coordinates, _coordinates...)
		}
	// If we're on the egde, add a perimiter
	} else {
		perimiter++

	}

	// Down
	if canDown {
		if garden[y+1][x] != letter {
			perimiter++
		} else if !slices.Contains(coordinates, coord(x,y+1)) {
			_coordinates, perimiter, corners = dfs(x, y+1, garden, coordinates, perimiter, corners)
			coordinates = appendUnique(coordinates, _coordinates...)
		}
	} else {
		perimiter++
	}

	// Left
	if canLeft {
		if garden[y][x-1] != letter {
			perimiter++
		} else if !slices.Contains(coordinates, coord(x-1,y)) {
			_coordinates, perimiter, corners = dfs(x-1, y, garden, coordinates, perimiter, corners)
			coordinates = appendUnique(coordinates, _coordinates...)
		}
	} else {
		perimiter++
	}

	// Right
	if canRight {
		if garden[y][x+1] != letter {
			perimiter++
		} else if !slices.Contains(coordinates, coord(x+1,y)) {
			_coordinates, perimiter, corners = dfs(x+1, y, garden, coordinates, perimiter, corners)
			coordinates = appendUnique(coordinates, _coordinates...)
		}
	} else {
		perimiter++
	}

	return coordinates, perimiter, corners

}

func main() {

	input, _ := utils.ReadFile("input.txt")

	var garden [][]string = make([][]string, 0)

	for _, row := range input {
		garden = append(garden, strings.Split(row, ""))
	}

	var zoneArea, zonesScanned []string
	var perimiter, corners int

	var total1, total2 int = 0, 0

	for y := range garden {
		for x := range garden[y] {

			if slices.Contains(zonesScanned, coord(x, y)) {
				continue
			}

			zoneArea, perimiter, corners = dfs(x, y, garden, []string{}, 0, 0)

			total1 += len(zoneArea)*perimiter
			total2 += len(zoneArea)*corners

			zonesScanned = append(zonesScanned, zoneArea...)

			// fmt.Printf("Zone %s, Area=%d, perimiter=%d, corners=%d\n", garden[y][x], len(zoneArea), perimiter, corners)

		}
	}

	fmt.Println("Part 1:", total1)
	fmt.Println("Part 2:", total2)
}