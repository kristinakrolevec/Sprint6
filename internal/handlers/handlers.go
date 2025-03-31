package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// Для корневого эндпоинта / нужно реализовать хендлер, который возвращает HTML из файла index.html.
func HtmlHandle(resp http.ResponseWriter, req *http.Request) {
	root, err := os.OpenRoot("../internal")
	if err != nil {
		http.Error(resp, "внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	defer root.Close()

	data, err := os.ReadFile("../index.html")
	if err != nil {
		http.Error(resp, "ошибка при загрузке страницы", http.StatusInternalServerError)
		return
	}
	resp.Header().Set("Content-Type", "text/html; charset=utf-8")
	resp.WriteHeader(http.StatusOK)
	resp.Write(data)
}

// Второй хендлер для эндпоинта /upload должен выполнять следующие действия:
func ParsHandle(resp http.ResponseWriter, req *http.Request) {
	// 1.Парсить html-форму из файла index.html.

	req.ParseMultipartForm(10 << 20)
	// 2.Получить файл из формы (не забудьте его закрыть).

	file, handler, err := req.FormFile("myFile")
	if err != nil {
		http.Error(resp, "ошибка при получении файла", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	// 3.Прочитать данные из файла.
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(resp, "внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}
	// 4.Передать эти данные в функцию автоопределения из пакета service,
	// которую вы создали, чтобы получить переконвертируемую строку.
	converted := service.ConvMorseString(string(data))

	// 5.Создать локальный файл. Эта операция обычно небезопасна и так делать не рекомендуется,
	// но в рамках нашего задания хотелось бы более наглядного результата, поэтому мы решились на этот шаг,
	// ради видимого результата. А вообще, обычно используют временные файлы.
	ext := filepath.Ext(handler.Filename)
	filename := strings.ReplaceAll(time.Now().UTC().String(), " ", "_")
	filename = strings.ReplaceAll(filename, ":", "-") + ext

	localFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		http.Error(resp, "внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}
	defer localFile.Close()
	// 6.Записать в локальный файл результат конвертации строки.
	// Для генерации имени файла вы можете использовать время с помощью time.Now().UTC().String().
	// Чтобы получить расширения файла, используйте filepath.Ext().

	if _, err = localFile.Write([]byte(converted)); err != nil {
		http.Error(resp, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	resp.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// 7.Вернуть результат конвертации строки.

	if _, err := resp.Write([]byte(converted)); err != nil {
		log.Printf("Error sending response: %v", err)
	}

}
