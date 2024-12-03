package main

import (
	"aoc24/utils"
	"fmt"
	"regexp"
	"strconv"
)

func main() {

	input, _ := utils.ReadFile("input.txt")

	// Part 1
	re1 := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	
	var sum int = 0
	
	for _, row := range(input) {

		matches := re1.FindAllStringSubmatch(row, -1)

		for _, match := range(matches) {

			lNum, _ := strconv.Atoi(match[1])
			rNum, _ := strconv.Atoi(match[2])

			sum += (lNum * rNum)
		}
	}
	fmt.Println("Part 1:", sum)

	// Part 2
	re2 := regexp.MustCompile(`((mul)\((\d+),(\d+)\)|do(n't)?)`)
	
	sum = 0

	var enabled bool = true
	for _, row := range(input) {

		matches := re2.FindAllStringSubmatch(row, -1)

		for _, match := range(matches) {

			if match[2] == "mul" {
				
				if enabled {

					lNum, _ := strconv.Atoi(match[3])
					rNum, _ := strconv.Atoi(match[4])

					sum += (lNum * rNum)

				}

			} else if match[1] == "don't" {
				enabled = false
			} else if match[1] == "do" {
				enabled = true
			} else {
				fmt.Println(match)
				panic("Failed to parse match.")
			}

		}

	}

	fmt.Println("Part 2:", sum)

}