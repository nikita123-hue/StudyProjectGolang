package main

import (
	"fmt"
	"math"
)

//Интерфейсы - именнованые коллекции сигнатур методов
//Допустим у нас есть две структуры - круг и треугольник

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

type geometry interface {
	area() float64
	perim() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	//Интерфейсы испоьзуются для предоеления набора методов, которые должны быть реализованы типами данных. Это позволяет
	// в последствии использовать одни и теже функции для различных объектов
	r := rect{width: 5, height: 4}
	c := circle{radius: 10}

	fmt.Println("Rectangle:")
	measure(r)

	fmt.Println("Circle:")
	measure(c)

}
