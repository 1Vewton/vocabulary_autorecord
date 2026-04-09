package config

import (
	"github.com/joho/godotenv"
)

// initialize configuration
func init() {
	godotenv.Load(".env")
}
