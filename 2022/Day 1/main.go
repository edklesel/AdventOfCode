package main

import (
	"fmt"
	"sort"
	"strconv"

	tools "github.com/edklesel/AdventOfCode/2022/GO/tools"
)

var filename string = "day1.txt"

func main() {

	var input []string
	input = tools.ReadInput(filename)

	var elf_no int32 = 0
	var elves = [][2]int32{}
	var max = [2]int32{0, 0}
	var max_arr = [][2]int32{}

	elves = append(elves, [2]int32{0, 0})
	for i := 0; i < len(input); i++ {

		if input[i] == "" {

			if elves[elf_no][1] > max[1] {
				max = [2]int32{elf_no, elves[elf_no][1]}
				max_arr = append(max_arr, max)
			}

			elf_no++
			elves = append(elves, [2]int32{elf_no, 0})
			continue
		} else {
			weight, _ := strconv.Atoi(input[i])
			elves[elf_no][1] += int32(weight)
		}

	}

	fmt.Println("Part 1")
	fmt.Printf("Maximum weight: %d at elf number %d\n", max[1], max[0])
	fmt.Println("\n\n")

	sort.Slice(elves, func(i, j int) bool {
		return elves[i][1] > elves[j][1]
	})

	fmt.Println("Part 2")
	const top_x_elves int = 3
	var top_x_elves_total int32 = 0
	fmt.Printf("Top %d elves carrying the most weight.\n", top_x_elves)
	for i := 0; i < top_x_elves; i++ {
		fmt.Printf("Number %d: Elf %d with weight %d\n", i, elves[i][0], elves[i][1])
		top_x_elves_total += elves[i][1]
	}

	fmt.Println("")
	fmt.Printf("Total weight from the top %d elves: %d\n", top_x_elves, top_x_elves_total)

}
