package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tanayarun/Binance-TUI/internal/binance"
)

type Model struct {
	client   *binance.Client
	balances []binance.Balance
}

type balancesLoadedMsg struct {
	balances []binance.Balance
}

type errMsg struct {
	err error
}

func NewPortfolioModel(client *binance.Client) Model {
	return Model{client: client}
}

func (m Model) Init() tea.Cmd {
	return m.fetchBalances()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return ""
}

func (m Model) fetchBalances() tea.Cmd {
	return func() tea.Msg {
		result, err := m.client.GetBalances()
		if err != nil {
			return errMsg{err: err}
		}

		return balancesLoadedMsg{balances: result}
	}
}
