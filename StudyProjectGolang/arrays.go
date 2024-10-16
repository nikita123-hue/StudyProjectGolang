package main

import "fmt"

func main() {
	//Инициализация массива
	testArray2 := [...]int{1, 2, 3, 4, 5}
	for _, t := range testArray2 {
		fmt.Print(t)
	}

	//Инифиализация двухмерного массива
	var testArray3 [2][2]int

	testArray3 = [2][2]int{
		{1, 2},
		{2, 1},
	}

	for _, t := range testArray3 {
		fmt.Print(t)
	}
}
