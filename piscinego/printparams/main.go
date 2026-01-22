package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			for _, char := range arg {
				z01.PrintRune(char)
			}
			z01.PrintRune('\n')
		}
	}
}
