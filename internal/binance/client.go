// Package binance is used to establish connection with binance
package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	apiKey    string
	secretKey string
}

type Balance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

type AccountResponse struct {
	Balances []Balance `json:"balances"`
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

	url := fmt.Sprintf("https://demo-api.binance.com/api/v3/account?timestamp=%d&signature=%s", timestamp, signature)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req.Header.Set("X-MBX-APIKEY", c.apiKey)
	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	fmt.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}
