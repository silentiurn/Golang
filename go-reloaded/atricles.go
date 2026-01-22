package main

import (
	"regexp"
	"strings"
)

var anExceptions = map[string]bool{
	"hour":      true,
	"honor":     true,
	"heir":      true,
	"honest":    true,
	"honorable": true,
	"heirloom":  true,
	"honestly":  true,
	"honesty":   true,
}

var aExceptions = map[string]bool{
	"university": true,
	"user":       true,
	"unit":       true,
	"european":   true,
	"one":        true,
	"and":        true,
	"or":         true,
}

var articleRe = regexp.MustCompile(`\b([Aa][Nn]?)\s+([A-Za-z'][A-Za-z]+)`)

// функция подбора артикля
func replaceArticle(m string) string {
	sub := articleRe.FindStringSubmatch(m)
	if len(sub) < 3 {
		return m
	}

	article := sub[1]
	word := sub[2]

	lw := strings.ToLower(word)
	lw = strings.Trim(lw, "'\"")

	// проверяем исключения
	shouldUseAn := false
	if anExceptions[lw] {
		shouldUseAn = true
	} else if aExceptions[lw] {
		shouldUseAn = false
	} else {
		switch lw[0] {
		case 'a', 'e', 'i', 'o', 'u':
			shouldUseAn = true
		}
	}

	resultArticle := "a"
	if shouldUseAn {
		resultArticle = "an"
	}

	//регистр артикля
	if article[0] >= 'A' && article[0] <= 'Z' {
		resultArticle = strings.Title(resultArticle)
	}

	return resultArticle + " " + word
}

func FixArticles(line string) string {
	return articleRe.ReplaceAllStringFunc(line, replaceArticle)
}

/*func FixArticles(line string) string {
	words := strings.Fields(line)

	for i := 0; i < len(words)-1; i++ {
		w := strings.ToLower(words[i])

		if w != "a" && w != "an" {
			continue
		}

		next := strings.ToLower(words[i+1])
		next = strings.Trim(next, "'\"")

		shouldUseAn := false

		if anExceptions[next] {
			shouldUseAn = true
		} else if aExceptions[next] {
			shouldUseAn = false
		} else if len(next) > 0 {
			switch next[0] {
			case 'a', 'e', 'i', 'o', 'u':
				shouldUseAn = true
			}
		}

		if shouldUseAn {
			if words[i][0] >= 'A' && words[i][0] <= 'Z' {
				words[i] = "An"
			} else {
				words[i] = "an"
			}
		} else {
			if words[i][0] >= 'A' && words[i][0] <= 'Z' {
				words[i] = "A"
			} else {
				words[i] = "a"
			}
		}
	}

	return strings.Join(words, " ")
}
*/
