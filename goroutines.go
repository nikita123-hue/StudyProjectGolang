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
	iterate("Обычный вызов:")
	go iterate("Вызов с go:")

	time.Sleep(13 * time.Millisecond)
}
