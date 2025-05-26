package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pigen-dev/pigen-cli/internal/core"
	"github.com/spf13/viper"
)

type Config struct {
	PigenCore PigenCore `yaml:"pigen_core"`
}

type PigenCore struct {
	CloudProvider CloudProvider `yaml:"cloud_provider"`
	Endpoint	string `yaml:"endpoint"`
}

type CloudProvider struct {
	Type string `yaml:"type"`
	ProjectID string `yaml:"project_id"`
	Region    string `yaml:"region"`
}

func (c *Config) InitConfig() error {
	if c.PigenCore.CloudProvider.Type == "GCP" {
		gcp := core.PigenCoreGCP{
			ProjectID: c.PigenCore.CloudProvider.ProjectID,
			Region:    c.PigenCore.CloudProvider.Region,
		}
		endpoint, err := gcp.DeployPigenCore()
		if err != nil {
			return fmt.Errorf("error deploying PigenCore on GCP: %w", err)
		}
		c.PigenCore.Endpoint = endpoint
		
	} else{
		return fmt.Errorf("unsupported cloud provider: %s", c.PigenCore.CloudProvider.Type)
	}
	viper.Set("config", c)
	err := writeViperConfig()
	if err != nil {
		return fmt.Errorf("error writing config: %w", err)
	}
	
	if err := viper.ReadInConfig(); err != nil {
			return fmt.Errorf("error reading config: %w", err)
	}
	return nil
}

func writeViperConfig() error {

	configPath := viper.ConfigFileUsed()

	// Ensure the directory exists
	if err := os.MkdirAll(filepath.Dir(configPath), os.ModePerm); err != nil {
			return err
	}

	// Decide whether to init or update
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
			// File doesn't exist -> safe to init
			err = viper.SafeWriteConfig()
	} else {
			// File exists -> overwrite with new values
			err = viper.WriteConfig()
	}
	return nil
}