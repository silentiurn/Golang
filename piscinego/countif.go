package piscine

func CountIf(f func(string) bool, tab []string) int {
	counter := 0
	i := 0
	for i < len(tab) {
		if f(tab[i]) {
			counter++
		}
		i++
	}
	return counter
}
