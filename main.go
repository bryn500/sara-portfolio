package main

import (
	"log"
	"os"
	"saraweb/startup"
	"saraweb/templates"
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

	config, err := startup.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config,json: %v", err)
		return
	}

	if err := templates.BuildTemplates(config); err != nil {
		log.Fatalf("Error building templates: %v", err)
		return
	}

	if isDev {
		Start(config)
	}
}
