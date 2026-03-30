package command

import (
	"bytes"
	"os/exec"
	"fmt"

	"github.com/spf13/cobra"
)

//var cmd *exec.Cmd

var list string

// shows all listening ports
var listCmd = &cobra.Command {
	Use: "list", 
	Short: "list listening ports",
	RunE: func(cmd *cobra.Command, args []string) error {
		// run command to list ports macos, linux
		output := exec.Command("lsof", "-i" ,"-P" ,"-n")
	
		out, err := output.Output()
		if err != nil {
			return err
		}

		str := bytes.Split(out, []byte("\n"))
		for _, s := range str {
			if bytes.Contains(s, []byte("LISTEN")) {
				fmt.Println(string(s))
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	// arguments
	listCmd.Flags().StringVarP(&list, "list", "l", "POSTS", "list ports")
}