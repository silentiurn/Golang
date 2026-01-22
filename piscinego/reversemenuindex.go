package piscine

func ReverseMenuIndex(menu []string) []string {
	reversed := make([]string, len(menu))
	for i := 0; i < len(menu); i++ {
		reversed[i] = menu[len(menu)-1-i]
	}
	return reversed
}
