package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/CosyOranges/AdventOfCode/src/utils"
)

// Main function to take input from a .txt file
func Day_01() {
	readFile, err := os.Open(utils.GetHomeDir() + "/AOC/data/2022/day01/input.txt")

	// If we can open the file successfully then read it line by line
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var sum int = 0

	var ans int = 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if len(line) == 0 {
			if sum > ans {
				ans = sum
			}

			sum = 0
		}

		tmp, _ := strconv.Atoi(line)
		sum += tmp
	}

	readFile.Close()

	fmt.Println("The most calories held by an elf is: ", ans)
}
