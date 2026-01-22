package main

import "regexp"

var quoteRe = regexp.MustCompile(`'\s*(.*?)\s*'`)
var quote2Re = regexp.MustCompile(`"\s*(.*?)\s*"`)

func FixQuotes(line string) string {
	line = quoteRe.ReplaceAllString(line, "'$1'")
	line = quote2Re.ReplaceAllString(line, "\"$1\"")
	return line
}

//func FixQuotes2(line string) string {
//	return quote2Re.ReplaceAllString(line, "\"$1\"")
//}
