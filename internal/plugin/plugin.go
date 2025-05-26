package plugin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	//"github.com/pigen-dev/pigen-cli/internal/templater"
	shared "github.com/pigen-dev/shared"
	"github.com/spf13/viper"
)

type PluginFile struct {
	Plugins []shared.PluginStruct `yaml:"plugins"`
}




func PluginPostRequest(plugin shared.PluginStruct, endpoint string) (*http.Response, error) {

	pigenCoreEndpoint := fmt.Sprintf("%s/api/v1/plugin%s", viper.GetString("config.pigen_core.endpoint"), endpoint)
	// 1. Send the request to the Pigen Core
	jsonData, err := json.MarshalIndent(plugin, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal plugin data into json: %w", err)
	}
	resp, err := http.Post(pigenCoreEndpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	return resp, nil

}
