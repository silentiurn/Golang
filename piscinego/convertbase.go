package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Сначала переводим строку nbr из baseFrom в десятичное число
	decimal := atoiBase(nbr, baseFrom)

	// Затем переводим десятичное число в baseTo
	result := itoaBase(decimal, baseTo)

	return result
}

func atoiBase(nbr string, base string) int {
	baseLen := len(base)
	result := 0

	for _, c := range nbr {
		index := 0
		for i, b := range base {
			if b == c {
				index = i
				break
			}
		}
		result = result*baseLen + index
	}
	return result
}

func itoaBase(n int, base string) string {
	if n == 0 {
		return string(base[0])
	}

	baseLen := len(base)
	result := ""

	for n > 0 {
		result = string(base[n%baseLen]) + result
		n /= baseLen
	}
	return result
}
