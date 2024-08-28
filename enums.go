package main

import "fmt"

//В голанге нет такого понятия как enum. Вместо него используется набор констант
//Ключевое слово  iota позволяет задать индекс константам в перечислении

const (
	Mn = iota
	Ts
	Wd
	Tu
	Fr
	St
	Su
)

const (
	apple = "apple"
	dog   = "dog"
)

type Fruit string

const (
	orange Fruit = "orange"
	banana Fruit = "banana"
)

func main() {
	//При выводе такого мы получим лишь индексы констант
	fmt.Println(Mn, Ts, Wd, Tu, Fr, St, Su)
	//Из недостатков: 1) нет проверки типов во время компиляции. 2) Невозможно добавить новые значения в перечисление после компиляции

	//Альтернативные решения
	//1) Использование const со string
	fmt.Println(apple)

	//2) Использование type со string
	fmt.Println(orange)
}
