package tui

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2" 
)

// this function describes filterform primitive
func FilterForm() *tview.Form {
	form := tview.NewForm().
	AddInputField("pid number", "", 40, tview.InputFieldInteger, nil).
	AddInputField("port number", "", 40, tview.InputFieldInteger, nil).
	AddDropDown("proto", []string{"any","tcp", "udp", "icmp", "ip", "udplite", "icmpmsg"}, 0, nil). // available for gopsutil ip,icmp,icmpmsg,tcp,udp,udplite
	AddInputField("local address", "", 40, nil, nil).
	AddInputField("foreign address", "", 40, nil, nil).
	AddDropDown("state", []string{"listening", "established"}, 0, nil).
	AddButton("Filter", nil)
	form.SetTitle(" Filter ").SetBorder(true).SetTitleAlign(tview.AlignLeft)
	form.SetFieldBackgroundColor(tcell.NewRGBColor(35,39,42))
	form.SetBackgroundColor(tcell.NewRGBColor(30,33,36)) 
	form.SetLabelColor(tcell.NewRGBColor(153,170,181)) // labels
	form.SetButtonBackgroundColor(tcell.NewRGBColor(153,170,181))// filter button
	form.SetButtonActivatedStyle(tcell.StyleDefault.Background(tcell.NewRGBColor(116,132,142)))
	form.SetBorder(true)

	return form 
}
