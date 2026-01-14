package coinbase

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const (
	createOrderMethod = http.MethodPost
	createOrderPath   = "/api/v3/brokerage/orders"
)

type Order struct {
	ProductId string `json:"product_id"`
	Side      string `json:"side"`
	QuoteSize string `json:"quote_size,omitempty"`
	BaseSize  string `json:"base_size,omitempty"`
}

type CreateOrderRequest struct {
	ClientOrderId      string             `json:"client_order_id"`
	ProductId          string             `json:"product_id"`
	Side               string             `json:"side"`
	OrderConfiguration OrderConfiguration `json:"order_configuration"`
}

type CreateOrderResponse struct {
	Success            bool               `json:"success"`
	SuccessResponse    *SuccessResponse   `json:"success_response,omitempty"`
	ErrorResponse      *ErrorResponse     `json:"error_response,omitempty"`
	OrderConfiguration OrderConfiguration `json:"order_configuration"`
}

func (c *Client) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*CreateOrderResponse, error) {
	var resp CreateOrderResponse
	err := c.request(ctx, createOrderMethod, createOrderPath, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func NewCreateOrderRequest(order *Order) *CreateOrderRequest {
	cfg := OrderConfiguration{
		MarketIoc: &MarketIoc{
			QuoteSize: order.QuoteSize,
			BaseSize:  order.BaseSize,
		},
	}

	return &CreateOrderRequest{
		ClientOrderId:      uuid.NewString(),
		ProductId:          order.ProductId,
		Side:               order.Side,
		OrderConfiguration: cfg,
	}
}
