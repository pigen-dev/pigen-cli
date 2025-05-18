package plugin

import (
	"encoding/json"
	"fmt"
	"io"
	shared "github.com/pigen-dev/shared"
)

func GetOutput(filePath, pluginName string) (shared.GetOutputResponse) {
	var coreResp shared.GetOutputResponse
	// 1. Destroy logic here
	pluginFile := &PluginFile{}
	// Read the file content
	coreEndpoint, err := pluginFile.GetPlugins(filePath)
	if err != nil {
		return shared.GetOutputResponse{
			Error: fmt.Errorf("failed to read plugin file: %w", err),
			Output: nil,
		}
	}
	DestroyEndpoint := fmt.Sprintf("%s/get_output", coreEndpoint)
	found := false
	for _, plugin := range pluginFile.Plugins {
		if plugin.Plugin.Label == pluginName {
			found = true
			fmt.Println("⏳ Outputting plugin:", pluginName)
			// Send the destroy request to the Pigen Core
			resp, err := PluginPostRequest(plugin, DestroyEndpoint)
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			err = json.Unmarshal(body, &coreResp)
			if err != nil {
				return shared.GetOutputResponse{
					Error: fmt.Errorf("failed to unmarshal response: %w", err),
					Output: nil,
				}
			}
			break
		}
	}
	if !found {
		return shared.GetOutputResponse{
			Error: fmt.Errorf("plugin %s not found in the plugin file, this can be caused if plugin is not installed or plugin isn't installed from this plugins file", pluginName),
			Output: nil,
		}
	}
	
	return coreResp
}