package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {

		args := os.Args[1:]

		n := len(args)
		for i := 0; i < n-1; i++ {
			for j := 0; j < n-1-i; j++ {
				if args[j] > args[j+1] {
					args[j], args[j+1] = args[j+1], args[j]
				}
			}
		}

		for _, arg := range args {
			for _, char := range arg {
				z01.PrintRune(char)
			}
			z01.PrintRune('\n')
		}
	}
}
