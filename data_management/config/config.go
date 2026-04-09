package config

import (
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

// initialize configuration
func InitializeConfig(env_path string) (err error) {
	Err := godotenv.Load(env_path)
	return Err
}
