package basic_config

import (
	"encoding/json"
	"os"

	"github.com/1Vewton/vocabulary_autorecord/data_management/config"
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
	content, err := os.ReadFile(config.Settings.BaiscConfigPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &BasicConfig)
	if err != nil {
		initializeBasicConfigDefault()
		bytes, _ := json.MarshalIndent(BasicConfig, "", "  ")
		err = os.WriteFile(config.Settings.BaiscConfigPath, bytes, 0666)
		if err != nil {
			return err
		}
	}
	return nil
}
