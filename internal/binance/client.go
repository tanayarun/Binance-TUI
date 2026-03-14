// Package binance is used to establish connection with binance
package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Client struct {
	apiKey    string
	secretKey string
}

func NewClient(apiKey, secretKey string) *Client {
	return &Client{
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

func (c *Client) GetBalances() error {
	timestamp := time.Now().UnixMilli()
	queryString := fmt.Sprintf("timestamp=%d", timestamp)
	fmt.Println(queryString)

	mac := hmac.New(sha256.New, []byte(c.secretKey))
	mac.Write([]byte(queryString))
	signature := hex.EncodeToString(mac.Sum(nil))
	fmt.Println(signature)

	return nil
}
