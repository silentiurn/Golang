package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	ans := ""
	count := 0
	i := 0
	for i < len(str) {
		if str[i] != ' ' {
			ans += string(str[i])
			count++
		}
		if count == 5 {
			count = 0
			ans += " "
			i++
		}
		i++
	}
	if len(ans) > 0 && ans[len(ans)-1] == ' ' {
		ans = ans[:len(ans)-1]
	}
	return ans + "\n"
}
