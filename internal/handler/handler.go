package handler

import "context"

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Run(ctx context.Context) error {
	return nil
}
