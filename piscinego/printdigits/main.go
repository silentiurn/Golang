package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for i := 0; i < 10; i++ {
		r := rune(i + '0')
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
