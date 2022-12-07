package day02

import (
	"fmt"
	"strings"

	"github.com/CosyOranges/AdventOfCode/src/utils"
)

/*
Scoring Guide:
A 	X	1	= Rock
B 	Y	2	= Paper
C 	Z	3	= Scissors

Win		= 6
Draw	= 3
Lost	= 0
*/

var in string

const (
	Win      = 6
	Draw     = 3
	Loss     = 0
	Rock     = 1
	Paper    = 2
	Scissors = 3
)

func match(p1 string, p2 string) int {
	return 0
}

func transformInput() [][]string {
	in = strings.TrimRight(utils.ReadInputFile("02"), "\n")
	if len(in) == 0 {
		panic("Empty input!")
	}

	var ans [][]string

	for _, line := range strings.Split(in, "\n") {
		ans = append(ans, strings.Split(line, " "))
	}

	return ans
}

func Day02Part1() {
	games := transformInput()

	choices := map[string]int{
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}

	score := 0
	for _, l := range games {
		// Always add score for our choice
		score += choices[l[1]]
		// Just use nested switch cases... ugly but it works
		switch l[1] {
		case "X": // Rock
			switch l[0] {
			case "A": // Rock
				{
					score += Draw
				}
			case "B":
				{
					score += Loss
				}
			case "C":
				{
					score += Win
				}
			}
		case "Y": // Paper
			switch l[0] {
			case "A": // Rock
				{
					score += Win
				}
			case "B":
				{
					score += Draw
				}
			case "C":
				{
					score += Loss
				}
			}
		case "Z": // Scissors
			switch l[0] {
			case "A": // Rock
				{
					score += Loss
				}
			case "B":
				{
					score += Win
				}
			case "C":
				{
					score += Draw
				}
			}
		}
	}

	fmt.Printf("PART 1: You would have won: %d points\n", score)
}

func Day02Part2() {
	games := transformInput()

	choices := map[string]int{
		"X": Loss,
		"Y": Draw,
		"Z": Win,
	}

	score := 0

	for _, l := range games {
		score += choices[l[1]]

		switch l[1] {
		case "X":
			{
				switch l[0] {
				case "A":
					{
						score += Scissors
					}
				case "B":
					{
						score += Rock
					}
				case "C":
					{
						score += Paper
					}
				}
			}
		case "Y":
			{
				switch l[0] {
				case "A":
					{
						score += Rock
					}
				case "B":
					{
						score += Paper
					}
				case "C":
					{
						score += Scissors
					}
				}
			}
		case "Z":
			{
				switch l[0] {
				case "A":
					{
						score += Paper
					}
				case "B":
					{
						score += Scissors
					}
				case "C":
					{
						score += Rock
					}
				}
			}
		}
	}

	fmt.Printf("PART 2: You would have won: %d points\n", score)
}
