package main

import (
	"fmt"
	"strings"
)

func main() {
	var targetWord string
	var userGuess string

	targetWord = "yes"
	fmt.Scan(&userGuess)

	strings.Contains(targetWord, userGuess)
}
