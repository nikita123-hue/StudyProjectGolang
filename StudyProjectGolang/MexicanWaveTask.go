package main

import (
	"fmt"
)

func main() {

	var words string = "hello"

	var result []string

	for i, char := range words {
		if char != ' ' {
			waveWord := words[:i] + string(char-32) + words[i+1:]
			result = append(result, waveWord)
		}
	}

	fmt.Print(result)
}
