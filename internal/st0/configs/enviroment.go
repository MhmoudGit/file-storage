package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config() map[string]string {
	// Load the environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	configVars := map[string]string{
		"DB_USERNAME": os.Getenv("DB_USERNAME"),
		"DB_PASSWORD": os.Getenv("DB_PASSWORD"),
		"DB_HOSTNAME": os.Getenv("DB_HOSTNAME"),
		"DB_NAME":     os.Getenv("DB_NAME"),
	}
	return configVars
}
