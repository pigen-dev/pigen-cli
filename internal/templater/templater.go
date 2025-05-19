package templater

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/joho/godotenv"
	"github.com/pigen-dev/pigen-cli/internal/plugin"
)

type PigenTemplater struct {
	ENV map[string]string
	Plugins map[string]any
}

func PigenReplacer(pigenFile []byte, pigenPluginsFilePath string) ([]byte, error) {
	pigenTemplater := PigenTemplater{}
	// Replace the secrets in the pigen file
	err := pigenTemplater.LoadENV()
	if err != nil {
		return nil, err
	}
	// Replace the plugin outputs in the pigen file
	err = pigenTemplater.LoadOutputs(pigenPluginsFilePath)
	if err != nil {
		return nil, err
	}
	// Replace the plugin outputs in the pigen file
	tmpl, err := template.New("pigen").Parse(string(pigenFile))
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, pigenTemplater)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}
	return buf.Bytes(), nil
}

func (t *PigenTemplater) LoadOutputs(pigenPluginsFilePath string) (error) {
	pluginOutputs, err := plugin.GetAllOutputs(pigenPluginsFilePath)
	if err != nil {
		return fmt.Errorf("failed to get all outputs: %w", err)
	}
	t.Plugins = pluginOutputs
	return nil
}

func (t *PigenTemplater) LoadENV() (error) {
	// Load the .env.pigen file
	secretsMap, err := loadEnvFile(".env.pigen")
	if err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}
	t.ENV = secretsMap
	return nil
}

func loadEnvFile(path string) (map[string]string, error) {
    return godotenv.Read(path)
}