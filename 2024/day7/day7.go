package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func operatorNumbers(nums []int, final int, tryConcat bool) bool {

	numConcat, _ := strconv.Atoi(
		fmt.Sprintf("%s%s", strconv.Itoa(nums[0]), strconv.Itoa(nums[1])),
	)

	if len(nums) == 2 {


		if nums[0] * nums[1] == final {
			return true
		} else if nums[0] + nums[1] == final {
			return true
		} else if tryConcat && (numConcat == final) {
			return true
		} else {
			return false
		}
		
	} else if operatorNumbers(append([]int{nums[0] + nums[1]}, nums[2:]...), final, tryConcat) {
		return true

	} else if operatorNumbers(append([]int{nums[0] * nums[1]}, nums[2:]...), final, tryConcat) {
		return true
	
	} else if tryConcat && operatorNumbers(append([]int{numConcat}, nums[2:]...), final, tryConcat) {
		return true
	}

	return false

}

func run(input []string, tryConcat bool) int {

	var total int = 0
	for _, row := range(input) {

		split := strings.Split(row, ": ")

		_final := split[0]
		final, _ := strconv.Atoi(_final)
		
		_nums := strings.Split(split[1], " ")
		var nums []int
		
		for _, num := range(_nums) {
			numInt, _ := strconv.Atoi(num)
			nums = append(nums, numInt)
		}

		if operatorNumbers(nums, final, tryConcat) {
			total += final
		}
		
	}

	return total
}

func main() {

	input, _ := utils.ReadFile("input.txt")

	// Part 1
	fmt.Println("Part 1:", run(input, false))

	// Part 2
	fmt.Println("Part 2:", run(input, true))

}