package tui

import (
	"fmt"
	"github.com/bettaburger/pobby/internal/port"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// this function displays the 'other' primitive 
func DisplayOther() *tview.TextView {
	other := tview.NewTextView().SetText(fmt.Sprintf("Total PID: %v", port.GetTotalCon()))
	other.SetTitle(" Other ").SetBorder(true).SetTitleAlign(tview.AlignLeft)
	other.SetBackgroundColor(tcell.NewRGBColor(30,33,36))
	return other
}
