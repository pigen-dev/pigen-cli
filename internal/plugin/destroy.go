package plugin

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/pigen-dev/pigen-cli/helpers"
	"github.com/pigen-dev/pigen-cli/pkg"
)

func DestroyPlugin(filePath, pluginName string) error {
	// 1. Destroy logic here
	pluginFile := &PluginFile{}
	// Read the file content
	coreEndpoint, err := pluginFile.GetPlugins(filePath)
	if err != nil {
		return fmt.Errorf("failed to read plugin file: %w", err)
	}
	DestroyEndpoint := fmt.Sprintf("%s/destroy_plugin", coreEndpoint)

	for _, plugin := range pluginFile.Plugins {
		var coreResp pkg.PigenCoreResponse
		if plugin.Plugin.Label == pluginName {
			fmt.Println("⏳ Destroying plugin:", plugin.Plugin.Label)
			// Send the destroy request to the Pigen Core
			resp, err := PluginPostRequest(plugin, DestroyEndpoint)
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			err = json.Unmarshal(body, coreResp)
			if err != nil {
				return fmt.Errorf("failed to unmarshal response: %w", err)
			}
			
			if coreResp.Error != "" {
				return fmt.Errorf("failed to destroy plugin: %s", coreResp.Error)
			}
			fmt.Println("✅ Plugin destroyed successfully:", plugin.Plugin.Label)
			break
		}
	}
	
	return nil
}

func UpdatePluginYaml(filePath, pluginName string) error {
	// 2. Update the plugin.yaml file
	pluginFile := &PluginFile{}
	// Read the file content
	_, err := pluginFile.GetPlugins(filePath)
	if err != nil {
		return fmt.Errorf("failed to read plugin file: %w", err)
	}

	for index, plugin := range pluginFile.Plugins {
		if plugin.Plugin.Label == pluginName {
			fmt.Println("⏳ Updating plugin.yaml:", plugin.Plugin.Label)
			// Send the destroy request to the Pigen Core
			pluginFile.Plugins = append(pluginFile.Plugins[:index], pluginFile.Plugins[index+1:]...)
			break
		}
	}
	err = helpers.WriteYamlFile(filePath, pluginFile)
	if err != nil {
		return fmt.Errorf("failed to write plugin file: %w", err)
	}
	return nil
}