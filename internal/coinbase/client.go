package coinbase

import (
	"crypto/ecdsa"
	"net/http"
	"strings"
	"time"

	"coinbase-advanced-recurring/internal/secret"
)

type Client struct {
	host       string
	keyName    string
	privateKey *ecdsa.PrivateKey
	httpClient *http.Client
}

func NewClient(appEnv string, s *secret.Secret) (*Client, error) {
	host := hostForEnv(appEnv)

	privateKey, err := parsePrivateKey(s.PrivateKey)
	if err != nil {
		return nil, err
	}

	return &Client{
		host:       host,
		keyName:    s.Name,
		privateKey: privateKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}, nil
}

func hostForEnv(appEnv string) string {
	env := strings.ToLower(appEnv)

	switch env {
	case "prod", "production":
		return "api.coinbase.com"
	default:
		return "api-sandbox.coinbase.com"
	}
}
