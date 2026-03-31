package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/table"
	"charm.land/bubbles/v2/textarea"
	"charm.land/lipgloss/v2"
)

type model struct { 
	table 		table.Model
	rows		[]table.Row
	searchBar 	textarea.Model 
}

func (m model) Init() tea.Cmd { 
	return textarea.Blink  
}

// table styling 
var (
	baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("#424549")).
	Padding(0, 1).
	Margin(1, 2)
)
// title style
var titleStyle = 
	lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#64ff71"))

// search style
var searchStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#507fcb")).
	Bold(true)

// footer style
var footStyle = 
	lipgloss.NewStyle().
	Foreground(lipgloss.Color("#34383e"))

// selection style
var selectStyle = 
	lipgloss.NewStyle().
	Foreground(lipgloss.Color("#ffffff")).
	Background(lipgloss.Color("#282b30")).
	Bold(true)

func newModel(rows []table.Row) model {
	columns := []table.Column{
		{Title: "COMMAND", Width: 15},
		{Title: "PID", Width: 10},
		{Title: "USER", Width: 15},
		{Title: "PORT", Width: 25},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	t.SetWidth(80)  
	t.SetHeight(15)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)

	s.Selected = selectStyle
	t.SetStyles(s)
	// Textarea for search bar
	search := textarea.New()
	search.Placeholder = "Search port"
	search.Focus()
	search.SetWidth(50)
	search.SetHeight(0)
	search.ShowLineNumbers = false 
	
	return model{
		table:     t,
		searchBar: search,
		rows:      rows,
	}
}


// --- Update ---
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	// Update search bar
	var cmd tea.Cmd
	m.searchBar, cmd = m.searchBar.Update(msg)
	cmds = append(cmds, cmd)

	// handle filtered rows/items
	filtered := []table.Row{}
	query := strings.ToLower(strings.TrimSpace(m.searchBar.Value()))
	// either nothing in the search bar or port does not exist
	if query == "" {
		filtered = m.rows
	} else {
		for _, r := range m.rows {
			portCol := strings.ToLower(r[3])		// port num
			commandName := strings.ToLower(r[0])	// process name
			if strings.Contains(portCol, query) || strings.Contains(commandName, query){
				filtered = append(filtered, r)
			}
		}
	}
	m.table.SetRows(filtered)
	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}


func (m model) View() tea.View {
    tableView := m.table.View()
	searchView := searchStyle.Render("Search: " + m.searchBar.View())

	// table render
    renderTable := fmt.Sprintf(
        "%s\n\n%s\n\n%s\n\n%s",
        titleStyle.Render("Ports in listen..."),
		searchView,
        tableView, // leave untouched
        footStyle.Render("up/down: navigate | q: quit | enter: more info"),
    )
    var v tea.View
    v.SetContent(baseStyle.Render(renderTable))
    return v
}

func StartTable(rows []table.Row) error {
	p := tea.NewProgram(newModel(rows))
	_, err := p.Run() 
	return err
}

