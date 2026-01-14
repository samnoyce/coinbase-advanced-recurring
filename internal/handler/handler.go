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

type Event struct {
	Orders []coinbase.Order `json:"orders"`
}

func New(client *coinbase.Client, logger *slog.Logger) *Handler {
	return &Handler{
		client: client,
		logger: logger,
	}
}

func (h *Handler) Run(ctx context.Context, event Event) error {
	for _, order := range event.Orders {
		req := coinbase.NewCreateOrderRequest(&order)

		resp, err := h.client.CreateOrder(ctx, req)
		if err != nil {
			h.logger.Error(
				"Failed to create order",
				"client_order_id", req.ClientOrderId,
				"side", req.Side,
				"product_id", req.ProductId,
				"error", err,
			)
			continue
		}

		if !resp.Success {
			h.logger.Error(
				"Order rejected",
				"client_order_id", req.ClientOrderId,
				"side", req.Side,
				"product_id", req.ProductId,
				"error", resp.ErrorResponse.ErrorDetails,
			)
			continue
		}

		h.logger.Info(
			"Order created",
			"client_order_id", resp.SuccessResponse.ClientOrderId,
			"side", resp.SuccessResponse.Side,
			"product_id", resp.SuccessResponse.ProductId,
			"order_id", resp.SuccessResponse.OrderId,
		)
	}

	return nil
}
