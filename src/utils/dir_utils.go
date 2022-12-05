package utils

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

/*
Get the users home directory (UNIX).
*/
func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	return homeDir
}

/*
Get the current working directory.
*/
func GetCurrentDir() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("getting caller function")
	}

	return filename
}

/*
Write an array of bytes to the ~/AOC/data/year/day/.txt file.
*/
func WriteToFile(filename string, content []byte) {
	// https://gosamples.dev/create-directory/#:~:text=To%20create%20a%20single%20directory,MkdirAll()%20.
	err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	if err != nil {
		log.Fatalf("Error making dir: %s", err)
	}

	// Using 0644 file mode for the bytes we get back from GetWithAOCCookie()
	// https://stackoverflow.com/questions/24811770/how-to-write-to-a-file-in-golang
	err = os.WriteFile(filename, content, os.FileMode(0644))
	if err != nil {
		log.Fatalf("Error writing file: %s", err)
	}
}
