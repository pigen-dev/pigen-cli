package plugin

import (
	"encoding/json"
	"fmt"
	"io"
	shared "github.com/pigen-dev/shared"
)

func GetAllOutputs(filePath string) (map[string]any, error) {
	pluginFile := &PluginFile{}
	// Read the file content
	coreEndpoint, err := pluginFile.GetPlugins(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read plugin file: %w", err)
	}
	outputs := make(map[string]any)
	for _, plugin := range pluginFile.Plugins {
		output := GetOutputData(plugin, coreEndpoint)
		if output.Error != nil {
				return nil, fmt.Errorf("failed to get output for plugin %s: %w", plugin.Plugin.Label, output.Error)
		}
		outputs[plugin.Plugin.Label] = make(map[string]any)
		for key, value := range output.Output {
			outputs[plugin.Plugin.Label].(map[string]any)[key] = value
		}
	}
	
	return outputs, nil
}

func GetOutput(filePath, pluginName string) (shared.GetOutputResponse) {
	outputResp := shared.GetOutputResponse{}
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
	
	found := false
	for _, plugin := range pluginFile.Plugins {
		if plugin.Plugin.Label == pluginName {
			found = true
			fmt.Println("⏳ Outputting plugin:", pluginName)
			// Send the get output request to the Pigen Core
			outputResp = GetOutputData(plugin, coreEndpoint)
			break
		}
	}
	
	if !found {
		outputResp.Error = fmt.Errorf("plugin %s not found in the plugin file, this can be caused if plugin is not installed or plugin isn't installed from this plugins file", pluginName)
	}
	
	return outputResp
}

func GetOutputData(plugin shared.PluginStruct, coreEndpoint string) (shared.GetOutputResponse) {
	outputResp := shared.GetOutputResponse{}
	// Send the get output request to the Pigen Core
	getOutputEndpoint := fmt.Sprintf("%s/get_output", coreEndpoint)
	resp, err := PluginPostRequest(plugin, getOutputEndpoint)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &outputResp)
	if err != nil {
		return shared.GetOutputResponse{
			Error: fmt.Errorf("failed to unmarshal response: %w", err),
			Output: nil,
		}
	}
	return outputResp
}