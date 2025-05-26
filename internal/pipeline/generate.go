package pipeline

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	//"github.com/pigen-dev/pigen-cli/internal/plugin"
	"github.com/pigen-dev/pigen-cli/helpers"
	"github.com/pigen-dev/pigen-cli/internal/plugin"
	"github.com/pigen-dev/pigen-cli/pkg"
	shared "github.com/pigen-dev/shared"
	"github.com/spf13/viper"
)

type GenerateScriptResponse struct {
	pkg.PigenCoreResponse
	Content string
}


func GenerateScript(pigenStepsPath string) error {
	yamlFile, err := helpers.ReadYamlFile(pigenStepsPath)
	if err != nil {
		return fmt.Errorf("failed to read pigen steps file: %w", err)
	}
	// Replace the plugin outputs in the pigen file
	yamlFile, err = plugin.PigenFileParser(yamlFile)
	if err != nil {
		return fmt.Errorf("failed to replace secrets: %w", err)
	}
	fmt.Println("yamlfile: ", string(yamlFile))
	err = generate(yamlFile)
	if err != nil {
		return fmt.Errorf("failed to generate script: %w", err)
	}
	return nil
}

func generate(pigenSteps []byte) error {
	var coreResp GenerateScriptResponse
	var pigenStepsFile shared.PigenStepsFile
	err := helpers.YamlToStruct(pigenSteps, &pigenStepsFile)
	if err != nil {
		return fmt.Errorf("failed to unmarshal pigen steps file: %w", err)
	}
	jsonData, err := helpers.StructToJson(pigenStepsFile)
	if err != nil {
		return fmt.Errorf("failed to convert pigen steps file to json: %w", err)
	}
	pigenCoreEndpoint := fmt.Sprintf("%s/api/v1/cicd/gen_script", viper.GetString("config.pigen_core.endpoint"))
	resp, err := http.Post(pigenCoreEndpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &coreResp)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %v", err.Error())
	}
	if coreResp.Error != "" {
		return fmt.Errorf("core response error : %v", coreResp.Error)
	}
	data, err := base64.StdEncoding.DecodeString(coreResp.Content)
	if err != nil {
		return fmt.Errorf("failed to decode base64 content: %v", err.Error())
	}
	err = os.WriteFile("cloudbuild.yaml", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
	return nil
}