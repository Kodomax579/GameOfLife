package main

import (
	"gameOfLife/internal/handlers"
	"gameOfLife/internal/services"
	"math/rand"
	"time"
)

func main() {
	width := 40
	height := 60
	Board := SeedRandom(0.3, width, height)

	handlers.PrintMenu(Board)
	for {
		Board = services.Rules(Board, width, height)
		handlers.PrintMenu(Board)
		time.Sleep(100 * time.Millisecond)
	}
}

func SeedRandom(density float64, width int, height int) [][]bool {
	board := make([][]bool, width)
	for y := range board {
		board[y] = make([]bool, height)
	}

	//Generate random seed
	for y := range board {
		for x := range board[y] {
			if rand.Float64() < density {
				board[y][x] = true
			}
		}
	}
	return board
}
