package main

import "fmt"

// Объявление констант
const testConst int = 0

func main() {
	//Объвление переменных
	var a int
	fmt.Println(a)

	var b int
	b = 3
	fmt.Println(b)

	c := 4
	fmt.Println(c)

	var d, e int = 1, 2
	fmt.Println(d + e)

	fmt.Println(testConst)
}
