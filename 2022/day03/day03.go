package day03

import (
	"fmt"
	"strings"

	"github.com/CosyOranges/AdventOfCode/src/utils"
)

var in string

func transformInput() (ans []string) {
	in = utils.ReadInputFile("03")
	if len(in) == 0 {
		panic("Empty input!")
	}

	// Now split by line
	ans = append(ans, strings.Split(in, "\n")...)

	return ans
}

func createMap(backpack string) map[byte]int {
	ans := map[byte]int{}
	for i := 0; i < len(backpack)/2; i++ {
		ans[backpack[i]] = 0
	}

	return ans
}

func addPriority(char byte) int {
	ascii := int(char)

	if ascii > 96 {
		return ascii - 96
	} else {
		return ascii - (65 - 27)
	}
}

/*
Sum of the items of priority in both backpack compartments
*/
func Day03Part1() {
	backpacks := transformInput()

	ans := 0
	for _, bag := range backpacks {
		first_compartment := createMap(bag)

		for j := len(bag) / 2; j < len(bag); j++ {
			if _, ok := first_compartment[bag[j]]; ok && first_compartment[bag[j]] < 1 {
				first_compartment[bag[j]] += 1
				ans += addPriority(bag[j])
			}
		}
	}

	fmt.Println("The sum of Priority items in backpacks is: ", ans)
}

/*
Something about badges?
*/
func Day03Part2() {

}
