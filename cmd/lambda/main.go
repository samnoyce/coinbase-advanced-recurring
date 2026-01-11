package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"coinbase-advanced-recurring/internal/handler"
)

func main() {
	h := handler.New()
	lambda.Start(h.Run)
}
