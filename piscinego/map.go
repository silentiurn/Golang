package piscine

func Map(f func(int) bool, a []int) []bool {
	output := make([]bool, len(a))
	i := 0
	for i < len(a) {
		output[i] = f(a[i])
		i++
	}
	return output
}
