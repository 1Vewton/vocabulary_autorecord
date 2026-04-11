package basic_config

import (
	"encoding/json"
	"os"

	"github.com/1Vewton/vocabulary_autorecord/data_management/config"
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
				// Overwrite the file if the json file does not correspond to the schema
				initializeBasicConfigDefault()
				bytes, _ := json.MarshalIndent(BasicConfig, "", "  ")
				err = os.WriteFile(config.Settings.BaiscConfigPath, bytes, 0666)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
