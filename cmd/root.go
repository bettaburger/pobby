package cmd


import (
	"fmt"
	"os"
	"github.com/spf13/cobra"

)

// descriptor
var rootCmd = &cobra.Command {
	Use: "pobby",
	Short: "Pobby is a tui for managing ports",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}