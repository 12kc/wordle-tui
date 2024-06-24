package main

import (
	"fmt"
)

func main() {
	/*	var word string
		fmt.Scan(&word)
		charSlice := make([]string, len(word))
		for i, char := range word {
			charSlice[i] = string(char)
		}
		fmt.Println(charSlice) */

	var i int = 0
	for i = 0; i != 10000000; {
		i++
		fmt.Println(i)
	}
}
