package plugin

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/pigen-dev/pigen-cli/helpers"
	shared "github.com/pigen-dev/shared"
)

func GetAllOutputs(pluginFile PluginFile) (map[string]any, error) {
	
	outputs := make(map[string]any)
	for _, plugin := range pluginFile.Plugins {
		output := GetOutputData(plugin)
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

func GetOutput(pluginName string) (shared.GetOutputResponse) {
	outputResp := shared.GetOutputResponse{}
	pluginFile := &PluginFile{}
	err := helpers.GetYamlFileSafe("pigen-plugins.yaml", pluginFile)
	if err != nil {
		outputResp.Error = fmt.Errorf("failed to read plugin file: %w", err)
		return outputResp
	}

	found := false
	for _, plugin := range pluginFile.Plugins {
		if plugin.Plugin.Label == pluginName {
			found = true
			fmt.Println("⏳ Outputting plugin:", pluginName)
			// Send the get output request to the Pigen Core
			err := PluginParser(&plugin)
			if err != nil {
				outputResp.Error = fmt.Errorf("can't parse plugin: %w", err)
				return outputResp
			}
			outputResp = GetOutputData(plugin)
			break
		}
	}
	
	if !found {
		outputResp.Error = fmt.Errorf("plugin %s not found in the plugin file, this can be caused if plugin is not installed or plugin isn't installed from this plugins file", pluginName)
	}
	
	return outputResp
}

func GetOutputData(plugin shared.PluginStruct) (shared.GetOutputResponse) {
	outputResp := shared.GetOutputResponse{}
	// Send the get output request to the Pigen Core
	resp, err := PluginPostRequest(plugin, "/get_output")
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return shared.GetOutputResponse{
			Error: fmt.Errorf("failed to read response body: %w", err),
			Output: nil,
		}
	}
	err = json.Unmarshal(body, &outputResp)
	if err != nil {
		return shared.GetOutputResponse{
			Error: fmt.Errorf("failed to unmarshal response: %w", err),
			Output: nil,
		}
	}
	outputResp.Output = helpers.CleanPluginOutput(outputResp.Output)
	return outputResp
}