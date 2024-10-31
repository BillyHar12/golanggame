package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Player struct {
	name string
}

func NewPlayer(name string) *Player {
	return &Player{name: name}
}

func (p *Player) MakeGuess() (int, error) {
	var input string
	fmt.Printf("%s, enter your guess: ", p.name)
	fmt.Scanln(&input)

	guess, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return 0, fmt.Errorf("invalid input, please enter a number")
	}
	return guess, nil
}
