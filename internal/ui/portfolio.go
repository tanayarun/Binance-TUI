package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tanayarun/Binance-TUI/internal/binance"
)

type Model struct {
	balances []binance.Balance
}

func NewPortfolioModel() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (m Model) View() string {
	return ""
}
