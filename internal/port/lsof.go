// list open files/processes
package port

import (
	"fmt"
	gopsutil "github.com/shirou/gopsutil/v4/net"

)

var (
	totalConnections uint16
	connections []gopsutil.ConnectionStat // stores all connections ; make sure to reset the list or make it dynamic
	listening []gopsutil.ConnectionStat // stores listening connections
)

func GetTotalCon() uint16 { return totalConnections }

// function that stores all the connections
func StoreCon() {
	connections, err := gopsutil.Connections("inet")
	if err != nil {
		fmt.Printf("./internal/port/lsof.go - unable to list open connections: %v\n", err)
	}
	for _, con := range connections {
		connections = append(connections, con)
		totalConnections += 1
	}
}

// print the connection list
func PrintConnections(CList []gopsutil.ConnectionStat) {
	for _, c := range CList {
		fmt.Printf("PID=%-6d %-12s %-22s -> %-22s %s\n", c.Pid, Prototype(c.Type), Addr(c.Laddr.IP, c.Laddr.Port), Addr(c.Raddr.IP, c.Raddr.Port), c.Status,)
	}
}

func Prototype(pt uint32) string {
	switch pt {
	case 1:
		return "TCP"
	case 2:
		return "UDP"
	default:
		return "UNKNOWN"
	}
}

func Addr(ip string, port uint32) string {
	if ip == "" {
		ip = "*"
	}
	return fmt.Sprintf("%s:%d", ip, port)
}

// this function lists all open connections 
func ListConnection() {
	connections, err := gopsutil.Connections("inet")
	if err != nil {
		fmt.Printf("./internal/port/lsof.go - unable to list open connections: %v\n", err)
	}
	PrintConnections(connections)
}

// this function lists all listening connections 
func ListListeningCon() {
	connections, err := gopsutil.Connections("inet")
	if err != nil {
		fmt.Printf("./internal/port/lsof.go - unable to list open listening connections: %v\n", err)
	}
	for _, c := range connections {
		if c.Status == "LISTEN" {
			listening = append(listening, c)
		}
	}
	PrintConnections(listening)
}
