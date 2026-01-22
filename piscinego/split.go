package piscine

func Split(s, sep string) []string {
	var result []string
	start := 0
	sepLen := len(sep)

	for i := 0; i+sepLen <= len(s); {
		if s[i:i+sepLen] == sep {
			result = append(result, s[start:i])
			i += sepLen
			start = i
		} else {
			i++
		}
	}
	result = append(result, s[start:])
	return result
}
