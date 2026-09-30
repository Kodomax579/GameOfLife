package handlers

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func PrintMenu(pixels []bool, width int) {
	clearScreen()

	var displayBoard strings.Builder

	for i, pixel := range pixels {

		if i%width == 0 {
			displayBoard.WriteString("\n")
		}
		if pixel == true {
			displayBoard.WriteString("██")
		} else {
			displayBoard.WriteString("  ")
		}
	}
	fmt.Print(displayBoard.String())
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
