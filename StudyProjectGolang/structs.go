package main

import "fmt"

//Структры в го - коллекция полей. Так как в го нет классов, структуры их как бы "заменяют"

type Person struct {
	name string
	age  int
}

func newPerson(name string, age int, p *Person) {
	p.name = name
	p.age = age
}

func (p Person) printName() {
	fmt.Println("Name:", p.name)
}

func (p Person) printAge() {
	fmt.Println("Age:", p.age)
}

func main() {
	var p Person
	newPerson("Nikita", 18, &p)
	p.printName()
	p.printAge()
}
