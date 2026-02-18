package main

import (
	"bufio"
	"os"
	"strings"
)

func loadBanner(filename string) (map[rune][]string, error) {
	sym := make(map[rune][]string)
	file, err := os.Open("banners/" + filename + ".txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	j := rune(32)
	i := 0
	for scanner.Scan() {
		if i == 0 {
			i++
			continue
		}
		if i%9 == 0 {
			j++
			i++
			continue
		}
		i++
		sym[j] = append(sym[j], scanner.Text())
	}
	return sym, nil
}

func GenerateASCII(text string, bannerName string) (string, error) {
	sym, err := loadBanner(bannerName)
	if err != nil {
		return "", err
	}
	text = strings.ReplaceAll(text, "\r\n", "\n")
	var output strings.Builder
	height := 8
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		if line == "" {
			output.WriteString("\n")
			continue
		}
		runeStr := []rune(line)
		for i := 0; i < height; i++ {
			for _, ch := range runeStr {
				if charLines, ok := sym[ch]; ok {
					output.WriteString(charLines[i])
				}
			}
			output.WriteString("\n")
		}
	}
	return output.String(), nil
}
