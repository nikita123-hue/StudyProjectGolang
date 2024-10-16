package main

import (
	"fmt"
	"time"
)

func iterate(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str, i)
	}
}

func main() {
	//Создание каналов
	//bufChannel := make(chan int, 1)
	//unbufChannel := make(chan int)

	//Однонаправленные каналы
	//readingChannel := make(<-chan int) //канал для чтения
	//writingChannel := make(chan<- int) //канал для записи
	iterate("Обычный вызов:")
	go iterate("Вызов с go:")

	time.Sleep(13 * time.Millisecond)
}
