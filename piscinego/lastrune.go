package piscine

func LastRune(s string) rune {
	r_s := []rune(s)
	return r_s[len(r_s)-1]
}
