package pipeline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/pigen-dev/pigen-cli/helpers"
	"github.com/pigen-dev/pigen-cli/pkg"
	shared "github.com/pigen-dev/shared"
	"github.com/spf13/viper"
)

type ConnectRepoResponse struct {
    ActionRequired   string `json:"Action is required,omitempty"`
		Error      			error  `json:"error,omitempty"`
  	Message         string `json:"message,omitempty"`
}

func SetupPipeline(pigenStepsPath string) error {
	var pigenStepsFile shared.PigenStepsFile
	err := helpers.ReadYamlFile(pigenStepsPath, &pigenStepsFile)
	if err != nil {
		return fmt.Errorf("failed to read pigen steps file: %w", err)
	}
	jsonData, err := helpers.StructToJson(pigenStepsFile)
	if err != nil {
		return fmt.Errorf("failed to convert pigen steps file to json: %w", err)
	}
	// Link github repo to cicd tool
	resp, err := connect_repo(jsonData)
	if err != nil {
		return fmt.Errorf("failed to connect repo: %w", err)
	}

	if resp.ActionRequired != "" {
		fmt.Println("Action is required, please follow this link to complete the action:", resp.ActionRequired)
		
		timeout := 0
		// Check if action is required
		// If action is required, wait for 5 seconds and check again
		// Timeout after 36 * 5 seconds
		// We will wait for 3 minutes
		for resp.ActionRequired != "" && timeout < 36 {
			time.Sleep(5 * time.Second)
			resp, err = connect_repo(jsonData)
			if err != nil {
				return fmt.Errorf("failed to connect repo: %w", err)
			}
			timeout++
		}
	
		if timeout == 36 {
			return fmt.Errorf("timeout while waiting for action to be completed please run again the command")
		}
	}
	fmt.Println(resp.Message)
	// Create trigger
	err = create_trigger(jsonData)
	if err != nil {
		return fmt.Errorf("failed to create trigger: %w", err)
	}
	return nil
}

func connect_repo(jsonPigenSteps []byte) (repoResp *ConnectRepoResponse, err error) {
	pigenCoreEndpoint := fmt.Sprintf("%s/api/v1/cicd/connect_repo", viper.GetString("config.pigen_core.endpoint"))
	resp, err := http.Post(pigenCoreEndpoint, "application/json", bytes.NewBuffer(jsonPigenSteps))
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &repoResp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return repoResp, nil
}

func create_trigger(jsonPigenSteps []byte) error {
	var coreResp pkg.PigenCoreResponse
	pigenCoreEndpoint := fmt.Sprintf("%s/api/v1/cicd/create_trigger", viper.GetString("config.pigen_core.endpoint"))
	resp, err := http.Post(pigenCoreEndpoint, "application/json", bytes.NewBuffer(jsonPigenSteps))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	err = json.Unmarshal(body, &coreResp)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if coreResp.Error != "" {
		return fmt.Errorf("failed to create trigger: %s", coreResp.Error)
	}
	//fmt.Println(coreResp.Message)
	//Fix returned message
	fmt.Println("Trigger created successfully")
	return nil
}