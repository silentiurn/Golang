package main

import (
	"os"

	"github.com/01-edu/z01"
)

func stringToInt(str string) (int, bool) {
	result := 0
	for _, char := range str {
		if char < '0' || char > '9' {
			return 0, false
		}
		result = result*10 + int(char-'0')
	}
	return result, true
}

func main() {
	upper := false
	if len(os.Args) > 1 && os.Args[1] == "--upper" {
		upper = true

		os.Args = os.Args[1:]
	}

	if len(os.Args) == 1 {
		return
	}

	for _, arg := range os.Args[1:] {
		n, valid := stringToInt(arg)
		if !valid || n < 1 || n > 26 {
			z01.PrintRune(' ')
		} else {
			letter := rune('a' + n - 1)
			if upper {
				letter = rune('A' + n - 1)
			}
			z01.PrintRune(letter)
		}
	}

	z01.PrintRune('\n')
}
