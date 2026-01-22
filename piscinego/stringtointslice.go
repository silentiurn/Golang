package piscine

func StringToIntSlice(str string) []int {
	var result []int
	for _, value := range str {
		result = append(result, int(value))
	}
	return result
}
