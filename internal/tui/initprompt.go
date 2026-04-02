package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/table"
	"charm.land/bubbles/v2/textarea"
	"charm.land/lipgloss/v2"

	zone "github.com/lrstanley/bubblezone/v2"

)

type model struct { 
	table 		table.Model
	rows		[]table.Row
	searchBar 	textarea.Model 
	deleteBar 	textarea.Model
}

func (m model) Init() tea.Cmd { 
	return textarea.Blink  
}

func newModel(rows []table.Row) model {
	columns := []table.Column{
		{Title: "COMMAND", Width: 15},
		{Title: "PID", Width: 10},
		{Title: "USER", Width: 15},
		{Title: "PORT", Width: 25},
	}
	// the table image
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)
	t.SetWidth(70)  
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
	search.SetWidth(50)
	search.SetHeight(0)
	search.CharLimit = 50
	search.ShowLineNumbers = false
	search.Focus()

	// Text area for delete bar
	delete := textarea.New()
	delete.Placeholder = "Kill port or pid"
	delete.SetWidth(50)
	delete.SetHeight(0)
	delete.CharLimit = 50
	delete.ShowLineNumbers = false 
	
	return model{
		table:     	t,
		searchBar:  search,
		deleteBar: 	delete,
		rows:      	rows,
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {	// where "input" goes msg.String()
		case "q", "ctrl+c":
			return m, tea.Quit
		
		// reseting the table state after pressing enter with no search 
		case "enter":
			query := strings.ToLower(strings.TrimSpace(m.searchBar.Value()))
			var filtered []table.Row
			if query == "" {
				filtered = m.rows
			} else {
				for _, r := range m.rows {
					portCol := strings.ToLower(r[3])
					commandName := strings.ToLower(r[0])
					if strings.Contains(portCol, query) || strings.Contains(commandName, query) {
						filtered = append(filtered, r)
					}
				}
			}
			m.table.SetRows(filtered)
			return m, nil
		}

	case tea.MouseReleaseMsg:
		if msg.Button != tea.MouseLeft {
			return m, nil
		}
		// zones
		switch {
		case zone.Get("search").InBounds(msg):
			m.searchBar.Focus()
			m.deleteBar.Blur()
			

		case zone.Get("delete").InBounds(msg):
			m.deleteBar.Focus()
			m.searchBar.Blur()

		}
	}
	
	// Update search bar and delete bar
	var cmd tea.Cmd
	if m.searchBar.Focused() {
		m.searchBar, cmd = m.searchBar.Update(msg)
		cmds = append(cmds, cmd)
	} else if m.deleteBar.Focused() {
		m.deleteBar, cmd = m.deleteBar.Update(msg)
		cmds = append(cmds, cmd)
	}

	// handle filtered rows/items
	filtered := []table.Row{}
	query := strings.ToLower(strings.TrimSpace(m.searchBar.Value()))
	if query == "" { 	// either nothing in the search bar or port does not exist
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

	// killing a port
	// rmber that ports owned by root requires sudo access. 
	// ...

	m.table.SetRows(filtered)
	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}


func (m model) View() tea.View {
    tableView := m.table.View()

	// wrapping marked zones: searchbar and deletebar
	searchView := 
		zone.Mark("search", searchStyle.Render("Search: "+m.searchBar.View()),
	)
	deleteView := 
		zone.Mark("delete", delStyle.Render("Delete: "+m.deleteBar.View()),
	)

	// render components for table
    renderTable := fmt.Sprintf(
        "%s\n\n%s\n\n%s\n\n%s\n%s",
        titleStyle.Render("Ports in listen..."),
		searchView,
		deleteView,
        tableView, 
        footStyle.Render(
			"Controls:\nup/down: navigate | q: quit | enter: more info"),
    )
    var v tea.View
	// bubblezone requires altscreen
	v.AltScreen = true				// *** not sure if i will keep this 
	v.MouseMode = tea.MouseModeCellMotion
    v.SetContent(zone.Scan(baseStyle.Render(renderTable)))
    return v
}

func StartTable(rows []table.Row) error {
	p := tea.NewProgram(newModel(rows))
	_, err := p.Run() 
	return err
}

