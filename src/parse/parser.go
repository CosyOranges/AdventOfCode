package parse

import (
	"flag"
	"log"
	"time"

	"github.com/CosyOranges/AdventOfCode/src/utils"
)

func Parser() (day int, year int, aoc_cookie string) {
	today := time.Now()
	flag.IntVar(&day, "day", today.Day(), "day number to fetch, 1-25")
	flag.IntVar(&year, "year", today.Year(), "AOC year")
	// defaults to env variable
	flag.StringVar(&aoc_cookie, "aoc_cookie", "", "AOC session cookie")
	flag.Parse()

	// Check users inputs
	if day > 25 || day < 1 {
		log.Fatalf("Day out of range [%d], must be between 1 and 25 [inclusive]", day)
	}

	if year < 2022 {
		log.Fatal("No solutions available for years earlier than 2022 :(")
	}

	aoc_cookie = utils.CookieChecker(aoc_cookie)

	return day, year, aoc_cookie
}
