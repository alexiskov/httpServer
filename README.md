#Простая заготовка сервера стандартного пакета net/http
#Определены часто используемые HTTP-методы

пример запуска:

package main

import (
	"log"
	httpserver "packagetes/httpServer"
)

func main() {
	pipe := httpserver.NewPipe()  //определяем пайп содержащий канал
	go httpserver.New(pipe).Run() //инициализируем сервер с созданным пайпом, запускаем сервер
	//обрабатываем ошибки переданные сервером по каналу пайпа
	for {
		log.Println(<-*pipe.ErorLog)
	}
}