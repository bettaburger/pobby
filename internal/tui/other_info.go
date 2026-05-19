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

// this function displays the process information primitive
func DisplayInfo() *tview.TextView {
	pid := GetPIDInput(FilterForm())
	//_, proto := GetPortInput(FilterForm()) // conversion error 
	//_, state := GetStateInput(FilterForm())
	
	information := tview.NewTextView()
	information.SetTitle(" Information ").SetBorder(true).SetTitleAlign(tview.AlignLeft)
	information.SetBackgroundColor(tcell.NewRGBColor(30,33,36))

	information.SetText(fmt.Sprintf("pid number: %v", pid))
	//information.SetText(fmt.Sprintf("port number: %v", proto))
	//information.SetText(fmt.Sprintf("state: %v", state))
	return information
}