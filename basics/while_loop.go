package main

import (
	"fmt"
	"math/rand"
	"time"
)

func while_loop() {
	guessingGame()
}

func guessingGame() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(100) + 1

	var guess int
	for {
		fmt.Println("Enter your guess 1 - 100:")
		fmt.Scanln(&guess)

		switch {
		case guess == target:
			fmt.Println("Congratulation! Correct answer is: ", target)
			break
		case guess > target:
			fmt.Println("Lower")
		case guess < target:
			fmt.Println("Higher")

		}
	}
}
