package plugin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	shared "github.com/pigen-dev/shared"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type PluginFile struct {
	Plugins []shared.PluginStruct `yaml:"plugins"`
}

func (p *PluginFile) GetPlugins(filePath string) (string, error) {
	fmt.Println("Plugins file path:", filePath)
	// Read the file content
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read plugin file: %w", err)
	}
	err = yaml.Unmarshal(fileContent, p)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal plugin file: %w", err)
	}
	pigenCoreEndpoint := fmt.Sprintf("%s/api/v1/plugin", viper.GetString("config.pigen_core.endpoint"))
	return pigenCoreEndpoint, nil
}

func PluginPostRequest(plugin shared.PluginStruct, endpoint string) (*http.Response, error) {
	// 1. Send the request to the Pigen Core
	jsonData, err := json.MarshalIndent(plugin, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal plugin data into json: %w", err)
	}
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	return resp, nil

}