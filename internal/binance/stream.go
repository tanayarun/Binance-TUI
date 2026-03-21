package binance

import (
	"context"
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
)

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

		fmt.Println(string(p))
	}
}
