package piscine

import "github.com/01-edu/z01"

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	baseLen := len(base)

	if nbr == -9223372036854775808 {
		z01.PrintRune('-')
		PrintNbrBase(-(nbr / baseLen), base)
		z01.PrintRune(rune(base[-(nbr % baseLen)]))
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	var result []rune
	for nbr > 0 {
		remainder := nbr % baseLen
		result = append([]rune{rune(base[remainder])}, result...)
		nbr /= baseLen
	}

	if len(result) == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	for _, r := range result {
		z01.PrintRune(r)
	}
}
