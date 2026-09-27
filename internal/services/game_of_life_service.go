package services

func Rules(currentPixels [][]bool, width int, height int) [][]bool {
	//Create new Board
	newBoard := make([][]bool, width)
	for x := range newBoard {
		newBoard[x] = make([]bool, height)
	}
	for x, pixelX := range currentPixels {
		for y, pixelY := range pixelX {
			newBoard[x][y] = pixelRules(pixelY, currentPixels, y, x)
		}
	}
	return newBoard
}

func pixelRules(pixel bool, currentPixels [][]bool, y int, x int) bool {
	if pixel == false {
		if getAliveNeighbors(currentPixels, y, x) == 3 {
			return true
		}
	}
	if pixel == true {

		aliveNeighbors := getAliveNeighbors(currentPixels, y, x)
		if aliveNeighbors == 2 || aliveNeighbors == 3 {
			return true
		}
	}
	return false
}

func getAliveNeighbors(currentPixels [][]bool, y int, x int) int {
	//number of neighbors
	counter := 0

	if isAlive(currentPixels, x+1, y) {
		counter++
	}
	if isAlive(currentPixels, x+1, y-1) {
		counter++
	}
	if isAlive(currentPixels, x+1, y+1) {
		counter++
	}
	if isAlive(currentPixels, x, y+1) {
		counter++
	}
	if isAlive(currentPixels, x, y-1) {
		counter++
	}
	if isAlive(currentPixels, x-1, y) {
		counter++
	}
	if isAlive(currentPixels, x-1, y+1) {
		counter++
	}
	if isAlive(currentPixels, x-1, y-1) {
		counter++
	}
	return counter
}

func isAlive(board [][]bool, x int, y int) bool {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return false
	}
	return board[x][y]
}
