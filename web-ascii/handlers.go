package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// Структура для передачи данных в шаблон ошибки
type ErrorData struct {
	Code    int
	Message string
}

// Универсальная функция для отображения ошибок
func renderError(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		// Если файл шаблона ошибки не найден, отдаем текст
		http.Error(w, fmt.Sprintf("Error %d: %s", status, msg), status)
		return
	}
	tmpl.Execute(w, ErrorData{Code: status, Message: msg})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderError(w, http.StatusNotFound, "Page Not Found")
		return
	}
	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	tmpl.Execute(w, nil)
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || (banner != "standard" && banner != "shadow" && banner != "thinkertoy") {
		renderError(w, http.StatusBadRequest, "Bad Request: Missing text or banner")
		return
	}

	for _, ch := range text {
		if (ch < 32 || ch > 126) && ch != '\n' && ch != '\r' {
			renderError(w, http.StatusBadRequest, "Bad Request: Non-ASCII characters")
			return
		}
	}

	result, err := GenerateASCII(text, banner)
	if err != nil {
		if os.IsNotExist(err) {
			renderError(w, http.StatusNotFound, "Banner file not found")
		} else {
			renderError(w, http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	tmpl.Execute(w, result)
}
