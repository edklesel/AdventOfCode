package main

import (
	"fmt"
	"strconv"
	"strings"

	tools "github.com/edklesel/AdventOfCode/2022/GO/tools"
)

var filename string = "day4.txt"

func main() {

	input := tools.ReadInput(filename)

	var total int = 0
	var total2 int = 0

	var breakj bool

	for i := 0; i < len(input); i++ {

		var elf1 string = strings.Split(input[i], ",")[0]
		var elf2 string = strings.Split(input[i], ",")[1]

		elf1_l, _ := strconv.Atoi(strings.Split(elf1, "-")[0])
		elf1_r, _ := strconv.Atoi(strings.Split(elf1, "-")[1])
		elf2_l, _ := strconv.Atoi(strings.Split(elf2, "-")[0])
		elf2_r, _ := strconv.Atoi(strings.Split(elf2, "-")[1])

		if (elf1_l <= elf2_l && elf1_r >= elf2_r) ||
			(elf2_l <= elf1_l && elf2_r >= elf1_r) {
			total++
		}

		for j := elf1_l; j <= elf1_r; j++ {
			breakj = false
			for k := elf2_l; k <= elf2_r; k++ {
				if j == k {
					breakj = true
					total2++
					break
				}
			}
			if breakj {
				break
			}
		}

	}

	fmt.Println("Part 1")
	fmt.Printf("Total ranges fully contained: %d\n", total)

	fmt.Println("")
	fmt.Println("Part 2")
	fmt.Printf("Total ranges overlapping: %d\n", total2)

}
