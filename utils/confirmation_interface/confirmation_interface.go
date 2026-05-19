package confirmation_interface

import (
	"fmt"
)

// Confirmation interface: [y/N] or [Y/n]
func ConfirmationInterface(prompt string, defaultAns bool) bool {
	var suffix string
	if defaultAns {
		suffix = " [Y/n]"
	} else {
		suffix = " [y/N]"
	}
	fmt.Printf("%s %s: ", prompt, suffix)
	fmt.Println()
	// Input
	var input string
	fmt.Scan(&input)
	// Check input
	if input == "y" || input == "Y" {
		return true
	} else if input == "n" || input == "N" {
		return false
	} else {
		return defaultAns
	}
}
