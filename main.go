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
	guessAmount int

	boardLayout = [5]string{
		{"guess1"},
		{"guess2"},
		{"guess3"},
		{"guess4"},
		{"guess5"},
	}
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

func checkGuess(string) {
	if userGuess != targetWord {
		guessAmount++
	} else {
		gameOver = true
	}
}

func isGameOver() {
	if guessAmount >= 5 {
		gameOver = true
	}
}

func main() {
	generateTargetWord()
	fmt.Scan(userGuess)
	checkGuess(userGuess)

}
