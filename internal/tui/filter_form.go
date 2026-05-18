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
	AddButton("Filter", nil).
	AddButton("Clear", nil)
	form.SetTitle(" Filter ").SetBorder(true).SetTitleAlign(tview.AlignLeft)
	form.SetFieldBackgroundColor(tcell.NewRGBColor(35,39,42))
	form.SetBackgroundColor(tcell.NewRGBColor(30,33,36)) 
	form.SetLabelColor(tcell.NewRGBColor(153,170,181)) // labels
	form.SetBorder(true)

	// buttons
	form.SetButtonBackgroundColor(tcell.NewRGBColor(153,170,181))
	form.SetButtonActivatedStyle(tcell.StyleDefault.Background(tcell.NewRGBColor(116,132,142)))

	return form 
}

func GetPIDInput(form *tview.Form) string {
	pid := form.GetFormItemByLabel("pid number").(*tview.InputField)
	return pid.GetText()
}

// returns the current selected option by its index and the port type name
func GetPortInput(form *tview.Form) (int, string) {
	port := form.GetFormItemByLabel("port number").(*tview.DropDown) // converts to tview state 
	return port.GetCurrentOption()
}

func GetProtoInput(form *tview.Form) string {
	proto := form.GetFormItemByLabel("proto").(*tview.InputField)
	return proto.GetText()
}

func GetLaddrInput(form *tview.Form) string {
	laddr := form.GetFormItemByLabel("local address").(*tview.InputField)
	return laddr.GetText()
}

func GetFaddrInput(form *tview.Form) string {
	faddr := form.GetFormItemByLabel("foreign address").(*tview.InputField)
	return faddr.GetText()
}

// returns the current selected option by its index and the state name
func GetStateInput(form *tview.Form) (int, string) {
	state := form.GetFormItemByLabel("state").(*tview.DropDown)
	return state.GetCurrentOption()
}
