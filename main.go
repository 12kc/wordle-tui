package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// FIX VARIABLE NAMES LATER
// apiRequest, apiData, targetWord, jsonFormat etc
type jsonFormat struct {
	id                int
	solution          string
	days_since_launch int
	editor            string
}

var currentDay = time.Now()

func getDate() {
	fmt.Println("Welcome to Gordle, the Wordle app made in Go!")
	fmt.Printf("Today is %s %d!", currentDay.Month(), currentDay.Day(), "\n")
}

func (targetWord string) getWordOfTheDay() {
	data, err := http.Get(fmt.Sprintf("https://www.nytimes.com/svc/wordle/v2/%d-%d-%d.json", currentDay.Year(), currentDay.Month(), currentDay.Day()))

	if err != nil {
		log.Fatalln("The API was not available at this time, maybe check your internet connection?")
	}

	err = json.Unmarshal(apiRequest, &data)
	targetWord = data.solution
}

func main() {
	getDate()
}
