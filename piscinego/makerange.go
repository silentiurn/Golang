package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	result := make([]int, max-min)
	for index := range result {
		result[index] = index + min
	}

	return result
}
