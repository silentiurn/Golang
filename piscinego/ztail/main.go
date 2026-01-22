package main

import (
	"os"
)

// Пишет байты в stdout
func writeStdout(b []byte) {
	if len(b) > 0 {
		os.Stdout.Write(b)
	}
}

// Пишет строку в stdout с переносом строки
func writeLineStdout(s string) {
	os.Stdout.Write([]byte(s + "\n"))
}

// Пишет строку в stderr с переносом строки
func writeLineStderr(s string) {
	os.Stderr.Write([]byte(s + "\n"))
}

// Преобразует строку в положительное число
func atoiPositive(s string) (int, bool) {
	if s == "" {
		return 0, false
	}
	n := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		d := int(c - '0')
		if n > (1<<31-1-d)/10 {
			return 0, false
		}
		n = n*10 + d
	}
	if n <= 0 {
		return 0, false
	}
	return n, true
}

// Выводит последние n байт файла
func printLastBytes(name string, n int) bool {
	f, err := os.Open(name)
	if err != nil {
		writeLineStderr(err.Error())
		return false
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		writeLineStderr(err.Error())
		return false
	}

	size := info.Size()
	toRead := int64(n)
	if toRead > size {
		toRead = size
	}
	if toRead <= 0 {
		return true
	}

	start := size - toRead
	buf := make([]byte, toRead)
	readN, _ := f.ReadAt(buf, start)
	if readN > 0 {
		writeStdout(buf[:readN])
	}
	return true
}

// Выводит сообщение об использовании
func usageFail() {
	writeLineStderr("Usage: go run . -c <byte_count> <file1> [file2] ...")
	os.Exit(1)
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		usageFail()
	}

	if args[0] != "-c" {
		usageFail()
	}

	n, ok := atoiPositive(args[1])
	if !ok {
		usageFail()
	}

	files := args[2:]
	if len(files) == 0 {
		usageFail()
	}

	exitStatus := 0
	needBlank := false

	for _, fname := range files {
		// Проверяем, можно ли открыть файл
		f, err := os.Open(fname)
		if err != nil {
			writeLineStderr(err.Error())
			exitStatus = 1
			needBlank = true
			continue
		}
		f.Close()

		// Если уже был успешный файл или ошибка — вставляем пустую строку
		if needBlank {
			writeLineStdout("")
		}
		needBlank = true

		// Заголовок и содержимое
		writeLineStdout("==> " + fname + " <==")
		ok := printLastBytes(fname, n)
		if !ok {
			exitStatus = 1
		}
	}

	if exitStatus != 0 {
		os.Exit(1)
	}
}
