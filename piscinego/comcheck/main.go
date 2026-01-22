package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	for _, value := range arg {
		if value == "01" || value == "galaxy" || value == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
	return
}
