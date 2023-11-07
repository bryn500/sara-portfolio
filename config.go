package main

import (
	"encoding/json"
	"os"
)

type PageData struct {
	Title       string
	Description string
	Body        string
}

type AppConfig struct {
	InputDir  string
	OutputDir string
	Port      string
	PageData  map[string]PageData
}

func LoadConfig(filename string) (AppConfig, error) {
	var config AppConfig

	file, err := os.Open(filename)
	if err != nil {
		return AppConfig{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return AppConfig{}, err
	}

	return config, nil
}
