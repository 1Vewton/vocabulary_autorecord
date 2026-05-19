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
}

// Basic Config
var BasicConfig basicConfig

// Initialize Basic Config to default values
func initializeBasicConfigDefault() {
	BasicConfig.DefinitionFieldName = "definition"
	BasicConfig.VocabFieldName = "vocabulary"
}

// Instruction for user to set the basic config
func instruction4SettingBasicConfig() {
	fmt.Println("You refuse to overwrite the configuration file. However, the program cannot run without the basic configuration.")
	fmt.Println("You can set the json file manually. ")
	fmt.Println("The followings are the meaning for the fields: ")
	fmt.Println("- VocabFieldName: the name of the field in the vocabulary list file that contains the raw vocabulary.")
	fmt.Println("- DefinitionFieldName: the name of the field in the vocabulary list file that contains the definition of the vocabulary.")
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
	fmt.Println("\033[1;34mYou can change the configuration by inputting the new values.\033[0m")
	// Do get info
	do_input := confirmation_interface.ConfirmationInterface("Do you want to alter the configuration?", true)
	if do_input {
		fmt.Println("Input the name of the field you want to change and the new value. ")
		var continue_inputting bool = true
		for continue_inputting {
			var input string
			fmt.Scan(&input)
			switch input {
			case "VocabFieldName":
				fmt.Print("New VocabFieldName: ")
				var new_vocab_field_name string
				fmt.Scan(&new_vocab_field_name)
				if new_vocab_field_name == "" {
					fmt.Println("The VocabFieldName cannot be empty.")
				} else {
					BasicConfig.VocabFieldName = new_vocab_field_name
					fmt.Println("VocabFieldName changed to ", new_vocab_field_name)
				}
			case "DefinitionFieldName":
				fmt.Print("New DefinitionFieldName: ")
				var new_definition_field_name string
				fmt.Scan(&new_definition_field_name)
				if new_definition_field_name == "" {
					fmt.Println("The DefinitionFieldName cannot be empty.")
				} else {
					BasicConfig.DefinitionFieldName = new_definition_field_name
					fmt.Println("DefinitionFieldName changed to ", new_definition_field_name)
				}
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
