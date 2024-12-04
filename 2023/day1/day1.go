package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	input, _ := utils.ReadFile("input.txt")

	// Part 1
	var total int = 0
	for _, row := range(input) {

		var firstNum, lastNum int = -1, -1
		for _, char := range(strings.Split(row, "")) {

			num, err := strconv.Atoi(char)
			if err == nil {
				if firstNum == -1 {
					firstNum = num
				}
				lastNum = num
			}
		}

		bigNumStr := fmt.Sprintf("%d%d",firstNum,lastNum)
		bigNum, _ := strconv.Atoi(bigNumStr)

		total += bigNum

	}

	fmt.Println("Part 1:", total)

	// Part 2
	var numbers map[string]int = map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
	}

	total = 0
	for _, row := range(input) {

		var firstNum, lastNum int = -1, -1
		chars := strings.Split(row, "")
		for i, char := range(chars) {

			// First check if it's a number
			num, err := strconv.Atoi(char)
			if err == nil {
				if firstNum == -1 {
					firstNum = num
				}
				lastNum = num
			} else {
				for numberName, num := range(numbers) {

					numberLen := len(numberName)

					// If we're too near the end of the row, skip it
					if i + numberLen > len(chars) {
						continue
					}
					numberString := strings.Join(chars[i:i+numberLen], "")

					if numberString == numberName {
						if firstNum == -1 {
							firstNum = num

						}
						lastNum = num
						break
					}

				}
			}
		}

		bigNumStr := fmt.Sprintf("%d%d",firstNum,lastNum)
		bigNum, _ := strconv.Atoi(bigNumStr)

		total += bigNum

	}

	fmt.Println("Part 2:", total)

}