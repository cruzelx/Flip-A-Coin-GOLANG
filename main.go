package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var STATES = []string{"HEAD", "TAIL"}

func main() {
	repeat := flag.Int("R", 0, "Number of times to toss a coin")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

	results := []string{}

	if *repeat != 0 {

		for i := 0; i < *repeat; i++ {
			results = append(results, flipACoin())
		}
		total, heads, tails := calcPercentage(results)
		p_heads := float64(heads) * 100.0 / float64(total)
		p_tails := float64(tails) * 100.0 / float64(total)
		fmt.Printf("Total toss: %v, Heads: (%v, %v%%), Tails: (%v, %v%%)", total, heads, p_heads, tails, p_tails)
		return
	}
	side := flipACoin()
	fmt.Printf("You have got a %v", side)
	return
}

// random choice
func randChoice(choiceArr []string) string {
	length := len(choiceArr)
	randIndex := rand.Intn(length)
	return choiceArr[randIndex]
}

func flipACoin() string {

	side := randChoice(STATES)
	return side
}

func calcPercentage(tossArr []string) (int, int, int) {
	freq := make(map[string]int)
	for _, num := range tossArr {
		freq[num] = freq[num] + 1
	}

	heads := freq[STATES[0]]
	tails := freq[STATES[1]]

	total := heads + tails

	return total, heads, tails
}
