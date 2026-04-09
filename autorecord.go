package main

import (
	"fmt"

	"github.com/1Vewton/vocabulary_autorecord/data_management/config"
)

// Initialize: Initializes the env file
func init() {
	Error := config.InitializeConfig(".env")
	if Error != nil {
		fmt.Println("Error loading.env file due to ", Error)
	} else {
		fmt.Println("Successfully loaded env file")
	}
}

// main function
func main() {
	fmt.Println("Welcome to vocabulary_autorecord")
}
