package plugin

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/pigen-dev/pigen-cli/helpers"
	"github.com/pigen-dev/pigen-cli/pkg"
)


func PluginInstall(filePath string) error {
	var coreResp pkg.PigenCoreResponse
	pluginFile := &PluginFile{}
	// Read the file content
	err := helpers.GetYamlFileSafe("pigen-plugins.yaml", pluginFile)
	if err != nil {
		return fmt.Errorf("can't get plugin file safely: %w", err)
	}
	installEndpoint := "/setup_plugin"

	for _, plugin := range pluginFile.Plugins {
		fmt.Println("⏳ Installing plugin:", plugin.Plugin.Label)
		err = PluginParser(&plugin)
		if err != nil {
			return fmt.Errorf("can't parse plugin: %w", err)
		}
		resp, err := PluginPostRequest(plugin, installEndpoint)
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