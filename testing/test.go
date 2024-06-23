package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	// Seed the random number generator (fixed line)
	rand.New(rand.NewSource(time.Now().UnixNano()))

	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	randomLineNumber := rand.Intn(len(lines))
	randomWord := lines[randomLineNumber]

	fmt.Println("Random word from file:", randomWord)
}
