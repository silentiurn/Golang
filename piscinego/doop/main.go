package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	arg1 := os.Args[1]
	operator := os.Args[2]
	arg3 := os.Args[3]

	// Проверяем валидность оператора
	validOperators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
		"%": true,
	}
	if !validOperators[operator] {
		return
	}

	// Преобразуем аргументы в числа
	a, ok1 := safeAtoi(arg1)
	b, ok2 := safeAtoi(arg3)

	if !ok1 || !ok2 {
		return
	}

	// Выполняем операцию
	var result int
	var overflow bool

	switch operator {
	case "+":
		result, overflow = addWithOverflow(a, b)
	case "-":
		result, overflow = subWithOverflow(a, b)
	case "*":
		result, overflow = mulWithOverflow(a, b)
	case "/":
		if b == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		result = a / b
	case "%":
		if b == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		result = a % b
	}

	if overflow {
		return
	}

	// Выводим результат
	output := itoa(result)
	os.Stdout.WriteString(output + "\n")
}

// Безопасное преобразование строки в int
func safeAtoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	result := 0
	sign := 1
	start := 0

	// Обрабатываем знак
	if s[0] == '+' {
		start = 1
	} else if s[0] == '-' {
		sign = -1
		start = 1
	}

	// Преобразуем каждый символ
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		digit := int(s[i] - '0')

		// Проверка на переполнение
		if result > (1<<63-1-digit)/10 {
			return 0, false
		}
		result = result*10 + digit
	}

	return result * sign, true
}

// Преобразование int в строку
func itoa(n int) string {
	if n == 0 {
		return "0"
	}

	negative := n < 0
	if negative {
		n = -n
	}

	var digits []byte
	for n > 0 {
		digits = append([]byte{byte('0' + n%10)}, digits...)
		n /= 10
	}

	if negative {
		digits = append([]byte{'-'}, digits...)
	}

	return string(digits)
}

// Сложение с проверкой переполнения
func addWithOverflow(a, b int) (int, bool) {
	if b > 0 && a > (1<<63-1)-b {
		return 0, true
	}
	if b < 0 && a < (-1<<63)-b {
		return 0, true
	}
	return a + b, false
}

// Вычитание с проверкой переполнения
func subWithOverflow(a, b int) (int, bool) {
	if b < 0 && a > (1<<63-1)+b {
		return 0, true
	}
	if b > 0 && a < (-1<<63)+b {
		return 0, true
	}
	return a - b, false
}

// Умножение с проверкой переполнения
func mulWithOverflow(a, b int) (int, bool) {
	if a == 0 || b == 0 {
		return 0, false
	}

	result := a * b
	if a == -1<<63 && b == -1 {
		return 0, true
	}
	if b != 0 && result/b != a {
		return 0, true
	}
	return result, false
}
