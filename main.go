package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <filename.json | filename.yaml>\n", os.Args[0])
		os.Exit(1)
	}
	filename := os.Args[1]

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file '%s': %v", filename, err)
	}

	var content interface{}
	var prettyData []byte

	jsonColor := color.New(color.FgCyan)
	yamlColor := color.New(color.FgMagenta)

	ext := filepath.Ext(filename)

	switch ext {
	case ".json":
		
		if err := json.Unmarshal(data, &content); err != nil {
			log.Fatalf("Invalid JSON file: %v", err)
		}
		
		prettyData, err = json.MarshalIndent(content, "", "  ")
		if err != nil {
			log.Fatalf("Failed to format JSON: %v", err)
		}

		jsonColor.Println(string(prettyData))

	case ".yml", ".yaml":

		if err := yaml.Unmarshal(data, &content); err != nil {
			log.Fatalf("Invalid YAML file: %v", err)
		}

		prettyData, err = yaml.Marshal(content)
		if err != nil {
			log.Fatalf("Failed to format YAML: %v", err)
		}

		yamlColor.Println(string(prettyData))

	default:
		log.Fatalf("Unsupported file type: %s. Please use .json, .yml, or .yaml", ext)
	}

}
