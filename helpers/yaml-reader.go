package helpers

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadYamlFile(filePath string) ([]byte,error) {
	// Read the file content
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return fileContent, nil
}

func YamlToStruct(file []byte, out any) error {
	// Unmarshal the YAML data into the struct
	err := yaml.Unmarshal(file, out)
	if err != nil {
		return fmt.Errorf("failed to unmarshal YAML: %w", err)
	}
	return nil
}
func StructToJson(in any) ([]byte, error) {
	// Marshal the struct into JSON
	jsonData, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal YAML to JSON: %w", err)
	}
	return jsonData, nil
}

func WriteYamlFile(filePath string, in any) error {
	// Marshal the struct into YAML
	yamlData, err := yaml.Marshal(in)
	if err != nil {
		return fmt.Errorf("failed to marshal struct to YAML: %w", err)
	}
	// Write the YAML data to the file
	err = os.WriteFile(filePath, yamlData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write YAML file: %w", err)
	}
	fmt.Printf("✅ Successfully wrote to %s\n", filePath)
	return nil
}