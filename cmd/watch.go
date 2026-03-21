package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tanayarun/Binance-TUI/internal/binance"
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
		err = binance.ReadOrderbook(conn)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	watchCmd.Flags().StringP("symbol", "s", "", "Trading pair e.g. BTCUSDT")
}
