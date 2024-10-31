package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("You have to guess a number between 1 and 100.")

	player := NewPlayer("Player1")
	game := NewGame(player)

	if err := game.Start(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Thanks for playing!")
}
