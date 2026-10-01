package services

func Rules(currentPixels []bool, width int, height int, tmpBoard []bool) {

	for i, pixel := range currentPixels {

		tmpBoard[i] = pixelRules(pixel, currentPixels, i%width, i/width, width, height)
	}
}

func pixelRules(pixel bool, currentPixels []bool, x int, y int, width int, height int) bool {
	if pixel == false {
		if getAliveNeighbors(currentPixels, x, y, width, height) == 3 {
			return true
		}
	} else {
		aliveNeighbors := getAliveNeighbors(currentPixels, x, y, width, height)
		if aliveNeighbors == 2 || aliveNeighbors == 3 {
			return true
		}
	}
	return false
}

func getAliveNeighbors(currentPixels []bool, x int, y int, width int, height int) int {
	//number of neighbors
	counter := 0

	if isAlive(currentPixels, y+1, x, width, height) {
		counter++
	}
	if isAlive(currentPixels, y+1, x-1, width, height) {
		counter++
	}
	if isAlive(currentPixels, y+1, x+1, width, height) {
		counter++
	}
	if isAlive(currentPixels, y, x+1, width, height) {
		counter++
	}
	if isAlive(currentPixels, y, x-1, width, height) {
		counter++
	}
	if isAlive(currentPixels, y-1, x, width, height) {
		counter++
	}
	if isAlive(currentPixels, y-1, x+1, width, height) {
		counter++
	}
	if isAlive(currentPixels, y-1, x-1, width, height) {
		counter++
	}
	return counter
}

func isAlive(board []bool, y int, x int, width int, height int) bool {
	// now there is no real border anymore

	// X will continoue on the other side of the board
	if x < 0 {
		x = width - 1
	} else if x >= width {
		x = 0
	}

	// > will continoue on the other side of the board
	if 0 > y {
		y = height - 1
	} else if y >= height {
		y = 0
	}

	index := y*width + x

	return board[index]
}
