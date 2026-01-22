package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			z01.PrintRune('9')
			n = -223372036854775808
		}
		n = -n
	}

	var digits []rune
	for n > 0 {
		d := rune(n%10 + '0')
		digits = append([]rune{d}, digits...)
		n /= 10
	}

	for _, r := range digits {
		z01.PrintRune(r)
	}
}
