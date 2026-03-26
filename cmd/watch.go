package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/tanayarun/Binance-TUI/internal/binance"
	"github.com/tanayarun/Binance-TUI/internal/ui"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "To wathch the orderbook of symbol",
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol, err := cmd.Flags().GetString("symbol")
		if err != nil {
			return err
		}

		conn, err := binance.ConnectOrderbook(symbol)
		if err != nil {
			return err
		}

		m := ui.NewOrderbookModel(symbol, conn)
		p := tea.NewProgram(m, tea.WithAltScreen())
		_, err = p.Run()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	watchCmd.Flags().StringP("symbol", "s", "", "Trading pair e.g. BTCUSDT")
}
