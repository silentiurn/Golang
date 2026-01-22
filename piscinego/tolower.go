package piscine

func ToLower(s string) string {
	r_s := []rune(s)
	for i, char := range r_s {
		if char >= 'A' && char <= 'Z' {
			r_s[i] = char + 32
		}
	}
	return string(r_s)
}
