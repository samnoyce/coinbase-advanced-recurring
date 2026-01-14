package coinbase

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
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

func (c *Client) request(ctx context.Context, method, path string, request, response any) error {
	url := fmt.Sprintf("https://%s%s", c.host, path)

	token, err := createToken(method, c.host, path, c.keyName, c.privateKey)
	if err != nil {
		return fmt.Errorf("creating token: %w", err)
	}

	var body io.Reader
	if request != nil {
		b, err := json.Marshal(request)
		if err != nil {
			return fmt.Errorf("encoding request: %w", err)
		}
		body = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return fmt.Errorf("building request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}

	return nil
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
