package piscine

func Rot14(s string) string {
	result := ""

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			result += string('a' + (char-'a'+14)%26)
		} else if char >= 'A' && char <= 'Z' {
			result += string('A' + (char-'A'+14)%26)
		} else {
			result += string(char) // Оставляем другие символы без изменений
		}
	}

	return result
}
