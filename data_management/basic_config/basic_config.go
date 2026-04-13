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
		schema_content, err := os.ReadFile("config_schema.json")
		if err != nil {
			return err
		} else {
			schema_content_string := string(schema_content)
			result, err := json_validator.Validate(schema_content_string, contentString)
			if err != nil {
				return err
			}
			if !result {
			ASK:
				choice := confirmation_interface.ConfirmationInterface(
					"The schema of the configuration file is not correspond to the schema, do you want to overwrite the configuration file with default setting?",
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
				} else {
					// The part to ask the user whether to exit if there is a problem with the configuration file
					instruction4SettingBasicConfig()
					exit_choice := confirmation_interface.ConfirmationInterface(
						"The program will exit now, do you want to exit?",
						false,
					)
					if exit_choice {
						return errors.New("The user refuse to overwrite the configuration file, program exit.")
					} else {
						goto ASK
					}
				}
			}
		}
	}
	return nil
}
