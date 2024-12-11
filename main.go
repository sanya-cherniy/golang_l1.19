package main

import "fmt"

func main() {
	var input string

	fmt.Scan(&input)

	rune_input := []rune(input)
	for i := len(rune_input) - 1; i >= 0; i-- {
		fmt.Print(string(rune_input[i]))
	}

}
