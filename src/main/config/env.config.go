package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	DB_NAME string
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASS string
}

func LoadVars(envFilePath string) *EnvVars {

	if envFilePath != "" {
		godotenv.Load(envFilePath)
	}

	return &EnvVars{
		DB_NAME: os.Getenv("DB_NAME"),
		DB_HOST: os.Getenv("DB_HOST"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_USER: os.Getenv("DB_USER"),
		DB_PASS: os.Getenv("DB_PASS"),
	}

}
