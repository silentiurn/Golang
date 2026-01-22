package piscine

func NRune(s string, n int) rune {
	r_s := []rune(s)
	for index, value := range r_s {
		if index == n-1 {
			return value
		}
	}
	return 0
}
