package tui

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2" 
	"fmt"
)

type ProcessFilter struct {
	PID            string
	Port           string
	Proto          string
	LocalAddress   string
	ForeignAddress string
	State          string
}

// this function describes filterform primitive
func FilterForm(app *tview.Application, table *tview.Table, data *ProcessTableData) *tview.Form {
	form := tview.NewForm()
	pidField := tview.NewInputField().SetLabel("pid").SetFieldWidth(40).SetAcceptanceFunc(tview.InputFieldInteger)
	portField := tview.NewInputField().SetLabel("port").SetFieldWidth(40).SetAcceptanceFunc(tview.InputFieldInteger)
	protoDrop := tview.NewDropDown().SetLabel("proto").SetOptions([]string{"any", "tcp", "udp", "icmp", "ip", "udplite", "icmpmsg"}, nil)
	localField := tview.NewInputField().SetLabel("local address").SetFieldWidth(40)
	foreignField := tview.NewInputField().SetLabel("foreign address").SetFieldWidth(40)
	stateDrop := tview.NewDropDown().SetLabel("state").SetOptions([]string{"all", "listening", "established"}, nil)
	form.AddFormItem(pidField).AddFormItem(portField).AddFormItem(protoDrop).AddFormItem(localField).AddFormItem(foreignField).AddFormItem(stateDrop)
	// this button filters the "query"
	form.AddButton("Filter", func() {
		data.mu.Lock()
		_, proto := protoDrop.GetCurrentOption()
		_, state := stateDrop.GetCurrentOption()
		filter := ProcessFilter{
			PID: pidField.GetText(),
			Port: portField.GetText(),
			Proto: proto,
			LocalAddress: localField.GetText(),
			ForeignAddress: foreignField.GetText(),
			State: state,
		}
		row, ok := FilterLogic(data, filter)
		data.mu.Unlock()
		if ok {
			table.Select(row, 0)
		} else {
			fmt.Printf(" unable to find row\n")
		}
	})
	// this button clears the filter
	form.AddButton("Clear", func() {
		pidField.SetText("")
		portField.SetText("")
		localField.SetText("")
		foreignField.SetText("")
		protoDrop.SetCurrentOption(0)
		stateDrop.SetCurrentOption(0)
	})
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
	pid := form.GetFormItemByLabel("pid").(*tview.InputField)
	return pid.GetText()
}

// returns the current selected option by its index and the port type name
func GetPortInput(form *tview.Form) (int, string) {
	port := form.GetFormItemByLabel("port").(*tview.DropDown) // converts to tview state 
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

/*func GetFilterState(form *tview.Form) {
	filter := form.GetFormItemByLabel("filter").(tview.Button)
}*/
