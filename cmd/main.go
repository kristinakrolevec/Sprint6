package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	//Здесь, в функции main() нужно создать логгер,
	// далее создать сервер с помощью вашей функции из пакета server,
	// и запустить его. Если при запуске сервера возникают ошибки,
	// выведите её с помощью логгера на уровне Fatal.

	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	myLog := log.New(file, "serv", log.LstdFlags|log.Lshortfile)

	srv := server.MyServer(myLog)

	err = srv.Serv.ListenAndServe()
	if err != nil {
		log.Fatal()
	}

}
