package piscine

func TrimAtoi(s string) int {
	sign := 1
	num := 0
	foundDigit := false

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char == '-' && !foundDigit {
			sign = -1
		} else if char >= '0' && char <= '9' {
			num = num*10 + int(char-'0')
			foundDigit = true
		}
	}

	return num * sign
}
