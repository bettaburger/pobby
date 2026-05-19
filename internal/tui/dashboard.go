package tui

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2" 
)

// general layout of the tui
func Dashboard() {
	dashboard := tview.NewApplication()
	dashboard.EnableMouse(true)

	/*newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(text)
	}*/
	// layouts
	header := tview.NewTextView().SetTitle(" Pobby Dashboard ").SetBorder(true).SetTitleAlign(tview.AlignCenter).SetBackgroundColor(tcell.NewRGBColor(30,33,36))
	filter := FilterForm()
	processes := ProcessTable(dashboard)
	information := DisplayInfo()
	help := tview.NewTextView().SetText("Navigation - Up, Down, press Ctrl-C to exit")
	help.SetTitle(" Help ").SetBorder(true).SetTitleAlign(tview.AlignCenter)
	help.SetTextColor((tcell.NewRGBColor(153,170,181)))
	help.SetTitleColor((tcell.NewRGBColor(153,170,181)))
	help.SetBorderColor((tcell.NewRGBColor(153,170,181)))
	help.SetBackgroundColor(tcell.NewRGBColor(30,33,36))
	// other layout
	other := DisplayOther() // quick search? kill pid?
	// actual grid layout
	grid := tview.NewGrid().
	SetRows(3, 20, 0, 3).SetColumns(30, 30, 0, 30).
	SetBorders(false).
	AddItem(header, 0, 0, 1, 4, 0, 0, false). // span 4 cols
	AddItem(help, 3, 0, 1, 4, 0, 0, false) // span 4 cols

	// Layout for screens wider than 100 cells.
	grid.AddItem(filter, 1, 0, 1, 2, 0, 100, true).
	AddItem(information, 1, 2, 1, 1, 0, 100, false).
	AddItem(other, 1, 3, 1, 1, 0, 100, false).
	AddItem(processes, 2, 0, 1, 4, 0, 100, false)

	// Layout for screens narrower than 100 cells (are hidden).
	grid.AddItem(filter, 0, 0, 0, 0, 0, 0, false).
	AddItem(processes, 2, 0, 1, 5, 0, 0, false).
	AddItem(information, 1, 0, 1, 5, 0, 0, false).
	AddItem(other, 0, 0, 0, 0, 0, 0, false)

	dashboard.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'k':
		return nil // deactivate key
	case 'j':
		return nil
	}
	return event
	})

	//go Refresh(dashboard)

	// run the app
	if err := dashboard.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		fmt.Printf("tui/dashboard.go - unable to load application, %\n", err)
		panic(err)
	}
}


