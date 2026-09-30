package main

import (
	"gameOfLife/internal/handlers"
	"gameOfLife/internal/services"
	"math/rand"
	"time"
)

func main() {
	width := 60
	height := 40

	//Generate two 1d arrays => we don't need to create a new array every time
	// Reduce timecomplexity
	board := SeedRandom(0.3, width, height)
	tmpBoard := make([]bool, len(board))

	handlers.PrintMenu(board, width)
	for {
		services.Rules(board, width, height, tmpBoard)

		// We will switch the references between two array fields
		// board => 1,2,3
		// tmpBoard => 4,5,6
		// When we switch, I can manipulate the tmpBoard board with new data without destroying the normal board
		board, tmpBoard = tmpBoard, board
		handlers.PrintMenu(board, width)
		time.Sleep(100 * time.Millisecond)
	}
}

func SeedRandom(density float64, width int, height int) []bool {

	board := make([]bool, width*height)

	//Generate random seed
	for i := range board {
		if rand.Float64() < density {
			board[i] = true
		}
	}
	return board
}
