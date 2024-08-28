package main

import "fmt"

//В го принято сообщать об ошибках через отдельное возвращаемое значение.
//panic это встроенная функция, которая останавливает поток выполнения программы и запускает механизм паники
//recovery - это механизм восстановления дляобработки паники. Используется, когда паника не должна привести к завершению работы всей программы
//1)Используется только внутри defer-функций
//2)Работает только в той горутине, где была вызвана паника

func DivideByNum(a, b float64) float64 {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovery: ", err)
		}
	}()

	if b == 0 {
		panic("Division by zero")
	}
	return a / b

}

func main() {
	fmt.Println(DivideByNum(10, 0))
	fmt.Println(DivideByNum(10, 5))
}
