package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Check if oapi-codegen is installed
	if _, err := exec.LookPath("oapi-codegen"); err != nil {
		log.Fatal("oapi-codegen not found. Please install with: go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest")
	}

	// Define the APIs to generate
	apis := []struct {
		Name string
		File string
	}{
		{Name: "todos", File: "api-specs/todos-api.yaml"},
		{Name: "categories", File: "api-specs/categories-api.yaml"},
	}

	// Create output directory
	err := os.MkdirAll("generated", 0755)
	if err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Generate code for each API
	for _, api := range apis {
		outputDir := filepath.Join("generated", api.Name)
		err := os.MkdirAll(outputDir, 0755)
		if err != nil {
			log.Fatalf("Failed to create output directory for %s: %v", api.Name, err)
		}

		outputFile := filepath.Join(outputDir, api.Name+".go")
		fmt.Printf("Generating code for %s API...\n", api.Name)

		// Run the code generator
		cmd := exec.Command(
			"oapi-codegen",
			"--package", api.Name,
			"--generate", "types,client,server,spec",
			"-o", outputFile,
			api.File,
		)

		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("Failed to generate code for %s: %v\n%s", api.Name, err, output)
		}

		fmt.Printf("Generated %s\n", outputFile)
	}

	fmt.Println("Code generation complete!")
}
