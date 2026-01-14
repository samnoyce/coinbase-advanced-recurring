package handler

import (
	"context"
	"log/slog"

	"coinbase-advanced-recurring/internal/coinbase"
)

type Handler struct {
	client *coinbase.Client
	logger *slog.Logger
}

func New(client *coinbase.Client, logger *slog.Logger) *Handler {
	return &Handler{
		client: client,
		logger: logger,
	}
}

func (h *Handler) Run(ctx context.Context) error {
	return nil
}
