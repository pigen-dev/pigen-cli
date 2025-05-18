package plugin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/pigen-dev/pigen-cli/pkg"
)


func PluginInstall(filePath string) error {
	var coreResp pkg.PigenCoreResponse
	pluginFile := &PluginFile{}
	// Read the file content
	coreEndpoint, err := pluginFile.GetPlugins(filePath)
	if err != nil {
		return fmt.Errorf("failed to read plugin file: %w", err)
	}
	installEndpoint := fmt.Sprintf("%s/setup_plugin", coreEndpoint)

	for _, plugin := range pluginFile.Plugins {
		fmt.Println("⏳ Installing plugin:", plugin.Plugin.Label)
		jsonData, err := json.MarshalIndent(plugin, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal plugin data into json: %w", err)
		}
		resp, err := http.Post(installEndpoint, "application/json", bytes.NewBuffer(jsonData))
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
		fmt.Println("✅ Plugin installed successfully:", plugin.Plugin.Label)
	}
	return nil
}