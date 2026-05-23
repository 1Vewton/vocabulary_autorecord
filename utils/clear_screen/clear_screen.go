package clear_screen

import (
	"fmt"
)

// ClearScreen clears the screen based on the operating system.
func ClearScreen() {
	fmt.Print("\033[2J\033[H")
}
