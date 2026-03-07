package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "binance-tui",
	Short: "A terminal UI for binance",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(watchCmd)
  rootCmd.AddCommand(portfolioCmd)
	rootCmd.AddCommand(tickersCmd)
}
