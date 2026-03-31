package command

import (
	"bytes"
	"os/exec"
	//"fmt"

	"github.com/spf13/cobra"
	"charm.land/bubbles/v2/table"
	"github.com/bettaburger/pobby/internal/tui"

)

var list string

// shows all listening ports
// run task list
var listCmd = &cobra.Command {
	Use: "list", 
	Short: "list listening ports",
	RunE: func(cmd *cobra.Command, args []string) error {
		// run command to list ports macos
		output := exec.Command("lsof", "-i" ,"-P" ,"-n")
	
		out, err := output.Output()
		if err != nil {
			return err
		}

		str := bytes.Split(out, []byte("\n"))
		var rows []table.Row
		
		for _, s := range str {
			if bytes.Contains(s, []byte("LISTEN")) {
				fields := bytes.Fields(s)
				if len(fields) < 9 {
					continue
				}
				// command, pid, user, group id, file descriptor, 
				// rest of info can be seen using select + enter
				command := string(fields[0]) 	// cmd
				pid := string(fields[1])	// pid
				user := string(fields[2])	// user id
				port := string(fields[8])	// port
				
				rows = append(rows, table.Row{command, pid, user, port})
			}
		}
		//fmt.Println(rows)
		return tui.StartTable(rows)
	},
}


func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&list, "list", "l", "LIST", "list all listening ports")
}