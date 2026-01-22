package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n < 1 || n >= 10 {
		return
	}

	comb := make([]int, n)
	isFirst := true

	var generate func(index int, startDigit int)

	generate = func(index int, startDigit int) {
		if index == n {
			if !isFirst {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}

			for _, d := range comb {
				z01.PrintRune(int32(d) + '0')
			}

			isFirst = false
			return
		}

		maxEnd := 10 - (n - index) + 1

		for i := startDigit; i < maxEnd; i++ {
			comb[index] = i
			generate(index+1, i+1)
		}
	}

	generate(0, 0)

	z01.PrintRune('\n')
}
