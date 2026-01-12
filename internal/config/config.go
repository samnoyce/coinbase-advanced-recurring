package config

import (
	"fmt"
	"os"
)

const (
	envAppEnv     = "APP_ENV"
	envSecretName = "COINBASE_SECRET_NAME"
)

type Config struct {
	AppEnv     string
	SecretName string
}

func Load() (*Config, error) {
	appEnv := envOrDefault(envAppEnv, "dev")

	secretName, err := requireEnv(envSecretName)
	if err != nil {
		return nil, err
	}

	return &Config{
		AppEnv:     appEnv,
		SecretName: secretName,
	}, nil
}

func envOrDefault(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}

func requireEnv(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return "", fmt.Errorf("environment variable %q must be set and non-empty", key)
	}
	return v, nil
}
