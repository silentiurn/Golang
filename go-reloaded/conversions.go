package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func ProcessCaseModifiers(line string) string {
	// Сначала нормализуем все модификаторы, убирая лишние пробелы
	line = normalizeModifiers(line)

	words := strings.Fields(line)

	for i := 0; i < len(words); i++ {
		word := words[i]

		// Одиночные модификаторы: (up), (low), (cap)
		if len(word) > 2 && word[0] == '(' && word[len(word)-1] == ')' {
			mod := word[1 : len(word)-1]

			switch mod {
			case "up", "low", "cap", "hex", "bin":
				if i > 0 {
					switch mod {
					case "up":
						words[i-1] = strings.ToUpper(words[i-1])

					case "low":
						words[i-1] = strings.ToLower(words[i-1])

					case "cap":
						words[i-1] = strings.Title(strings.ToLower(words[i-1]))

					case "hex":
						if n, err := strconv.ParseInt(words[i-1], 16, 64); err == nil {
							words[i-1] = strconv.FormatInt(n, 10)
						}

					case "bin":
						if n, err := strconv.ParseInt(words[i-1], 2, 64); err == nil {
							words[i-1] = strconv.FormatInt(n, 10)
						}
					}
				}

				words = append(words[:i], words[i+1:]...)
				i--
				continue
			}
		}

		// Модификаторы с параметрами: (up,2), (low,3), (cap,4)
		if len(word) > 4 && word[0] == '(' && strings.Contains(word, ",") {
			content := word[1 : len(word)-1]
			parts := strings.Split(content, ",")

			if len(parts) == 2 {
				mod := parts[0]
				nStr := parts[1]

				if mod == "up" || mod == "low" || mod == "cap" {
					n, err := strconv.Atoi(nStr)
					if err == nil {
						count := 0

						for j := i - 1; j >= 0 && count < n; j-- {
							if !isWord(words[j]) {
								continue
							}

							switch mod {
							case "up":
								words[j] = strings.ToUpper(words[j])
							case "low":
								words[j] = strings.ToLower(words[j])
							case "cap":
								words[j] = strings.Title(strings.ToLower(words[j]))
							}

							count++
						}

						words = append(words[:i], words[i+1:]...)
						i--
					}
				}
			}
		}
	}

	return strings.Join(words, " ")
}

// функция убирает пробелы в модификаторах, оставляя пробелы между словами
func normalizeModifiers(line string) string {
	// Регулярное выражение для поиска модификаторов с пробелами
	modRe := regexp.MustCompile(`\(\s*([^)]+?)\s*\)`)

	line = modRe.ReplaceAllStringFunc(line, func(m string) string {
		content := m[1 : len(m)-1]
		content = strings.ReplaceAll(content, " ", "")
		content = strings.ToLower(content)
		return "(" + content + ")"
	})
	// пробелы вокруг модификатора до и после
	line = regexp.MustCompile(
		`(\S)(\((up|low|cap|hex|bin)(,\d+)?\))`,
	).ReplaceAllString(line, `$1 $2`)

	line = regexp.MustCompile(
		`(\((up|low|cap|hex|bin)(,\d+)?\))(\S)`,
	).ReplaceAllString(line, `$1 $4`)

	return line
}

func isWord(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return true
		}
	}
	return false
}
