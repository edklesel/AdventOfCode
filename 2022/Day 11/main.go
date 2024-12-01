package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	tools "github.com/edklesel/AdventOfCode/2022/GO/tools"
)

func removeItem(arr [][]int64) [][]int64 {
	var newArr = [][]int64{}
	for index, val := range arr {
		if index != 0 {
			newArr = append(newArr, val)
		}
	}
	return newArr
}

func addNumbers(arr1 []int64, arr2 []int64, mod int64) (result []int64) {

	start := time.Now()

	var maxLen int
	var arrBig []int64
	var arrSmall []int64

	if len(arr1) > len(arr2) {
		maxLen = len(arr1)
		arrBig = arr1
		arrSmall = arr2
	} else {
		maxLen = len(arr2)
		arrBig = arr2
		arrSmall = arr1
	}

	result = make([]int64, maxLen)

	for i := 0; i < maxLen; i++ {

		result[i] += arrBig[i]
		if len(arrSmall) > i {
			result[i] += arrSmall[i]
		}

	}

	end := time.Since(start)
	if end.Microseconds() > 1000000 {
		fmt.Printf("\naddNumbers took %v\n", time.Since(start))
		fmt.Println(arr1)
		fmt.Println(arr2)
	}

	return normaliseNumber(result, mod)

}

func multiplyNumbers(arr1 []int64, arr2 []int64, mod int64) (result []int64) {

	var square bool = true

	start := time.Now()

	// Initialise the result
	result = make([]int64, len(arr1)+len(arr2)-1)
	for i := 0; i < len(result); i++ {
		result[i] = 0
	}

	// Check if the arrays are the same
	if len(arr1) == len(arr2) {
		for index := range arr1 {
			if arr1[index] != arr2[index] {
				square = false
				break
			}
		}
	} else {
		square = false
	}

	if end := time.Since(start); end.Microseconds() > 1000 {
		fmt.Printf("Checked for a square in %v (arrsize = %d, %d)\n", time.Since(start), len(arr1), len(arr2))
	}

	// If we're squaring the number, do some optimisation
	if square {
		return normaliseNumber(sqaureNumber(arr1), mod)
	} else {
		// Otherwise we have to brute force it

		// Cycle through each element and multiply
		//	it with every other element in the other array
		for i := 0; i < len(arr1); i++ {
			for j := 0; j < len(arr2); j++ {
				result[i+j] += arr1[i] * arr2[j]
			}
		}

	}

	end := time.Since(start)
	if end.Microseconds() > 1000000 {
		fmt.Printf("multiplyNumbers (nonSquare) took %v\n", time.Since(start))
	}

	return normaliseNumber(result, mod)

}

func normaliseNumber(arr []int64, mod int64) []int64 {

	start := time.Now()

	var index int = 0

	for {

		// If the value is bigger than the modulus
		if arr[index] >= mod {

			// Check if we need to extend the number
			if len(arr)-1 < index+1 {
				arr = append(arr, 0)
			}
			arr[index+1] += arr[index] / mod
			arr[index] = arr[index] % mod
		}

		// If we've reached the end of the index, return
		if index == len(arr)-1 {
			break
		} else {
			index += 1
		}

	}

	end := time.Since(start)
	if end.Microseconds() > 1000000 {
		fmt.Printf("normaliseNumber took %v\n", time.Since(start))
		fmt.Println(arr)
	}
	return arr

}

func sqaureNumber(arr []int64) (result []int64) {

	var numCoefficients []int
	var num int
	var indexStart int
	var numSteps int
	var multiple int64

	start := time.Now()

	// Initialise the result
	result = make([]int64, 2*len(arr)-1)
	for i := 0; i < len(result); i++ {
		result[i] = 0
	}

	numCoefficients = make([]int, len(result))
	num = 0
	for index := range numCoefficients {
		if float32(index) < float32(len(numCoefficients))/2 {
			num += 1
		} else {
			num -= 1
		}
		numCoefficients[index] = num
	}

	for indexR := range result {

		indexStart = indexR / 2
		if indexR%2 == 0 {
			numSteps = (numCoefficients[indexR] + 1) / 2
			for step := 0; step < numSteps; step++ {
				if step == 0 {
					multiple = 1
				} else {
					multiple = 2
				}
				result[indexR] += multiple * arr[indexStart+step] * arr[indexStart-step]
			}
		} else if indexR%2 == 1 {
			numSteps = numCoefficients[indexR] / 2
			for step := 0; step < numSteps; step++ {
				result[indexR] += 2 * arr[indexStart-step] * arr[indexR-(indexStart-step)]
			}
		}
	}

	if end := time.Since(start); end.Microseconds() > 1000000 {
		fmt.Printf("multiplyNumbers (square) took %v (arrSize = %d)\n", time.Since(start), len(arr))
		// fmt.Println(arr)
	}

	return result

}

func manageMonkeys(rounds int, worry bool, input []string) int64 {

	var totalInspections = []int64{}
	var monkeysCurrent = [][][]int64{}
	var monkeyIndexes = []int{}
	var productDivisible int64 = 1
	var item []int64

	// Set up the slices
	monkeyMatch := regexp.MustCompile(`^Monkey \d+\:`)
	for index, row := range input {
		if monkeyMatch.MatchString(row) {
			monkeyIndexes = append(monkeyIndexes, index)

			// Multiply together the "divisible by" numbers
			divisibleBy, _ := strconv.Atoi(strings.Split(strings.Split(input[index+3], ": ")[1], " ")[2])
			productDivisible *= int64(divisibleBy)

		}
	}

	if worry {
		productDivisible *= 3
	}

	// monkeysStarting = make([][][]int64, len(monkeyIndexes))
	monkeysCurrent = make([][][]int64, len(monkeyIndexes))

	// Extract the starting items
	for monkeyNo, index := range monkeyIndexes {
		for _, _item := range strings.Split(strings.Split(input[index+1], ": ")[1], ", ") {
			_item, _ := strconv.Atoi(_item)

			item = []int64{int64(_item)}

			monkeysCurrent[monkeyNo] = append(monkeysCurrent[monkeyNo], normaliseNumber(item, productDivisible))
		}
	}
	var operation string
	var test string
	var test_true string
	var test_false string
	var itemNew []int64
	var valOp []int64
	var valTest int64
	var newMonkeyNo int
	var product int64

	totalInspections = make([]int64, len(monkeyIndexes))

	// fmt.Println("")
	// fmt.Println(0)
	// for index := range monkeysCurrent {
	// 	fmt.Println(monkeysCurrent[index])
	// }

	// Begin the rounds
	for i := 1; i <= rounds; i++ {

		for monkeyNo, index := range monkeyIndexes {

			// fmt.Println("\n")
			// fmt.Printf("Monkey %d\n", monkeyNo)

			operation = strings.Split(input[index+2], ": ")[1]
			test = strings.Split(input[index+3], ": ")[1]
			test_true = strings.Split(input[index+4], ": ")[1]
			test_false = strings.Split(input[index+5], ": ")[1]

			for _, item := range monkeysCurrent[monkeyNo] {

				// fmt.Println("")
				// fmt.Printf("item: %d\n", item)

				// Check if the value to be operated with is a number of
				//  the item itself
				switch strings.Split(operation, " ")[4] {
				case "old":
					valOp = item
				default:
					_valOp, _ := strconv.Atoi(strings.Split(operation, " ")[4])
					valOp = normaliseNumber([]int64{int64(_valOp)}, productDivisible)
				}

				// fmt.Printf("\tOp = %s, valOp = %d\n", strings.Split(operation, " ")[3], valOp)

				// Extract the operation
				switch strings.Split(operation, " ")[3] {
				case "*":
					itemNew = multiplyNumbers(item, valOp, productDivisible)
				case "+":
					itemNew = addNumbers(item, valOp, productDivisible)
				default:
					panic(fmt.Sprintf("Unexpected operation %s\n", strings.Split(operation, " ")[3]))
				}

				// fmt.Printf("\titemNew: %d (productDivisible = %d)\n", itemNew, productDivisible)

				// When the monkey gets bored, divide by 3
				// Hacked for part 1 because I cba to work out division in polynomial arithmetic
				if worry {
					product = 0
					for index, val := range itemNew {
						product += val * int64(math.Pow(float64(productDivisible), float64(index)))
					}

					itemNew = normaliseNumber([]int64{product / 3}, productDivisible)
					// fmt.Printf("\tAfter /3, itemNew = %d\n", itemNew)
				}

				// Extract the test
				_valTest, _ := strconv.Atoi(strings.Split(strings.Split(test, ": ")[0], " ")[2])
				valTest = int64(_valTest)

				// fmt.Printf("\tvalTest = %d\n", valTest)

				if itemNew[0] == 0 || itemNew[0]%valTest == 0 {
					newMonkeyNo, _ = strconv.Atoi(strings.Split(strings.Split(test_true, ": ")[0], " ")[3])
				} else {
					newMonkeyNo, _ = strconv.Atoi(strings.Split(strings.Split(test_false, ": ")[0], " ")[3])
				}

				// fmt.Printf("Thrown to monkey %d\n", newMonkeyNo)

				monkeysCurrent[newMonkeyNo] = append(monkeysCurrent[newMonkeyNo], itemNew)
				monkeysCurrent[monkeyNo] = removeItem(monkeysCurrent[monkeyNo])
				totalInspections[monkeyNo]++
			}

		}
		// fmt.Println("")
		fmt.Printf("Round %d\n", i)
		// for index := range monkeysCurrent {
		// 	fmt.Println(monkeysCurrent[index])
		// }
		// fmt.Println("")
		// fmt.Println(totalInspections)
		// var test string
		// fmt.Scanln(&test)
		// if i == 1 {
		// 	// break
		// }

	}

	sort.Slice(totalInspections, func(i, j int) bool { return totalInspections[i] > totalInspections[j] })
	return totalInspections[0] * totalInspections[1]

}

func main() {

	var filename string = "day11.test.txt"

	input := tools.ReadInput(filename)

	// fmt.Println("")
	// fmt.Println("Part 1")
	// fmt.Printf("Level of monkey business: %d\n", manageMonkeys(20, true, input))

	// fmt.Println(sqaureNumber([]int64{675602}))
	// fmt.Println(multiplyNumbers([]int64{675602}, []int64{675602}, 9327116929))

	// return

	fmt.Println("")
	fmt.Println("Part 2")
	fmt.Printf("Level of monkey business: %d\n", manageMonkeys(10000, false, input))

}
