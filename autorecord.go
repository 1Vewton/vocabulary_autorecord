package main

import (
	"fmt"
	"os"

	"github.com/1Vewton/vocabulary_autorecord/commands/root"
	"github.com/1Vewton/vocabulary_autorecord/data_management/basic_config"
	"github.com/1Vewton/vocabulary_autorecord/data_management/config"
	"github.com/1Vewton/vocabulary_autorecord/data_management/vocabulary_manager"
	"github.com/1Vewton/vocabulary_autorecord/utils/confirmation_interface"
	"github.com/1Vewton/vocabulary_autorecord/utils/welcome_text"
)

// Initialize: Initializes the configuration for the program
func init() {
INIT:
	// Process the error when initiating
	errorProcessInitiating := func(service string, Error error) bool {
		if Error != nil {
			fmt.Printf("Error loading %s due to %s\n", service, Error)
			will_quite := confirmation_interface.ConfirmationInterface(
				"\033[31mThe program cannot startup, do you want to quit?\033[0m",
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
	fmt.Println("\033[32mInitializing the program...\033[0m")
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
	Err := root.Execute()
	if Err != nil {
		fmt.Printf("\033[31mExecution failed due to %s\033[0m", Err)
	}
}
