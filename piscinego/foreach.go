package piscine

func ForEach(f func(int), a []int) {
	i := 0
	for i < len(a) {
		f(a[i])
		i++
	}
}
