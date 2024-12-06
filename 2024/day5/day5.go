package main

import (
	"aoc/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {

	input, _ := utils.ReadFile("input.txt")

	// Find the rules and the updates
	var rules, updates, correctUpdates, incorrectUpdates []string
	var updateCorrect bool = true
	var pageIndexes map[string]int
	var lNum, rNum string
	var lNumIndex, rNumIndex int
	var iterations, total int

	for i, row := range(input) {
		if row == "" {
			rules = input[:i]
			updates = input[i+1:]
			break
		}
	}

	for _, update := range(updates) {

		pages := strings.Split(update, ",")

		// Repeat until the update is correct
		iterations = 0
		updateCorrect = false
		for !updateCorrect {
			updateCorrect = true

			// Get the index of all the pages
			pageIndexes = make(map[string]int)
			for i, page := range(pages) {
				pageIndexes[page] = i
			}
			
			// Go through the rules and check them one by one
			for _, rule := range(rules) {
				
				lNum = strings.Split(rule, "|")[0]
				rNum = strings.Split(rule, "|")[1]

				if !(slices.Contains(pages, lNum) && slices.Contains(pages, rNum)) {
					continue
				}

				// If the left number comes after the right number,
				// the rule is incorrect, so try swapping the numbers.
				if pageIndexes[lNum] >= pageIndexes[rNum] {
					updateCorrect = false
					// fmt.Println("Update",strings.Join(pages, ","),"is incorrect")
					
					lNumIndex = pageIndexes[lNum]
					rNumIndex = pageIndexes[rNum]

					// Swap the numbers in the update
					pages[lNumIndex] = rNum
					pages[rNumIndex] = lNum

					// Swap the indexes in the index map
					pageIndexes[rNum] = lNumIndex
					pageIndexes[lNum] = rNumIndex

					break
				}
			}

			iterations++
			if updateCorrect {
				// fmt.Println("Update",strings.Join(pages, ","), "is correct")
				break
			} else {
			}
		}

		// If it was right first time, add it to the "Correct" bucket
		if iterations == 1 {
			correctUpdates = append(correctUpdates, update)
			// fmt.Println("Update",strings.Join(pages, ","),"was correct first time")
		// Otherwise it counts as an incorrect update
		} else {
			incorrectUpdates = append(incorrectUpdates, strings.Join(pages, ","))
			// fmt.Println("Update",strings.Join(pages, ","),"was correct after",iterations,"iterations.")
		}
	
	}

	// Part 1
	total = 0
	for _, correctUpdate := range(correctUpdates) {

		pages := strings.Split(correctUpdate, ",")
		middleNum, _ := strconv.Atoi(pages[(len(pages)-1)/2])
		total += middleNum
	}

	fmt.Println("Part 1:", total)

	// Part 2
	total = 0
	for _, incorrectUpdate := range(incorrectUpdates) {

		pages := strings.Split(incorrectUpdate, ",")
		middleNum, _ := strconv.Atoi(pages[(len(pages)-1)/2])
		total += middleNum
	}

	fmt.Println("Part 2:", total)

}