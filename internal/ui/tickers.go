package ui

import (
	"encoding/json"

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
	return m, nil
}

func (m TickersModel) View() string {
	return ""
}
