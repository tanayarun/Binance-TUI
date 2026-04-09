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

func sign(query, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(query))
	signature := hex.EncodeToString(mac.Sum(nil))
	return signature
}

func (c *Client) GetBalances() ([]Balance, error) {
	timestamp := time.Now().UnixMilli()
	queryString := fmt.Sprintf("timestamp=%d", timestamp)

	signature := sign(queryString, c.secretKey)
	url := fmt.Sprintf("https://demo-api.binance.com/api/v3/account?timestamp=%d&signature=%s", timestamp, signature)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-MBX-APIKEY", c.apiKey)
	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var account AccountResponse
	err = json.Unmarshal(body, &account)
	if err != nil {
		return nil, err
	}

	var result []Balance
	for _, b := range account.Balances {
		free, err := strconv.ParseFloat(b.Free, 64)
		if err != nil {
			continue
		}
		if free > 0 {
			result = append(result, b)
		}
	}

	return result, nil
}
