package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/tanayarun/Binance-TUI/internal/binance"
	"github.com/tanayarun/Binance-TUI/internal/ui"
)

var tickersCmd = &cobra.Command{
	Use:   "tickers",
	Short: "Get the top tickers",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := binance.ConnectTickers()
		if err != nil {
			return err
		}

		m := ui.NewTickersModel(conn)
		p := tea.NewProgram(m, tea.WithAltScreen())
		_, err = p.Run()
		if err != nil {
			return err
		}

		return nil
	},
}
