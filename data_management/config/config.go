package config

import (
	"github.com/1Vewton/vocabulary_autorecord/utils/env_reader"
	"github.com/joho/godotenv"
)

// Env struct
type env_config struct {
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
}
