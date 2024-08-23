package main

import (
	"fmt"
)

func main() {
	//Map - это ассоциативный тип данных или "хэш"(в некоторых языках его так называют). Предаставляет собой коллекицию ключ-значение

	//Инициализация map
	testMap := make(map[string]int)

	//Инициализация с объявлением
	testMap2 := map[string]int{"value1": 1, "value2": 2}
	fmt.Println(testMap2)

	testMap["val1"] = 1
	testMap["val2"] = 2
	testMap["val3"] = 3

	fmt.Println(testMap)

	//С помощью метода delete можно удалять какие-то элементы из карты
	delete(testMap, "val2")
	fmt.Println(testMap)

	//Clear очищает карту
	clear(testMap)
	fmt.Println(testMap)
}
