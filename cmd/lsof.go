package cmd

import (
	"github.com/spf13/cobra"
	"github.com/bettaburger/pobby/internal/port"
	
)
var (
	listC bool
)
// This function describes the ports command
var PortsCmd = &cobra.Command {
	Use: "ports",
	Short: "List all open files/processes",
	RunE: func(cmd *cobra.Command, args []string) error {
		port.ListConnection()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(PortsCmd)
	PortsCmd.Flags().BoolVarP(&listC, "list", "l", false, "list all open files/processes" ) // ./pobby ports -l
}

