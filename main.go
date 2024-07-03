package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	targetWord  string
	userGuess   string
	gameOver    bool
	gameWon     bool
	guessAmount int

	boardColorsLayout = [6][5]int{
		// 0 = Empty
		// 1 = Gray    / Wrong letter
		// 2 = Yellow / Right letter, wrong place
		// 3 = Green / Right letter, right place
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	boardWords = [6]string{}
)

func main() {
	generateTargetWord()

	if guessAmount <= 6 {
		fmt.Scan(&userGuess)
		checkGuess(userGuess, targetWord)
		writeGuess(userGuess)
	}
}

func generateTargetWord() string {

	rand.New(rand.NewSource(time.Now().UnixNano()))

	file, err := os.Open("src/words.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	randomLineNumber := rand.Intn(len(lines))
	targetWord := lines[randomLineNumber]

	return targetWord
}

func getUserGuess() string {
	fmt.Printf("What do you think the word of the day is?\n")
	for {
		fmt.Scan(&userGuess)
		if len(userGuess) == 5 {
			break
		} else {
			fmt.Println("Please enter a 5 letter word.")
		}
	}
	return userGuess
}

func writeGuess(userGuess string) {
	switch guessAmount {
	case 1:
		boardWords[0] = userGuess
	case 2:
		boardWords[1] = userGuess
	case 3:
		boardWords[2] = userGuess
	case 4:
		boardWords[3] = userGuess
	case 5:
		boardWords[4] = userGuess
	case 6:
		boardWords[5] = userGuess
	}
}

func checkGuess(userGuess, targetWord string) {
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
	if userGuess == targetWord {
		gameOver = true
		gameWon = true
	}
}
