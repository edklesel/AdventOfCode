package main

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func getMinKey(dict map[int]int) int {
	var max int = 1e9
	for k, _ := range(dict) {
		if k < max {
			max = k
		}
	}
	return max
}

func sortMap(dict map[int]int) map[int]int {

	keys := make([]int, 0, len(dict))
	newDict := make(map[int]int)

	for k := range(dict) {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		newDict[k] = dict[k]
	}

	return newDict

}

func main() {

	_input, _ := utils.ReadFile("input copy.txt")
	input := _input[0]

	var blocks []string = make([]string, 0)
	var ID string

	for i, char := range(strings.Split(input, "")) {
		
		isFile := i%2 == 0
		isSpace := i%2 == 1

		num, _ := strconv.Atoi(char)

		if isFile {
			ID = strconv.Itoa(i/2)
		} else if isSpace {
			ID = "."
		}

		for j := 0; j < num; j++ {
			blocks = append(blocks, ID)
		}

	}

	var blocksStart []string = make([]string, len(blocks))
	_ = copy(blocksStart, blocks)

	// Part 1

	var charsTotal int = 0
	var charsOrdered int = 0
	for _, char := range(blocks) {
		if char != "." {
			charsTotal++
		}
	}

	// fmt.Println(strings.Join(blocks, ""))

	for i, char := range(blocks) {

		if char == "." {

			for j := len(blocks)-1; j > 0; j-- {

				if blocks[j] != "." {
					blocks[i] = blocks[j]
					blocks[j] = "."
					charsOrdered++
					break
				}
			}

			// fmt.Println(strings.Join(blocks, ""))

		} else {
			charsOrdered++
		}

		if charsTotal == charsOrdered {
			break
		}

	}

	var total int = 0
	for id, char := range(blocks) {

		num, _ := strconv.Atoi(char)

		total += id * num

	}

	fmt.Println("Part 1:", total)

	// Part 2
	blocks = make([]string, len(blocksStart))
	_ = copy(blocks, blocksStart)

	// Make a map of the file sizes and the indexes the file is at
	var fileSizes map[int]int = make(map[int]int)
	var fileIndexes map[int][]int = make(map[int][]int)
	var maxId int
	var inGap bool
	var gapId int
	var gapMap map[int]int = make(map[int]int)
	for i, char := range(blocks) {
		if char != "." {
			inGap = false
			num, _ := strconv.Atoi(char)

			maxId = num

			if _, ok := fileSizes[num]; !ok {
				fileSizes[num] = 0
			}
			fileSizes[num]++

			if _, ok := fileIndexes[num]; !ok {
				fileIndexes[num] = []int{}
			}
			fileIndexes[num] = append(fileIndexes[num], i)
		} else {
			if !inGap {
				gapId = i
			}
			if _, ok := gapMap[gapId]; !ok {
				gapMap[gapId] = 0
			}
			inGap = true
			gapMap[gapId] ++
		}
	}

	// fmt.Println(gapMap)

	// Iterate down the file sizes
	var fileLength, oldGapIndex, oldGapLength int
	var gapFound bool
	// var atId bool
	for fileId := maxId; fileId > 0; fileId-- {
		fileLength = fileSizes[fileId]

		// fmt.Println()
		
		// fmt.Println(strings.Join(blocks, ""))
		// fmt.Println(gapMap)

		// Find the first gap in the slice
		if getMinKey(gapMap) >= fileIndexes[fileId][0] {
				fmt.Println("No gaps to the left of file", fileId)
				continue
		}

		// Search the gap map to see if there are any big enough
		gapFound = false
		for gapIndex, gapLength := range(gapMap) {
			if gapLength >= fileSizes[fileId] {
				gapFound = true
				oldGapLength = gapMap[gapIndex]
				oldGapIndex = gapIndex
				break
			}
		}

		if !gapFound {
			fmt.Println("File",fileId,"is too big.")
			continue
		} else {
			fmt.Println("Found a gap big enough for file",fileId)
		}

		// Replace the old file in the string with the dots
		for _, i := range(fileIndexes[fileId]) {
			blocks[i] = "."
		}

		// Remove any gaps which start after this file
		for k := range(gapMap) {
			if k >= fileIndexes[fileId][0] {
				delete(gapMap, k)
			}
		}
		

		// Replace the dots in the gap with the file ID
		for i := oldGapIndex; i < oldGapIndex+fileLength; i++ {
			blocks[i] = strconv.Itoa(fileId)
		}

		// Clsoe up the gap in the gap map
		delete(gapMap, oldGapIndex)
		if oldGapLength > fileSizes[fileId] {
			newGapIndex := oldGapIndex + fileSizes[fileId]
			gapMap[newGapIndex] = oldGapLength - fileSizes[fileId]
			fmt.Println("Gap starting at", oldGapIndex, "is now size",gapMap[newGapIndex], "and starts at index", newGapIndex)
		}
		gapMap = sortMap(gapMap)
		fmt.Println(len(gapMap))

	}

	// Now calculate the checksum
	total = 0
	var total2 int64 = 0

	for i := range(blocks) {

		if blocks[i] == "." {
			continue
		}

		fileId, _ := strconv.Atoi(blocks[i])
		total2 += int64(fileId) * int64(i)

	}
	fmt.Println(strings.Join(blocks, ""))
	fmt.Println("Part 2:",total2)


}