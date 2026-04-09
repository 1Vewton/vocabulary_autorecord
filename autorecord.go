package main

import (
	"fmt"

	"github.com/1Vewton/vocabulary_autorecord/data_management/config"
	"github.com/1Vewton/vocabulary_autorecord/utils/welcome_text"
)

// Initialize: Initializes the env file
func init() {
	Error := config.InitializeConfig(".env")
	if Error != nil {
		fmt.Println("Error loading.env file due to ", Error)
	} else {
		fmt.Println("Successfully loaded env file")
	}
	config.InitializeSettings()
}

// main function
func main() {
	fmt.Println("Welcome to ")
	welcome_text.WelcomeText()
}
