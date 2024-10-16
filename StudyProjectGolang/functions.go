package main

import "fmt"

func main() {
	//!Важно: в голанге не существует перегрузки функций из-за принципа простоты языка
	//Синткасис функции: func(параметры) тип возвращаемого значения

	a := 3
	b := 5
	c := 1

	fmt.Println(plus(a, b))
	fmt.Println(plusPlus(a, b, c))
	fmt.Println(multiply(a, b, c))

	fmt.Println("Check:", check(1, 2, 3, 4, 5, 6))
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

// Вместо перегрузки функции можно не передавать точное кол-во параметров в функцию просто
func multiply(nums ...int) int {
	result := 1

	for _, num := range nums {
		result *= num
	}
	return result
}

func check(nums ...int) []int {

	var checkSlice []int

	for _, num := range nums {
		if num%2 == 0 {
			checkSlice = append(checkSlice, num)
		}
	}

	return checkSlice
}
