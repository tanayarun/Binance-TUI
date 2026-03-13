// Package binance is used to establish connection with binance
package binance

import (
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
	return nil
}
