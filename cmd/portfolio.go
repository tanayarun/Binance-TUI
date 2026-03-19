// Package cmd is used to get data from api
package cmd

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/tanayarun/Binance-TUI/internal/binance"
	"github.com/tanayarun/Binance-TUI/internal/ui"
)

var portfolioCmd = &cobra.Command{
	Use:   "portfolio",
	Short: "see your portfolio",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := os.Getenv("BINANCE_API_KEY")
		secretKey := os.Getenv("BINANCE_SECRET_KEY")
		client := binance.NewClient(apiKey, secretKey)

		m := ui.NewPortfolioModel(client)
		p := tea.NewProgram(m, tea.WithAltScreen())
		_, err := p.Run()

		return err
	},
}
