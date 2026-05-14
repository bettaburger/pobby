// list open files/processes
package port

import (
	"fmt"
	gopsutil "github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
)

// this type describes each process or connection
type ProcessStat struct {
	PID int32
	Port uint32
	Proto string
	ProcessName string 
	LocalAddress gopsutil.Addr
	ForeignAddress gopsutil.Addr
	State string 
}

var (
	totalConnections uint16
	//connections []gopsutil.ConnectionStat // stores all connections ; make sure to reset the list or make it dynamic
	listening []gopsutil.ConnectionStat // stores listening connections
	processes []ProcessStat // list of all process 
)

func GetTotalCon() uint16 { return totalConnections }

// function that stores all the connections
func StoreCon() []ProcessStat {
	connections, err := gopsutil.Connections("inet")
	if err != nil {
		fmt.Printf("./internal/port/lsof.go - unable to list open connections: %v\n", err)
	}
	for _, con := range connections {
		name := ""
		// get process name from PID
		if con.Pid != 0 {
			proc, err := process.NewProcess(con.Pid)
			if err == nil {
				pname, err := proc.Name()
				if err == nil {
					name = pname
				}
			}
		}
		p := ProcessStat{
			PID: con.Pid, 
			Port: con.Laddr.Port,
			Proto: GetProto(con.Type),
			ProcessName: name,
			LocalAddress: con.Laddr,
			ForeignAddress: con.Raddr,
			State: con.Status, 
		}
		processes = append(processes, p)
		totalConnections += 1
	}
	return processes
}

// print the connection list
func PrintConnections(CList []gopsutil.ConnectionStat) {
	for _, c := range CList {
		fmt.Printf("PID=%-6d %-12s %-22s -> %-22s %s\n", c.Pid, GetProto(c.Type), Addr(c.Laddr.IP, c.Laddr.Port), Addr(c.Raddr.IP, c.Raddr.Port), c.Status,)
	}
}
// since gopsutil doesn't display the proto name 
var netProtocols = []string{
	"ip",
	"icmp",
	"icmpmsg",
	"tcp",
	"udp",
	"udplite",
}
// this function returns the string representation of the proto uint
func GetProto(pt uint32) string {
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

// this function lists all established connections


