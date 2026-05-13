package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	GeminiAPIKey string `json:"gemini_api_key"`
	ArchivePath string `json:"archive_path"`
	OutputPath string `json:"output_path"`
}

func LoadConfig(path string) (*Config, error) {
	// Read file
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Create config variable
	var config Config

	// Convert JSOn -> struct
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}