package config

import (
	"github.com/1Vewton/vocabulary_autorecord/utils/env_reader"
	"github.com/joho/godotenv"
)

// Env struct
type env_config struct {
	// Paths
	BaiscConfigPath  string
	VocabListPath    string
	ConfigSchemaPath string
	VocabSchemaPath  string
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

// get pL0
func GetpL0() float64 {
	return Settings.pL0
}

// get pS
func GetpS() float64 {
	return Settings.pS
}

// get pS
func GetpG() float64 {
	return Settings.pG
}

// get pT
func GetpT() float64 {
	return Settings.pT
}

// Initialize settings
func InitializeSettings() {
	Settings.pL0 = env_reader.GetEnvFloat64("PL0", 0.5)
	Settings.pT = env_reader.GetEnvFloat64("PT", 0.14)
	Settings.pG = env_reader.GetEnvFloat64("PG", 0.14)
	Settings.pS = env_reader.GetEnvFloat64("PS", 0.09)
	Settings.BaiscConfigPath = env_reader.GetEnvString("BASIC_CONFIG_PATH", "configuration.json")
	Settings.VocabListPath = env_reader.GetEnvString("VOCAB_LIST_PATH", "vocab_list.json")
	Settings.VocabFieldNane = env_reader.GetEnvString("VOCAB_FIELD_NAME", "vocabulary")
	Settings.DefinitionFieldNane = env_reader.GetEnvString("DEFINITION_FIELD_NAME", "definition")
	Settings.ConfigSchemaPath = env_reader.GetEnvString("CONFIG_SCHEMA_PATH", "config_schema.json")
	Settings.VocabSchemaPath = env_reader.GetEnvString("VOCAB_SCHEMA_PATH", "vocab_schema.json")
}
