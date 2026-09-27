package handlers

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func PrintMenu(pixels [][]bool) {
	clearScreen()
	for _, pixelX := range pixels {
		for _, pixelY := range pixelX {
			if pixelY == true {
				fmt.Print("██")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
