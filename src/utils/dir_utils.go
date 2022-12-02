package utils

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	return homeDir
}

func GetCurrentDir() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("getting caller function")
	}

	return filename
}

func WriteToFile(filename string, content []byte) {
	err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	if err != nil {
		log.Fatalf("Error making dir: %s", err)
	}
	err = os.WriteFile(filename, content, os.FileMode(0644))
	if err != nil {
		log.Fatalf("Error writing file: %s", err)
	}
}
