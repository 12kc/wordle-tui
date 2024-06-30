package main

import (
	"fmt"
	"strings"
)

func main() {
	var userGuess string
	var boardColorsLayout = [6][3]int {
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	var guessAmount int = 0
	targetWord := "tes"

	for i := 0; i < 6; i++ {
		fmt.Scan(&userGuess)
		guessAmount++
	}

	for i := range userGuess {
		if i < len(userGuess) {
			if userGuess[i] == targetWord[i] {
				boardColorsLayout[guessAmount][i] = 3
			} else if strings.Contains(string(userGuess[i]), targetWord) { // Convert to string to be able to compare
				boardColorsLayout[guessAmount][i] = 2
			} else {
				boardColorsLayout[guessAmount][i] = 1
			}
		}
	}	

	for i := 0; i < 6; i++ {
		fmt.Println(boardColorsLayout)
	}
}
