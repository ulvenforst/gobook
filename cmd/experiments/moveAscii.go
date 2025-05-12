package experiments

import (
	"fmt"
	"time"
)

func MoveAscii() {
	// The character to animate
	char := ">"

	// Terminal width (adjust as needed)
	width := 50

	// Hide cursor during animation
	fmt.Print("\033[?25l")

	// Animation loop
	for i := range width {
		// Move cursor to starting position (row 1, column 1)
		fmt.Print("\033[1;1H")

		// Create padding space before the character
		padding := ""
		for range i {
			padding += " "
		}

		// Print the character with padding
		fmt.Println(padding + char)

		// Delay between frames
		time.Sleep(100 * time.Millisecond)
	}

	// Show cursor again
	fmt.Print("\033[?25h")
}
