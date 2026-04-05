package ui

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
	"github.com/tanayarun/Binance-TUI/internal/binance"
)

type TickersModel struct {
	tickers []binance.MiniTicker
	conn    *websocket.Conn
	loading bool
	err     error
}

type tickerUpdateMsg struct {
	tickers []binance.MiniTicker
}

func NewTickersModel(conn *websocket.Conn) TickersModel {
	return TickersModel{
		conn:    conn,
		loading: true,
	}
}

func (m TickersModel) fetchTickers() tea.Cmd {
	return func() tea.Msg {
		_, p, err := m.conn.ReadMessage()
		if err != nil {
			return errMsg{err: err}
		}

		var update []binance.MiniTicker
		err = json.Unmarshal(p, &update)
		if err != nil {
			return errMsg{err: err}
		}

		return tickerUpdateMsg{
			tickers: update,
		}
	}
}

func (m TickersModel) Init() tea.Cmd {
	return m.fetchTickers()
}

func (m TickersModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickerUpdateMsg:
		m.tickers = msg.tickers
		m.loading = false
		return m, m.fetchTickers()
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

func (m TickersModel) View() string {
	if m.loading {
		return "Fetching Tickers"
	}
	if m.err != nil {
		return "Error: " + m.err.Error()
	}

	var sb strings.Builder

	sb.WriteString(StyleHeader.Render("Tickers: ") + "\n")

	watchlist := map[string]bool{
		"BTCUSDT": true,
		"ETHUSDT": true,
		"SOLUSDT": true,
		"BNBUSDT": true,
		"XRPUSDT": true,
	}

	for _, ticker := range m.tickers {
		if watchlist[ticker.Symbol] {
			current, _ := strconv.ParseFloat(ticker.Current, 64)
			open, _ := strconv.ParseFloat(ticker.Open, 64)
			change := ((current - open) / open) * 100
			row := fmt.Sprintf("%-12s %15s %+.2f%%", ticker.Symbol, ticker.Current, change)
			sb.WriteString(PriceColor(change).Render(row) + "\n")
		}
	}

	return sb.String()
}
