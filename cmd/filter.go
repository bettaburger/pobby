package cmd 

import (
	"github.com/spf13/cobra"
	"github.com/bettaburger/pobby/internal/tui"
)
var (
	filter string
)
// This function describes the filter command
var FilterCmd = &cobra.Command {
	Use: "filter",
	Short: "displays a filter for finding specific processes",
	RunE: func(cmd *cobra.Command, args []string) error {
		tui.Run()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(FilterCmd)
	FilterCmd.Flags().StringVarP(&filter, "filter", "f", "", "opens the filter form" ) // ./pobby filter
}


