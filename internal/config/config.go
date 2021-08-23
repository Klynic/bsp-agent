package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type GcpConfig struct {
	ServiceAccount string `envconfig:"GCP_SERVICE_ACCOUNT"`
	ProjectId      string `envconfig:"GCP_PROJECT_ID"`
	BucketName     string `envconfig:"GCP_BUCKET_NAME"`
}

type RedisConfig struct {
	Address  string `envconfig:"REDIS_ADDRESS"`
	Password string `envconfig:"REDIS_PASSWORD" default:""`
	DB       int    `envconfig:"REDIS_DB" default:"0"`
	Key      string `envconfig:"REDIS_STREAM_KEY" default:"replication"`
	Group    string `envconfig:"REDIS_CONSUMER_GROUP"`
}

type Config struct {
	GcpConfig   GcpConfig
	RedisConfig RedisConfig
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to parse configuration: %w", err)
	}

	return cfg, nil
}
