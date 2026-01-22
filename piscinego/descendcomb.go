package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	hashmap := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for i := 99; i > -1; i-- {
		for j := i - 1; j > -1; j-- {
			z01.PrintRune(hashmap[i/10])
			z01.PrintRune(hashmap[i%10])
			z01.PrintRune(' ')
			z01.PrintRune(hashmap[j/10])
			z01.PrintRune(hashmap[j%10])
			if i != 1 || j != 0 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
