package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// FIX VARIABLE NAMES LATER
// apiRequest, apiData, targetWord, jsonFormat, number etc
var currentDay = time.Now()

func getDate() {
	fmt.Println("Welcome to Gordle, the Wordle app made in Go!")
	fmt.Printf("Today is %s %d!", currentDay.Month(), currentDay.Day(), "\n")
}

func (targetWord string, downloadOption byte) getTargetWord() {

	// Generate random line number
	var number int
	number = rand.Intn(1000)
	fmt.Println(number) // For testing purposes, remove on commit

	file, err := os.Open("src/words.txt")
	if err != nil {
		log.Fatal(err)
	}

}
func main() {
}
