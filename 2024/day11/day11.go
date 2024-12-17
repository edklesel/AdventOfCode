package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func increment(counter map[int]int, stone int, inc int) map[int]int {
	if _, ok := counter[stone]; !ok {
		counter[stone] = 0
	}
	counter[stone] += inc
	return counter
}

func run(stones []int, blinks int, debug bool) int {

	var counter map[int]int = make(map[int]int)

	// First up fill the counter with the stones
	for _, stone := range(stones) {
		counter = increment(counter, stone, 1)
	}

	// Then loop through and increment each number
	var newCounter map[int]int
	
	if debug {fmt.Println(counter)}
	for range(blinks) {

		newCounter = make(map[int]int)

		for stone, count := range(counter) {
			if stone == 0 {
				newCounter = increment(newCounter, 1, count)
			} else if len(strconv.Itoa(stone)) % 2 == 0 {
				stoneStr := strconv.Itoa(stone)
				halfway := len(stoneStr)/2
				stone1, _ := strconv.Atoi(stoneStr[:halfway])
				stone2, _ := strconv.Atoi(stoneStr[halfway:])
				newCounter = increment(newCounter, stone1, count)
				newCounter = increment(newCounter, stone2, count)
			} else {
				newCounter = increment(newCounter, stone * 2024, count)
			}
		}

		counter = newCounter
		if debug {fmt.Println(counter)}
	}

	total := 0
	for _, count := range(counter) {
		total += count
	}

	return total

}

func main() {

	_input, _ := utils.ReadFile("input.txt")
	input := strings.Split(_input[0], " ")

	var stones []int = make([]int, 0)
	for _, stone := range(input) {
		num, _ := strconv.Atoi(stone)
		stones = append(stones, num)
	}

	fmt.Println("Part 1:", run(stones, 25, false))
	fmt.Println("Part 2:", run(stones, 75, false))

}