package piscine

func StrRev(s string) string {
	runes := []rune(s)
	reversed := ""
	for i := len(runes) - 1; i >= 0; i-- {
		reversed += string(runes[i])
	}
	return reversed
}
