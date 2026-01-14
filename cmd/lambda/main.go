package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"coinbase-advanced-recurring/internal/coinbase"
	"coinbase-advanced-recurring/internal/config"
	"coinbase-advanced-recurring/internal/handler"
	"coinbase-advanced-recurring/internal/secret"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		logger.Error("Failed to load config", "error", err)
		return
	}

	secretClient, err := secret.NewClient(ctx)
	if err != nil {
		logger.Error("Failed to create Secrets Manager client", "error", err)
		return
	}

	s, err := secretClient.Fetch(ctx, cfg.SecretName)
	if err != nil {
		logger.Error("Failed to fetch secret", "secret_name", cfg.SecretName, "error", err)
		return
	}

	coinbaseClient, err := coinbase.NewClient(cfg.AppEnv, s)
	if err != nil {
		logger.Error("Failed to create Coinbase client", "app_env", cfg.AppEnv, "error", err)
		return
	}

	h := handler.New(coinbaseClient, logger)

	lambda.Start(h.Run)
}
