package main

import (
	"fmt"
	"strings"

	tools "github.com/edklesel/AdventOfCode/2022/GO/tools"
)

var filename string = "day2.txt"

func main() {

	input := tools.ReadInput(filename)

	var score int = 0

	var scores = make(map[string]int)
	scores["X"] = 1 //Rock
	scores["Y"] = 2 //Paper
	scores["Z"] = 3 //Scissors

	var opponent_move string
	var your_move string
	for i := 0; i < len(input); i++ {

		opponent_move = strings.Split(input[i], " ")[0]
		your_move = strings.Split(input[i], " ")[1]

		// Win scenarios
		var win = []string{"AY", "BZ", "CX"}
		var lose = []string{"AZ", "BX", "CY"}
		var tie = []string{"AX", "BY", "CZ"}

		// Win scenrarios
		score += scores[your_move]
		if tools.Find(win, opponent_move+your_move) {
			// fmt.Println("You Win")
			score += 6
		} else if tools.Find(lose, opponent_move+your_move) {
			// fmt.Println("You lose")
		} else if tools.Find(tie, opponent_move+your_move) {
			// fmt.Println("Tied game")
			score += 3
		}

	}

	fmt.Println("Part 1")
	fmt.Printf("Your final score: %d\n", score)

	var outcome string

	score = 0
	scores["A"] = 1 //Rock
	scores["B"] = 2 //Paper
	scores["C"] = 3 //Scissors

	for i := 0; i < len(input); i++ {

		opponent_move = strings.Split(input[i], " ")[0]
		outcome = strings.Split(input[i], " ")[1]

		switch outcome {
		case "X": //Lose
			switch opponent_move {
			case "A":
				score += scores["Z"]
			case "B":
				score += scores["X"]
			case "C":
				score += scores["Y"]
			}

		case "Y": //Tie
			score += 3 + scores[opponent_move]

		case "Z": //Win
			score += 6
			switch opponent_move {
			case "A":
				score += scores["Y"]
			case "B":
				score += scores["Z"]
			case "C":
				score += scores["X"]
			}

		}

	}

	fmt.Println("")
	fmt.Println("Part 2")
	fmt.Printf("Your final score: %d\n", score)

}
