package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func deleteIndex(slice []int, index int) ([]int) {

	var newSlice []int = []int{}
	for i := 0; i < len(slice); i++ {
		if i == index {
			continue
		}
		newSlice = append(newSlice, slice[i])
	}

	return newSlice
}

func isRowSafe(row []int) (bool) {

	var rowInc, rowDec bool = false, false

	var lNum, rNum int
	for i := 0; i < len(row)-1; i++ {

		lNum = row[i]
		rNum = row[i+1]
		

		// If the numbers are the same, discard
		if rNum == lNum {
			return false
		}

		// If the numbers are more than 3 apart, discard
		if math.Abs(float64(rNum - lNum)) > 3 {
			return false
		}

		// Otherwise, inc or dec
		if rNum > lNum {
			rowInc = true
		} else if rNum < lNum {
			rowDec = true
		}

		if rowInc && rowDec {
			return false
		}
	}

	return true
}

func parseRow(row string) []int {

	var _numbers []string = strings.Split(row, " ")
	var numbers []int = []int{}

	for _, i := range(_numbers) {
		num, _ := strconv.Atoi(i)
		numbers = append(numbers, num)
	}

	return numbers
}

func main() {
	input, _ := utils.ReadFile("input.txt")
	var numbers []int
	var numSafe int = 0
	
	// Part 1
	for _, row := range(input) {
		
		numbers = parseRow(row)
		safe := isRowSafe(numbers)
		if safe {
			numSafe += 1
		}

	}

	fmt.Println("Part 1:", numSafe)

	// Part2
	numSafe = 0
	for _, row := range(input) {
		numbers := parseRow(row)

		if isRowSafe(numbers) {
			numSafe++
		} else {
			for i := 0; i < len(numbers); i++ {
				if isRowSafe(deleteIndex(numbers, i)) {
					numSafe++
					break
				}
			}
		}

	}

	fmt.Println("Part 2:", numSafe)
	
}