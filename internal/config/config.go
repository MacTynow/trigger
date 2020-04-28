package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config from environment definition
type Config struct {
	ProjectID     string
	TopicID       string
	ContainerName string
}

// GetConfigFromEnv return the config parsed from the env
func GetConfigFromEnv(prefix string) *Config {
	var config Config
	err := envconfig.Process(prefix, &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &config
}
