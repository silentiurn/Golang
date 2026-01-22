package main

import "github.com/01-edu/z01"

type ptr struct {
	x int
	y int
}

func setPoint(ptr *ptr) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	message := [15]rune{
		'x', ' ', '=', ' ',
		'0' + 42/10,
		'0' + 42%10,
		',', ' ', 'y', ' ', '=', ' ',
		'0' + 21/10,
		'0' + 21%10,
		'\n',
	}
	for i := 0; i < 15; i++ {
		z01.PrintRune(message[i])
	}
}
