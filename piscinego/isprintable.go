package piscine

func IsPrintable(s string) bool {
	for _, char := range s {
		// Проверяем, что символ в диапазоне от 32 до 126 (включая пробелы)
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}
