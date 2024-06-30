package main

import (
	"fmt"
	"strings"
)

var guessAmount int
var gameOver bool = false
var guessCorrect bool = false
var boardColorsLayout = [6][3]int{
	{0, 0, 0},
	{0, 0, 0},
	{0, 0, 0},
	{0, 0, 0},
	{0, 0, 0},
	{0, 0, 0},
}

func main() {
	var userGuess string
	targetWord := "tar"
	fmt.Printf("The target word is: %s\n", targetWord)

	for i := 0; i < 6; i++ {
		fmt.Scan(&userGuess)
		checkGuess(userGuess, targetWord)
		guessAmount++
		if gameOver && guessCorrect {
			fmt.Printf("You win! The word was %s\n", targetWord)
		} else if gameOver {
			fmt.Printf("You lost! The word was %s\n", targetWord)
		}
	}
}

func checkGuess(userGuess, targetWord string) {
	for i := range userGuess {
		if i < len(userGuess) {
			if userGuess[i] == targetWord[i] {
				boardColorsLayout[guessAmount][i] = 3
				fmt.Println(boardColorsLayout)
			} else if strings.Contains(targetWord, string(userGuess[i])) { // Convert to string to be able to compare
				boardColorsLayout[guessAmount][i] = 2
				fmt.Println(boardColorsLayout)
			} else {
				boardColorsLayout[guessAmount][i] = 1
				fmt.Println(boardColorsLayout)
			}
		}
	}
	if userGuess == targetWord {
		gameOver = true
		guessCorrect = true
	}
}
