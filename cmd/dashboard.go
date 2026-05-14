package cmd 

import (
	"github.com/spf13/cobra"
	"github.com/bettaburger/pobby/internal/tui"
)
var (
	open string
)
// This function describes the open command
var OpenCmd = &cobra.Command {
	Use: "open",
	Short: "starts the dashboard application",
	RunE: func(cmd *cobra.Command, args []string) error {
		tui.Dashboard()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(OpenCmd)
	OpenCmd.Flags().StringVarP(&open, "open", "", "", "starts the dashboard" ) // ./pobby open
}


