package day02

import (
	"fmt"
	"io/ioutil"
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
	// read the entire file & return the byte slice as a string
	content, err := ioutil.ReadFile(utils.GetHomeDir() + "/AOC/data/2022/day02/input.txt") // TODO: Tidy up this side of things
	if err != nil {
		panic(err)
	}
	// trim off new lines and tabs at end of input files
	strContent := string(content)
	interim := strings.TrimRight(strContent, "\n")
	in = strings.TrimRight(interim, "\n")
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

	fmt.Printf("You would have won: %d points\n", score)
}
