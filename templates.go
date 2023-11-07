package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

func BuildTemplates(config AppConfig) error {
	sharedTemplates := []string{
		config.InputDir + "shared/base.html",
		config.InputDir + "shared/nav.html",
		config.InputDir + "shared/footer.html",
	}

	files, err := os.ReadDir(config.InputDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()

		if filepath.Ext(fileName) != ".html" {
			continue
		}

		fmt.Printf("building: %s\n", config.InputDir+fileName)

		pageTemplates := append(sharedTemplates, config.InputDir+fileName)
		baseTmpl, err := template.ParseFiles(pageTemplates...)
		if err != nil {
			return err
		}

		outputFile, err := os.Create(config.OutputDir + fileName)
		if err != nil {
			return err
		}
		defer outputFile.Close()

		fmt.Printf("output: %s\n", config.OutputDir+fileName)

		if err := baseTmpl.Execute(outputFile, config.PageData[fileName]); err != nil {
			return err
		}
	}

	return nil
}
