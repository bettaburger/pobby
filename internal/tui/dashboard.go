package tui

import (
	"fmt"
	"github.com/rivo/tview"
)

// general layout of the tui
func Dashboard() {
	/*newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(text)
	}*/
	// layouts
	header := tview.NewTextView().SetTitle("Pobby Dashboard").SetBorder(true).SetTitleAlign(tview.AlignCenter)
	filter := tview.NewTextArea().SetTitle("Filter").SetBorder(true).SetTitleAlign(tview.AlignLeft)
	processes := tview.NewTextView().SetTitle("Processes").SetBorder(true).SetTitleAlign(tview.AlignLeft)
	information := tview.NewTextView().SetTitle("Information").SetBorder(true).SetTitleAlign(tview.AlignLeft)
	help := tview.NewTextArea().SetTitle("Help").SetBorder(true).SetTitleAlign(tview.AlignLeft)
	// actual grid layout
	grid := tview.NewGrid().
	SetRows(3, 0, 0, 3).SetColumns(30, 0, 0, 30).
	SetBorders(false).
	AddItem(header, 0, 0, 1, 5, 0, 0, false). // span 4 cols
	AddItem(help, 3, 0, 1, 5, 0, 0, false) // span 4 cols
	
	// Layout for screens wider than 100 cells.
	grid.AddItem(filter, 1, 0, 1, 3, 0, 100, false).
	AddItem(information, 1, 3, 1, 2, 0, 100, false).
	AddItem(processes, 2, 0, 1, 5, 0, 100, false)

	// Layout for screens narrower than 100 cells 
	grid.AddItem(filter, 0, 0, 0, 0, 0, 0, false).
		AddItem(processes, 2, 0, 1, 5, 0, 0, false).
		AddItem(information, 1, 0, 1, 5, 0, 0, false)

	if err := tview.NewApplication().SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		fmt.Printf("tui/dashboard.go - unable to load application, %\n", err)
		panic(err)
	}
}
