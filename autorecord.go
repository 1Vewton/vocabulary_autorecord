package main

import (
	"fmt"
	"os"

	"github.com/1Vewton/vocabulary_autorecord/data_management/basic_config"
	"github.com/1Vewton/vocabulary_autorecord/data_management/config"
	"github.com/1Vewton/vocabulary_autorecord/data_management/vocabulary_manager"
	"github.com/1Vewton/vocabulary_autorecord/utils/confirmation_interface"
	"github.com/1Vewton/vocabulary_autorecord/utils/welcome_text"
)

// Initialize: Initializes the env file
func init() {
INIT:
	// Process the error when initiating
	errorProcessInitiating := func(service string, Error error) bool {
		if Error != nil {
			fmt.Printf("Error loading %s due to %s\n", service, Error)
			will_quite := confirmation_interface.ConfirmationInterface(
				"The program cannot start up. Do you want to quit?",
				true,
			)
			if will_quite {
				os.Exit(1)
			} else {
				return false
			}
		} else {
			fmt.Printf("Successfully loaded %s \n", service)
		}
		return true
	}
	Error := config.InitializeConfig(".env")
	if !errorProcessInitiating("env file", Error) {
		goto INIT
	}
	config.InitializeSettings()
	Error = basic_config.InitializeBasicConfig()
	if !errorProcessInitiating("basic config", Error) {
		goto INIT
	}
	Error = vocabulary_manager.InitializeVocabularyManager()
	if !errorProcessInitiating("vocabulary list", Error) {
		goto INIT
	}
}

// main function
func main() {
	fmt.Println("Welcome to ")
	welcome_text.WelcomeText()
}
