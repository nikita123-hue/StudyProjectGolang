package main

import "fmt"

//Встраивание структур - это использование другой структуры как поле в новой структуры(можно использовать также с интефрейсами)
//Нужно для того, чтобы код был более читаем, и чтобы к взаимосвязанным объектам был доступ через один элемент

type person struct {
	name string
	age  int
}
type employee struct {
	person
	salary int
	title  string
}

// Однако при использовании функции создания нового объекта нужно будет обращаться к вложенным функциям отдельно
func newEmployee(name string, age int, salary int, title string) *employee {
	return &employee{
		person: person{
			name: name,
			age:  age,
		},
		salary: salary,
		title:  title,
	}
}

func main() {
	e := newEmployee("Nikita", 18, 1200, "programer")
	fmt.Println(e)
}
