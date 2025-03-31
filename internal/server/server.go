package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// LogAndServ - 1. Создайте структуру сервера с полями для логгера (log.Logger) и http-сервера (http.Server).
type LogAndServ struct {
	Log  *log.Logger
	Serv *http.Server
}

//	MyServer - 2. Создайте функцию, в которой нужно создать http-роутер. Функция принимает log.Logger и возвращает
//
// экземпляр структуры вашего сервера.
func MyServer(l *log.Logger) *LogAndServ {

	// 3. Зарегистрируйте ваши хендлеры в http-роутере.
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HtmlHandle)
	mux.HandleFunc("/upload", handlers.ParsHandle)

	// 4. Создайте экземпляр структуры http.Server. Для настойки вашего сервера используйте следующие поля:
	// Addr — используйте порт 8080.
	// Handler — передайте ваш http-роутер.
	// ErrorLog — передайте ваш логгер.
	// ReadTimeout — таймаут для чтения. 5 секунд.
	// WriteTimeout — таймаут для записи. 10 секунд.
	// IdleTimeout — таймаут ожидания следующего запроса. 15 секунд.
	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	//5. Верните ссылку на ваш сервер.
	return &LogAndServ{
		Log:  l,
		Serv: s,
	}
}
