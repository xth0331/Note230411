package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	heads int
	tails int
)

func coinflip() string {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(2)

	if result == 0 {
		return "heads"
	} else {
		return "tails"
	}
}

func main() {

	switch coinflip() {
	case "heads":
		heads++
	case "tails":
		tails++
	default:
		fmt.Println("landed on edge!")
	}

	fmt.Printf("Heads: %d, Tails: %d\n", heads, tails)
}
