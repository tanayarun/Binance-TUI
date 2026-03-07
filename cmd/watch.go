package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "To wathch the orderbook of symbol",
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol, err := cmd.Flags().GetString("symbol")
		if err != nil {
			return err
		}
		fmt.Println(symbol)
		return nil
	},
}

func init() {
	watchCmd.Flags().StringP("symbol", "s", "", "Trading pair e.g. BTCUSDT")
}
