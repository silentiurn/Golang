package piscine

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	}

	nb = nb * RecursivePower(nb, power-1)
	return nb
}
