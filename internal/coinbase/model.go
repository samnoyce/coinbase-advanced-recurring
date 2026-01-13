package coinbase

type SuccessResponse struct {
	OrderId       string `json:"order_id"`
	ProductId     string `json:"product_id"`
	Side          string `json:"side"`
	ClientOrderId string `json:"client_order_id"`
}

type ErrorResponse struct {
	Error                 string `json:"error"`
	Message               string `json:"message"`
	ErrorDetails          string `json:"error_details"`
	PreviewFailureReason  string `json:"preview_failure_reason"`
	NewOrderFailureReason string `json:"new_order_failure_reason"`
}

type OrderConfiguration struct {
	MarketIoc *MarketIoc `json:"market_market_ioc"`
}

type MarketIoc struct {
	QuoteSize string `json:"quote_size,omitempty"`
	BaseSize  string `json:"base_size,omitempty"`
}
