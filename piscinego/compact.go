package piscine

func Compact(ptr *[]string) int {
	var new_str []string
	count := 0
	for _, value := range *ptr {
		if value != "" && value != " " {
			new_str = append(new_str, value)
			count++
		}
	}
	*ptr = new_str
	return count
}
