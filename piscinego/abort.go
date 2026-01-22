package piscine

func Abort(a, b, c, d, e int) int {
	ar := []int{a, b, c, d, e}
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}
	return ar[2]
}
