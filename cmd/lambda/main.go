package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"

	"coinbase-advanced-recurring/internal/coinbase"
	"coinbase-advanced-recurring/internal/config"
	"coinbase-advanced-recurring/internal/handler"
	"coinbase-advanced-recurring/internal/secret"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	secretClient, err := secret.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create Secrets Manager client: %v", err)
	}

	s, err := secretClient.Fetch(ctx, cfg.SecretName)
	if err != nil {
		log.Fatalf("Failed to fetch secret %q: %v", cfg.SecretName, err)
	}

	coinbaseClient, err := coinbase.NewClient(cfg.AppEnv, s)
	if err != nil {
		log.Fatalf("Failed to create Coinbase client: %v", err)
	}

	h := handler.New(coinbaseClient)

	lambda.Start(h.Run)
}
