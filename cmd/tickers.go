package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tanayarun/Binance-TUI/internal/binance"
)

var tickersCmd = &cobra.Command{
	Use:   "tickers",
	Short: "Get the top tickers",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := binance.ConnectTickers()
		if err != nil {
			return err
		}

		err = binance.ReadTickers(conn)
		if err != nil {
			return err
		}

		return nil
	},
}
