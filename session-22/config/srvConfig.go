package config

import "strings"

func ParseSrvConfig() (map[string]string, error) {
	result := make(map[string]string)
	envVars, err := GetEnvVars()
	if err != nil {
		return nil, err
	}
	for _, variable := range envVars {
		if variable[:4] == "SRV_" {
			split := strings.SplitN(variable, "=", 2)
			result[split[0]] = split[1]
		}
	}

	return result, nil
}
