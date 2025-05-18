package plugin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/pigen-dev/pigen-cli/pkg"
	shared "github.com/pigen-dev/shared"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type PluginFile struct {
	Plugins []shared.PluginStruct `yaml:"plugins"`
}

func PluginInstall(filePath string) error {
	var coreResp pkg.PigenCoreResponse
	var pluginFile PluginFile
	fmt.Println("Installing plugins from file:", filePath)
	// Read the file content
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read plugin file: %w", err)
	}
	err = yaml.Unmarshal(fileContent, &pluginFile)
	if err != nil {
		return fmt.Errorf("failed to unmarshal plugin file: %w", err)
	}
	pigenCoreEndpoint := fmt.Sprintf("%s/api/v1/plugin/setup_plugin", viper.GetString("config.pigen_core.endpoint"))

	for _, plugin := range pluginFile.Plugins {
		fmt.Println("Installing plugin:", plugin.Plugin.Label)
		jsonData, err := json.MarshalIndent(plugin, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal plugin data into json: %w", err)
		}
		resp, err := http.Post(pigenCoreEndpoint, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			return fmt.Errorf("failed to send request: %w", err)
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		err = json.Unmarshal(body, &coreResp)
		if err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
		if coreResp.Error != "" {
			return fmt.Errorf("failed to create trigger: %s", coreResp.Error)
		}
		fmt.Println("Plugin installed successfully:", plugin.Plugin.Label)
	}
	return nil
}