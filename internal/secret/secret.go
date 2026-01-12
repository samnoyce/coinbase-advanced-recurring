package secret

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type Client struct {
	svc *secretsmanager.Client
}

type Secret struct {
	Name       string `json:"name"`
	PrivateKey string `json:"privateKey"`
}

func NewClient(ctx context.Context) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("loading default config: %w", err)
	}

	return &Client{
		svc: secretsmanager.NewFromConfig(cfg),
	}, nil
}

func (c *Client) Fetch(ctx context.Context, secretName string) (*Secret, error) {
	resp, err := c.svc.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		return nil, fmt.Errorf("fetching secret %q: %w", secretName, err)
	}

	secretString := aws.ToString(resp.SecretString)
	if secretString == "" {
		return nil, fmt.Errorf("secret %q must be set and non-empty", secretName)
	}

	var s Secret
	if err := json.Unmarshal([]byte(secretString), &s); err != nil {
		return nil, fmt.Errorf("decoding secret %q: %w", secretName, err)
	}

	return &s, nil
}
