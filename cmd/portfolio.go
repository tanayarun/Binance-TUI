// Package cmd is used to get data from api
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var portfolioCmd = &cobra.Command{
	Use: "portfolio",
	Short: "see your portfolio",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Portfolio coming soon")
		return nil
	},
}
