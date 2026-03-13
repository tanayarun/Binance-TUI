// Package binance is used to establish connection with binance
package binance

type Client struct {
	apiKey    string
	secretKey string
}

func NewClient(apiKey, secretKey string) *Client {
	return &Client{
		apiKey : apiKey,
	secretKey : secretKey,
	}
}
