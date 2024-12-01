package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	tools "github.com/edklesel/AdventOfCode/2022/GO/tools"
)

var filename string = "day7.txt"

func main() {

	input := tools.ReadInput(filename)

	var dirTree = []string{}
	var dirSizes = make(map[string]int)
	var newDir string
	var cdDir string

	var re_MatchFile = regexp.MustCompile(`\d+ \w+(\.\w+)?`)

	for _, command := range input {

		if command == "$ cd .." {

			dirTree = dirTree[:len(dirTree)-1]

		} else if strings.Contains(command, "$ cd") {

			cdDir = strings.Split(command, " ")[2]

			if len(dirTree) == 0 {

				dirTree = append(dirTree, cdDir)
				newDir = cdDir

			} else {

				newDir = dirTree[len(dirTree)-1] + "_" + strings.Split(command, " ")[2]
				dirTree = append(dirTree, newDir)

			}

			if _, keyExists := dirSizes[newDir]; !keyExists {

				dirSizes[newDir] = 0

			}
		} else if re_MatchFile.MatchString(command) {

			for _, dirInDirTree := range dirTree {

				fileSize, _ := strconv.Atoi(strings.Split(command, " ")[0])
				dirSizes[dirInDirTree] += fileSize

			}

		} else if command == "$ ls" || strings.Contains(command, "dir ") {

			continue

		} else {

		}

	}

	var total int = 0
	for _, val := range dirSizes {
		total += val
	}

	fmt.Println("Part 1")
	fmt.Printf("Total directory size: %d\n", total)

	var totalSize int = 70000000
	var requiredSpace int = 30000000
	var currentFree = totalSize - dirSizes["/"]
	var needToFree = requiredSpace - currentFree
	var currentMinimum = 1000000000000

	for _, val := range dirSizes {

		if val < currentMinimum && val > needToFree {
			currentMinimum = val
		}

	}

	fmt.Println("")
	fmt.Println("Part 2")
	fmt.Printf("Smallest directory size to delete: %d\n", currentMinimum)

}
