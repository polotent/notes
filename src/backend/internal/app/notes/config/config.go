package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

type Config struct {
	Database struct {
		Dialect  string
		Host     string
		Port     string
		DBname   string
		Username string
		Password string
	}

	Server struct {
		Port string
	}
}

var (
	AppConfigPath = "../../internal/app/notes/resources/config/application.%s.json"
)

func LoadAppConfig() (*Config, string) {
	var env *string
	if appEnv := os.Getenv("WEB_APP_ENV"); appEnv != "" {
		env = &appEnv
	} else {
		env = flag.String("env", "develop", "To switch configurations.")
	}

	jsonFilePath := fmt.Sprintf(AppConfigPath, *env)
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		fmt.Printf("Error while opening file for reading at %s", jsonFilePath)
		os.Exit(2)
	}

	defer jsonFile.Close()

	config := &Config{}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("Error while reading file content at %s", jsonFilePath)
		os.Exit(2)
	}

	if err := json.Unmarshal([]byte(byteValue), config); err != nil {
		fmt.Printf("Error while unmarshalling config: %s", err)
		os.Exit(2)
	}
	return config, *env
}
