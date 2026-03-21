package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
)

type DepthUpdate struct {
	Symbol string     `json:"s"`
	Bids   [][]string `json:"b"`
	Asks   [][]string `json:"a"`
}

func ConnectOrderbook(symbol string) (*websocket.Conn, error) {
	url := fmt.Sprintf("wss://demo-stream.binance.com:9443/ws/%s@depth", strings.ToLower(symbol))
	conn, _, err := websocket.DefaultDialer.DialContext(context.Background(), url, nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func ReadOrderbook(conn *websocket.Conn) error {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			return err
		}

		var update DepthUpdate
		err = json.Unmarshal(p, &update)
		if err != nil {
			return err
		}

		if len(update.Bids) > 0 && len(update.Asks) > 0 {
			fmt.Printf("Best Bid: %s | Best Ask: %s\n", update.Bids[0][0], update.Asks[0][0])
		}
	}
}
