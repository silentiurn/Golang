package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programName := os.Args[0]

	lastSlash := -1
	for i, char := range programName {
		if char == '/' {
			lastSlash = i
		}
	}

	if lastSlash != -1 {
		programName = programName[lastSlash+1:]
	}

	for _, char := range programName {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
