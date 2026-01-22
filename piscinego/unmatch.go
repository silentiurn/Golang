package piscine

func Unmatch(a []int) int {
	nums := make(map[int]int, len(a))
	for _, n := range a {
		nums[n]++
	}
	for _, n := range a {
		if nums[n]%2 == 1 {
			return n
		}
	}
	return -1
}
