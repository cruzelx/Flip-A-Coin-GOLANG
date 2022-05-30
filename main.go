package main

import (
	"fmt"
	"math/rand"
	"time"
)

var STATES = []string{"HEAD", "TAIL"}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	side := randChoice(STATES)
	fmt.Printf("You've got a %v \n", side)
}

// random choice
func randChoice(choiceArr []string) string {
	length := len(choiceArr)
	randIndex := rand.Intn(length)
	return choiceArr[randIndex]
}
