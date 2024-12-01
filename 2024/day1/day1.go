package main

import (
	"aoc24/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func parseInput(input string) ([]int, []int) {

	var raw_data []string
	raw_data, err := utils.ReadFile(input)
	if err != nil {
		panic(err)
	}

	var _left int
	var left []int

	var _right int
	var right []int

	var LR []string
	for _, line := range raw_data {

		LR = strings.Split(line, "   ")
		
		_left, _ = strconv.Atoi(LR[0])
		_right, _ = strconv.Atoi(LR[1])

		left = append(left, _left)
		right = append(right, _right)
	}

	return left, right

}

func main() {

	var left []int

	var right []int

	// Part 1
	left, right = parseInput("input.txt")

	sort.Ints(left)
	sort.Ints(right)

	var sum float64
	for i := range(left) {
		sum += math.Abs(float64(left[i]) - float64(right[i]))
	}

	fmt.Println("Part 1:", int64(sum))

	// Part 2
	var similarity int
	var occurrences map[string]int = make(map[string]int)

	left, right = parseInput("input.txt")

	for i := range(left) {
		leftNumber := left[i]
		leftStr := strconv.Itoa(leftNumber)
		if _, exists := occurrences[leftStr]; ! exists {
			occurrences[leftStr] = 0
			for _, rightNumber := range right {
				if leftNumber == rightNumber {
					occurrences[leftStr] += 1
				}
			}
		}
		similarity += leftNumber * occurrences[leftStr]
	}

	fmt.Println("Part 2:", similarity)

}
