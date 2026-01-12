package coinbase

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func parsePrivateKey(privateKey string) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, fmt.Errorf("private key must be PEM-encoded")
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parsing EC private key: %w", err)
	}

	return key, nil
}

func createToken(method, host, path, keyName string, privateKey *ecdsa.PrivateKey) (string, error) {
	now := time.Now()

	claims := jwt.MapClaims{
		"sub": keyName,
		"iss": "coinbase-cloud",
		"nbf": now.Unix(),
		"exp": now.Add(2 * time.Minute).Unix(),
		"uri": fmt.Sprintf("%s %s%s", method, host, path),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header["kid"] = keyName
	token.Header["nonce"] = uuid.NewString()

	signed, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("signing token: %w", err)
	}

	return signed, nil
}
