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
	// --- 1. Check for Command-Line Arguments ---
	if len(os.Args) < 2 {
		// os.Args[0] is the program name itself
		fmt.Fprintf(os.Stderr, "Usage: %s <filename.json | filename.yaml>\n", os.Args[0])
		os.Exit(1)
	}
	filename := os.Args[1]

	// --- 2. Read the File ---
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file '%s': %v", filename, err)
	}

	// This generic `interface{}` will hold our unmarshalled data
	var content interface{}
	var prettyData []byte // This will hold our formatted output

	// Create our color printers
	jsonColor := color.New(color.FgCyan)
	yamlColor := color.New(color.FgMagenta)

	// --- 3. Detect File Type and Process ---
	ext := filepath.Ext(filename)

	switch ext {
	case ".json":
		// --- 4. Process JSON ---
		
		// Unmarshal the messy JSON into our generic interface
		if err := json.Unmarshal(data, &content); err != nil {
			log.Fatalf("Invalid JSON file: %v", err)
		}

		// Marshal it back into a pretty, indented byte slice
		// "" for no prefix, "  " for 2-space indentation
		prettyData, err = json.MarshalIndent(content, "", "  ")
		if err != nil {
			log.Fatalf("Failed to format JSON: %v", err)
		}

		// Print the colorized result
		jsonColor.Println(string(prettyData))

	case ".yml", ".yaml":
		// --- 5. Process YAML ---

		// Unmarshal the messy YAML
		if err := yaml.Unmarshal(data, &content); err != nil {
			log.Fatalf("Invalid YAML file: %v", err)
		}

		// Marshal it back into pretty YAML.
		// gopkg.in/yaml.v3 formats by default.
		prettyData, err = yaml.Marshal(content)
		if err != nil {
			log.Fatalf("Failed to format YAML: %v", err)
		}

		// Print the colorized result
		yamlColor.Println(string(prettyData))

	default:
		// --- 6. Handle Unknown Files ---
		log.Fatalf("Unsupported file type: %s. Please use .json, .yml, or .yaml", ext)
	}
}