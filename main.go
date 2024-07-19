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
	gameOver          bool
	gameWon           bool
	guessAmount       int
	boardColorsLayout = [6][5]int{
		// 0 = Empty
		// 1 = Gray
		// 2 = Yellow
		// 3 = Green
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
	var targetWord string = generateTargetWord()

	for ; guessAmount < 6; guessAmount++ {
		userGuess := getUserGuess()
		checkGuess(userGuess, targetWord)
		writeGuess(userGuess)
		if guessAmount == 6 {
			gameOver = true
		}
		if gameWon && gameOver {
			fmt.Printf("Congratulations! The word was %s", targetWord)
		} else if gameOver {
			fmt.Printf("You lost! The word was %s", targetWord)
		}
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
	var userGuess string
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
			if string(userGuess[i]) == string(targetWord[i]) {
				boardColorsLayout[guessAmount][i] = 3
				fmt.Println(boardColorsLayout[guessAmount])
			} else if strings.Contains(targetWord, string(userGuess[i])) { // Convert to string to be able to compare
				boardColorsLayout[guessAmount][i] = 2
				fmt.Println(boardColorsLayout[guessAmount])
			} else {
				boardColorsLayout[guessAmount][i] = 1
				fmt.Println(boardColorsLayout[guessAmount])
			}
		}
	}
	if userGuess == targetWord {
		gameOver = true
		gameWon = true
	}
}
