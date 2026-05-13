package tui

import (
	"fmt"
	"strconv"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
)

// temporary vars
var (
	pid string
	portNumber string
	proto string 
	connectionName string 
	laddr string 
	faddr string
	state string
)

type formModel struct {
	form *huh.Form
}

// this filter form model describes the text form to filter the lsof call. 
func initFormModel() formModel {
	form := huh.NewForm(
		huh.NewGroup (
			huh.NewInput().Title("pid number").Prompt("?").Validate(isNumber).Value(&pid),
			huh.NewInput().Title("port number").Prompt("?").Validate(isNumber).Value(&portNumber),
			huh.NewSelect[string]().Title("proto").Options(
				huh.NewOption("tcp", "tcp"),
				huh.NewOption("udp", "udp"),
				huh.NewOption("icmp", "icmp"),
				huh.NewOption("ip", "ip"),
				huh.NewOption("unix", "unix"),
				huh.NewOption("socket", "socket"),
			).Value(&proto),
			huh.NewInput().Title("process name").Prompt("?").Value(&connectionName),
			huh.NewInput().Title("local address").Prompt("?").Validate(isEmpty).Value(&laddr),
			huh.NewInput().Title("foreign address").Prompt("?").Validate(isEmpty).Value(&faddr),
			huh.NewSelect[string]().Title("state").Options(
				huh.NewOption("listening", "listening"),
				huh.NewOption("established", "established"),
			).Value(&state),
			huh.NewConfirm().Key("done").Title("Filter?").Affirmative("Yes").Negative("No"),
		),
	).WithWidth(50).WithShowHelp(false).WithShowErrors(false)
	return formModel{form: form}
}

// init the form model
func (m formModel) Init() tea.Cmd { return m.form.Init() }

func (m formModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	var cmds []tea.Cmd
	// process form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}
	// check if form is completed
	if m.form.State == huh.StateCompleted {
		// run the process 
		// for now just quit the form when it is done
		//cmds = append(cmds, tea.Quit)
		// make sure to append and switch back to the main menu
		// or reset back to form
		m = initFormModel()
	}
	return m, tea.Batch(cmds...)
}

func (m formModel) View() tea.View { return tea.NewView(m.form.View()) }

// validates if the user entered a number otherwise leave it be 
func isNumber(s string) error {
	s = strings.TrimSpace(s)
	// empty
	if s == "" {
		return nil
	}
	if _, err := strconv.Atoi(s); err != nil {
		return fmt.Errorf("must be a number")
	}
	return nil
}

func isEmpty(s string) error {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	} 
	return nil
}