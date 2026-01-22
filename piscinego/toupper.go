package piscine

func ToUpper(s string) string {
	r_s := []rune(s)
	for i, char := range r_s {
		if char >= 'a' && char <= 'z' {
			r_s[i] = char - 32
		}
	}
	return string(r_s)
}
