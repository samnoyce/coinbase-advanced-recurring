package handler

import (
	"context"

	"coinbase-advanced-recurring/internal/coinbase"
)

type Handler struct {
	client *coinbase.Client
}

func New(client *coinbase.Client) *Handler {
	return &Handler{client: client}
}

func (h *Handler) Run(ctx context.Context) error {
	return nil
}
