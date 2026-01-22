package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb < 0 || nb > 20 {
		return 0
	}
	for i := 1; i < nb+1; i++ {
		result *= i
		if result < 0 {
			return 0
		}
	}
	return result
}
