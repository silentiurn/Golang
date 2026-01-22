package piscine

func AtoiBase(s string, base string) int {
	if !IsValidBase(base) {
		return 0
	}

	baseLen := len(base)
	result := 0

	for _, char := range s {
		index := indexOf(char, base)
		if index == -1 {
			return 0
		}
		result = result*baseLen + index
	}

	return result
}

func IsValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' || seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

func indexOf(char rune, base string) int {
	for i, b := range base {
		if b == char {
			return i
		}
	}
	return -1
}
