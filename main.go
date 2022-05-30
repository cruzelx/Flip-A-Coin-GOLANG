package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// Define two sides of the coin
var STATES = []string{"HEAD", "TAIL"}

func main() {

	/*
		Get repeat integer value to toss the coin n number of times.
		If no number is provided 0 is set as default and the flipACoin() function is run once.
	*/
	repeat := flag.Int("R", 0, "Number of times to toss a coin")
	flag.Parse()

	// Generate random seed value for each execution of the application
	rand.Seed(time.Now().UTC().UnixNano())

	// Store all repeated toss results in an array of strings
	results := []string{}

	// Check if repeat value is not the default value 0
	if *repeat != 0 {

		// Append repeated coin toss results
		for i := 0; i < *repeat; i++ {
			results = append(results, flipACoin())
		}

		// Get frequency of heads and tails in the repeated toss
		total, heads, tails := calcFrequency(results)

		// Calculate the percentage of the heads and tails in the tosses
		p_heads := float64(heads) * 100.0 / float64(total)
		p_tails := float64(tails) * 100.0 / float64(total)
		fmt.Printf("Total toss: %v, Heads: (%v, %v%%), Tails: (%v, %v%%)", total, heads, p_heads, tails, p_tails)
		return
	}

	// If repeat is 0 run only once and print the result
	side := flipACoin()
	fmt.Printf("You have got a %v", side)
	return
}

// Randomly choose the coin state
func randChoice(choiceArr []string) string {
	length := len(choiceArr)
	randIndex := rand.Intn(length)
	return choiceArr[randIndex]
}

// Flip for head or tail
func flipACoin() string {
	side := randChoice(STATES)
	return side
}

// Convert tosses result array to map to calculate the frequency of heads and tails
func calcFrequency(tossArr []string) (int, int, int) {
	freq := make(map[string]int)
	for _, num := range tossArr {
		freq[num] = freq[num] + 1
	}

	heads := freq[STATES[0]]
	tails := freq[STATES[1]]

	total := heads + tails

	return total, heads, tails
}
