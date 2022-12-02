// This is the main file that will be used to call the solution to each days problem
package main

import (
	"fmt"

	"github.com/CosyOranges/AdventOfCode/2022/day01"
	"github.com/CosyOranges/AdventOfCode/src/parse"
	"github.com/CosyOranges/AdventOfCode/src/utils"
)

var SOLVED_DAYS [25]int

func main() {
	fmt.Println("Welcome to the Advent of Code!")

	day, year, aoc_cookie := parse.Parser()
	utils.GetInput(day, year, aoc_cookie)
	day01.Day_01()
}
