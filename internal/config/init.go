package config

import (
	"fmt"

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
			return err
		}
		c.PigenCore.Endpoint = endpoint
		
	} else{
		return fmt.Errorf("unsupported cloud provider: %s", c.PigenCore.CloudProvider.Type)
	}
	viper.Set("config", c)
	viper.WriteConfig()
	if err := viper.ReadInConfig(); err != nil {
			return err
	}
	return nil
}