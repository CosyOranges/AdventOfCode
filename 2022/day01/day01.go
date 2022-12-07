package day01

import (
	"fmt"
	"sort"
	"strings"

	"github.com/CosyOranges/AdventOfCode/src/maths"
	"github.com/CosyOranges/AdventOfCode/src/utils"
	"github.com/spf13/cast"
)

// variable to assign the input to
var in string

func transformInput() (ans [][]int) {
	in = utils.ReadInputFile("01")
	if len(in) == 0 {
		panic("Empty input!")
	}

	// We need to split initially on "\n\n" because elves are sperarated from the next
	// by an empty line
	for _, group := range strings.Split(in, "\n\n") {
		row := []int{}
		for _, line := range strings.Split(group, "\n") {
			row = append(row, cast.ToInt(line))
		}
		ans = append(ans, row)
	}

	return ans
}

// Main function to take input from a .txt file
func Day01Part1() {
	elves := transformInput()

	ans := 0

	for _, line := range elves {
		// TODO: Factor this out into a simple maths module
		sum := maths.SumIntSlice(line)

		if sum > ans {
			ans = sum
		}
	}
	fmt.Println("The Elf with the most calories is: ", ans)
}

func Day01Part2() {
	elves := transformInput()

	totals := []int{}

	// Create array of totals per elf
	for _, line := range elves {
		// TODO: Factor this out into a simple maths module
		sum := maths.SumIntSlice(line)
		totals = append(totals, sum)
	}

	// sort the totals
	sort.Ints(totals)

	// extract the top three as a sum
	ans := 0
	for i := 0; i < 3; i++ {
		ans += totals[len(totals)-1-i]
	}

	fmt.Println("The total for the top three elves is: ", ans)
}
