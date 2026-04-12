package confirmation_interface

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Confirmation interface: [y/N] or [N/y]
func ConfirmationInterface(prompt string, defaultAns bool) bool {
	var suffix string
	if defaultAns {
		suffix = " [Y/n]"
	} else {
		suffix = " [y/N]"
	}
	fmt.Printf("%s %s: ", prompt, suffix)
	// Input
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	// Error: Get default answer
	if err != nil {
		return defaultAns
	}
	// Trim input
	input = strings.TrimSpace(input)
	// Check input
	if input == "y" || input == "Y" {
		return true
	} else if input == "n" || input == "N" {
		return false
	} else {
		return defaultAns
	}
}
