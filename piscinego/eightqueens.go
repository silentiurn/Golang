package piscine

import "github.com/01-edu/z01"

func printSolution(pos [8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(pos[i] + '0'))
	}
	z01.PrintRune('\n')
}

func isSafe(pos [8]int, row, col int) bool {
	for i := 0; i < row; i++ {
		if pos[i] == col || pos[i]-col == i-row || col-pos[i] == i-row {
			return false
		}
	}
	return true
}

func solve(pos [8]int, row int) {
	if row == 8 {
		printSolution(pos)
		return
	}
	for col := 1; col <= 8; col++ {
		if isSafe(pos, row, col) {
			pos[row] = col
			solve(pos, row+1)
		}
	}
}

func EightQueens() {
	var pos [8]int
	solve(pos, 0)
}
