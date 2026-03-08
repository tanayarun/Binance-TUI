package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tickersCmd = &cobra.Command{
	Use: "tickers",
	Short: "Get the top tickers",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("tickers coming soon")
		return nil
	},
}
