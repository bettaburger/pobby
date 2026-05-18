package tui

import (
	"github.com/bettaburger/pobby/internal/port"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"fmt"
	"time"
)

type ProcessTableData struct {
	tview.TableContentReadOnly
	TableHeaders []string
	Content []port.ProcessStat
	Length int
	RowMap map[int32]int
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
	process := p.Content[row-1] // content of a process at a specific row and column
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

// this function displays the 'other' primitive 
func DisplayOther() *tview.TextView {
	other := tview.NewTextView().SetText(fmt.Sprintf("Total PID: %v", port.GetTotalCon()))
	other.SetTitle(" Other ").SetBorder(true).SetTitleAlign(tview.AlignLeft)
	other.SetBackgroundColor(tcell.NewRGBColor(30,33,36))
	return other
}

// this function updates the table if there is a new row added, if a row exists keep it 
func UpdateTable(app *tview.Application, table *tview.Table, data *ProcessTableData) {
	app.QueueUpdateDraw(func() {
		newMap := make(map[int32]int)
		row := 1 // start at row 1 
		for _, p := range data.Content {
			existingRow, ok := data.RowMap[p.PID]
			if ok {
				// update existing row
				table.GetCell(existingRow, 0).SetText(fmt.Sprintf("%d", p.PID))
				table.GetCell(existingRow, 1).SetText(p.ProcessName)
				table.GetCell(existingRow, 2).SetText(fmt.Sprintf("%d", p.Port))
				table.GetCell(existingRow, 3).SetText(p.Proto)
				table.GetCell(existingRow, 4).SetText(fmt.Sprintf("%s:%d", p.LocalAddress.IP, p.LocalAddress.Port))
				table.GetCell(existingRow, 5).SetText(fmt.Sprintf("%s:%d", p.ForeignAddress.IP, p.ForeignAddress.Port))
				table.GetCell(existingRow, 6).SetText(p.State)
				newMap[p.PID] = existingRow
			} else {
				// create new row
				table.SetCell(row, 0, tview.NewTableCell(fmt.Sprintf("%d", p.PID)))
				table.SetCell(row, 1, tview.NewTableCell(p.ProcessName))
				table.SetCell(row, 2, tview.NewTableCell(fmt.Sprintf("%d", p.Port)))
				table.SetCell(row, 3, tview.NewTableCell(p.Proto))
				table.SetCell(row, 4, tview.NewTableCell(fmt.Sprintf("%s:%d", p.LocalAddress.IP, p.LocalAddress.Port)))
				table.SetCell(row, 5, tview.NewTableCell(fmt.Sprintf("%s:%d", p.ForeignAddress.IP, p.ForeignAddress.Port)))
				table.SetCell(row, 6, tview.NewTableCell(p.State))
				newMap[p.PID] = row
				row++
			}
		}
	})
}

// this function creates the ticker, interval of 1 second per tick
func Refresh(app *tview.Application, table *tview.Table, data *ProcessTableData) {
	tick := time.NewTicker(1 * time.Second) 
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			data.Content = port.StoreCon()
			UpdateTable(app, table, data)
			//currentTime := time.Now().Format("15:04:05") // debugging the queueupdatedraw
				//table.SetTitle(fmt.Sprintf(" Processes %s ", currentTime))
		}
	}		
}

// this function displays the 'process' primitive
func ProcessTable(app *tview.Application) *tview.Table {
	// form the connection 
	processes := port.StoreCon()
	// table
	table := tview.NewTable()
	table.SetBorder(true).SetTitleAlign(tview.AlignCenter).SetBackgroundColor(tcell.NewRGBColor(30, 33, 36))
	table.SetFixed(1, 0) 
	table.SetTitle(" Processes ")
	tableHeaders := []string{"pid", "port", "proto", "process name", "local address", "foreign address", "state"}
	data := &ProcessTableData{
		TableHeaders: tableHeaders,
		Content: processes,
		Length: len(processes),
	}
	table.SetContent(data)
	table.SetSelectable(true, false)
	// call refresh
	go Refresh(app, table, data)
	return table
}

