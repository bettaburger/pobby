package tui

import (
	"github.com/bettaburger/pobby/internal/port"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"fmt"
)

type ProcessTableData struct {
	tview.TableContentReadOnly
	TableHeaders []string
	Content []port.ProcessStat
}

// return the row of the table
func (p *ProcessTableData) GetRowCount() int {
	return 1 + len(p.Content)
}

// return the col of the table, basically the length of table headers 
func (p *ProcessTableData) GetColumnCount() int {
	return len(p.TableHeaders)
}

// returns the cell display
func (p *ProcessTableData) GetCell(row, column int) *tview.TableCell {
	if row == 0 {
		return tview.NewTableCell(p.TableHeaders[column]).SetSelectable(false).SetAttributes(tcell.AttrBold).SetBackgroundColor(tcell.NewRGBColor(35,39,42))
	}
	// process rows for column 0 ... 6 
	process := p.Content[row-1]
	var text string
	switch column {
	case 0:
		text = fmt.Sprintf("%d",process.PID)
	case 1:
		text = fmt.Sprintf("%d", process.Port)
	case 2:
		text = process.Proto
	case 3:
		text = process.ProcessName
	case 4:
		text = fmt.Sprintf("%s:%d", process.LocalAddress.IP, process.LocalAddress.Port,)
	case 5:
		text = fmt.Sprintf("%s:%d", process.ForeignAddress.IP, process.ForeignAddress.Port)
	case 6:
		text = process.State
	}
	return tview.NewTableCell(text)
}

func ProcessTable() *tview.Table {
	// form the connection 
	processes := port.StoreCon()
	// table
	table := tview.NewTable()
	table.SetTitle(" Processes ").SetBorder(true).SetTitleAlign(tview.AlignCenter).SetBackgroundColor(tcell.NewRGBColor(30, 33, 36))
	tableHeaders := []string{"pid", "port", "proto", "process name", "local address", "foreign address", "state"}
	data := &ProcessTableData{
		TableHeaders: tableHeaders,
		Content: processes,
	}
	table.SetContent(data)
	table.SetFixed(1, 0)
	table.SetSelectable(true, false)
	return table
}

// refresh in the background