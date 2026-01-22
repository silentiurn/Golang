package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	var digits []rune
	for n > 0 {
		digits = append(digits, rune(n%10)+'0')
		n /= 10
	}

	for i := 0; i < len(digits)-1; i++ {
		for j := 0; j < len(digits)-1-i; j++ {
			if digits[j] > digits[j+1] {
				digits[j], digits[j+1] = digits[j+1], digits[j]
			}
		}
	}

	for _, digit := range digits {
		z01.PrintRune(digit)
	}
}
