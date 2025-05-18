package pipeline

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/pigen-dev/pigen-cli/helpers"
	"github.com/pigen-dev/pigen-cli/pkg"
	shared "github.com/pigen-dev/shared"
	"github.com/spf13/viper"
)

type GenerateScriptResponse struct {
	pkg.PigenCoreResponse
	Content string
}


func GenerateScript(pigenStepsPath string) error {
	var pigenStepsFile shared.PigenStepsFile
	err := helpers.ReadYamlFile(pigenStepsPath, &pigenStepsFile)
	if err != nil {
		return fmt.Errorf("failed to read pigen steps file: %w", err)
	}
	jsonData, err := helpers.StructToJson(pigenStepsFile)
	if err != nil {
		return fmt.Errorf("failed to convert pigen steps file to json: %w", err)
	}
	// Replace secrets in the JSON data
	jsonData, err = helpers.ReplaceSecrets(jsonData)
	if err != nil {
		return fmt.Errorf("failed to replace secrets: %w", err)
	}
	err = generate(jsonData)
	if err != nil {
		return fmt.Errorf("failed to generate script: %w", err)
	}
	return nil
}

func generate(jsonPigenSteps []byte) error {
	var coreResp GenerateScriptResponse
	pigenCoreEndpoint := fmt.Sprintf("%s/api/v1/cicd/gen_script", viper.GetString("config.pigen_core.endpoint"))
	resp, err := http.Post(pigenCoreEndpoint, "application/json", bytes.NewBuffer(jsonPigenSteps))
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
		return fmt.Errorf("failed to create trigger: %v", coreResp.Error)
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