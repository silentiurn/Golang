package main

import (
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}
		return
	}

	for _, fileName := range args {
		file, err := os.Open(fileName)
		if err != nil {

			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}

		defer file.Close()

		_, err = io.Copy(os.Stdout, file)
		if err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}
	}
}
