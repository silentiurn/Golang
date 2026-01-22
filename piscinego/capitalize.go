package piscine

func Capitalize(s string) string {
	result := []byte(s)
	isNewWord := true

	for i := 0; i < len(result); i++ {
		if (result[i] >= 'a' && result[i] <= 'z') || (result[i] >= 'A' && result[i] <= 'Z') || (result[i] >= '0' && result[i] <= '9') {
			if isNewWord && result[i] >= 'a' && result[i] <= 'z' {
				result[i] -= 'a' - 'A' // Делаем букву заглавной
			} else if !isNewWord && result[i] >= 'A' && result[i] <= 'Z' {
				result[i] += 'a' - 'A' // Делаем букву строчной
			}
			isNewWord = false
		} else {
			isNewWord = true
		}
	}

	return string(result)
}
