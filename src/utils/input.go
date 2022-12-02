package utils

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

func GetWithAOCCookie(url string, cookie string) []byte {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Build the request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatalf("Error making request: %s", err)
	}

	// Build the cookie and add to the request
	session_cookie := http.Cookie{
		Name:  "session",
		Value: cookie,
	}
	req.AddCookie(&session_cookie)

	// Fetch the response
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %s", err)
	}

	// Get the body of the response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading body of response: %s", err)
	}

	return body
}

func GetInput(day, year int, cookie string) {
	fmt.Printf("Fetching for day %d, year %d\n", day, year)

	// make the request
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	body := GetWithAOCCookie(url, cookie)

	if strings.HasPrefix(string(body), "Puzzle inputs differ by user") {
		panic("'Puzzle inputs differ by user' response")
	}

	// write to file
	filename := filepath.Join(GetHomeDir(), "AOC/data/", fmt.Sprintf("%d/day%02d/input.txt", year, day))

	WriteToFile(filename, body)

	fmt.Printf("Wrote the input data for %02d/25/%d to file: %s", day, year, filename)

	fmt.Println("Done!")
}
