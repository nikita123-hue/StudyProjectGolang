package main

import "fmt"

// Ну, методы они и в Африке методы, сказать нечего
func main() {
	a := newAnimal("Tiger", 120)
	fmt.Println("Characteristics of animal: Name -", a.name, ", weight -", a.weight)
	a.sound()

}

type Animal struct {
	name   string
	weight int
}

func (a Animal) sound() {
	fmt.Println("Animal sound")
}

func newAnimal(name string, weight int) *Animal {
	a := Animal{name: name, weight: weight}
	return &a
}
