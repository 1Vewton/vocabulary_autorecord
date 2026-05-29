package basic_config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/1Vewton/vocabulary_autorecord/data_management/config"
	"github.com/1Vewton/vocabulary_autorecord/utils/confirmation_interface"
	"github.com/1Vewton/vocabulary_autorecord/utils/json_validator"
)

// basic config struct
type basicConfig struct {
	VocabFieldName      string
	DefinitionFieldName string
	LLMApiKey           string
	LLMProvider         string
	LLMModelName        string
	LLMBaseURL          string
}

// Basic Config
var BasicConfig basicConfig

// Initialize Basic Config to default values
func initializeBasicConfigDefault() {
	BasicConfig.DefinitionFieldName = "definition"
	BasicConfig.VocabFieldName = "vocabulary"
	BasicConfig.LLMApiKey = "YOUR_LLM_API_KEY"
	BasicConfig.LLMProvider = "openai"
	BasicConfig.LLMModelName = "gpt-3.5-turbo"
	BasicConfig.LLMBaseURL = "https://api.openai.com/v1/completions"
}

// Instruction for user to set the basic config
func instruction4SettingBasicConfig() {
	fmt.Println("You refuse to overwrite the configuration file. However, the program cannot run without the basic configuration.")
	fmt.Println("You can set the json file manually. ")
	fmt.Println("The followings are the meaning for the fields: ")
	fmt.Println("- VocabFieldName: the name of the field in the vocabulary list file that contains the raw vocabulary.")
	fmt.Println("- DefinitionFieldName: the name of the field in the vocabulary list file that contains the definition of the vocabulary.")
	fmt.Println("- LLMApiKey: the API key for the LLM provider.")
	fmt.Println("- LLMProvider: the name of the LLM provider.")
	fmt.Println("- LLMModelName: the name of the LLM model.")
	fmt.Println("- LLMBaseURL: the base URL for the LLM provider.")
}

// Initialize Basic Config
func InitializeBasicConfig() (Error error) {
	// Check whether the configuration file exists
	_, err := os.Stat(config.Settings.BaiscConfigPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Create file if not exist
			initializeBasicConfigDefault()
			_, err := os.Create(config.Settings.BaiscConfigPath)
			if err == nil {
				bytes, _ := json.MarshalIndent(BasicConfig, "", "  ")
				err = os.WriteFile(config.Settings.BaiscConfigPath, bytes, 0666)
				if err != nil {
					return err
				}
			}
			return err
		} else {
			return err
		}
	}
	// Read the file
	content, err := os.ReadFile(config.Settings.BaiscConfigPath)
	if err != nil {
		return err
	}
	contentString := string(content)
	// After reading the file
	err = json.Unmarshal(content, &BasicConfig)
	if err != nil {
		// Overwrite the file if the json file does not have correct style of json
		initializeBasicConfigDefault()
		bytes, _ := json.MarshalIndent(BasicConfig, "", "  ")
		err = os.WriteFile(config.Settings.BaiscConfigPath, bytes, 0666)
		if err != nil {
			return err
		}
	} else {
		// Validate the config
		schema_content, err := os.ReadFile(config.Settings.ConfigSchemaPath)
		if err != nil {
			return err
		} else {
			schema_content_string := string(schema_content)
			result, err := json_validator.Validate(schema_content_string, contentString)
			if err != nil {
				return err
			}
			if !result {
				// choice for user
				for {
					choice := confirmation_interface.ConfirmationInterface(
						"\033[31mThe schema of the configuration file is not correspond to the schema, do you want to overwrite the configuration file with default setting?\033[0m",
						false,
					)
					if choice {
						fmt.Println("Overwriting the configuration file with default setting...")
						// Overwrite the file if the json file does not correspond to the schema
						initializeBasicConfigDefault()
						bytes, err := json.MarshalIndent(BasicConfig, "", "  ")
						if err != nil {
							return err
						}
						err = os.WriteFile(config.Settings.BaiscConfigPath, bytes, 0666)
						if err != nil {
							return err
						}
						break
					} else {
						// The part to ask the user whether to exit if there is a problem with the configuration file
						instruction4SettingBasicConfig()
						exit_choice := confirmation_interface.ConfirmationInterface(
							"\033[31mThe config reading session will end now, do you want to exit?\033[0m",
							false,
						)
						if exit_choice {
							return errors.New("The user refuse to overwrite the configuration file, program exit.")
						} else {
							continue
						}
					}
				}
			}
		}
	}
	return nil
}

// Input the config
func getInputConfig(field *string, fieldName string) {
	fmt.Printf("New %s: ", fieldName)
	var new_field_val string
	fmt.Scan(&new_field_val)
	if new_field_val == "" {
		fmt.Printf("The %s cannot be empty.\n", fieldName)
	} else {
		*field = new_field_val
		fmt.Printf("\033[32m%s changed to %s\n\033[0m", fieldName, *field)
	}
}

// Configurations changing
func ChangeConfig() (Error error) {
	// Check whether the configuration file exists
	_, err := os.Stat(config.Settings.BaiscConfigPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Create file if not exist
			initializeBasicConfigDefault()
			_, err := os.Create(config.Settings.BaiscConfigPath)
			if err == nil {
				bytes, _ := json.MarshalIndent(BasicConfig, "", "  ")
				err = os.WriteFile(config.Settings.BaiscConfigPath, bytes, 0666)
				if err != nil {
					return err
				}
			}
			return err
		} else {
			return err
		}
	}
	// Show info
	fmt.Println("The current configuration is:")
	fmt.Println("VocabFieldName: ", BasicConfig.VocabFieldName)
	fmt.Println("DefinitionFieldName: ", BasicConfig.DefinitionFieldName)
	fmt.Println("LLMApiKey: ", BasicConfig.LLMApiKey)
	fmt.Println("LLMProvider: ", BasicConfig.LLMProvider)
	fmt.Println("LLMModelName: ", BasicConfig.LLMModelName)
	fmt.Println("LLMBaseURL: ", BasicConfig.LLMBaseURL)
	fmt.Println("\033[1;34mYou can change the configuration by inputting the new values.\033[0m")
	// Do get info
	do_input := confirmation_interface.ConfirmationInterface("Do you want to alter the configuration?", true)
	if do_input {
		var continue_inputting bool = true
		for continue_inputting {
			fmt.Println("Input the name of the field you want to change and the new value. ")
			var input string
			fmt.Scan(&input)
			switch input {
			case "VocabFieldName":
				getInputConfig(&BasicConfig.VocabFieldName, "VocabFieldName")
			case "DefinitionFieldName":
				getInputConfig(&BasicConfig.DefinitionFieldName, "DefinitionFieldName")
			case "LLMApiKey":
				getInputConfig(&BasicConfig.LLMApiKey, "LLMApiKey")
			case "LLMProvider":
				getInputConfig(&BasicConfig.LLMProvider, "LLMProvider")
			case "LLMModelName":
				getInputConfig(&BasicConfig.LLMModelName, "LLMModelName")
			case "LLMBaseURL":
				getInputConfig(&BasicConfig.LLMBaseURL, "LLMBaseURL")
			default:
				fmt.Println("Invalid input. ")
			}
			// If the user wants to end the inputting
			continue_inputting = confirmation_interface.ConfirmationInterface("Continue changing the configuration?", false)
			if !continue_inputting {
				fmt.Println("Configuration changing finished. ")
				bytes, _ := json.MarshalIndent(BasicConfig, "", "  ")
				err = os.WriteFile(config.Settings.BaiscConfigPath, bytes, 0666)
				if err != nil {
					return err
				}
				break
			}
		}
	} else {
		fmt.Println("Configuration unchanged.")
	}
	return nil
}
