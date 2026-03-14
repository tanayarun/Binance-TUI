// Package binance is used to establish connection with binance
package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
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

	mac := hmac.New(sha256.New, []byte(c.secretKey))
	mac.Write([]byte(queryString))
	signature := hex.EncodeToString(mac.Sum(nil))

	url := fmt.Sprintf("https://demo-api.binance.com/api/v3/account?timestamp=%d&signature=%s", timestamp, signature)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-MBX-APIKEY", c.apiKey)
	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var account AccountResponse
	err = json.Unmarshal(body, &account)
	if err != nil {
		return err
	}

	for _, b := range account.Balances {
		free, err := strconv.ParseFloat(b.Free, 64)
		if err != nil {
			continue
		}
		if free > 0 {
			fmt.Printf("Asset: %s | Free: %s | Locked: %s\n", b.Asset, b.Free, b.Locked)
		}
	}

	return nil
}
