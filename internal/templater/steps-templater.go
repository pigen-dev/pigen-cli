package templater

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/joho/godotenv"
)

type PigenTemplater struct {
	Plugins map[string]any
	ENV		 map[string]string
}

func PigenReplacer(file []byte, pluginOutputs map[string]any) ([]byte, error) {
	// Replace the plugin outputs in the pigen file
	templater := PigenTemplater{
		Plugins: pluginOutputs,
	}
	secretsMap, err := loadEnvFile(".env.pigen")
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}
	templater.ENV = secretsMap
	temp, err := template.New("pigen").Parse(string(file))
	if err != nil {
		return nil, err
	}
	var replacedFile bytes.Buffer
	err = temp.Execute(&replacedFile, templater)
	if err != nil {
		return nil, err
	}
	return replacedFile.Bytes(), nil
}

func loadEnvFile(path string) (map[string]string, error) {
    return godotenv.Read(path)
}