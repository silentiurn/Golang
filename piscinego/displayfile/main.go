package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("File name missing")
	} else if len(args) >= 2 {
		fmt.Println("Too many arguments")
	} else {
		file, err := os.ReadFile(args[0]) // Используем os.ReadFile
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Print(string(file))
		}
	}
}
