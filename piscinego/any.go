package piscine

func Any(f func(string) bool, a []string) bool {
	i := 0
	for i < len(a) {
		if f(a[i]) {
			return true
		}
		i++
	}
	return false
}
