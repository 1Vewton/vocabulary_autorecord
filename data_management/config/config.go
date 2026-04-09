package config

import (
	"github.com/1Vewton/vocabulary_autorecord/utils/env_reader"
	"github.com/joho/godotenv"
)

// Env struct
type env_config struct {
	// Paths
	BaiscConfigPath string
	VocabListPath   string
	// Storage
	VocabFieldNane      string
	DefinitionFieldNane string
	// BKT related
	pL0 float64
	pT  float64
	pG  float64
	pS  float64
}

var Settings env_config

// initialize configuration
func InitializeConfig(env_path string) (err error) {
	Err := godotenv.Load(env_path)
	return Err
}

// Initialize settings
func InitializeSettings() {
	Settings.pL0 = env_reader.GetEnvFloat64("PL0", 0.000001)
	Settings.pT = env_reader.GetEnvFloat64("PT", 0.3)
	Settings.pG = env_reader.GetEnvFloat64("PG", 0.2)
	Settings.pS = env_reader.GetEnvFloat64("PS", 0.05)
	Settings.BaiscConfigPath = env_reader.GetEnvString("BASIC_CONFIG_PATH", "configuration.json")
	Settings.VocabListPath = env_reader.GetEnvString("VOCAB_LIST_PATH", "vocab_list.json")
	Settings.VocabFieldNane = env_reader.GetEnvString("VOCAB_FIELD_NAME", "vocabulary")
	Settings.DefinitionFieldNane = env_reader.GetEnvString("DEFINITION_FIELD_NAME", "definition")
}
