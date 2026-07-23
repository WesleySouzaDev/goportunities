package config

import (
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("error loading .env file")
	}
}
