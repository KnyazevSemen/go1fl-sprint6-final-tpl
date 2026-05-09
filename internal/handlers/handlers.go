package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

var tmpl *template.Template

func init() {
	var err error
	tmpl, err = template.ParseFiles("static/index.html")
	if err != nil {
		log.Fatal("ошибка загрузки формы:", err)
	}
}

func CurrentHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	const maxUploadSize = 10 << 20
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

	err := r.ParseMultipartForm(maxUploadSize)
	if err != nil {
		http.Error(w, "ошибка парсинга формы: "+err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "ошибка получения файла: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "ошибка чтения файла: "+err.Error(), http.StatusInternalServerError)
		return
	}

	inputText := string(data)
	convertedText, err := service.ConvertTextOrMorse(inputText)
	if err != nil {
		http.Error(w, "ошибка конвертации: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fileName := generateFileName(header.Filename)
	err = os.WriteFile(fileName, []byte(convertedText), 0644)
	if err != nil {
		http.Error(w, "ошибка записи файла: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(convertedText))
}

func generateFileName(originalName string) string {
	ext := filepath.Ext(originalName)
	baseName := originalName[:len(originalName)-len(ext)]
	timestamp := time.Now().UTC().String()
	return baseName + "_" + timestamp + ext
}
