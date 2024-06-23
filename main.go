package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	targetWord  string
	userGuess   string
	gameOver    bool
	gameWon     bool
	guessAmount int = 0

	// 0 = Empty
	// 1 = Gray
	// 2 = Yellow
	// 3 = Green
	boardColorsLayout = [6][5]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	boardWords = [6]string{}
)

func generateTargetWord() string {

	rand.New(rand.NewSource(time.Now().UnixNano()))

	file, err := os.Open("words.txt")
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

func splitGuess() {

}

func checkGuess(userGuess string) {
	if userGuess == targetWord {
		gameOver = true
		gameWon = true
	} else {
		guessAmount++
	}
}

func appendGuess(userGuess string) {
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

func isGameOver() {
	if guessAmount >= 5 {
		gameOver = true
	}
}

func getUserGuess() {
	fmt.Scan(userGuess)
	fmt.Printf("What do you think the word of the day is?\n")
}
func main() {
	generateTargetWord()
	getUserGuess()
	checkGuess(userGuess)
	appendGuess(userGuess)
}
