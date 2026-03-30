package command


import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

// descriptors
var rootCmd = &cobra.Command {
	Use: "pobby",
	Short: "Pobby is a CLI tool for ports.",
	Long: "Simple CLI tool to display and manage listening ports in your terminal",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}