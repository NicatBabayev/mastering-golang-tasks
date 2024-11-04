package json

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	App_name   string   `json: "app_name"`
	Version    string   `json: "version"`
	Debug_mode bool     `json: "debug_mode"`
	Services   []string `json: "services"`
}

func parseJSON2(fileName string) config {
	var config config

	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config)

	return config
}

func RunTask4() {
	config := parseJSON2("task/json/config.json")
	var debug_mode string
	if config.Debug_mode {
		debug_mode = "Enabled"
	} else {
		debug_mode = "Disabled"
	}
	fmt.Println("Configuration Details:")
	fmt.Printf("- App Name: %s\n", config.App_name)
	fmt.Printf("- Version: %s\n", config.Version)
	fmt.Printf("- Debug Mode: %s\n", debug_mode)
	fmt.Printf("- Services: %v\n", config.Services)
}
