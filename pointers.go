package main

import "fmt"

// Есть передача по значению и по указателю. При передаче по значению функция получает именно копию аргумента и лбые операции внутри функции не меняют переменную
// А при передачеч по указателю функция получает именно указатель на память, выделенную под переменную. Так что люое изменение в функции будет менять переменную
func main() {
	i := 1
	fmt.Println("Start", i)

	byValue(i)
	fmt.Println("Step 1", i)

	byPointers(&i)
	fmt.Println("Step 2", i)
}

func byValue(test1 int) {
	test1 = 0
}

func byPointers(test1 *int) {
	*test1 = 0
}
