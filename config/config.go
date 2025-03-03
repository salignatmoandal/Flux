package config

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
}
