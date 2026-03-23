package ui

import (
	"encoding/json"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
	"github.com/tanayarun/Binance-TUI/internal/binance"
)

type OrderbookModel struct {
	symbol  string
	conn    *websocket.Conn
	bids    [][]string
	asks    [][]string
	loading bool
	err     error
}

type orderbookUpdateMsg struct {
	bids [][]string
	asks [][]string
}

func NewOrderbookModel(symbol string, conn *websocket.Conn) OrderbookModel {
	return OrderbookModel{
		symbol:  symbol,
		conn:    conn,
		loading: true,
	}
}

func (m OrderbookModel) Init() tea.Cmd {
	return m.fetchOrderbook()
}

func (m OrderbookModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case orderbookUpdateMsg:
		m.bids = msg.bids
		m.asks = msg.asks
		m.loading = false
		return m, m.fetchOrderbook()
	case errMsg:
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m OrderbookModel) View() string {
	return ""
}

func (m OrderbookModel) fetchOrderbook() tea.Cmd {
	return func() tea.Msg {
		_, p, err := m.conn.ReadMessage()
		if err != nil {
			return errMsg{err: err}
		}

		var update binance.DepthUpdate
		err = json.Unmarshal(p, &update)
		if err != nil {
			return errMsg{err: err}
		}

		return orderbookUpdateMsg{
			bids: update.Bids,
			asks: update.Asks,
		}
	}
}
