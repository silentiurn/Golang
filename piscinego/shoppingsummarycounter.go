package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	slice := []string{}
	var index int
	for i, char := range str {
		if char == ' ' {
			slice = append(slice, str[index:i])
			index = i + 1
		}
	}
	slice = append(slice, str[index:])
	count := make(map[string]int)
	for _, num := range slice {
		count[num]++
	}
	return count
}
