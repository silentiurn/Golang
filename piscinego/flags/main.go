package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Пузырьковая сортировка
func bubbleSort(str string) string {
	runes := []rune(str)
	n := len(runes)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	return string(runes)
}

// Проверка, начинается ли строка с префикса
func hasPrefix(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}

// Печать помощи
func printHelp() {
	helpText := `--insert
  -i
	 This flag inserts the string into the string passed as argument.
--order
  -o
	 This flag will behave like a boolean, if it is called it will order the argument.`
	for _, ch := range helpText {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) == 1 || os.Args[1] == "--help" || os.Args[1] == "-h" {
		printHelp()
		return
	}

	args := os.Args[1:]
	insertStr := ""
	orderFlag := false

	for i := 0; i < len(args); i++ {
		a := args[i]

		if hasPrefix(a, "--insert=") {
			insertStr = a[9:] // длина "--insert="
		} else if hasPrefix(a, "-i=") {
			insertStr = a[3:] // длина "-i="
		} else if a == "--insert" || a == "-i" {
			if i+1 < len(args) {
				insertStr = args[i+1]
				i++
			}
		} else if a == "--order" || a == "-o" {
			orderFlag = true
		}
	}

	mainStr := args[len(args)-1]

	if insertStr != "" {
		mainStr = mainStr + insertStr
	}
	if orderFlag {
		mainStr = bubbleSort(mainStr)
	}

	for _, ch := range mainStr {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
