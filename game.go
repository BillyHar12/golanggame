package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	player   *Player
	number   int
	attempts int
}

func NewGame(player *Player) *Game {
	rand.Seed(time.Now().UnixNano())
	return &Game{
		player:   player,
		number:   rand.Intn(100) + 1, // Random number between 1 and 100
		attempts: 0,
	}
}

func (g *Game) Start() error {
	for {
		guess, err := g.player.MakeGuess()
		if err != nil {
			return err
		}
		g.attempts++

		if guess < g.number {
			fmt.Println("Too low! Try again.")
		} else if guess > g.number {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congratulations! You've guessed the number %d in %d attempts.\n", g.number, g.attempts)
			break
		}
	}
	return nil
}
