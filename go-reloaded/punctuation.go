package main

import "regexp"

var spaceBeforePunct = regexp.MustCompile(`\s+([.,!?;:])`)
var spaceAfterPunct = regexp.MustCompile(`([.,!?;:])([^\s.,!?;:'])`)

func FixPunctuation(line string) string {
	line = spaceBeforePunct.ReplaceAllString(line, "$1")
	line = spaceAfterPunct.ReplaceAllString(line, "$1 $2")
	return line
}
