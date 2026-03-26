package ui

import (
	"encoding/json"
	"fmt"
	"strings"

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
	if m.loading {
		return "Conncting to " + m.symbol + "..."
	}
	if m.err != nil {
		return "Error: " + m.err.Error()
	}

	var sb strings.Builder

	sb.WriteString(StyleHeader.Render("Orderbook:") + "\n")

	for i := range 5 {
		if i < len(m.asks) {
			row := fmt.Sprintf("Ask: %s | %s", m.asks[i][0], m.asks[i][1])
			sb.WriteString(StyleRed.Render(row) + "\n")
		}
	}

	sb.WriteString(StyleGray.Render("--------------") + "\n")

	for i := range 5 {
		if i < len(m.bids) {
			column := fmt.Sprintf("Bids: %s | %s", m.bids[i][0], m.bids[i][1])
			sb.WriteString(StyleGreen.Render(column) + "\n")
		}
	}

	sb.WriteString("\n" + StyleGray.Render("q: quit"))

	return sb.String()
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
