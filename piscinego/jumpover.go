package piscine

func JumpOver(str string) string {
	if len(str) < 3 {
		return "\n"
	}
	var result []rune
	for i := 2; i < len(str); i += 3 {
		result = append(result, rune(str[i]))
	}
	result = append(result, '\n')
	return string(result)
}
