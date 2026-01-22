package piscine

func ActiveBits(n int) int {
	count := 0
	for n > 0 {
		lastBit := n % 2
		count += lastBit
		n /= 2
	}
	return count
}
