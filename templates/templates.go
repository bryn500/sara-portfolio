package templates

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"saraweb/startup"
)

func BuildTemplates(config startup.AppConfig) error {
	sharedTemplates := []string{
		config.InputDir + "shared/base.html",
		config.InputDir + "shared/nav.html",
		config.InputDir + "shared/footer.html",
	}

	fmt.Printf("input dir: %s\n", config.InputDir)

	files, err := os.ReadDir(config.InputDir)
	if err != nil {
		return err
	}

	// // example parallel
	// errChan := make(chan error, len(files))
	// var wg sync.WaitGroup
	// for _, file := range files {
	// 	wg.Add(1)
	// 	go func(fileName string) {
	// 		defer wg.Done()
	// 	}(fileName)
	// }
	// // Wait for all goroutines to finish.
	// wg.Wait()
	// close(errChan)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()

		if filepath.Ext(fileName) != ".html" {
			continue
		}

		fmt.Printf("building file for: %s\n", fileName)

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

		if err := baseTmpl.Execute(outputFile, config.PageData[fileName]); err != nil {
			return err
		}
	}

	return nil
}
