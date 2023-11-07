package main

import (
	"log"
	"os"
)

func main() {
	// Get the command-line arguments
	args := os.Args

	// Check if the "--dev" argument is present
	isDev := false
	for _, arg := range args {
		if arg == "--dev" {
			isDev = true
			break
		}
	}

	config, err := LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config,json: %v", err)
		return
	}

	if err := BuildTemplates(config); err != nil {
		log.Fatalf("Error building templates: %v", err)
		return
	}

	if isDev {
		StartServer(config)
	}
}
