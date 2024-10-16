package main

import "fmt"

//Дженерики - механизм, который позволяет создавать функции и типы данных, способные работать с различными типами данных

func SumOfValues[T int | float64](s []T) T {
	var result T
	result = 0
	for _, v := range s {
		result += v
	}
	return result
}

func main() {
	intValues := []int{1, 2, 3, 4, 5}
	floatValues := []float64{1.01, 1.02, 1.03, 1.04}

	str := fmt.Sprintf("%.1f", SumOfValues(floatValues))

	fmt.Println("Sum of integer:", SumOfValues(intValues))
	fmt.Println("Sum of float:", str)
}
