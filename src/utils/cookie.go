package utils

import (
	"log"
	"os"
	"strings"
)

func CookieChecker(aoc_cookie string) string {
	if len(aoc_cookie) == 0 {
		log.Println("No AOC Cookie provided, falling back to AOC_SESSION_COOKIE")
	} else {
		return aoc_cookie
	}

	aoc_cookie = os.Getenv("AOC_SESSION_COOKIE")
	if len(aoc_cookie) == 0 {
		log.Println("No AOC_SESSION_COOKIE environment variable, falling back to ~/.config/aoc/token")
	} else {
		return aoc_cookie
	}

	path := GetHomeDir() + "/.config/aoc/token"

	cookie, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error: %s", err)
		log.Fatalf("Error no token file found at path: %s", path)
	}

	// Trim any potential new lines:
	aoc_cookie = string(cookie)
	aoc_cookie = strings.TrimSuffix(aoc_cookie, "\n")

	return aoc_cookie
}
