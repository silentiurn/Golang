package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func run() error {
	if len(os.Args) != 3 {
		return errors.New("usage: go run . <sample.txt> <result.txt>")
	}

	if os.Args[1] == os.Args[2] {
		return errors.New("input and output files must be different")
	}

	in, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(os.Args[2])
	if err != nil {
		return err
	}
	defer out.Close()

	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)
	defer writer.Flush()

	for scanner.Scan() {
		line := ProcessLine(scanner.Text())
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
