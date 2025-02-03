package config

import (
	"fmt"
	"strings"
)

func parseDBConfig() (map[string]string, error) {
	result := make(map[string]string)
	envVars, err := GetEnvVars()
	if err != nil {
		return nil, err
	}
	for _, variable := range envVars {
		if variable[:3] == "DB_" {
			split := strings.SplitN(variable, "=", 2)
			result[split[0]] = split[1]
		}
	}

	return result, nil
}

func GenerateDSN(dbType string) (string, error) {
	var result string
	vars, _ := parseDBConfig()
	switch dbType {
	case "postgres":
		result = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s %s", vars["DB_HOST"], vars["DB_USER"], vars["DB_PASS"], vars["DB_NAME"], vars["DB_PORT"], vars["DB_OPTIONS"])
	case "mysql":
		result = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", vars["DB_USER"], vars["DB_PASS"], vars["DB_HOST"], vars["DB_PORT"], vars["DB_NAME"], vars["DB_OPTIONS"])
	}

	return result, nil
}
