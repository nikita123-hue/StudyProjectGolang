package main

import (
	"log"
	"os"
)

//Логирование - процесс записи информации о всех событиях, которые происходят в программе
//Для логирование в го используется пакет log

func main() {

	file, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Panic("Ошибка открытия файла логов:", err)
	}
	defer file.Close() // отложенное закрытие файла

	log.SetOutput(file) // перенаправление вывода логов в файл logfile.log
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)

	log.Println("New message!")
}
