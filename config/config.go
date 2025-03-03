package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Cloud struct {
		AWS struct {
			Region    string `yaml:"region"`
			AccessKey string `yaml:"access_key"`
			SecretKey string `yaml:"secret_key"`
		} `yaml:"aws"`
		GCP struct {
			ProjectID      string `yaml:"project_id"`
			CredentialFile string `yaml:"credential_file"`
		} `yaml:"gcp"`
		Azure struct {
			SubscriptionID string `yaml:"subscription_id"`
			TenantID       string `yaml:"tenant_id"`
		} `yaml:"azure"`
	} `yaml:"cloud"`
	Analysis struct {
		ResourceThreshold float64 `yaml:"resource_threshold"`
		ScanInterval      string  `yaml:"scan_interval"`
	} `yaml:"analysis"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
