// Package cmd is used to get data from api
package cmd

import "github.com/spf13/cobra"

var portfolioCmd = &cobra.Command{
	Use: "portfolio",
	Short: "see your portfolio",
}

func init() {
}
