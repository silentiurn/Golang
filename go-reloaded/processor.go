package main

func ProcessLine(line string) string {

	line = FixQuotes(line)
	//line = FixQuotes2(line)
	line = ProcessCaseModifiers(line)
	line = FixQuotes(line)

	line = FixPunctuation(line)
	line = FixArticles(line)

	return line
}
