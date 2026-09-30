package services

func Rules(currentPixels []bool, width int, height int, tmpBoard []bool) {

	for i, pixel := range currentPixels {

		tmpBoard[i] = pixelRules(pixel, currentPixels, i%width, i/width, width)
	}
}

func pixelRules(pixel bool, currentPixels []bool, x int, y int, width int) bool {
	if pixel == false {
		if getAliveNeighbors(currentPixels, x, y, width) == 3 {
			return true
		}
	} else {
		aliveNeighbors := getAliveNeighbors(currentPixels, x, y, width)
		if aliveNeighbors == 2 || aliveNeighbors == 3 {
			return true
		}
	}
	return false
}

func getAliveNeighbors(currentPixels []bool, x int, y int, width int) int {
	//number of neighbors
	counter := 0

	if isAlive(currentPixels, y+1, x, width) {
		counter++
	}
	if isAlive(currentPixels, y+1, x-1, width) {
		counter++
	}
	if isAlive(currentPixels, y+1, x+1, width) {
		counter++
	}
	if isAlive(currentPixels, y, x+1, width) {
		counter++
	}
	if isAlive(currentPixels, y, x-1, width) {
		counter++
	}
	if isAlive(currentPixels, y-1, x, width) {
		counter++
	}
	if isAlive(currentPixels, y-1, x+1, width) {
		counter++
	}
	if isAlive(currentPixels, y-1, x-1, width) {
		counter++
	}
	return counter
}

func isAlive(board []bool, y int, x int, width int) bool {
	if x < 0 || x >= width || 0 > y || y >= (len(board)/width) {
		return false
	}

	index := y*width + x

	return board[index]
}
