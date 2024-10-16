package main

import (
	"fmt"
)

func main() {
	var test int
	fmt.Scan(&test)

	switch test {
	case 1:
		fmt.Println("1")

	case 2:
		fmt.Println("2")

	case 3:
		fmt.Println("3")

	default:
		fmt.Println("Why?")
	}

}
