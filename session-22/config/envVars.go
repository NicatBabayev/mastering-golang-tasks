package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func GetEnvVars() ([]string, error) {
	var result []string

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file")
	}
	envVars := os.Environ()
	result = envVars

	return result, nil
}
