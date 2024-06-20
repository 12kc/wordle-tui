package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var boardLayout = [6][5]string{
	{"a", "b", "c", "d", "e"},
	{"a", "b", "c", "d", "e"},
	{"a", "b", "c", "d", "e"},
	{"a", "b", "c", "d", "e"},
	{"a", "b", "c", "d", "e"},
	{"a", "b", "c", "d", "e"},
}

var guess0 string
var guess1 string
var guess2 string
var guess3 string
var guess4 string
var guess5 string

func getTargetWord() string {

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

func formatGuess(string) {
	formattedGuess0 := []string(guess0)
	fmt.Println(formattedGuess0)
}

func main() {
	fmt.Scan(&guess0)
	formatGuess(guess0)
}
