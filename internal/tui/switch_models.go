package tui

import (
	tea "charm.land/bubbletea/v2"
)

type switchToModel struct {
	name string
}

func SwitchModelCmd(name string) tea.Cmd {
	return func() tea.Msg {
		return switchToModel{name: name}
	}
}