// This is the main file that will be used to call the solution to each days problem
package main

import (
	"fmt"

	"github.com/CosyOranges/AdventOfCode/2022/day01"
	"github.com/CosyOranges/AdventOfCode/2022/day02"
	"github.com/CosyOranges/AdventOfCode/src/parse"
	"github.com/CosyOranges/AdventOfCode/src/utils"
)

func answerPrinter(day int) {
	fmt.Printf("---------------------\n| ANSWERS for DAY %d |\n---------------------\n", day)
}

func main() {
	fmt.Println("Welcome to the Advent of Code!")

	day, year, aoc_cookie := parse.Parser()
	utils.GetInput(day, year, aoc_cookie)

	// Using a switch case here for the day...
	// Is this the best idea? 25 days... kinda repetitive...
	// TODO: Scope out a better way to solve this?
	switch day {
	case 1:
		{
			answerPrinter(day)
			day01.Day01Part1()
			day01.Day01Part2()
		}
	case 2:
		{
			answerPrinter(day)
			day02.Day02Part1()
		}
	}

}
