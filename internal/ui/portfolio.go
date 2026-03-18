package ui

import (
	"fmt"
	"strings"

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
	switch msg := msg.(type) {
	case balancesLoadedMsg:
		m.balances = msg.balances
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

func (m Model) View() string {
	if len(m.balances) == 0 {
		return "Loading..."
	}
	var sb strings.Builder
	for _, b := range m.balances {
		fmt.Fprintf(&sb, "Asset: %s | Free: %s | Locked: %s\n", b.Asset, b.Free, b.Locked)
	}
	return sb.String()
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
