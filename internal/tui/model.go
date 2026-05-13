package tui

import (
	tea "charm.land/bubbletea/v2"
)

// this describes the main model for the whole application 
type mainModel struct {
	models map[string]tea.Model
	selectedModel string
}

func initMainModel() mainModel {
	return mainModel {
		models: map[string]tea.Model {
			"form": initFormModel(), 
			// dashboard model
		},
		// default selected model
		selectedModel: "form",
	}
}
func (m mainModel) Init() tea.Cmd { return m.models[m.selectedModel].Init() }

func (m *mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
			// quit program
		case "ctrl+c":
			return m, tea.Quit
		}
		// switch the current model
	case switchToModel:
		m.selectedModel = msg.name
		return m, m.models[msg.name].Init()
	}
	currentModel := m.models[m.selectedModel]
	updated, cmd := currentModel.Update(msg)
	m.models[m.selectedModel] = updated
	return m, cmd
}

func (m mainModel) View() tea.View { return m.models[m.selectedModel].View() }

func Run() error {
	model := initMainModel()
	p := tea.NewProgram(&model)
	_, err := p.Run()
	return err
}