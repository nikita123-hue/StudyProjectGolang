package main

import "fmt"

// Строки в голанге - массив байтов. В других языках строки состоят из символов, а в го - из рун.
func main() {
	str := "qwerty"

	//здесь каждый символ будет выводится как номер символа по UTF-8
	for _, s := range str {
		fmt.Print(s, " ")
	}

	fmt.Println("\nLength of str:", len(str))

	//Чтобы перебрать нормально массив рун, нужно преобразовывать руну в строку
	for _, v := range str {
		fmt.Print(string(v))
	}

	fmt.Println("\nSmall test: enter a some letter ")
	var str2 string
	fmt.Scan(&str2)

	var q rune
	found := false

	for _, q = range str {
		if string(q) == str2 {
			fmt.Println("Letter", str2, "is found!")
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Letter", str2, "is not found!")
	}
}
